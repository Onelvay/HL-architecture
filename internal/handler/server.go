package handler

import (
	"context"
	"github.com/Onelvay/HL-architecture/config"
	routes "github.com/Onelvay/HL-architecture/internal/handler/http/v1"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

type Server struct {
	*http.Server
}

// @title Testing App
// @version 1.0
// description -

// @host localhost:8080
// @BasePath /
func NewServer(cfg config.Config, s service.Service) *Server {
	router := gin.New()

	router.Use(gin.Recovery(),
		gin.Logger(),
	)

	router.MaxMultipartMemory = 8 << 20
	routes.InitRoutes(router, s)
	//
	//docs.SwaggerInfo.BasePath = "/"
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	caw := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedOrigins: []string{"http://localhost:4200"},
	}).Handler(router)

	return &Server{
		&http.Server{
			Addr:           ":" + cfg.Http.Port,
			Handler:        caw,
			ReadTimeout:    cfg.Http.ReadTimeout,
			WriteTimeout:   cfg.Http.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}

}

func (s Server) Run() (err error) {
	return s.ListenAndServe()
}

func (s Server) Stop(ctx context.Context) (err error) {
	return s.Shutdown(ctx)
}
