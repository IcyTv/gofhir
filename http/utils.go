package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gofhir/generated"
)

func Error(c *gin.Context, err generated.OperationOutcome) {
	//TODO figure out which format we want to respond with
	c.JSON(400, err)
}

// ======
// Functions that generate OperationOutcomes for the most common errors
// ======

func NotFound(c *gin.Context) {
	Error(c, generated.OperationOutcome{
		Issue: &generated.OperationOutcomeIssue{
			Severity:    "error",
			Code:        "not-found",
			Diagnostics: "The resource could not be found",
		},
	})
}

func InvalidFormat(c *gin.Context) {
	Error(c, generated.OperationOutcome{
		Issue: &generated.OperationOutcomeIssue{
			Severity:    "error",
			Code:        "invalid",
			Diagnostics: "The format specified is not supported",
		},
	})
}

func InvalidBundleType(c *gin.Context) {
	Error(c, generated.OperationOutcome{
		Issue: &generated.OperationOutcomeIssue{
			Severity:    "error",
			Code:        "invalid",
			Diagnostics: "The bundle type specified is not supported, should be either 'batch' or 'transaction'",
		},
	})
}
