package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:singleusername", rt.wrap(rt.setMyUserName))

	rt.router.PUT("/users/:singleusername/photos/:singlephoto", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:singleusername/photos/:singlephoto", rt.wrap(rt.deletePhoto))

	rt.router.PUT("/users/:singleusername/photos/:singlephoto/likes/:singlelike", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:singleusername/photos/:singlephoto/likes/:singlelike", rt.wrap(rt.unlikePhoto))

	rt.router.PUT("/users/:singleusername/photos/:singlephoto/comments/:singlecomment", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:singleusername/photos/:singlephoto/comments/:singlecomment", rt.wrap(rt.uncommentPhoto))

	rt.router.PUT("/users/:singleusername/banned/:singlebanneduser", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:singleusername/banned/:singlebanneduser", rt.wrap(rt.unbanUser))

	rt.router.PUT("/users/:singleusername/followed/:singlefolloweduser", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:singleusername/followed/:singlefolloweduser", rt.wrap(rt.unfollowUser))

	rt.router.GET("/users/:singleusername/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:singleusername/stream", rt.wrap(rt.getMyStream))

	rt.router.GET("/users/:singleusername/photos", rt.wrap(rt.getPhotos))

	rt.router.GET("/users/:singleusername/banned", rt.wrap(rt.getBan))

	rt.router.GET("/users/:singleusername/followed", rt.wrap(rt.getFollowers))

	rt.router.GET("/users/:singleusername/photos/:singlephoto/comments", rt.wrap(rt.getComments))

	rt.router.GET("/users/:singleusername/search/:searchedusername", rt.wrap(rt.searchUsers))

	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
