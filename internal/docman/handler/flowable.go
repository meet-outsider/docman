package handler

import (
	"docman/pkg/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlowableHandler struct {
	flowable *kit.Flowable
}

func NewFlowableHandler(handler *kit.Flowable) *FlowableHandler {
	return &FlowableHandler{
		flowable: handler,
	}
}

// @description 流程相关接口
func (s *FlowableHandler) GetUsers(ctx *gin.Context) {
	resp, err := s.flowable.DoRequest("GET", "/service/identity/users", nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	var responseBody []byte
	_, err = resp.Body.Read(responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Data(http.StatusOK, resp.Header.Get("Content-Type"), responseBody)
}
