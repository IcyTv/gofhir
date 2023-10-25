package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofhir/generated"
)

func BatchRoutes(r *gin.Engine) {

	// https://www.hl7.org/fhir/http.html#transaction
	r.POST("/", func(c *gin.Context) {
		var bundle generated.Bundle

		//TODO i think you can specify the format with the Content-Type header as well
		format := c.DefaultQuery("_format", "json")

		var bindingType binding.Binding

		if format == "json" {
			bindingType = binding.JSON
		} else if format == "xml" {
			bindingType = binding.XML
		} else {
			return
		}

		if err := c.ShouldBindWith(&bundle, bindingType); err != nil {
			InvalidFormat(c)
			return
		}

		if bundle.Type == "batch" {

		} else if bundle.Type == "transaction" {

		} else {
			InvalidBundleType(c)
		}

		c.JSON(200, gin.H{
			"message": "batch",
		})
	})
}
