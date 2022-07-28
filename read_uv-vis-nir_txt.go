package spectrum

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// スペクトルデータ抽出
func ExtractSpectrum(file_name string) *UvVisNir {
	s := readTextFile(file_name)
	var wavelengths []float64
	var reflectances []float64
	for i := 2; i < len(s); i++ {
		split_s := strings.Split(s[i], ",")
		wavelength, _ := strconv.ParseFloat(split_s[0], 64)
		reflectance, _ := strconv.ParseFloat(split_s[1], 64)

		//リストに追加
		wavelengths = append(wavelengths, wavelength)
		reflectances = append(reflectances, reflectance)
	}

	return SetSpectraData(file_name, wavelengths, reflectances)

}

// データを１行ずつ読み込みスライスに格納する
func readTextFile(file_name string) []string {
	fp, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	s := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	return s
}

//
