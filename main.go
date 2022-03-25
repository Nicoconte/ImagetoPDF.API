package main

import (
	"imagetopdf/routes"
	"imagetopdf/services"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(3).Minutes().Do(services.DeleteAllSessions)
	cron.StartAsync()

	http.ListenAndServe("localhost:8080", routes.Router)
}
