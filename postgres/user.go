package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rafaelcarvalhocaetano/meetup/graphql/model"
)

type UserRepo struct {
	DB *pg.DB
}

func (m *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User

	err := m.DB.Model(&user).Where("id=?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User

	err := m.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}
