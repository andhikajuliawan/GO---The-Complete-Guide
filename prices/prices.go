package prices

type taxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	taxIncludedPrices map[float64][]float64
}