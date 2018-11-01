package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetActivityQuery ...
func GetActivityQuery(params graphql.ResolveParams) (interface{}, error) {
	var activities []models.Activity

	// ... Implémenter la logique de base de données ici

	return activities, nil
}
