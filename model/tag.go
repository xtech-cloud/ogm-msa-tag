package model

type Tag struct {
	Code    string            `bson:"code"`
	Name    string            `bson:"name"`
	Flag    int64             `bson:"flag"`
	Alias   map[string]string `bson:"alias"`
	Keyword []string          `bson:"keyword"`
	Dummy   []string          `bson:"dummy"`
}
