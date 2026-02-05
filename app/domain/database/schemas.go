package database

// var sqlLogin = `SELECT count(nombre) from seguridad.usuarios WHERE nombre = $1 AND contrasena = $2`
var sqlInventario = `SELECT 
	codigo, nombre, marca, unidad, costo, costoa, costopr, precio1, precio2, precio3, cantidad, enser, exento, 
  clasif, tipo, empaque, cantemp, pedido, disponible, preciom1, preciom2, preciom3, costodolar, dirfoto, foto, 
	descripcion, codservicio, preciovar, compuesto, mateprima, global, cantvar, espresent FROM empre001.inventario`

var sqlGetInventario = sqlInventario + ` ORDER BY codigo desc;`

var sqlGetInventarioFormal = sqlInventario + ` p 
where (compuesto = false and espresent = false) or mateprima = true or (
codigo in (select distinct p.codinv from empre001.presinventario p inner join empre001.itemspres i 
on p.id = i.codpres where p.codinv = i.codinv)
) order by nombre;`

var sqlGetItemsInventario = `select coditem, nombre, cantidad FROM EMPRE001.itemsinventario i where codinventario = $1;`
var sqlGetPresenInventario = `select id, codinv, presentacion, cantidad, precio FROM EMPRE001.presinventario i where codinv = $1;`

var sqlGetInventarioCompacto = sqlInventario + ` where global = true order by global desc, nombre asc;`

var sqlGetInventarioNombre = `SELECT codigo, nombre, precio1 FROM empre001.inventario ORDER BY nombre;`
var sqlAddInventario = `INSERT INTO empre001.inventario (codigo, nombre, marca, unidad, costo, costoa, costopr, precio1, precio2, precio3, cantidad, enser, exento, clasif, tipo, empaque, cantemp, pedido, disponible, preciom1, preciom2, preciom3, costodolar, dirfoto, foto, descripcion, preciovar, compuesto, mateprima, espresent) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28, $29, $30);`

const sqlAddItemsInventario = `INSERT INTO empre001.itemsinventario (codinventario, coditem, nombre, cantidad) VALUES($1, $2, $3, $4);`
const sqlAddPresenInventario = `INSERT INTO empre001.presinventario (codinv, presentacion, cantidad, precio) VALUES($1, $2, $3, $4) RETURNING id;`
const sqlDelItemsInventario = `DELETE FROM empre001.itemsinventario WHERE codinventario = $1);`
const sqlDelPresenInventario = `DELETE FROM empre001.presinventario WHERE codinv = $1;`

var sqlUpdInventario = `UPDATE empre001.inventario SET 
nombre = $2, 
marca = $3, 
unidad = $4, 
costo = $5, 
costoa = $6, 
costopr = $7, 
precio1 = $8, 
precio2 = $9, 
precio3 = $10, 
cantidad = $11, 
enser = $12, 
exento = $13, 
clasif = $14, 
tipo = $15, 
empaque = $16, 
cantemp = $17, 
pedido = $18, 
disponible = $19, 
preciom1 = $20, 
preciom2 = $21, 
preciom3 = $22, 
costodolar = $23, 
dirfoto = $24, 
foto = $25, 
descripcion = $26, 
preciovar = $27, 
compuesto = $28, 
mateprima = $29,
espresent = $30 
WHERE codigo = $1;`

var sqlDelInventario = `DELETE FROM empre001.inventario WHERE codigo = $1;`
var sqlGetMenu = `SELECT * FROM empre001.menu order by nombre;`
var sqlGetClase = `SELECT * FROM empre001.clasemenu order by id;`
var sqlGetMenuClase = `SELECT * FROM empre001.menu where idclase = $1`
var sqlAddClase = `INSERT INTO empre001.clasemenu (nombre) VALUES($1);`
var sqlUpdClase = `UPDATE empre001.clasemenu SET nombre=$2 WHERE id=$1;`
var sqlDelClase = `DELETE FROM empre001.clasemenu WHERE id=$1;`
var sqlAddMenu = `INSERT INTO empre001.menu(codigo, nombre, descripcion, idclase, precio1, precio2, precio3, cantidad, preciom1, preciom2, preciom3, dirfoto, foto) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`
var sqlUpdMenu = `UPDATE empre001.menu SET nombre=$2, descripcion=$3, idclase=$4, precio1=$5, precio2=$6, precio3=$7, cantidad=$8, preciom1=$9, preciom2=$10, preciom3=$11, dirfoto=$12, foto=$13 where codigo=$1`
var sqlUpdInventCod = `UPDATE empre001.inventario SET nombre = $2, descripcion = $3, precio1=$4, precio2=$5, precio3=$6, preciom1=$7, preciom2=$8, preciom3=$9, dirfoto=$10, foto=$11 WHERE codigo = $1;`
var sqlDelMenu = `DELETE FROM empre001.menu where codigo=$1`
var sqlDelAllMenu = `DELETE FROM empre001.menu`
var sqlGetEmpre = `SELECT id, rif, rasocial, dirfisc, ciudad, estado, telf, logo, comercial, slogan, iva, correo, instagram, whatsapp FROM empre001.empresa;`
var sqlAddEmpresa = `INSERT INTO empre001.empresa (id, rif, rasocial, dirfisc, ciudad, estado, telf, logo, comercial, slogan, iva, correo, instagram, whatsapp) 
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`
var sqlUpdEmpresa = `UPDATE empre001.empresa SET rif=$2, rasocial=$3, dirfisc=$4, ciudad=$5, estado=$6, telf=$7, logo=$8, comercial=$9, slogan=$10, iva=$11,
correo=$12, instagram=$13, whatsapp=$14 WHERE id=$1;`
var sqlDelEmpresa = `DELETE FROM empre001.empresa WHERE id=$1;`
var sqlGetPrefacturas = `SELECT id, idcliente, fecha, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, tasadiv, monetodiv FROM empre001.prefacturas;`

