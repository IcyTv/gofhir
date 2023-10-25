package generated

import (
	"context"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

func resourcesCreateCollections(db *mongo.Database) error {
	var err error

	err = db.CreateCollection(context.TODO(), "Account", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ActivityDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ActorDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "AdministrableProductDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "AdverseEvent", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "AllergyIntolerance", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Appointment", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "AppointmentResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ArtifactAssessment", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "AuditEvent", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Basic", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Binary", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "BiologicallyDerivedProduct", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "BiologicallyDerivedProductDispense", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "BodyStructure", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Bundle", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CapabilityStatement", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CarePlan", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CareTeam", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ChargeItem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ChargeItemDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Citation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Claim", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ClaimResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ClinicalImpression", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ClinicalUseDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CodeSystem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Communication", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CommunicationRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CompartmentDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Composition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ConceptMap", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Condition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ConditionDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Consent", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Contract", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Coverage", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CoverageEligibilityRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "CoverageEligibilityResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DetectedIssue", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Device", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceAssociation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceDispense", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceMetric", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DeviceUsage", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DiagnosticReport", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "DocumentReference", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Encounter", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EncounterHistory", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Endpoint", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EnrollmentRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EnrollmentResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EpisodeOfCare", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EventDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Evidence", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EvidenceReport", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "EvidenceVariable", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExampleScenario", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ExplanationOfBenefit", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "FamilyMemberHistory", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Flag", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "FormularyItem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "GenomicStudy", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Goal", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "GraphDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Group", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "GuidanceResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "HealthcareService", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ImagingSelection", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ImagingStudy", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Immunization", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ImmunizationEvaluation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ImmunizationRecommendation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ImplementationGuide", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Ingredient", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "InsurancePlan", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "InventoryItem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "InventoryReport", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Invoice", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Library", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Linkage", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "List", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Location", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ManufacturedItemDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Measure", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MeasureReport", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Medication", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MedicationAdministration", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MedicationDispense", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MedicationKnowledge", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MedicationRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MedicationStatement", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MedicinalProductDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MessageDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MessageHeader", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "MolecularSequence", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "NamingSystem", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "NutritionIntake", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "NutritionOrder", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "NutritionProduct", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Observation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ObservationDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "OperationDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "OperationOutcome", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Organization", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "OrganizationAffiliation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PackagedProductDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Parameters", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Patient", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PaymentNotice", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PaymentReconciliation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Permission", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Person", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PlanDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Practitioner", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "PractitionerRole", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Procedure", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Provenance", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Questionnaire", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "QuestionnaireResponse", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "RegulatedAuthorization", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "RelatedPerson", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "RequestOrchestration", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Requirements", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ResearchStudy", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ResearchSubject", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "RiskAssessment", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Schedule", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SearchParameter", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ServiceRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Slot", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Specimen", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SpecimenDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "StructureDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "StructureMap", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Subscription", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubscriptionStatus", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubscriptionTopic", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Substance", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubstanceDefinition", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubstanceNucleicAcid", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubstancePolymer", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubstanceProtein", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubstanceReferenceInformation", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SubstanceSourceMaterial", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SupplyDelivery", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "SupplyRequest", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Task", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "TerminologyCapabilities", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "TestPlan", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "TestReport", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "TestScript", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "Transport", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "ValueSet", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "VerificationResult", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	err = db.CreateCollection(context.TODO(), "VisionPrescription", &options.CreateCollectionOptions{ChangeStreamPreAndPostImages: bson.M{"enabled": true}})
	if err != nil {
		return err
	}

	return nil
}
