package versioning

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/gofhir/db"
	"github.com/gofhir/generated"
	"github.com/r3labs/diff/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DiffObject(a, b interface{}) (diff.Changelog, error) {
	// 1. Compare the resourceType fields

	return diff.Diff(a, b)
}

func SetupChangeStream(waitGroup *sync.WaitGroup, database *mongo.Database) (context.CancelFunc, error) {
	var required = options.Required
	stream, err := database.Watch(context.TODO(), mongo.Pipeline{}, &options.ChangeStreamOptions{
		FullDocumentBeforeChange: &required,
	})

	if err != nil {
		return nil, err
	}

	waitGroup.Add(1)

	routineCtx, cancelFn := context.WithCancel(context.Background())
	go iterateChangeStream(routineCtx, waitGroup, stream)

	return cancelFn, nil
}

type ChangeStreamEvent struct {
	OperationType            string `bson:"operationType"`
	FullDocument             bson.M `bson:"fullDocument"`
	FullDocumentBeforeChange bson.M `bson:"fullDocumentBeforeChange"`
	Namespace                struct {
		Database   string `bson:"db"`
		Collection string `bson:"coll"`
	} `bson:"ns"`
	bson.M `bson:",inline"`
}

func iterateChangeStream(routineCtx context.Context, waitGroup *sync.WaitGroup, stream *mongo.ChangeStream) {
	defer stream.Close(routineCtx)
	defer waitGroup.Done()

	for stream.Next(routineCtx) {
		var data ChangeStreamEvent
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}

		// TODO we should do some (optional) sanity checks here. I.e. check that we're in the fhir database, that the collection == resourceType, etc.

		if data.OperationType == "replace" {
			fullDocument := reflect.New(generated.TypeRegistry[data.Namespace.Collection]).Interface()

			fullDocumentMarshaled, err := bson.Marshal(data.FullDocument)
			if err != nil {
				fmt.Printf("Error marshaling %+v\n", err)
				continue
			}

			if err := bson.Unmarshal(fullDocumentMarshaled, fullDocument); err != nil {
				fmt.Printf("Error unmarshalling %+v\n", err)
				continue
			}

			fullDocumentBeforeChange := reflect.New(generated.TypeRegistry[data.Namespace.Collection]).Interface()

			fullDocumentBeforeChangeMarshaled, err := bson.Marshal(data.FullDocumentBeforeChange)
			if err != nil {
				fmt.Printf("Error marshaling %+v\n", err)
				continue
			}

			if err := bson.Unmarshal(fullDocumentBeforeChangeMarshaled, fullDocumentBeforeChange); err != nil {
				fmt.Printf("Error unmarshalling %+v\n", err)
				continue
			}

			id := reflect.ValueOf(fullDocument).Elem().FieldByName("Id").Interface().(*primitive.ObjectID)
			versionId := reflect.ValueOf(fullDocument).Elem().FieldByName("Meta").Elem().FieldByName("VersionId").Interface().(int)

			difference, err := diff.Diff(fullDocument, fullDocumentBeforeChange)
			if err != nil {
				fmt.Printf("Error diffing %+v\n", err)
				continue
			}

			fmt.Printf("Difference: %+v\n", difference)

			_, err = db.FhirDatabase.Collection(data.Namespace.Collection).UpdateByID(
				routineCtx,
				id,
				bson.M{
					"$inc": bson.M{
						// TODO I should fix the capitalization of bson fields...
						"meta.versionid": 1,
					},
					"$currentDate": bson.M{
						"meta.lastupdated": true,
					},
				},
				&options.UpdateOptions{
					BypassDocumentValidation: &[]bool{true}[0],
				},
			)

			if err != nil {
				fmt.Printf("Error updating resource %+v\n", err)
				continue
			}

			_, err = db.VersioningDatabase.Collection(data.Namespace.Collection).InsertOne(
				routineCtx,
				bson.M{
					"resourceId": id,
					"versionId":  versionId,
					"delta":      difference,
				},
			)

			if err != nil {
				fmt.Printf("Error inserting version %+v\n", err)
				continue
			}
		}
	}

	if err := stream.Err(); err != nil {
		if err.Error() == "context canceled" {
			return
		}
		panic(err)
	}
}
