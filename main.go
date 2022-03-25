package main

import (
	"imagetopdf/routes"
	"net/http"
)

func main() {
	//Cron will be setup here
	http.ListenAndServe("localhost:8080", routes.Router)
}
