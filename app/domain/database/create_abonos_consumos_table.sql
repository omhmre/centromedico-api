-- Script de creación de tablas para Cuadro de Abonos y Consumos (Digitel / Pacientes)

CREATE TABLE IF NOT EXISTS medi001.paciente_abonos (
    id SERIAL PRIMARY KEY,
    cedula_paciente VARCHAR(20) NOT NULL,
    patrocinante VARCHAR(100) NOT NULL DEFAULT 'Particular',
    fecha_abono TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    monto NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    tasa NUMERIC(12, 4) NOT NULL DEFAULT 1.0000,
    metodo_pago VARCHAR(50) DEFAULT 'Transferencia',
    referencia VARCHAR(100) DEFAULT '',
    observaciones TEXT DEFAULT '',
    creado_por VARCHAR(100) DEFAULT '',
    fecha_creacion TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS medi001.paciente_consumos (
    id SERIAL PRIMARY KEY,
    id_abono INT NULL REFERENCES medi001.paciente_abonos(id) ON DELETE SET NULL,
    cedula_paciente VARCHAR(20) NOT NULL,
    id_cita INT NULL,
    fecha_consumo TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    especialidad VARCHAR(100) NOT NULL DEFAULT 'General',
    servicio VARCHAR(150) DEFAULT '',
    monto NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    observaciones TEXT DEFAULT '',
    creado_por VARCHAR(100) DEFAULT '',
    fecha_creacion TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Índices para optimizar búsquedas por paciente y patrocinante
CREATE INDEX IF NOT EXISTS idx_abonos_cedula ON medi001.paciente_abonos(cedula_paciente);
CREATE INDEX IF NOT EXISTS idx_abonos_patrocinante ON medi001.paciente_abonos(patrocinante);
CREATE INDEX IF NOT EXISTS idx_consumos_cedula ON medi001.paciente_consumos(cedula_paciente);
CREATE INDEX IF NOT EXISTS idx_consumos_abono ON medi001.paciente_consumos(id_abono);