const sqlDocumentos = `SELECT a.id, a.idcliente, b.nombre as cliente, b.dirfiscal, b.rif, b.persconta, b.tlfconta, to_char(a.fecha, 'DD/MM/YYYY') as fecha, a.subtotal, a.dscto, a.mototal, 
a.deimp, a.tasaimp, a.moimp, a.moneto, a.tasadiv, a.monetodiv, a.idvendedor, COALESCE(v.nombre, '') as vendedor, a.idsesion, a.idmesa, a.idmesonero, a.pagado, a.porpagar, a.cambio,
b.cxcbs, b.cxcdiv, a.valido `

var sqlGetFacturas = sqlDocumentos +
	`, case 
	when p.tasa > 1 then true 
	else false
end as esdivisa
FROM empre001.facturas a 
left join empre001.clientes b on a.idcliente = b.id 
left join empre001.vendedores v on a.idvendedor = v.id
left join (select idfact, max(tasa) as tasa from empre001.detpagos d group by idfact order by idfact) p on a.id = p.idfact `

const sqlGetPresupuestos = `SELECT a.id, a.idcliente, b.nombre as cliente, b.dirfiscal, b.rif, b.persconta, b.tlfconta, to_char(a.fecha, 'DD/MM/YYYY') as fecha, diasvence, to_char(a.vence, 'DD/MM/YYYY') as vence, a.subtotal, a.dscto, a.mototal, 
a.deimp, a.tasaimp, a.moimp, a.moneto, a.tasadiv, a.monetodiv, a.idvendedor, COALESCE(v.nombre, '') as vendedor, a.idsesion, a.idmesa, a.idmesonero, a.condiciones
FROM empre001.presupuestos a left join empre001.clientes b on a.idcliente = b.id left join empre001.vendedores v on a.idvendedor = v.id `

const sqlWherePresupuestos = `WHERE a.id = $1;`

var sqlGetNotasEntrega = sqlDocumentos + `FROM empre001.entregas a left join empre001.clientes b on a.idcliente = b.id left join empre001.vendedores v
on a.idvendedor = v.id`

var sqlGetFactura = ` WHERE id = $1 order by a.id desc;`

var sqlGetVentasFecha = ` WHERE a.fecha between $1 and $2 order by a.id desc;`

var sqlAnularFactura = `update empre001.facturas set valido = false where id = $1`
var sqlAnularInventario = `CALL empre001.dev_inv($1)`

var sqlGetPrefactura = `SELECT p.id, p.idcliente, c.nombre, p.fecha, p.subtotal, p.dscto, p.mototal, p.deimp, p.tasaimp, p.moimp, p.moneto, p.idvendedor, p.idsesion, p.idmesa, p.idmesonero, p.pagado, p.porpagar, p.cambio, p.tasadiv, p.monetodiv,
	c.cxcbs, c.cxcdiv, a.valido  FROM empre001.prefacturas p left join empre001.clientes c on p.idcliente = c.id where p.id = $1;`
var sqlPostPrefactura = `UPDATE empre001.prefacturas SET 
idcliente = $2, 
fecha = $3, 
subtotal = $4, 
dscto = $5, 
mototal = $6, 
deimp = $7, 
tasaimp = $8, 
moimp = $9, 
moneto = $10, 
idvendedor = $11, 
idsesion = $12,
idmesa = $13, 
idmesonero = $14,
pagado = 0.0,
porpagar = $15,
cambio = 0.0, 
tasadiv = $16,
monetodiv = $17
WHERE id = $1;`

var sqlPostPrefacturaNueva = `insert into empre001.prefacturas(idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, porpagar, tasadiv, monetodiv, condiciones)
values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,$18,$19) RETURNING id;`

