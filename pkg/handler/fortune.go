package handler

import (
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/model"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/service"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/utils"
	"net/http"
)

type FortuneHandler struct {
	service service.FortuneInterface
}

func NewFortuneHandlers(service service.FortuneInterface) *FortuneHandler {
	return &FortuneHandler{service: service}
}

func (h *FortuneHandler) ProcessURLsParentCate(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, "ProcessURLsCate")
	var urls model.Data
	if err := r.GinCtx.ShouldBindJSON(&urls); err != nil {
		log.Println("Can not Decode Urls.", err)
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if len(urls.APIs) == 0 {
		log.Fatal("No APIs provided.")
		return nil, ginext.NewError(http.StatusNotFound, "")
	}
	client := &http.Client{}
	rs, err := h.service.ProcessURLsParentCate(client, urls)
	if err != nil {
		log.Fatal("Fail to Process Service URLsCate.")
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *FortuneHandler) ProcessURLsShop(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, "ProcessURLsShop")
	var urls model.Data
	if err := r.GinCtx.ShouldBindJSON(&urls); err != nil {
		log.Println("Can not Decode Urls.", err)
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if len(urls.APIs) == 0 {
		log.Fatal("No APIs provided.")
		return nil, ginext.NewError(http.StatusNotFound, "")
	}
	client := &http.Client{}
	rs, err := h.service.ProcessURLsShop(client, urls)
	if err != nil {
		log.Fatal("Fail to Process Service URLsShop.")
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *FortuneHandler) ProcessURLsItem(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, "ProcessURLsShop")
	var urls model.Data
	if err := r.GinCtx.ShouldBindJSON(&urls); err != nil {
		log.Println("Can not Decode Urls.", err)
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if len(urls.APIs) == 0 {
		log.Fatal("No APIs provided.")
		return nil, ginext.NewError(http.StatusNotFound, "")
	}
	client := &http.Client{}
	rs, err := h.service.ProcessURLsItem(client, urls)
	if err != nil {
		log.Fatal("Fail to Process Service URLsShop.")
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *FortuneHandler) ProcessURLsItem1(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, "ProcessURLsShop")
	var urls model.Data
	if err := r.GinCtx.ShouldBindJSON(&urls); err != nil {
		log.Println("Can not Decode Urls.", err)
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if len(urls.APIs) == 0 {
		log.Fatal("No APIs provided.")
		return nil, ginext.NewError(http.StatusNotFound, "")
	}
	client := &http.Client{}
	rs, err := h.service.ProcessURLsItem1(client, urls)
	if err != nil {
		log.Fatal("Fail to Process Service URLsShop.")
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}
