package mutations

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available mutations.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"addAchievement": &graphql.Field{
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
		"addActivity": &graphql.Field{
			Type: models.AchievementT,
			Args: graphql.FieldConfigArgument{
				"message": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		"addAlert": &graphql.Field{
			Type: models.AlertT,
			Args: graphql.FieldConfigArgument{
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"details": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"post_ref": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"user_ref": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"comment_ref": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
		},
		"addSection": &graphql.Field{
			Type: models.SectionT,
		},
		"removeAchievement": &graphql.Field{
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
