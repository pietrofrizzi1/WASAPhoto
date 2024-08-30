package api

import (
	"log"
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	log.Print("ciao")
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	rt.router.POST("/user/:singleusername", rt.wrap(rt.setMyUserName))

	rt.router.POST("/user/:singleusername/photos/:singlephoto", rt.wrap(rt.uploadPhoto))

	rt.router.DELETE("/user/:singleusername/photos/:singlephoto", rt.wrap(rt.deletePhoto))

	rt.router.PUT("/user/:singleusername/photos/:singlephoto/likes/:singlelike", rt.wrap(rt.likePhoto))

	rt.router.DELETE("/user/:singleusername/photos/:singlephoto/likes/:singlelike", rt.wrap(rt.unlikePhoto))

	rt.router.PUT("/user/:singleusername/photos/:singlephoto/comments/:singlecomment", rt.wrap(rt.commentPhoto))

	rt.router.DELETE("/user/:singleusername/photos/:singlephoto/comments/:singlecomment", rt.wrap(rt.uncommentPhoto))

	rt.router.PUT("/user/:singleusername/banned/:singlebanneduser", rt.wrap(rt.banUser))

	rt.router.DELETE("/user/:singleusername/banned/:singlebanneduser", rt.wrap(rt.unbanUser))

	rt.router.PUT("/user/:singleusername/followed/:singlefolloweduser", rt.wrap(rt.followUser))

	rt.router.DELETE("/user/:singleusername/followed/:singlefolloweduser", rt.wrap(rt.unfollowUser))

	rt.router.GET("/user/:singleusername/profile", rt.wrap(rt.getUserProfile))

	rt.router.GET("/user/:singleusername/stream", rt.wrap(rt.getMyStream))

	// Special routes (?)
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
