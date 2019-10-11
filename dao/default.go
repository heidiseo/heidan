package dao

import (
	"context"
	"github.com/heroku/go-getting-started/model"
)

type DAO interface {
	ReadAdventures(ctx context.Context) error
	AddAdventure(ctx context.Context, adventure model.Adventure) (*model.Adventure, error)
	UpdateInFirestore(ctx context.Context) error
	DeleteFromFirestore(ctx context.Context) error
}