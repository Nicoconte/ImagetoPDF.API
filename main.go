package main

import (
	"imagetopdf/data"
	"imagetopdf/routes"
	"imagetopdf/services"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(1).Hour().Do(services.DeleteAllSessions)
	cron.StartAsync()

	http.ListenAndServe(data.Config.Host, routes.Router)
}
