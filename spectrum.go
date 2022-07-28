package spectrum

type Spectrum interface {
	GetSpectrumData() map[float64]float64
	error
	GetDataName() string
}
