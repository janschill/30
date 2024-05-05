package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"30.janschill.de/main/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if len(os.Args) < 2 {
			log.Fatal("No command provided")
	}

	switch os.Args[1] {
	case "reset":
			reset()
	case "seed":
			seed()
	default:
			log.Fatal("Unknown command: " + os.Args[1])
	}
}

func reset() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("bday").Collection("users")

	_, err = collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
			log.Fatal(err)
	}

	log.Println("All records deleted successfully.")
}

func shuffleSlice(slice []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func seed() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("bday").Collection("users")

	_, err = collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
			log.Fatal(err)
	}

	emojis := []string{"🐶", "🐱", "🐭", "🐰", "🦊", "🐻", "🐼", "🐻‍❄️", "🐨", "🐯", "🦁", "🐮", "🐷", "🐸", "🐵", "🐥", "🦉", "🐙", "🕷️", "🐒", "🪿", "🦆", "🐺", "🐗", "🐴", "🐝", "🐛", "🦋", "🐌", "🐞", "🐜", "🦂", "🐟", "🦭", "🦧", "🦤", "🦦", "🦔", "🐓", "🦜", "🦫", "🍆", "🌽", "🍌", "🥕", "🌶️", "🥖"}
	shuffleSlice(emojis)
	usersNames := []string{"Jan", "Birgit", "Lisa", "NJ", "Aaron", "Roman", "Beatrice", "Anne", "Hans", "Loli", "Ralle", "Karl", "Malthe", "Joshua", "Mads", "Raphael", "Liv", "Pia", "William", "Rasmus", "Emma", "Jakob", "Jonas", "Caroline"}

	for index, userName := range usersNames {
		fmt.Println("Inserting ", userName)
		_, err = collection.InsertOne(context.Background(), models.User{
			Name:       userName,
			URL:        uuid.New().String(),
			ComesBy:    "",
			BringsWith: []string{},
			Stays:      []bool{false, false, false, false, false},
			Emoji:      emojis[index%len(emojis)],
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

