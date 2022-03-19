package db

import (
	"context"
	"log"
	"time"

	"github.com/sonr-io/webauthn.io/models"
	"go.mongodb.org/mongo-driver/bson"
)

type RecordNameObj struct {
	Name  string   `json:"name"`
	Names []string `json:"names"`
	// TimeStamp time.Time
}

func (db *MongoClient) FindDid(did string) *models.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	record := models.User{}
	collection.FindOne(ctx, bson.M{"did": did}).Decode(&record)
	return &record
}

func (db *MongoClient) AddDid(did string, jwt models.Jwt) error {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	record := &models.User{
		Did: did,
		Jwt: jwt,
	}

	res, err := collection.InsertOne(ctx, record)
	if err != nil || res == nil {
		log.Print("\nunable to insert entry into DB in database package\n")
		return err
	}
	return nil
}

func (db *MongoClient) StoreRecord(nameToRecord string, did string) bool {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q1 := bson.M{"did": did}
	q2 := bson.M{"$addToSet": bson.M{"names": nameToRecord}}

	record := models.User{}
	collection.FindOne(ctx, q1).Decode(&record)

	if record.Did == "" {
		return false
	}

	collection.FindOneAndUpdate(ctx, q1, q2)
	return true
}

// check if name is available, if available return true
func (db *MongoClient) CheckName(name string) (bool, error) {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.CountDocuments(ctx, bson.M{"names": name})

	if err != nil {
		return false, err
	}

	//fmt.Println(result)

	return result == 0, nil
}

func (db *MongoClient) NewUser(user models.User) error {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user.Credentials = []models.Credential{}
	collection.InsertOne(ctx, user)
	return nil
}

func (db *MongoClient) FindUserByName(name string) *models.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := &models.User{}
	collection.FindOne(ctx, bson.M{"names": name}).Decode(user)
	return user
}

func (db *MongoClient) AttachDid(placeHolderDid string, newDid string) error {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := models.User{}
	collection.FindOneAndUpdate(ctx, bson.M{"did": placeHolderDid}, bson.M{"$set": bson.M{"did": newDid}}).Decode(user)
	return nil
}

func (db *MongoClient) GetUser(id uint) *models.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := &models.User{}
	collection.FindOne(ctx, bson.M{"model.id": id}).Decode(user)
	return user
}

func (db *MongoClient) GetUserByUsername(username string) *models.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := &models.User{}
	collection.FindOne(ctx, bson.M{"username": username}).Decode(user)
	return user
}

func (db *MongoClient) PutUser(u *models.User) *models.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndUpdate(ctx, bson.M{"id": u.ID}, u).Decode(u)
	return u
}

func (db *MongoClient) RecordPayment(name string) error {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndUpdate(ctx, bson.M{"username": name}, bson.M{"$set": bson.M{"paid": true}})
	return nil
}

func (db *MongoClient) AttachIntent(piID string, name string) {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndUpdate(ctx, bson.M{"username": name}, bson.M{"$set": bson.M{"piid": piID}})
}

func (db *MongoClient) SuccessfulPayment(piID string) {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOneAndUpdate(ctx, bson.M{"piid": piID}, bson.M{"$set": bson.M{"paid": true}})
}
