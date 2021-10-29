package domain

import "time"

type Meow struct {
	ID        int64
	Body      string
	CreatedOn time.Time
}
