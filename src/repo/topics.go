package repo

import "time"

type Topic struct {
	Name    string
	Created string
	Updated string
	Author  string
}

func GetTopic() *Topic {
	t, _ := GetTopics()
	return t[0]
}

func GetTopics() ([]*Topic, error) {
	topics := make([]*Topic, 0)
	topics = append(topics, &Topic{
		Name:    "Aboba",
		Author:  "admin",
		Created: time.Now().Format(time.RFC822),
		Updated: time.Now().Format(time.RFC822),
	})
	return topics, nil
}
