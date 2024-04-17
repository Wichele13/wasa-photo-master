package api


import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)


// struct for the Profile
type Profile struct {
	Id             uint64 `json:"userId"`
	RequesterId    uint64 `json:"requesterId"`
	Username       string `json:"username"`
	FollowerCount  int    `json:"followerCount"`
	FollowingCount int    `json:"followingCount"`
	PostCount      int    `json:"postCount"`
	IsFollowing    bool   `json:"isFollowing"`
	CheckBanned    bool   `json:"checkBanned"`
	Banned         bool   `json:"banned"`
}

// struct for the User
type User struct {
	Id       string `json:"userId"`
	Username string `json:"username"`
}

// struct for the PhotoStream
type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

// struct for the Follow
type Follow struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"userId"`
	FollowId uint64 `json:"followId"`
}

// struct for the Like
type Like struct {
	LikeId     uint64 `json:"id"`
	UserId     uint64 `json:"userId"`
	PhotoId    uint64 `json:"photoId"`
	PhotoOwner string `json:"likeType"`
}

// struct for the Ban
type Ban struct {
	RequestingBanId uint64 `json:"id"`
	UserId          uint64 `json:"userId"`
	BanId           uint64 `json:"banId"`
}

// struct for the Photo
type Photo struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"like"`
	CommentCount int    `json:"comment"`
}

// struct for the Comment

type Comment struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"userId"`
	PhotoId    uint64 `json:"photoId"`
	PhotoOwner string `json:"photoOwner"`
	Comment    string `json:"comment"`
}


func (u *User) FromDatabase(user database.User) {
	u.Id = user.Id
	u.Username = user.Username
}

func (u *User) ToDatabase() database.User {
	return database.User{
		Id:       u.Id,
		Username: u.Username,
	}
}