const sqlAddItems = `INSERT INTO empre001.detprefacturas (idprefact, codprod, cant, precio, subtotal, descuento, neto, descripcion, cantmp, cantpres, iditempres, nombre) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
const sqlAddItemsPres = `insert into empre001.itemspres(codpres, codinv, cantidad) values($1, $2, $3);`
const sqlDelItems = `delete from empre001.detprefacturas where idprefact = $1;`

var sqlPutPrefactura = `UPDATE empre001.prefacturas SET 
idcliente = $2, 
fecha = $3, 
subtotal = $4, 
dscto = $5, 
mototal = $6, 
deimp = $7, 
tasaimp = $8, 
moimp = $9, 
moneto = $10, 
idvendedor = $11, 
idsesion= $12,
tasadiv = $13,
monetodiv = $14
where id = $1;`

const sqlGetClientes = `SELECT id, tipo, nombre, rif, dirfiscal, ciudad, estado, telf, correo, twitter, facebook, whatsapp, instagram, status, clasif, dscto, cred, diascr, cxcbs, persconta, tlfconta, codvend, cxcdiv
FROM empre001.clientes order by id;`

const sqlUpdCliente = `UPDATE empre001.clientes
SET 
tipo=$2, 
nombre=$3, 
rif=$4, 
dirfiscal=$5, 
ciudad=$6, 
estado=$7, 
telf=$8, 
correo=$9, 
twitter=$10, 
facebook=$11, 
whatsapp=$12, 
instagram=$13, 
status=$14, 
clasif=$15, 
dscto=$16, 
cred=$17, 
diascr=$18, 
cxcbs=$19, 
persconta=$20, 
tlfconta=$21, 
codvend=$22, 
cxcdiv=$23
WHERE id=$1`

const sqlDelCliente = `DELETE FROM empre001.clientes WHERE id = $1;`

const sqlPostCliente = `INSERT INTO empre001.clientes
	(id, tipo, nombre, rif, dirfiscal, ciudad, estado, telf, correo, twitter, facebook, whatsapp, instagram, status, clasif, dscto, cred, diascr, cxcbs, persconta, tlfconta, codvend, cxcdiv)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23);`

const sqlGetItemsCxcCliente = `select id, idfact, fecha, montobs, cobradobs, saldobs, montodiv, cobradodiv, saldodiv 
from empre001.cxc where codclie = $1;`

const sqlAddMesa = `INSERT INTO empre001.mesas(id, nombre) VALUES($1, $2)`
const sqlUpdMesa = `UPDATE empre001.mesas SET nombre = $2 WHERE id = $1`
const sqlDelMesa = `DELETE FROM empre001.mesas WHERE id = $1`
const sqlAbrirMesa = `UPDATE empre001.mesas SET abierta = $1, subtotal = $2, inicio = now(), idprefactura = $3, idcliente = cast($4 as varchar(15)),
cliente = (select c.nombre from empre001.clientes c where c.id = cast($4 as varchar(15))), idmesonero = cast($5 as varchar(15)), 
mesonero = (select m.nombre from empre001.mesoneros m where m.id = cast($5 as varchar(15))) WHERE id = $6;`
const sqlActualizarMesa = `UPDATE empre001.mesas SET subtotal = $1, idcliente = cast($2 as varchar(15)),
cliente = (select c.nombre from empre001.clientes c where c.id = cast($2 as varchar(15))), 
idmesonero = cast($3 as varchar(15)), mesonero = (select m.nombre from empre001.mesoneros m where m.id = cast($3 as varchar(15))) WHERE id = $4;`
const sqlCerrarMesa = `UPDATE empre001.mesas SET abierta = 0, subtotal = 0.0, inicio = now(), idprefactura = 0, idcliente = '', cliente = '', idmesonero = '', mesonero = '' WHERE id = $1`

const sqlUpdMesonero = `UPDATE empre001.mesoneros SET nombre = $2, direccion = &3, telefono = $4 WHERE id = $1`
const sqlDelMesonero = `DELETE FROM empre001.mesoneros WHERE id = $1`
const sqlAddMesonero = `INSERT INTO empre001.mesoneros(id, nombre, direccion, telefono) VALUES($1,$2,$3,$4)`

const sqlGetItemsPreFacturas = `select d.idprefact, d.codprod, d.nombre, d.cant, d.precio, d.subtotal, d.descuento, d.neto, d.descripcion, iditempres from empre001.detprefacturas d where d.idprefact = $1;`

const sqlPostFactura = `insert into empre001.facturas(idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, pagado, porpagar, cambio, condiciones) 
SELECT idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, 
	 $2 as pagado, 0.00, 0.00, condiciones FROM empre001.prefacturas where id = $1 returning id;`

const sqlPostPresupuesto = `insert into empre001.presupuestos(idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, pagado, porpagar, cambio, condiciones) 
SELECT idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, 
moneto, 0.00, 0.00, condiciones FROM empre001.prefacturas where id = $1 returning id;`

// const sqlPostFactura = `insert into empre001.facturas(idcliente, fecha, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, pagado, porpagar, cambio)
// SELECT idcliente, fecha, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv,
// 	 $2 as pagado,
// 				case
// 				when $2 > moneto then 0.00
// 				else moneto - $2
// 			end
// 			as porpagar,
// 			case
// 				when $2 > moneto then $2 - moneto
// 				else 0.00
// 			end
// 			as cambio
// FROM empre001.prefacturas where id = $1 returning id;`

