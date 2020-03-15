package main

import (
	"fmt"
	"net/http"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-2/database"
	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-2/models"
	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-2/routes"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {

	log.Info("Create new cron")

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {

		log.Info("[Job 1]Every second job\n")

		message := models.CheckProductQuantity()

		fmt.Println(message)

	})

	//Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()

	database.GetConnection()
	database.DbConnection()
	database.Db.Ping()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(routes.MainRoutes())

	server := &http.Server{
		Addr:    "0.0.0.0:8006",
		Handler: n,
	}

	log.Info("server Running")

	server.ListenAndServe()

}
