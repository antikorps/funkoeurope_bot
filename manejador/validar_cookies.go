package manejador

import (
	"log"
	"net/http"
)

func (m *Manejador) ValidarCookies() {
	peticion, peticionError := http.NewRequest("GET", "https://funkoeurope.com/account", nil)
	if peticionError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido preparar la petición para validar las cookies incorporadas", peticionError)
	}

	respuesta, respuestaError := m.Cliente.Do(peticion)
	if respuestaError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener la respuesta que permite validar las cookies incorporadas", respuestaError)
	}
	defer respuesta.Body.Close()
	if respuesta.StatusCode != 200 {
		log.Fatalln("ERROR CRÍTICO: las cookies incorporadas no han permitido la autentificación del usuario, la petición de validación ha devuelto un status code de", respuesta.Status)
	}
}
