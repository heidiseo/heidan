package main

import (
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
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
	ReadFromFirestore(client)
	AddToFirestore(client)
	ReadFromFirestore(client)
	UpdateInFirestore(client)
	ReadFromFirestore(client)
	DeleteFromFirestore(client)
	ReadFromFirestore(client)
	defer client.Close()
}

//ReadFromFirestore function will read data from firestore
func ReadFromFirestore(c *firestore.Client) error {
	ctx := context.Background()
	iter := c.Collection("adventures").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}
	return nil
}

//AddToFirestore will add data to firestore
func AddToFirestore(c *firestore.Client) error {
	ctx := context.Background()
	_, _, err := c.Collection("adventures").Add(ctx, map[string]interface{}{
		"type":    "libraries",
		"name": "British Library",
		"location": "London King's Cross",
		"price": 0,
		"estimated_duration": "2 hours",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func UpdateInFirestore(c *firestore.Client) error {
	ctx := context.Background()
	_, err := c.Collection("adventures").Doc("activities").Set(ctx, map[string]interface{}{
		"priority": "boom boom",
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func DeleteFromFirestore(c *firestore.Client) error {
	ctx := context.Background()
	_, err := c.Collection("adventures").Doc("activities").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}