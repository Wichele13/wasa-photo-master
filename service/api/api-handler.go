package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// userActions
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/user/:username/setUsername", rt.wrap(rt.setMyUserName))
	rt.router.GET("/user/:username/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:username/profile", rt.wrap(rt.getUserProfile))

	// PhotoActions
	rt.router.PUT("/users/:username/photo/:photoId", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:username/photo/:photoId", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:username/photo", rt.wrap(rt.getPhotos))

	// BanActions
	rt.router.PUT("/users/:username/ban/:banId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/ban/:banId", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:username/ban", rt.wrap(rt.getBansList))

	// FollowActions
	rt.router.PUT("/users/:username/follow/:followId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/follow/:followId", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:username/follow", rt.wrap(rt.getFollowers))

	// LikeActions
	rt.router.PUT("/users/:username/photo/:photoId/like/:likeId", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/photo/:photoId/like/:likeId", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/users/:username/photo/:photoId/like", rt.wrap(rt.getLikes))

	// CommentActions
	rt.router.PUT("/users/:username/photo/:photoId/comment/:commentId", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:username/photo/:photoId/comment/:commentId", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/users/:username/photo/:photoId/comment", rt.wrap(rt.getComments))

	// Liveness
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
