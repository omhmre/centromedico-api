-- Migration: Add retirement fields to specialists
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS activo BOOLEAN DEFAULT TRUE;
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS fecha_retiro VARCHAR(20) DEFAULT '';
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS motivo_retiro TEXT DEFAULT '';

-- Update existing records to be active
UPDATE medi001.doctores SET activo = TRUE WHERE activo IS NULL;
