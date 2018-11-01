package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetAchievementQuery ...
func GetAchievementQuery(params graphql.ResolveParams) (interface{}, error) {
	var achievements []models.Achievement

	// ... Implémenter la logique de base de données ici

	return achievements, nil
}
