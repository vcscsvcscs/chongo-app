package model

// Only required to break import cycle

type Token struct {
	Username     string `json:"username" bson:"userName"`
	Token        string `json:"token" bson:"token"`
	TimeAccessed int64  `json:"timeAccessed" bson:"timeAccessed"`
}
