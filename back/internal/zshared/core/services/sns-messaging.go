package shared_services

import (
	"context"
	"encoding/json"
	"fmt"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/rs/zerolog"
)

type SnsMessaging struct {
	SNS       *sns.SNS
	AccountId string
	Region    string
	logger    *zerolog.Logger
}

func newSnsMessaging(logger *zerolog.Logger) shared_ports.IEventMessaging {
	log := logger.With().Str("resource", "bus_messaging").Logger()
	region := utils.GetConfig("aws_region")
	accoundId := utils.GetConfig("aws_account_id")

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		logger.Panic().Msgf("Error ocurred creating aws client session: %+v", err.Error())
	}

	sns := sns.New(session)

	return &SnsMessaging{
		AccountId: accoundId,
		Region:    region,
		SNS:       sns,
		logger:    &log,
	}
}

func (b *SnsMessaging) SendMessage(ctx context.Context, data *shared_domain.MessageEvent) {
	topicArn := fmt.Sprintf("arn:aws:sns:%s:%s:%s", b.Region, b.AccountId, data.Topic)
	rawData, err := json.Marshal(data)

	if err != nil {
		b.logger.Err(err).Msg("Invalid message")
		return
	}

	input := &sns.PublishInput{
		Message:  aws.String(string(rawData)),
		TopicArn: aws.String(topicArn),
	}

	result, err := b.SNS.Publish(input)
	if err != nil {
		b.logger.Err(err).Interface("message", data).Str("topic", topicArn).Msg("<SNS MESSAGING> Faild to send message")
		return
	}

	b.logger.Info().Interface("message", data).Interface("result", result).Msg("<SNS MESSAGING> Message published successfully")
}
