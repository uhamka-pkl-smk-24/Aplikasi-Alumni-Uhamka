package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jeypc/go-auth/go-auth/config"
	"github.com/jeypc/go-auth/go-auth/entities"
	"github.com/jeypc/go-auth/go-auth/model"
	"golang.org/x/crypto/bcrypt"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}

	if len(session.Values) == 0 || session.Values["loggedIn"] != true || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	id, ok := session.Values["id"].(int)
	if !ok {
		http.Error(w, "User ID not found in session", http.StatusInternalServerError)
		return
	}

	lamaranList, err := userModel.GetAllLamaran()
	if err != nil {
		http.Error(w, "Unable to fetch lamaran data", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"username": session.Values["username"],
		"lamaran":  lamaranList,
	}

	temp, err := template.ParseFiles("go-auth/views/html/dashboard_Admin.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}
	if err := temp.Execute(w, data); err != nil {
		log.Println("error : ", err.Error())
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}

	lamaran, err := model.NewUserModel().GetNewNotifications(id)
	if err != nil {
		http.Error(w, "Unable to fetch notifications", http.StatusInternalServerError)
		return
	}

	data = map[string]interface{}{
		"username": session.Values["username"],
		"lamaran":  lamaran,
	}

	tmpl, err := template.ParseFiles("go-auth/views/html/dashboard_Admin.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}

func AdminLowonganDashboard(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}

	if session.Values["role"] != "admin_lowongan" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Instantiate UserModel and get all lowongan
	lowonganList, err := userModel.GetAllLowongan()
	if err != nil {
		http.Error(w, "Failed to fetch job postings", http.StatusInternalServerError)
		return
	}

	// Siapkan data yang akan dikirim ke template
	data := struct {
		LowonganList []*entities.Lowongan
	}{
		LowonganList: lowonganList,
	}

	// Render admin lowongan dashboard
	tmpl, err := template.ParseFiles("go-auth/views/html/adminLowongan.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("data :", err)
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}

func TambahPekerjaan(w http.ResponseWriter, r *http.Request) {
	// Mengambil sesi dan memastikan user adalah mahasiswa
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil || session.Values["role"] != "mahasiswa" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	// Menangani permintaan POST untuk menambahkan pekerjaan
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20) // Batas ukuran file 10 MB
		if err != nil {
			log.Println("Error parsing form data:", err)
			http.Error(w, "Unable to process form data", http.StatusInternalServerError)
			return
		}

		alias := r.FormValue("alias")

		file, handler, err := r.FormFile("surat_lamaran")
		if err != nil {
			log.Println("Error reading file:", err)
			http.Error(w, "Unable to read file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("directory", dir)

		filename := handler.Filename
		if alias != "" {
			filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
		}

		fileLocation := filepath.Join(dir, "files", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		nim := fmt.Sprintf("%v", session.Values["nim"])

		nimInt, _ := strconv.Atoi(nim)

		// Membuat objek Lamaran dan mengisi data dari form
		pekerjaan := &entities.Lamaran{
			NamaPekerjaan: r.FormValue("nama_pekerjaan"),
			Perusahaan:    r.FormValue("perusahaan"),
			Nim:           nimInt,
			SuratLamaran:  filename,
		}

		err = userModel.TambahPekerjaan(pekerjaan)
		if err != nil {
			log.Println("Error adding job data:", err)
			http.Redirect(w, r, "/error?msg=Gagal menambahkan data pekerjaan", http.StatusSeeOther)
			return
		}

		if _, err := io.Copy(targetFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect ke halaman mahasiswa dengan pesan sukses
		http.Redirect(w, r, "/mahasiswa?msg=Sukses menambahkan data pekerjaan", http.StatusSeeOther)
		return
	}

	// Menampilkan halaman dashboard mahasiswa jika metode GET
	tmpl, err := template.ParseFiles("go-auth/views/html/dashboard_Mhs.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func TambahData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		session, err := config.Store.Get(r, config.SESSION_ID)
		if err != nil {
			http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
			return
		}

		if len(session.Values) == 0 || session.Values["loggedIn"] != true || session.Values["role"] != "admin" {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}

		data := map[string]interface{}{
			"username": session.Values["username"],
		}

		// Parse dan eksekusi template
		tmpl, err := template.ParseFiles("go-auth/views/html/tambahData.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}

	} else if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20) // Batas ukuran file 10 MB
		if err != nil {
			log.Println("Error parsing form data:", err)
			http.Error(w, "Unable to process form data", http.StatusInternalServerError)
			return
		}

		// insert photo to storage
		alias := r.FormValue("alias")

		file, handler, err := r.FormFile("photo")
		if err != nil {
			log.Println("Error reading file:", err)
			http.Error(w, "Unable to read file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("directory", dir)

		filename := handler.Filename
		if alias != "" {
			filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
		}

		fileLocation := filepath.Join(dir, "photo", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		// Copy the contents of the uploaded file to the target file
		_, err = io.Copy(targetFile, file)
		if err != nil {
			log.Println("Error copying file:", err)
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		// Konversi NIM dari string ke int64
		nimStr := r.Form.Get("nim")
		log.Println(nimStr)
		nim, err := strconv.ParseInt(nimStr, 10, 64)
		if err != nil {
			http.Error(w, "NIM harus berupa angka", http.StatusBadRequest)
			return
		}

		// Konversi IPK dari string ke float64
		ipkStr := r.Form.Get("ipk")
		ipk, err := strconv.ParseFloat(ipkStr, 64)
		if err != nil {
			http.Error(w, "IPK harus berupa angka desimal", http.StatusBadRequest)
			return
		}

		// Hash password menggunakan MD5
		password := r.Form.Get("password")
		hashPasword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		passwordDecr := string(hashPasword)

		// Buat objek Mahasiswa dengan nilai yang telah dikonversi
		mahasiswa := &entities.Mahasiswa{
			NIM:          nim,
			NoIjazah:     r.FormValue("noIjazah"),
			NamaLengkap:  r.FormValue("namaLengkap"),
			TempatLahir:  r.FormValue("tempatLahir"),
			TanggalLahir: r.FormValue("tanggalLahir"),
			Agama:        r.FormValue("agama"),
			Alamat:       r.FormValue("alamat"),
			NoTelp:       r.FormValue("noTelp"),
			Email:        r.FormValue("email"),
			IPK:          ipk, // Sudah dikonversi ke float
			BidangStudi:  r.FormValue("bidangStudi"),
			Photo:        filename,
			Angkatan:     r.FormValue("angkatan"),
			Tahunlulus:   r.FormValue("tahunlulus"),
		}

		// Buat objek User untuk data user
		user := &entities.User{
			Nim:      nim,
			Username: r.FormValue("username"),
			Password: passwordDecr, // Simpan password yang telah di-hash
			Role:     r.FormValue("role"),
			No_telp:  r.FormValue("noTelp"),
			Email:    r.FormValue("email"),
		}

		err = userModel.TambahMahasiswa(mahasiswa, user)
		if err != nil {
			log.Println("Error adding student data:", err)
			http.Error(w, "Unable to add student data", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
func UpdateApprovalStatus(w http.ResponseWriter, r *http.Request) {

	log.Println("Received request for updating approval status") // Logging untuk debug

	// Validasi request method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Menguraikan form data dari request body
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Mendapatkan nilai 'id' dan 'isApproved' dari form data
	idStr := r.FormValue("id")
	isApprovedStr := r.FormValue("isApproved")

	// Konversi 'id' menjadi integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Konversi 'isApproved' menjadi boolean
	isApprovedInt, err := strconv.Atoi(isApprovedStr)
	if err != nil {
		http.Error(w, "Invalid approval status", http.StatusBadRequest)
		return
	}
	isApproved := isApprovedInt == 1

	log.Println("id :", id)
	log.Println("isApproved :", isApproved)

	// Memanggil fungsi model untuk mengupdate status approval
	err = userModel.UpdateApprovalStatus(id, isApproved)
	if err != nil {
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}

	// Jika sukses, kirimkan respon sukses
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status updated successfully"))
}

func DownloadPekerjaan(w http.ResponseWriter, r *http.Request) {

	// Mendapatkan nama file dari query string
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "File not specified", http.StatusBadRequest)
		return
	}

	// Mencari file di direktori "files"
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Unable to get working directory", http.StatusInternalServerError)
		return
	}
	fileLocation := filepath.Join(dir, "files", filename)

	// Memastikan file ada
	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Menyetel header untuk mendownload file
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", getFileSize(fileLocation)))

	// Mengirim file ke response writer
	http.ServeFile(w, r, fileLocation)
}

func getFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func ImportMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	var rawData interface{}

	// Menerima data mentah dari body untuk debugging
	err := json.NewDecoder(r.Body).Decode(&rawData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Invalid input:", err)
		return
	}

	// Log the raw data received for inspection
	rawBytes, _ := json.Marshal(rawData)
	log.Println("Raw JSON received:", string(rawBytes))

	// Decode the raw data into slice of MahasiswaImport
	var mahasiswaList []entities.MahasiswaImport
	err = json.Unmarshal(rawBytes, &mahasiswaList)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusInternalServerError)
		log.Println("Failed to decode mahasiswaList:", err)
		return
	}

	// Log the decoded Mahasiswa list
	log.Printf("Decoded Mahasiswa List: %+v\n", mahasiswaList)

	// Simpan data mahasiswa ke database
	for _, mahasiswa := range mahasiswaList {
		log.Println("Tanggal Lahir:", mahasiswa.TanggalLahir)
		err := userModel.Create(&mahasiswa)
		if err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			log.Println("Failed to save data:", err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func ImportUserHandler(w http.ResponseWriter, r *http.Request) {
	var rawData interface{}

	// Menerima data mentah dari body untuk debugging
	err := json.NewDecoder(r.Body).Decode(&rawData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Invalid input:", err)
		return
	}

	// Cetak data mentah untuk melihat struktur yang diterima
	log.Printf("Received data: %+v\n", rawData)

	// Lakukan decoding ulang ke dalam slice UserImport
	var userList []entities.UserImport
	rawBytes, _ := json.Marshal(rawData) // Konversi kembali ke JSON untuk di-decode ulang
	log.Println("Raw JSON Data:", string(rawBytes))

	// Unmarshal JSON ke dalam struct userList
	err = json.Unmarshal(rawBytes, &userList)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusInternalServerError)
		log.Println("Failed to decode userList:", err)
		return
	}

	// Berikan response sukses
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data imported successfully"))
}
