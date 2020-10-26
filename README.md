# Slowflow

根据 github.com/antlinker/flow 修改

## 工作流设计器

- [Camunda 下载地址](https://camunda.com/download/modeler/)
- [文档参考](https://docs.awspaas.com/reference-guide/aws-paas-process-reference-guide/process_structure/activities.html)

## 获取项目

```bash
go get github.com/redshiftkeying/slowflow-server
```

## 直接使用

### 1.初始化工作流引擎（仅引擎）

```go
package main
import (
    sw "github.com/redshiftkeying/slowflow-server"
    "github.com/redshiftkeying/slowflow-server/engine"
    _ "github.com/go-sql-driver/mysql"
)
func main() {
   	//use default configs
	sw.SetServerOptions("")
	// Instance
	e := sw.InitEngine()
}
```

配置文件样例：

```yaml
# config.yml
dsn: root:123456@tcp(127.0.0.1:3306)/sflow?charset=utf8
```

### 2. 加载工作流文件

```go
	err := e.LoadFile("leave.bpmn")
	if err != nil {
		// 处理错误
	}
```

### 3. 发起流程

```go
  input := map[string]interface{}{
	"day": 1,
  }

	result, err := e.StartFlow("流程编号", "开始节点编号", "流程发起人ID", input)
	if err != nil {
		// 处理错误
	}
```

### 4. 查询待办流程列表

```go
	todos, err := e.QueryTodoFlows("流程编号", "流程待办人ID")
	if err != nil {
		// 处理错误
	}
```

### 5. 处理流程

```go
  input := map[string]interface{}{
	"action": "pass",
  }

  result, err = e.HandleFlow("待办流程节点实例ID", "流程处理人ID", input)
	if err != nil {
		// 处理错误
	}
```

### 6. 停止流程

```go
	err := e.StopFlow("待办流程节点实例ID", func(flowInstance *schema.FlowInstance) bool {
		return flowInstance.Launcher == "XXX"
	})
	if err != nil {
		// 处理错误
	}
```

### 7. 查询流程历史数据

```go
result,err := e.QueryFlowHistory("待办流程实例ID")
if err != nil {
	// 处理错误
}
```

### 8. 查询已办理的流程实例 ID 列表

```go
ids,err := e.QueryDoneFlowIDs("流程编号","流程处理人ID")
if err != nil {
	// 处理错误
}
```

### 9. 查询节点实例的候选人 ID 列表

```go
ids,err := e.QueryNodeCandidates("待办流程节点实例ID")
if err != nil {
	// 处理错误
}
```

### 10. 停止流程实例

```go
	err := e.StopFlowInstance("待办流程节点实例ID", func(flowInstance *schema.FlowInstance) bool {
		return flowInstance.Launcher == "XXX"
	})
	if err != nil {
		// 处理错误
	}
```

## 使用接口

### 1. 初始化工作流引擎

#### 默认配置

```go
package main
import (
    sw "github.com/redshiftkeying/slowflow-server"
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
func filter(c *gin.Context) {
	fmt.Printf("请求参数：%s \n", c.FullPath())
	c.Next()
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

## TODO 列表

- [x] 移除 ionic
- [x] 使用 GIN 代替 gear 框架
- [ ] 移除 qlang 的部分
- [ ] 接入 ORM 支持多种数据库
- [ ] 支持子流程
- [ ] 重构并完善客户端
- [ ] 新增 APP 客户端
- [ ] 增加丰富的例子
- [ ] 完善 BPMN2.0 的支持和实现
- [ ] 打包 Docker
