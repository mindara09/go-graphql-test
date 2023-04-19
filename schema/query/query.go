package query

import (
	"fmt"
	"go-graphql-test/config"
	"go-graphql-test/schema/types"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ProductList": &graphql.Field{
			/*
				Get List Products
				http://localhost:8080/graphql?query={ProductList{id,name_product,type_product,created_at}}
			*/
			Type:        graphql.NewList(types.ProductsTypes),
			Description: "Get List Product",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var a types.Products
				var b []types.Products

				db, err := config.GetConnection()
				if err != nil {
					panic(err.Error())
				}

				result, err := db.Query("SELECT id, name_product, type_product, created_at FROM products")
				if err != nil {
					panic(err.Error())
				}

				for result.Next() {
					err = result.Scan(&a.Id, &a.NameProduct, &a.TypeProduct, &a.CreatedAt)
					if err != nil {
						panic(err.Error())
					}
					b = append(b, a)
				}

				return b, nil

			},
		},
		"ProductDetail": &graphql.Field{
			/*
				Get List Products by id
				http://localhost:8080/graphql?query={ProductDetail(id:1){id,name_product,type_product,created_at}}
			*/
			Type:        graphql.NewList(types.ProductsTypes),
			Description: "Get List Products by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					var a types.Products
					var b []types.Products

					db, err := config.GetConnection()
					if err != nil {
						panic(err.Error())
					}

					query := fmt.Sprintf("SELECT id, name_product, type_product, created_at FROM products WHERE id='%v'", id)
					result, err := db.Query(query)
					if err != nil {
						panic(err.Error())
					}

					for result.Next() {
						err = result.Scan(&a.Id, &a.NameProduct, &a.TypeProduct, &a.CreatedAt)
						if err != nil {
							panic(err.Error())
						}
						b = append(b, a)
					}
					return b, nil
				}
				return nil, nil
			},
		},
	},
})
