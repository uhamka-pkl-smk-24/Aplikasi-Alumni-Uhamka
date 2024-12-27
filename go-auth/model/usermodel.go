package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jeypc/go-auth/go-auth/config"
	"github.com/jeypc/go-auth/go-auth/entities"
)

type UserModel struct {
	DB *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &UserModel{
		DB: conn,
	}
}

// Existing function for fetching user by a field (e.g., username)
func (m *UserModel) Where(user *entities.User, field1, value1 string) error {
	query := "SELECT * FROM user WHERE " + field1 + " = ? LIMIT 1"
	return m.DB.QueryRow(query, value1).Scan(
		&user.Nim, &user.Username, &user.Email, &user.Password,
		&user.Role, &user.No_telp,
	)
}

// New function to get student data by NIM
func (m *UserModel) GetMahasiswaByNIM(nim string) (*entities.Mahasiswa, error) {
	mahasiswa := &entities.Mahasiswa{}
	query := `
        SELECT 
            nim, no_ijazah, nama_lengkap, tempat_lahir, tanggal_lahir, agama, 
            alamat, no_telp, email, ipk, photo, bidang_studi,photo, angkatan, tahun_lulus 
        FROM mahasiswa 
        WHERE nim = ?`

	log.Printf("Executing query: %s with NIM: %s", query, nim)

	err := m.DB.QueryRow(query, nim).Scan(
		&mahasiswa.NIM, &mahasiswa.NoIjazah, &mahasiswa.NamaLengkap,
		&mahasiswa.TempatLahir, &mahasiswa.TanggalLahir, &mahasiswa.Agama,
		&mahasiswa.Alamat, &mahasiswa.NoTelp, &mahasiswa.Email, &mahasiswa.IPK, &mahasiswa.Photo,
		&mahasiswa.BidangStudi, &mahasiswa.Photo, &mahasiswa.Angkatan, &mahasiswa.Tahunlulus)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No student found with the given NIM: %s", nim)
			return nil, errors.New("no student found with the given NIM")
		}
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	log.Printf("Fetched student data: %+v", mahasiswa)

	return mahasiswa, nil
}

