package workflow

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redshiftkeying/slowflow-server/api"
	"github.com/redshiftkeying/slowflow-server/engine"
	"github.com/redshiftkeying/slowflow-server/service/db"
	"github.com/spf13/viper"
)

// AddServerMiddleware 中间件
func (a *Server) AddServerMiddleware(middlewares ...gin.HandlerFunc) {
	for _, m := range middlewares {
		a.app.Use(m)
	}
}

// StartServer 启动管理服务
func (a *Server) StartServer() {
	port := viper.GetString("port")
	if port == "" {
		port = "6062"
	}
	a.app.Run(":" + port)
}

// Server 流程管理服务
type Server struct {
	engine *engine.Engine
	app    *gin.Engine
}

// Init 初始化
func Init() *Server {
	a := new(Server)
	// init engine
	dsn := viper.GetString("dsn")
	a.engine = engine.Init(db.SetDSN(dsn), db.SetTrace(true))

	// debug ? debug:release
	if viper.GetString("mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	a.app = gin.Default()
	// set routers
	setStatic(a)
	setAPI(a)
	return a
}

func setStatic(a *Server) {
	staticRoot := viper.GetString("static.root")
	staticPath := viper.GetString("static.path")
	if staticRoot != "" {
		if staticPath != "" {
			a.app.Static(staticPath, staticRoot)
		}
	}
}

func setAPI(a *Server) {
	// rest api config
	restConfig := viper.Sub("api.rest")
	if restConfig != nil {
		var rests []api.RestConfig
		err := restConfig.Unmarshal(&rests)
		if err != nil {
			log.Printf("unable to decode into struct, %v \n", err)
			return
		}
	}
	// graphql api config
	graphqlConfig := viper.Sub("api.graphql")
	if graphqlConfig != nil {
		var graphqls []api.GraphqlConfig
		err := restConfig.Unmarshal(&graphqls)
		if err != nil {
			log.Printf("unable to decode into struct, %v \n", err)
			return
		}
	}
	// protobuf api config
	protobufConfig := viper.Sub("api.protobuf")
	if protobufConfig != nil {
		var protobufs []api.ProtobufConfig
		err := restConfig.Unmarshal(&protobufs)
		if err != nil {
			log.Printf("unable to decode into struct, %v \n", err)
			return
		}
	}
}
