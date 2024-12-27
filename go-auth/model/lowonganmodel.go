package model

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jeypc/go-auth/go-auth/entities"
)

func (m *UserModel) TambahLowongan(lowongan *entities.Lowongan) error {
	query := `
        INSERT INTO lowongan (nama_pekerjaan, perusahaan, lokasi, gaji, deskripsi, syarat)
        VALUES (?, ?, ?, ?, ?, ?)
    `

	_, err := m.DB.Exec(query,
		lowongan.NamaPekerjaan, lowongan.Perusahaan,
		lowongan.Lokasi, lowongan.Gaji, lowongan.Deskripsi, lowongan.Persyaratan,
	)
	if err != nil {
		log.Printf("Error adding student data: %v", err)
		return err
	}

	log.Println("Successfully")
	return nil
}

func (m *UserModel) GetAllLowongan() ([]*entities.Lowongan, error) {
	query := "SELECT id, nama_pekerjaan, perusahaan, lokasi, gaji, deskripsi, syarat FROM lowongan"
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lowonganList []*entities.Lowongan
	for rows.Next() {
		lowongan := &entities.Lowongan{}
		err := rows.Scan(&lowongan.ID, &lowongan.NamaPekerjaan, &lowongan.Perusahaan, &lowongan.Lokasi, &lowongan.Gaji, &lowongan.Deskripsi, &lowongan.Persyaratan)
		if err != nil {
			return nil, err
		}
		lowonganList = append(lowonganList, lowongan)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lowonganList, nil
}

func (m *UserModel) GetLowonganByID(id int) (*entities.Lowongan, error) {
	lowongan := &entities.Lowongan{}

	// Query untuk mengambil data lowongan berdasarkan ID
	query := `SELECT id, nama_pekerjaan, perusahaan, lokasi, gaji, deskripsi, syarat
			  FROM lowongan WHERE id = ?`

	// Eksekusi query
	err := m.DB.QueryRow(query, id).Scan(
		&lowongan.ID,
		&lowongan.NamaPekerjaan,
		&lowongan.Perusahaan,
		&lowongan.Lokasi,
		&lowongan.Gaji,
		&lowongan.Deskripsi,
		&lowongan.Persyaratan,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no lowongan found with the given ID") // Tidak ditemukan
		}
		log.Printf("Error fetching lowongan by ID %d: %v", id, err)
		return nil, err // Kembalikan error
	}

	return lowongan, nil // Kembalikan data lowongan
}

func (m *UserModel) UpdateLowongan(lowongan *entities.Lowongan) error {
	// Start a transaction
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	// Update the lowongan table
	query := `
        UPDATE lowongan SET 
            nama_pekerjaan = ?, 
            perusahaan = ?, 
            lokasi = ?, 
            gaji = ?, 
            deskripsi = ?, 
            syarat = ? 
        WHERE id = ?`

	_, err = tx.Exec(query, lowongan.NamaPekerjaan, lowongan.Perusahaan,
		lowongan.Lokasi, lowongan.Gaji, lowongan.Deskripsi, lowongan.Persyaratan, lowongan.ID)
	if err != nil {
		tx.Rollback()
		log.Printf("Error updating job posting: %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	log.Println("Successfully updated job posting")
	return nil
}

func (m *UserModel) DeleteLowongan(id int) error {

	query := "DELETE FROM lowongan WHERE id = ?"
	_, err := m.DB.Exec(query, id)
	return err
}
