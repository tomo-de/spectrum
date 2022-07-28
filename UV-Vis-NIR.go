package spectrum

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// uv_vis_nirのデータを格納する
type UvVisNir struct {
	DataName      *string
	Wavelengths   *[]float64
	Reflectances  *[]float64
	Pitch         *decimal.Decimal
	MaxWaveLength *float64
	MinWaveLength *float64
}

//コンストラクター関数
func SetSpectraData(data_name string, wavelengths []float64, reflectances []float64) *UvVisNir {

	data := &UvVisNir{
		DataName:     &data_name,
		Wavelengths:  &wavelengths,
		Reflectances: &reflectances,
	}

	data.checkPitch()
	return data
}

//pitchを導き構造体に保存
func (u *UvVisNir) checkPitch() error {
	l := *u.Wavelengths
	var validationPitch decimal.Decimal

	for i := 0; i <= len(l)-2; i++ {

		// decimalを使用してpitchを正確に導く
		v := decimal.NewFromFloat(l[i])
		n := decimal.NewFromFloat(l[i+1])

		if i == 0 {
			//初回は検証用の値を用意
			validationPitch = n.Sub(v)
		} else {
			// 以降は検証用のピッチの比較する
			pitch := n.Sub(v)
			if validationPitch.Equal(pitch) {
				return errors.New("pitchが一致しません")
			}

		}

	}
	u.Pitch = &validationPitch
	return nil
}


// データ名を取得する
func (u *UvVisNir) GetDataName() string {
	return *u.DataName
}

// スペクトルのピッチを取得する
func(u *UvVisNir) GetPitch() float64 {
	pitch := *u.Pitch
	return pitch.InexactFloat64()
}

//スペクトルのデータを取得する
func (u *UvVisNir) GetSpectrumData() ([]float64, error) {
	if u.Reflectances == nil {
		msg := "Reflectances == nil"
		return nil, fmt.Errorf("err %s", msg)
	} else {
		x := *u.Reflectances
		return x, nil
	}
}
