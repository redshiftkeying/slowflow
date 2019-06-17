package router

import (
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/redshiftkeying/slowflow/schema"
)

// GraphqlHandler post router
func GraphqlHandler() gin.HandlerFunc {
	helloSchema := graphql.MustParseSchema(schema.GetRootSchema(), &schema.Resolver{})

	h := gin.WrapH(&relay.Handler{Schema: helloSchema})
	return h
}
