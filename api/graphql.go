package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/redshiftkeying/slowflow-server/service/graph"
	"github.com/redshiftkeying/slowflow-server/service/graph/generated"
)

// GraphqlConfig 配置
type GraphqlConfig struct {
	Method  Method
	Path    string
	Handler gin.HandlerFunc
}

// SetHandler 与Gin框架的接口
func (g *GraphqlConfig) SetHandler(h *gin.HandlerFunc) (err error) {
	g.Handler = *h
	return
}

// GraphqlHandler POST graphql query endpoint
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
