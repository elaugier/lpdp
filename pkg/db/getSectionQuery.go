package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetSectionQuery ...
func GetSectionQuery(params graphql.ResolveParams) (interface{}, error) {
	var sections []models.Section

	// ... Implémenter la logique de base de données ici

	return sections, nil
}
