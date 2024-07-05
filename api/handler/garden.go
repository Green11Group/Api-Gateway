package handler

import (
	garden "apigateway/genproto/gardenmangement"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGarden(c *gin.Context) {
	req := garden.CreateGardenRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading from body",
		})
		return
	}

	resp, err := h.garden.CreateGarden(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetGarden(c *gin.Context) {
	id := c.Param("id")
	req := garden.GetGardenRequest{Id: id}

	resp, err := h.garden.GetGarden(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateGarden(c *gin.Context) {
	req := garden.UpdateGardenRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading from body",
		})
		return
	}

	resp, err := h.garden.UpdateGarden(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteGarden(c *gin.Context) {
	id := c.Param("id")
	req := garden.DeleteGardenRequest{Id: id}

	resp, err := h.garden.DeleteGareden(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserGardens(c *gin.Context) {
	userId := c.Param("user_id")
	req := garden.GetUserGardensRequest{UserId: userId}

	resp, err := h.garden.GetUserGardens(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreatePlant(c *gin.Context) {
	req := garden.CreatePlantRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading from body",
		})
		return
	}

	resp, err := h.garden.CreatePlant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetPlant(c *gin.Context) {
	gardenId := c.Param("garden_id")
	req := garden.GetPlantRequest{GardenId: gardenId}

	resp, err := h.garden.GetPlant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdatePlant(c *gin.Context) {
	req := garden.UpdatePlantRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading from body",
		})
		return
	}

	resp, err := h.garden.UpdatePlant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeletePlant(c *gin.Context) {
	id := c.Param("id")
	req := garden.DeletePlantRequest{Id: id}

	resp, err := h.garden.DeletePlant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateCareLog(c *gin.Context) {
	req := garden.CreateCareLogRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading from body",
		})
		return
	}

	resp, err := h.garden.CreateCareLog(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCareLog(c *gin.Context) {
	plantId := c.Param("plant_id")
	req := garden.GetCareLogRequest{PlantId: plantId}

	resp, err := h.garden.GetCareLog(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
