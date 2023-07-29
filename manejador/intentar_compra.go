package manejador

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func (m *Manejador) IntentarCompra() {
	var dataCarrito = strings.NewReader(`{"form_type":"product","utf8":"✓","id":"` + m.Id + `","productId":"` + m.IdProducto + `","quantity":"` + fmt.Sprint(m.Configuracion.Cantidad) + `"}`)
	peticionCarrito, peticionCarritoError := http.NewRequest("POST", "https://funkoeurope.com/cart/add.js", dataCarrito)
	if peticionCarritoError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido crear la petición para añadir el producto al carrito", peticionCarritoError)
	}
	peticionCarrito.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0")
	peticionCarrito.Header.Set("Accept", "*/*")
	peticionCarrito.Header.Set("Accept-Language", "es-ES,es;q=0.8,en-US;q=0.5,en;q=0.3")
	peticionCarrito.Header.Set("Referer", m.Configuracion.URL)
	peticionCarrito.Header.Set("Content-Type", "application/json")
	peticionCarrito.Header.Set("X-Requested-With", "XMLHttpRequest")
	peticionCarrito.Header.Set("Origin", "https://funkoeurope.com")
	peticionCarrito.Header.Set("Connection", "keep-alive")
	peticionCarrito.Header.Set("Sec-Fetch-Dest", "empty")
	peticionCarrito.Header.Set("Sec-Fetch-Mode", "cors")
	peticionCarrito.Header.Set("Sec-Fetch-Site", "same-origin")
	peticionCarrito.Header.Set("Pragma", "no-cache")
	peticionCarrito.Header.Set("Cache-Control", "no-cache")
	peticionCarrito.Header.Set("TE", "trailers")
	respuestaCarrito, respuestaCarritoError := m.Cliente.Do(peticionCarrito)
	if respuestaCarritoError != nil {
		log.Fatalln("ERROR CRÍTICO: no se ha podido obtener respuesta de la petición para añadir al carrito", respuestaCarritoError)
	}
	defer respuestaCarrito.Body.Close()
	if respuestaCarrito.StatusCode == 200 {
		m.ProductoIncorporado = true
	}
	if respuestaCarrito.StatusCode == 429 {
		log.Println("ADVERTENCIA: se ha rechazado la petición por un exceso de peticiones. Por seguridad, se esperarán 5 segundos antes de realizar la próxima petición")
		time.Sleep(5 * time.Second)
	}
}
