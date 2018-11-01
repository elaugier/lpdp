package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetTagQuery ...
func GetTagQuery(params graphql.ResolveParams) (interface{}, error) {
	var tags []models.Tag

	// ... Implémenter la logique de base de données ici

	return tags, nil
}
