package domain

import (
	"time"
)

//null.string worked fine comparitevly : "github.com/volatiletech/null"

type Pickupmain struct {
	Pickuprequestid           int       `json:"pickuprequestid"`
	Customerid                string    `json:"customerid"`
	Pickupdroptype            string    `json:"pickupdroptype"`
	Pickuplocation            string    `json:"pickuplocation"`
	Droplocation              string    `json:"droplocation"`
	Pickupscheduleslot        string    `json:"pickupscheduleslot"`
	Pickupscheduledate        time.Time `json:"pickupscheduledate"`
	Actualpickupdatetime      time.Time `json:"actualpickupdatetime"`
	Pickupagentid             int       `json:"pickupagentid"`
	Pickupfacilityid          string    `json:"pickupfacilityid"`
	Pickupstatus              string    `json:"pickupstatus"`
	Paymentstatus             string    `json:"paymentstatus"`
	Createddatetime           time.Time `json:"createddatetime"`
	Pickupaddress             string    `json:"pickupaddress"`
	Domesticforeignidentifier string    `json:"domesticforeignidentifier"`
	Pickuplong                string    `json:"pickuplong"`
	Pickuplat                 string    `json:"pickuplat"`
	Modifieddatetime          time.Time `json:"modifieddatetime"`
	Pickuprequestedpincode    string    `json:"pickuprequestedpincode"`
	//Customername              *string   `json:"customername"`
	//Customername         sql.NullString `json:"customername"`
	//Customername         pgtype.Text `json:"customername"`
	//Customername null.String `json:"customername"`
	Customername         string    `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
}
