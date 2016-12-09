package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct{}

var (
	db  *gorm.DB
	err error
)

type Wisdom struct {
	ID                int `gorm:"primary_key,AUTO_INCREMENT"`
	Title             string
	WithoutDiacritics string
	Text              string
	Subject           string
}

func Open() *Database {
	db, err = gorm.Open("sqlite3", "database.sqlite3")
	if err != nil {
		panic(err)
	}
	return &Database{}
}

func (d *Database) SearchWisdom(query string) (list []Wisdom) {
	db.Table("wisdoms").Where("without_diacritics LIKE ?", "%"+query+"%").Scan(&list)
	return
}

func (d *Database) GetWisdom(a int) (u Wisdom) {
	db.Table("wisdoms").Find(&u, "id = ?", a)
	return
}

func (d *Database) SearchWisdomByCategory(query string) (list []Wisdom) {
	db.Table("wisdoms").Where("Subject LIKE ?", "%"+query+"%").Scan(&list)
	return
}
