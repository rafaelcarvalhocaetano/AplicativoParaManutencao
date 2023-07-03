package graphql

import "github.com/rafaelcarvalhocaetano/meetup/postgres"

type Resolver struct {
	MeetupRepo *postgres.MeetupRepo
	UserRepo   *postgres.UserRepo
}
