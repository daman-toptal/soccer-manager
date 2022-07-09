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

type TeamDbManager interface {
	Create(context.Context, *model.Team) (*model.Team, error)
	Get(context.Context, id.TeamID) (*model.Team, error)
	Find(context.Context, map[string]interface{}) ([]*model.Team, error)
	Update(context.Context, *model.Team, ...map[string]interface{}) (*model.Team, error)
}

type team struct {
	collection *mongo.Collection
}

func NewTeamDbManager(collection *mongo.Collection) TeamDbManager {
	return team{
		collection: collection,
	}
}

func (t team) Create(ctx context.Context, tm *model.Team) (*model.Team, error) {
	tm.CreatedAt = time.Now()

	_, err := t.collection.InsertOne(ctx, tm)
	return tm, err
}

func (t team) Get(ctx context.Context, txnID id.TeamID) (*model.Team, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: txnID,
	}}
	team := &model.Team{}
	if err := t.collection.FindOne(ctx, filter).Decode(team); err != nil {
		return nil, err
	}
	return team, nil
}

func (t team) Update(ctx context.Context, updateModel *model.Team, filters ...map[string]interface{}) (*model.Team, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: updateModel.Id,
	}}
	update := bson.M{"$set": t.getUpdateMap(updateModel)}
	team := &model.Team{}
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: (*options.ReturnDocument)(&After)}
	if err := t.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(team); err != nil {
		return nil, err
	}
	return team, nil
}

func (t team) Find(ctx context.Context, filters map[string]interface{}) ([]*model.Team, error) {
	dbFilters := bson.M{}
	for key, val := range filters {
		dbFilters[key] = val
	}
	cur, err := t.collection.Find(ctx, dbFilters)
	if err != nil {
		return nil, err
	}
	var teams []*model.Team
	for cur.Next(ctx) {
		team := &model.Team{}
		if err := cur.Decode(&team); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)
	return teams, nil
}

func (t team) getUpdateMap(updateModel *model.Team) bson.M {
	updateMap := bson.M{}
	if !(updateModel.Name == "") {
		updateMap["name"] = updateModel.Name
	}
	if !(updateModel.Country == "") {
		updateMap["country"] = updateModel.Country
	}
	if !(updateModel.Value == nil) {
		updateMap["value"] = *updateModel.Value
	}
	if !(updateModel.Budget == nil) {
		updateMap["budget"] = *updateModel.Budget
	}
	return updateMap
}
