package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Gets called automatically when the package is loaded
func init() {
  username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")

  var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@localhost:27017", username, password)))
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
  URL        string             `bson:"url"`
  Name       string             `bson:"name"`
  ComesBy    string             `bson:"comesby"`
  BringsWith []string           `bson:"bringswith"`
  Stays      []bool              `bson:"stays"`
  Emoji      string             `bson:"emoji"`
}

func GetUserByUrl(url string) (User, error) {
  log.Println("Finding user by:", url)
  collection := client.Database("bday").Collection("users")
  filter := bson.D{{Key: "url", Value: url}}
  var user User
  err := collection.FindOne(context.Background(), filter).Decode(&user)
  if err != nil {
    if err == mongo.ErrNoDocuments {
      fmt.Println("No user with this URL found")
    } else {
      fmt.Println("Error finding user:", err)
    }
  }
  return user, nil
}

func GetAllUsers() ([]User, error) {
  collection := client.Database("bday").Collection("users")

  cursor, err := collection.Find(context.Background(), bson.D{})
  if err != nil {
    return nil, err
  }
  defer cursor.Close(context.Background())

  var users []User
  if err = cursor.All(context.Background(), &users); err != nil {
    return nil, err
  }

  return users, nil
}

func UpdateUser(user User) (User, error) {
  collection := client.Database("bday").Collection("users")

  filter := bson.M{"url": user.URL}
  update := bson.M{
    "$set": bson.M{
      "name":       user.Name,
      "comesby":    user.ComesBy,
      "bringswith": user.BringsWith,
      "stays":      user.Stays,
    },
  }
  _, err := collection.UpdateOne(context.Background(), filter, update)

  if err != nil {
    return User{}, err
  }

  updatedUser, err := GetUserByUrl(user.URL)

  return updatedUser, err
}
