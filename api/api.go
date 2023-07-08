// Package api contains all http-related code
package api

import (
	"net/http"

	"github.com/desteves/gospell/speller"
	"github.com/gin-gonic/gin"
)

type API struct {
	SpellerProvider speller.Speller
}

// error messages
const (
	ErrorMsgStatusBadRequest          = "invalid request"
	ErrorMsgStatusNotFound            = "not found"
	ErrorMsgStatusInternalServerError = "something went wrong"
)

// endpoints
const (
	SpellEndpoint       = "/spell"
	HealthcheckEndpoint = "/ping"
)

type SpellQueryParams struct {
	Input string `json:"input" form:"input"`
}

func (a *API) SpellEndpointHandler(c *gin.Context) {
	var parameters SpellQueryParams

	if c.Request.URL.Query().Get("input") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": ErrorMsgStatusBadRequest})
	}

	err := c.BindQuery(&parameters)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": ErrorMsgStatusBadRequest})
	}

	bodyResponse, err := a.SpellerProvider.GetCodeWords(parameters.Input)
	if err != nil {
		if err.Error() == ErrorMsgStatusNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": ErrorMsgStatusNotFound})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": ErrorMsgStatusInternalServerError})
	}
	c.JSON(http.StatusOK, bodyResponse)
}
