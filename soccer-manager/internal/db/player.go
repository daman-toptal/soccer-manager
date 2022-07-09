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

type PlayerDbManager interface {
	Create(context.Context, *model.Player) (*model.Player, error)
	Get(context.Context, id.PlayerID) (*model.Player, error)
	Find(context.Context, map[string]interface{}) ([]*model.Player, error)
	Update(context.Context, *model.Player, ...map[string]interface{}) (*model.Player, error)
}

type player struct {
	collection *mongo.Collection
}

func NewPlayerDbManager(collection *mongo.Collection) PlayerDbManager {
	return player{
		collection: collection,
	}
}

func (p player) Create(ctx context.Context, pm *model.Player) (*model.Player, error) {
	pm.CreatedAt = time.Now()

	_, err := p.collection.InsertOne(ctx, pm)
	return pm, err
}

func (p player) Get(ctx context.Context, txnID id.PlayerID) (*model.Player, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: txnID,
	}}
	player := &model.Player{}
	if err := p.collection.FindOne(ctx, filter).Decode(player); err != nil {
		return nil, err
	}
	return player, nil
}

func (p player) Update(ctx context.Context, updateModel *model.Player, filters ...map[string]interface{}) (*model.Player, error) {
	var filter interface{}

	if len(filters) > 0 {
		filterMap := bson.M{}
		for key, val := range filters[0] {
			filterMap[key] = val
		}
		filter = filterMap
	} else {
		filter = bson.D{{
			Key:   "_id",
			Value: updateModel.Id,
		}}
	}

	update := bson.M{"$set": p.getUpdateMap(updateModel)}
	player := &model.Player{}
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: (*options.ReturnDocument)(&After)}
	if err := p.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(player); err != nil {
		return nil, err
	}
	return player, nil
}

func (p player) Find(ctx context.Context, filters map[string]interface{}) ([]*model.Player, error) {
	dbFilters := bson.M{}
	for key, val := range filters {
		dbFilters[key] = val
	}
	cur, err := p.collection.Find(ctx, dbFilters)
	if err != nil {
		return nil, err
	}
	var players []*model.Player
	for cur.Next(ctx) {
		player := &model.Player{}
		if err := cur.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)
	return players, nil
}

func (p player) getUpdateMap(updateModel *model.Player) bson.M {
	updateMap := bson.M{}
	if !(updateModel.FirstName == "") {
		updateMap["firstName"] = updateModel.FirstName
	}
	if !(updateModel.LastName == "") {
		updateMap["lastName"] = updateModel.LastName
	}
	if !(updateModel.Country == "") {
		updateMap["country"] = updateModel.Country
	}
	if !(updateModel.Value == nil) {
		updateMap["value"] = *updateModel.Value
	}
	if !(updateModel.IsListed == nil) {
		updateMap["isListed"] = *updateModel.IsListed
	}
	if !(updateModel.AskValue == nil) {
		updateMap["askValue"] = *updateModel.AskValue
	}
	if !(updateModel.TeamId.IsZero()) {
		updateMap["teamId"] = updateModel.TeamId
	}
	return updateMap
}