const sqlPostEntrega = `insert into empre001.entregas(idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, pagado, porpagar, cambio, condiciones) 
SELECT idcliente, fecha, diasvence, vence, subtotal, dscto, mototal, deimp, tasaimp, moimp, moneto, idvendedor, idsesion, idmesa, idmesonero, tasadiv, monetodiv, 
	 $2 as pagado,  
				case 
				when $2 > moneto then 0 
				else moneto - $2
			end 
			as porpagar,  
			case 
				when $2 > moneto then $2 - moneto 
				else 0 
			end 
			as cambio,
            condiciones
FROM empre001.prefacturas where id = $1 returning id;`

const insertItemsFacturas = `codprod, cant, precio, subtotal, descuento, neto, descripcion, cantmp, cantpres, iditempres, nombre)
SELECT $1, codprod, cant, precio, subtotal, descuento, neto, descripcion, cantmp, cantpres, iditempres, nombre 
FROM empre001.detprefacturas where idprefact = $2;`

const sqlPostDetFactura = `INSERT INTO empre001.detfacturas (idfact, ` + insertItemsFacturas

const sqlPostDetEntrega = `INSERT INTO empre001.detentregas (identrega, ` + insertItemsFacturas

const sqlPostDetPresupuesto = `INSERT INTO empre001.detpresupuestos (idpre, ` + insertItemsFacturas

const getItemsDocumentos = `select d.idfact, d.codprod, d.nombre as producto, d.cant, d.precio, d.subtotal, d.descuento, d.neto, d.descripcion, d.cantmp, d.cantpres, iditempres `

const sqlGetItemsFacturas = getItemsDocumentos + `from empre001.detfacturas d where d.idfact =  $1;`

const sqlGetItemsEntregas = getItemsDocumentos + `from empre001.detentregas d where d.identrega =  $1;`

var sqlGetVentasFactura = `select count(id) as CantFact, sum(subtotal) as Monto, sum(dscto) as Descuento, 
sum(moneto) as SubTotal, sum(monetodiv) as SubTotalDivisa from empre001.facturas f where fecha between $1 and $2 and valido = true`

var sqlGetVentasProductos = `SELECT d.codprod, d.nombre,  sum(d.cant) as cant, sum(d.subtotal) as subtotal, sum(d.subtotal/f.tasadiv) as subtotaldivisa,
sum(d.cantmp) as CantMP, sum(d.cantpres * d.cant) as cantpres FROM empre001.detfacturas d inner join empre001.facturas f on d.idfact = f.id 
where f.fecha between $1 and $2 and f.valido = true group by 1, 2 order by 4 desc;`

var sqlUpdDivisa = `UPDATE empre001.divisas SET tasabs = $1, fechatasa = $2;`
var sqlUpdDivisaInst = `UPDATE empre001.instpagos SET tasa = $1 WHERE simbolo = '$';`

const sqlPostPagos = `INSERT INTO empre001.pagos (idcliente, fecha, monto, dscto, total, idsesion) VALUES($1, now(), $2, $3, $4, $5) RETURNING id;`
const sqlAddItemsPago = `INSERT INTO empre001.detpagos (idpago, idinstpago, comenta, monto, tasa, total, idfact) VALUES($1, $2, $3, $4, $5, $6, $7);`
const sqlGetDetPagos = `select d.id, d.idpago, d.idinstpago, i.descripcion, d.comenta, d.monto, d.tasa, d.total, d.idfact 
from empre001.detpagos d left join empre001.instpagos i on d.idinstpago = i.id 
where d.idpago in (select distinct id from empre001.pagos p where p.fecha between $1 and $2);`
const sqlGetResumenDetPagos = `select d.idinstpago, i.descripcion, count(d.id) as cant, sum(d.monto) as montos, d.tasa, sum(d.total) as totales
from empre001.detpagos d left join empre001.instpagos i on d.idinstpago = i.id 
where d.idpago in (select distinct id from empre001.pagos p where p.fecha between $1 and $2)
group by 1,2,5;`

const sqlGetDivisas = `select id, divisa, simbolo, tasabs, fechatasa from empre001.divisas d ;`

const sqlGetVendedores = `SELECT id, cedula, nombre, direccion, telefono, correo, codvend FROM empre001.vendedores;`

