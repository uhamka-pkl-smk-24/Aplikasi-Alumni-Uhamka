package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeypc/go-auth/go-auth/entities"
)

func TambahLowonganController(w http.ResponseWriter, r *http.Request) {
	// Handle GET request: Render the form page
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("go-auth/views/html/pasangLoker.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}
		return
	}

	// Handle POST request: Process form submission
	if r.Method == http.MethodPost {
		// Parse form values
		namaPekerjaan := r.FormValue("nama_pekerjaan")
		perusahaan := r.FormValue("perusahaan")
		lokasi := r.FormValue("lokasi")
		gajiStr := r.FormValue("gaji")
		deskripsi := r.FormValue("deskripsi")
		persyaratan := r.FormValue("syarat")

		// Convert ID and Gaji to integers

		gaji, err := strconv.Atoi(gajiStr)
		if err != nil {
			http.Error(w, "Invalid Gaji", http.StatusBadRequest)
			return
		}

		// Create a Lowongan struct
		lowongan := &entities.Lowongan{
			NamaPekerjaan: namaPekerjaan,
			Perusahaan:    perusahaan,
			Lokasi:        lokasi,
			Gaji:          gaji,
			Deskripsi:     deskripsi,
			Persyaratan:   persyaratan,
		}

		// Call the TambahLowongan method
		err = userModel.TambahLowongan(lowongan)
		if err != nil {
			http.Error(w, "Failed to add job posting", http.StatusInternalServerError)
			return
		}

		// Log success and redirect
		log.Println("Job posting successfully added")
		// Redirect ke halaman mahasiswa dengan pesan sukses
		http.Redirect(w, r, "/?msg=Sukses menambahkan lowongan", http.StatusSeeOther)
		return
	}

	// Handle unsupported methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func EditLowonganController(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the query parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	log.Println(id)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Handle GET request: Render the edit form
	if r.Method == http.MethodGet {
		lowongan, err := userModel.GetLowonganByID(id)
		if err != nil {
			http.Error(w, "Failed to fetch job posting", http.StatusInternalServerError)
			return
		}
		if lowongan == nil {
			http.Error(w, "Job posting not found", http.StatusNotFound)
			return
		}

		// Prepare data for the template
		data := struct {
			Lowongan *entities.Lowongan
		}{
			Lowongan: lowongan,
		}

		tmpl, err := template.ParseFiles("go-auth/views/html/editlowongan.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}

		log.Println(tmpl)
		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err.Error())
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}
		return
	}

	// Handle POST request: Process form submission
	if r.Method == http.MethodPost {
		// Parse form values
		namaPekerjaan := r.FormValue("namaPekerjaan")
		perusahaan := r.FormValue("perusahaan")
		lokasi := r.FormValue("lokasi")
		gajiStr := r.FormValue("gaji")
		deskripsi := r.FormValue("deskripsi")
		persyaratan := r.FormValue("persyaratan")

		// Convert Gaji to integer
		gaji, err := strconv.Atoi(gajiStr)
		if err != nil {
			http.Error(w, "Invalid Gaji", http.StatusBadRequest)
			return
		}

		// Create a Lowongan struct with ID
		lowongan := &entities.Lowongan{
			ID:            id, // Set ID for the update
			NamaPekerjaan: namaPekerjaan,
			Perusahaan:    perusahaan,
			Lokasi:        lokasi,
			Gaji:          gaji,
			Deskripsi:     deskripsi,
			Persyaratan:   persyaratan,
		}

		// Call the UpdateLowongan method
		err = userModel.UpdateLowongan(lowongan)
		if err != nil {
			http.Error(w, "Failed to update job posting", http.StatusInternalServerError)
			return
		}

		// Log success and redirect
		log.Println("Job posting successfully updated")
		http.Redirect(w, r, "/adminlowongan", http.StatusSeeOther)
		return
	}

	// Handle unsupported methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func LowonganListController(w http.ResponseWriter, r *http.Request) {
	// Ambil data lowongan dari database
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

	// Render template
	tmpl, err := template.ParseFiles("go-auth/views/html/index.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}

func DashboardMhsHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil data dari query string
	jobTitle := r.URL.Query().Get("nama_pekerjaan")
	companyName := r.URL.Query().Get("perusahaan")

	// Membuat data untuk diteruskan ke template HTML
	data := struct {
		JobTitle    string
		CompanyName string
	}{
		JobTitle:    jobTitle,
		CompanyName: companyName,
	}

	// Render halaman dashboard_mhs dengan data yang diperoleh
	tmpl := template.Must(template.ParseFiles("go-auth/views/html/dashboard_Mhs.html"))
	tmpl.Execute(w, data)
}

func DeleteLowonganHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Panggil metode untuk menghapus lowongan berdasarkan ID
	err = userModel.DeleteLowongan(id)
	if err != nil {
		http.Error(w, "Failed to delete job posting", http.StatusInternalServerError)
		return
	}

	// Redirect atau beri notifikasi sukses
	http.Redirect(w, r, "/adminlowongan?msg=Lowongan berhasil dihapus", http.StatusSeeOther)
}
