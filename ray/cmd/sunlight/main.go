package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"ray/embrice/entity"
	"ray/embrice/jsonsql"
	//"strconv"
)

func main() {
	//	GetDataFromExcel()
	//GetDataFromExcel2()
	/*
		for _, goods := range GetDataFromExcel2() {
			fmt.Println(goods)
		}
	*/
	//	fmt.Println(ExcelParse("sunlight.xlsx"))
	for _, goods := range jsonsql.FuzzySearch(GetDataFromExcel2(), "进口") {
		fmt.Println("goods:", goods)
	}

	http.HandleFunc("/get_all_goods", GetAllGoods)
	http.HandleFunc("/get_goods_by", GetGoodsBy)
	//http.HandleFunc("/auth_by", AuthBy)

	http.ListenAndServe(":8888", nil)
}

// GetAllGoods Get list of all goods
func GetAllGoods(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		response := entity.Response{}
		goodsList := GetDataFromExcel2()

		if !response.HanderSuccess(w, &goodsList) {
			response.HandlerFailed(w, "Parse excel failed...")
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Only GET accepted")
	}
	return
}

// GetGoodsBy Get goods By name
func GetGoodsBy(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		if len(values.Get("name")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Wrong input")
			return
		}

		name := string(values.Get("name"))

		response := entity.Response{}
		goodsList := SearchDataFromExcel(name)

		if !response.HanderSuccess(w, &goodsList) {
			response.HandlerFailed(w, "Parse excel failed...")
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Only GET accepted")
	}
	return
}
