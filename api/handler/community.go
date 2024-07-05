package handler

import (
	community "apigateway/genproto/community"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCommunity(c *gin.Context) {
	req := community.CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.CreateCommunity(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCommunityByID(c *gin.Context) {
	id := c.Param("community_id")
	req := community.GetCommunityRequest{CommunityID: id}

	resp, err := h.community.GetCommunityByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateCommunityByID(c *gin.Context) {
	req := community.UpdateCommunityRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.UpdateCommunityByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteCommunityByID(c *gin.Context) {
	id := c.Param("community_id")
	req := community.DeleteCommunityRequest{CommunityID: id}

	resp, err := h.community.DeleteCommunityByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllCommunities(c *gin.Context) {
	req := community.GetCommunitiesRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.GetAllCommunities(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinCommunityMember(c *gin.Context) {
	req := community.JoinMemberRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.JoinCommunityMember(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) LeftCommunityMember(c *gin.Context) {
	req := community.LeftMemberRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.LeftCommunityMember(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinCommunityEvent(c *gin.Context) {
	req := community.JoinEventRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.JoinCommunityEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetEventByID(c *gin.Context) {
	id := c.Param("event_id")
	req := community.GetEventsRequest{EventID: id}

	resp, err := h.community.GetEventByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinCommunityForum(c *gin.Context) {
	req := community.JoinForumRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.JoinCommunityForum(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetForumByID(c *gin.Context) {
	id := c.Param("forum_id")
	req := community.GetForumRequest{ForumID: id}

	resp, err := h.community.GetForumByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) AddCommentForum(c *gin.Context) {
	req := community.AddCommentRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading from body"})
		return
	}

	resp, err := h.community.AddCommentForum(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetForumCommentByID(c *gin.Context) {
	id := c.Param("comment_id")
	req := community.GetForumCommentRequest{UserID: id}

	resp, err := h.community.GetForumCommentByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
