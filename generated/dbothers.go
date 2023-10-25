package generated

import (
	"context"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

func othersCreateCollections(db *mongo.Database) error {
	var err error

	err = db.CreateCollection(context.TODO(), "CDSHooksRequestOrchestration", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EBMRecommendation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableMeasure", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableMeasure", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ActualGroup", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "GroupDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "FamilyMemberHistoryForGeneticsAnalysis", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableActivityDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableActivityDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ProvenanceRelevantHistory", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableConceptMap", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableConceptMap", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableValueSet", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableValueSet", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ComputableValueSet", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExecutableValueSet", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableCodeSystem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableCodeSystem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CDSHooksGuidanceResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceMetricObservationProfile", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationvitalsigns", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationvitalspanel", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationresprate", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationoxygensat", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationheartrate", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationheadcircum", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationbp", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationbodyweight", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationbodytemp", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationbodyheight", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observationbmi", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "LogicLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CQLLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ELMLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "FHIRPathLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ModelInfoLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ModuleDefinitionLibrary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableTestScript", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExampleLipidProfile", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExampleLipidProfile", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExampleLipidProfile", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExampleLipidProfile", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExampleLipidProfile", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ClinicalDocument", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ProfileForCatalog", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "HistoryBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SearchSetBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubscriptionNotificationBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "BatchBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "BatchResponseBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "TransactionBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "TransactionResponseBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DocumentBundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareableNamingSystem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishableNamingSystem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ShareablePlanDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CDSHooksServicePlanDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PublishablePlanDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ComputablePlanDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	return nil
}
