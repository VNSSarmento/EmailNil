package campaign

import (
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

func NewCampaign(name string, content string, emails []string) *Campaign {
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
	}
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

func NewPublished(name string, urlImag string, categori string, publics []string) *Published {
	public := make([]Public, len(publics))

	for i, value := range publics {
		public[i].Name = value
	}

	return &Published{
		Name:     "aiopsjdaspjkod",
		Img:      "abc",
		Category: "acbx",
		Public:   public,
	}
}
