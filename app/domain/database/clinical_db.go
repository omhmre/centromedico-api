package database

import (
	"log"
	"omhmre.com/centromedico/app/domain/models"
)

// GetClinicalHistory returns the clinical history base object for a patient
func (d *DB) GetClinicalHistory(cedula string) (*models.ClinicalHistory, error) {
	row := d.db.QueryRow(sqlGetClinicalHistory, cedula)
	var h models.ClinicalHistory
	err := row.Scan(&h.ID, &h.CedulaPaciente, &h.AntecedentesFamiliares, &h.PatologiasPrevias, &h.Alergias, &h.Habitos, &h.CreatedAt, &h.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// UpsertClinicalHistory inserts or updates the clinical history base object
func (d *DB) UpsertClinicalHistory(h *models.ClinicalHistory) (int, error) {
	var id int
	err := d.db.QueryRow(sqlUpsertClinicalHistory, h.CedulaPaciente, h.AntecedentesFamiliares, h.PatologiasPrevias, h.Alergias, h.Habitos).Scan(&id)
	if err != nil {
		log.Printf("Error upserting clinical history: %v", err)
		return 0, err
	}
	return id, nil
}

// GetClinicalRecords returns all clinical records for a patient, along with their attachments
func (d *DB) GetClinicalRecords(cedula string) ([]models.ClinicalRecord, error) {
	rows, err := d.db.Query(sqlGetClinicalRecords, cedula)
	if err != nil {
		log.Printf("Error querying clinical records: %v", err)
		return nil, err
	}
	defer rows.Close()

	var records []models.ClinicalRecord
	for rows.Next() {
		var r models.ClinicalRecord
		err := rows.Scan(&r.ID, &r.CedulaPaciente, &r.IDEspecialista, &r.NombreEspecialista, &r.EspecialidadEspecialista,
			&r.MotivoConsulta, &r.ExamenFisico, &r.Diagnostico, &r.Tratamiento, &r.Observaciones, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// fetch attachments for this record
		attachments, _ := d.GetClinicalAttachments(r.ID)
		r.Attachments = attachments

		records = append(records, r)
	}
	return records, nil
}

// InsertClinicalRecord creates a new clinical record note
func (d *DB) InsertClinicalRecord(r *models.ClinicalRecord) (int, error) {
	var id int
	err := d.db.QueryRow(sqlInsertClinicalRecord, r.CedulaPaciente, r.IDEspecialista, r.MotivoConsulta, r.ExamenFisico, r.Diagnostico, r.Tratamiento, r.Observaciones).Scan(&id)
	if err != nil {
		log.Printf("Error inserting clinical record: %v", err)
		return 0, err
	}
	return id, nil
}

// GetClinicalAttachments retrieves attachments for a specific clinical record
func (d *DB) GetClinicalAttachments(recordID int) ([]models.ClinicalAttachment, error) {
	rows, err := d.db.Query(sqlGetClinicalAttachments, recordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []models.ClinicalAttachment
	for rows.Next() {
		var a models.ClinicalAttachment
		err := rows.Scan(&a.ID, &a.IDRegistro, &a.NombreArchivo, &a.TipoArchivo, &a.UrlArchivo, &a.CreatedAt)
		if err != nil {
			continue
		}
		attachments = append(attachments, a)
	}
	return attachments, nil
}

// InsertClinicalAttachment saves a new file attachment info into DB
func (d *DB) InsertClinicalAttachment(a *models.ClinicalAttachment) (int, error) {
	var id int
	err := d.db.QueryRow(sqlInsertClinicalAttachment, a.IDRegistro, a.NombreArchivo, a.TipoArchivo, a.UrlArchivo).Scan(&id)
	if err != nil {
		log.Printf("Error inserting clinical attachment: %v", err)
		return 0, err
	}
	return id, nil
}
