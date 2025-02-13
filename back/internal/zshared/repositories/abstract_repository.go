package shared_repositories

import (
	"context"
	"errors"
	"reflect"

	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AbstractRepository[T any, Q any] struct {
	collection *mongo.Collection
	logger     *zerolog.Logger
}

func BuildNewRepository[T any, Q any](ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) shared_ports.Repository[T, Q] {
	coll := connection.Collection(collection)

	return &AbstractRepository[T, Q]{
		collection: coll,
		logger:     logger,
	}
}

func (r *AbstractRepository[T, Q]) FindOne(ctx context.Context, opts shared_ports.FindOneOpts, result *T) error {
	filterBson := bson.D{}
	for key, value := range opts.Filter {
		filterBson = append(filterBson, bson.E{Key: key, Value: validateOID(value)})
	}

	findOtps := options.FindOne()

	res := r.collection.FindOne(ctx, filterBson, findOtps)
	nilFind, _ := res.Raw()
	if nilFind == nil {
		return errors.New("not found")
	}

	err := res.Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (r *AbstractRepository[T, Q]) FindMany(ctx context.Context, opts shared_ports.FindManyOpts, result *[]T, returnCount bool) (*int64, error) {
	filterBson := bson.D{}
	for key, value := range opts.Filter {
		filterBson = append(filterBson, bson.E{Key: key, Value: validateOID(value)})
	}

	findOtps := options.Find()
	if opts.Take > 0 {
		findOtps.SetLimit(opts.Take)
	}
	if opts.Skip > 0 {
		findOtps.SetSkip(opts.Skip)
	}
	if len(opts.Sort) > 0 {
		findOtps.SetSort(opts.Sort)
	}

	cur, err := r.collection.Find(ctx, filterBson, findOtps)
	if err != nil {
		r.logger.Err(err).Msg("[FindMany Repo]: Failed to fetch data in")
		return nil, err
	}

	if err = cur.All(ctx, result); err != nil {
		r.logger.Err(err).Msg("[FindMany Repo]: Failed to decode data")
		return nil, err
	}

	if returnCount {
		count, err := r.collection.CountDocuments(ctx, filterBson)
		if err != nil {
			r.logger.Err(err).Msg("[FindMany Repo]: Failed to fetch data in Count")
			return nil, err
		}
		return &count, nil
	}

	return nil, nil
}

func (r *AbstractRepository[T, Q]) InsertOne(ctx context.Context, entity T) error {
	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *AbstractRepository[T, Q]) DeleteOne(ctx context.Context, opts shared_ports.DeleteOpts) (bool, error) {
	var filterBson bson.D
	for key, value := range opts.Filter {
		filterBson = append(filterBson, bson.E{Key: key, Value: validateOID(value)})
	}

	res, err := r.collection.DeleteOne(ctx, filterBson)
	if err != nil {
		return false, err
	}

	if res.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

func validateOID(val interface{}) interface{} {
	if reflect.ValueOf(val).Kind() == reflect.String {
		oid, err := primitive.ObjectIDFromHex(val.(string))
		if err != nil {
			return val
		}

		return oid
	}

	return val
}

func (r *AbstractRepository[T, Q]) UpdateOne(ctx context.Context, opts shared_ports.UpdateOpts) (*T, error) {
	filterBson := bson.D{}
	for key, value := range opts.Filter {
		filterBson = append(filterBson, bson.E{Key: key, Value: validateOID(value)})
	}

	payloadBson := bson.D{}
	if opts.Payload != nil {
		for key, value := range *opts.Payload {
			payloadBson = append(payloadBson, bson.E{Key: key, Value: value})
		}
	}

	after := options.After
	updtOpts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{{Key: "$set", Value: payloadBson}}

	res := r.collection.FindOneAndUpdate(ctx, filterBson, update, updtOpts)
	err := res.Err()
	if err != nil {
		r.logger.Err(err).Interface("filter", filterBson).Interface("payload", payloadBson).Msg("ERROR | Update One Repo")
		return nil, err
	}

	var newData T
	if err := res.Decode(&newData); err != nil {
		r.logger.Err(err).Interface("filter", filterBson).Interface("payload", payloadBson).Msg("ERROR | Update One Repo <Decode Response>")
		return nil, err
	}

	return &newData, nil
}
