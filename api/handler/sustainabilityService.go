package handler

import (
	"context"
	"net/http"

	sustainability "apigateway/genproto/SustainabilityService"
	"github.com/gin-gonic/gin"
)

func (h *Handler) LogImpact(c *gin.Context) {
	req := sustainability.LogImpactRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.sustainability.LogImpact(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserImpact(c *gin.Context) {
	id := c.Param("user_id")
	req := sustainability.GetUserImpactRequest{UserId: id}

	resp, err := h.sustainability.GetUserImpact(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCommunityImpact(c *gin.Context) {
	id := c.Param("community_id")
	req := sustainability.GetCommunityImpactRequest{CommunityId: id}

	resp, err := h.sustainability.GetCommunityImpact(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetChallenges(c *gin.Context) {
	req := sustainability.GetChallengesRequest{}

	resp, err := h.sustainability.GetChallenges(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinChallenge(c *gin.Context) {
	req := sustainability.JoinChallengeRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.sustainability.JoinChallenge(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateChallengeProgress(c *gin.Context) {
	req := sustainability.UpdateChallengeProgressRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.sustainability.UpdateChallengeProgress(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserChallenges(c *gin.Context) {
	id := c.Param("user_id")
	req := sustainability.GetUserChallengesRequest{UserId: id}

	resp, err := h.sustainability.GetUserChallenges(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserLeaderboard(c *gin.Context) {
	req := sustainability.GetUserLeaderboardRequest{}

	resp, err := h.sustainability.GetUserLeaderboard(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCommunityLeaderboard(c *gin.Context) {
	req := sustainability.GetCommunityLeaderboardRequest{}

	resp, err := h.sustainability.GetCommunityLeaderboard(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) PostChallenges(c *gin.Context) {
	req := sustainability.PostChallengesRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.sustainability.PostChallenges(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
