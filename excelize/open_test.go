package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"testing"
)

func TestOpen(t *testing.T) {
	file, err := excelize.OpenFile("C://Users//曹帅//Desktop//Book1.xlsx")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Panic(err)
		}
	}()
	//输出所有的sheet
	list := file.GetSheetList()
	fmt.Println(list)
	//获取第一个sheet index 从0开始
	index, err := file.GetSheetIndex("Sheet1")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(index)
	//获取第一个sheet的A2的值
	value, err := file.GetCellValue("Sheet1", "A2")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(value)

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Panic(err)
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
