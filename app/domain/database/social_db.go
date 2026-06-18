package database

import (
	"fmt"
	"log"
	"omhmre.com/centromedico/app/domain/models"
)

// GetSocialEvaluation returns the social evaluation for a patient
func (d *DB) GetSocialEvaluation(cedula string) (*models.SocialEvaluation, error) {
	row := d.db.QueryRow(sqlGetSocialEvaluation, cedula)
	var e models.SocialEvaluation
	err := row.Scan(
		&e.ID,
		&e.CedulaPaciente,
		&e.IDEspecialista,
		&e.NombreEspecialista,
		&e.GrupoFamiliar,
		&e.SituacionEconomica,
		&e.ViviendaEntorno,
		&e.AspectoSalud,
		&e.DiagnosticoSocial,
		&e.PlanAccion,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error al obtener evaluacion social: %w", err)
	}
	return &e, nil
}

// UpsertSocialEvaluation inserts or updates a social evaluation
func (d *DB) UpsertSocialEvaluation(e *models.SocialEvaluation) (int, error) {
	var id int
	err := d.db.QueryRow(
		sqlUpsertSocialEvaluation,
		e.CedulaPaciente,
		e.IDEspecialista,
		e.GrupoFamiliar,
		e.SituacionEconomica,
		e.ViviendaEntorno,
		e.AspectoSalud,
		&e.DiagnosticoSocial,
		e.PlanAccion,
	).Scan(&id)
	if err != nil {
		log.Printf("Error upserting social evaluation: %v", err)
		return 0, fmt.Errorf("error al guardar evaluacion social: %w", err)
	}
	return id, nil
}
