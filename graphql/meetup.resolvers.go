package graphql

import (
	"context"

	"github.com/rafaelcarvalhocaetano/meetup/graphql/model"
)

type meetupResolver struct{ *Resolver }

func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return r.UserRepo.GetUserByID(obj.UserID)
}
