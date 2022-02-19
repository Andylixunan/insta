package global

import "fmt"

func GetDBConnectionStr(user, passwd, dsn, dbname, option string) string {
	s := "%s:%s@tcp(%s)/%s?%s"
	return fmt.Sprintf(s, user, passwd, dsn, dbname, option)
}
