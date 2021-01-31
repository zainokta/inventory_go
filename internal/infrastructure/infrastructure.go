package infrastructure

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type GinServerMode int

type GinServer struct {
	port   int
	Router *gin.Engine
}

func NewServer(port int, mode string) GinServer {
	s := GinServer{}
	s.port = port
	s.Router = gin.New()

	switch mode {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "TEST":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s.Router.Use(gin.Recovery())

	SetCors(s.Router, "*")

	return s
}

func SetCors(engine *gin.Engine, allowedOrigins string) {
	engine.Use(cors.Middleware(cors.Config{
		Origins:         allowedOrigins,
		Methods:         strings.Join([]string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions, http.MethodPatch}, ","),
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}

func (s GinServer) Start() {
	s.Router.Run(":" + strconv.Itoa(s.port))
}
