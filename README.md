```
// txtファイルを読み込む
a := spectrum.ExtractSpectrum("グラフキューブ 6B 20sec.txt")
//データのプリント
fmt.Println(a.GetSpectrumData())

//excelへの書き出し
var f spectrum.ExcelFile
//ファイルの設定
f.NewFile()
//既存のファイルはOpenFile

//シートの設定
f.SetSheet("Sheet1")

//書き出す列番号
f.SetColumn(1)
//波長の書き出し
a.ExportWavelength(f)


//書き出す列番号
f.SetColumn(2)
//反射強度の書き出し
a.ExportReflectances(f)
```
