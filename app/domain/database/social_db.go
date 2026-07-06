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
		&e.LugarNacimiento,
		&e.GradoEscolar,
		&e.Escolaridad,
		&e.ReferidoPor,
		&e.MadreNombre,
		&e.MadreEdad,
		&e.MadreCI,
		&e.MadreTelefono,
		&e.MadreOcupacion,
		&e.MadreCorreo,
		&e.MadreDireccion,
		&e.PadreNombre,
		&e.PadreEdad,
		&e.PadreCI,
		&e.PadreTelefono,
		&e.PadreOcupacion,
		&e.PadreDireccion,
		&e.AntecedentesDesarrollo,
		&e.GrupoFamiliar,
		&e.SituacionEconomica,
		&e.ViviendaEntorno,
		&e.AspectoSalud,
		&e.DiagnosticoSocial,
		&e.Conclusion,
		&e.PlanAccion,
		&e.Entregado,
		&e.UnlockedBy,
		&e.UnlockedAt,
		&e.UnlockReason,
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
		e.LugarNacimiento,
		e.GradoEscolar,
		e.Escolaridad,
		e.ReferidoPor,
		e.MadreNombre,
		e.MadreEdad,
		e.MadreCI,
		e.MadreTelefono,
		e.MadreOcupacion,
		e.MadreCorreo,
		e.MadreDireccion,
		e.PadreNombre,
		e.PadreEdad,
		e.PadreCI,
		e.PadreTelefono,
		e.PadreOcupacion,
		e.PadreDireccion,
		e.AntecedentesDesarrollo,
		e.GrupoFamiliar,
		e.SituacionEconomica,
		e.ViviendaEntorno,
		e.AspectoSalud,
		e.DiagnosticoSocial,
		e.Conclusion,
		e.PlanAccion,
		e.Entregado,
	).Scan(&id)
	if err != nil {
		log.Printf("Error upserting social evaluation: %v", err)
		return 0, fmt.Errorf("error al guardar evaluacion social: %w", err)
	}
	return id, nil
}

// UnlockSocialEvaluation unlocks a social evaluation and registers the admin who unlocked it
func (d *DB) UnlockSocialEvaluation(cedula string, unlockedBy int, reason string) (int, error) {
	var id int
	err := d.db.QueryRow(sqlUnlockSocialEvaluation, unlockedBy, reason, cedula).Scan(&id)
	if err != nil {
		log.Printf("Error unlocking social evaluation for %s: %v", cedula, err)
		return 0, fmt.Errorf("error al desbloquear evaluacion social: %w", err)
	}
	return id, nil
}
