package projections

import "go.mongodb.org/mongo-driver/bson"

func CreateUserProjection(fields string) bson.M {
	defaultFields := []string{"_id"}

	return CreateProjection(fields, defaultFields)
}