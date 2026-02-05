package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenidos a Control de Citas")
	}
}

// Health Check Endpoint - Versión mejorada
func (a *App) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configurar headers
		w.Header().Set("Content-Type", "application/json")

		// Verificar conexión a DB
		if err := a.DB.Ping(); err != nil {
			utils.CreateLog(fmt.Sprintf("Health Check FAILED: %v", err))
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":    "DOWN",
				"error":     err.Error(),
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			})
			return
		}

		// Verificar otros servicios si es necesario (ej: Redis, APIs externas)
		// ...

		// Respuesta exitosa
		utils.CreateLog("Health Check OK")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "UP",
			"services":  []string{"database"},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	}
}

func (a *App) MenuWeb() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Handle("/menuweb", http.FileServer(http.Dir("app/templates")))
		// http.FileServer(http.Dir("app/templates"))
		// template, _ := template.ParseFiles("app/templates/index.html")
		// template.Execute(w, nil)
	}
}

func (a *App) ConfigVar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Handle("/configvar", http.FileServer(http.Dir("app/templates")))
		http.FileServer(http.Dir("app/templates"))
		template, _ := template.ParseFiles("app/templates/configuracion.html")
		template.Execute(w, nil)
	}
}

func (a *App) GetInventario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespInventario
		datos, err := a.DB.GetInventarios(0)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Inventario"
		}
		rp.Status = 10
		rp.Mensaje = "Inventario listado correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetInventarioFormal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespInventario
		datos, err := a.DB.GetInventarios(2)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Inventario"
		}
		rp.Status = 10
		rp.Mensaje = "Inventario listado correctamente!"
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetInventarioCompacto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespInventario
		datos, err := a.DB.GetInventarios(1)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Inventario"
		}
		rp.Status = 10
		rp.Mensaje = "Inventario listado correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetInventarioMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespInventarioNombre
		datos, err := a.DB.GetInventarioMenu()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Inventario"
		}
		rp.Status = 10
		rp.Mensaje = "Inventario listado correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
		// fmt.Fprintf(w, data)
	}
}

func (a *App) AddInventario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Inventario
		// err :=
		json.NewDecoder(r.Body).Decode(&p)
		// if err != nil {
		// 	// log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddInventario(p)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) AddItemsInventario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p []models.ItemsInventario

		json.NewDecoder(r.Body).Decode(&p)
		// if err != nil {
		// 	// log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddItemsInventario(p)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) AddPresentaciones() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p []models.PresenInventario

		json.NewDecoder(r.Body).Decode(&p)
		// if err != nil {
		// 	// log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddPresenInventario(p)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) AddPresenInventario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p []models.PresenInventario

		json.NewDecoder(r.Body).Decode(&p)
		rp := a.DB.AddPresenInventario(p)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdInventario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var p models.Inventario

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			// fmt.Println(respo.Mensaje)
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdInventario(p)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelInventario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Inventario
		var resp models.RespInventario
		// err :=
		json.NewDecoder(r.Body).Decode(&p)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.DelInventario(p)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar el Producto en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespMenuLista
		datos, err := a.DB.GetMenu()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Menu"
		}
		rp.Status = 10
		rp.Mensaje = "Menu listado correctamente!"
		rp.Items = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
		// fmt.Fprintf(w, data)
	}
}

