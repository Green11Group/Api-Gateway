package api

import (
	"apigateway/api/handler"
	garden "apigateway/genproto/gardenmangement"
	community "apigateway/genproto/community"
	user "apigateway/genproto/users"
	sustainability "apigateway/genproto/SustainabilityService"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conngar,conncom,connuser,connsus *grpc.ClientConn) *gin.Engine{
	router:=gin.Default()

	garden:=garden.NewGardenServiceClient(conngar)
	community:=community.NewCommunityServiceClient(conncom)
	users:=user.NewUserClient(connuser)
	sustainibility:=sustainability.NewSustainabilityServiceClient(connsus)

	
	handlers:=handler.NewHandler(garden,community,users,sustainibility)

	// Garden Service apilari
	router.POST("/creategarden", handlers.CreateGarden)
	router.GET("/getgarden/:id", handlers.GetGarden)
	router.PUT("/updategarden", handlers.UpdateGarden)
	router.DELETE("/deletegarden/:id", handlers.DeleteGarden)
	router.GET("/usergardens/:user_id", handlers.GetUserGardens)

	router.POST("/createplant", handlers.CreatePlant)
	router.GET("/getplant/:garden_id", handlers.GetPlant)
	router.PUT("/updateplant", handlers.UpdatePlant)
	router.DELETE("/deleteplant/:id", handlers.DeletePlant)

	router.POST("/createcarelog", handlers.CreateCareLog)
	router.GET("/getcarelog/:plant_id", handlers.GetCareLog)


	// Community Service apilari
	router.POST("/createcommunity", handlers.CreateCommunity)
	router.GET("/getcommunity/:community_id", handlers.GetCommunityByID)
	router.PUT("/updatecommunity", handlers.UpdateCommunityByID)
	router.DELETE("/deletecommunity/:community_id", handlers.DeleteCommunityByID)
	router.GET("/getcommunities", handlers.GetAllCommunities)

	router.POST("/createcommunity/member", handlers.JoinCommunityMember)
	router.DELETE("/deletecommunity/member", handlers.LeftCommunityMember)
	router.POST("/updatecommunity/event", handlers.JoinCommunityEvent)
	router.GET("/getevent/:event_id", handlers.GetEventByID)

	router.POST("/createcommunity/forum", handlers.JoinCommunityForum)
	router.GET("/getforum/:forum_id", handlers.GetForumByID)
	router.POST("/updateforum/comment", handlers.AddCommentForum)
	router.GET("/getforum/comment/:comment_id", handlers.GetForumCommentByID)

	// User service apilari
	router.GET("/user/:user_id", handlers.GetUser)
	router.PUT("/user", handlers.UpdateUser)
	router.DELETE("/user/:user_id", handlers.DeleteUser)
	router.GET("/user/profile/:user_id", handlers.GetUserProfile)
	router.PUT("/user/profile", handlers.UpdateUserProfile)

	// sustainibility service apilari
	router.POST("/api/impact/log", handlers.LogImpact)
	router.GET("/api/users/:user_id/impact", handlers.GetUserImpact)
	router.GET("/api/communities/:community_id/impact", handlers.GetCommunityImpact)
	router.GET("/api/challenges", handlers.GetChallenges)
	router.POST("/api/challenges/:id/join", handlers.JoinChallenge)
	router.PUT("/api/challenges/:id/progress", handlers.UpdateChallengeProgress)
	router.GET("/api/users/:user_id/challenges", handlers.GetUserChallenges)
	router.GET("/api/leaderboard/users", handlers.GetUserLeaderboard)
	router.GET("/api/leaderboard/communities", handlers.GetCommunityLeaderboard)
	router.POST("/api/challenges", handlers.PostChallenges)

	return router
}