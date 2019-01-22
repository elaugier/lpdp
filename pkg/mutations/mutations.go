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
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Alert type.",
				},
				"details": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Alert description.",
				},
				"user_ref": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "User affected by the alert.",
				},
				"post_ref": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Post affected by the the alert.",
				},
				"comment_ref": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Comment affected by the alert.",
				},
			},
			Description: "Add a new alert",
		},
		"addAlertAction": &graphql.Field{
			Type: models.AlertActionT,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Alert action title",
				},
				"details": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Alert action details",
				},
				"alert_ref": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Alert id reference",
				},
				"user_ref": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "User id reference",
				},
			},
			Description: "Add alert action",
		},
		"addSection": &graphql.Field{
			Type: models.SectionT,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Section name.",
				},
				"shortname": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Section shortname.",
				},
				"description": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Describe section content",
				},
			},
			Description: "Add new section.",
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
