package db

import (
	"context"
	"soccer-manager/internal/model"
	"soccer-manager/util/id"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	After int8 = 1
)

type UserDbManager interface {
	Create(context.Context, *model.User) (*model.User, error)
	Get(context.Context, id.UserID) (*model.User, error)
	Find(context.Context, map[string]interface{}) ([]*model.User, error)
	Update(context.Context, *model.User, ...map[string]interface{}) (*model.User, error)
}

type user struct {
	collection *mongo.Collection
}

func NewUserDbManager(collection *mongo.Collection) UserDbManager {
	return user{
		collection: collection,
	}
}

func (u user) Create(ctx context.Context, um *model.User) (*model.User, error) {
	um.CreatedAt = time.Now()
	_, err := u.collection.InsertOne(ctx, um)
	return um, err
}

func (u user) Get(ctx context.Context, txnID id.UserID) (*model.User, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: txnID,
	}}
	user := &model.User{}
	if err := u.collection.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u user) Update(ctx context.Context, updateModel *model.User, filters ...map[string]interface{}) (*model.User, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: updateModel.Id,
	}}
	update := bson.M{"$set": u.getUpdateMap(updateModel)}
	user := &model.User{}
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: (*options.ReturnDocument)(&After)}
	if err := u.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u user) Find(ctx context.Context, filters map[string]interface{}) ([]*model.User, error) {
	dbFilters := bson.M{}
	for key, val := range filters {
		dbFilters[key] = val
	}
	cur, err := u.collection.Find(ctx, dbFilters)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for cur.Next(ctx) {
		user := &model.User{}
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)
	return users, nil
}

func (u user) getUpdateMap(updateModel *model.User) bson.M {
	updateMap := bson.M{}
	if !(updateModel.Name == "") {
		updateMap["name"] = updateModel.Name
	}
	return updateMap
}
