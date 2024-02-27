package domain

import "time"

type SubPiecedetails struct {
	Subid                               int       `json:"subid"`
	Intlid                              int       `json:"intlid"`
	Identifierpieceid                   int       `json:"identifierpieceid"`
	Subpiececatproductcode              string    `json:"subpiececatproductcode"`
	Hscode                              string    `json:"hscode"`
	Productcustomstariffhead            string    `json:"productcustomstariffhead"`
	Productdescription                  string    `json:"productdescription"`
	Isocodefororigincountry             string    `json:"isocodefororigincountry"`
	Unitforsubpiecequantity             string    `json:"unitforsubpiecequantity"`
	Subpiecequantitycount               string    `json:"subpiecequantitycount"`
	Producttotalvalueasperinvoice       string    `json:"producttotalvalueasperinvoice"`
	Isocodeforcurrency                  string    `json:"isocodeforcurrency"`
	Subpieceweight                      string    `json:"subpieceweight"`
	Subpieceweightnett                  string    `json:"subpieceweightnett"`
	Productinvoicenumber                string    `json:"productinvoicenumber"`
	Productinvoicedate                  time.Time `json:"productinvoicedate"`
	Statusforecommerce                  string    `json:"statusforecommerce"`
	Urlforecommerceconsignment          string    `json:"urlforecommerceconsignment"`
	Ecommercepaymenttransactionid       string    `json:"ecommercepaymenttransactionid"`
	Ecommerceskuno                      string    `json:"ecommerceskuno"`
	Taxinvoicenumber                    string    `json:"taxinvoicenumber"`
	Taxinvoicedate                      time.Time `json:"taxinvoicedate"`
	Serialnumberforsubpieceintaxinvoice string    `json:"serialnumberforsubpieceintaxinvoice"`
	Valueofsubpieceaspertaxinvoice      float64   `json:"valueofsubpieceaspertaxinvoice"`
	Assessablefreeonboardvalue          float64   `json:"assessablefreeonboardvalue"`
	Isocodeforassessablecurrency        string    `json:"isocodeforassessablecurrency"`
	Exchangerateforasblcurr             float64   `json:"exchangerateforasblcurr"`
	Assessableamount                    float64   `json:"assessableamount"`
	Rateforexportduty                   float64   `json:"rateforexportduty"`
	Exportdutyamount                    float64   `json:"exportdutyamount"`
	Rateforcess                         float64   `json:"rateforcess"`
	Cessamount                          float64   `json:"cessamount"`
	Igstrate                            float64   `json:"igstrate"`
	Igstamount                          float64   `json:"igstamount"`
	Compensationrate                    float64   `json:"compensationrate"`
	Compensationamount                  float64   `json:"compensationamount"`
	Detailsofletterofundertakingorbond  bool      `json:"detailsofletterofundertakingorbond"`
	Modeofpayment                       string    `json:"modeofpayment"`
	Paymenttransactionid                string    `json:"paymenttransactionid"`
	Createdon                           time.Time `json:"createdon"`
	Createdby                           string    `json:"createdby"`
	Updatedon                           time.Time `json:"updatedon"`
	Updatedby                           string    `json:"updatedby"`
	Authorisedon                        time.Time `json:"authorisedon"`
	Authorisedby                        string    `json:"authorisedby"`
	Facilityid                          string    `json:"facilityid"`
	Ipaddress                           string    `json:"ipaddress"`
	Bookingchanneltype                  string    `json:"bookingchanneltype"`
}
