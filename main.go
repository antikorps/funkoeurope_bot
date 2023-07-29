package main

import (
	"funkoeurope_bot/manejador"
	"log"
	"time"
)

func main() {
	manejador := manejador.Crear()
	manejador.ValidarCookies()
	manejador.BuscarId()

	var intentosCompra int
	for !manejador.ProductoIncorporado {
		intentosCompra++
		log.Println("INFO: intento", intentosCompra, "de añadir el producto al carrito")
		manejador.IntentarCompra()
		time.Sleep(time.Duration(manejador.Configuracion.EsperaMs) * time.Millisecond)
	}

	log.Println("ÉXITO: producto añadido al carrito")
}