func (a *App) GetMenuClases() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespMenu
		datos, err := a.DB.GetMenuListado()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Menu"
		}
		rp.Status = 10
		rp.Mensaje = "Menu listado correctamente!"
		rp.Items = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetClases() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespClase
		datos, err := a.DB.GetClases()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Clases"
		}
		rp.Status = 10
		rp.Mensaje = "Clases listadas correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) AddClase() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Clase
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			// log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.AddClase(c)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdClase() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var c models.Clase
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdClase(c)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelClase() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Clase
		var resp models.RespInventario
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			// log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.DelClase(c)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar la clase en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) AddMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m models.Menu
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			// log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.AddMenu(m)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) AddAllMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m []models.Menu

		// fmt.Println("datos de la peticion http")
		// err :=
		json.NewDecoder(r.Body).Decode(&m)
		// fmt.Println(len(m))
		// fmt.Println(m[0].Nombre)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddAllMenu(m)
		// fmt.Println("Status")
		// fmt.Println(rp.Status)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var m models.Menu
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdMenu(m)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m models.Menu
		var resp models.RespMenu
		// err :=
		json.NewDecoder(r.Body).Decode(&m)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.DelMenu(m)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar la clase en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) DelAllMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m models.Menu
		var resp models.RespMenu
		// err :=
		json.NewDecoder(r.Body).Decode(&m)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.DelAllMenu()
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar el menu completo en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetEmpre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespEmpre
		datos, err := a.DB.GetEmpre()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Clases"
		}
		rp.Status = 10
		rp.Mensaje = "Empresas listadas correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) AddEmpresa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Empresa
		// err :=
		json.NewDecoder(r.Body).Decode(&e)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddEmpre(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdEmpresa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var e models.Empresa
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdEmpresa(e)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelEmpresa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Empresa
		var resp models.RespMenu
		// err :=
		json.NewDecoder(r.Body).Decode(&e)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.DelEmpresa(e)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar la empresa en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetMenuCompleto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespMenuCompleto
		datos, err := a.DB.GetMenuCompleto()
		// log.Println(datos)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Menu"
		}
		rp.Status = 10
		rp.Mensaje = "Menu Listado Completo!"
		rp.Data = datos.Data
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) Images() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/form-data")

		// Abrir la carpeta de imágenes
		imgPath := "app/assets/images"
		files, err := os.ReadDir(imgPath)
		if err != nil {
			// fmt.Println(err)
			return
		}

		// Crear un escritor multipart
		writer := multipart.NewWriter(w)
		defer writer.Close()

		for _, file := range files {
			filePath := filepath.Join(imgPath, file.Name())
			// fmt.Println("Archivo ", filePath)
			fileHandler, err := os.Open(filePath)
			if err != nil {
				// fmt.Println(err)
				return
			}
			defer fileHandler.Close()

			// Añadir el archivo adjunto al escritor multipart
			part, err := writer.CreateFormFile(file.Name(), file.Name())
			if err != nil {
				// fmt.Println(err)
				return
			}

			// Copiar el contenido del archivo al archivo adjunto
			_, err = io.Copy(part, fileHandler)
			if err != nil {
				// fmt.Println(err)
				return
			}
		}
	}
}

func (a *App) GetPrefacturas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespPrefactura
		datos, err := a.DB.GetPrefacturas()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Clases"
		}
		rp.Status = 10
		rp.Mensaje = "Facturas listadas correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetPrefactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Prefactura
		var rp models.RespPrefactura

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.GetPrefactura(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando PreFacturas"
		}
		rp.Status = 10
		rp.Mensaje = "Factura listada correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) PostPreFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Prefactura
		// log.Println("ingresando a funcion PostPrefactura. codificando datos...")
		json.NewDecoder(r.Body).Decode(&e)
		// log.Println("datos codificados")
		// log.Println(e)
		// log.Println("preparando para ingresar a funcion PostPrefactura en la base de datos")
		p, rp := a.DB.PostPrefactura(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, p, http.StatusOK)
		}
	}
}

