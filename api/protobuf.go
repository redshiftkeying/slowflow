package api

import "github.com/gin-gonic/gin"

// ProtobufConfig 提供API管理
type ProtobufConfig struct {
	Method  Method
	Path    string
	Handler gin.HandlerFunc
}
