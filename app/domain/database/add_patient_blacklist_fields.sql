-- add_patient_blacklist_fields.sql
ALTER TABLE medi001.pacientes ADD COLUMN is_blacklisted bool DEFAULT false;
ALTER TABLE medi001.pacientes ADD COLUMN blacklist_reason text;
ALTER TABLE medi001.pacientes ADD COLUMN blacklist_date timestamp with time zone;
