package abstractions

import (
	"context"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type LaneService interface {
	Create(ctx context.Context, lane *entities.Lane) (err error)
	List(ctx context.Context) (lanes []*entities.Lane, err error)
	Lane(ctx context.Context, id int64) (lane *entities.Lane, err error)
	Delete(ctx context.Context, id int64) (err error)
}
