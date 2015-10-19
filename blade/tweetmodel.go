package blade

import (
	_ "database/sql"
)

type Tweetmodel struct {
	id      int
	message string
}
