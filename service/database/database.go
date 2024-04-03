package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Errors
var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrBanDoesNotExist = errors.New("Ban does not exist")
var ErrFollowDoesNotExist = errors.New("Follow does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

// User is the struct for the user
type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

// Stream is the struct for the photo stream
type Stream struct {
	Identifier uint64        `json:"identifier"`
	Photos     []PhotoStream `json:"photoStream"`
}

// PhotoStream is the struct for the photo stream
type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	Username     string `json:"username"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
	LikeStatus   bool   `json:"likeStatus"`
}

// Followers is the struct for the followers
type Followers struct {
	Id        uint64   `json:"identifier"`
	Followers []Follow `json:"Followers"`
}

// Follow is the struct for the follow
type Follow struct {
	FollowId   uint64 `json:"followId"`
	FollowedId uint64 `json:"followedId"`
	UserId     uint64 `json:"userId"`
}

// Bans is the struct for the bans
type Bans struct {
	Identifier uint64 `json:"identifier"`
	Username   string `json:"username"`
	Bans       []Ban  `json:"bans"`
}

// Ban is the struct for the ban
type Ban struct {
	BanId    uint64 `json:"banId"`
	BannedId uint64 `json:"bannedId"`
	UserId   uint64 `json:"userId"`
}

// Photos is the struct for the photos
type Photos struct {
	RequestUser uint64  `json:"requestUser"`
	Identifier  uint64  `json:"identifier"`
	Photos      []Photo `json:"photos"`
}

// Photo is the struct for the photo
type Photo struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	File          []byte `json:"file"`
	Date          string `json:"date"`
	LikesCount    int    `json:"likesCount"`
	CommentsCount int    `json:"commentsCount"`
	LikeStatus    bool   `json:"likeStatus"`
}

// Like is the struct for the like
type Like struct {
	LikeId          uint64 `json:"likeId"`
	UserIdentifier  uint64 `json:"identifier"`
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	PhotoOwner      uint64 `json:"photoOwner"`
}

// Comments is the struct for the comments
type Comments struct {
	RequestIdentifier uint64    `json:"requestIdentifier"`
	PhotoIdentifier   uint64    `json:"photoIdentifier"`
	PhotoOwner        uint64    `json:"identifier"`
	Comments          []Comment `json:"comments"`
}

// Comment is the struct for the comment
type Comment struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	PhotoId       uint64 `json:"photoId"`
	PhotoOwner    uint64 `json:"photoOwner"`
	OwnerUsername string `json:"ownerUsername"`
	Username      string `json:"username"`
	Content       string `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// User
	CreateUser(User) (User, error)
	SetUsername(User, string) (User, error)
	GetUserId(string) (User, error)
	CheckUserById(User) (User, error)
	CheckUserByUsername(User) (User, error)
	CheckUser(User) (User, error)
	GetMyStream(User) ([]PhotoStream, error)

	// Follow
	SetFollow(Follow) (Follow, error)
	RemoveFollow(uint64, uint64, uint64) error
	GetFollowingId(user1 uint64, user2 uint64) (Follow, error)
	GetFollowers(User, uint64) (Follow, error)
	GetFollowersCount(uint64) (int, error)
	GetFollowingsCount(uint64) (int, error)
	GetFollowStatus(uint64, uint64) (bool, error)

	// Ban
	AddBan(Ban) (Ban, error)
	RemoveBan(Ban) error
	GetBans(User, uint64) (Ban, error)
	GetBanById(Ban) (Ban, error)
	UpdateBanStatus(int, uint64, uint64) error
	GetBanStatus(uint64, uint64) (bool, error)
	CheckIfBanned(uint64, uint64) (bool, error)

	// Photo
	UploadPhoto(Photo) (Photo, error)
	RemovePhoto(uint64) error
	GetPhotos(User, uint64) ([]Photo, error)
	GetPhotosCount(uint64) (int, error)
	CheckPhoto(Photo) (Photo, error)

	// Like
	AddLike(Like) (Like, error)
	RemoveLike(Like) error
	RemoveLikes(uint64, uint64) error
	GetLike(uint64, uint64) (Like, error)
	GetLikeById(Like) (Like, error)
	GetLikesCount(uint64) (int, error)

	// Comment
	AddComment(Comment) (Comment, error)
	RemoveComment(Comment) error
	GetComments(photoId uint64) ([]Comment, error)
	GetCommentsCount(uint64) (int, error)
	GetCommentById(Comment) (Comment, error)
	RemoveComments(uint64, uint64) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
