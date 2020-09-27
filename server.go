package workflow

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redshiftkeying/slowflow-server/api"
	"github.com/redshiftkeying/slowflow-server/engine"
)

type serverOptions struct {
	prefix      string
	staticRoot  string
	middlewares []gin.HandlerFunc
}

// ServerOption 流程服务配置
type ServerOption func(*serverOptions)

// ServerPrefixOption 访问前缀
func ServerPrefixOption(prefix string) ServerOption {
	return func(opts *serverOptions) {
		opts.prefix = prefix
	}
}

// ServerStaticRootOption 静态文件目录
func ServerStaticRootOption(staticRoot string) ServerOption {
	return func(opts *serverOptions) {
		opts.staticRoot = staticRoot
	}
}

// ServerMiddlewareOption 中间件
func ServerMiddlewareOption(middlewares ...gin.HandlerFunc) ServerOption {
	return func(opts *serverOptions) {
		opts.middlewares = middlewares
	}
}

// Server 流程管理服务
type Server struct {
	opts   serverOptions
	engine *engine.Engine
	app    *gin.Engine
}

// Init 初始化
func (a *Server) Init(engine *engine.Engine, opts ...ServerOption) *Server {
	a.engine = engine

	var o serverOptions
	for _, opt := range opts {
		opt(&o)
	}

	if o.prefix == "" {
		o.prefix = "/"
	}
	a.opts = o

	app := gin.Default()

	for _, m := range a.opts.middlewares {
		app.Use(m)
	}

	if a.opts.staticRoot != "" {
		app.Static("/flow", a.opts.staticRoot)
	}

	a.app = app
	newRouterMiddleware(a)

	return a
}

func (a *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.app.ServeHTTP(w, r)
}

func newRouterMiddleware(srv *Server) {
	router := srv.app
	// router := gear.NewRouter(gear.RouterOptions{
	// 	Root:       fmt.Sprintf("%sapi", srv.opts.prefix),
	// 	IgnoreCase: true,
	// })

	// restfull api
	restAPI := new(api.RestAPI).Init(srv.engine)
	router.GET("/api/v1/flow", restAPI.QueryFlowPage)
	router.GET("/api/v1/flow/:id", restAPI.GetFlow)
	router.DELETE("/api/v1/flow/:id", restAPI.DeleteFlow)
	router.POST("/api/v1/flow", restAPI.SaveFlow)
}

// StartServer 启动管理服务
func StartServer(engine *engine.Engine, opts ...ServerOption) http.Handler {
	srv := new(Server).Init(engine, opts...)
	return srv
}
