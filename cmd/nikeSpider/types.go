package main

type DataCategory struct {
	Id    string `bson:"_id" json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type DataDetail struct {
	Id       string `bson:"_id" json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
