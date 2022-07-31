package spectrum

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Writer interface{
	WriteColumn(dataName string,data []float64, columnNumber int) error
}
type ExcelFile struct{
	//ファイル名
	fileName string
	// fileオブジェクト
	file *excelize.File
	// $sheet1など
	sheet string
}

// 波長を読み込みexcelに書き出す
func (u *UvVisNir) ExportWavelength(e Writer, columnNumber int)error{
	wavelengths := *u.Wavelengths
	if err := e.WriteColumn("波長", wavelengths, columnNumber); err != nil{
		return fmt.Errorf("波長の書き出しに失敗 %w", err)
	}
	return nil
}

// 反射強度を読み込みexcelに書き出す
func (u *UvVisNir) ExportReflectances(e ExcelFile, columnNumber int)error{
	dataName := *u.DataName
	reflectances := *u.Reflectances
	if err := e.WriteColumn(dataName, reflectances, columnNumber); err != nil{
		return fmt.Errorf("反射強度の書き出しに失敗 %w", err)
	}
	return nil
}




// 列をexcelに書き出す
func (e ExcelFile)WriteColumn(dataName string, data []float64, columnNumber int)error{
    column := columnNumber
	f := e.file

	// dataNameを列の最初の行に書き出す
	// 数値の座標を文字に直す
	cellPosition, err := excelize.CoordinatesToCellName(column, 1)
	if err != nil {
		return fmt.Errorf("最初の行の書き出しに失敗 %w", err)
	}
	// dataNameを列の最初のcellに入れる
	if err:= f.SetCellValue(e.sheet, cellPosition, dataName); err!=nil{
		return fmt.Errorf("最初の行の書き出しに失敗 %w", err)
	}

	// A1,A2などのセルの位置のスライスを作成
	var cellPositions []string
	for i:=1; i<=len(data); i++ {
		cellPosition, err := excelize.CoordinatesToCellName(column, i+1)
		if err != nil{
			return fmt.Errorf("行の書き出しに失敗 %w", err)
		}
		cellPositions = append(cellPositions, cellPosition)
	}
	//excelへのデータの書き出し
	for i:=0; i<len(data); i++ {
		err:= f.SetCellValue(e.sheet, cellPositions[i], data[i])
		if err != nil{
			return fmt.Errorf("行の書き出しに失敗 %w", err)
		}
    }

	// ブックを保存する
    if err := e.saveFile(); err != nil {
        return fmt.Errorf("ブックの保存に失敗 %w", err)
	}
	return nil
}

// ブックを保存する
func (e *ExcelFile)saveFile() error{
	if err := e.file.SaveAs(e.fileName + ".xlsx"); err != nil {
         return fmt.Errorf("ブックの保存に失敗 %w", err)
	}
	return nil
}


// 現存するexcelのファイルを開く
func (e *ExcelFile)OpenFile(filePath string) error {
	var err error
	e.file, err = excelize.OpenFile(filePath)
	if err != nil{
		return fmt.Errorf("ファイルの読み込みに失敗 %w", err)
	}
	return nil
}


// 新しいファイルを作成する
func (e *ExcelFile)NewFile(fileName string){
	e.fileName = fileName
	e.file = excelize.NewFile()
}

// sheetを設定する
func (e *ExcelFile)SetSheet(sheetName string){
	// もしシートが存在しない場合はシートを作成
	if e.file.GetSheetIndex(sheetName) == -1{
		e.file.NewSheet(sheetName)
	}
	e.sheet = sheetName
}

// sheetを消す
func(e ExcelFile)DeleteSheet(sheetName string){
	//sheetNameのsheetを消す
	e.file.DeleteSheet(sheetName)
}


