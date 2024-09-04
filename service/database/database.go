package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Errori comuni
var (
	ErrUserNotFound    = errors.New("User does not exist")
	ErrPhotoNotFound   = errors.New("Photo does not exist")
	ErrBanNotFound     = errors.New("Ban does not exist")
	ErrFollowNotFound  = errors.New("Follow does not exist")
	ErrCommentNotFound = errors.New("Comment does not exist")
	ErrLikeNotFound    = errors.New("Like does not exist")
)

// Tipi di Dati (Struct)

// User rappresenta un utente
type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

// PhotoStream rappresenta uno stream di foto
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

// Stream rappresenta uno stream di foto
type Stream struct {
	Identifier uint64        `json:"identifier"`
	Photos     []PhotoStream `json:"photoStream"`
}

// Followers rappresenta i follower di un utente
type Followers struct {
	Id        uint64   `json:"identifier"`
	Followers []Follow `json:"Followers"`
}

// Follow rappresenta una relazione di follow
type Follow struct {
	FollowId   uint64 `json:"followId"`
	FollowedId uint64 `json:"followedId"`
	UserId     uint64 `json:"userId"`
}

// Bans rappresenta i bans di un utente
type Bans struct {
	Identifier uint64 `json:"identifier"`
	Username   string `json:"username"`
	Bans       []Ban  `json:"bans"`
}

// Ban rappresenta un ban
type Ban struct {
	BanId    uint64 `json:"banId"`
	BannedId uint64 `json:"bannedId"`
	UserId   uint64 `json:"userId"`
}

// Photos rappresenta una collezione di foto
type Photos struct {
	RequestUser uint64  `json:"requestUser"`
	Identifier  uint64  `json:"identifier"`
	Photos      []Photo `json:"photos"`
}

// Photo rappresenta una foto
type Photo struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	File          []byte `json:"file"`
	Date          string `json:"date"`
	LikesCount    int    `json:"likesCount"`
	CommentsCount int    `json:"commentsCount"`
	LikeStatus    bool   `json:"likeStatus"`
}

// Like rappresenta un like
type Like struct {
	LikeId          uint64 `json:"likeId"`
	UserIdentifier  uint64 `json:"identifier"`
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	PhotoOwner      uint64 `json:"photoOwner"`
}

// Comments rappresenta i commenti su una foto
type Comments struct {
	RequestIdentifier uint64    `json:"requestIdentifier"`
	PhotoIdentifier   uint64    `json:"photoIdentifier"`
	PhotoOwner        uint64    `json:"identifier"`
	Comments          []Comment `json:"comments"`
}

// Comment rappresenta un commento
type Comment struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	PhotoId       uint64 `json:"photoId"`
	PhotoOwner    uint64 `json:"photoOwner"`
	OwnerUsername string `json:"ownerUsername"`
	Username      string `json:"username"`
	Content       string `json:"content"`
}

// AppDatabase è l'interfaccia ad alto livello per il database
type AppDatabase interface {
	CreateUser(User) (User, error)
	SetUsername(User, string) (User, error)
	GetUserId(string) (User, error)
	CheckUserById(User) (User, error)
	CheckUserByUsername(User) (User, error)
	CheckUser(User) (User, error)
	SearchUsersByUsername(username string) ([]User, error)
	GetMyStream(User) ([]PhotoStream, error)

	SetFollow(Follow) (Follow, error)
	RemoveFollow(uint64, uint64) error
	GetFollowingId(user1 uint64, user2 uint64) (Follow, error)
	GetFollowers(User, uint64) (Follow, error)
	GetFollowersCount(uint64) (int, error)
	GetFollowingsCount(uint64) (int, error)
	GetFollowStatus(uint64, uint64) (bool, error)

	CreateBan(Ban) (Ban, error)
	RemoveBan(Ban) error
	GetBan(User, uint64) (Ban, error)
	GetBanById(Ban) (Ban, error)
	UpdateBanStatus(int, uint64, uint64) error
	GetBanStatus(uint64, uint64) (bool, error)
	CheckIfBanned(uint64, uint64) (bool, error)

	SetPhoto(Photo) (Photo, error)
	RemovePhoto(uint64) error
	GetPhotos(User, uint64) ([]Photo, error)
	CheckPhoto(Photo) (Photo, error)

	SetLike(Like) (Like, error)
	RemoveLike(Like) error
	RemoveLikes(uint64, uint64) error
	GetLike(uint64, uint64) (Like, error)
	GetLikeById(Like) (Like, error)
	GetLikesCount(photoid uint64) (int, error)

	SetComment(Comment) (Comment, error)
	RemoveComment(Comment) error
	RemoveComments(uint64, uint64) error
	GetComments(photoid uint64) ([]Comment, error)
	GetCommentById(Comment) (Comment, error)
	GetCommentsCount(uint64) (int, error)

	Ping() error
}

// appdbimpl è l'implementazione dell'interfaccia AppDatabase
type appdbimpl struct {
	c *sql.DB
}

// New restituisce una nuova istanza di AppDatabase basata sulla connessione SQLite `db`.
// `db` è obbligatorio - verrà restituito un errore se `db` è `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Abilita le chiavi esterne
	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	// Verifica se la tabella esiste. In caso contrario, crea la struttura del database
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		usersDatabase := `CREATE TABLE users (
			Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE
			);`
		photosDatabase := `CREATE TABLE photos (
			Id INTEGER NOT NULL PRIMARY KEY, 
			userId INTEGER NOT NULL,
			photo BLOB,
			date TEXT,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		likesDatabase := `CREATE TABLE likes (
			Id INTEGER NOT NULL PRIMARY KEY,
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			photoOwner INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		commentsDatabase := `CREATE TABLE comments (
			Id INTEGER NOT NULL PRIMARY KEY,
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			photoOwner INTEGER NOT NULL,
			content TEXT NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		bansDatabase := `CREATE TABLE bans (
			banId INTEGER NOT NULL PRIMARY KEY,
			bannedId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		followersDatabase := `CREATE TABLE followers (
			Id INTEGER NOT NULL PRIMARY KEY,
			followerId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`

		// Crea le tabelle nel database
		_, err = db.Exec(usersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(photosDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(likesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(commentsDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(bansDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(followersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// Ping controlla la connessione al database.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
