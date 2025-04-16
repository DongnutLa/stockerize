package shared_services

import (
	"context"
	"fmt"
	"sync"
	"time"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SharedOrderService struct {
	logger           *zerolog.Logger
	orderSummaryRepo shared_repositories.IOrderSummaryRepository
}

func NewSharedOrdersSummaryService(
	logger *zerolog.Logger,
	orderSummaryRepo shared_repositories.IOrderSummaryRepository,
) shared_ports.ISharedOrdersSummaryService {
	return &SharedOrderService{
		logger:           logger,
		orderSummaryRepo: orderSummaryRepo,
	}
}

func (s *SharedOrderService) HandleOrdersSummary(ctx context.Context, payload map[string]interface{}, topic string) {
	s.logger.Info().Interface("body", payload).Msgf("Message received for topic %s", topic)

	totals := utils.EventDataToStruct[order_domain.Totals](payload["totals"])
	orderType := utils.EventDataToStruct[order_domain.OrderType](payload["orderType"])
	paymentMethod := utils.EventDataToStruct[order_domain.PaymentMethod](payload["paymentMethod"])

	s.CalculateOrdersSummary(ctx, totals, *orderType, *paymentMethod)
}

func getUTCMinus5Date() time.Time {
	now := time.Now().UTC()
	utcMinus5 := now.Add(-5 * time.Hour)
	year, month, day := utcMinus5.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func getCurrentWeekRangeUTC5() (start, end time.Time) {
	now := time.Now().UTC()
	utcMinus5 := now.Add(-5 * time.Hour)

	// Calcular el lunes de esta semana (00:00)
	weekday := utcMinus5.Weekday()
	daysSinceMonday := int(weekday - time.Monday)
	if daysSinceMonday < 0 {
		daysSinceMonday += 7 // Ajuste para domingo
	}
	start = utcMinus5.AddDate(0, 0, -daysSinceMonday).Truncate(24 * time.Hour)

	// Calcular el domingo (23:59:59.999)
	end = start.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second + 999*time.Millisecond)

	return start, end
}

func getCurrentMonthRangeUTC5() (start, end time.Time) {
	now := time.Now().UTC()
	utcMinus5 := now.Add(-5 * time.Hour)

	// Primer día del mes (00:00:00 en UTC-5)
	start = time.Date(utcMinus5.Year(), utcMinus5.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Último día del mes (23:59:59.999 en UTC-5)
	end = time.Date(utcMinus5.Year(), utcMinus5.Month()+1, 0, 23, 59, 59, 999999, time.UTC)

	return start, end
}

func (s *SharedOrderService) GetOrdersSummary(ctx context.Context) ([]shared_domain.OrderSummary, error) {
	currentDate := getUTCMinus5Date()
	weekStart, weekEnd := getCurrentWeekRangeUTC5()
	monthStart, monthEnd := getCurrentMonthRangeUTC5()

	//filter["$text"] = map[string]string{"$search": queryParams.Search}

	opts := shared_ports.FindManyOpts{
		Filter: map[string]interface{}{
			"$or": []bson.M{
				{ // Resumen diario para el 15/04/2025
					"summaryType": "DAILY",
					"start":       bson.M{"$gte": currentDate},
					"end":         bson.M{"$lte": currentDate},
				},
				{ // Resumen semanal que incluye el 15/04/2025
					"summaryType": "WEEKLY",
					"start":       bson.M{"$gte": weekStart},
					"end":         bson.M{"$lte": weekEnd},
				},
				{ // Resumen mensual de abril 2025
					"summaryType": "MONTHLY",
					"start":       bson.M{"$gte": monthStart},
					"end":         bson.M{"$lte": monthEnd},
				},
			},
		},
	}

	var result []shared_domain.OrderSummary

	_, err := s.orderSummaryRepo.FindMany(ctx, opts, &result, false)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SharedOrderService) CalculateOrdersSummary(ctx context.Context, totals *order_domain.Totals, orderType order_domain.OrderType, paymentMethod order_domain.PaymentMethod) {
	currentDate := getUTCMinus5Date()
	weekStart, weekEnd := getCurrentWeekRangeUTC5()
	monthStart, monthEnd := getCurrentMonthRangeUTC5()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		s.saveOrderSummary(ctx, 1, totals.Total, shared_domain.DailySummaryType, orderType, paymentMethod, currentDate, currentDate)
	}()

	go func() {
		defer wg.Done()
		s.saveOrderSummary(ctx, 1, totals.Total, shared_domain.WeeklySummaryType, orderType, paymentMethod, weekStart, weekEnd)
	}()

	go func() {
		defer wg.Done()
		s.saveOrderSummary(ctx, 1, totals.Total, shared_domain.MonthlySummaryType, orderType, paymentMethod, monthStart, monthEnd)
	}()

	wg.Wait()
}

func (s *SharedOrderService) saveOrderSummary(
	ctx context.Context,
	count int,
	total float64,
	summaryType shared_domain.SummaryType,
	orderType order_domain.OrderType,
	paymentMethod order_domain.PaymentMethod,
	start, end time.Time,
) {
	now := time.Now()

	fmt.Println()
	fmt.Printf("Time now: %s\n", now)
	fmt.Printf("Start Date: %s\n", start)
	fmt.Printf("End Date: %s\n", end)
	fmt.Println()

	filter := bson.D{
		{Key: "orderType", Value: string(orderType)},
		{Key: "summaryType", Value: string(summaryType)},
		{Key: "paymentMethod", Value: string(paymentMethod)},
		{Key: "start", Value: start},
		{Key: "end", Value: end},
	}

	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "count", Value: count},
			{Key: "total", Value: total},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "orderType", Value: string(orderType)},
			{Key: "summaryType", Value: string(summaryType)},
			{Key: "paymentMethod", Value: string(paymentMethod)},
			{Key: "updatedAt", Value: &now},
		}},
		{Key: "$setOnInsert", Value: bson.D{
			{Key: "start", Value: start},
			{Key: "end", Value: end},
		}},
	}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var orderSummary shared_domain.OrderSummary

	coll, logger := s.orderSummaryRepo.GetCollection()

	err := coll.FindOneAndUpdate(
		ctx,
		filter,
		update,
		opts,
	).Decode(&orderSummary)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to update orders summary")
		return
	}

	logger.Info().Interface("orderSummary", orderSummary).Msg("Orders summary updated successfully")
}
