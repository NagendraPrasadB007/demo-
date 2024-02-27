package domain

import "time"

type Paymentdetails struct {
	Paymenttranid   string    `json:"paymenttranid"`
	Pickuprequestid int       `json:"pickuprequestid"`
	Articleid       string    `json:"articleid"`
	Paymenttype     string    `json:"paymenttype"`
	Modeofpayment   string    `json:"modeofpayment"`
	Paymentstatus   string    `json:"paymentstatus"`
	Paymentdatetime time.Time `json:"paymentdatetime"`
	Paidamount      float64   `json:"paidamount"`
	Paymentid       int       `json:"paymentid"`
}
