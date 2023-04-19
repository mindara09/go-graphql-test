package mutation

import (
	"go-graphql-test/config"
	"go-graphql-test/schema/types"

	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"LoginAccount": &graphql.Field{
			/*
				Login Account
				http://localhost:8080/graphql?query=mutation+_{LoginAccount(username:"",password:""){id,username,password,created_at}}
			*/
			Type:        graphql.NewList(types.Login),
			Description: "Login Account",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		"CreateProduct": &graphql.Field{
			/*
				Create Product
				http://localhost:8080/graphql?query=mutation+_{CreateProduct(name_product:"",type_product:2){name_product,type_product}}
			*/
			Type:        graphql.NewList(types.ProductsTypes),
			Description: "Create Data Product",
			Args: graphql.FieldConfigArgument{
				"name_product": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type_product": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				var a []types.Products

				db, err := config.GetConnection()
				if err != nil {
					panic(err.Error())
				}

				name_product := params.Args["name_product"].(string)
				type_product := params.Args["type_product"].(int)

				_, err = db.Query("INSERT INTO products(name_product, type_product) VALUES (?,?)", name_product, type_product)

				if err != nil {
					panic(err.Error())
				}

				product := types.Products{
					NameProduct: name_product,
					TypeProduct: type_product,
				}

				a = append(a, product)

				return a, err

			},
		},
		"UpdateProduct": &graphql.Field{
			/*
				Update Product by id
				http://localhost:8080/graphql?query=mutation+_{UpdateProduct(id:3,name_product:"Kursi Rexus",type_product:2){id,name_product,type_product}}
			*/
			Type:        graphql.NewList(types.ProductsTypes),
			Description: "Update Data Product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"name_product": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type_product": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				var a []types.Products

				db, err := config.GetConnection()
				if err != nil {
					panic(err.Error())
				}

				id := params.Args["id"].(int)
				name_product := params.Args["name_product"].(string)
				type_product := params.Args["type_product"].(int)

				_, err = db.Query("UPDATE products SET name_product=?, type_product=? WHERE id=?", name_product, type_product, id)
				if err != nil {
					panic(err.Error())
				}

				product := types.Products{
					Id:          id,
					NameProduct: name_product,
					TypeProduct: type_product,
				}

				a = append(a, product)

				return a, err
			},
		},
		"DeleteProduct": &graphql.Field{
			/*
				Delete Product by id
				http://localhost:8080/graphql?query=mutation+_{DeleteProduct(id:1){id}}
			*/
			Type:        graphql.NewList(types.ProductsTypes),
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var a []types.Products

				db, err := config.GetConnection()
				if err != nil {
					panic(err.Error())
				}

				id := params.Args["id"].(int)

				_, err = db.Query("DELETE FROM products WHERE id=?", id)
				if err != nil {
					panic(err.Error())
				}

				product := types.Products{
					Id: id,
				}

				a = append(a, product)

				return a, err

			},
		},
	},
})
