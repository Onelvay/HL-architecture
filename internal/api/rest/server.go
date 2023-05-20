package rest

import (
	"context"
	"github.com/Onelvay/HL-architecture/config"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(cfg config.Config, s service.Service) *Server {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 16 << 20

	p := NewProductHandler(s.Product)
	productRoutes(router, p)

	return &Server{
		&http.Server{
			Addr:         ":" + cfg.Http.Port,
			Handler:      router,
			ReadTimeout:  cfg.Http.ReadTimeout,
			WriteTimeout: cfg.Http.WriteTimeout,
		},
	}

}
func (s Server) Run() (err error) {
	return s.ListenAndServe()
}
func (s Server) Stop(ctx context.Context) (err error) {
	return s.Shutdown(ctx)
}
