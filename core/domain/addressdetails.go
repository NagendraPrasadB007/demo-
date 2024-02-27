package domain

type Addressdetails struct {
	Customerid   string `json:"customerid" update:"customer_id"`
	Addressid    int    `json:"addressid"`
	Firstname    string `json:"firstname" update:"firstname"`
	Lastname     string `json:"lastname" update:"lastname"`
	Addressline1 string `json:"addressline1" update:"addressline1"`
	Addressline2 string `json:"addressline2" update:"addressline2"`
	Landmark     string `json:"landmark" update:"landmark"`
	City         string `json:"city" update:"city"`
	State        string `json:"state" update:"state"`
	Country      string `json:"country" update:"country"`
	Pincode      string `json:"pincode" update:"pincode"`
	Mobilenumber string `json:"mobilenumber" update:"mobilenumber"`
	Emailid      string `json:"emailid" update:"email_id"`
	Geocode      string `json:"geocode" update:"geo_code"`
	Addresstype  string `json:"addresstype" update:"address_type"`
	Fromtopickup string `json:"fromtopickup" update:"fromtopickup"`
	Isverified   bool   `json:"isverified" update:"is_verified"`
}

/*
type Addressdetails struct {
	Customerid   string `json:"customerid"`
	Addressid    int    `json:"addressid"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Addressline1 string `json:"addressline1"`
	Addressline2 string `json:"addressline2"`
	Landmark     string `json:"landmark"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode"`
	Mobilenumber string `json:"mobilenumber"`
	Emailid      string `json:"emailid"`
	Geocode      string `json:"geocode"`
	Addresstype  string `json:"addresstype"`
	Fromtopickup string `json:"fromtopickup"`
	Isverified   bool   `json:"isverified"`
}
*/
