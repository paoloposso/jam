package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws/session"
	musicalevents "github.com/paoloposso/jam/src/musical_events"
)

type Repository struct {
	session *session.Session
}

func NewRepository() (*Repository, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return &Repository{
		session: sess,
	}, nil
}

func (repo *Repository) Create(musicalEvent musicalevents.MusicalEvent) (id string, err error) {
	return "", nil
}
