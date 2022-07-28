package spectrum

import (
	"github.com/xuri/excelize/v2"
)

type Writer interface{
	WriteColumn(dataName string,data []float64) error
}
type ExcelFile struct{
	file *excelize.File
	// $sheet1など
	sheet string
	// 列番号
	column int
}

// 波長を読み込みexcelに書き出す
func (u *UvVisNir) ExportWavelength(e Writer){
	wavelengths := *u.Wavelengths
	e.WriteColumn("波長", wavelengths)
}

// 反射強度を読み込みexcelに書き出す
func (u *UvVisNir) ExportReflectances(e ExcelFile){
	dataName := *u.DataName
	reflectances := *u.Reflectances
	e.WriteColumn(dataName, reflectances)
}

// 列をexcelに書き出す
func (e ExcelFile)WriteColumn(dataName string, data []float64)error{
    column := e.column
	f := e.file

	// dataNameを列の最初の行に書き出す
	// 数値の座標を文字に直す
	cellPosition, err := excelize.CoordinatesToCellName(column, 1)
	if err != nil {
		return err
	}
	// dataNameを列の最初のcellに入れる
	if err:= f.SetCellValue(e.sheet, cellPosition, dataName); err!=nil{
		return err
	}

	// A1,A2などのセルの位置のスライスを作成
	var cellPositions []string
	for i:=1; i<=len(data); i++ {
		cellPosition, err := excelize.CoordinatesToCellName(column, i+1)
		if err != nil{
			return err
		}
		cellPositions = append(cellPositions, cellPosition)
	}
	//excelへのデータの書き出し
	for i:=0; i<len(data); i++ {
		err:= f.SetCellValue(e.sheet, cellPositions[i], data[i])
		if err != nil{
			return err
		}
    }
		// ブックを保存する
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        return err
	}
	return nil
}

// 現存するexcelのファイルを開く
func (e *ExcelFile)OpenFile(filePath string) error {
	var err error
	e.file, err = excelize.OpenFile(filePath)
	if err != nil{
		return err
	}
	return nil
}


// 新しいファイルを作成する
func (e *ExcelFile)NewFile(){
	e.file = excelize.NewFile()
}

// sheetを設定する
func (e *ExcelFile)SetSheet(sheetName string){
	e.sheet = sheetName
}

// 列番号を設定する
func(e *ExcelFile)SetColumn(Column int){
	e.column = Column
}