package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetAlertQuery ...
func GetAlertQuery(params graphql.ResolveParams) (interface{}, error) {
	var alerts []models.Alert

	// ... Implémenter la logique de base de données ici

	return alerts, nil
}