const sqlGetVendedoresCxc = `SELECT id, cedula, nombre, direccion, telefono, correo, codvend, 0.00 as totcxc, 0.00 as totcxcbs, 0.00 as totcxcdiv, null FROM empre001.vendedores;`

const sqlGetInstrumentos = `Select id, descripcion, tasa, simbolo from empre001.instpagos order by id;`

const sqlGetDetFact = `select idfact, codprod, nombre, cant, precio, subtotal, descuento, neto, descripcion, cantmp, iditempres  
from empre001.detfacturas d where idfact = $1`

const sqlGetDetPre = `select idpre, codprod, nombre, cant, precio, subtotal, descuento, neto, descripcion, cantmp, iditempres  
from empre001.detpresupuestos d where idpre = $1`

const sqlGetDetNotasEnt = `select identrega as idfact, codprod, nombre, cant, precio, subtotal, descuento, neto, descripcion, cantmp, iditempres  
from empre001.detentregas d where identrega = $1`

const sqlGetDetPagosFact = `select d.id, d.idpago, d.idinstpago, i.descripcion, d.comenta, d.monto, d.tasa, d.total, d.idfact from empre001.detpagos d inner join empre001.instpagos i 
on d.idinstpago = i.id where d.idfact = $1`

const sqlGetInvoicePdf = `SELECT 
a.id, 
a.idcliente, 
b.nombre, 
b.dirfiscal,
b.rif,
b.persconta, 
b.tlfconta,
to_char(a.fecha, 'DD/MM/YYYY') as fecha, 
a.subtotal, 
a.dscto, 
a.mototal, 
a.deimp,
a.tasaimp,
a.moimp,
a.moneto,
a.tasadiv,
a.monetodiv,
a.idvendedor,
COALESCE(v.nombre, '') as vendedor,
a.idsesion, 
a.idmesa,
a.idmesonero,
a.pagado,
a.porpagar,
a.cambio,
b.cxcbs,
b.cxcdiv 
FROM empre001.facturas a left join empre001.clientes b on a.idcliente = b.id left join empre001.vendedores v
on a.idvendedor = v.id where a.id = $1;`

const sqlGetEntregaPdf = `SELECT 
a.id, a.idcliente, b.nombre, b.dirfiscal,b.rif,b.persconta, b.tlfconta,to_char(a.fecha, 'DD/MM/YYYY') as fecha, a.subtotal, a.dscto, 
a.mototal, a.deimp,a.tasaimp,a.moimp,a.moneto,a.tasadiv,a.monetodiv,a.idvendedor,COALESCE(v.nombre, '') as vendedor,a.idsesion, a.idmesa,
a.idmesonero, a.pagado,a.porpagar,a.cambio,b.cxcbs,b.cxcdiv 
FROM empre001.entregas a left join empre001.clientes b on a.idcliente = b.id left join empre001.vendedores v
on a.idvendedor = v.id where a.id = $1;`

const sqlGetNotaEntregaPdf = `SELECT 
a.id, 
a.idcliente, 
b.nombre, 
b.dirfiscal,
b.rif,
b.persconta, 
b.tlfconta,
to_char(a.fecha, 'DD/MM/YYYY') as fecha, 
a.subtotal, 
a.dscto, 
a.mototal, 
a.deimp,
a.tasaimp,
a.moimp,
a.moneto,
a.tasadiv,
a.monetodiv,
a.idvendedor,
COALESCE(v.nombre, '') as vendedor,
a.idsesion, 
a.idmesa,
a.idmesonero,
a.pagado,
a.porpagar,
a.cambio,
b.cxcbs,
b.cxcdiv 
FROM empre001.entregas a left join empre001.clientes b on a.idcliente = b.id left join empre001.vendedores v
on a.idvendedor = v.id where a.id = $1;`

const sqlGetUsuarios = `SELECT u.id, u.codigo, u.clave, u.nombre, u.idtipouser, t.tipo, u.idperfil, u.status, 
u.direccion, u.direccion2, u.ciudad, u.estado, u.telf, u.cel, u.correo, COALESCE(facebook, '') as facebook,
    COALESCE(whatsapp, '') as whatsapp,
    COALESCE(instagram, '') as instagram, u.idvendedor
FROM seguridad.usuarios u left join seguridad.tipouser t on u.idtipouser = t.id`

const sqlGetUsuarioWhere = ` where u.codigo = $1 and u.clave = $2;`

const sqlGetTopVentas = `SELECT d.codprod, d.nombre, sum(d.cant) as cantidad, sum(d.neto) as venta FROM empre001.detfacturas d 
WHERE d.idfact in (select distinct id from empre001.facturas f WHERE f.fecha >= CURRENT_DATE - INTERVAL '30 days'
  AND f.fecha < CURRENT_DATE and f.valido = true)
GROUP BY d.codprod, d.nombre
ORDER BY 3 DESC
LIMIT 10;`

