package services

import (
	"errors"
	db "music-library-management/models/db"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUsrEmailk(email string) error {
	user := &db.User{}
	userCollection := mgm.Coll(user)
	err := userCollection.First(bson.M{user.Email: email}, user)

	if err != nil {
		return errors.New("Email is already in use")
	}

	return nil
}
