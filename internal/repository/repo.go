package repository

import (
	"context"

	"github.com/MurashovVen/outsider-sdk/mongo"

	"outsider-whether/internal/models"
)

type Repository struct {
	db *mongo.Client
}

func New(db *mongo.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) WhetherConfigurationSave(
	ctx context.Context, wc *models.WhetherConfigurationMDBv1,
) (*models.WhetherConfigurationMDBv1, error) {
	_, err := r.db.Database(whetherDBName).Collection(configurationCollectionName).InsertOne(
		ctx, wc)
	if err != nil {
		return nil, err
	}

	return wc, nil
}
