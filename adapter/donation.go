package adapter

import (
	"fmt"
	"log"
	"opn/model"
	"sync"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/shopspring/decimal"
)



func MakeDonation(card *model.Card, donate *model.Donation, mu *sync.Mutex, errCh chan<- error, client *omise.Client) {
	// defer wg.Done()
	// newAmount := decimal.NewFromInt(card.Amount)

	res := &omise.Card{}
	err := client.Do(res, &operations.CreateToken{
		Name:            card.Name,
		Number:          card.CCNumber,
		ExpirationMonth: card.ExpMonth,
		ExpirationYear:  card.ExpYear,
	})
	card.Token = res.ID
	if err != nil {
		switch e:= err.(type) {
		case *omise.ErrInternal:
			errCh <- fmt.Errorf("ErrInternal %v",e.Error())
		case *omise.ErrTransport:
			errCh <- fmt.Errorf("ErrTransport %v",e.Error())
		case *omise.Error:
			if e.StatusCode==400 {
				errCh <- fmt.Errorf("bad request")
			}else if e.StatusCode== 401{
				errCh <- fmt.Errorf("unauthorized")
			}else if e.StatusCode==403{
				errCh <- fmt.Errorf("forbidden")
			}else if e.StatusCode==404{
				errCh <- fmt.Errorf("not found")
			}else if e.StatusCode==422{
				errCh <- fmt.Errorf("unprocessable entity")
			}else if e.StatusCode==429{
				errCh <- fmt.Errorf("too many requests")
			}else if e.StatusCode==500{
				errCh <- fmt.Errorf("internal server error")
			}else if e.StatusCode==503{
				errCh <- fmt.Errorf("service unavailable")
			}
		default:
			errCh <- fmt.Errorf("unknown error: %v",err)
		}
	}
	CreateCharge(card, donate, client, mu,errCh)

}

func CreateCharge(card *model.Card, donate *model.Donation, client *omise.Client, mu *sync.Mutex, errCh chan<- error) {
	newAmount := decimal.NewFromInt(card.Amount)
	result := &omise.Charge{}
	err := client.Do(result, &operations.CreateCharge{
		Amount:   card.Amount,
		Currency: "thb",
		Card:     card.Token,
	})
	if err != nil {
		switch e:= err.(type) {
		case *omise.ErrInternal:
			errCh <- fmt.Errorf("ErrInternal %v",e.Error())
		case *omise.ErrTransport:
			errCh <- fmt.Errorf("ErrTransport %v",e.Error())
		case *omise.Error:
			if e.StatusCode==400 {
				errCh <- fmt.Errorf("bad request")
			}else if e.StatusCode== 401{
				errCh <- fmt.Errorf("unauthorized")
			}else if e.StatusCode==403{
				errCh <- fmt.Errorf("forbidden")
			}else if e.StatusCode==404{
				errCh <- fmt.Errorf("not found")
			}else if e.StatusCode==422{
				errCh <- fmt.Errorf("unprocessable entity")
			}else if e.StatusCode==429{
				errCh <- fmt.Errorf("too many requests")
			}else if e.StatusCode==500{
				errCh <- fmt.Errorf("internal server error")
			}else if e.StatusCode==503{
				errCh <- fmt.Errorf("service unavailable")
			}
		default:
			errCh <- fmt.Errorf("unknown error: %v",err)
		}
	}
	mu.Lock()
	defer mu.Unlock()
	if result.Status == "successful" {
		log.Println("Successful charge")
		donate.Donors = append(donate.Donors, model.Donor{Name: card.Name, Amount: card.Amount})
	} else {
		donate.Fail_donated = donate.Fail_donated.Add(newAmount)
	}
	donate.Total_received = donate.Total_received.Add(newAmount)

}
