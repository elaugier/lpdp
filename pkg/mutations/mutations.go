package mutations

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available mutations.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createAchievement": &graphql.Field{
			Type: models.AchievementT,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"owner_ref": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Description: "Create a new achievement.",
		},
		"createSection": &graphql.Field{
			Type: models.SectionT,
		},
		"updateAchievement": &graphql.Field{
			Type: models.AchievementT,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "id of the achievement",
				},
			},
			Description: "Delete an achievement.",
		},
	}
}
