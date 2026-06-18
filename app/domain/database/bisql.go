package database

const sqlBIResumenGeneral = `
SELECT 
    COUNT(*)::integer AS total_citas,
    COUNT(CASE WHEN status = 'Completada' THEN 1 END)::integer AS completadas,
    COUNT(CASE WHEN status = 'Cancelada' THEN 1 END)::integer AS canceladas,
    COUNT(CASE WHEN status = 'Pendiente' THEN 1 END)::integer AS pendientes,
    COALESCE(SUM(CASE WHEN status = 'Completada' THEN montoref ELSE 0 END), 0)::double precision AS total_ingresos_usd,
    COALESCE(AVG(CASE WHEN status = 'Completada' THEN montoref END), 0)::double precision AS ingreso_por_cita,
    COUNT(DISTINCT cedula)::integer AS pacientes_unicos,
    COALESCE((COUNT(CASE WHEN status = 'Completada' THEN 1 END)::double precision / NULLIF(COUNT(*), 0)) * 100, 0)::double precision AS tasa_completadas
FROM 
    medi001.citas
WHERE 
    inicio >= $1::timestamp AND inicio <= $2::timestamp;`

const sqlBICitasPorDia = `
SELECT 
    TO_CHAR(inicio, 'YYYY-MM-DD') AS fecha,
    COUNT(*)::integer AS citas,
    COALESCE(SUM(CASE WHEN status = 'Completada' THEN montoref ELSE 0 END), 0)::double precision AS ingresos
FROM 
    medi001.citas
WHERE 
    inicio >= $1::timestamp AND inicio <= $2::timestamp
GROUP BY 
    TO_CHAR(inicio, 'YYYY-MM-DD')
ORDER BY 
    fecha ASC;`

const sqlBICitasPorEspecialidad = `
SELECT 
    d.espec AS especialidad,
    COUNT(*)::integer AS total_citas,
    COUNT(CASE WHEN c.status = 'Completada' THEN 1 END)::integer AS completadas,
    COALESCE(SUM(CASE WHEN c.status = 'Completada' THEN c.montoref ELSE 0 END), 0)::double precision AS total_ingresos,
    COALESCE((COUNT(CASE WHEN c.status = 'Completada' THEN 1 END)::double precision / NULLIF(COUNT(*), 0)) * 100, 0)::double precision AS tasa_eficiencia
FROM 
    medi001.citas c
INNER JOIN 
    medi001.doctores d ON c.iddoctor = d.id
WHERE 
    c.inicio >= $1::timestamp AND c.inicio <= $2::timestamp
GROUP BY 
    d.espec
ORDER BY 
    total_ingresos DESC;`

const sqlBIRendimientoDoctor = `
SELECT 
    d.nombres AS nombre,
    COUNT(*)::integer AS total_citas,
    COUNT(CASE WHEN c.status = 'Completada' THEN 1 END)::integer AS completadas,
    COALESCE(SUM(CASE WHEN c.status = 'Completada' THEN c.montoref ELSE 0 END), 0)::double precision AS ingresos,
    COALESCE((COUNT(CASE WHEN c.status = 'Completada' THEN 1 END)::double precision / NULLIF(COUNT(*), 0)) * 100, 0)::double precision AS eficiencia
FROM 
    medi001.citas c
INNER JOIN 
    medi001.doctores d ON c.iddoctor = d.id
WHERE 
    c.inicio >= $1::timestamp AND c.inicio <= $2::timestamp
GROUP BY 
    d.id, d.nombres
ORDER BY 
    ingresos DESC;`

const sqlBIMetodosPago = `
SELECT 
    p.paymentmethod AS metodo,
    COALESCE(SUM(p.amount), 0)::double precision AS total,
    COALESCE((SUM(p.amount) / NULLIF(SUM(SUM(p.amount)) OVER(), 0)) * 100, 0)::double precision AS porcentaje
FROM 
    medi001.payments p
WHERE 
    p.date >= $1::timestamp AND p.date <= $2::timestamp
GROUP BY 
    p.paymentmethod
ORDER BY 
    total DESC;`

const sqlBIHeatmap = `
SELECT 
    EXTRACT(DOW FROM inicio)::integer AS dia_semana,
    EXTRACT(HOUR FROM inicio)::integer AS hora,
    COUNT(*)::integer AS cantidad
FROM 
    medi001.citas
WHERE 
    inicio >= $1::timestamp AND inicio <= $2::timestamp
GROUP BY 
    dia_semana, hora
ORDER BY 
    dia_semana, hora;`
