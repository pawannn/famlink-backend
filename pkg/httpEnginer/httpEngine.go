package httpEngine

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	metadb "github.com/pawannn/famly/core/services/metaDB"
	"github.com/pawannn/famly/core/services/sms"
	appconfig "github.com/pawannn/famly/pkg/appConfig"
	token "github.com/pawannn/famly/port/token"
)

type FamlyRoute struct {
	Route       string
	Method      string
	Middleware  []gin.HandlerFunc
	Controller  gin.HandlerFunc
	Description string
}

type FamlyEngine struct {
	config appconfig.Config
	Router *gin.Engine
	DB     *sql.DB
	Token  *token.TokenPort
	MetaDB metadb.MetaDBService
	Sms    sms.SmsService
}

func InitfamlyEngine(c appconfig.Config, DB *sql.DB, tS token.TokenPort, mDb metadb.MetaDBService, sms sms.SmsService) *FamlyEngine {
	g := gin.Default()
	fE := FamlyEngine{
		config: c,
		Router: g,
		MetaDB: mDb,
		DB:     DB,
		Token:  &tS,
		Sms:    sms,
	}
	return &fE
}

func (fE *FamlyEngine) AddRoute(routes []FamlyRoute) {
	for _, route := range routes {
		fmt.Printf("%s : %s : %s\n", route.Method, route.Route, route.Description)
		handlers := append(route.Middleware, route.Controller)
		switch route.Method {
		case "GET":
			fE.Router.GET(route.Route, handlers...)
		case "POST":
			fE.Router.POST(route.Route, handlers...)
		case "PUT":
			fE.Router.PUT(route.Route, handlers...)
		case "DELETE":
			fE.Router.DELETE(route.Route, handlers...)
		case "PATCH":
			fE.Router.PATCH(route.Route, handlers...)
		default:
			fmt.Printf("Unsupported method %s for route %s\n", route.Method, route.Route)
		}
	}
}

func (fE *FamlyEngine) StartServer() {
	addr := fmt.Sprintf(":%d", fE.config.App_port)
	fE.Router.Run(addr)
}
