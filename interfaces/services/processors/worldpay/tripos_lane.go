package worldpay

import (
	"context"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
	"github.com/thermeon/boilerplate-golang/infrastructures/clients/triposlane"
	"github.com/thermeon/boilerplate-golang/infrastructures/clients/triposlane/lanes"
)

func NewLaneAPI() abstractions.LaneService {
	return &laneAPI{
		API: triposlane.NewHTTPClientWithConfig(nil, nil),
	}
}

type laneAPI struct {
	API *triposlane.TriPOSLaneAPI
}

func (l laneAPI) Create(ctx context.Context, lane *entities.Lane) error {
	lpp := lanes.LanesPostParams{}
	laneOk, err := l.API.Lanes.LanesPost(&lpp)
	lane.LaneID = int64(laneOk.Payload.LaneID)
	return err
}

func (l laneAPI) List(ctx context.Context) ([]*entities.Lane, error) {
	//TODO implement me
	panic("implement me")
}

func (l laneAPI) Lane(ctx context.Context, i int64) (*entities.Lane, error) {
	//TODO implement me
	panic("implement me")
}

func (l laneAPI) Delete(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}
