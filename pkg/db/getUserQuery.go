package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetUserQuery ...
func GetUserQuery(params graphql.ResolveParams) (interface{}, error) {
	var users []models.User

	// ... Implémenter la logique de base de données ici

	return users, nil
}
