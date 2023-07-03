// Package mongodb implements mongodb database functionalities.
package mongodb

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateUser creates a new user.
func (m *mongodb) CreateUser(ctx context.Context, user *models.User) error {
	collection := m.db.Collection("users")

	total, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return err
	}

	user.ID = total + 1

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user.
func (m *mongodb) GetUser(ctx context.Context, telegramID int64) (*models.User, error) {
	user := models.User{}

	collection := m.db.Collection("users")

	filter := bson.D{{Key: "telegram_id", Value: telegramID}}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("can not fetch user %d from database: %v", telegramID, err)
	}

	return &user, nil
}

// UpdateLanguage changes user's language.
func (m *mongodb) UpdateLanguage(ctx context.Context, lang models.Language, telegramID int64) error {
	collection := m.db.Collection("users")

	filter := bson.D{{Key: "telegram_id", Value: telegramID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "language", Value: lang}}}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
