package model

import (
	"fmt"

	"github.com/jeypc/go-auth/go-auth/entities"
)

func (m *UserModel) GetLamarans() ([]*entities.Lamaran, error) {
	var lamarans []*entities.Lamaran

	// Assuming you have a global database connection `db`
	rows, err := m.DB.Query("SELECT id, nama_pekerjaan, perusahaan, surat_lamaran, approve FROM lamaran")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		lamaran := new(entities.Lamaran) // Inisialisasi lamaran sebagai instance baru dari entities.Lamaran
		if err := rows.Scan(&lamaran.ID, &lamaran.NamaPekerjaan, &lamaran.Perusahaan, &lamaran.SuratLamaran, &lamaran.Approve); err != nil {
			return nil, err
		}
		lamarans = append(lamarans, lamaran)
	}

	return lamarans, nil
}

func (m *UserModel) GetAllLamaran() ([]*entities.Lamaran, error) {
	rows, err := m.DB.Query("SELECT id, nama_pekerjaan, perusahaan, surat_lamaran, approve FROM lamaran")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lamarans []*entities.Lamaran // Menggunakan []Lamaran
	for rows.Next() {
		var lamaran entities.Lamaran // Menggunakan Lamaran
		err := rows.Scan(&lamaran.ID, &lamaran.NamaPekerjaan, &lamaran.Perusahaan, &lamaran.SuratLamaran, &lamaran.Approve)
		if err != nil {
			return nil, err
		}
		lamarans = append(lamarans, &lamaran)
	}

	return lamarans, nil // Mengembalikan slice lamarans dan error
}

func (m *UserModel) UpdateApprovalStatus(id int, isApproved bool) error {
	query := "UPDATE lamaran SET approve = ? WHERE id = ?"
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(isApproved, id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil
}
