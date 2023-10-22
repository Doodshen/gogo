//学习reflect 使用go中reflect实现一个SQL构造器

package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) string {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	if v.Kind() != reflect.Struct {
		panic("unspporte argument type ")
	}

	tableName := t.Name() //通过结构体类型提取出sql的表名

	sql := fmt.Sprintf("INSERT INTO %s", tableName)

	columns := "("

	values := "VALUES("

	for i := 0; i < v.NumField(); i++ {
		// 注意reflect.Value 也实现了NumField,Kind这些方法
		// 这里的v.Field(i).Kind()等价于t.Field(i).Type.Kind()
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				columns += fmt.Sprintf("%s", t.Field(i).Name)

				values += fmt.Sprintf("%d", v.Field(i).Int())
			} else {
				columns += fmt.Sprintf(",%s", t.Field(i).Name)
				values += fmt.Sprintf(",%d", v.Field(i).Int())

			}
		case reflect.String:
			if i == 0 {
				columns += fmt.Sprintf("%s", t.Field(i).Name)
				values += fmt.Sprintf("'%s'", v.Field(i).String())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", '%s'", v.Field(i).String())
			}
		}
	}
	columns += ");"
	values += ");"

	sql += columns + values
	fmt.Println(sql)
	return sql
}

func main() {
	o := order{
		ordID:      123,
		customerId: 33,
	}

	createQuery(o)

	e := employee{
		name:    "wzs",
		id:      123,
		address: "lalalla",
		salary:  55555,
		country: "wwww",
	}
	createQuery(e)

}
