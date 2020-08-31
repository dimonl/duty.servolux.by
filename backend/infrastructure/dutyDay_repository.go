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

type DutyDayRepository interface {
	//TODO make func to realize methods of this repo
	CreateDutyDay(ctx context.Context, newDutyDay *domain.DutyDay) (*domain.DutyDay, error)
	UpdateDutyDay(ctx context.Context, dutyId string, newDutyName string) error
	DeleteDutyDay(ctx context.Context, dutyId string) error
	GetDutyDayByID(ctx context.Context, dutyId string) (*domain.DutyDay, error)
	GetDutyDays(ctx context.Context) ([]*domain.DutyDay, error)
}

type dutyDayRep struct {
}

func NewDutyDayRepository() DutyDayRepository {
	return &dutyDayRep{}
}

//func (cp *companyRep) CreateCompany(ctx context.Context, companyName string) (*domain.Company, error) {
func (cp *dutyDayRep) CreateDutyDay(ctx context.Context, newDutyDay *domain.DutyDay) (*domain.DutyDay, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyDayTable)

	//comp := domain.NewCompany()
	//comp.ID = primitive.NewObjectID()
	//comp.Name = companyName
	res, err := collection.InsertOne(ctx, newDutyDay)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return newDutyDay, nil
}

func (cp *dutyDayRep) UpdateDutyDay(ctx context.Context, dutyId string, newDutyName string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyDayTable)

	//id, _ := primitive.ObjectIDFromHex(companyId)
	res, err := cp.GetDutyDayByID(ctx, dutyId)
	if err != nil {
		return err
	}
	if res.Name != newDutyName {
		filter := bson.M{"_id": res.ID}
		update := bson.M{"$set": bson.M{"Day": newDutyName}}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			//fmt.Printf("update fail %v\n", err)
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (cp *dutyDayRep) DeleteDutyDay(ctx context.Context, dutyDayId string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyDayTable)
	id, _ := primitive.ObjectIDFromHex(dutyDayId)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(deleteResult)
	return nil
}

func (cp *dutyDayRep) GetDutyDayByID(ctx context.Context, ids string) (*domain.DutyDay, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyDayTable)
	id, _ := primitive.ObjectIDFromHex(ids)
	filter := bson.D{{"_id", id}}
	var res domain.DutyDay
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}

func (cp *dutyDayRep) GetDutyDays(ctx context.Context) ([]*domain.dutyDay, error) {

	var specArray []*domain.dutyDay

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DBname).Collection(database.DBDutyDayTable)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(ctx) {
		//
		var elem domain.dutyDay
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		specArray = append(specArray, &elem)
	}
	cur.Close(ctx)
	return specArray, nil
}
