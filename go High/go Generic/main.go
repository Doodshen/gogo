package main

import (
	"errors"
	"fmt"
	"reflect"
	//"reflect"
)

//解决方法1 使用反射
func twoSum(a, b interface{}) (interface{}, error)  {
    if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
        return nil, errors.New("two value type different")
    }

    switch reflect.TypeOf(a).Kind() {
    case reflect.Int:
        return reflect.ValueOf(a).Int() + reflect.ValueOf(b).Int(), nil
    case reflect.Float64:
        return reflect.ValueOf(a).Float() + reflect.ValueOf(b).Float(), nil
    case reflect.String:
        return reflect.ValueOf(a).String() + " " + reflect.ValueOf(b).String(), nil
    default:
        return nil, errors.New("unknow value type")
    }
}
//泛型解决
func ToSum2[T int|float64|string](a T,b T)T{
	return a+b 
}

func main(){
	sum := ToSum2[int](111,222)
	fmt.Println(sum)

	sum2 := ToSum2[string]("hello","world")
	fmt.Println(sum2)
}