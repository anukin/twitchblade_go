package blade

import "database/sql"

type User struct {
	Name        string
	Password    string
	Transaction *sql.Tx
}

func Stream(username string, transaction *sql.Tx) []Tweetmodel {
	rows, err := transaction.Query("Select id, tweet from tweets where username=$1", username)
	if err != nil {
		panic(err)
	}
	tweets := make([]Tweetmodel, 0)
	for rows.Next() {
		var id int
		var msg string
		rows.Scan(&id, &msg)
		tweet := Tweetmodel{id, msg}
		tweets = append(tweets, tweet)
	}
	return tweets
}
