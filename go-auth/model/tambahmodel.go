package model

import (
	"log"

	"github.com/jeypc/go-auth/go-auth/entities"
)

func (m *UserModel) TambahMahasiswa(mahasiswa *entities.Mahasiswa, user *entities.User) error {
	// Query untuk menyimpan data user
	query1 := `INSERT INTO user
        (nim, username, email, password, role, no_telp) 
        VALUES (?, ?, ?, ?, ?, ?)`

	// Eksekusi query untuk tabel user
	_, err := m.DB.Exec(query1, user.Nim, user.Username, user.Email, user.Password, user.Role, user.No_telp)
	if err != nil {
		log.Printf("Error adding user data: %v", err)
		return err
	}

	// Query untuk menyimpan data mahasiswa
	query2 := `INSERT INTO mahasiswa 
        (nim, no_ijazah, nama_lengkap, tempat_lahir, tanggal_lahir, agama, alamat, no_telp, email, ipk, bidang_studi, photo, angkatan, tahunlulus) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Eksekusi query untuk tabel mahasiswa
	_, err = m.DB.Exec(query2, mahasiswa.NIM, mahasiswa.NoIjazah, mahasiswa.NamaLengkap, mahasiswa.TempatLahir,
		mahasiswa.TanggalLahir, mahasiswa.Agama, mahasiswa.Alamat, mahasiswa.NoTelp, mahasiswa.Email, mahasiswa.IPK,
		mahasiswa.BidangStudi, mahasiswa.Photo, mahasiswa.Angkatan, mahasiswa.Tahunlulus)

	if err != nil {
		log.Printf("Error adding student data: %v", err)
		return err
	}

	log.Println("Successfully added student and user data")
	return nil
}

func (m *UserModel) TambahPekerjaan(lamaran *entities.Lamaran) error {
	query := `
        INSERT INTO lamaran (nama_pekerjaan, perusahaan, surat_lamaran, nim)
        VALUES (?, ?, ?, ?)
    `

	_, err := m.DB.Exec(query,
		lamaran.NamaPekerjaan, lamaran.Perusahaan,
		lamaran.SuratLamaran, lamaran.Nim,
	)
	if err != nil {
		log.Printf("Error adding student data: %v", err)
		return err
	}

	log.Println("Successfully")
	return nil
}
