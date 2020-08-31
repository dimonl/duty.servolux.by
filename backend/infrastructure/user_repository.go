package infrastructure

import (
	"context"
	"fmt"
	"log"
	"main/domain"
	"main/infrastructure/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	//TODO make func to realize methods of this repo
	CreateUser(ctx context.Context, newUser *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, dutyId string, newDutyName string) error
	DeleteUser(ctx context.Context, dutyId string) error
	GetUserByID(ctx context.Context, dutyId string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)
}

type UserRep struct {
}

func NewUserRepository() UserRepository {
	return &UserRep{}
}

//func (cp *companyRep) CreateCompany(ctx context.Context, companyName string) (*domain.Company, error) {
func (cp *UserRep) CreateUser(ctx context.Context, newUser *domain.Users) (*domain.User, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBUsersTable)

	//comp := domain.NewCompany()
	//comp.ID = primitive.NewObjectID()
	//comp.Name = companyName
	res, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return newUser, nil
}

func (cp *UserRep) UpdateUser(ctx context.Context, dutyId string, newDutyName string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBUsersTable)

	//id, _ := primitive.ObjectIDFromHex(companyId)
	res, err := cp.GetUserByID(ctx, dutyId)
	if err != nil {
		return err
	}
	if res.Name != newDutyName {
		filter := bson.M{"_id": res.ID}
		update := bson.M{"$set": bson.M{"FirstName": newDutyName}}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			//fmt.Printf("update fail %v\n", err)
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (cp *UserRep) DeleteUser(ctx context.Context, UserId string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBUsersTable)
	id, _ := primitive.ObjectIDFromHex(UserId)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(deleteResult)
	return nil
}

func (cp *UserRep) GetUserByID(ctx context.Context, ids string) (*domain.User, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBUsersTable)
	id, _ := primitive.ObjectIDFromHex(ids)
	filter := bson.D{{"_id", id}}
	var res domain.User
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}

func (cp *UserRep) GetUsers(ctx context.Context) ([]*domain.Users, error) {

	var specArray []*domain.User

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBUsersTable)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(ctx) {
		//
		var elem domain.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		specArray = append(specArray, &elem)
	}
	cur.Close(ctx)
	return specArray, nil
}
