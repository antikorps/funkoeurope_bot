package manejador

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func (m *Manejador) BuscarId() {
	peticionProducto, peticionError := http.NewRequest("GET", m.Configuracion.URL, nil)
	if peticionError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido preparar la petición para buscar el id", peticionError)
	}
	respuesta, respuestaError := m.Cliente.Do(peticionProducto)
	if respuestaError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener respuesta de la petición para obtener el id", respuestaError)
	}
	if respuesta.StatusCode != 200 {
		log.Fatalln("ERROR CRÍTICO: se ha obtenido un status code incorrecto en la petición para obtener el id", respuesta.Status)
	}
	defer respuesta.Body.Close()

	html, htmlError := io.ReadAll(respuesta.Body)
	if htmlError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener el html de la respuesta para obtener el id", htmlError)
	}
	regexMin := regexp.MustCompile(`\n`)
	html = regexMin.ReplaceAll(html, []byte(""))
	regexDataVariant := regexp.MustCompile(`.*?<span class="id">(.*?)</span>.*`)
	id := regexDataVariant.ReplaceAll(html, []byte("$1"))
	idRecuperado := string(id)

	regexProductId := regexp.MustCompile(`.*?<span class="product_id">(.*?)</span>.*`)
	productoId := regexProductId.ReplaceAll(html, []byte("$1"))
	productoIdRecuperado := string(productoId)

	if idRecuperado == "" {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener el id de la url introducida")
	}
	if productoIdRecuperado == "" {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener el id del producto de la url introducida")
	}
	if len(idRecuperado) > 20 || len(productoId) > 20 {
		log.Fatalln("ERROR CRÍTICO: ha fallado la recuperación del id o del producto id")
	}
	m.Id = idRecuperado
	m.IdProducto = productoIdRecuperado
}
