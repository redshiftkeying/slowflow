package api

import "github.com/gin-gonic/gin"

// GraphqlConfig 配置
type GraphqlConfig struct {
	Method  Method
	Path    string
	Handler gin.HandlerFunc
}
