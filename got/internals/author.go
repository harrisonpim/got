package internals

import (
	"fmt"
	"strconv"
	"time"
)

type Author struct {
	Name  string
	Email string
	Time  time.Time
}

func NewAuthor(name, email string) *Author {
	return &Author{
		Name:  name,
		Email: email,
		Time:  time.Now(),
	}
}

func (author *Author) String() string {
	timestamp := strconv.FormatInt(author.Time.Unix(), 10) + " " + author.Time.Format("-0700")
	return fmt.Sprintf("%s <%s> %s", author.Name, author.Email, timestamp)
}
