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

type DutiesRepository interface {
	//TODO make func to realize methods of this repo
	CreateDuty(ctx context.Context, newDuty *domain.Duties) (*domain.Duties, error)
	UpdateDuty(ctx context.Context, dutyId string, newDuty string) error
	DeleteDuty(ctx context.Context, dutyId string) error
	GetDutyByID(ctx context.Context, dutyId string) (*domain.Duties, error)
	GetDuties(ctx context.Context) ([]*domain.Duties, error)
}

type dutiesRep struct {
}

func NewDutiesRepository() DutiesRepository {
	return &dutiesRep{}
}

//func (cp *companyRep) CreateCompany(ctx context.Context, companyName string) (*domain.Company, error) {
func (cp *dutiesRep) CreateDuty(ctx context.Context, newDuty *domain.Duties) (*domain.Duties, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyTable)

	//comp := domain.NewCompany()
	//comp.ID = primitive.NewObjectID()
	//comp.Name = companyName
	res, err := collection.InsertOne(ctx, newDuty)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return newDuty, nil
}

func (cp *dutiesRep) UpdateDuty(ctx context.Context, dutyId string, newDuty string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyTable)

	//id, _ := primitive.ObjectIDFromHex(companyId)
	res, err := cp.GetDutyByID(ctx, dutyId)
	if err != nil {
		return err
	}
	if res.Name != newDuty {
		filter := bson.M{"_id": res.ID}
		update := bson.M{"$set": bson.M{"Name": newDuty}}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			//fmt.Printf("update fail %v\n", err)
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (cp *dutiesRep) DeleteDuty(ctx context.Context, dutyId string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyTable)
	id, _ := primitive.ObjectIDFromHex(companyId)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(deleteResult)
	return nil
}

func (cp *dutiesRep) GetDutyByID(ctx context.Context, ids string) (*domain.Duties, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyTable)
	id, _ := primitive.ObjectIDFromHex(ids)
	filter := bson.D{{"_id", id}}
	var res domain.Duties
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}

func (cp *dutiesRep) GetDuties(ctx context.Context) ([]*domain.Duties, error) {

	var specArray []*domain.Duties

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyTable)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(ctx) {
		//
		var elem domain.Duties
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		specArray = append(specArray, &elem)
	}
	cur.Close(ctx)
	return specArray, nil
}
