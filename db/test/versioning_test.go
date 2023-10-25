package test

// func TestDiffBasicObservation(t *testing.T) {
// 	observation1 := generated.Observation{
// 		Language: "en",
// 		DataAbsentReason: &generated.CodeableConcept{
// 			Coding: []*generated.Coding{
// 				{
// 					System: "http://terminology.hl7.org/CodeSystem/data-absent-reason",
// 					Code:   "unknown",
// 				},
// 			},
// 		},
// 	}
// 	observation2 := generated.Observation{
// 		Language: "de",
// 	}

// 	diff, err := db.DiffObject(observation1, observation2)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	marshaled, err := json.MarshalIndent(diff, "", "  ")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(string(marshaled))
// }
