package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stat4market_task/model"
	"stat4market_task/service"
)

type Event interface {
	Post(c *gin.Context)
}

type event struct {
	svc service.Event
}

func EventHandler(svc service.Event) Event {
	return &event{svc: svc}
}

// Post
// @Summary Post event
// @Description Post event
// @Param eventInput body model.Event true "Event Info"
// @Produce json
// @Success 201
// @Router /api/event [post]
func (h *event) Post(c *gin.Context) {
	var eventItem model.Event
	if err := c.BindJSON(&eventItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Errorf("handler.Post.BindJson: %v", err),
		})
		return
	}

	if err := h.svc.Create(c, eventItem); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": fmt.Errorf("handler.Post.Create: %v", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "event created successfully",
	})
}
