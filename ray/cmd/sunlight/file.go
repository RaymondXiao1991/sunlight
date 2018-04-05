package main

import (
	"fmt"
	"log"
	"ray/embrice/constant"
	"ray/embrice/entity"

	"github.com/excelize"
	"github.com/tealeg/xlsx"
	"strconv"
)

//func GetDataFromExcel() map[int]map[int]map[int]string {
func GetDataFromExcel() {

	var (
		excelFilePath string                         = "/home/qingyang/Desktop/sunlight.xlsx"
		fileResult    map[int]map[int]map[int]string = make(map[int]map[int]map[int]string)
		sheetResult   map[int]map[int]string         = make(map[int]map[int]string)
	)
	//打开一个excel文件资源
	f, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		log.Println(err.Error())
	}
	//循环文件中所有工作表
	for sheetKey, sheet := range f.Sheets {
		//循环对应工作表中行数
		for key, row := range sheet.Rows {
			rowResult := make(map[int]string)
			//循环工作表行数的每一列
			for k, cell := range row.Cells {
				rowResult[k] = cell.Value
			}
			//如果为空不添加对应值到 数组
			if !entity.Empty(rowResult) {
				sheetResult[key] = rowResult
			}
		}
		//如果为空不添加对应值到 数组
		if !entity.Empty(sheetResult) {
			fileResult[sheetKey] = sheetResult
		}

	}

	//输出表格的结果
	for _, sheet := range fileResult {
		for k, _ := range sheet {

			log.Printf("%d=%v\n", k, sheet[k])
		}

	}
	//return fileResult
}

func GetDataFromExcel2() []*entity.Goods {
	//func GetDataFromExcel2() {
	goods := make([]*entity.Goods, 0)
	xlsx, err := excelize.OpenFile("/home/qingyang/Desktop/sunlight.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Get value from cell by given worksheet name and axis.
	//	cell := xlsx.GetCellValue("Sheet1", "B2")
	//	fmt.Println("cell:", cell)

	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	for r, row := range rows {
		if r == 0 {
			continue
		}
		product := new(entity.Goods)
		for c, colCell := range row {
			switch c {
			case constant.TYPE_NAME:
				product.Name = colCell
			case constant.TYPE_PRICE:
				product.Price, _ = strconv.ParseFloat(colCell, 64)
			case constant.TYPE_INVENTORY:
				product.Inventory, _ = strconv.Atoi(colCell)
			default:
			}
		}
		goods = append(goods, product)
	}
	return goods
}

func SearchDataFromExcel(name string) []*entity.Goods {
	goods := make([]*entity.Goods, 0)
	xlsx, err := excelize.OpenFile("/home/qingyang/Desktop/sunlight.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	for r, row := range rows {
		if r == 0 {
			continue
		}
		product := new(entity.Goods)
		for c, colCell := range row {
			switch c {
			case constant.TYPE_NAME:
				product.Name = colCell
			case constant.TYPE_PRICE:
				f, _ := strconv.ParseFloat(colCell, 64)
				product.Price = f
			case constant.TYPE_INVENTORY:
				i, _ := strconv.Atoi(colCell)
				product.Inventory = i
			default:
			}
		}
		goods = append(goods, product)
	}
	return goods
}

// CheckErr 检查错误
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

// ExcelParse xlsx文件解析
func ExcelParse(fileName string) []string {
	filePath := "/home/qingyang/Desktop/" + fileName
	xlFile, err := xlsx.OpenFile(filePath)
	CheckErr(err) //自己定义的函数
	//获取行数
	length := len(xlFile.Sheets[0].Rows)
	//开辟除表头外的行数的数组内存
	resourceArr := make([]string, length-1)
	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		//遍历每一行
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
			if rowIndex == 0 {
				// for _, cell := range row.Cells {
				//  text := cell.String()
				//  fmt.Printf("%s\n", text)
				// }
				continue
			}
			//遍历每一个单元
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					//如果是每一行的第一个单元格
					if cellIndex == 0 {
						resourceArr[rowIndex-1] = text
					}
				}
			}
		}
	}
	return resourceArr
}

/*
func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}
*/
