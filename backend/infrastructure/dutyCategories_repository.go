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

type DutyCategoriesRepository interface {
	//TODO make func to realize methods of this repo
	CreateDutyCategory(ctx context.Context, newDutyCategory *domain.DutyCategories) (*domain.DutyCategories, error)
	UpdateDutyCategory(ctx context.Context, dutyId string, newDutyName string) error
	DeleteDutyCategory(ctx context.Context, dutyId string) error
	GetDutyCategoryByID(ctx context.Context, dutyId string) (*domain.DutyCategories, error)
	GetDutyCategories(ctx context.Context) ([]*domain.DutyCategories, error)
}

type dutyCategoriesRep struct {
}

func NewDutyCategoriesRepository() DutyCategoriesRepository {
	return &dutyCategoriesRep{}
}

//func (cp *companyRep) CreateCompany(ctx context.Context, companyName string) (*domain.Company, error) {
func (cp *dutyCategoriesRep) CreateDutyCategory(ctx context.Context, newDutyCategory *domain.DutyCategories) (*domain.DutyCategories, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyCategoriesTable)

	//comp := domain.NewCompany()
	//comp.ID = primitive.NewObjectID()
	//comp.Name = companyName
	res, err := collection.InsertOne(ctx, newDutyCategory)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return newDutyCategory, nil
}

func (cp *dutyCategoriesRep) UpdateDutyCategory(ctx context.Context, dutyId string, newDutyName string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyCategoriesTable)

	//id, _ := primitive.ObjectIDFromHex(companyId)
	res, err := cp.GetDutyCategoryByID(ctx, dutyId)
	if err != nil {
		return err
	}
	if res.Name != newDutyName {
		filter := bson.M{"_id": res.ID}
		update := bson.M{"$set": bson.M{"Category": newDutyName}}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			//fmt.Printf("update fail %v\n", err)
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (cp *dutyCategoriesRep) DeleteDutyCategory(ctx context.Context, dutyCategoryId string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyCategoriesTable)
	id, _ := primitive.ObjectIDFromHex(dutyCategoryId)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(deleteResult)
	return nil
}

func (cp *dutyCategoriesRep) GetDutyCategoryByID(ctx context.Context, ids string) (*domain.DutyCategories, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyCategoriesTable)
	id, _ := primitive.ObjectIDFromHex(ids)
	filter := bson.D{{"_id", id}}
	var res domain.DutyCategories
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}

func (cp *dutyCategoriesRep) GetDuties(ctx context.Context) ([]*domain.DutyCategories, error) {

	var specArray []*domain.DutyCategories

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyCategoriesTable)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(ctx) {
		//
		var elem domain.DutyCategories
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		specArray = append(specArray, &elem)
	}
	cur.Close(ctx)
	return specArray, nil
}
