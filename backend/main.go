package main

import (
	"log"
	"net/http"

	common "github.com/alx-t/cookbook/backend/common"
	"github.com/alx-t/cookbook/backend/routes"
	"github.com/codegangsta/negroni"
)

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}

	// defer common.Db.Close()

	log.Println("Listening...")
	server.ListenAndServe()
}
