-- Add new fields to doctores table
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS cedula character varying(20);
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS fecha_nacimiento date;
ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS fecha_ingreso date;

-- Add new fields to personal table if any are missing (though we saw they exist)
-- Let's just double check and ensure they are there
DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema='medi001' AND table_name='personal' AND column_name='cedula') THEN
        ALTER TABLE medi001.personal ADD COLUMN cedula character varying(20);
    END IF;
END $$;
