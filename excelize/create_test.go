package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"testing"
)

func TestCases(t *testing.T) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			panic(err)

		}
	}()
	sheet, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Panic(err)
	}
	for i := 1; i <= 100; i++ {
		if err = f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), "Hello world."+strconv.Itoa(i)); err != nil {
			log.Panic(err)
		}

		if err = f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), 1*i); err != nil {
			log.Panic(err)
		}
	}

	f.SetActiveSheet(sheet)
	if err = f.SaveAs("C://Users//曹帅//Desktop//Book1.xlsx"); err != nil {
		log.Panic(err)
	}
}
