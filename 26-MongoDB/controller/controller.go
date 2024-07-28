package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RahulGIT24/gomongo/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "netflix"
const colname = "watchlist"

// Most IMP
var collection *mongo.Collection

// connect with db

func init(){
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connected to mongo
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	collection = client.Database(dbName).Collection(colname)

	// collection is ready
	fmt.Println("Collection Refrence is ready")
}

// MONGODB helpers ~ file

// insert 1 record

func insertOneMovie(movie model.Netflix){
	inserted,err := collection.InsertOne(context.Background(),movie)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id: ",inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id" : id}
	update := bson.M{"$set": bson.M{"watched":true}}
	_,err := collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Document Updated")
}

// Delete 1 record
func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	delCount,err := collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count: ",delCount)
}

// Delete all records from mongodb
func deleteAllMovie() int64 {
	deleteResult,err := collection.DeleteMany(context.Background(),bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from mongo

func getAllMovies() []primitive.M{
	cur,err := collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies = append(movies,movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual Controllers ~ File
func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Marked")
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}