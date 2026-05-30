-- Historial Clínico - Tablas de Apoyo

-- Antecedentes Genales del Paciente
CREATE TABLE IF NOT EXISTS medi001.paciente_antecedentes (
    id_paciente INTEGER PRIMARY KEY REFERENCES medi001.pacientes(id) ON DELETE CASCADE,
    medicos TEXT,
    quirurgicos TEXT,
    alergicos TEXT,
    familiares TEXT,
    habitos TEXT,
    otros TEXT,
    ultima_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Registro de Signos Vitales
CREATE TABLE IF NOT EXISTS medi001.paciente_signos_vitales (
    id SERIAL PRIMARY KEY,
    id_paciente INTEGER REFERENCES medi001.pacientes(id) ON DELETE CASCADE,
    id_cita INTEGER REFERENCES medi001.citas(id) ON DELETE SET NULL,
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tension_arterial VARCHAR(20),
    frecuencia_cardiaca INTEGER,
    frecuencia_respiratoria INTEGER,
    temperatura DECIMAL(4,2),
    saturacion_oxigeno INTEGER,
    peso DECIMAL(5,2),
    talla DECIMAL(5,2),
    imc DECIMAL(4,2),
    notas TEXT,
    usuario_operacion VARCHAR(50)
);

-- Evolución Clínica / Informes Médicos
-- (Si ya existe la tabla, alteramos si es necesario, si no, la creamos)
CREATE TABLE IF NOT EXISTS medi001.informe_medico (
    id SERIAL PRIMARY KEY,
    id_paciente INTEGER REFERENCES medi001.pacientes(id) ON DELETE CASCADE,
    id_doctor INTEGER REFERENCES medi001.doctores(id) ON DELETE CASCADE,
    id_cita INTEGER REFERENCES medi001.citas(id) ON DELETE SET NULL,
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    diagnostico TEXT,
    evolucion TEXT,
    plan TEXT,
    recomendaciones TEXT,
    contenido TEXT,
    entregado BOOLEAN DEFAULT FALSE,
    fecha_entrega TIMESTAMP,
    modificado_post_entrega BOOLEAN DEFAULT FALSE,
    usuario_operacion VARCHAR(50)
);

-- Migraciones para tablas existentes que no tengan las columnas nuevas
ALTER TABLE medi001.informe_medico ADD COLUMN IF NOT EXISTS contenido TEXT;
ALTER TABLE medi001.informe_medico ADD COLUMN IF NOT EXISTS entregado BOOLEAN DEFAULT FALSE;
ALTER TABLE medi001.informe_medico ADD COLUMN IF NOT EXISTS fecha_entrega TIMESTAMP;
ALTER TABLE medi001.informe_medico ADD COLUMN IF NOT EXISTS modificado_post_entrega BOOLEAN DEFAULT FALSE;

-- Tabla de Licenciamiento Profesional
CREATE TABLE IF NOT EXISTS seguridad.licencia (
    id SERIAL PRIMARY KEY,
    key_hash VARCHAR(64) UNIQUE NOT NULL,
    license_key TEXT NOT NULL,
    client_name VARCHAR(150),
    rif VARCHAR(50),
    valid_until TIMESTAMP,
    is_premium BOOLEAN DEFAULT FALSE,
    activated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_verified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

