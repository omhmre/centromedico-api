-- Crear tabla de Evaluaciones Sociales (Informe de Trabajadora Social)
CREATE TABLE IF NOT EXISTS medi001.evaluaciones_sociales (
    id SERIAL PRIMARY KEY,
    cedula_paciente VARCHAR(50) NOT NULL REFERENCES medi001.pacientes(cedula) ON DELETE CASCADE,
    id_especialista INTEGER NOT NULL REFERENCES medi001.doctores(id) ON DELETE CASCADE,
    grupo_familiar TEXT DEFAULT '',
    situacion_economica TEXT DEFAULT '',
    vivienda_entorno TEXT DEFAULT '',
    aspecto_salud TEXT DEFAULT '',
    diagnostico_social TEXT DEFAULT '',
    plan_accion TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(cedula_paciente)
);

-- Habilitar RLS (Row Level Security) en Supabase para las nuevas tablas
ALTER TABLE medi001.evaluaciones_sociales ENABLE ROW LEVEL SECURITY;

-- Crear política de seguridad (Supabase)
CREATE POLICY "Permitir acceso a especialistas y gerencia" ON medi001.evaluaciones_sociales
    FOR ALL
    TO authenticated
    USING (true);
