package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var usersCollection *mongo.Collection

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FullName string             `bson:"fullName,omitempty"`
	Age      int                `bson:"age,omitempty"`
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	usersCollection = client.Database("db").Collection("users")

	// insertOneUser()
	// insertMultipleUser()
	// findAllUsers()
	// findUserById()

	// filterUserByAge()
	// deleteUserByAge()
	// deleteAllUsers()
	// deleteUserById()

	// updateUserById()
}

func insertOneUser() {
	// user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	user := User{FullName: "Mert", Age: 24}
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func insertMultipleUser() {
	// users := []interface{}{
	// 	bson.D{{"fullName", "User 2"}, {"age", 25}},
	// 	bson.D{{"fullName", "User 3"}, {"age", 20}},
	// 	bson.D{{"fullName", "User 4"}, {"age", 28}},
	// }
	userArr := []User{
		User{FullName: "user 5", Age: 25},
		User{FullName: " user 6", Age: 26},
		User{FullName: "user 7", Age: 27},
	}
	users := make([]interface{}, len(userArr))
	for i := range userArr {
		users[i] = userArr[i]
	}

	results, err := usersCollection.InsertMany(context.TODO(), users)
	if err != nil {
		panic(err)
	}
	fmt.Println(results.InsertedIDs)
}

func findAllUsers() {
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	// var results []bson.M
	var results []User
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println("displaying all results in a collection")
	for _, result := range results {
		fmt.Println(result)
	}
}

func findUserById() {

	id, err := primitive.ObjectIDFromHex("632f7212fc6c304094bf3088")
	if err != nil {
		fmt.Println(err)
	}

	findone_result := usersCollection.FindOne(context.TODO(), bson.M{"_id": id})
	// var bson_obj bson.M
	var user User
	if err2 := findone_result.Decode(&user); err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("user:", user)
}

func filterUserByAge() {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	// var results []bson.M
	var results []User
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}
}

func deleteUserByAge() {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// result, err := usersCollection.DeleteOne(context.TODO(), filter)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("deleting the first result from the search filter")
	// fmt.Println("Number of documents deleted:", result.DeletedCount)

	results, err := usersCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("deleting every result from the search filter")
	fmt.Println("Number of documents deleted:", results.DeletedCount)
}

func deleteAllUsers() {

	filter := bson.D{}
	results, err := usersCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("deleting every result from the search filter")
	fmt.Println("Number of documents deleted:", results.DeletedCount)
}

func deleteUserById() {

	id, _ := primitive.ObjectIDFromHex("632f792b94802c501aa2f504")

	filter := bson.M{"_id": id}
	results, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of documents deleted:", results.DeletedCount)
}

func updateUserById() {
	id, _ := primitive.ObjectIDFromHex("632f792b94802c501aa2f504")
	filter := bson.M{"_id": id}
	update := bson.D{{"$set", bson.D{{"age", 24}}}, {"$set", bson.D{{"fullName", "mert dz"}}}}
	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.ModifiedCount)
}
