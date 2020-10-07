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

// online database section
const URIOnline = "mongodb+srv://" + UserOnline + ":" + PassOnline + "@duty.1cotf.gcp.mongodb.net/" + DbNameOnline + "?retryWrites=true&w=majority"
const UserOnline = "adminDB"
const PassOnline = "3nHF5VCqi6w3Cgc3"
const DbNameOnline = "duty"

type DBConnection interface {
	ConnectDB(ctx context.Context, path string) (*mongo.Client, error)
	DisconnectDB(ctx context.Context, db *mongo.Client) error
	GetDB(db *mongo.Client, name string) *mongo.Database
}


type connDB struct {
}

// constructor
func NewConnectDB() DBConnection {
	return &connDB{}
}

// methods to implement interface
func (conn *connDB) ConnectDB(ctx context.Context, path string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(path)
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

func (conn *connDB) GetDB(cn *mongo.Client, name string) *mongo.Database {
	return cn.Database(name)
}