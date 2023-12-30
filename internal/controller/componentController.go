package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/jacobslunga/bubblegum-server/internal/model"
  "github.com/jacobslunga/bubblegum-server/internal/service"
)

type ComponentController struct {
  service *service.ComponentService
}

func NewComponentController(service *service.ComponentService) *ComponentController {
  return &ComponentController{service: service}
}

func (c *ComponentController) GetAllComponents(ctx *gin.Context) {
  components, err := c.service.GetAllComponents()

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  ctx.JSON(http.StatusOK, components)
}

func (c *ComponentController) GetComponentById(ctx *gin.Context) {
  componentId := ctx.Param("id")

  component, err := c.service.GetComponentById(componentId)

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  ctx.JSON(http.StatusOK, component)
}

func (c *ComponentController) CreateComponent(ctx *gin.Context) {
  var component model.Component

  err := ctx.BindJSON(&component)

  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  err = c.service.CreateComponent(&component)

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  ctx.JSON(http.StatusCreated, component)
}

func (c *ComponentController) UpdateComponent(ctx *gin.Context) {
  var component model.Component

  err := ctx.BindJSON(&component)

  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  err = c.service.UpdateComponent(&component)

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  ctx.JSON(http.StatusOK, component)
}

func (c *ComponentController) DeleteComponent(ctx *gin.Context) {
  var component model.Component

  err := ctx.BindJSON(&component)

  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  err = c.service.DeleteComponent(&component)

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  ctx.JSON(http.StatusOK, component)
}
