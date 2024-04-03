package model

import (
	"log"
	"opn/utils"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type Donation struct {
	Total_received decimal.Decimal
	Fail_donated   decimal.Decimal
	Donors         []Donor
}
type Donor struct {
	Name   string
	Amount int64
}
type Card struct {
	CCNumber string
	Name     string
	Token    string
	ExpYear  int
	ExpMonth time.Month
	Amount   int64
}
func (donation *Donation) PrintDonationSummary() {
	people := decimal.NewFromInt(int64(len(donation.Donors)))
	if people.IsZero() {
		people = decimal.NewFromInt(1)
	}
	log.Printf("total received: THB %v\n", utils.FormatDecimalWithCommas(donation.Total_received))
	log.Printf("successfully donated: THB %v\n", utils.FormatDecimalWithCommas(donation.Total_received.Sub(donation.Fail_donated)))
	log.Printf("faulty donation: THB %v\n", utils.FormatDecimalWithCommas(donation.Fail_donated))
	log.Printf("average per person: THB %v\n", utils.FormatDecimalWithCommas(donation.Total_received.Sub(donation.Fail_donated).Div(people)))
	log.Printf("top donors: ")
	for _, donor := range findTopthreeDonors(donation.Donors) {
		log.Printf("%s\n", donor)
	}
}
func findTopthreeDonors(donors []Donor) (topthree []string) {
	if len(donors) == 0 {
		return topthree
	}
	sort.Slice(donors, func(i, j int) bool {
		return donors[i].Amount > donors[j].Amount
	})
	topthree = append(topthree, donors[0].Name, donors[1].Name, donors[2].Name)
	return topthree
}
