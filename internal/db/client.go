package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllColection(DB *mongo.Database, colection string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := DB.Collection(colection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func FilterCollection(DB *mongo.Database, collection string, filter bson.M) ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if filter["id"] != nil && filter["id"] != "" {
		objID, err := primitive.ObjectIDFromHex(filter["id"].(string))
		if err != nil {
			return nil, fmt.Errorf("invalid id format: %v", err)
		}
		filter["_id"] = objID
		delete(filter, "id")
	}

	cursor, err := DB.Collection(collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func PutRecord(DB *mongo.Database, collection string, record bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertResult, err := DB.Collection(collection).InsertOne(ctx, record)
	if err != nil {
		return fmt.Errorf("failed to insert record: %v to collection: %s", err, collection)
	}
	fmt.Printf("Inserted record: %s, to collection: %s\n", insertResult.InsertedID, collection)
	return nil
}

func DeleteRecord(DB *mongo.Database, collection string, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format: %v", err)
	}

	fmt.Printf("Deleting record: %s from collection: %s\n", objID, collection)

	filter := bson.M{"_id": objID}
	deleteResult, err := DB.Collection(collection).DeleteOne(ctx, filter)
	fmt.Println(deleteResult.DeletedCount)
	if err != nil || deleteResult.DeletedCount == int64(0) {
		return fmt.Errorf("failed deleting record")
	}

	fmt.Println("Deleted record ID:", deleteResult)

	return nil
}

func UpdateRecord(DB *mongo.Database, collection string, id string, record bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format: %v", err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": record}

	result, err := DB.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == int64(0) {
		return fmt.Errorf("failed to update record: %v to collection: %s", err, collection)
	}
	fmt.Printf("Updated record: %s, to collection: %s\n", id, collection)
	return nil
}
