package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const UsersCollectionName = "users"
const DBDutyCategoriesTable = "dutyCategories"
const DBDutyWorkersTable = "dutyWorkers"
const DBDutyDayTable = "dutyDay"
const DBDutyTable = "duties"
const DBname = "dutyDB"
const URI = "mongodb://127.0.0.1:27017"

type DBConnection interface {
	ConnectDB(ctx context.Context) (*mongo.Client, error)
	DisconnectDB(ctx context.Context, db *mongo.Client) error
	GetDB(db *mongo.Client) *mongo.Database
}

type connDB struct {
}

// constructor
func NewConnectDB() DBConnection {
	return &connDB{}
}

// methods to implement interface
func (conn *connDB) ConnectDB(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func (conn *connDB) DisconnectDB(ctx context.Context, cl *mongo.Client) error {
	err := cl.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
	return err
}

func (conn *connDB) GetDB(cn *mongo.Client) *mongo.Database {
	return cn.Database(DBname)
}
