-- Add audit fields to citas table
ALTER TABLE medi001.citas ADD COLUMN IF NOT EXISTS usuario_creacion VARCHAR(100);
ALTER TABLE medi001.citas ADD COLUMN IF NOT EXISTS fecha_creacion TIMESTAMP;

-- Update existing records to have at least some value if null (optional, but good for consistency)
UPDATE medi001.citas SET usuario_creacion = usuario_operacion WHERE usuario_creacion IS NULL;
UPDATE medi001.citas SET fecha_creacion = fecha_operacion WHERE fecha_creacion IS NULL;
UPDATE medi001.citas SET fecha_creacion = inicio WHERE fecha_creacion IS NULL;
