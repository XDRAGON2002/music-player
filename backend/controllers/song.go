package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"example.com/backend/models"
)

var songCollection *mongo.Collection

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}
	connectionUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("SONGS_COLLECTION")
	clientOption := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	songCollection = client.Database(dbName).Collection(colName)
	fmt.Println("Song collection ready")
}


func GetSongs(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	page64 := int64(page)
	options := &options.FindOptions{}
	options.SetSkip((page64 - 1) * 25)
	options.SetLimit(25)
	cur, err := songCollection.Find(context.Background(), bson.D{{}}, options)
	if err != nil {
		log.Fatal(err)
	}
	var songs []primitive.M
	for cur.Next(context.Background()) {
		var song bson.M
		err := cur.Decode(&song)
		if err != nil {
			json.NewEncoder(w).Encode("database exhausted")
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	json.NewEncoder(w).Encode(songs)
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var song bson.M
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := songCollection.FindOne(context.Background(), filter).Decode(&song)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode("Song doesn't exist")
			return
		}
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(song)
}

// TODO: Fix add song to model
func AddSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	var foundSong models.Song
	_ = json.NewDecoder(r.Body).Decode(&song)
	song.Likes = 0
	filter := bson.M{"songid": song.SongID}
	err := songCollection.FindOne(context.Background(), filter).Decode(&foundSong)
	if err == nil {
		json.NewEncoder(w).Encode("Song already exists")
		return
	}
	inserted, err := songCollection.InsertOne(context.Background(), song)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(inserted)
}

// TODO: Fix like/dislike logic
func LikeSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"likes": 1}}
	inserted, err := songCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(inserted)
}

func SearchSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchString := params["search"]
	filter := bson.M{"songname": bson.M{"$regex": primitive.Regex{Pattern: searchString, Options: "i"}}}
	cur, err := songCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var songs []primitive.M
	for cur.Next(context.Background()) {
		var song bson.M
		err := cur.Decode(&song)
		if err != nil {
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	json.NewEncoder(w).Encode(songs)
}
