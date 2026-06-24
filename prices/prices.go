package prices

type taxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	taxIncludedPrices map[float64][]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *taxIncludedPriceJob {
	return &taxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30, 40},
		TaxRate:     taxRate,
	}
}
