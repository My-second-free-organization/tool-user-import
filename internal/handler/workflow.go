package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/My-second-free-organization/platform-api/internal/config"
)

type WorkflowHandler struct { cfg *config.Config }
func NewWorkflowHandler(cfg *config.Config) *WorkflowHandler { return &WorkflowHandler{cfg: cfg} }
func (h *WorkflowHandler) Create(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"message": "created"}) }
func (h *WorkflowHandler) List(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"workflows": []interface{}{}}) }
func (h *WorkflowHandler) Get(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"id": c.Param("id")}) }
func (h *WorkflowHandler) Delete(c *gin.Context) { c.Status(http.StatusNoContent) }
