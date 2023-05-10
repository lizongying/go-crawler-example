package main

type DataWord struct {
	Id  string `bson:"_id" json:"id"`
	Fan string `json:"fan"`
}
