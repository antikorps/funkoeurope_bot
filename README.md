# FunkoEurope Bot - Proyecto Experimental
# Importante
Este proyecto tiene caracter experimental. Surge de la petición de un particular, pero no soy usuario de la web y desconozco su funcionamiento, por lo que me resulta muy difícil toda la fase de testeo.
# Propósito
La idea del bot es monitorizar la URL de un producto e incorporarlo automáticamente al carrito de la compra del usuario cuando haya stock.
# Uso
El bot se ejecuta desde el binario que se encuentra en la carpeta bin. Para su funcionamiento es necesario que en la misma carpeta del ejecutable se encuentre un archivo de configuración y de cookies, ambos en formato json.

El archivo configuracion.json debe tener la siguiente estructura:

```json
{
    "url": "https://funkoeurope.com/products/soba-mask-one-piece#",
    "cantidad": 1,
    "espera_ms": 2000
}
```
El archivo cookies.json debe crearse con la extensión [Cookie Quick Manager ](https://addons.mozilla.org/es/firefox/addon/cookie-quick-manager/) y tendrá una estructura del siguiente tipo:
```json
[{
	"Host raw": "http://.funkoeurope.com/",
	"Name raw": "_landing_page",
	"Path raw": "/",
	"Content raw": "%2F",
	"Expires": "12-08-2023 10:14:38",
	"Expires raw": "1691828078",
	"Send for": "Any type of connection",
	"Send for raw": "false",
	"HTTP only raw": "true",
	"SameSite raw": "lax",
	"This domain only": "Valid for subdomains",
	"This domain only raw": "false",
	"Store raw": "firefox-default",
	"First Party Domain": ""
}]
```
Una vez se disponga de los archivos simplemente es necesario ejecutar el binario desde la línea de comandos.

Primeramente hará una petición para comprobar si las cookies incorporadas permiten la identificación del usuario.

Posteriormente hará una petición a la URL incorporada para buscar el id y el productId.

Finalmente entrará en un bucle infinito en el que intentará incorparar el producto. Si finalmente consigue añadirlo al carrito de compra acabará su ejecución. Se puede configurar el tiempo de espera en milisegundos entre cada intento desde el archivo configuracion.json

# Instrucciones y ejemplo de funcionamiento
![Instrucciones y ejemplo de funcionamiento](https://i.imgur.com/9BqBy64.gif )