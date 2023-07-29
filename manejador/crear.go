package manejador

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func Crear() Manejador {
	rutaEjecutable, rutaEjecutableError := os.Executable()
	if rutaEjecutableError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener la ruta del ejecutable", rutaEjecutableError)
	}
	rutaRaiz := filepath.Dir(rutaEjecutable)

	rutaConfiguracion := filepath.Join(rutaRaiz, "configuracion.json")
	archivoConfiguracion, archivoConfiguracionError := os.ReadFile(rutaConfiguracion)
	if archivoConfiguracionError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido leer el archivo configuracion.json", archivoConfiguracionError)
	}
	var configuracion Configuracion
	configuracionJsonError := json.Unmarshal(archivoConfiguracion, &configuracion)
	if configuracionJsonError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido parsear el contenido de configuracion.json", configuracionJsonError)
	}

	rutaArchivoCookies := filepath.Join(rutaRaiz, "cookies.json")
	archivoCookies, archivoCookiesError := os.ReadFile(rutaArchivoCookies)
	if archivoCookiesError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido leer el archivo cookies.json", archivoCookiesError)
	}
	var cookiesNavegador CookiesNavegador
	cookiesNavegadorError := json.Unmarshal(archivoCookies, &cookiesNavegador)
	if cookiesNavegadorError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido parsear el contenido de cookies.json", cookiesNavegadorError)
	}

	var cookies []*http.Cookie
	for _, v := range cookiesNavegador {
		var httpOnly bool
		if v.HTTPOnlyRaw == "true" {
			httpOnly = true
		}
		cookies = append(cookies, &http.Cookie{
			Name:       v.NameRaw,
			Value:      v.ContentRaw,
			Path:       v.PathRaw,
			RawExpires: v.ExpiresRaw,
			HttpOnly:   httpOnly,
		})
	}

	jar, jarError := cookiejar.New(nil)
	if jarError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido crear el cookier jar", jarError)
	}
	url, _ := url.Parse("https://funkoeurope.com")
	jar.SetCookies(url, cookies)

	cliente := http.Client{
		Jar:     jar,
		Timeout: 12 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return Manejador{
		Cliente:       cliente,
		Configuracion: configuracion,
	}

}
