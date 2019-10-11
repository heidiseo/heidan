package dao

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"google.golang.org/api/iterator"
	"log"
	"context"
)

type Firestore struct {
	fs *firestore.Client
	collection string
}

func NewFirestoreDAO(fs *firestore.Client, collection string) DAO {
	return &Firestore{
		fs:         fs,
		collection: collection,
	}
}

//ReadFromFirestore function will read data from firestore
func (fs *Firestore) ReadAdventures(ctx context.Context) error {
	iter := fs.fs.Collection("adventures").Documents(ctx)
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
	return  nil
}

//AddToFirestore will add data to firestore
func (fs *Firestore) AddAdventure(ctx context.Context, adventure model.Adventure) (*model.Adventure, error) {
	_, _, err := fs.fs.Collection("adventures").Add(ctx, adventure)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return nil, err
	}
	return &adventure,nil
}

func (fs *Firestore) UpdateInFirestore(ctx context.Context) error {
	_, err := fs.fs.Collection("adventures").Doc("activities").Set(ctx, map[string]interface{}{
		"priority": "boom boom",
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func (fs *Firestore) DeleteFromFirestore(ctx context.Context) error {
	_, err := fs.fs.Collection("adventures").Doc("activities").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}