package route

import (
	"github.com/gin-contrib/pprof"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/service"
	handlerFortune "gitlab.com/merakilab9/meracrawler/fortune/pkg/handler"
	serviceFortune "gitlab.com/merakilab9/meracrawler/fortune/pkg/service"
	"net/http"
)

type Service struct {
	*service.BaseApp
}

func NewService() *Service {
	s := &Service{
		service.NewApp("fortuneService", "v1.0"),
	}
	client := &http.Client{}
	fortuneService := serviceFortune.NewFortuneService(client)
	fortuneHandle := handlerFortune.NewFortuneHandlers(fortuneService)

	pprof.Register(s.Router)

	v1Api := s.Router.Group("/api/v1")
	v1Api.POST("/cate/parse-to-json", ginext.WrapHandler(fortuneHandle.ProcessURLsParentCate))
	v1Api.POST("/shop/parse-to-json", ginext.WrapHandler(fortuneHandle.ProcessURLsShop))
	v1Api.POST("/item/parse-to-json", ginext.WrapHandler(fortuneHandle.ProcessURLsItem))
	v1Api.POST("/item1/parse-to-json", ginext.WrapHandler(fortuneHandle.ProcessURLsItem1))

	return s
}
