package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetBadgeQuery ...
func GetBadgeQuery(params graphql.ResolveParams) (interface{}, error) {
	var badges []models.Badge

	// ... Implémenter la logique de base de données ici

	return badges, nil
}
