package spectrum

import (
	"fmt"
	"testing"
)

func TestSpectrum(t *testing.T) {
	a := ExtractSpectrum("グラフキューブ 6B 20sec.txt")
	fmt.Println(a.GetSpectrumData())
    
    
    f := NewFile("aba", "asd")
    t.Log(a.ExportWavelength(f, 1))
    a.ExportReflectances(f, 2)
    fmt.Println(f)
   
    
}
