package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name       string
	URL        string
	ComesBy    string
	BringsWith []string
	Stays      []int
	Emoji      string
}


func shuffleSlice(slice []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func main() {
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
		_, err = collection.InsertOne(context.Background(), User{
			Name:       userName,
			URL:        uuid.New().String(),
			ComesBy:    "",
			BringsWith: []string{},
			Stays:      []int{},
			Emoji:      emojis[index%len(emojis)],
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
