package model

import (
	"fmt"

	"github.com/jeypc/go-auth/go-auth/entities"
)

func (m *UserModel) GetNotifications(nim string) ([]*entities.Lamaran, error) {
	query := `
        SELECT id, nama_pekerjaan, perusahaan, surat_lamaran
        FROM lamaran
        WHERE nim = ? AND approve = 1`

	rows, err := m.DB.Query(query, nim)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []*entities.Lamaran
	for rows.Next() {
		var lamaran entities.Lamaran
		if err := rows.Scan(&lamaran.ID, &lamaran.NamaPekerjaan, &lamaran.Perusahaan, &lamaran.SuratLamaran); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		notifications = append(notifications, &lamaran)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return notifications, nil
}

func (nm *UserModel) GetNewNotifications(id int) ([]*entities.Lamaran, error) {
	query := "SELECT id, nama_pekerjaan, perusahaan, surat_lamaran FROM lamaran WHERE id = ?"
	rows, err := nm.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*entities.Lamaran
	for rows.Next() {
		var newnotification entities.Lamaran
		if err := rows.Scan(&newnotification.ID, &newnotification.NamaPekerjaan, &newnotification.Perusahaan, &newnotification.SuratLamaran); err != nil {
			return nil, err
		}
		notifications = append(notifications, &newnotification)
	}

	return notifications, nil
}
