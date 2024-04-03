package main

import (
	"bufio"
	"fmt"
	"log"

	"opn/adapter"
	"opn/cipher"
	"opn/config"
	"opn/model"
	"opn/utils"
	"os"
	"strconv"

	"strings"
	"sync"

	"github.com/omise/omise-go"
)

func main() {
	conf := new(model.Configuration)
	config.SetUpConfiguration(conf)
	file, err := os.Open("./data/fng.1000.csv.rot128")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rotReader, err := cipher.NewRot128Reader(file)
	if err != nil {
		panic(err)
	}
	errCh := make(chan error, 400)
	go func() {
		for err := range errCh {
			if err != nil {
				log.Printf("error channel : %v", err)
			}
		}
	}()
	scanner := bufio.NewScanner(rotReader)
	donation := model.Donation{}
	log.Println("performing donations...")
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	client, err := omise.NewClient(conf.App.PublicKey, conf.App.SecretKey)
	if err != nil {
		fmt.Println(err)
	}
	for scanner.Scan() {
		plaintext := scanner.Text()
		data := strings.Split(plaintext, ",")
		card := &model.Card{}
		err := createCardPayload(data, card)
		if err != nil {
			//skip csv headers which is not being calculate
			continue
		}
		wg.Add(1)
		go func(card *model.Card) {
			defer wg.Done()
			adapter.MakeDonation(card, &donation, mu, errCh, client)
		}(card)

	}

	wg.Wait()
	log.Println("done.")
	donation.PrintDonationSummary()
}

func createCardPayload(data []string, card *model.Card) error {
	amount, err := strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		//skip csv headers which is not being calculate
		return err
	}
	// expYear, err := strconv.Atoi(strings.TrimSpace(data[5]))
	expYear, err := strconv.Atoi(data[5])
	if err != nil {
		//skip csv headers which is not being calculate
		return err
	}
	month, err := strconv.Atoi(data[4])
	if err != nil {
		//skip csv headers which is not being calculate
		return err
	}
	expMonth, err := utils.GetMonth(month)
	if err != nil {
		//skip csv headers which is not being calculate
		return err
	}
	card.Name = data[0]
	card.CCNumber = data[2]
	card.ExpYear = expYear
	card.ExpMonth = expMonth
	card.Amount = amount
	return nil

}
