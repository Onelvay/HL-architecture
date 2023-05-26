package handler

import (
	"context"
	"github.com/Onelvay/HL-architecture/config"
	routes "github.com/Onelvay/HL-architecture/internal/handler/http"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(cfg config.Config, s service.Service) *Server {
	router := gin.New()

	router.Use(gin.Recovery(),
		gin.Logger(),
	)

	router.MaxMultipartMemory = 8 << 20
	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	routes.InitRoutes(router, s)

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
