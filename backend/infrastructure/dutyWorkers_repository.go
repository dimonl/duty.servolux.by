package infrastructure

import (
	"context"
	"dsv/domain"
	"dsv/infrastructure/database"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DutyWorkersRepository ..
type DutyWorkersRepository interface {
	//TODO make func to realize methods of this repo
	CreateDutyWorker(ctx context.Context, newDutyWorker *domain.DutyWorkers) (*domain.DutyWorkers, error)
	UpdateDutyWorker(ctx context.Context, dutyId, newDutyName, newDutySurname, newIDCategories string) error
	DeleteDutyWorker(ctx context.Context, dutyId string) error
	GetDutyWorkerByID(ctx context.Context, dutyId string) (*domain.DutyWorkers, error)
	GetDutyWorkers(ctx context.Context) ([]*domain.DutyWorkers, error)
}

type dutyWorkerRep struct {
}

func NewDutyWorkersRepository() DutyWorkersRepository {
	return &dutyWorkerRep{}
}

//func (cp *companyRep) CreateCompany(ctx context.Context, companyName string) (*domain.Company, error) {
func (cp *dutyWorkerRep) CreateDutyWorker(ctx context.Context, newDutyWorker *domain.DutyWorkers) (*domain.DutyWorkers, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx, database.URIOnline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DbNameOnline).Collection(database.DBDutyWorkersTable)

	//comp := domain.NewCompany()
	//comp.ID = primitive.NewObjectID()
	//comp.Name = companyName
	res, err := collection.InsertOne(ctx, newDutyWorker)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return newDutyWorker, nil
}

func (cp *dutyWorkerRep) UpdateDutyWorker(ctx context.Context, dutyId, newDutyName, newDutySurname, newIDCategories string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx, database.URIOnline)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DbNameOnline).Collection(database.DBDutyWorkersTable)

	//id, _ := primitive.ObjectIDFromHex(companyId)
	res, err := cp.GetDutyWorkerByID(ctx, dutyId)
	if err != nil {
		return err
	}
	if res.FirstName != newDutyName {
		filter := bson.M{"_id": res.ID}
		update := bson.M{"$set": bson.M{"FirstName": newDutyName, "LastName": newDutySurname,"IDCategory": newIDCategories}}
		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			//fmt.Printf("update fail %v\n", err)
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (cp *dutyWorkerRep) DeleteDutyWorker(ctx context.Context, dutyWorkerId string) error {

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx, database.URIOnline)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DbNameOnline).Collection(database.DBDutyWorkersTable)
	id, _ := primitive.ObjectIDFromHex(dutyWorkerId)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(deleteResult)
	return nil
}

func (cp *dutyWorkerRep) GetDutyWorkerByID(ctx context.Context, ids string) (*domain.DutyWorkers, error) {
	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx, database.URIOnline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DbNameOnline).Collection(database.DBDutyWorkersTable)
	id, _ := primitive.ObjectIDFromHex(ids)
	filter := bson.D{{"_id", id}}
	var res domain.DutyWorkers
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}

func (cp *dutyWorkerRep) GetDutyWorkers(ctx context.Context) ([]*domain.DutyWorkers, error) {

	var specArray []*domain.DutyWorkers

	con := database.NewConnectDB()
	client, err := con.ConnectDB(ctx, database.URIOnline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer con.DisconnectDB(ctx, client)

	collection := client.Database(database.DbNameOnline).Collection(database.DBDutyWorkersTable)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(ctx) {
		//
		var elem domain.DutyWorkers
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		specArray = append(specArray, &elem)
	}
	cur.Close(ctx)
	return specArray, nil
}