func (m *UserModel) EditMahasiswa(mahasiswa *entities.Mahasiswa, user *entities.User) error {
	// Start a transaction
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	// Update the user table
	query1 := `UPDATE user SET 
		email = ?, 
		no_telp = ?  
		WHERE nim = ?`

	_, err = tx.Exec(query1, user.Email, user.No_telp, mahasiswa.NIM)
	if err != nil {
		tx.Rollback()
		log.Printf("Error updating user data: %v", err)
		return err
	}

	// Update the mahasiswa table
	query2 := `UPDATE mahasiswa SET 
		nama_lengkap = ?, 
		tempat_lahir = ?, 
		tanggal_lahir = ?, 
		agama = ?, 
		alamat = ?, 
		no_telp = ?, 
		email = ?, 
		bidang_studi = ?,
		photo = ?,
		angkatan = ?,
		tahun_lulus = ?,
		WHERE nim = ?`

	_, err = tx.Exec(query2, mahasiswa.NamaLengkap, mahasiswa.TempatLahir,
		mahasiswa.TanggalLahir, mahasiswa.Agama, mahasiswa.Alamat, mahasiswa.NoTelp, mahasiswa.Email,
		mahasiswa.BidangStudi, mahasiswa.Photo, mahasiswa.Angkatan, mahasiswa.Tahunlulus, mahasiswa.NIM)
	if err != nil {
		tx.Rollback()
		log.Printf("Error updating mahasiswa data: %v", err)
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	log.Println("Successfully updated student and user data")
	return nil
}

func (m *UserModel) GetAllMahasiswa() ([]*entities.Mahasiswa, error) {
	rows, err := m.DB.Query("SELECT nim, no_ijazah, nama_lengkap, tempat_lahir, tanggal_lahir, agama, alamat, no_telp, email, ipk, bidang_studi, photo, angkatan, tahun_lulus FROM mahasiswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mahasiswaList []*entities.Mahasiswa

	for rows.Next() {
		// Inisialisasi objek Mahasiswa baru
		mahasiswa := &entities.Mahasiswa{}

		// Pindai data dari hasil query ke dalam objek Mahasiswa
		if err := rows.Scan(
			&mahasiswa.NIM,
			&mahasiswa.NoIjazah,
			&mahasiswa.NamaLengkap,
			&mahasiswa.TempatLahir,
			&mahasiswa.TanggalLahir,
			&mahasiswa.Agama,
			&mahasiswa.Alamat,
			&mahasiswa.NoTelp,
			&mahasiswa.Email,
			&mahasiswa.IPK,
			&mahasiswa.BidangStudi,
			&mahasiswa.Photo,
			&mahasiswa.Angkatan,
			&mahasiswa.Tahunlulus,
		); err != nil {
			return nil, err
		}

		// Tambahkan objek Mahasiswa ke dalam slice mahasiswaList
		mahasiswaList = append(mahasiswaList, mahasiswa)
	}

	return mahasiswaList, nil
}

func (m *UserModel) FilterMahasiswa(nim, noIjazah, nama string) ([]entities.Mahasiswa, error) {
	var mahasiswaList []entities.Mahasiswa

	// Membangun query SQL
	query := "SELECT nim, no_ijazah, nama_lengkap, tempat_lahir, tanggal_lahir, agama, alamat, no_telp, email, ipk, bidang_studi, angkatan, tahun_lulus FROM mahasiswa "
	// var args string

	if nim != "" {
		query += " ORDER BY nim"
	}
	if noIjazah != "" {
		query += " ORDER BY no_ijazah"
	}
	if nama != "" {
		query += " ORDER BY nama_lengkap"
	}

	log.Println(m.DB.Query(query))

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	log.Println(query)

	for rows.Next() {
		var mhs entities.Mahasiswa
		if err := rows.Scan(
			&mhs.NIM, &mhs.NoIjazah, &mhs.NamaLengkap, &mhs.TempatLahir,
			&mhs.TanggalLahir, &mhs.Agama, &mhs.Alamat, &mhs.NoTelp,
			&mhs.Email, &mhs.IPK, &mhs.BidangStudi, &mhs.Angkatan, &mhs.Tahunlulus,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		mahasiswaList = append(mahasiswaList, mhs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return mahasiswaList, nil
}

func (m *UserModel) ImageFileNameByNim(nim string) (string, error) {
	var filename string
	query := "SELECT photo FROM mahasiswa WHERE nim = ?"
	err := m.DB.QueryRow(query, nim).Scan(&filename)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Jika tidak ada baris yang ditemukan, kembalikan string kosong
		}
		return "", fmt.Errorf("error fetching image filename: %w", err)
	}
	return filename, nil
}

func (m *UserModel) Create(mahasiswa *entities.MahasiswaImport) error {
	// Membuat koneksi ke database dengan menggunakan fungsi DBConn dari package config
	db, err := config.DBConn()
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return err
	}
	defer db.Close()

	// Memulai transaksi
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return err
	}

	// Query untuk memasukkan data ke tabel user, dengan password default '123' dan role 'mahasiswa'
	queryUser := `
		INSERT INTO user (nim, username, email, password, role, no_telp) 
		VALUES (?, ?, ?, ?, ?, ?)
	`

	// Password default '123' dan role 'mahasiswa'
	defaultPassword := "$2a$12$8.PbJh49qw8FBMHvX7mvbe1vEMDn3hUq8BqubzPT/w49XSV3EFd9m" // PASSWORD DEFAULT "123"
	defaultRole := "mahasiswa"

	// Menjalankan query untuk memasukkan data ke tabel user
	_, err = tx.Exec(queryUser, mahasiswa.NIM, mahasiswa.NamaLengkap, mahasiswa.Email, defaultPassword, defaultRole, mahasiswa.NoTelp)
	if err != nil {
		log.Println("Failed to insert user data:", err)
		tx.Rollback() // Batalkan transaksi jika terjadi kesalahan
		return err
	}

	// Query untuk memasukkan data mahasiswa
	queryMahasiswa := `
		INSERT INTO mahasiswa (nim, no_ijazah, nama_lengkap, tempat_lahir, tanggal_lahir, agama, alamat, no_telp, email, ipk, bidang_studi, angkatan, tahun_lulus) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Menjalankan query untuk memasukkan data ke tabel mahasiswa
	_, err = tx.Exec(queryMahasiswa, mahasiswa.NIM, mahasiswa.NoIjazah, mahasiswa.NamaLengkap, mahasiswa.TempatLahir,
		mahasiswa.TanggalLahir, mahasiswa.Agama, mahasiswa.Alamat, mahasiswa.NoTelp, mahasiswa.Email, mahasiswa.IPK,
		mahasiswa.BidangStudi, mahasiswa.Angkatan, mahasiswa.Tahunlulus)
	if err != nil {
		log.Println("Failed to insert mahasiswa data:", err)
		tx.Rollback() // Batalkan transaksi jika terjadi kesalahan
		return err
	}

	// Commit transaksi jika kedua operasi berhasil
	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction:", err)
		return err
	}

	log.Println("Mahasiswa and user data successfully saved")
	return nil
}
