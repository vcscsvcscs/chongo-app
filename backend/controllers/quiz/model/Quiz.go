package model

type Quiz struct {
	Id        string              `json:"id" bson:"id"`
	Owner     string              `json:"username" bson:"username"`
	Name      string              `json:"name" bson:"name"`
	Questions map[string]Question `json:"questions" bson:"questions"`
	Solutions map[string]int      `json:"solutions" bson:"solutions"`
}

type Question struct {
	Type     string         `json:"type" bson:"type"`
	Question string         `json:"question" bson:"question"`
	Media    string         `json:"media" bson:"media"`
	Answers  map[int]string `json:"answers" bson:"answers"`
}
