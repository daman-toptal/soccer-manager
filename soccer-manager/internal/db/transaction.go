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

type TransactionDbManager interface {
	Create(context.Context, *model.Transaction) (*model.Transaction, error)
	Get(context.Context, id.TransactionID) (*model.Transaction, error)
	Find(context.Context, map[string]interface{}) ([]*model.Transaction, error)
	Update(context.Context, *model.Transaction, ...map[string]interface{}) (*model.Transaction, error)
}

type transaction struct {
	collection *mongo.Collection
}

func NewTransactionDbManager(collection *mongo.Collection) TransactionDbManager {
	return transaction{
		collection: collection,
	}
}

func (t transaction) Create(ctx context.Context, tm *model.Transaction) (*model.Transaction, error) {
	tm.CreatedAt = time.Now()
	_, err := t.collection.InsertOne(ctx, tm)
	return tm, err
}

func (t transaction) Get(ctx context.Context, txnID id.TransactionID) (*model.Transaction, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: txnID,
	}}
	transaction := &model.Transaction{}
	if err := t.collection.FindOne(ctx, filter).Decode(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t transaction) Update(ctx context.Context, updateModel *model.Transaction, filters ...map[string]interface{}) (*model.Transaction, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: updateModel.Id,
	}}
	update := bson.M{"$set": t.getUpdateMap(updateModel)}
	transaction := &model.Transaction{}
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: (*options.ReturnDocument)(&After)}
	if err := t.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t transaction) Find(ctx context.Context, filters map[string]interface{}) ([]*model.Transaction, error) {
	dbFilters := bson.M{}
	for key, val := range filters {
		dbFilters[key] = val
	}
	cur, err := t.collection.Find(ctx, dbFilters)
	if err != nil {
		return nil, err
	}
	var transactions []*model.Transaction
	for cur.Next(ctx) {
		transaction := &model.Transaction{}
		if err := cur.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)
	return transactions, nil
}

func (t transaction) getUpdateMap(updateModel *model.Transaction) bson.M {
	updateMap := bson.M{}
	return updateMap
}
