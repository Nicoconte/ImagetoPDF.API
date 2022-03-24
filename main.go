package main

import (
	"imagetopdf/routes"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:8080", routes.Router)
}
