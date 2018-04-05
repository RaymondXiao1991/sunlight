package jsonsql

import (
	"encoding/json"
	"fmt"
	"github.com/elgs/jsonql"
	"ray/embrice/entity"
	"reflect"
	"strings"
)

type UniversalDTO struct {
	Data interface{} `json:"data"`
	// more fields with important meta-data about the message...
}

// FuzzySearch 模糊查询
func FuzzySearch(goods []*entity.Goods, pattern string) []*entity.Goods {

	//struct 到json str
	b, err := json.Marshal(goods)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	parser, err := jsonql.NewStringQuery(string(b))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	pattern = "name ~= '" + pattern + ".*'"
	maps, err := parser.Query(pattern)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//fmt.Println("maps:", maps)

	ggoods := make([]*entity.Goods, 0)

	if err = json.Unmarshal([]byte(maps), &ggoods); err != nil {
		fmt.Println("error")
	}

	fmt.Println("ggoods:", ggoods)

	//maps = appendAssetsToTemplateData(maps)

	//fmt.Println("maps:", maps)

	v, ok := maps.(*entity.Goods)
	fmt.Println("v", v)
	fmt.Println("ok:", ok)
	if !ok {
		return nil
	}
	fmt.Println("vvvvvvvvvvvv")
	fmt.Printf("%+v\n", v)

	return nil
}

func appendAssetsToTemplateData(t interface{}) interface{} {
	fmt.Println("reflect.TypeOf(t).Kind()---------------", reflect.TypeOf(t).Kind())
	fmt.Println("reflect.ValueOf(t)---------------", reflect.ValueOf(t))

	//DoFiledAndMethod(t)

	switch reflect.TypeOf(t).Kind() {
	case reflect.Struct:
		fmt.Println("struct---------------")
		s := reflect.ValueOf(t)
		fmt.Println("s:", s)
	case reflect.Slice:
		PrintVar(t, 0)

		for j := 0; j < reflect.ValueOf(t).Len(); j++ {
			fmt.Println("j:", j)
			fmt.Println(reflect.ValueOf(t).Index(j))
			fmt.Println("reflect.TypeOf(reflect.ValueOf(t).Index(j)).Kind()---------------", reflect.TypeOf(reflect.ValueOf(t).Index(j)).Kind())

		}

	}

	return t
}

//type:interface value:sturct
func PrintStruct(t reflect.Type, v reflect.Value, i int) {
	fmt.Println("------------------------------1")
	fmt.Println("")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(strings.Repeat(" ", i), t.Field(i).Name, ":")
		value := v.Field(i)
		PrintVar(value.Interface(), i+2)
		fmt.Println("")
	}
}

func PrintArraySlice(v reflect.Value, pc int) {
	fmt.Println("------------------------------2")
	for j := 0; j < v.Len(); j++ {
		PrintVar(v.Index(j).Interface(), pc+2)
	}
}
func PrintMap(v reflect.Value, pc int) {
	fmt.Println("------------------------------3")
	for _, k := range v.MapKeys() {
		PrintVar(k.Interface(), pc)
		PrintVar(v.MapIndex(k).Interface(), pc)
	}
}

func PrintVar(i interface{}, ident int) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(i).Elem()
		t = v.Type()
	}
	switch v.Kind() {
	case reflect.Array:
		PrintArraySlice(v, ident)
	case reflect.Chan:
		fmt.Println("Chan")
	case reflect.Func:
		fmt.Println("Func")
	case reflect.Interface:
		fmt.Println("Interface")
	case reflect.Map:
		PrintMap(v, ident)
	case reflect.Slice:
		PrintArraySlice(v, ident)
	case reflect.Struct:
		PrintStruct(t, v, ident)
	case reflect.UnsafePointer:
		fmt.Println("UnsafePointer")
	default:
		fmt.Println("v.Kind():", v.Kind())
		fmt.Print(strings.Repeat(" ", ident), v.Interface())
	}
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
