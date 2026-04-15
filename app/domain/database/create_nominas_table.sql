-- Actualizar tabla personal
ALTER TABLE medi001.personal ADD COLUMN IF NOT EXISTS frecuencia_pago VARCHAR(20) DEFAULT 'Mensual';

-- Crear tabla de nóminas
CREATE TABLE IF NOT EXISTS medi001.nominas (
    id SERIAL PRIMARY KEY,
    personal_id INT NOT NULL REFERENCES medi001.personal(id) ON DELETE CASCADE,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    tipo_periodo VARCHAR(20) NOT NULL, -- Semanal, Quincenal, Mensual
    monto_base DECIMAL(15, 2) NOT NULL,
    bonificaciones DECIMAL(15, 2) DEFAULT 0,
    deducciones DECIMAL(15, 2) DEFAULT 0,
    monto_total DECIMAL(15, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'Pendiente', -- Pendiente, Pagado, Cancelado
    fecha_pago TIMESTAMP,
    egreso_id INT, -- Referencia al ID del gasto generado (si aplica)
    notas TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
