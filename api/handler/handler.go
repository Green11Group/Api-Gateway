package handler

import (
	garden "apigateway/genproto/gardenmangement"
	community "apigateway/genproto/community"
	user "apigateway/genproto/users"
	sustainability "apigateway/genproto/SustainabilityService"
)
type Handler struct{
	garden garden.GardenServiceClient
	community community.CommunityServiceClient
	user user.UserClient
	sustainability sustainability.SustainabilityServiceClient
}

func NewHandler(garden garden.GardenServiceClient,community community.CommunityServiceClient,user user.UserClient,sustainability sustainability.SustainabilityServiceClient) *Handler{
	return &Handler{garden: garden,community: community,user: user,sustainability: sustainability}
}