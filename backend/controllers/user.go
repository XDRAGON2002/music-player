package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"example.com/backend/models"
)

var userCollection *mongo.Collection

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	connectionUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("USERS_COLLECTION")
	clientOption := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	userCollection = client.Database(dbName).Collection(colName)
	fmt.Println("User collection ready")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cur, err := userCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var users []primitive.M
	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	defer cur.Close(context.Background())
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user bson.M
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode("User doesn't exist")
			return
		}
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var foundUser models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	filter := bson.M{"email": user.Email}
	err := userCollection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err == nil {
		json.NewEncoder(w).Encode("User already exists")
		return
	}
	values := map[string]string{"name": "Liked Songs"}
	data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:3000/api/playlist/add/", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]string
	json.NewDecoder(resp.Body).Decode(&res)
	user.Playlists = []string{res["InsertedID"]}
	password := []byte(user.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)
	inserted, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(inserted)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var foundUser models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	filter := bson.M{"email": user.Email}
	err := userCollection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
		json.NewEncoder(w).Encode("Incorrect email")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		json.NewEncoder(w).Encode("Incorrect password")
		return
	}
	token := GenerateJWT(foundUser.ID.Hex(), foundUser.Email)
	json.NewEncoder(w).Encode(token)
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)).Time,
	})
}

type CustomClaims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string, email string) string {
	claims := CustomClaims{
		userId,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return string(signed)
}
