package store

import (
	"crypto/tls"
	"fmt"
	"hex/pkg/types"
	"log"
	"net"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	database string = os.Getenv("DB_NAME")
	password string = os.Getenv("DB_PASSWORD")
)

type MongoStore struct {
	productCollection *mgo.Collection
}

func NewMongoStore() *MongoStore {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s.mongo.cosmos.azure.com:10255", database)}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: database,
		Username: database,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect, go error %v\n", err)
		os.Exit(1)
	}

	// defer session.Close()

	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB(database).C("product")

	return &MongoStore{
		productCollection: collection,
	}
}

func (m *MongoStore) All() ([]types.Product, error) {
	result := []types.Product{}
	err := m.productCollection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return nil, err
	}

	return result, nil
}

func (m *MongoStore) Find(id string) (*types.Product, error) {
	result := types.Product{}
	err := m.productCollection.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return nil, err
	}

	return &result, nil
}

func (m *MongoStore) Insert(p types.Product) error {
	err := m.productCollection.Insert(&types.Product{
		Id:    p.Id,
		Name:  p.Name,
		Brand: p.Brand,
		Price: p.Price,
	})

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return err
	}

	return nil
}

func (m *MongoStore) Delete(id string) error {
	query := bson.M{"_id": id}
	err := m.productCollection.Remove(query)
	if err != nil {
		log.Fatal("Error deleting record: ", err)
		return err
	}

	return nil
}