const sqlGetVentasMes = `select m.numero_mes as nro, m.nombre_mes as mes, SUM(f.moneto) AS Bs, SUM(f.monetodiv) AS Divisas
FROM empre001.facturas f
JOIN empre001.meses m ON EXTRACT(MONTH FROM f.fecha) = m.numero_mes
where f.fecha >= CURRENT_DATE - INTERVAL '6 months' and f.valido = true GROUP BY 1, 2 ORDER BY 1  limit 6;`

const sqlGetProveedores = `SELECT id, tipo, proveedor, rif, dirfiscal, ciudad, estado, telf, correo, twitter, facebook, status, obs, clasif, credito, diascred, cxp
FROM empre001.proveedores;`

const sqlGetCxcResumen = `select c.id, c.idfact, c.codclie, c2.nombre, c.saldobs , c.saldodiv 
from empre001.cxc c inner join empre001.clientes c2 on c.codclie = c2.id 
where c.fecha between $1 and $2 and c.saldobs > 0.01;`

const sqlGetCxcVencida = `WITH cxc_vencida AS (
    SELECT
        *,
        CASE 
            WHEN dias_vencidos <= '30 days'::interval THEN '0-30 días'
            WHEN dias_vencidos BETWEEN '31 days'::interval AND '60 days'::interval THEN '31-60 días'
            WHEN dias_vencidos BETWEEN '61 days'::interval AND '90 days'::interval THEN '61-90 días'
            WHEN dias_vencidos BETWEEN '91 days'::interval AND '120 days'::interval THEN '91-120 días'  -- Nuevo rango
            ELSE 'Más de 120 días'
        END AS rango
    FROM
        (SELECT saldodiv, AGE(CURRENT_DATE, fecha) AS dias_vencidos FROM empre001.cxc) AS subquery
)
SELECT
    rango,
    SUM(saldodiv) AS saldo
FROM
    cxc_vencida
GROUP BY
    rango;`

const sqlGetCompras = `SELECT c.id, c.idprov, p.proveedor, c.fecha, c.subtotal, c.dscto, c.mototal, c.deimp, c.tasaimp, c.moimp, c.moneto, c.idsesion 
	FROM empre001.compras c INNER JOIN empre001.proveedores p ON c.idprov = p.id ORDER BY c.id DESC`

const sqlGetParametros = `select id, parametro, descripcion, valor, valores, descvalor from seguridad.parametros;`
const sqlAddParametros = `INSERT INTO seguridad.parametros(parametro, descripcion, valor, valores, descvalor) VALUES($1, $2, $3, $4, $5);`
const sqlUpdateParametros = `UPDATE seguridad.parametros SET descripcion = $2, valor = $3,  valores = $4, descvalor = $5 where parametro = $1`

const sqlAddProveedor = `INSERT INTO empre001.proveedores
(tipo, proveedor, rif, dirfiscal, ciudad, estado, telf, correo, twitter, facebook, status, obs, clasif, credito, diascred, cxp)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);`

const sqlUpdateProveedor = `UPDATE empre001.proveedores SET 
	tipo=$2, 
	proveedor=$3, 
	rif=$4, 
	dirfiscal=$5, 
	ciudad=$6, 
	estado=$7, 
	telf=$8, 
	correo=$9, 
	twitter=$10, 
	facebook=$11, 
	status=$12, 
	obs=$13, 
	clasif=$14, 
	credito=$15, 
	diascred=$16, 
	cxp=$17 
	WHERE id = $1;`

const sqlDelProveedor = `DELETE FROM empre001.proveedores WHERE id = $1;`

const sqlDelCompra = `DELETE FROM empre001.compras WHERE id = $1;`

const sqlGetNextIdCompra = `select nextval('empre001.compras_id_seq') ;`

const sqlInsertCompra = `INSERT INTO empre001.compras( 
          id, 
          idprov,
          fecha,
          subtotal,
          dscto,
          mototal,
          deimp,
          tasaimp,
          moimp,
          moneto,
          idsesion
        ) values(
            $1,
            $2,
            $3,
            $4,
            $5,
            $6,
            $7,
            $8,
            $9,
            $10,
            $11)`

const sqlInsertDetCompra = `INSERT INTO empre001.detcompra (idcompra, codprod, deprod, cant, costo, subtotalcosto)
          VALUES($1, $2, $3, $4, $5, $6);`

const sqlGetTasaActual = `select tasabs from empre001.divisas where id = 1;`

const sqlGetNextIdCxp = `select nextval('empre001.cxp_id_seq');`

const sqlInsertCxp = `INSERT INTO empre001.cxp (
            id,
            idcompra,
            codprov,
            fecha,
            montobs,
            pagadobs,
            saldobs,
            montodiv,
            pagadodiv,
            saldodiv,
            idsesion 
            ) VALUES(
              $1, 
              $2, 
              $3,       
              $4, 
              $5, 
              0.0, 
              $6, 
              $7, 
              0.0, 
              $8, 
              $9
              );`

