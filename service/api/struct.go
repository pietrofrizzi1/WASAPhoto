package api

import (
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

// Profile rappresenta un profilo utente
type Profile struct {
	RequestId      uint64 `json:"requestId"`
	Id             uint64 `json:"id"`
	Username       string `json:"username"`
	FollowersCount int    `json:"followersCount"`
	FollowingCount int    `json:"followingCount"`
	FollowStatus   bool   `json:"followStatus"`
	BanStatus      bool   `json:"banStatus"`
	CheckIfBanned  bool   `json:"checkIfBanned"`
}

// User rappresenta un utente
type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

// PhotoStream rappresenta uno stream di foto
type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

// Follow rappresenta una relazione di follow
type Follow struct {
	FollowId   uint64 `json:"followId"`
	FollowedId uint64 `json:"followedId"`
	UserId     uint64 `json:"userId"`
}

// Ban rappresenta un ban
type Ban struct {
	BanId    uint64 `json:"banId"`
	BannedId uint64 `json:"bannedId"`
	UserId   uint64 `json:"userId"`
}

// Photo rappresenta una foto
type Photo struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	File          []byte `json:"file"`
	Date          string `json:"date"`
	LikesCount    int    `json:"likesCount"`
	CommentsCount int    `json:"commentsCount"`
}

// Like rappresenta un like
type Like struct {
	LikeId          uint64 `json:"likeId"`
	UserIdentifier  uint64 `json:"identifier"`
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	PhotoOwner      uint64 `json:"photoOwner"`
}

// Comment rappresenta un commento
type Comment struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"userId"`
	PhotoId    uint64 `json:"photoId"`
	PhotoOwner uint64 `json:"photoOwner"`
	Content    string `json:"content"`
}

// Metodi di Conversione (Conversion Methods)

// Metodi di conversione per User
func (u *User) ConvertForApplication(user database.User) {
	u.Id = user.Id
	u.Username = user.Username
}

func (u *User) ConvertForDatabase() database.User {
	return database.User{
		Id:       u.Id,
		Username: u.Username,
	}
}

// Metodi di conversione per PhotoStream
func (s *PhotoStream) PhotoStreamConvertForApplication(photoStream database.PhotoStream) {
	s.Id = photoStream.Id
	s.UserId = photoStream.UserId
	s.File = photoStream.File
	s.Date = photoStream.Date
	s.LikeCount = photoStream.LikeCount
	s.CommentCount = photoStream.CommentCount
}

func (s *PhotoStream) PhotoStreamConvertForDatabase() database.PhotoStream {
	return database.PhotoStream{
		Id:           s.Id,
		UserId:       s.UserId,
		File:         s.File,
		Date:         s.Date,
		LikeCount:    s.LikeCount,
		CommentCount: s.CommentCount,
	}
}

// Metodi di conversione per Follow
func (f *Follow) FollowConvertForApplication(follow database.Follow) {
	f.FollowId = follow.FollowId
	f.FollowedId = follow.FollowedId
	f.UserId = follow.UserId
}

func (f *Follow) FollowConvertForDatabase() database.Follow {
	return database.Follow{
		FollowId:   f.FollowId,
		FollowedId: f.FollowedId,
		UserId:     f.UserId,
	}
}

// Metodi di conversione per Ban
func (b *Ban) BanConvertForApplication(ban database.Ban) {
	b.BanId = ban.BanId
	b.BannedId = ban.BannedId
	b.UserId = ban.UserId
}

func (b *Ban) BanConvertForDatabase() database.Ban {
	return database.Ban{
		BanId:    b.BanId,
		BannedId: b.BannedId,
		UserId:   b.UserId,
	}
}

// Metodi di conversione per Photo
func (p *Photo) PhotoConvertForApplication(photo database.Photo) {
	p.Id = photo.Id
	p.UserId = photo.UserId
	p.File = photo.File
	p.Date = photo.Date
	p.LikesCount = photo.LikesCount
	p.CommentsCount = photo.CommentsCount
}

func (p *Photo) PhotoConvertForDatabase() database.Photo {
	return database.Photo{
		Id:            p.Id,
		UserId:        p.UserId,
		File:          p.File,
		Date:          p.Date,
		LikesCount:    p.LikesCount,
		CommentsCount: p.CommentsCount,
	}
}

// Metodi di conversione per Like
func (l *Like) LikeConvertForApplication(like database.Like) {
	l.LikeId = like.LikeId
	l.UserIdentifier = like.UserIdentifier
	l.PhotoIdentifier = like.PhotoIdentifier
	l.PhotoOwner = like.PhotoOwner
}

func (l *Like) LikeConvertForDatabase() database.Like {
	return database.Like{
		LikeId:          l.LikeId,
		UserIdentifier:  l.UserIdentifier,
		PhotoIdentifier: l.PhotoIdentifier,
		PhotoOwner:      l.PhotoOwner,
	}
}

// Metodi di conversione per Comment
func (c *Comment) CommentConvertForApplication(comment database.Comment) {
	c.Id = comment.Id
	c.UserId = comment.UserId
	c.PhotoId = comment.PhotoId
	c.Content = comment.Content
}

func (c *Comment) CommentConvertForDatabase() database.Comment {
	return database.Comment{
		Id:         c.Id,
		UserId:     c.UserId,
		PhotoId:    c.PhotoId,
		PhotoOwner: c.PhotoOwner,
		Content:    c.Content,
	}
}
