package tracks_test

import (
	"context"
	"testing"

	"FT26/internal/tracks"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestRepository_GetAll(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		repo := tracks.NewRepository(mt.Client, "testdb")

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			"testdb.tracks",
			mtest.FirstBatch,
			bson.D{{Key: "name", Value: "Monza"}},
			bson.D{{Key: "name", Value: "Imola"}},
		))

		res, err := repo.GetAll(context.Background())

		if err != nil {
			t.Fatal(err)
		}

		if len(res) != 2 {
			t.Fatalf("expected 2 circuits, got %d", len(res))
		}
	})
}

func TestRepository_GetByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("found", func(mt *mtest.T) {
		repo := tracks.NewRepository(mt.Client, "testdb")

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			"testdb.tracks",
			mtest.FirstBatch,
			bson.D{{Key: "name", Value: "Silverstone"}},
		))

		res, err := repo.GetByID(context.Background(), "507f1f77bcf86cd799439011")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if res.Name != "Silverstone" {
			t.Fatalf("expected Silverstone, got %s", res.Name)
		}
	})
}
