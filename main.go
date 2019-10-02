package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./firebaseKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	ReadFromFirestore(ctx, client)
	AddToFirestore(ctx, client)
	//ReadFromFirestore(ctx, client)

	defer client.Close()
}

//ReadFromFirestore function will read data from firestore
func ReadFromFirestore(ctx context.Context, c *firestore.Client) {
	iter := c.Collection("adventures").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

//AddToFirestore will add data to firestore
func AddToFirestore(ctx context.Context, c *firestore.Client) {
	_, _, err := c.Collection("adventures").Add(ctx, map[string]interface{}{
		"concert": "Jon Bellion",
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}
