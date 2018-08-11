package main

import (
	"flag"
	"os"
	"fmt"

	"github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrgin/v1"
	"github.com/gin-gonic/gin"

	"github.com/urbn/ordernumbergenerator/app"
	"github.com/urbn/ordernumbergenerator/app/clients"
	"github.com/urbn/ordernumbergenerator/app/daos"
	"github.com/urbn/ordernumbergenerator/app/config"
	"github.com/urbn/ordernumbergenerator/app/handlers"
)

var port *string

func init() {
	var err error
	port = flag.String("port", "7070", "Webserver Port to listen") //TODO change port to be unique to ordernum?
	flag.Parse()

	app.Configuration, err = config.LoadConfig()
	if err != nil {
		ErrorExit(err)
	}
}

func main() {
	session := clients.CreateMongoSession(app.Configuration.MongoHost)
	collection := session.DB(app.MongoDatabase).C(app.MongoOrdersCollection)

	dao := daos.CreateOrderNumberDao(session, collection)

	router := gin.Default()

	if app.Configuration.AppName != "" && app.Configuration.NewRelicLicense != "" {
		newRelicConfig := newrelic.NewConfig(app.Configuration.AppName, app.Configuration.NewRelicLicense)
		newRelicApp, err := newrelic.NewApplication(newRelicConfig)
		if err != nil {
			println("NewRelic initialization error occurred")
		} else {
			router.Use(nrgin.Middleware(newRelicApp))
		}
	}

	var hh = handlers.HealthHandler{
		app.Configuration,
	}

	var nh = handlers.OrderNumberHandler{
		DataCenterId: app.Configuration.DatacenterId,
		Dao:          dao,
	}

	router.GET("/health", hh.GetHealth)
	router.POST("/v0/:siteId/sterling-order-number", nh.GetOrderNumber)

	println("Webserver ready and listening on Port", *port)

	err := router.Run(":7070")
	if err != nil {
		panic(fmt.Sprintf("Webserver stopped with error %s, quitting application", err.Error()))
	}
}

func ErrorExit(err error) {
	fmt.Printf("Application initialization Error occurred: %s", err.Error())
	os.Exit(1)
}
