package pikpay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
	"time"
)

type Mode int

const (
	TestMode Mode = iota
	ProductionMode
)

var urls = map[Mode]string{
	TestMode:       "https://ipgtest.pikpay.ba",
	ProductionMode: "https://ipgtest.pikpay.ba",
}

var NewDeclinedError = errors.New("Transaction declined.")
var NewUnknownError = errors.New("Unknown error occured.")

type PikPay struct {
	mode         Mode
	apiKey       string
	secret       string
	LastResponse *Response
}

type postData struct {
	Transaction TransactionData `xml:"transaction"`
}

func NewPikPay(apiKey, secret string, mode Mode) *PikPay {
	return &PikPay{
		mode:   mode,
		apiKey: apiKey,
		secret: secret,
	}
}

func (p *PikPay) newRequest(payload interface{}) (*Transaction, error) {
	xmlPayload, err := xml.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urls[p.mode], bytes.NewBuffer(xmlPayload))
	if err != nil {
		return nil, err
	}

	return &Transaction{
		r:       req,
		payload: xmlPayload,
		Timeout: time.Second * 30,
		pikpay:  p,
	}, nil
}

func (p *PikPay) NewTransaction(transactionData TransactionData, t TransactionType) (*Transaction, error) {
	payload := postData{
		Transaction: transactionData,
	}

	// TODO
	// calculate digest
	// set other processing data fields
	// determine transaction type
	digest := ""

	if transactionData.OrderDetails != nil {
		// Calculate digest
	}

	pData := &processingData{
		transactionType: transactionTypeText[t],
		digest:          digest,
	}

	payload.Transaction.processingData = pData

	return p.newRequest(payload)
}
