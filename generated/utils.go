package generated

// func GenerateMongoVersion(resourceType string, id primitive.ObjectID) {
// 	coll := db.FhirDatabase.Collection(resourceType)
// 	revColl := db.FhirDatabase.Collection(resourceType + "_rev")

// 	// What we want to do:
// 	// 1. Find the current version of the resource (in the resource collection)
// 	// 2. Insert a new version of the resource (in the resource_rev collection) (automatically increment the version and add a backreference to the new document)
// 	// 3. Update the current version of the resource (in the resource collection) (automatically increment the version and add a backreference to the new document)

// 	// Some important notes:
// 	// - The version field is a nested field in the "meta" structure, so "meta.version" == <documentversion>

// 	c

// 	mongo.Pipeline{}
// }
