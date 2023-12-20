package handlers

import (
	"clase_5_mux_gorm/database"
	"clase_5_mux_gorm/dto"
	"clase_5_mux_gorm/jwt"
	"clase_5_mux_gorm/modelos"
	"clase_5_mux_gorm/validaciones"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Seguridad_protegido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Recurso protegido",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
func Seguridad_login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var registro dto.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&registro); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	if len(registro.Correo) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El E-Mail es obligatorio",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if validaciones.Regex_correo.FindStringSubmatch(registro.Correo) == nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El E-Mail ingresado no es válido",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	if !validaciones.ValidarPassword(registro.Password) {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "La contraseña debe tener al menos 1 número, una mayúscula, y un largo entre 6 y 20 caracteres",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	//p2gHNiENUw
	//validamos si existe el correo informado
	usuario := modelos.Usuario{}
	if database.Database.
		Where("correo = ? ", registro.Correo).
		Limit(1).
		Find(&usuario).RowsAffected > 0 {
		//usamos bcrypt para comparar password
		passwordBytes := []byte(registro.Password)
		passwordBD := []byte(usuario.Password)
		err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
		if err != nil {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "La credenciales ingresadas con inválidas",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(respuesta)
			return
		} else {
			jwtKey, errJwt := jwt.GenerarJWT(usuario)
			if errJwt != nil {
				respuesta := map[string]string{
					"estado":  "error",
					"mensaje": "Ocurrió un error al intentar general el Token correspondiente " + err.Error(),
				}

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(respuesta)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dto.LoginRespuestaDto{
				Nombre: usuario.Nombre,
				Token:  jwtKey,
			})
		}
	} else {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "La credenciales ingresadas con inválidas",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
}
func Seguridad_registro(w http.ResponseWriter, r *http.Request) {
	var registro dto.UsuarioDto
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&registro); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	if len(registro.Nombre) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El Nombre es obligatorio",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	if len(registro.Correo) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El E-Mail es obligatorio",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if validaciones.Regex_correo.FindStringSubmatch(registro.Correo) == nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El E-Mail ingresado no es válido",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	if !validaciones.ValidarPassword(registro.Password) {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "La contraseña debe tener al menos 1 número, una mayúscula, y un largo entre 6 y 20 caracteres",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	// p2gHNiENUw
	// validamos si existe el correo informado
	usuario := modelos.Usuario{}
	if database.Database.
		Where("correo = ? ", registro.Correo).
		Limit(1).
		Find(&usuario).RowsAffected > 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El E-Mail " + registro.Correo + " ya está siendo usado por otro usuario",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	} else {
		// generamos hash con bcrypt
		costo := 8
		bytes, _ := bcrypt.GenerateFromPassword([]byte(registro.Password), costo)
		datos := modelos.Usuario{Nombre: registro.Nombre, Correo: registro.Correo, Telefono: registro.Telefono, PerfilID: registro.PerfilID, Password: string(bytes), Fecha: time.Now()}
		database.Database.Save(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se creó el registro exitosamente",
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(respuesta)
	}
}
