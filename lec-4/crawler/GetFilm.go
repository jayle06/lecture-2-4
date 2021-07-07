package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"ocg.com/test/model"
	"strconv"
	"strings"
)

func GetFilmInfo(pathURL string) ([]model.Film, error) {
	doc, err := goquery.NewDocument(pathURL)
	if err != nil {
		log.Println(err)
	}

	infoList := make([]model.Film, 0)
	doc.Find("table tbody").Each(func (index int, tableHtml *goquery.Selection){
		var info model.Film
		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
			row := make([]string, 0)
			rowHtml.Find("td").Each(func(indexTd int, tableCell *goquery.Selection) {
				row = append(row, tableCell.Text())
				tableCell.Find("a").Each(func(indexA int, imgTag *goquery.Selection) {
					imgTag.Find("img").Each(func(indexImg int, attribute *goquery.Selection) {
						info.Poster, _ = attribute.Attr("src")
					})
				})
			})
			info.ID++
			info.Title = strings.ReplaceAll(strings.TrimSpace(row[1]),"\n","")
			info.Rating, _ = strconv.ParseFloat(strings.TrimSpace(row[2]), 64)
			infoList = append(infoList, info)
		})
	})

	return infoList, nil
}
