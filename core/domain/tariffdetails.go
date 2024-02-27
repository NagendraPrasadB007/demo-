package domain

type Tariffdetails struct {
	Pickuprequestid     int     `json:"pickuprequestid"`
	Articleid           string  `json:"articleid"`
	Totalamount         float64 `json:"totalamount"`
	Pickupcharges       float64 `json:"pickupcharges"`
	Registrationfee     float64 `json:"registrationfee"`
	Postage             float64 `json:"postage"`
	Ackorpodfee         float64 `json:"ackorpodfee"`
	Valueinsurance      float64 `json:"valueinsurance"`
	Valueforvpcod       float64 `json:"valueforvpcod"`
	Doordeliverycharges float64 `json:"doordeliverycharges"`
	Packingfee          float64 `json:"packingfee"`
	Cgst                float64 `json:"cgst"`
	Sgst                float64 `json:"sgst"`
	Othercharges        float64 `json:"othercharges"`
	Tariffid            int     `json:"tariffid"`
}
