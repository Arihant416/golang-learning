package prices

import (
	"bufio"
	"fmt"
	"os"
	"price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	filePointer, error := os.Open("prices.txt")
	if error != nil {
		fmt.Println("Could not open file ", error)
		return
	}

	scanner := bufio.NewScanner(filePointer)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	error = scanner.Err()
	if error != nil {
		fmt.Println("Could not read file ", error)
		filePointer.Close()
		return
	}

	prices, error := conversion.StringsToFloats(lines)
	if error != nil {
		fmt.Println(error)
		filePointer.Close()
		return
	}
	job.InputPrices = prices
	filePointer.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
