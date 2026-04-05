-- Migration: Add scheduling fields to medi001.doctores
-- Run this script in your PostgreSQL database to support the new specialist scheduling features.

ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS days_of_week JSONB DEFAULT '[1,2,3,4,5]';
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS start_time VARCHAR(10) DEFAULT '08:00';
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS end_time VARCHAR(10) DEFAULT '18:00';
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS slot_duration INTEGER DEFAULT 45;

-- Update existing rows to have the default values if they were NULL (if IF NOT EXISTS was used on an existing table)
UPDATE medi001.doctores SET days_of_week = '[1,2,3,4,5]' WHERE days_of_week IS NULL;
UPDATE medi001.doctores SET start_time = '08:00' WHERE start_time IS NULL;
UPDATE medi001.doctores SET end_time = '18:00' WHERE end_time IS NULL;
UPDATE medi001.doctores SET slot_duration = 45 WHERE slot_duration IS NULL;
