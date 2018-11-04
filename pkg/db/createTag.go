package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//CreateTag ...
func CreateTag(params graphql.ResolveParams) (interface{}, error) {
	tag := &models.Tag{
		Code:    params.Args["code"].(string),
		Label:   params.Args["label"].(string),
		Enabled: params.Args["enabled"].(bool),
	}

	// ... Implémenter la logique de base de données ici

	return tag, nil
}
