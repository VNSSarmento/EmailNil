package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID       string
	Name     string
	Content  string
	CreatAt  time.Time
	Contacts []Contact
}

type Contact struct {
	email string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	contact := make([]Contact, len(emails))

	for index, email := range emails {
		contact[index].email = email
	}
	return &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		CreatAt:  time.Now(),
		Contacts: contact,
	}, nil
}

type Published struct {
	Name     string
	Img      string
	Category string
	Public   []Public
}

type Public struct {
	Name string
}

func NewPublished(name string, urlImag string, categori string, publics []string) (*Published, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	public := make([]Public, len(publics))

	for i, value := range publics {
		public[i].Name = value
	}

	return &Published{
		Name:     "aiopsjdaspjkod",
		Img:      "abc",
		Category: "acbx",
		Public:   public,
	}, nil
}
