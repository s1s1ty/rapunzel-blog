package storage

import (
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"github.com/s4kibs4mi/rapunzel-blog/proto/defs"
	"gopkg.in/mgo.v2/bson"
)

/**
 * := Coded with love by Sakib Sami on 20/1/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type UserStorage interface {
	Init() bool
	Save(user models.User) bool
	Update(user *models.User) bool
	Delete(user models.User) bool
	Count() int
	FindByID(ID bson.ObjectId) *models.User
	FindByUsername(username string) *models.User
	FindByEmail(username string) *models.User
	FindAll() []models.User
	FindAllByQuery(query defs.Query) []models.User
}

type SessionStorage interface {
	Init() bool
	SaveSession(session *models.Session) bool
	DeleteSession(session *models.Session) bool
	FindSessionByAccessToken(ID string) *models.Session
}

type PostStorage interface {
	Init() bool
	SavePost(post *models.Post) bool
	UpdatePost(post *models.Post) bool
	DeletePost(post *models.Post) bool
	FindPostsByQuery(query []*defs.Query) []*models.Post
	FindPostByID(postID string) *models.Post
}

type CommentStorage interface {
	Init() bool
	SaveComment(comment *models.Comment) bool
	FindCommentsByQuery(query []*defs.Query) []*models.Comment
	FindCommentByID(commentID string) *models.Comment
}

var mongodbStorage *MongodbStorage

func NewUserStorage() UserStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}

func NewSessionStorage() SessionStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}

func NewPostStorage() PostStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}

func NewCommentStorage() CommentStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}
