package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jeypc/go-auth/go-auth/config"
	"github.com/jeypc/go-auth/go-auth/entities"
	"github.com/jeypc/go-auth/go-auth/model"
)

func MahasiswaDashboard(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}

	if len(session.Values) == 0 || session.Values["loggedIn"] != true || session.Values["role"] != "mahasiswa" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	nim := fmt.Sprintf("%v", session.Values["nim"])
	log.Println("NIM SESSION ", nim)

	// Get student data from the database
	mahasiswa, err := model.NewUserModel().GetMahasiswaByNIM(nim)
	if err != nil {
		http.Error(w, "Unable to fetch student data", http.StatusInternalServerError)
		return
	}

	// Log for checking student data
	log.Printf("Mahasiswa data: %+v\n", mahasiswa)

	lamarans, err := model.NewUserModel().GetNotifications(nim)
	if err != nil {
		http.Error(w, "Unable to fetch notifications", http.StatusInternalServerError)
		return
	}

	log.Println("lamaran", lamarans)

	// // Mencari file di direktori "files"
	// dir, err := os.Getwd()
	// if err != nil {
	// 	http.Error(w, "Unable to get working directory", http.StatusInternalServerError)
	// 	return
	// }
	// fileLocation := filepath.Join(dir, "photo", mahasiswa.Photo)
	// log.Println("file", fileLocation)

	// // Memastikan file ada
	// if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
	// 	http.Error(w, "File not found", http.StatusNotFound)
	// 	return
	// }

	// w.Header().Set("Content-Type", "image/jpeg") // Adjust based on your image type
	// http.ServeFile(w, r, fileLocation)

	// Data to be passed to the template
	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
		"username":  session.Values["username"],
		"Lamarans":  lamarans,
	}

	log.Println("data", data)

	// Parsing and executing the template
	temp, err := template.ParseFiles("go-auth/views/html/dashboard_Mhs.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}
	if err := temp.Execute(w, data); err != nil {
		log.Println("error", err.Error())
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}

func EditMahasiswa(w http.ResponseWriter, r *http.Request) {

	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}

	if len(session.Values) == 0 || session.Values["loggedIn"] != true || session.Values["role"] != "mahasiswa" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	nim := fmt.Sprintf("%v", session.Values["nim"])

	if r.Method == http.MethodGet {

		log.Println("NIM SESSION ", nim)

		// vars := mux.Vars(r)
		// nim := vars["nim"]

		if nim == "" {
			http.Error(w, "NIM tidak diberikan", http.StatusBadRequest)
			return
		}

		mahasiswa, err := model.NewUserModel().GetMahasiswaByNIM(nim)
		if err != nil {
			http.Error(w, "Unable to fetch student data", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"mahasiswa": mahasiswa,
		}

		tmpl, err := template.ParseFiles("go-auth/views/html/edit-profile.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		nimInt, err := strconv.ParseInt(nim, 10, 64)
		if err != nil {
			http.Error(w, "NIM harus berupa angka", http.StatusBadRequest)
			return
		}

		// ipkStr := r.Form.Get("ipk")
		// ipk, err := strconv.ParseFloat(ipkStr, 64)
		// if err != nil {
		// 	http.Error(w, "IPK harus berupa angka desimal", http.StatusBadRequest)
		// 	return
		// }

		mahasiswa := &entities.Mahasiswa{
			NIM: nimInt,
			// NoIjazah:     r.Form.Get("no-ijazah"),
			NamaLengkap:  r.Form.Get("nama-lengkap"),
			TempatLahir:  r.Form.Get("tempat-lahir"),
			TanggalLahir: r.Form.Get("tanggal-lahir"),
			Agama:        r.Form.Get("agama"),
			Alamat:       r.Form.Get("alamat"),
			NoTelp:       r.Form.Get("no-telp"),
			Email:        r.Form.Get("email"),
			// IPK:          ipk,
			BidangStudi: r.Form.Get("bidang-studi"),
			// Angkatan: r.From.Get("angkatan"),
			// Tahunlulus: r.From.Get("tahunlulus")
		}

		user := &entities.User{
			No_telp: r.Form.Get("no-elp"),
			Email:   r.Form.Get("email"),
		}

		err = model.NewUserModel().EditMahasiswa(mahasiswa, user)
		if err != nil {
			http.Error(w, "Unable to update student data", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func FilterMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Ambil sesi untuk validasi
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}

	// Validasi session
	if len(session.Values) == 0 || session.Values["loggedIn"] != true || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		// Ambil parameter filter dari query string
		NIM := r.URL.Query().Get("nim")
		NoIjazah := r.URL.Query().Get("no_ijazah")
		NamaLengkap := r.URL.Query().Get("nama")

		// Panggil fungsi model untuk mendapatkan data mahasiswa berdasarkan filter
		mahasiswaList, err := model.NewUserModel().FilterMahasiswa(NIM, NoIjazah, NamaLengkap)
		if err != nil {
			http.Error(w, "Unable to fetch student data", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"mahasiswa": mahasiswaList,
			"username":  session.Values["username"],
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
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DataMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Helper function to handle errors
	handleError := func(msg string, err error) {
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
	}

	// Helper function to parse and execute the template
	renderTemplate := func(dataMahasiswa interface{}, dataLamaran interface{}) {
		tmpl, err := template.ParseFiles("go-auth/views/html/dashboard_Admin.html")
		if err != nil {
			handleError("Error parsing template", err)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Mahasiswa": dataMahasiswa,
			"Lamaran":   dataLamaran,
		})
		if err != nil {
			handleError("Error executing template", err)
		}
	}

	// Ambil parameter filter dari query string
	NIM := r.URL.Query().Get("nim")
	NoIjazah := r.URL.Query().Get("no_ijazah")
	NamaLengkap := r.URL.Query().Get("nama")

	// Ambil data lamaran dari model
	dataLamaran, err := model.NewUserModel().GetAllLamaran()
	if err != nil {
		handleError("Error fetching data lamaran", err)
		return
	}

	var dataMahasiswa interface{}

	// Filter data mahasiswa berdasarkan parameter yang ada
	if NIM != "" || NoIjazah != "" || NamaLengkap != "" {
		dataMahasiswa, err = model.NewUserModel().FilterMahasiswa(NIM, NoIjazah, NamaLengkap)
		if err != nil {
			handleError("Unable to fetch student data", err)
			return
		}
	} else {
		// Ambil data mahasiswa tanpa filter
		dataMahasiswa, err = model.NewUserModel().GetAllMahasiswa()
		if err != nil {
			handleError("Error fetching data mahasiswa", err)
			return
		}
	}

	renderTemplate(dataMahasiswa, dataLamaran)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		log.Println("test")
		// Ambil sesi untuk validasi
		session, err := config.Store.Get(r, config.SESSION_ID)
		if err != nil {
			http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
			return
		}

		// Validasi session
		if len(session.Values) == 0 || session.Values["loggedIn"] != true || session.Values["role"] != "mahasiswa" {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}

		log.Println(session)

		log.Println("test")
		nim := fmt.Sprintf("%v", session.Values["nim"])
		if nim == "" {
			http.Error(w, "Missing NIM", http.StatusBadRequest)
			return

		}

		filename, err := model.NewUserModel().ImageFileNameByNim(nim)
		if err != nil {
			log.Println("Failed to retrieve image filename for NIM:", nim, "Error:", err)
			http.Error(w, "Failed to retrieve image filename", http.StatusInternalServerError)
			return
		}

		if filename == "" {
			log.Println("No image filename found for NIM:", nim)
			http.Error(w, "No image found for this user", http.StatusNotFound)
			return
		}

		// Construct the file path
		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, "Unable to get working directory", http.StatusInternalServerError)
			return
		}
		log.Println(dir)

		// Construct the file path
		filePath := filepath.Join(dir, "photo", filename)
		log.Println("Constructed file path:", filePath)

		// Check if the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Println("File not found:", filePath)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		log.Println("filepath", filePath)
		// Determine the content type
		contentType := "image/jpeg" // Default to JPEG
		switch filepath.Ext(filename) {
		case ".png":
			contentType = "image/png"
		case ".gif":
			contentType = "image/gif"
		}

		w.Header().Set("Content-Type", contentType)
		http.ServeFile(w, r, filePath)

	}
}
