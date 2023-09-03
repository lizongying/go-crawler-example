package main

type DataQuestions struct {
	Data []struct {
		Id      string `_css:"@id"`
		Content string `_css:".rich-content-container"`
	} `_css:"div.answer"`
}