const sqlGetEmailConfig = `SELECT id, smtp, puerto, usuario, clave, tls FROM empre001.emailconfig;`

const sqlUpdEmailConfig = `UPDATE empre001.emailconfig
SET smtp=$2, puerto=$3, usuario=$4, clave=$5, tls=$6 where id = $1;`
const sqlAddEmailConfig = `INSERT INTO empre001.emailconfig
(smtp, puerto, usuario, clave, tls)
VALUES($1, $2, $3, $4, $5);`

const sqlDelEmailConfig = `DELETE FROM empre001.emailconfig WHERE id = $1;`

const sqlDelPresupuesto = `DELETE FROM empre001.presupuestos `

const sqlGetDoctores = `select d.id, d.nombres, d.espec, d.dir, d.tlf, d.correo, d.whatsapp, d.instagram, 
d.tasapago from medi001.doctores d order by d.nombres;`

const sqlUpdDoctores = `update medi001.doctores set 
nombres = $2, espec = $3, dir = $4, tlf = $5, correo = $6, whatsapp = $7, instagram = $8, tasapago = $9 
where id = $1;`

const sqlGetPacientes = `SELECT id, cedula, nombres, fenac, representante, whatsapp, direccion, correo, diagnostico, cxc, created_at
FROM medi001.pacientes;`

const sqlPostPaciente = `INSERT INTO medi001.pacientes
(cedula, nombres, fenac, representante, whatsapp, direccion, correo, diagnostico, cxc, created_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;`

const sqlUpdPaciente = `UPDATE medi001.pacientes
SET cedula=$2, nombres=$3, fenac=$4, representante=$5, whatsapp=$6, direccion=$7, correo=$8, diagnostico=$9, cxc=$10
WHERE id=$1;`

const sqlDelPaciente = `DELETE FROM medi001.pacientes WHERE id = $1;`

const sqlPostDoctor = `INSERT INTO medi001.doctores
(id, nombres, espec, dir, tlf, correo, whatsapp, instagram, tasapago)
VALUES(nextval('medi001.doctores_id_seq'::regclass), $1, $2, $3, $4, $5, $6, $7, $8);`

const sqlDelDoctor = `DELETE FROM medi001.doctores WHERE id = $1;`

const sqlGetCitas = `SELECT
    c.id AS cita_id,
    c.iddoctor,
    d.nombres AS especialista,
    d.espec AS especialidad,
    c.cedula AS paciente_cedula,
    p.nombres AS paciente,
    c.motivo,
    c.inicio,
    c.fin,
    c.diagnostico,
    c.status AS cita_status,
    c.color,
    c.montoref,
    c.tasa,
    c.montobs,
    c.pagado,
    c.saldo,
    c.group_id
FROM
    medi001.citas c
INNER JOIN
    medi001.doctores d ON c.iddoctor = d.id
INNER JOIN
    medi001.pacientes p ON c.cedula = p.cedula
ORDER BY
    c.id;`

const sqlDelCita = `DELETE FROM medi001.citas WHERE id = $1;`

const sqlDelCitaAll = `DELETE FROM medi001.citas WHERE group_id = (SELECT group_id FROM medi001.citas WHERE id = $1) AND inicio >= (SELECT inicio FROM medi001.citas WHERE id = $1);`

const sqlUpdCita = `UPDATE medi001.citas SET 
		iddoctor = $2, 
		cedula = $3, 
		motivo = $4, 
		inicio = $5, 
		fin = $6, 
		status = $7, 
		color = $8,
        montoref = $9,
        tasa = $10,
        montobs = $11,
        pagado = $12,
        group_id = $13
		WHERE id = $1`

const sqlUpdDiagnostico = `UPDATE medi001.citas SET diagnostico = $1, status = 'Completada' WHERE id = $2`

const sqlPostCita = `INSERT INTO medi001.citas (id, iddoctor, cedula, motivo, inicio, fin, status, color, montoref, tasa, montobs, pagado, saldo, group_id) 
		VALUES (nextval('medi001.citas_id_seq'::regclass), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $8, $12);`

const sqlGetCitaPaciente = `SELECT
    c.id AS cita_id,
    c.iddoctor,
    d.nombres AS especialista,
    d.espec AS especialidad,
    c.cedula AS paciente_cedula,
    p.nombres AS paciente,
    c.motivo,
    c.inicio,
    c.fin,
    c.diagnostico,
    c.status AS cita_status,
    c.color,
    c.montoref,
    c.tasa,
    c.montobs,
    c.pagado,
    c.saldo
    ,c.group_id
FROM
    medi001.citas c
INNER JOIN
    medi001.doctores d ON c.iddoctor = d.id
INNER JOIN
    medi001.pacientes p ON c.cedula = p.cedula
where c.cedula = $1
ORDER BY c.id;`

