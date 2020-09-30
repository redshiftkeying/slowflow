package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redshiftkeying/slowflow-server/engine"
	"github.com/redshiftkeying/slowflow-server/schema"
)

// Method API支持的方法
type Method string

const (
	MethodGET     Method = "GET"
	MethodPUT     Method = "PUT"
	MethodPOST    Method = "POST"
	MethodDELETE  Method = "DELETE"
	MethodPATCH   Method = "PATCH"
	MethodOPTIONS Method = "OPTIONS"
	MethodHEAD    Method = "HEAD"
)

// RestConfig 配置
type RestConfig struct {
	Method  Method
	Path    string
	Handler gin.HandlerFunc
}

// 	router.GET("/api/v1/flow", restAPI.QueryFlowPage)
// 	router.GET("/api/v1/flow/:id", restAPI.GetFlow)
// 	router.DELETE("/api/v1/flow/:id", restAPI.DeleteFlow)
// 	router.POST("/api/v1/flow", restAPI.SaveFlow)

// RestAPI 提供API管理
type RestAPI struct {
	engine *engine.Engine
}

// 获取分页的页索引
func (a *RestAPI) pageIndex(ctx *gin.Context) uint {
	if v := ctx.Query("current"); v != "" {
		i, _ := strconv.Atoi(v)
		return uint(i)
	}
	return 1
}

// 获取分页的页大小
func (a *RestAPI) pageSize(ctx *gin.Context) uint {
	if v := ctx.Query("pageSize"); v != "" {
		i, _ := strconv.Atoi(v)
		if i > 40 {
			i = 40
		}
		return uint(i)
	}
	return 10
}

// QueryFlowPage 查询流程分页数据
func (a *RestAPI) QueryFlowPage(ctx *gin.Context) {
	pageIndex, pageSize := a.pageIndex(ctx), a.pageSize(ctx)
	params := schema.FlowQueryParam{
		Code: ctx.Query("code"),
		Name: ctx.Query("name"),
	}

	total, items, err := a.engine.FlowBll().QueryAllFlowPage(params, pageIndex, pageSize)
	if err != nil {
		return
	}

	response := map[string]interface{}{
		"list": items,
		"pagination": map[string]interface{}{
			"total":    total,
			"current":  pageIndex,
			"pageSize": pageSize,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

// GetFlow 获取流程数据
func (a *RestAPI) GetFlow(ctx *gin.Context) {
	item, err := a.engine.FlowBll().GetFlow(ctx.Param("id"))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, item)
}

type saveFlowRequest struct {
	XML string `json:"xml"`
}

func (a *saveFlowRequest) Validate() error {
	if len(a.XML) == 0 {
		return errors.New("请求含有空数据")
	}
	return nil
}

// SaveFlow 保存流程
func (a *RestAPI) SaveFlow(ctx *gin.Context) {
	var req saveFlowRequest
	if err := ctx.ShouldBindXML(&req); err != nil {
		return
	}

	_, err := a.engine.CreateFlow([]byte(req.XML))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

// DeleteFlow 删除流程数据
func (a *RestAPI) DeleteFlow(ctx *gin.Context) {
	err := a.engine.FlowBll().DeleteFlow(ctx.Param("id"))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
