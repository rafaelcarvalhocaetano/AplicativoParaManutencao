package graphql

import (
	"context"

	"github.com/rafaelcarvalhocaetano/meetup/graphql/model"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *queryResolver) Meetup(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupRepo.GetMeetups()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserRepo.GetUserByID(id)
}
