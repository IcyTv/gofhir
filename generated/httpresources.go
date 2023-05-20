package generated

import (
	"context"
	"fmt"
	gin "github.com/gin-gonic/gin"
	db "github.com/gofhir/db"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AccountRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Account")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Account")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Account{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Account"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Account{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Account")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Account{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Account")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ActivityDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ActivityDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ActivityDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ActivityDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ActivityDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ActivityDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ActivityDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ActivityDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ActivityDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ActorDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ActorDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ActorDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ActorDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ActorDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ActorDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ActorDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ActorDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ActorDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func AdministrableProductDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/AdministrableProductDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("AdministrableProductDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := AdministrableProductDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "AdministrableProductDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := AdministrableProductDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("AdministrableProductDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := AdministrableProductDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("AdministrableProductDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func AdverseEventRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/AdverseEvent")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("AdverseEvent")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := AdverseEvent{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "AdverseEvent"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := AdverseEvent{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("AdverseEvent")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := AdverseEvent{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("AdverseEvent")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func AllergyIntoleranceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/AllergyIntolerance")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("AllergyIntolerance")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := AllergyIntolerance{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "AllergyIntolerance"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := AllergyIntolerance{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("AllergyIntolerance")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := AllergyIntolerance{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("AllergyIntolerance")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func AppointmentRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Appointment")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Appointment")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Appointment{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Appointment"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Appointment{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Appointment")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Appointment{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Appointment")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func AppointmentResponseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/AppointmentResponse")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("AppointmentResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := AppointmentResponse{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "AppointmentResponse"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := AppointmentResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("AppointmentResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := AppointmentResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("AppointmentResponse")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ArtifactAssessmentRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ArtifactAssessment")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ArtifactAssessment")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ArtifactAssessment{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ArtifactAssessment"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ArtifactAssessment{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ArtifactAssessment")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ArtifactAssessment{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ArtifactAssessment")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func AuditEventRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/AuditEvent")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("AuditEvent")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := AuditEvent{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "AuditEvent"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := AuditEvent{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("AuditEvent")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := AuditEvent{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("AuditEvent")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func BasicRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Basic")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Basic")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Basic{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Basic"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Basic{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Basic")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Basic{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Basic")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func BinaryRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Binary")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Binary")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Binary{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Binary"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Binary{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Binary")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Binary{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Binary")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func BiologicallyDerivedProductRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/BiologicallyDerivedProduct")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("BiologicallyDerivedProduct")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := BiologicallyDerivedProduct{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "BiologicallyDerivedProduct"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := BiologicallyDerivedProduct{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("BiologicallyDerivedProduct")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := BiologicallyDerivedProduct{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("BiologicallyDerivedProduct")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func BiologicallyDerivedProductDispenseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/BiologicallyDerivedProductDispense")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("BiologicallyDerivedProductDispense")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := BiologicallyDerivedProductDispense{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "BiologicallyDerivedProductDispense"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := BiologicallyDerivedProductDispense{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("BiologicallyDerivedProductDispense")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := BiologicallyDerivedProductDispense{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("BiologicallyDerivedProductDispense")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func BodyStructureRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/BodyStructure")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("BodyStructure")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := BodyStructure{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "BodyStructure"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := BodyStructure{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("BodyStructure")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := BodyStructure{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("BodyStructure")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func BundleRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Bundle")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Bundle")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Bundle{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Bundle"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Bundle{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Bundle")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Bundle{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Bundle")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CapabilityStatementRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CapabilityStatement")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CapabilityStatement")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CapabilityStatement{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CapabilityStatement"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CapabilityStatement{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CapabilityStatement")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CapabilityStatement{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CapabilityStatement")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CarePlanRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CarePlan")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CarePlan")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CarePlan{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CarePlan"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CarePlan{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CarePlan")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CarePlan{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CarePlan")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CareTeamRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CareTeam")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CareTeam")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CareTeam{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CareTeam"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CareTeam{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CareTeam")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CareTeam{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CareTeam")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ChargeItemRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ChargeItem")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ChargeItem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ChargeItem{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ChargeItem"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ChargeItem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ChargeItem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ChargeItem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ChargeItem")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ChargeItemDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ChargeItemDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ChargeItemDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ChargeItemDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ChargeItemDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ChargeItemDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ChargeItemDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ChargeItemDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ChargeItemDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CitationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Citation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Citation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Citation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Citation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Citation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Citation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Citation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Citation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ClaimRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Claim")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Claim")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Claim{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Claim"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Claim{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Claim")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Claim{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Claim")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ClaimResponseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ClaimResponse")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ClaimResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ClaimResponse{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ClaimResponse"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ClaimResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ClaimResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ClaimResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ClaimResponse")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ClinicalImpressionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ClinicalImpression")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ClinicalImpression")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ClinicalImpression{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ClinicalImpression"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ClinicalImpression{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ClinicalImpression")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ClinicalImpression{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ClinicalImpression")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ClinicalUseDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ClinicalUseDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ClinicalUseDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ClinicalUseDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ClinicalUseDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ClinicalUseDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ClinicalUseDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ClinicalUseDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ClinicalUseDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CodeSystemRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CodeSystem")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CodeSystem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CodeSystem{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CodeSystem"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CodeSystem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CodeSystem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CodeSystem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CodeSystem")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CommunicationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Communication")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Communication")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Communication{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Communication"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Communication{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Communication")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Communication{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Communication")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CommunicationRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CommunicationRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CommunicationRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CommunicationRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CommunicationRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CommunicationRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CommunicationRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CommunicationRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CommunicationRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CompartmentDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CompartmentDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CompartmentDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CompartmentDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CompartmentDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CompartmentDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CompartmentDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CompartmentDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CompartmentDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CompositionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Composition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Composition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Composition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Composition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Composition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Composition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Composition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Composition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ConceptMapRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ConceptMap")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ConceptMap")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ConceptMap{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ConceptMap"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ConceptMap{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ConceptMap")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ConceptMap{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ConceptMap")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ConditionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Condition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Condition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Condition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Condition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Condition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Condition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Condition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Condition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ConditionDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ConditionDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ConditionDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ConditionDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ConditionDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ConditionDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ConditionDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ConditionDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ConditionDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ConsentRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Consent")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Consent")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Consent{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Consent"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Consent{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Consent")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Consent{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Consent")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ContractRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Contract")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Contract")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Contract{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Contract"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Contract{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Contract")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Contract{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Contract")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CoverageRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Coverage")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Coverage")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Coverage{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Coverage"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Coverage{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Coverage")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Coverage{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Coverage")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CoverageEligibilityRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CoverageEligibilityRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CoverageEligibilityRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CoverageEligibilityRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CoverageEligibilityRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CoverageEligibilityRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CoverageEligibilityRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CoverageEligibilityRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CoverageEligibilityRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func CoverageEligibilityResponseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/CoverageEligibilityResponse")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("CoverageEligibilityResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := CoverageEligibilityResponse{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "CoverageEligibilityResponse"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := CoverageEligibilityResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("CoverageEligibilityResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := CoverageEligibilityResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("CoverageEligibilityResponse")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DetectedIssueRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DetectedIssue")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DetectedIssue")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DetectedIssue{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DetectedIssue"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DetectedIssue{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DetectedIssue")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DetectedIssue{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DetectedIssue")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Device")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Device")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Device{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Device"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Device{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Device")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Device{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Device")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceAssociationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DeviceAssociation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DeviceAssociation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DeviceAssociation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DeviceAssociation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DeviceAssociation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DeviceAssociation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DeviceAssociation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DeviceAssociation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DeviceDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DeviceDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DeviceDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DeviceDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DeviceDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DeviceDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DeviceDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DeviceDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceDispenseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DeviceDispense")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DeviceDispense")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DeviceDispense{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DeviceDispense"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DeviceDispense{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DeviceDispense")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DeviceDispense{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DeviceDispense")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceMetricRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DeviceMetric")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DeviceMetric")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DeviceMetric{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DeviceMetric"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DeviceMetric{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DeviceMetric")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DeviceMetric{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DeviceMetric")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DeviceRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DeviceRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DeviceRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DeviceRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DeviceRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DeviceRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DeviceRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DeviceRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DeviceUsageRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DeviceUsage")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DeviceUsage")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DeviceUsage{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DeviceUsage"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DeviceUsage{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DeviceUsage")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DeviceUsage{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DeviceUsage")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DiagnosticReportRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DiagnosticReport")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DiagnosticReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DiagnosticReport{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DiagnosticReport"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DiagnosticReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DiagnosticReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DiagnosticReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DiagnosticReport")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func DocumentReferenceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/DocumentReference")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("DocumentReference")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := DocumentReference{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "DocumentReference"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := DocumentReference{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("DocumentReference")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := DocumentReference{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("DocumentReference")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EncounterRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Encounter")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Encounter")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Encounter{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Encounter"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Encounter{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Encounter")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Encounter{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Encounter")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EncounterHistoryRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EncounterHistory")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EncounterHistory")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EncounterHistory{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EncounterHistory"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EncounterHistory{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EncounterHistory")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EncounterHistory{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EncounterHistory")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EndpointRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Endpoint")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Endpoint")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Endpoint{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Endpoint"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Endpoint{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Endpoint")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Endpoint{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Endpoint")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EnrollmentRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EnrollmentRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EnrollmentRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EnrollmentRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EnrollmentRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EnrollmentRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EnrollmentRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EnrollmentRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EnrollmentRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EnrollmentResponseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EnrollmentResponse")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EnrollmentResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EnrollmentResponse{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EnrollmentResponse"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EnrollmentResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EnrollmentResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EnrollmentResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EnrollmentResponse")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EpisodeOfCareRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EpisodeOfCare")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EpisodeOfCare")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EpisodeOfCare{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EpisodeOfCare"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EpisodeOfCare{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EpisodeOfCare")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EpisodeOfCare{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EpisodeOfCare")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EventDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EventDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EventDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EventDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EventDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EventDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EventDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EventDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EventDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EvidenceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Evidence")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Evidence")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Evidence{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Evidence"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Evidence{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Evidence")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Evidence{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Evidence")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EvidenceReportRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EvidenceReport")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EvidenceReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EvidenceReport{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EvidenceReport"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EvidenceReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EvidenceReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EvidenceReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EvidenceReport")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func EvidenceVariableRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/EvidenceVariable")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("EvidenceVariable")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := EvidenceVariable{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "EvidenceVariable"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := EvidenceVariable{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("EvidenceVariable")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := EvidenceVariable{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("EvidenceVariable")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ExampleScenarioRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ExampleScenario")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ExampleScenario")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ExampleScenario{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ExampleScenario"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ExampleScenario{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ExampleScenario")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ExampleScenario{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ExampleScenario")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ExplanationOfBenefitRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ExplanationOfBenefit")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ExplanationOfBenefit")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ExplanationOfBenefit{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ExplanationOfBenefit"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ExplanationOfBenefit{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ExplanationOfBenefit")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ExplanationOfBenefit{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ExplanationOfBenefit")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func FamilyMemberHistoryRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/FamilyMemberHistory")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("FamilyMemberHistory")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := FamilyMemberHistory{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "FamilyMemberHistory"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := FamilyMemberHistory{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("FamilyMemberHistory")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := FamilyMemberHistory{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("FamilyMemberHistory")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func FlagRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Flag")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Flag")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Flag{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Flag"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Flag{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Flag")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Flag{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Flag")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func FormularyItemRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/FormularyItem")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("FormularyItem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := FormularyItem{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "FormularyItem"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := FormularyItem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("FormularyItem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := FormularyItem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("FormularyItem")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func GenomicStudyRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/GenomicStudy")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("GenomicStudy")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := GenomicStudy{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "GenomicStudy"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := GenomicStudy{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("GenomicStudy")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := GenomicStudy{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("GenomicStudy")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func GoalRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Goal")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Goal")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Goal{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Goal"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Goal{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Goal")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Goal{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Goal")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func GraphDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/GraphDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("GraphDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := GraphDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "GraphDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := GraphDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("GraphDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := GraphDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("GraphDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func GroupRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Group")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Group")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Group{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Group"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Group{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Group")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Group{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Group")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func GuidanceResponseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/GuidanceResponse")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("GuidanceResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := GuidanceResponse{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "GuidanceResponse"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := GuidanceResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("GuidanceResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := GuidanceResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("GuidanceResponse")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func HealthcareServiceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/HealthcareService")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("HealthcareService")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := HealthcareService{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "HealthcareService"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := HealthcareService{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("HealthcareService")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := HealthcareService{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("HealthcareService")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ImagingSelectionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ImagingSelection")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ImagingSelection")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ImagingSelection{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ImagingSelection"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ImagingSelection{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ImagingSelection")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ImagingSelection{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ImagingSelection")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ImagingStudyRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ImagingStudy")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ImagingStudy")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ImagingStudy{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ImagingStudy"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ImagingStudy{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ImagingStudy")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ImagingStudy{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ImagingStudy")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ImmunizationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Immunization")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Immunization")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Immunization{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Immunization"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Immunization{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Immunization")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Immunization{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Immunization")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ImmunizationEvaluationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ImmunizationEvaluation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ImmunizationEvaluation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ImmunizationEvaluation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ImmunizationEvaluation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ImmunizationEvaluation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ImmunizationEvaluation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ImmunizationEvaluation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ImmunizationEvaluation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ImmunizationRecommendationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ImmunizationRecommendation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ImmunizationRecommendation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ImmunizationRecommendation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ImmunizationRecommendation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ImmunizationRecommendation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ImmunizationRecommendation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ImmunizationRecommendation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ImmunizationRecommendation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ImplementationGuideRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ImplementationGuide")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ImplementationGuide")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ImplementationGuide{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ImplementationGuide"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ImplementationGuide{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ImplementationGuide")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ImplementationGuide{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ImplementationGuide")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func IngredientRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Ingredient")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Ingredient")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Ingredient{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Ingredient"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Ingredient{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Ingredient")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Ingredient{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Ingredient")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func InsurancePlanRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/InsurancePlan")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("InsurancePlan")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := InsurancePlan{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "InsurancePlan"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := InsurancePlan{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("InsurancePlan")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := InsurancePlan{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("InsurancePlan")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func InventoryItemRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/InventoryItem")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("InventoryItem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := InventoryItem{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "InventoryItem"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := InventoryItem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("InventoryItem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := InventoryItem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("InventoryItem")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func InventoryReportRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/InventoryReport")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("InventoryReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := InventoryReport{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "InventoryReport"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := InventoryReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("InventoryReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := InventoryReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("InventoryReport")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func InvoiceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Invoice")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Invoice")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Invoice{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Invoice"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Invoice{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Invoice")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Invoice{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Invoice")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func LibraryRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Library")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Library")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Library{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Library"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Library{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Library")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Library{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Library")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func LinkageRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Linkage")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Linkage")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Linkage{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Linkage"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Linkage{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Linkage")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Linkage{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Linkage")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ListRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/List")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("List")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := List{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "List"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := List{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("List")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := List{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("List")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func LocationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Location")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Location")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Location{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Location"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Location{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Location")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Location{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Location")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ManufacturedItemDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ManufacturedItemDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ManufacturedItemDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ManufacturedItemDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ManufacturedItemDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ManufacturedItemDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ManufacturedItemDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ManufacturedItemDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ManufacturedItemDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MeasureRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Measure")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Measure")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Measure{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Measure"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Measure{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Measure")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Measure{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Measure")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MeasureReportRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MeasureReport")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MeasureReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MeasureReport{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MeasureReport"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MeasureReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MeasureReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MeasureReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MeasureReport")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Medication")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Medication")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Medication{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Medication"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Medication{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Medication")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Medication{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Medication")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicationAdministrationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MedicationAdministration")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MedicationAdministration")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MedicationAdministration{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MedicationAdministration"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MedicationAdministration{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MedicationAdministration")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MedicationAdministration{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MedicationAdministration")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicationDispenseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MedicationDispense")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MedicationDispense")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MedicationDispense{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MedicationDispense"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MedicationDispense{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MedicationDispense")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MedicationDispense{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MedicationDispense")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicationKnowledgeRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MedicationKnowledge")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MedicationKnowledge")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MedicationKnowledge{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MedicationKnowledge"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MedicationKnowledge{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MedicationKnowledge")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MedicationKnowledge{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MedicationKnowledge")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicationRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MedicationRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MedicationRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MedicationRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MedicationRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MedicationRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MedicationRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MedicationRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MedicationRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicationStatementRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MedicationStatement")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MedicationStatement")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MedicationStatement{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MedicationStatement"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MedicationStatement{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MedicationStatement")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MedicationStatement{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MedicationStatement")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MedicinalProductDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MedicinalProductDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MedicinalProductDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MedicinalProductDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MedicinalProductDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MedicinalProductDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MedicinalProductDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MedicinalProductDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MedicinalProductDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MessageDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MessageDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MessageDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MessageDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MessageDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MessageDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MessageDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MessageDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MessageDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MessageHeaderRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MessageHeader")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MessageHeader")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MessageHeader{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MessageHeader"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MessageHeader{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MessageHeader")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MessageHeader{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MessageHeader")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func MolecularSequenceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/MolecularSequence")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("MolecularSequence")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := MolecularSequence{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "MolecularSequence"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := MolecularSequence{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("MolecularSequence")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := MolecularSequence{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("MolecularSequence")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func NamingSystemRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/NamingSystem")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("NamingSystem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := NamingSystem{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "NamingSystem"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := NamingSystem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("NamingSystem")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := NamingSystem{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("NamingSystem")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func NutritionIntakeRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/NutritionIntake")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("NutritionIntake")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := NutritionIntake{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "NutritionIntake"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := NutritionIntake{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("NutritionIntake")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := NutritionIntake{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("NutritionIntake")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func NutritionOrderRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/NutritionOrder")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("NutritionOrder")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := NutritionOrder{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "NutritionOrder"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := NutritionOrder{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("NutritionOrder")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := NutritionOrder{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("NutritionOrder")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func NutritionProductRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/NutritionProduct")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("NutritionProduct")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := NutritionProduct{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "NutritionProduct"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := NutritionProduct{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("NutritionProduct")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := NutritionProduct{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("NutritionProduct")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ObservationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Observation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Observation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Observation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Observation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Observation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Observation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Observation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Observation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ObservationDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ObservationDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ObservationDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ObservationDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ObservationDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ObservationDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ObservationDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ObservationDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ObservationDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func OperationDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/OperationDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("OperationDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := OperationDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "OperationDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := OperationDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("OperationDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := OperationDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("OperationDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func OperationOutcomeRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/OperationOutcome")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("OperationOutcome")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := OperationOutcome{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "OperationOutcome"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := OperationOutcome{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("OperationOutcome")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := OperationOutcome{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("OperationOutcome")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func OrganizationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Organization")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Organization")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Organization{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Organization"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Organization{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Organization")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Organization{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Organization")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func OrganizationAffiliationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/OrganizationAffiliation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("OrganizationAffiliation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := OrganizationAffiliation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "OrganizationAffiliation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := OrganizationAffiliation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("OrganizationAffiliation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := OrganizationAffiliation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("OrganizationAffiliation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PackagedProductDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/PackagedProductDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("PackagedProductDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := PackagedProductDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "PackagedProductDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := PackagedProductDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("PackagedProductDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := PackagedProductDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("PackagedProductDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PatientRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Patient")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Patient")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Patient{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Patient"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Patient{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Patient")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Patient{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Patient")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PaymentNoticeRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/PaymentNotice")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("PaymentNotice")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := PaymentNotice{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "PaymentNotice"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := PaymentNotice{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("PaymentNotice")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := PaymentNotice{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("PaymentNotice")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PaymentReconciliationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/PaymentReconciliation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("PaymentReconciliation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := PaymentReconciliation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "PaymentReconciliation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := PaymentReconciliation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("PaymentReconciliation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := PaymentReconciliation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("PaymentReconciliation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PermissionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Permission")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Permission")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Permission{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Permission"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Permission{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Permission")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Permission{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Permission")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PersonRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Person")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Person")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Person{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Person"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Person{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Person")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Person{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Person")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PlanDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/PlanDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("PlanDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := PlanDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "PlanDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := PlanDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("PlanDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := PlanDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("PlanDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PractitionerRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Practitioner")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Practitioner")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Practitioner{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Practitioner"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Practitioner{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Practitioner")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Practitioner{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Practitioner")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func PractitionerRoleRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/PractitionerRole")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("PractitionerRole")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := PractitionerRole{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "PractitionerRole"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := PractitionerRole{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("PractitionerRole")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := PractitionerRole{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("PractitionerRole")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ProcedureRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Procedure")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Procedure")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Procedure{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Procedure"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Procedure{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Procedure")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Procedure{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Procedure")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ProvenanceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Provenance")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Provenance")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Provenance{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Provenance"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Provenance{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Provenance")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Provenance{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Provenance")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func QuestionnaireRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Questionnaire")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Questionnaire")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Questionnaire{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Questionnaire"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Questionnaire{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Questionnaire")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Questionnaire{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Questionnaire")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func QuestionnaireResponseRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/QuestionnaireResponse")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("QuestionnaireResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := QuestionnaireResponse{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "QuestionnaireResponse"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := QuestionnaireResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("QuestionnaireResponse")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := QuestionnaireResponse{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("QuestionnaireResponse")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func RegulatedAuthorizationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/RegulatedAuthorization")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("RegulatedAuthorization")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := RegulatedAuthorization{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "RegulatedAuthorization"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := RegulatedAuthorization{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("RegulatedAuthorization")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := RegulatedAuthorization{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("RegulatedAuthorization")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func RelatedPersonRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/RelatedPerson")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("RelatedPerson")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := RelatedPerson{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "RelatedPerson"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := RelatedPerson{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("RelatedPerson")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := RelatedPerson{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("RelatedPerson")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func RequestOrchestrationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/RequestOrchestration")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("RequestOrchestration")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := RequestOrchestration{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "RequestOrchestration"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := RequestOrchestration{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("RequestOrchestration")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := RequestOrchestration{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("RequestOrchestration")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func RequirementsRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Requirements")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Requirements")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Requirements{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Requirements"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Requirements{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Requirements")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Requirements{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Requirements")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ResearchStudyRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ResearchStudy")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ResearchStudy")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ResearchStudy{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ResearchStudy"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ResearchStudy{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ResearchStudy")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ResearchStudy{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ResearchStudy")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ResearchSubjectRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ResearchSubject")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ResearchSubject")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ResearchSubject{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ResearchSubject"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ResearchSubject{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ResearchSubject")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ResearchSubject{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ResearchSubject")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func RiskAssessmentRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/RiskAssessment")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("RiskAssessment")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := RiskAssessment{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "RiskAssessment"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := RiskAssessment{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("RiskAssessment")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := RiskAssessment{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("RiskAssessment")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ScheduleRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Schedule")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Schedule")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Schedule{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Schedule"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Schedule{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Schedule")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Schedule{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Schedule")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SearchParameterRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SearchParameter")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SearchParameter")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SearchParameter{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SearchParameter"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SearchParameter{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SearchParameter")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SearchParameter{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SearchParameter")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ServiceRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ServiceRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ServiceRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ServiceRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ServiceRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ServiceRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ServiceRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ServiceRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ServiceRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SlotRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Slot")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Slot")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Slot{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Slot"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Slot{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Slot")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Slot{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Slot")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SpecimenRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Specimen")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Specimen")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Specimen{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Specimen"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Specimen{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Specimen")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Specimen{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Specimen")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SpecimenDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SpecimenDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SpecimenDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SpecimenDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SpecimenDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SpecimenDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SpecimenDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SpecimenDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SpecimenDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func StructureDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/StructureDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("StructureDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := StructureDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "StructureDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := StructureDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("StructureDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := StructureDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("StructureDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func StructureMapRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/StructureMap")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("StructureMap")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := StructureMap{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "StructureMap"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := StructureMap{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("StructureMap")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := StructureMap{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("StructureMap")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubscriptionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Subscription")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Subscription")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Subscription{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Subscription"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Subscription{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Subscription")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Subscription{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Subscription")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubscriptionStatusRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubscriptionStatus")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubscriptionStatus")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubscriptionStatus{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubscriptionStatus"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubscriptionStatus{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubscriptionStatus")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubscriptionStatus{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubscriptionStatus")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubscriptionTopicRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubscriptionTopic")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubscriptionTopic")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubscriptionTopic{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubscriptionTopic"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubscriptionTopic{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubscriptionTopic")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubscriptionTopic{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubscriptionTopic")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstanceRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Substance")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Substance")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Substance{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Substance"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Substance{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Substance")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Substance{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Substance")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstanceDefinitionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubstanceDefinition")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubstanceDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubstanceDefinition{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubstanceDefinition"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubstanceDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubstanceDefinition")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubstanceDefinition{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubstanceDefinition")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstanceNucleicAcidRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubstanceNucleicAcid")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubstanceNucleicAcid")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubstanceNucleicAcid{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubstanceNucleicAcid"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubstanceNucleicAcid{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubstanceNucleicAcid")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubstanceNucleicAcid{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubstanceNucleicAcid")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstancePolymerRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubstancePolymer")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubstancePolymer")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubstancePolymer{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubstancePolymer"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubstancePolymer{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubstancePolymer")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubstancePolymer{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubstancePolymer")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstanceProteinRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubstanceProtein")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubstanceProtein")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubstanceProtein{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubstanceProtein"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubstanceProtein{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubstanceProtein")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubstanceProtein{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubstanceProtein")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstanceReferenceInformationRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubstanceReferenceInformation")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubstanceReferenceInformation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubstanceReferenceInformation{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubstanceReferenceInformation"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubstanceReferenceInformation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubstanceReferenceInformation")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubstanceReferenceInformation{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubstanceReferenceInformation")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SubstanceSourceMaterialRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SubstanceSourceMaterial")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SubstanceSourceMaterial")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SubstanceSourceMaterial{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SubstanceSourceMaterial"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SubstanceSourceMaterial{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SubstanceSourceMaterial")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SubstanceSourceMaterial{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SubstanceSourceMaterial")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SupplyDeliveryRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SupplyDelivery")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SupplyDelivery")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SupplyDelivery{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SupplyDelivery"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SupplyDelivery{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SupplyDelivery")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SupplyDelivery{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SupplyDelivery")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func SupplyRequestRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/SupplyRequest")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("SupplyRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := SupplyRequest{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "SupplyRequest"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := SupplyRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("SupplyRequest")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := SupplyRequest{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("SupplyRequest")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func TaskRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Task")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Task")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Task{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Task"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Task{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Task")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Task{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Task")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func TerminologyCapabilitiesRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/TerminologyCapabilities")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("TerminologyCapabilities")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := TerminologyCapabilities{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "TerminologyCapabilities"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := TerminologyCapabilities{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("TerminologyCapabilities")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := TerminologyCapabilities{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("TerminologyCapabilities")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func TestPlanRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/TestPlan")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("TestPlan")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := TestPlan{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "TestPlan"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := TestPlan{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("TestPlan")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := TestPlan{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("TestPlan")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func TestReportRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/TestReport")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("TestReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := TestReport{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "TestReport"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := TestReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("TestReport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := TestReport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("TestReport")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func TestScriptRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/TestScript")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("TestScript")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := TestScript{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "TestScript"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := TestScript{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("TestScript")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := TestScript{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("TestScript")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func TransportRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/Transport")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("Transport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := Transport{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "Transport"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := Transport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("Transport")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := Transport{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("Transport")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func ValueSetRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/ValueSet")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("ValueSet")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := ValueSet{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "ValueSet"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := ValueSet{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("ValueSet")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := ValueSet{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("ValueSet")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func VerificationResultRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/VerificationResult")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("VerificationResult")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := VerificationResult{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "VerificationResult"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := VerificationResult{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("VerificationResult")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := VerificationResult{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("VerificationResult")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func VisionPrescriptionRoutes(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/VisionPrescription")
	g.GET("/:id", func(c *gin.Context) {
		coll := db.FhirDatabase.Collection("VisionPrescription")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		res := coll.FindOne(context.TODO(), bson.M{"_id": objId})
		result := VisionPrescription{}
		err = res.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		format := c.DefaultQuery("_format", "json")
		println(format)
		if format == "json" {
			result.ResourceType = "VisionPrescription"
			c.JSON(http.StatusOK, result)
		} else if format == "xml" {
			c.XML(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _format parameter"})
		}
	})
	g.PUT("/:id", func(c *gin.Context) {
		a := VisionPrescription{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errA.Error()})
			fmt.Println(errA)
			return
		}
		coll := db.FhirDatabase.Collection("VisionPrescription")
		objId, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objId}, a)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource does not exist"})
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	g.POST("", func(c *gin.Context) {
		a := VisionPrescription{}
		if errA := c.ShouldBind(&a); errA != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errA.Error(),
			})
			fmt.Println(errA)
			return
		}
		objId := primitive.NewObjectID()
		a.Id = &objId
		coll := db.FhirDatabase.Collection("VisionPrescription")
		result, err := coll.InsertOne(context.TODO(), a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		c.String(http.StatusOK, "Success")
	})
	return g
}
func Routes(r *gin.Engine) {
	AccountRoutes(r)
	ActivityDefinitionRoutes(r)
	ActorDefinitionRoutes(r)
	AdministrableProductDefinitionRoutes(r)
	AdverseEventRoutes(r)
	AllergyIntoleranceRoutes(r)
	AppointmentRoutes(r)
	AppointmentResponseRoutes(r)
	ArtifactAssessmentRoutes(r)
	AuditEventRoutes(r)
	BasicRoutes(r)
	BinaryRoutes(r)
	BiologicallyDerivedProductRoutes(r)
	BiologicallyDerivedProductDispenseRoutes(r)
	BodyStructureRoutes(r)
	BundleRoutes(r)
	CapabilityStatementRoutes(r)
	CarePlanRoutes(r)
	CareTeamRoutes(r)
	ChargeItemRoutes(r)
	ChargeItemDefinitionRoutes(r)
	CitationRoutes(r)
	ClaimRoutes(r)
	ClaimResponseRoutes(r)
	ClinicalImpressionRoutes(r)
	ClinicalUseDefinitionRoutes(r)
	CodeSystemRoutes(r)
	CommunicationRoutes(r)
	CommunicationRequestRoutes(r)
	CompartmentDefinitionRoutes(r)
	CompositionRoutes(r)
	ConceptMapRoutes(r)
	ConditionRoutes(r)
	ConditionDefinitionRoutes(r)
	ConsentRoutes(r)
	ContractRoutes(r)
	CoverageRoutes(r)
	CoverageEligibilityRequestRoutes(r)
	CoverageEligibilityResponseRoutes(r)
	DetectedIssueRoutes(r)
	DeviceRoutes(r)
	DeviceAssociationRoutes(r)
	DeviceDefinitionRoutes(r)
	DeviceDispenseRoutes(r)
	DeviceMetricRoutes(r)
	DeviceRequestRoutes(r)
	DeviceUsageRoutes(r)
	DiagnosticReportRoutes(r)
	DocumentReferenceRoutes(r)
	EncounterRoutes(r)
	EncounterHistoryRoutes(r)
	EndpointRoutes(r)
	EnrollmentRequestRoutes(r)
	EnrollmentResponseRoutes(r)
	EpisodeOfCareRoutes(r)
	EventDefinitionRoutes(r)
	EvidenceRoutes(r)
	EvidenceReportRoutes(r)
	EvidenceVariableRoutes(r)
	ExampleScenarioRoutes(r)
	ExplanationOfBenefitRoutes(r)
	FamilyMemberHistoryRoutes(r)
	FlagRoutes(r)
	FormularyItemRoutes(r)
	GenomicStudyRoutes(r)
	GoalRoutes(r)
	GraphDefinitionRoutes(r)
	GroupRoutes(r)
	GuidanceResponseRoutes(r)
	HealthcareServiceRoutes(r)
	ImagingSelectionRoutes(r)
	ImagingStudyRoutes(r)
	ImmunizationRoutes(r)
	ImmunizationEvaluationRoutes(r)
	ImmunizationRecommendationRoutes(r)
	ImplementationGuideRoutes(r)
	IngredientRoutes(r)
	InsurancePlanRoutes(r)
	InventoryItemRoutes(r)
	InventoryReportRoutes(r)
	InvoiceRoutes(r)
	LibraryRoutes(r)
	LinkageRoutes(r)
	ListRoutes(r)
	LocationRoutes(r)
	ManufacturedItemDefinitionRoutes(r)
	MeasureRoutes(r)
	MeasureReportRoutes(r)
	MedicationRoutes(r)
	MedicationAdministrationRoutes(r)
	MedicationDispenseRoutes(r)
	MedicationKnowledgeRoutes(r)
	MedicationRequestRoutes(r)
	MedicationStatementRoutes(r)
	MedicinalProductDefinitionRoutes(r)
	MessageDefinitionRoutes(r)
	MessageHeaderRoutes(r)
	MolecularSequenceRoutes(r)
	NamingSystemRoutes(r)
	NutritionIntakeRoutes(r)
	NutritionOrderRoutes(r)
	NutritionProductRoutes(r)
	ObservationRoutes(r)
	ObservationDefinitionRoutes(r)
	OperationDefinitionRoutes(r)
	OperationOutcomeRoutes(r)
	OrganizationRoutes(r)
	OrganizationAffiliationRoutes(r)
	PackagedProductDefinitionRoutes(r)
	PatientRoutes(r)
	PaymentNoticeRoutes(r)
	PaymentReconciliationRoutes(r)
	PermissionRoutes(r)
	PersonRoutes(r)
	PlanDefinitionRoutes(r)
	PractitionerRoutes(r)
	PractitionerRoleRoutes(r)
	ProcedureRoutes(r)
	ProvenanceRoutes(r)
	QuestionnaireRoutes(r)
	QuestionnaireResponseRoutes(r)
	RegulatedAuthorizationRoutes(r)
	RelatedPersonRoutes(r)
	RequestOrchestrationRoutes(r)
	RequirementsRoutes(r)
	ResearchStudyRoutes(r)
	ResearchSubjectRoutes(r)
	RiskAssessmentRoutes(r)
	ScheduleRoutes(r)
	SearchParameterRoutes(r)
	ServiceRequestRoutes(r)
	SlotRoutes(r)
	SpecimenRoutes(r)
	SpecimenDefinitionRoutes(r)
	StructureDefinitionRoutes(r)
	StructureMapRoutes(r)
	SubscriptionRoutes(r)
	SubscriptionStatusRoutes(r)
	SubscriptionTopicRoutes(r)
	SubstanceRoutes(r)
	SubstanceDefinitionRoutes(r)
	SubstanceNucleicAcidRoutes(r)
	SubstancePolymerRoutes(r)
	SubstanceProteinRoutes(r)
	SubstanceReferenceInformationRoutes(r)
	SubstanceSourceMaterialRoutes(r)
	SupplyDeliveryRoutes(r)
	SupplyRequestRoutes(r)
	TaskRoutes(r)
	TerminologyCapabilitiesRoutes(r)
	TestPlanRoutes(r)
	TestReportRoutes(r)
	TestScriptRoutes(r)
	TransportRoutes(r)
	ValueSetRoutes(r)
	VerificationResultRoutes(r)
	VisionPrescriptionRoutes(r)
}