func (a *App) UpdPreFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Prefactura

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.UpdPrefactura(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetClientes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespClientes
		datos, err := a.DB.GetClientes()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Clientes"
		}
		rp.Status = 10
		rp.Mensaje = "Clientes listados correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) AddCliente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Cliente
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.AddCliente(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdCliente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var e models.Cliente
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdCliente(e)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelCliente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Cliente
		var resp models.RespMenu
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.DelCliente(e)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar el cliente en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetMesas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespMesas
		datos, err := a.DB.GetMesas()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Mesas"
		}
		rp.Status = 10
		rp.Mensaje = "Mesas listadas correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) AddMesa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Mesas
		// err :=
		json.NewDecoder(r.Body).Decode(&e)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddMesa(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdMesa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var e models.Mesas
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdMesa(e)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelMesa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Mesas
		var resp models.RespMenu
		// err :=
		json.NewDecoder(r.Body).Decode(&e)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.DelMesa(e)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar la mesa en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetMesoneros() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var rp models.RespMesas
		datos, err := a.DB.GetMesoneros()
		if err.Status > 200 {
			// rp.Status = err.Status
			// rp.Mensaje = "Error Cargando Mesas"
			fmt.Fprintf(w, "Cannot format json, err=%v/n", err.Mensaje)
		}
		// rp.Status = 10
		// rp.Mensaje = "Mesas listadas correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) AddMesonero() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Mesoneros
		// err :=
		json.NewDecoder(r.Body).Decode(&e)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AddMesonero(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdMesonero() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var e models.Mesoneros
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdMesonero(e)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) DelMesonero() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Mesoneros
		var resp models.RespMenu
		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.DelMesonero(e)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar la mesa en la base de datos. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) AbrirMesa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Mesas
		// err :=
		json.NewDecoder(r.Body).Decode(&e)
		// if err != nil {
		// 	log.Printf("Unable to decode the request body.  %v", err)
		// }
		rp := a.DB.AbrirMesa(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetInstrumentos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var rp models.RespMesas
		datos, err := a.DB.GetInstrumentos()
		if err.Status > 200 {
			// rp.Status = err.Status
			// rp.Mensaje = "Error Cargando Mesas"
			fmt.Fprintf(w, "Cannot format json, err=%v/n", err.Mensaje)
		}

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) cleanMesa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var e models.Mesas
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.CleanMesa(e)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) PostFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var e models.IdFactura

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.PostFactura(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PostNotaEntrega() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var e models.IdFactura

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.PostNotaEntrega(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PostPresupuesto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Id
		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.PostPresupuesto(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetFacturas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespFactura
		datos, err := a.DB.GetFacturas()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Facturas"
		}
		rp.Status = 10
		rp.Mensaje = "Facturas listadas correctamente!"
		// rp.Data = datos
		// log.Println(datos)
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetPresupuestos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespFactura
		datos, err := a.DB.GetPresupuestos()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Presupuestos"
		}
		rp.Status = 10
		rp.Mensaje = "Presupuestos listados correctamente!"
		// rp.Data = datos
		// log.Println(datos)
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetNotasEntrega() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespFactura
		datos, err := a.DB.GetNotasEntrega()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Notas de Entrega"
		}
		rp.Status = 10
		rp.Mensaje = "Notas de Entrega listadas correctamente!"
		// rp.Data = datos
		// log.Println(datos)
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetDivisas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetDivisas()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Divisas"
		}
		rp.Status = 10
		rp.Mensaje = "Facturas listadas correctamente!"
		// rp.Data = datos
		// log.Println(datos)
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetPreFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Prefactura
		var rp models.RespPrefactura

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.GetPrefactura(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando PreFacturas"
		}
		rp.Status = 10
		rp.Mensaje = "Factura listada correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetFacturaId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Factura
		var rp models.RespFactura

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.GetFacturaId(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Facturas"
		}
		rp.Status = 10
		rp.Mensaje = "Factura listada correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetPresupuestoId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Presupuestos
		var rp models.RespFactura

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.GetPresupuestoId(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Facturas"
			utils.CreateLog(err.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Factura listada correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) DelPresupuesto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Id
		// var resp models.RespPresupuestos
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.DelPresupuesto(f)
		if rp.Status >= 400 {
			rp.Status = 40
			rp.Mensaje = "No se pudo eliminar el presupuesto en la base de datos. Error=%v \n"
			// rp.Data = []
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		rp.Status = 201
		rp.Mensaje = "Presupuesto eliminado correctamente!"
		sendResponse(w, r, rp, http.StatusOK)
	}
}

func (a *App) UpdAnularFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var c models.Factura
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdAnularFactura(c)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) PostVentasFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f database.DT
		var rp models.RespFactura

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.PostVentasFactura(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Ventas"
		}
		rp.Status = 10
		rp.Mensaje = "Ventas listadas correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) PostVentasProductos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f database.DT
		var rp models.RespProductos

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.PostVentasProductos(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Ventas"
		}
		rp.Status = 10
		rp.Mensaje = "Ventas listadas correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) UpdDivisa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var c models.Divisas
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdDivisa(c)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) PostPagos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Pagos

		// log.Println("handler postpagos")
		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.PostPagos(e)
		log.Println(rp.Status)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetDetPagosFecha() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		var f models.Fechas

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.GetDetPagosFecha(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Pagos"
		}
		rp.Status = 10
		rp.Mensaje = "Pagos listados correctamente!"
		// rp.Data = datos
		// log.Println(datos)

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetResumenDetPagos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		var f models.Fechas

		json.NewDecoder(r.Body).Decode(&f)
		datos, err := a.DB.GetResDetPagos(f)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Pagos"
		}
		rp.Status = 10
		rp.Mensaje = "Pagos listados correctamente!"
		// rp.Data = datos
		// log.Println(datos)

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) SendVentasMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.MailSend
		json.NewDecoder(r.Body).Decode(&f)
		a.DB.SendMail(f)
	}
}

