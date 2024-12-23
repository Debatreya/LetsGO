package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"goAPImod/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://admin:password@localhost:27017/gonetflix?authSource=admin"
const dbName = "gonetflix"
const colName = "watchlist"

// MOST IMPORTANT: this is the collection instance that will be used in the controller functions
var collection *mongo.Collection

// connect to mongodb
func init() { // init is a special function that runs when the package is imported
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	col := client.Database(dbName).Collection(colName)

	// collection instance is now ready to be used
	fmt.Println("Collection instance created!")
	collection = col
}

// MONGODB helper functions (Not exported so small letter)

// insert one record
func insertOneMovie(movie models.Netflix) {
	inserted, error := collection.InsertOne(context.Background(), movie)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Inserted one movie: ", inserted.InsertedID)
}

// update one record
func updateOneMovie(movieId string) {
    id, err := primitive.ObjectIDFromHex(movieId)
    if err != nil {
        log.Fatalf("Error converting movieId to ObjectID: %v", err)
    }

    filter := bson.D{{Key: "_id", Value: bson.ObjectID(id)}}
    update := bson.D{{Key: "$set", Value: bson.D{{Key: "watched", Value: true}}}}

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Fatalf("Error updating document: %v", err)
    }

    fmt.Printf("Matched %v documents and modified %v documents.\n", result.MatchedCount, result.ModifiedCount)
}


// delete one record
func deleteOneMovie(movieId string) {
    id, err := primitive.ObjectIDFromHex(movieId)
    if err != nil {
        log.Fatalf("Error converting movieId to ObjectID: %v", err)
    }

    filter := bson.D{{Key: "_id", Value: bson.ObjectID(id)}}

    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        log.Fatalf("Error deleting document: %v", err)
    }

    fmt.Printf("Deleted %v documents.\n", result.DeletedCount)
}


// delete all records
func deleteAllMovies() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}) // delete all movies
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Count: ", deleted.DeletedCount) // check if the movies were deleted\
	return deleted.DeletedCount
}

// find all records
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}}) // find all movies returns a cursor (A cursor is a pointer to the result set of a query)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background()) // close the cursor once the function is done executing

	var movies []primitive.M // slice to store the decoded documents, primitive.M is an unordered representation of BSON documents in Go

	// iterate the cursor like a while loop
	for cur.Next(context.Background()) {
		// create a value into which the single document can be decoded
		var movie primitive.M     // primitive.M vs bson.M is the same thing. bson.M is a map type that is used to represent BSON documents in Go
		err := cur.Decode(&movie) // decode the document into a map, decode the document into the movie variable, &movie is a pointer to the movie variable
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) // append the movie to the movies slice
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return movies
}

// Actual controller - file (must go to other file) (These are Exported functions)

// GetMyAllMovies gets all movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //  application/json vs application/x-www-form-urlencoded is the same thing. It is a MIME type that tells the browser what type of data it’s looking at
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// CreateMovie creates a new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie) // as no retrun from insertOneMovie, we are returning the movie
}

// Mark as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteOneMovie deletes a movie
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteAllMovies deletes all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	// Create a struct to format the response in a better way
	type deletedCount struct {
		DeletedCount int64 `json:"Number of movies deleted"`
	}

	dltCount := deletedCount{
		DeletedCount: count,
	}

	json.NewEncoder(w).Encode(dltCount) // this will create a json of form {"DeletedCount":count}, to format the response in a better way, we can create a struct and then encode it
}