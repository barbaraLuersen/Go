package main

import (
	"Go/route"
	"net/http"
)

func main() {
	route.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
