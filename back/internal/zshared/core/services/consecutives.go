package shared_services

import (
	"context"
	"fmt"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConsecutiveService struct {
	logger          *zerolog.Logger
	consecutiveRepo shared_repositories.IConsecutiveRepository
}

func NewConsecutiveService(
	logger *zerolog.Logger,
	repository shared_repositories.IConsecutiveRepository,
) shared_ports.IConsecutiveService {
	return &ConsecutiveService{
		logger:          logger,
		consecutiveRepo: repository,
	}
}

// ObtenerConsecutivo obtiene e incrementa un consecutivo
func (c *ConsecutiveService) GetConsecutive(ctx context.Context, tipo shared_domain.ConsecutiveType) (string, error) {
	filter := bson.M{"_id": tipo}
	update := bson.M{"$inc": bson.M{"sequence": 1}}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var consecutive shared_domain.Consecutive

	coll, logger := c.consecutiveRepo.GetCollection()

	err := coll.FindOneAndUpdate(
		ctx,
		filter,
		update,
		opts,
	).Decode(&consecutive)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to get consecutive")
		return "", err
	}

	// Formatear el número con prefijo si existe
	formattedNumber := fmt.Sprintf("%d", consecutive.Sequence)
	if consecutive.Prefix != "" {
		formattedNumber = consecutive.Prefix + formattedNumber
	}

	return formattedNumber, nil
}

func (c *ConsecutiveService) RestoreConsecutive(ctx context.Context, tipo shared_domain.ConsecutiveType) error {
	filter := bson.M{"_id": tipo}
	update := bson.M{"$inc": bson.M{"sequence": -1}}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var consecutive shared_domain.Consecutive

	coll, logger := c.consecutiveRepo.GetCollection()

	err := coll.FindOneAndUpdate(
		ctx,
		filter,
		update,
		opts,
	).Decode(&consecutive)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to restore consecutive")
		return err
	}

	logger.Info().Interface("consecutive", consecutive).Msg("Consecutive restored successfully")

	return nil
}
