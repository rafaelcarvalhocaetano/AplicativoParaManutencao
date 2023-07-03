package graphql

import (
	"context"

	"github.com/rafaelcarvalhocaetano/meetup/graphql/model"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() UserResolver { return &userResolver{r} }

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	// var m []*model.Meetup
	// for _, meetup := range meetups {
	// 	if meetup.UserID == obj.ID {
	// 		m = append(m, meetup)
	// 	}
	// }
	// return m, nil
	return r.MeetupRepo.GetMeetupsForUser(obj)
}
