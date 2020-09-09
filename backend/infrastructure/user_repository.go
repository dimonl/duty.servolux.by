package infrastructure

import (
	"context"
	"crypto/md5"
	"dsv/domain"
	"dsv/infrastructure/database"
	"errors"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	//TODO make func to realize methods of this repo
	CreateUser(ctx context.Context, login string, pass string) (*domain.User, error)
	UpdateUser(ctx context.Context, userId string, newDutyName string) error
	DeleteUser(ctx context.Context, userId string) error
	GetUserByID(ctx context.Context, userId string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)

	UserLogin(ctx context.Context, login string, pass string) (*domain.User, error)
	GetUserByLogin(ctx context.Context, login string) (*domain.User, error)
}

type UserRep struct {
}

func NewUserRepository() UserRepository {
	return &UserRep{}
}

// func (cp *companyRep) CreateCompany(ctx context.Context, companyName string) (*domain.Company, error) {

// CreateUser ..
func (cp *UserRep) CreateUser(ctx context.Context, login string, pass string) (*domain.User, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	h := md5.New()
	collection := client.Database(database.DBname).Collection(database.UsersCollectionName)

	userForInsert := domain.NewUser()
	userForInsert.ID = primitive.NewObjectID()
	userForInsert.FirstName = login
	userForInsert.Password = string(h.Sum([]byte(pass)))

	res, err := collection.InsertOne(ctx, userForInsert)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return userForInsert, nil
}

func (cp *UserRep) UpdateUser(ctx context.Context, userId string, newUserName string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.UsersCollectionName)

	//id, _ := primitive.ObjectIDFromHex(companyId)
	res, err := cp.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}
	if res.FirstName != newUserName {
		filter := bson.M{"_id": res.ID}
		update := bson.M{"$set": bson.M{"FirstName": newUserName}}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			//fmt.Printf("update fail %v\n", err)
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (cp *UserRep) DeleteUser(ctx context.Context, userId string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.UsersCollectionName)
	id, _ := primitive.ObjectIDFromHex(userId)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(deleteResult)
	return nil
}

func (cp *UserRep) GetUserByID(ctx context.Context, userId string) (*domain.User, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.UsersCollectionName)
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"_id", id}}
	var res domain.User
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}

func (cp *UserRep) GetUsers(ctx context.Context) ([]*domain.User, error) {

	var specArray []*domain.User

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.UsersCollectionName)
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

// UserLogin ...
func (cp *UserRep) UserLogin(ctx context.Context, login string, pass string) (*domain.User, error) {

	if strings.TrimSpace(login) == "" {
		return nil, errors.New("empty login")
	}

	us, err := cp.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	if us != nil {
		h := md5.New()

		loginPass := string(h.Sum([]byte(pass)))
		if loginPass == us.Password {
			return us, nil
		}

		return nil, errors.New("password is not valid")
	}
	return nil, err
}

func (cp *UserRep) GetUserByLogin(ctx context.Context, login string) (*domain.User, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	var result domain.User

	collection := client.Database(database.DBname).Collection(database.UsersCollectionName)
	filter := bson.D{{"name", login}}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}