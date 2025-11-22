package tracks

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Circuit, error)
	GetByID(ctx context.Context, id string) (*Circuit, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, dbName string) Repository {
	return &repository{
		collection: client.Database(dbName).Collection("circuits"),
	}
}

func (r *repository) GetAll(ctx context.Context) ([]Circuit, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var circuits []Circuit
	if err := cursor.All(ctx, &circuits); err != nil {
		return nil, err
	}
	return circuits, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*Circuit, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var circuit Circuit
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&circuit)
	if err != nil {
		return nil, err
	}

	return &circuit, nil
}
