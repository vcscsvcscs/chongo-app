package controllers

type response struct {
	Types    []string      `json:"types" bson:"types`
	Contents []interface{} `json:"contents" bson:"contents"`
}

type Authenticationdata struct {
	Token string `json:"token" bson:"token`
}
