package main

type DataPosts struct {
	Data []struct {
		Id   string `_xpath:"./@data-pid"`
		Name string `_xpath:".//li[@class='d_name']/a/text()"`
	} `_xpath:"//div[contains(@class,'l_post')]"`
}
