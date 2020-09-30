# Slowflow

根据github.com/antlinker/flow 修改

## 工作流设计器

- [Camunda下载地址](https://camunda.com/download/modeler/)
- [文档参考](https://docs.awspaas.com/reference-guide/aws-paas-process-reference-guide/process_structure/activities.html)

## 获取项目

```bash
go get github.com/redshiftkeying/slowflow-server
```

## 使用

### 1. 初始化工作流引擎

#### 默认配置

```go
package main
import (
    sw "github.com/redshiftkeying/slowflow-server/"
    _ "github.com/go-sql-driver/mysql"
)
func main() {
   	//use default configs
	sw.SetServerOptions("")
	// Instance
	s := sw.Init()
	// add Middleware
	s.AddServerMiddleware(filter)
	// start server
    s.StartServer()
}
```

#### 使用配置文件

```go
package main
import (
    sw "github.com/redshiftkeying/slowflow-server/"
    _ "github.com/go-sql-driver/mysql"
)
func main() {
   	//use default configs
	sw.SetServerOptions("./config.yml")
	// Instance
	s := sw.Init()
	// add Middleware
	s.AddServerMiddleware(filter)
	// start server
    s.StartServer()
}
```

配置文件样例：

```yaml
# config.yml
dsn: root:123456@tcp(127.0.0.1:3306)/sflow?charset=utf8
mode: debug
staticroot: ./build
staticPath: /
port: 6062
```

### 2. 加载工作流文件

```go
	err := sw.LoadFile("leave.bpmn")
	if err != nil {
		// 处理错误
	}
```

### 3. 发起流程

```go
  input := map[string]interface{}{
	"day": 1,
  }

	result, err := sw.StartFlow("流程编号", "开始节点编号", "流程发起人ID", input)
	if err != nil {
		// 处理错误
	}
```

### 4. 查询待办流程列表

```go
	todos, err := sw.QueryTodoFlows("流程编号", "流程待办人ID")
	if err != nil {
		// 处理错误
	}
```

### 5. 处理流程

```go
  input := map[string]interface{}{
	"action": "pass",
  }

  result, err = sw.HandleFlow("待办流程节点实例ID", "流程处理人ID", input)
	if err != nil {
		// 处理错误
	}
```

### 6. 停止流程

```go
	err := sw.StopFlow("待办流程节点实例ID", func(flowInstance *schema.FlowInstance) bool {
		return flowInstance.Launcher == "XXX"
	})
	if err != nil {
		// 处理错误
	}
```

### 7. 接入WEB流程管理

```go
func main() {
serverOptions := []flow.ServerOption{
	    sw.ServerStaticRootOption("./build"),
	    sw.ServerPrefixOption("/flow/"),
	    sw.ServerMiddlewareOption(filter),
	}

	http.Handle("/flow/", sw.StartServer(engine.GetInstance(), serverOptions...))
}

func filter(c *gin.Context) {
    fmt.Printf("请求参数：%s \n", c.FullPath())
}
```

### 8. 查询流程待办数据

```go
	result,err := sw.QueryTodoFlows("流程编号","流程处理人ID")
	if err != nil {
		// 处理错误
	}
```

### 9. 查询流程历史数据

```go
result,err := sw.QueryFlowHistory("待办流程实例ID")
if err != nil {
	// 处理错误
}
```

### 10. 查询已办理的流程实例ID列表

```go
ids,err := sw.QueryDoneFlowIDs("流程编号","流程处理人ID")
if err != nil {
	// 处理错误
}
```

### 11. 查询节点实例的候选人ID列表

```go
ids,err := sw.QueryNodeCandidates("待办流程节点实例ID")
if err != nil {
	// 处理错误
}
```

### 12. 停止流程实例

```go
	err := sw.StopFlowInstance("待办流程节点实例ID", func(flowInstance *schema.FlowInstance) bool {
		return flowInstance.Launcher == "XXX"
	})
	if err != nil {
		// 处理错误
	}
```


## TODO列表
- [x] 移除ionic
- [x] 使用GIN代替gear框架
- [ ] 接入ORM支持多种数据库
- [ ] 支持子流程
- [ ] 重构并完善客户端
- [ ] 新增APP客户端
- [ ] 增加丰富的例子
- [ ] 完善BPMN2.0的支持和实现

