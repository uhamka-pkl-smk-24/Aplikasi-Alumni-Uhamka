package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/go-auth/go-auth/controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("GET", "POST")
	http.HandleFunc("/jobs", controller.LowonganListController)
	r.HandleFunc("/session-info", controller.SessionInfoHandler).Methods("GET")

	// r.HandleFunc("/login", controller.LoginHandler).Methods(http.MethodPost, http.MethodPost)
	r.HandleFunc("/logout", controller.Logout).Methods("GET")
	// r.HandleFunc("/profile", controller.MahasiswaProfile).Methods("GET")
	http.HandleFunc("/adminlowongan", controller.AdminLowonganDashboard)
	http.HandleFunc("/admin", controller.AdminDashboard)
	http.HandleFunc("/mahasiswa", controller.MahasiswaDashboard)
	r.HandleFunc("/foto", controller.ImageHandler).Methods("GET")
	r.HandleFunc("/filter-mahasiswa", controller.FilterMahasiswa).Methods(http.MethodGet)
	r.HandleFunc("/tambah-data", controller.TambahData).Methods("GET")
	r.HandleFunc("/edit-profile/{nim}", controller.EditMahasiswa).Methods("GET")
	r.HandleFunc("/edit-lowongan/{id}", controller.EditLowonganController).Methods("GET")
	r.HandleFunc("/delete/{id}", controller.DeleteLowonganHandler).Methods("DELETE")

	r.HandleFunc("/tambahdata", controller.TambahData).Methods("POST")
	r.HandleFunc("/pasang-lowongan", controller.TambahLowonganController).Methods(http.MethodGet, http.MethodPost)

	r.HandleFunc("/edit-mahasiswa/{nim}", controller.EditMahasiswa).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/edit-lowongan/{id}", controller.EditLowonganController).Methods(http.MethodGet, http.MethodPost)
	http.Handle("/", r)

	// Admin route
	r.HandleFunc("/get-lamaran-data", controller.AdminDashboard).Methods("GET")
	r.HandleFunc("/import-mahasiswa", controller.ImportMahasiswaHandler).Methods("POST")
	r.HandleFunc("/import-user", controller.ImportUserHandler).Methods("POST")

	r.HandleFunc("/update-approval-status", controller.UpdateApprovalStatus).Methods("POST")
	// http.HandleFunc("/update-approval-status", controller.UpdateApprovalStatus(db))

	r.HandleFunc("/admin", controller.DataMahasiswa).Methods("GET")

	r.HandleFunc("/admin", controller.AdminDashboard).Methods("GET")
	r.HandleFunc("/adminlowongan", controller.AdminLowonganDashboard).Methods("GET")

	// Mahasiswa route
	// r.HandleFunc("/notifications", controller.NotificationController).Method("GET")
	r.HandleFunc("/mahasiswa", controller.MahasiswaDashboard).Methods("GET")

	r.HandleFunc("/upload-lamaran", controller.TambahPekerjaan).Methods("POST")

	r.HandleFunc("/download", controller.DownloadPekerjaan).Methods("GET")

	// Serve files in the "files" directory at the "/files/" URL path
	// Define a route for serving static files (like CSS, JS, etc.)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("run in http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
