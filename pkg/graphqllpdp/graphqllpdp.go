package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

var (
	graphqlLpdpSchema graphql.Schema
)

// GetSchema ...
func GetSchema() (graphql.Schema, error) {

	achievementInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Achievement",
		Description: "Achievement for LPDP Member",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "achievement id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of achievement",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of achievement",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of achievement",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of achievement",
			},
			"OwnerRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of achievement",
			},
		},
	})

	activityInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Activity",
		Description: "Activity for LPDP Member",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "activity id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of activity",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of activity",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of activity",
			},
			"Message": &graphql.Field{
				Type:        graphql.String,
				Description: "message of activity",
			},
			"Type": &graphql.Field{
				Type:        graphql.String,
				Description: "type of activity",
			},
			"OwnerRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of activity",
			},
		},
	})

	alertInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Alert",
		Description: "Alert for LPDP Post or Comment",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "alert id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of alert",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of alert",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of alert",
			},
			"Type": &graphql.Field{
				Type:        graphql.String,
				Description: "type of alert",
			},
			"Details": &graphql.Field{
				Type:        graphql.String,
				Description: "Details of alert",
			},
			"PostRef": &graphql.Field{
				Type:        graphql.String,
				Description: "post targeted by the alert",
			},
			"CommentRef": &graphql.Field{
				Type:        graphql.String,
				Description: "comment targeted of alert",
			},
			"UserRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of alert",
			},
		},
	})

	alertActionInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "AlertAction",
		Description: "Alert action for LPDP Alert",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "alert action id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of alert action",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of alert action",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of alert action",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of alert action",
			},
			"Details": &graphql.Field{
				Type:        graphql.String,
				Description: "Details of alert action",
			},
			"AlertRef": &graphql.Field{
				Type:        graphql.String,
				Description: "alert targeted by the alert action",
			},
			"UserRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of alert action",
			},
		},
	})

	badgeInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Badge",
		Description: "Badge for LPDP Member",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "badge id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of badge",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of badge",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of badge",
			},
			"Message": &graphql.Field{
				Type:        graphql.String,
				Description: "message of badge",
			},
			"Type": &graphql.Field{
				Type:        graphql.String,
				Description: "type of badge",
			},
			"OwnerRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of badge",
			},
		},
	})

	userInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "User",
		Description: "Member of LPDP",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "user id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of user",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of user",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of user",
			},
			"Email": &graphql.Field{
				Type:        graphql.String,
				Description: "email address of user",
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"achievements": &graphql.Field{
				Type: graphql.NewList(achievementInterface),
			},
			"achievement": &graphql.Field{
				Type: achievementInterface,
			},
			"activity": &graphql.Field{
				Type: activityInterface,
			},
			"alerts": &graphql.Field{
				Type: graphql.NewList(alertInterface),
			},
			"alert": &graphql.Field{
				Type: alertInterface,
			},
			"alertactions": &graphql.Field{
				Type: graphql.NewList(alertActionInterface),
			},
			"alertaction": &graphql.Field{
				Type: alertActionInterface,
			},
			"badges": &graphql.Field{
				Type: graphql.NewList(badgeInterface),
			},
			"badge": &graphql.Field{
				Type: badgeInterface,
			},
			"users": &graphql.Field{
				Type: graphql.NewList(userInterface),
			},
			"user": &graphql.Field{
				Type: userInterface,
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: graphql.Fields{},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
