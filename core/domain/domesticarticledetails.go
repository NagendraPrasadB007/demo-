package domain

import "time"

type Domesticarticledetails struct {
	Domid                     int       `json:"domid"`
	Pickuprequestid           int       `json:"pickuprequestid"`
	Articleid                 string    `json:"articleid"`
	Articlestate              string    `json:"articlestate"`
	Articletype               string    `json:"articletype"`
	Articlecontent            string    `json:"articlecontent"`
	Articleimageid            int       `json:"articleimageid"`
	Articlepickupcharges      float64   `json:"articlepickupcharges"`
	Ispremailing              bool      `json:"ispremailing"`
	Isparcelpacking           bool      `json:"isparcelpacking"`
	Createddatetime           time.Time `json:"createddatetime"`
	Modifieddatetime          time.Time `json:"modifieddatetime"`
	Customerdacpickup         string    `json:"customerdacpickup"`
	Addresstype               string    `json:"addresstype"`
	Bkgtransactionid          string    `json:"bkgtransactionid"`
	Originpin                 int       `json:"originpin"`
	Destinationpin            int       `json:"destinationpin"`
	Physicalweight            float64   `json:"physicalweight"`
	Shape                     string    `json:"shape"`
	Dimensionlength           float64   `json:"dimensionlength"`
	Dimensionbreadth          float64   `json:"dimensionbreadth"`
	Dimensionheight           float64   `json:"dimensionheight"`
	Volumetricweight          float64   `json:"volumetricweight"`
	Chargedweight             float64   `json:"chargedweight"`
	Mailservicetypecode       string    `json:"mailservicetypecode"`
	Bkgtype                   string    `json:"bkgtype"`
	Mailform                  string    `json:"mailform"`
	Isprepaid                 bool      `json:"isprepaid"`
	Prepaymenttype            string    `json:"prepaymenttype"`
	Valueofprepayment         float64   `json:"valueofprepayment"`
	Vpcodflag                 bool      `json:"vpcodflag"`
	Valueforvpcod             float64   `json:"valueforvpcod"`
	Insuranceflag             bool      `json:"insuranceflag"`
	Insurancetype             string    `json:"insurancetype"`
	Valueinsurance            float64   `json:"valueinsurance"`
	Acknowledgementpod        bool      `json:"acknowledgementpod"`
	Instructionsrts           string    `json:"instructionsrts"`
	Addressrefsender          string    `json:"addressrefsender"`
	Addressrefreceiver        string    `json:"addressrefreceiver"`
	Addressrefsenderaltaddr   string    `json:"addressrefsenderaltaddr"`
	Addressrefreceiveraktaddr string    `json:"addressrefreceiveraktaddr"`
	Barcodenumber             string    `json:"barcodenumber"`
	Pickupflag                bool      `json:"pickupflag"`
	Basetariff                float64   `json:"basetariff"`
	Tax                       float64   `json:"tax"`
	Totaltariff               float64   `json:"totaltariff"`
	Modeofpayment             string    `json:"modeofpayment"`
	Paymenttranid             string    `json:"paymenttranid"`
	Status                    string    `json:"status"`
	Createdon                 time.Time `json:"createdon"`
	Createdby                 string    `json:"createdby"`
	Updatedon                 time.Time `json:"updatedon"`
	Updatedby                 string    `json:"updatedby"`
	Authorisedon              time.Time `json:"authorisedon"`
	Authorisedby              string    `json:"authorisedby"`
	Facilityid                string    `json:"facilityid"`
	Reqipaddress              string    `json:"reqipaddress"`
	Bookingchannel            string    `json:"bookingchannel"`
	Customerid                string    `json:"customerid"`
	Contractnumber            string    `json:"contractnumber"`
	Isparcel                  bool      `json:"isparcel"`
	Iscod                     bool      `json:"iscod"`
}
