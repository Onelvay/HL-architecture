package rest

import (
	"context"
	"fmt"
	"github.com/Onelvay/HL-architecture/config"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(cfg config.Config, s service.Service) *Server {
	router := gin.New()

	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 16 << 20
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	p := NewProductHandler(s.Product)
	productRoutes(router, p)

	a := NewAuthorizationHandler()
	authRoutes(router, a)

	return &Server{
		&http.Server{
			Addr:           ":" + cfg.Http.Port,
			Handler:        router,
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
