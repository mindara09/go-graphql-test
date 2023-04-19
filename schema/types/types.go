package types

import "github.com/graphql-go/graphql"

var ProductsTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Products",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name_product": &graphql.Field{
			Type: graphql.String,
		},
		"type_product": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var TypesProducts = graphql.NewObject(graphql.ObjectConfig{
	Name: "Type Products",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name_type": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var Login = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login Account",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
