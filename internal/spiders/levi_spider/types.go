package levi_spider

type DataDetail struct {
	Id       string `bson:"_id" json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
