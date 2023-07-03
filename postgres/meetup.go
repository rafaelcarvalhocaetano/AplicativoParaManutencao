package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rafaelcarvalhocaetano/meetup/graphql/model"
)

type MeetupRepo struct {
	DB *pg.DB
}

func (m *MeetupRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup

	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupRepo) Create(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	return meetup, err
}

func (m *MeetupRepo) GetByID(id string) (*model.Meetup, error) {
	var meetup model.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

func (m *MeetupRepo) Update(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

func (m *MeetupRepo) Delete(meetup *model.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}

func (m *MeetupRepo) GetMeetupsForUser(user *model.User) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}
