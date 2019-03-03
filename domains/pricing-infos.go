package domains

type PricingInfos struct {
	Period string `json:"period"`
	YearlyIptu      string `json:"yearlyIptu"`
	Price           string `json:"price"`
	RentalTotalPrice string `json:"rentalTotalPrice"`
	BusinessType    string `json:"businessType"`
	MonthlyCondoFee string `json:"monthlyCondoFee"`
}
