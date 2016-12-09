package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/PuerkitoBio/goquery"
)

type Wisdom struct {
	ID                int `gorm:"primary_key,AUTO_INCREMENT"`
	Title             string
	WithoutDiacritics string
	Text              string
	Subject           string
}

func main() {
	db, err := gorm.Open("sqlite3", "database.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for j := 0; j < 17; j++ {
		doc, err := goquery.NewDocument("http://farsi.balaghah.net/%D9%85%D8%AD%D9%85%D8%AF-%D8%AF%D8%B4%D8%AA%DB%8C/%D8%AA%D8%B1%D8%AC%D9%85%D9%87?field_translation_id_value=&page=" + strconv.Itoa(j))
		if err != nil {
			log.Fatal(err)
		}

		if !db.HasTable(&Wisdom{}) {
			db.CreateTable(&Wisdom{})
		}

		doc.Find(".cols-3 a").Each(func(i int, s *goquery.Selection) {
			link, exist := s.Attr("href")
			if !exist {
				return
			}

			title := s.Text()
			fmt.Println(title)

			doc, err := goquery.NewDocument("http://farsi.balaghah.net" + link)
			if err != nil {
				log.Fatal(err)
			}

			ltext := make([]string, 0)
			mode := make([]string, 0)
			doc.Find("#block-system-main > .content > .nb-node-body p strong").Each(func(i int, s *goquery.Selection) {
				text := s.Text()
				if strings.Contains(text, "اخلاقى") {
					mode = append(mode, "اخلاقی")
				}
				if strings.Contains(text, "سياسى") {
					mode = append(mode, "سياسى")
				}
				if strings.Contains(text, "اجتماعى") {
					mode = append(mode, "اجتماعى")
				}
				if strings.Contains(text, "علمى") {
					mode = append(mode, "علمى")
				}
				if strings.Contains(text, "اعتقادى") {
					mode = append(mode, "اعتقادى")
				}
				if strings.Contains(text, "معنوى") {
					mode = append(mode, "معنوى")
				}
				if len(text) < 40 {
					return
				}
				ltext = append(ltext, text)
			})

			wisdom := Wisdom{}
			for _, t := range ltext {
				wisdom.Text += t + "\n"
			}
			for _, t := range mode {
				wisdom.Subject += t + " "
			}
			wisdom.Title = title
			wisdom.WithoutDiacritics = RemoveDiacritics(wisdom.Text)

			db.Table("wisdoms").Save(&wisdom)
		})
	}
}
