package dbcontext

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDBConnection(dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func GetDBConnectionStr(user, passwd, dsn, dbname, option string) string {
	s := "%s:%s@tcp(%s)/%s?%s"
	return fmt.Sprintf(s, user, passwd, dsn, dbname, option)
}
