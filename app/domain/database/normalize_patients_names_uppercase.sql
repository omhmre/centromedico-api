-- normalize_patients_names_uppercase.sql
-- Normaliza todos los nombres de pacientes existentes en la base de datos a mayúsculas.

UPDATE medi001.pacientes SET nombres = UPPER(nombres);
