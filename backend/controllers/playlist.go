package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"example.com/backend/models"
)

var playlistCollection *mongo.Collection

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	connectionUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("PLAYLIST_COLLECTION")
	clientOption := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	playlistCollection = client.Database(dbName).Collection(colName)
	fmt.Println("Playlist collection ready")
}

func GetAllPlaylists(w http.ResponseWriter, r *http.Request) {
	cur, err := playlistCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var playlists []primitive.M
	for cur.Next(context.Background()) {
		var playlist bson.M
		err := cur.Decode(&playlist)
		if err != nil {
			log.Fatal(err)
		}
		playlists = append(playlists, playlist)
	}
	defer cur.Close(context.Background())
	json.NewEncoder(w).Encode(playlists)
}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var playlist bson.M
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := playlistCollection.FindOne(context.Background(), filter).Decode(&playlist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode("Playlist doesn't exist")
			return
		}
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(playlist)
}

func AddPlaylist(w http.ResponseWriter, r *http.Request) {
	var playlist models.Playlist
	_ = json.NewDecoder(r.Body).Decode(&playlist)
	playlist.Songs = []string{}
	inserted, err := playlistCollection.InsertOne(context.Background(), playlist)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(inserted)
}

func DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	id, _ := primitive.ObjectIDFromHex(params["userid"])
	filter := bson.M{"_id": id}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	var idx int
	for i, item := range user.Playlists {
		if item == params["playlistid"] {
			idx = i
			break
		}
	}
	user.Playlists = append(user.Playlists[:idx], user.Playlists[idx+1:]...)
	update := bson.M{"$set": bson.M{"playlists": user.Playlists}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	id, _ = primitive.ObjectIDFromHex(params["playlistid"])
	filter = bson.M{"_id": id}
	result, err := playlistCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

type IDStruct struct {
	PlaylistID string `json:"playlistId"`
	SongID     string `json:"songId"`
}

func AddPlaylistSong(w http.ResponseWriter, r *http.Request) {
	var ids IDStruct
	var playlist models.Playlist
	_ = json.NewDecoder(r.Body).Decode(&ids)
	playlistId, _ := primitive.ObjectIDFromHex(ids.PlaylistID)
	filter := bson.M{"_id": playlistId}
	err := playlistCollection.FindOne(context.Background(), filter).Decode(&playlist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode("Playlist doesn't exist")
			return
		}
		log.Fatal(err)
	}
	playlist.Songs = append(playlist.Songs, ids.SongID)
	update := bson.M{"$set": bson.M{"songs": playlist.Songs}}
	inserted, err := playlistCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(inserted)
}

func DeletePlaylistSong(w http.ResponseWriter, r *http.Request) {
	var ids IDStruct
	var playlist models.Playlist
	_ = json.NewDecoder(r.Body).Decode(&ids)
	playlistId, _ := primitive.ObjectIDFromHex(ids.PlaylistID)
	filter := bson.M{"_id": playlistId}
	err := playlistCollection.FindOne(context.Background(), filter).Decode(&playlist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode("Playlist doesn't exist")
			return
		}
		log.Fatal(err)
	}
	var idx int
	for i, item := range playlist.Songs {
		if item == ids.SongID {
			idx = i
			break
		}
	}
	playlist.Songs = append(playlist.Songs[:idx], playlist.Songs[idx+1:]...)
	update := bson.M{"$set": bson.M{"songs": playlist.Songs}}
	inserted, err := playlistCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(inserted)
}
