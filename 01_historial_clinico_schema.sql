-- Crear tabla de Historial Clínico (Datos estructurados / Antecedentes)
CREATE TABLE IF NOT EXISTS medi001.historial_clinico (
    id SERIAL PRIMARY KEY,
    cedula_paciente VARCHAR(50) NOT NULL REFERENCES medi001.pacientes(cedula) ON DELETE CASCADE,
    antecedentes_familiares TEXT DEFAULT '',
    patologias_previas TEXT DEFAULT '',
    alergias TEXT DEFAULT '',
    habitos TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(cedula_paciente)
);

-- Crear tabla de Registros Médicos (Notas de consultas)
CREATE TABLE IF NOT EXISTS medi001.registros_clinicos (
    id SERIAL PRIMARY KEY,
    cedula_paciente VARCHAR(50) NOT NULL REFERENCES medi001.pacientes(cedula) ON DELETE CASCADE,
    id_especialista INTEGER NOT NULL REFERENCES medi001.doctores(id) ON DELETE CASCADE,
    motivo_consulta TEXT NOT NULL,
    examen_fisico TEXT DEFAULT '',
    diagnostico TEXT DEFAULT '',
    tratamiento TEXT DEFAULT '',
    observaciones TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Crear tabla de Archivos Adjuntos (Exámenes, recetas, fotos)
CREATE TABLE IF NOT EXISTS medi001.archivos_clinicos (
    id SERIAL PRIMARY KEY,
    id_registro INTEGER NOT NULL REFERENCES medi001.registros_clinicos(id) ON DELETE CASCADE,
    nombre_archivo VARCHAR(255) NOT NULL,
    tipo_archivo VARCHAR(100) NOT NULL, -- ej: 'application/pdf', 'image/jpeg'
    url_archivo TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Habilitar RLS (Row Level Security) en Supabase para las nuevas tablas
ALTER TABLE medi001.historial_clinico ENABLE ROW LEVEL SECURITY;
ALTER TABLE medi001.registros_clinicos ENABLE ROW LEVEL SECURITY;
ALTER TABLE medi001.archivos_clinicos ENABLE ROW LEVEL SECURITY;

-- Crear políticas de seguridad (Supabase)
-- Se asume que el backend (Go) accede usando el Service Role, por lo que se salta RLS.
-- Si los clientes se conectaran directo, usaríamos estas políticas.
-- Para el rol autenticado de supabase:
CREATE POLICY "Permitir acceso a especialistas y gerencia" ON medi001.historial_clinico
    FOR ALL
    TO authenticated
    USING (true); -- Implementar validación de claims JWT si se accede directo desde Flutter.

CREATE POLICY "Permitir acceso a especialistas y gerencia" ON medi001.registros_clinicos
    FOR ALL
    TO authenticated
    USING (true);

CREATE POLICY "Permitir acceso a especialistas y gerencia" ON medi001.archivos_clinicos
    FOR ALL
    TO authenticated
    USING (true);
