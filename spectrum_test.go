package spectrum

import (
	"fmt"
	"testing"
)

func TestSpectrum(t *testing.T) {
	a := ExtractSpectrum("グラフキューブ 6B 20sec.txt")
	fmt.Println(a.GetSpectrumData())
    
    
    var f ExcelFile
    f.NewFile()
    f.SetSheet("Sheet1")
    f.SetColumn(2)
    a.ExportWavelength(f)
    f.SetColumn(3)
    a.ExportReflectances(f)
    fmt.Println(f)
   
    
}
