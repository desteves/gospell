package main

import (
	"github.com/desteves/gospell/api"
	"github.com/desteves/gospell/speller"
	"github.com/gin-gonic/gin"
)

func main() {
	// for now use hardcoded, internal phonetic alphabet
	s, err := speller.NewNATOProvider("internal")
	if err != nil {
		panic(err)
	}

	// set the api to use the internal speller
	a := api.API{
		SpellerProvider: s,
	}

	// http server
	router := gin.Default()
	v1 := router.Group("/v1")

	{
		v1.GET(api.SpellEndpoint, a.SpellEndpointHandler)
		v1.GET(api.HealthcheckEndpoint, func(c *gin.Context) {
			c.JSON(200, "pong")
		})
	}

	router.Run(":8080")

}
