package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.routes.POST("/session", rt.doLogin)

	rt.routes.POST("/user/:singleusername", rt.setMyUserName)

	rt.routes.POST("/user/:singleusername/photos/:singlephoto", rt.uploadPhoto)

	rt.routes.POST("/user/:singleusername/photos/:singlephoto", rt.deletePhoto)

	rt.routes.PUT("/user/:singleusername/photos/:singlephoto/likes/:singlelike", rt.likePhoto)

	rt.routes.DELETE("/user/:singleusername/photos/:singlephoto/likes/:singlelike", rt.unlikePhoto)

	rt.routes.PUT("/user/:singleusername/photos/:singlephoto/comments/:singlecomment", rt.commentPhoto)

	rt.routes.DELETE("/user/:singleusername/photos/:singlephoto/comments/:singlecomment", rt.uncommentPhoto)

	rt.routes.PUT("/user/:singleusername/banned/:singlebanneduser", rt.banUser)

	rt.routes.DELETE("/user/:singleusername/banned/:singlebanneduser", rt.unbanUser)

	rt.routes.PUT("/user/:singleusername/followed/:singlefolloweduser", rt.followUser)

	rt.routes.DELETE("/user/:singleusername/followed/:singlefolloweduser", rt.unfollowUser)

	rt.routes.GET("/user/:singleusername/profile", rt.getUserProfile)

	rt.routes.GET("/user/:singleusername/stream", rt.getMyStream)

	// Special routes (?)
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