func (a *App) GetEmailConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetEmailConfig()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Configuracion"
		}
		rp.Status = 10
		rp.Mensaje = "Parametros listados correctamente!"
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) PutEmailConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var c models.EmailConfig

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdEmailConfig(c)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) PostEmailConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.EmailConfig
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.AddEmailConfig(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelEmailConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Id
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.DelEmailConfig(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetVendedores() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetVendedores()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Vendedores"
		}
		rp.Status = 10
		rp.Mensaje = "Vendedores listados correctamente!"
		// rp.Data = datos
		// log.Println(datos)
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetTopVentas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespTopVentas
		datos, err := a.DB.GetTopVentas()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Ventas"
		}
		rp.Status = 10
		rp.Mensaje = "Ventas listadas correctamente!"
		rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetVentasMes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespVentasMensual
		datos, err := a.DB.GetVentasMes()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Ventas"
		}
		rp.Status = 10
		rp.Mensaje = "Ventas listadas correctamente!"
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetVentas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespVentasMensual
		var c models.Fechas

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			rp.Status = 502
			rp.Mensaje = "Unable to decode the request body. " + err.Error()
			utils.CreateLog("Unable to decode the request body. " + err.Error())
			sendResponse(w, r, rp, http.StatusInternalServerError)
		}
		datos, resp := a.DB.GetVentas(c)
		if resp.Status > 200 {
			rp.Status = resp.Status
			rp.Mensaje = "Error Cargando Ventas"
			utils.CreateLog("Error Cargando Ventas " + resp.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Ventas listadas correctamente!"
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			utils.CreateLog("Cannot format json")
			return
		}
	}
}

func (a *App) ClearLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		utils.ClearLog()
		rp.Status = 10
		rp.Mensaje = "Se vacio el archivo log!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(rp)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		archivo := "log.txt"
		http.ServeFile(w, r, archivo)
	}
}

func (a *App) GetProveedores() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.RespProveedores
		datos, err := a.DB.GetProveedores()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Proveedores"
		}
		rp.Status = 10
		rp.Mensaje = "Proveedores listados correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetCxcResumen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Fechas
		json.NewDecoder(r.Body).Decode(&e)
		var rp models.Respuesta
		datos, err := a.DB.GetCxcResumen(e)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Ventas"
		}
		rp.Status = 10
		rp.Mensaje = "Ventas listadas correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetCxcVencida() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetCxcVencida()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando CxC"
		}
		rp.Status = 10
		rp.Mensaje = "Cuentas por Cobrar listadas correctamente!"
		// rp.Data = datos
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetCompras() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta

		datos, err := a.DB.GetCompras()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando compras"
		}
		rp.Status = 10
		rp.Mensaje = "compras listadas correctamente!"
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetParametros() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		datos, err := a.DB.GetParametros()
		if err.Status > 200 {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", err.Mensaje)
		}
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) PostParametro() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Parametros

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.AddParametro(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PutParametro() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respo models.Respuesta
		var c models.Parametros
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			respo.Status = 502
			respo.Mensaje = "Unable to decode the request body. " + err.Error()
			sendResponse(w, r, respo, http.StatusInternalServerError)
		}
		resp := a.DB.UpdParametro(c)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, resp, http.StatusOK)
		}
	}
}

func (a *App) PostProveedor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Proveedor

		// log.Println("handler postpagos")
		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.PostProveedor(e)
		// log.Println(rp.Status)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PutProveedor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Proveedor

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.UpdProveedor(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelProveedor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Id

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.DelProveedor(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PostCompra() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Compra

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.AddCompra(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelCompra() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Id

		json.NewDecoder(r.Body).Decode(&e)
		rp := a.DB.DelCompra(e)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}
