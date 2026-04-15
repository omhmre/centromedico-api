CREATE TABLE IF NOT EXISTS medi001.personal (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    cedula VARCHAR(20) UNIQUE NOT NULL,
    telefono VARCHAR(20),
    correo VARCHAR(255),
    direccion TEXT,
    titulo VARCHAR(255),
    universidad VARCHAR(255),
    fecha_ingreso DATE,
    fecha_nacimiento DATE,
    cargo VARCHAR(100),
    sueldo DECIMAL(15, 2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'Activo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
