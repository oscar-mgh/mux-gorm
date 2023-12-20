package jwt

import (
	"clase_5_mux_gorm/modelos"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerarJWT(usuario modelos.Usuario) (string, error) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic("Error loading .env file")
	}
	miClave := []byte(os.Getenv("SECRET_JWT"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"correo":         usuario.Correo,
		"nombre":         usuario.Nombre,
		"generado_desde": "https://www.cesarcancino.com",
		"id":             usuario.Id,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Hour * 24).Unix(), //24 HORAS
	})
	tokenString, err := token.SignedString(miClave)
	return tokenString, err
}
