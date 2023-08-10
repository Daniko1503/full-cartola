package event

import (
	"context"
	"encoding/json"

	"github.com/Daniko1503/full-cartola/internal/domain/entity"
	"github.com/Daniko1503/full-cartola/internal/usecase"
	"github.com/Daniko1503/full-cartola/pkg/uow"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProcessNewAction struct{}

func (p ProcessNewAction) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.ActionAddInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	actionTable := entity.ActionTable{}
	actionTable.Init()
	addNewActionUsecase := usecase.NewActionAddUseCase(uow, &actionTable)
	err = addNewActionUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
