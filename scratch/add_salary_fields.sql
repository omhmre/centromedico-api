-- Add salary fields to doctores table
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS sueldo numeric(15,2) DEFAULT 0;
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS frecuencia_pago character varying(50) DEFAULT 'Mensual';
