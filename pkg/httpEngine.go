package pkg

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	appconfig "github.com/pawannn/famlink/pkg/appConfig"
)

type FamLinkRoute struct {
	Route       string
	Method      string
	Middleware  []gin.HandlerFunc
	Controller  gin.HandlerFunc
	Description string
}

type FamLinkEngine struct {
	config appconfig.Config
	Router *gin.Engine
	DB     *sql.DB
}

func InitFamLinkEngine(c appconfig.Config, DB *sql.DB) *FamLinkEngine {
	g := gin.Default()
	fE := FamLinkEngine{
		config: c,
		Router: g,
	}
	return &fE
}

func (fE *FamLinkEngine) AddRoute(routes []FamLinkRoute) {
	for _, route := range routes {
		fmt.Printf("%s : %s : %s\n", route.Method, route.Route, route.Description)
		switch route.Method {
		case "GET":
			handlers := append(route.Middleware, route.Controller)
			fE.Router.GET(route.Route, handlers...)
		case "POST":
			handlers := append(route.Middleware, route.Controller)
			fE.Router.POST(route.Route, handlers...)
		case "PUT":
			handlers := append(route.Middleware, route.Controller)
			fE.Router.PUT(route.Route, handlers...)
		case "DELETE":
			handlers := append(route.Middleware, route.Controller)
			fE.Router.DELETE(route.Route, handlers...)
		case "PATCH":
			handlers := append(route.Middleware, route.Controller)
			fE.Router.PATCH(route.Route, handlers...)
		default:
			fmt.Printf("Unsupported method %s for route %s\n", route.Method, route.Route)
		}
	}
}

func (fE *FamLinkEngine) StartServer() {
	addr := fmt.Sprintf(":%d", fE.config.App_port)
	fE.Router.Run(addr)
}
