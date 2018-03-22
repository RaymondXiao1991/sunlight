package main

import (
	//	"encoding/json"
	//"fmt"
	"net/http"
)

func main() {
	//	GetDataFromExcel()
	GetDataFromExcel2()
	//	fmt.Println(ExcelParse("sunlight.xlsx"))

	//http.HandleFunc("/get_all_goods", GetAllGoods)
	//http.HandleFunc("/get_goods_by", GetGoodsBy)
	//http.HandleFunc("/auth_by", AuthBy)

	http.ListenAndServe(":8888", nil)
}

/*
func GetAllGoods(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		value := GetDataFromExcel()
		fmt.Fprint(w, value)
		response, err := json.Marshal(value)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprint(w, string(response))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Only GET accepted")
	}
}
*/
