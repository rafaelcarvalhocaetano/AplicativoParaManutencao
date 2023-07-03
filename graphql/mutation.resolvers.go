package graphql

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaelcarvalhocaetano/meetup/graphql/model"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	if len(input.Name) < 5 {
		return nil, errors.New("name not long enough")
	}

	meetup := &model.Meetup{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		UserID:      "ecbbaab4-1550-4d2b-abf7-802298543e88",
	}

	return r.MeetupRepo.Create(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {

	meetup, err := r.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exist")
	}

	meetup.Name = *input.Name
	meetup.Description = *input.Description

	meetup, err = r.MeetupRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exist")
	}

	err = r.MeetupRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}

	return true, nil
}