const sqlGetPayments = `SELECT id, appointmentid, paymentmethod, amount, currency, reference, "date", status, notes
FROM medi001.payments;`

const sqlGetPaymentsByCita = `SELECT id, appointmentid, paymentmethod, amount, currency, reference, "date", status, notes
FROM medi001.payments where appointmentid = $1;`

const sqlPostPayments = `INSERT INTO medi001.payments (appointmentid, paymentmethod, amount, currency, 
reference, date, status, notes)
VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

const sqlUpdPayments = `UPDATE medi001.payments SET paymentmethod = $2, amount = $3, currency = $4, reference = $5, "date" = $6 , 
status = $7, notes = $8 WHERE id = $1;`

const sqlDelPayments = `DELETE FROM medi001.payments WHERE id = $1;`

const sqlGetRelPagos = `SELECT 
    d.id AS doctor_id, 
    d.nombres AS doctor_nombre, 
    c.id AS cita_id,
    paci.nombres AS paciente_nombre, 
    pay.date AS fecha_pago,
    SUM(pay.amount) AS monto_cobrado_cita,
    STRING_AGG(DISTINCT pay.paymentmethod, ', ') AS formas_pago,
    c.saldo AS saldo_cita,
    d.tasapago AS porcentaje_pago_doctor, 
    (SUM(pay.amount) * (d.tasapago / 100.0)) AS monto_correspondiente_doctor
FROM medi001.citas c
JOIN medi001.doctores d ON c.iddoctor = d.id
left JOIN medi001.payments pay ON c.id = pay.appointmentid
JOIN medi001.pacientes paci ON c.cedula = paci.cedula
WHERE DATE(c.inicio) BETWEEN $1 AND $2
GROUP BY 
    d.id, 
    d.nombres, 
    paci.nombres, 
    pay.date,
    c.id, 
    d.tasapago, 
    c.saldo
ORDER BY 
    d.nombres, 
    paci.nombres, 
    c.id;`

// `SELECT d.id AS doctor_id, d.nombres AS doctor_nombre, c.id AS cita_id,
//     paci.nombres AS paciente_nombre, SUM(pay.amount) AS monto_cobrado_cita,
//     d.tasapago AS porcentaje_pago_doctor, (SUM(pay.amount) * (d.tasapago / 100.0)) AS monto_correspondiente_doctor
// FROM medi001.citas c
// JOIN medi001.doctores d ON c.iddoctor = d.id
// JOIN medi001.payments pay ON c.id = pay.appointmentid
// JOIN medi001.pacientes paci ON c.cedula = paci.cedula
// WHERE DATE(c.inicio) between $1 and $2
// GROUP BY d.id, d.nombres, paci.nombres, c.id, d.tasapago
// ORDER BY d.nombres, paci.nombres, c.id;`

const sqlAddUsuario = `INSERT INTO seguridad.usuarios(codigo, clave, nombre, idtipouser, idperfil, status, direccion, 
direccion2, ciudad, estado, telf, cel, correo, facebook, whatsapp, instagram, idvendedor)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17);`

const sqlGetPreciosPorPaciente = `SELECT especialidad, precio FROM medi001.paciente_precios_especialidad WHERE id_paciente = $1;`

const sqlUpsertPrecioEspecialidad = `
INSERT INTO medi001.paciente_precios_especialidad (id_paciente, especialidad, precio)
VALUES ($1, $2, $3)
ON CONFLICT (id_paciente, especialidad) DO UPDATE SET precio = EXCLUDED.precio;`

const sqlDelPrecioEspecialidad = `DELETE FROM medi001.paciente_precios_especialidad WHERE id_paciente = $1 AND especialidad = $2;`

const sqlGetCitasFecha = `
	SELECT
    c.id AS cita_id,
    c.iddoctor,
    d.nombres AS especialista,
    d.espec AS especialidad,
    c.cedula AS paciente_cedula,
    p.nombres AS paciente,
    c.motivo,
    c.inicio,
    c.fin,
	c.diagnostico,
    c.status AS cita_status,
    c.color,
    c.montoref,
    c.tasa,
    c.montobs,
    c.pagado,
    c.saldo
    ,c.group_id
FROM
    medi001.citas c
INNER JOIN
    medi001.doctores d ON c.iddoctor = d.id
INNER JOIN
    medi001.pacientes p ON c.cedula = p.cedula
WHERE c.inicio BETWEEN $1 AND $2
	ORDER BY c.inicio`

const sqlUpdateUnpaidAppointmentsVESRate = `
UPDATE medi001.citas
SET tasa = $1,
montobs = montoref * $1
WHERE pagado = 0
  AND status != 'Cancelada'
  AND montoref != 0;
`
const sqlGetAllPacientesData = `SELECT medi001.get_all_pacientes_data($1, $2);`
