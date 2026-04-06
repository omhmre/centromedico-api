-- backend_audit_migration.sql
-- Run this script in your PostgreSQL database (schema medi001)

-- 1. Add columns to appointments (citas)
ALTER TABLE medi001.citas ADD COLUMN IF NOT EXISTS motivo_cancelacion TEXT;
ALTER TABLE medi001.citas ADD COLUMN IF NOT EXISTS usuario_operacion VARCHAR(100);
ALTER TABLE medi001.citas ADD COLUMN IF NOT EXISTS fecha_operacion TIMESTAMP;

-- 2. Add columns to payments
ALTER TABLE medi001.payments ADD COLUMN IF NOT EXISTS usuario_operacion VARCHAR(100);
ALTER TABLE medi001.payments ADD COLUMN IF NOT EXISTS fecha_operacion TIMESTAMP;

-- Optional: Add comments for clarity
COMMENT ON COLUMN medi001.citas.motivo_cancelacion IS 'Motivo por el cual se canceló la cita';
COMMENT ON COLUMN medi001.citas.usuario_operacion IS 'Usuario que realizó la última modificación';
COMMENT ON COLUMN medi001.citas.fecha_operacion IS 'Fecha y hora de la última modificación';

COMMENT ON COLUMN medi001.payments.usuario_operacion IS 'Usuario que registró el pago';
COMMENT ON COLUMN medi001.payments.fecha_operacion IS 'Fecha y hora del registro del pago';
