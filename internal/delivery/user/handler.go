package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/morris-zheng/go-slim-micro-api/internal/common/response"
	"github.com/morris-zheng/go-slim-micro-api/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/morris-zheng/go-slim-micro-usersvc/export/user"
)

type Handler struct {
	svc *domain.ServiceContext
}

func NewHandler(svc *domain.ServiceContext) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Logger.Info(c, fmt.Sprintf("get user:  %d", id))

	u, err := h.svc.UserCli.Get(c, &user.UserById{Id: 5})
	if err != nil {
		response.Fail(c, response.Response{
			Msg:      err.Error(),
			Code:     10404,
			HttpCode: http.StatusNotFound,
		})
		return
	}

	response.Success(c, response.Response{
		Data: u,
	})
}
