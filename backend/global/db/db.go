package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDBConnectionStr(user, passwd, dsn, dbname, option string) string {
	s := "%s:%s@tcp(%s)/%s?%s"
	return fmt.Sprintf(s, user, passwd, dsn, dbname, option)
}
