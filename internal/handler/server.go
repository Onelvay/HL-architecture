package handler

import (
	"context"
	"github.com/Onelvay/HL-architecture/config"
	_ "github.com/Onelvay/HL-architecture/docs"
	routes "github.com/Onelvay/HL-architecture/internal/handler/http/v1"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Server struct {
	*http.Server
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost
// @BasePath	/api/v1

func NewServer(cfg config.Config, s service.Service) *Server {
	router := gin.New()

	router.Use(gin.Recovery(),
		gin.Logger(),
	)

	router.MaxMultipartMemory = 8 << 20
	routes.InitRoutes(router, s)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
