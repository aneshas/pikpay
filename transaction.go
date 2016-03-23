package pikpay

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type TransactionType int

const (
	Authorization TransactionType = iota
	Purchase
	Capture
	Refund
	Void
)

var transactionTypeText = map[TransactionType]string{
	Authorization: "authorization",
	Purchase:      "purchase",
	Capture:       "capture",
	Refund:        "refund",
	Void:          "void",
}

// TODO - Add omitempty to all fields
type BuyerProfile struct {
}

type CardDetails struct {
}

type OrderDetails struct {
	OrderInfo   string `xml:"order-info,omitempty"`
	OrderNumber string `xml:"order-number,omitempty"`
	Amount      int    `xml:"amount,omitempty"`
	Currency    string `xml:"currency,omitempty"`
}

type processingData struct {
	transactionType string `xml:"transaction-type"`
	digest          string `xml:"digest"`
}

type TransactionData struct {
	*BuyerProfile
	*CardDetails
	*OrderDetails
	*processingData
}

type Transaction struct {
	r       *http.Request
	payload []byte
	Timeout time.Duration
	pikpay  *PikPay
}

type Response struct {
	OrderDetails
	ID              int
	Acquirer        string
	ResponseCode    int    `xml:"response-code"`
	ApprovalCode    int    `xml:"approval-code"`
	ResponseMessage string `xml:"response-message"`
	ReferenceNumber int    `xml:"reference-number"`
	Systan          int
	ECI             int
	XID             string
	ACSV            string
	CCType          string `xml:"cc-type"`
	Status          string
	TransactionType string    `xml:"transaction-type"`
	CreatedAt       time.Time `xml:"created-at"`
	Enrollment      string
	Authentication  string
}

func (r *Transaction) Do() (*Response, error) {
	client := http.Client{
		Timeout: r.Timeout,
	}

	resp, err := client.Do(r.r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody := struct {
		Transaction Response
	}{}

	err = xml.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	switch responseBody.Transaction.Status {
	case "declined":
		return nil, NewDeclinedError
	case "invalid":
		return nil, fmt.Errorf(responseText(responseBody.Transaction.ResponseCode))
	case "approved":
		r.pikpay.LastResponse = &responseBody.Transaction
		return &responseBody.Transaction, nil
	}

	return &responseBody.Transaction, NewUnknownError
}
