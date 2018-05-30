package main_test

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func TestExcel(t *testing.T) {
	xlsx, err := excelize.OpenFile("string.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	sheetMap := xlsx.GetSheetMap()
	t.Log(sheetMap)
	for _, val := range sheetMap {
		rows, err := xlsx.Rows(val)
		if err != nil {
			t.Fatal(err)
		}
		for rows.Next() {
			for _, col := range rows.Columns() {
				t.Log(col, []byte(col))
			}
		}
	}
}
