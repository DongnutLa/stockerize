package shared_repositories

import (
	"context"
	"fmt"
	"sync"

	utils "github.com/DongnutLa/stockio/internal/zshared/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionConfig struct {
	URI             string
	DatabaseName    string
	OtherClientOpts *options.ClientOptions
}

var once sync.Once
var databaseConn *mongo.Database

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(ctx context.Context) *MongoDB {
	var client *mongo.Client
	config := ConnectionConfig{
		URI:          utils.GetConfig("db_uri"),
		DatabaseName: utils.GetConfig("db_name"),
	}

	once.Do(func() {
		// Connect to MongoDB
		clientOptions := options.Client().ApplyURI(config.URI)
		clientDb, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			client = nil
			return
		}

		client = clientDb
	})

	if client == nil {
		fmt.Println("[MONGO DB] Failed to initialize MongoDB")
		return nil
	}

	database := client.Database(config.DatabaseName)
	databaseConn = database

	fmt.Println("[MONGO DB] MongoDB initialized successfully")
	return &MongoDB{
		Client:   client,
		Database: database,
	}
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
	if m.Client != nil {
		return m.Client.Disconnect(ctx)
	}
	return nil
}

func GetMongoConnection() *mongo.Database {
	return databaseConn
}
