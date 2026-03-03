package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/My-second-free-organization/platform-api/internal/config"
)

type TaskHandler struct { cfg *config.Config }
func NewTaskHandler(cfg *config.Config) *TaskHandler { return &TaskHandler{cfg: cfg} }
func (h *TaskHandler) Get(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"id": c.Param("id")}) }
func (h *TaskHandler) Complete(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"id": c.Param("id"), "status": "completed"}) }
