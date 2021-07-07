package main

import (
	"ocg.com/test/crawler"
	"ocg.com/test/dao"
)

func main() {
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	info, _ := crawler.GetFilmInfo(url)
	dao.InsertToDB(info)
}
