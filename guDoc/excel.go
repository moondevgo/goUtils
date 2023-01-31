package guDoc

import (
	"fmt"

	basic "github.com/moondevgo/goUtils/guBasic"
	excelize "github.com/xuri/excelize/v2"
)

// # Type
// ## Excel(struct)
type Excel struct {
	FilePath  string
	SheetName string
}

// // ## Sheets(interface)
// type Sheets interface {
// 	SetFilePath() string
// 	SetSheet() string
// 	Open() *excelize.File
// 	// Close() nil
// 	Read() string
// 	Write() bool
// }

// ## headerIndexes
//   - data 내에서 element의 인덱스
func headerIndexes(header []string, fields map[string]string) (indexes map[int]string) {
	indexes = make(map[int]string)
	if len(fields) < 1 { // fields가 없을 때
		for i, h := range header {
			indexes[i] = h
		}
	}
	for k, field := range fields {
		indexes[basic.IndexOf(field, header)] = k
	}
	return indexes
}

// func headerIndexes(header, fields []string) (indexes map[int]string) {
// 	indexes = make(map[int]string)
// 	if len(fields) < 1 {
// 		fields = header
// 	}
// 	for _, field := range fields {
// 		// fmt.Printf("field %v in header %v", field, header)
// 		indexes[basic.IndexOf(field, header)] = field
// 	}
// 	return indexes
// }

// slices [][]string -> []map[string]interface{}
// TODO: fields map[string]string
func mapsFromSlices(slices [][]string, fields map[string]string) (maps []map[string]interface{}) {
	indexes := headerIndexes(slices[0], fields)
	for _, row := range slices[1:] {
		dict := make(map[string]interface{})
		for i, k := range indexes {
			dict[k] = row[i]
		}
		maps = append(maps, dict)
	}
	return maps
}

// func mapsFromSlices(slices [][]string, fields []string) (maps []map[string]interface{}) {
// 	// maps = []map[string]interface{}{}
// 	dict := make(map[string]interface{})
// 	indexes := headerIndexes(slices[0], fields)
// 	for _, row := range slices[1:] {
// 		// dict = map[string]interface{}{}
// 		for i, cell := range row {
// 			dict[indexes[i]] = cell
// 		}
// 		maps = append(maps, dict)
// 	}
// 	return maps
// }

// ** Function(interface implement)
// * Get
// Get FilePath
func (excel *Excel) GetFilePath() string {
	return excel.FilePath
}

// Get SheetName
func (excel *Excel) GetSheetName() string {
	return excel.SheetName
}

// * Set
// Set FilePath
func (excel *Excel) SetFilePath(filePath string) {
	excel.FilePath = filePath
}

// Set FilePath
func (excel *Excel) SetSheetName(sheetName string) {
	excel.SheetName = sheetName
}

// Open Excel File
func (excel *Excel) Open() *excelize.File {
	f, err := excelize.OpenFile(excel.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

// TODO: Struct output
// Read Excel File
// func (excel *Excel) Read(sheetName string, fields []string) []map[string]interface{} {
func (excel *Excel) Read(sheetName string, fields map[string]string) []map[string]interface{} {
	f, err := excelize.OpenFile(excel.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// return rows
	return mapsFromSlices(rows, fields)
}

// // Write Excel File
// func (excel *Excel) Write(sheetName string) bool {
// 	f := excelize.NewFile()
// 	index := f.NewSheet(sheetName)
// 	f.SetCellValue(sheetName, "A2", "Hello World.")
// 	f.SetCellValue(sheetName, "B2", "100")
// 	f.SetActiveSheet(index)
// 	if err := f.SaveAs(excel.FilePath); err != nil {
// 		fmt.Println(err)
// 		return false
// 	}
// 	return true
// }
