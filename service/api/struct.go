package api

// struct for the Profile
type Profile struct {
	Id             uint64 `json:"userId"`
	requesterId    uint64 `json:"requesterId"`
	Username       string `json:"username"`
	FollowerCount  int    `json:"followerCount"`
	FollowingCount int    `json:"followingCount"`
	PostCount      int    `json:"postCount"`
	IsFollowing    bool   `json:"isFollowing"`
	checkBanned    bool   `json:"checkBanned"`
	banned         bool   `json:"banned"`
}

// struct for the User
type User struct {
	Id       string `json:"userId"`
	Username string `json:"username"`
}

type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

type Follow struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"userId"`
	FollowId uint64 `json:"followId"`
}

type Like struct {
	LikeId     uint64 `json:"id"`
	UserId     uint64 `json:"userId"`
	PhotoId    uint64 `json:"photoId"`
	PhotoOwner string `json:"likeType"`
}

type Ban struct {
	RequestingBanId uint64 `json:"id"`
	UserId          uint64 `json:"userId"`
	BanId           uint64 `json:"banId"`
}
