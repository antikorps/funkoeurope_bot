package manejador

import "net/http"

type Manejador struct {
	Cliente             http.Client
	Configuracion       Configuracion
	Id                  string
	IdProducto          string
	ProductoIncorporado bool
}

type CookiesNavegador []struct {
	HostRaw           string `json:"Host raw"`
	NameRaw           string `json:"Name raw"`
	PathRaw           string `json:"Path raw"`
	ContentRaw        string `json:"Content raw"`
	Expires           string `json:"Expires"`
	ExpiresRaw        string `json:"Expires raw"`
	SendFor           string `json:"Send for"`
	SendForRaw        string `json:"Send for raw"`
	HTTPOnlyRaw       string `json:"HTTP only raw"`
	SameSiteRaw       string `json:"SameSite raw"`
	ThisDomainOnly    string `json:"This domain only"`
	ThisDomainOnlyRaw string `json:"This domain only raw"`
	StoreRaw          string `json:"Store raw"`
	FirstPartyDomain  string `json:"First Party Domain"`
}

type Configuracion struct {
	URL      string `json:"url"`
	EsperaMs int    `json:"espera_ms"`
	Cantidad int    `json:"cantidad"`
}
