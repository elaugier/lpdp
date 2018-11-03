package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//CreateAchievement ...
func CreateAchievement(params graphql.ResolveParams) (interface{}, error) {
	achievement := &models.Achievement{
		Title:    params.Args["title"].(string),
		OwnerRef: params.Args["owner_ref"].(uuid.UUID),
	}

	// ... Implémenter la logique de base de données ici

	return achievement, nil
}
