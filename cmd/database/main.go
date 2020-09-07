package main

// 通过反射的方式直接映射数据的列
import (
	"database/sql"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

var query = `select name, age from info`

func main() {
	u := User{}
	exec(&u)
	log.Info().Interface("user", u).Msg("user info")
}

type User struct {
	Id   int64
	Name string
	Age  int
}

func exec(u interface{}) {
	db, err := sql.Open("mysql", "root:root123@/tester")
	if err != nil {
		log.Error().Err(err).Msg("database connection failed")
		return
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Error().Err(err).Msg("prepare failed")
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Error().Err(err).Msg("query failed")
		return
	}

	for rows.Next() {
		name := reflect.New(reflect.ValueOf("").Type()).Interface()
		age := reflect.New(reflect.ValueOf(0).Type()).Interface()
		if err = rows.Scan(name, age); err != nil {
			log.Error().Err(err).Msg("can't not scan record")
			continue
		}
		log.Info().Interface("name", name).Interface("age", age).Msg("output record info")
	}
}

//func value(u interface{}) []interface{} {
//	var itrs []interface{}
//	ty := reflect.TypeOf(u)
//	val := reflect.ValueOf(u)
//	if ty.Kind() == reflect.Ptr {
//		ty = ty.Elem()
//		val = val.Elem()
//	}
//
//	for i := 0; i < ty.NumField(); i++ {
//		itrs = append(itrs, val.Field(i).Interface())
//	}
//	return itrs
//}
