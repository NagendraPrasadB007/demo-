package handler

import (
	"net/http"
	"pickupmanagement/core/domain"
	"pickupmanagement/core/port"
	repository "pickupmanagement/repo"
	"time"

	"github.com/gin-gonic/gin"
)

// response represents a response body format
type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// errorResponse represents an error response body format
type errorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}

// meta represents metadata for a paginated response
type meta struct {
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip uint64) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

//////////////////////////////////// Image

type ImageResponse struct {
	Imageid         int    `json:"imageid"`
	Filename        string `json:"filename"`
	Size            int    `json:"size"`
	Mimetype        string `json:"mimetype"`
	Filepath        string `json:"filePath"`
	Pickuprequestid int    `json:"pickuprequestid"`
}

// Response function to upload the image of open article
func newImageResponse(image *domain.Image, filePath string) ImageResponse {
	return ImageResponse{
		Imageid:         repository.GeneratedImageID,
		Filename:        image.Filename,
		Size:            image.Size,
		Mimetype:        image.Mimetype,
		Filepath:        filePath,
		Pickuprequestid: image.Pickuprequestid,
	}
}

// Response function to get the image of the pickuprequest based on imageid
func newImageDetailsResponse(image *domain.Image) ImageResponse {
	return ImageResponse{
		Imageid:  image.Imageid,
		Filename: image.Filename,
		Size:     image.Size,
		Mimetype: image.Mimetype,
		Filepath: image.Filepath,
	}
}

////////////////////////////////////

type dashboardResponse struct {
	Imageid         int    `json:"imageid"`
	Filename        string `json:"filename"`
	Size            int    `json:"size"`
	Mimetype        string `json:"mimetype"`
	Filepath        string `json:"filePath"`
	Pickuprequestid int    `json:"pickuprequestid"`
}

//////////////////////////////////////////////pickupmain post api

type PickupmainResponse struct {
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
	// Customername              null.String `json:"customername"`
	Customername         string    `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
}

// response function to create a new request in pickupmain table
func newPickupmainResponse(pickupmain *domain.Pickupmain) PickupmainResponse {
	return PickupmainResponse{
		//Pickuprequestid: pickupmain.Pickuprequestid,
		Pickuprequestid:           repository.GeneratedPickupRequestID,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		Customername:              pickupmain.Customername,
		Customermobilenumber:      pickupmain.Customermobilenumber,
		Assigneddatetime:          pickupmain.Assigneddatetime,
	}
}

//ORIGINAL
// func newRaisedPickupRequestResponseDom(pickupmain *domain.Pickupmain, domestic *domain.Domesticarticledetails, international *domain.Internationalarticledetails, subpiece *domain.SubPiecedetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseDom {

// (&pickupmain, &domestic, &international, &tariff, &payment)
func newRaisedPickupRequestResponseInt(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, subpiece *domain.SubPiecedetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseInt {
	return RaisedPickupRequestResponseInt{

		// func newRaisedPickupRequestResponseInt(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseInt {
		// 	return RaisedPickupRequestResponseInt{
		//PICKUPMAIN//
		Pickuprequestid:           repository.SerialPickupRequestID,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		//INTERNATIONAL //
		Intlid: repository.SerialIntlID,
		//Pickuprequestid:      international.Pickupagentid,
		Articleid:            international.Articleid,
		Articlestate:         international.Articlestate,
		Articletype:          international.Articletype,
		Articlecontent:       international.Articlecontent,
		Articleimageid:       international.Articleimageid,
		Articlepickupcharges: international.Articlepickupcharges,
		Ispremailing:         international.Ispremailing,
		Isparcelpacking:      international.Isparcelpacking,
		//Createddatetime:           international.Createddatetime,
		//Modifieddatetime:          international.Modifieddatetime,
		Customerdacpickup:         international.Customerdacpickup,
		Addresstype:               international.Addresstype,
		Bkgtransactionid:          international.Bkgtransactionid,
		Origincountrycode:         international.Origincountrycode,
		Destinationcountrycode:    international.Destinationcountrycode,
		Physicalweight:            international.Physicalweight,
		Mailclass:                 international.Mailclass,
		Contenttype:               international.Contenttype,
		Shape:                     international.Shape,
		Dimensionlength:           international.Dimensionlength,
		Dimensionbreadth:          international.Dimensionbreadth,
		Dimensionheight:           international.Dimensionheight,
		Volumetricweight:          international.Volumetricweight,
		Chargedweight:             international.Chargedweight,
		Mailservicetypecode:       international.Mailservicetypecode,
		Bkgtype:                   international.Bkgtype,
		Mailform:                  international.Mailform,
		Isprepaid:                 international.Isprepaid,
		Prepaymenttype:            international.Prepaymenttype,
		Valueofprepayment:         international.Valueofprepayment,
		Vpcodflag:                 international.Vpcodflag,
		Valueforvpcod:             international.Valueforvpcod,
		Insuranceflag:             international.Insuranceflag,
		Insurancetype:             international.Insurancetype,
		Valueinsurance:            international.Valueinsurance,
		Acknowledgementpod:        international.Acknowledgementpod,
		Instructionsrts:           international.Instructionsrts,
		Addressrefsender:          international.Addressrefsender,
		Addressrefreceiver:        international.Addressrefreceiver,
		Addressrefsenderaltaddr:   international.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: international.Addressrefreceiveraktaddr,
		Barcodenumber:             international.Barcodenumber,
		Pickupflag:                international.Pickupflag,
		Basetariff:                international.Basetariff,
		Tax:                       international.Tax,
		Totaltariff:               international.Totaltariff,
		Modeofpayment:             international.Modeofpayment,
		Paymenttranid:             international.Paymenttranid,
		Status:                    international.Status,
		Createdon:                 international.Createdon,
		Createdby:                 international.Createdby,
		Updatedon:                 international.Updatedon,
		Updatedby:                 international.Updatedby,
		Authorisedon:              international.Authorisedon,
		Authorisedby:              international.Authorisedby,
		Facilityid:                international.Facilityid,
		Reqipaddress:              international.Reqipaddress,
		Bookingchannel:            international.Bookingchannel,
		Consignmentvalue:          international.Consignmentvalue,
		Mailexporttype:            international.Mailexporttype,
		Pbefilingtype:             international.Pbefilingtype,
		Declaration1:              international.Declaration1,
		Declaration23:             international.Declaration23,
		Declaration4:              international.Declaration4,
		Selffilingcusbroker:       international.Selffilingcusbroker,
		Cusbrokerlicno:            international.Cusbrokerlicno,
		Cusbrokername:             international.Cusbrokername,
		Cusbrokeraddress:          international.Cusbrokeraddress,
		//Customerid:          international.Customerid,
		Contractnumber: international.Contractnumber,
		Gstn:           international.Gstn,
		Ibccode:        international.Ibccode,
		Lut:            international.Lut,
		Adcode:         international.Adcode,
		Isparcel:       international.Isparcel,
		Iscod:          international.Iscod,
		//SUBPIECE//
		//Intlid:                              repository.SerialIntlID,
		Identifierpieceid:                   subpiece.Identifierpieceid,
		Subpiececatproductcode:              subpiece.Subpiececatproductcode,
		Hscode:                              subpiece.Hscode,
		Productcustomstariffhead:            subpiece.Productcustomstariffhead,
		Productdescription:                  subpiece.Productdescription,
		Isocodefororigincountry:             subpiece.Isocodefororigincountry,
		Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
		Subpiecequantitycount:               subpiece.Subpiecequantitycount,
		Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
		Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
		Subpieceweight:                      subpiece.Subpieceweight,
		Subpieceweightnett:                  subpiece.Subpieceweightnett,
		Productinvoicenumber:                subpiece.Productinvoicenumber,
		Productinvoicedate:                  subpiece.Productinvoicedate,
		Statusforecommerce:                  subpiece.Statusforecommerce,
		Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
		Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
		Ecommerceskuno:                      subpiece.Ecommerceskuno,
		Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
		Taxinvoicedate:                      subpiece.Taxinvoicedate,
		Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
		Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
		Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
		Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
		Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
		Assessableamount:                    subpiece.Assessableamount,
		Rateforexportduty:                   subpiece.Rateforexportduty,
		Exportdutyamount:                    subpiece.Exportdutyamount,
		Rateforcess:                         subpiece.Rateforcess,
		Cessamount:                          subpiece.Cessamount,
		Igstrate:                            subpiece.Igstrate,
		Igstamount:                          subpiece.Igstamount,
		Compensationrate:                    subpiece.Compensationrate,
		Compensationamount:                  subpiece.Compensationamount,
		Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
		//Modeofpayment:                       subpiece.Modeofpayment,
		Paymenttransactionid: subpiece.Paymenttransactionid,
		//Createdon:                           subpiece.Createdon,
		//Createdby:                           subpiece.Createdby,
		//Updatedon:                           subpiece.Updatedon,
		//Updatedby:                           subpiece.Updatedby,
		//Authorisedon:                        subpiece.Authorisedon,
		//Authorisedby:                        subpiece.Authorisedby,
		//Facilityid:                          subpiece.Facilityid,
		Ipaddress:          subpiece.Ipaddress,
		Bookingchanneltype: subpiece.Bookingchanneltype,
		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
	}
}

type SubPieceResponse struct {
	//Subid                               int       `json:"subid"`
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

/////////////////////////// raise pickuprequest ( international )

type RaisedPickupRequestResponseInt1 struct {
	//pickupmain
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
	// Customername              null.String `json:"customername"`
	Customername         string    `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
	Intlid               int       `json:"intlid"`
	//Pickuprequestid           int       `json:"pickuprequestid"`
	Articleid            string  `json:"articleid"`
	Articlestate         string  `json:"articlestate"`
	Articletype          string  `json:"articletype"`
	Articlecontent       string  `json:"articlecontent"`
	Articleimageid       int     `json:"articleimageid"`
	Articlepickupcharges float64 `json:"articlepickupcharges"`
	Ispremailing         bool    `json:"ispremailing"`
	Isparcelpacking      bool    `json:"isparcelpacking"`
	//Createddatetime           time.Time `json:"createddatetime"`
	//Modifieddatetime          time.Time `json:"modifieddatetime"`
	Customerdacpickup         string    `json:"customerdacpickup"`
	Addresstype               string    `json:"addresstype"`
	Bkgtransactionid          string    `json:"bkgtransactionid"`
	Origincountrycode         int       `json:"origincountrycode"`
	Destinationcountrycode    int       `json:"destinationcountrycode"`
	Physicalweight            float64   `json:"physicalweight"`
	Mailclass                 string    `json:"mailclass"`
	Contenttype               string    `json:"contenttype"`
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
	Consignmentvalue          int       `json:"consignmentvalue"`
	Mailexporttype            string    `json:"mailexporttype"`
	Pbefilingtype             string    `json:"pbefilingtype"`
	Declaration1              string    `json:"declaration1"`
	Declaration23             string    `json:"declaration23"`
	Declaration4              string    `json:"declaration4"`
	Selffilingcusbroker       string    `json:"selffilingcusbroker"`
	Cusbrokerlicno            string    `json:"cusbrokerlicno"`
	Cusbrokername             string    `json:"cusbrokername"`
	Cusbrokeraddress          string    `json:"cusbrokeraddress"`
	//Customerid          string `json:"customerid"`
	Contractnumber string             `json:"contractnumber"`
	Gstn           string             `json:"gstn"`
	Ibccode        string             `json:"ibccode"`
	Lut            string             `json:"lut"`
	Adcode         string             `json:"adcode"`
	Isparcel       bool               `json:"isparcel"`
	Iscod          bool               `json:"iscod"`
	SubPieces      []SubPieceResponse `json:"subpieces"`
	/*
		//SUBPEICE//
		Subid int `json:"subid"`
		//Intlid                              int       `json:"intlid"`
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
		//Modeofpayment                       string    `json:"modeofpayment"`
		Paymenttransactionid string `json:"paymenttransactionid"`
		//Createdon                           time.Time `json:"createdon"`
		//Createdby                           string    `json:"createdby"`
		//Updatedon                           time.Time `json:"updatedon"`
		//Updatedby                           string    `json:"updatedby"`
		//Authorisedon                        time.Time `json:"authorisedon"`
		//Authorisedby                        string    `json:"authorisedby"`
		//Facilityid                          string    `json:"facilityid"`
		Ipaddress          string `json:"ipaddress"`
		Bookingchanneltype string `json:"bookingchanneltype"`
	*/
	//TARIFF//
	//Pickuprequestid     int     `json:"pickuprequestid"`
	//Articleid           string  `json:"articleid"`
	Totalamount     float64 `json:"totalamount"`
	Pickupcharges   float64 `json:"pickupcharges"`
	Registrationfee float64 `json:"registrationfee"`
	Postage         float64 `json:"postage"`
	Ackorpodfee     float64 `json:"ackorpodfee"`
	//Valueinsurance      float64 `json:"valueinsurance"`
	//Valueforvpcod       float64 `json:"valueforvpcod"`
	Doordeliverycharges float64 `json:"doordeliverycharges"`
	Packingfee          float64 `json:"packingfee"`
	Cgst                float64 `json:"cgst"`
	Sgst                float64 `json:"sgst"`
	Othercharges        float64 `json:"othercharges"`
	Tariffid            int     `json:"tariffid"`
	//PAYMENT
	//Paymenttranid   string    `json:"paymenttranid"`
	//Pickuprequestid int       `json:"pickuprequestid"`
	//Articleid       string    `json:"articleid"`
	Paymenttype string `json:"paymenttype"`
	//Modeofpayment   string    `json:"modeofpayment"`
	//Paymentstatus   string    `json:"paymentstatus"`
	Paymentdatetime time.Time `json:"paymentdatetime"`
	Paidamount      float64   `json:"paidamount"`
	Paymentid       int       `json:"paymentid"`
}

// response function to raise pickuprequest ( international )
func newRaisedPickupRequestResponseInt1(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseInt1 {
	return RaisedPickupRequestResponseInt1{
		//PICKUPMAIN//
		Pickuprequestid:           repository.SerialPickupRequestID,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		Customername:              pickupmain.Customername,
		Customermobilenumber:      pickupmain.Customermobilenumber,
		Assigneddatetime:          pickupmain.Assigneddatetime,
		//INTERNATIONAL //
		Intlid: repository.SerialIntlID,
		//Pickuprequestid:      international.Pickupagentid,
		Articleid:            international.Articleid,
		Articlestate:         international.Articlestate,
		Articletype:          international.Articletype,
		Articlecontent:       international.Articlecontent,
		Articleimageid:       international.Articleimageid,
		Articlepickupcharges: international.Articlepickupcharges,
		Ispremailing:         international.Ispremailing,
		Isparcelpacking:      international.Isparcelpacking,
		//Createddatetime:           international.Createddatetime,
		//Modifieddatetime:          international.Modifieddatetime,
		Customerdacpickup:         international.Customerdacpickup,
		Addresstype:               international.Addresstype,
		Bkgtransactionid:          international.Bkgtransactionid,
		Origincountrycode:         international.Origincountrycode,
		Destinationcountrycode:    international.Destinationcountrycode,
		Physicalweight:            international.Physicalweight,
		Mailclass:                 international.Mailclass,
		Contenttype:               international.Contenttype,
		Shape:                     international.Shape,
		Dimensionlength:           international.Dimensionlength,
		Dimensionbreadth:          international.Dimensionbreadth,
		Dimensionheight:           international.Dimensionheight,
		Volumetricweight:          international.Volumetricweight,
		Chargedweight:             international.Chargedweight,
		Mailservicetypecode:       international.Mailservicetypecode,
		Bkgtype:                   international.Bkgtype,
		Mailform:                  international.Mailform,
		Isprepaid:                 international.Isprepaid,
		Prepaymenttype:            international.Prepaymenttype,
		Valueofprepayment:         international.Valueofprepayment,
		Vpcodflag:                 international.Vpcodflag,
		Valueforvpcod:             international.Valueforvpcod,
		Insuranceflag:             international.Insuranceflag,
		Insurancetype:             international.Insurancetype,
		Valueinsurance:            international.Valueinsurance,
		Acknowledgementpod:        international.Acknowledgementpod,
		Instructionsrts:           international.Instructionsrts,
		Addressrefsender:          international.Addressrefsender,
		Addressrefreceiver:        international.Addressrefreceiver,
		Addressrefsenderaltaddr:   international.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: international.Addressrefreceiveraktaddr,
		Barcodenumber:             international.Barcodenumber,
		Pickupflag:                international.Pickupflag,
		Basetariff:                international.Basetariff,
		Tax:                       international.Tax,
		Totaltariff:               international.Totaltariff,
		Modeofpayment:             international.Modeofpayment,
		Paymenttranid:             international.Paymenttranid,
		Status:                    international.Status,
		Createdon:                 international.Createdon,
		Createdby:                 international.Createdby,
		Updatedon:                 international.Updatedon,
		Updatedby:                 international.Updatedby,
		Authorisedon:              international.Authorisedon,
		Authorisedby:              international.Authorisedby,
		Facilityid:                international.Facilityid,
		Reqipaddress:              international.Reqipaddress,
		Bookingchannel:            international.Bookingchannel,
		Consignmentvalue:          international.Consignmentvalue,
		Mailexporttype:            international.Mailexporttype,
		Pbefilingtype:             international.Pbefilingtype,
		Declaration1:              international.Declaration1,
		Declaration23:             international.Declaration23,
		Declaration4:              international.Declaration4,
		Selffilingcusbroker:       international.Selffilingcusbroker,
		Cusbrokerlicno:            international.Cusbrokerlicno,
		Cusbrokername:             international.Cusbrokername,
		Cusbrokeraddress:          international.Cusbrokeraddress,
		//Customerid:          international.Customerid,
		Contractnumber: international.Contractnumber,
		Gstn:           international.Gstn,
		Ibccode:        international.Ibccode,
		Lut:            international.Lut,
		Adcode:         international.Adcode,
		Isparcel:       international.Isparcel,
		Iscod:          international.Iscod,
		SubPieces:      newSubpieceOrderresponse(international.SubPieces),

		/*
			//SUBPIECE//
			//Intlid:                              repository.SerialIntlID,
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
			//Modeofpayment:                       subpiece.Modeofpayment,
			Paymenttransactionid: subpiece.Paymenttransactionid,
			//Createdon:                           subpiece.Createdon,
			//Createdby:                           subpiece.Createdby,
			//Updatedon:                           subpiece.Updatedon,
			//Updatedby:                           subpiece.Updatedby,
			//Authorisedon:                        subpiece.Authorisedon,
			//Authorisedby:                        subpiece.Authorisedby,
			//Facilityid:                          subpiece.Facilityid,
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
		*/

		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		Tariffid:            tariff.Tariffid,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
		Paymentid:       payment.Paymentid,
	}
}

// raise pickuprequest ( international )

func newSubpieceOrderresponse(subpieces []domain.SubPiecedetails) []SubPieceResponse {
	var subPieceResponses []SubPieceResponse

	for _, subpiece := range subpieces {
		// subID, err := repository.GetNextSubID()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//below line for getting subid in response
		//for _, subID := range repository.SerialSubIDs {

		subPieceResponse := SubPieceResponse{
			//below line for getting subid in response
			//Subid:                               repository.SerialSubID,
			//Subid:                               subID,
			Intlid:                              repository.SerialIntlID,
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
			//Modeofpayment:                       subpiece.Modeofpayment,
			Paymenttransactionid: subpiece.Paymenttransactionid,
			//Createdon:                           subpiece.Createdon,
			//Createdby:                           subpiece.Createdby,
			//Updatedon:                           subpiece.Updatedon,
			//Updatedby:                           subpiece.Updatedby,
			//Authorisedon:                        subpiece.Authorisedon,
			//Authorisedby:                        subpiece.Authorisedby,
			//Facilityid:                          subpiece.Facilityid,
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
		}

		subPieceResponses = append(subPieceResponses, subPieceResponse)
		//}
	}
	return subPieceResponses
}

// Response function to fetch pickuprequest details based on pickuprequestid
func newFetchPickupRequestResponseInt1(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseInt1 {
	return RaisedPickupRequestResponseInt1{
		//PICKUPMAIN//
		Pickuprequestid:           pickupmain.Pickuprequestid,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		Customername:              pickupmain.Customername,
		Customermobilenumber:      pickupmain.Customermobilenumber,
		Assigneddatetime:          pickupmain.Assigneddatetime,
		//INTERNATIONAL //
		Intlid: international.Intlid,
		//Pickuprequestid:      international.Pickupagentid,
		Articleid:            international.Articleid,
		Articlestate:         international.Articlestate,
		Articletype:          international.Articletype,
		Articlecontent:       international.Articlecontent,
		Articleimageid:       international.Articleimageid,
		Articlepickupcharges: international.Articlepickupcharges,
		Ispremailing:         international.Ispremailing,
		Isparcelpacking:      international.Isparcelpacking,
		//Createddatetime:           international.Createddatetime,
		//Modifieddatetime:          international.Modifieddatetime,
		Customerdacpickup:         international.Customerdacpickup,
		Addresstype:               international.Addresstype,
		Bkgtransactionid:          international.Bkgtransactionid,
		Origincountrycode:         international.Origincountrycode,
		Destinationcountrycode:    international.Destinationcountrycode,
		Physicalweight:            international.Physicalweight,
		Mailclass:                 international.Mailclass,
		Contenttype:               international.Contenttype,
		Shape:                     international.Shape,
		Dimensionlength:           international.Dimensionlength,
		Dimensionbreadth:          international.Dimensionbreadth,
		Dimensionheight:           international.Dimensionheight,
		Volumetricweight:          international.Volumetricweight,
		Chargedweight:             international.Chargedweight,
		Mailservicetypecode:       international.Mailservicetypecode,
		Bkgtype:                   international.Bkgtype,
		Mailform:                  international.Mailform,
		Isprepaid:                 international.Isprepaid,
		Prepaymenttype:            international.Prepaymenttype,
		Valueofprepayment:         international.Valueofprepayment,
		Vpcodflag:                 international.Vpcodflag,
		Valueforvpcod:             international.Valueforvpcod,
		Insuranceflag:             international.Insuranceflag,
		Insurancetype:             international.Insurancetype,
		Valueinsurance:            international.Valueinsurance,
		Acknowledgementpod:        international.Acknowledgementpod,
		Instructionsrts:           international.Instructionsrts,
		Addressrefsender:          international.Addressrefsender,
		Addressrefreceiver:        international.Addressrefreceiver,
		Addressrefsenderaltaddr:   international.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: international.Addressrefreceiveraktaddr,
		Barcodenumber:             international.Barcodenumber,
		Pickupflag:                international.Pickupflag,
		Basetariff:                international.Basetariff,
		Tax:                       international.Tax,
		Totaltariff:               international.Totaltariff,
		Modeofpayment:             international.Modeofpayment,
		Paymenttranid:             international.Paymenttranid,
		Status:                    international.Status,
		Createdon:                 international.Createdon,
		Createdby:                 international.Createdby,
		Updatedon:                 international.Updatedon,
		Updatedby:                 international.Updatedby,
		Authorisedon:              international.Authorisedon,
		Authorisedby:              international.Authorisedby,
		Facilityid:                international.Facilityid,
		Reqipaddress:              international.Reqipaddress,
		Bookingchannel:            international.Bookingchannel,
		Consignmentvalue:          international.Consignmentvalue,
		Mailexporttype:            international.Mailexporttype,
		Pbefilingtype:             international.Pbefilingtype,
		Declaration1:              international.Declaration1,
		Declaration23:             international.Declaration23,
		Declaration4:              international.Declaration4,
		Selffilingcusbroker:       international.Selffilingcusbroker,
		Cusbrokerlicno:            international.Cusbrokerlicno,
		Cusbrokername:             international.Cusbrokername,
		Cusbrokeraddress:          international.Cusbrokeraddress,
		//Customerid:          international.Customerid,
		Contractnumber: international.Contractnumber,
		Gstn:           international.Gstn,
		Ibccode:        international.Ibccode,
		Lut:            international.Lut,
		Adcode:         international.Adcode,
		Isparcel:       international.Isparcel,
		Iscod:          international.Iscod,
		//SubPieces:      newfetchSubpieceOrderresponse(international.SubPieces),
		SubPieces: newfetchSubpieceOrderresponse(international.SubPieces),

		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		Tariffid:            tariff.Tariffid,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
		Paymentid:       payment.Paymentid,
	}
}

func newfetchSubpieceOrderresponse(subpieces []domain.SubPiecedetails) []SubPieceResponse {
	var subPieceResponses []SubPieceResponse

	for _, subpiece := range subpieces {

		subPieceResponse := SubPieceResponse{
			//below line for getting subid in response
			//Subid:                               repository.SerialSubID,
			//Subid:                               subpiece.Subid,
			Intlid:                              subpiece.Intlid,
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
			//Modeofpayment:                       subpiece.Modeofpayment,
			Paymenttransactionid: subpiece.Paymenttransactionid,
			//Createdon:                           subpiece.Createdon,
			//Createdby:                           subpiece.Createdby,
			//Updatedon:                           subpiece.Updatedon,
			//Updatedby:                           subpiece.Updatedby,
			//Authorisedon:                        subpiece.Authorisedon,
			//Authorisedby:                        subpiece.Authorisedby,
			//Facilityid:                          subpiece.Facilityid,
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
		}

		subPieceResponses = append(subPieceResponses, subPieceResponse)
		//}
	}
	return subPieceResponses
}

type SubPieceResponseFetch struct {
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

/*
func newSubpieceOrderresponse(inter []domain.Internationalarticledetails) []SubPieceResponse {
	var orderProductResponses []SubPieceResponse

	for _, subpiece := range inter {
		orderProductResponses = append(orderProductResponses, SubPieceResponse{})
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
			//Modeofpayment:                       subpiece.Modeofpayment,
			Paymenttransactionid: subpiece.Paymenttransactionid,
			//Createdon:                           subpiece.Createdon,
			//Createdby:                           subpiece.Createdby,
			//Updatedon:                           subpiece.Updatedon,
			//Updatedby:                           subpiece.Updatedby,
			//Authorisedon:                        subpiece.Authorisedon,
			//Authorisedby:                        subpiece.Authorisedby,
			//Facilityid:                          subpiece.Facilityid,
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
	}

	return SubPieceResponse
}
*/

////////////////////////////////////////// Raise pickuprequest

// Response function to raise a pickuprequest
func newRaisedPickupRequestResponseDom(pickupmain *domain.Pickupmain, domestic *domain.Domesticarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseDom {
	return RaisedPickupRequestResponseDom{
		//PICKUPMAIN//
		Pickuprequestid:           repository.SerialPickupRequestID,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		Customername:              pickupmain.Customername,
		Customermobilenumber:      pickupmain.Customermobilenumber,
		Assigneddatetime:          pickupmain.Assigneddatetime,

		//DOMESTIC//
		Domid:                repository.SerialDomID,
		Articleid:            domestic.Articleid,
		Articlestate:         domestic.Articlestate,
		Articletype:          domestic.Articletype,
		Articlecontent:       domestic.Articlecontent,
		Articleimageid:       domestic.Articleimageid,
		Articlepickupcharges: domestic.Articlepickupcharges,
		Ispremailing:         domestic.Ispremailing,
		Isparcelpacking:      domestic.Isparcelpacking,
		//Createddatetime:           domestic.Createddatetime,
		//Modifieddatetime:          domestic.Modifieddatetime,
		Customerdacpickup:         domestic.Customerdacpickup,
		Addresstype:               domestic.Addresstype,
		Bkgtransactionid:          domestic.Bkgtransactionid,
		Originpin:                 domestic.Originpin,
		Destinationpin:            domestic.Destinationpin,
		Physicalweight:            domestic.Physicalweight,
		Shape:                     domestic.Shape,
		Dimensionlength:           domestic.Dimensionlength,
		Dimensionbreadth:          domestic.Dimensionbreadth,
		Dimensionheight:           domestic.Dimensionheight,
		Volumetricweight:          domestic.Volumetricweight,
		Chargedweight:             domestic.Chargedweight,
		Mailservicetypecode:       domestic.Mailservicetypecode,
		Bkgtype:                   domestic.Bkgtype,
		Mailform:                  domestic.Mailform,
		Isprepaid:                 domestic.Isprepaid,
		Prepaymenttype:            domestic.Prepaymenttype,
		Valueofprepayment:         domestic.Valueofprepayment,
		Vpcodflag:                 domestic.Vpcodflag,
		Valueforvpcod:             domestic.Valueforvpcod,
		Insuranceflag:             domestic.Insuranceflag,
		Insurancetype:             domestic.Insurancetype,
		Valueinsurance:            domestic.Valueinsurance,
		Acknowledgementpod:        domestic.Acknowledgementpod,
		Instructionsrts:           domestic.Instructionsrts,
		Addressrefsender:          domestic.Addressrefsender,
		Addressrefreceiver:        domestic.Addressrefreceiver,
		Addressrefsenderaltaddr:   domestic.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: domestic.Addressrefreceiveraktaddr,
		Barcodenumber:             domestic.Barcodenumber,
		Pickupflag:                domestic.Pickupflag,
		Basetariff:                domestic.Basetariff,
		Tax:                       domestic.Tax,
		Totaltariff:               domestic.Totaltariff,
		Modeofpayment:             domestic.Modeofpayment,
		Paymenttranid:             domestic.Paymenttranid,
		Status:                    domestic.Status,
		Createdon:                 domestic.Createdon,
		Createdby:                 domestic.Createdby,
		Updatedon:                 domestic.Updatedon,
		Updatedby:                 domestic.Updatedby,
		Authorisedon:              domestic.Authorisedon,
		Authorisedby:              domestic.Authorisedby,
		Facilityid:                domestic.Facilityid,
		Reqipaddress:              domestic.Reqipaddress,
		Bookingchannel:            domestic.Bookingchannel,
		//Customerid:                domestic.Customerid,
		Contractnumber: domestic.Contractnumber,
		Isparcel:       domestic.Isparcel,
		Iscod:          domestic.Iscod,
		//INTERNATIONAL //
		/*
			Intlid: repository.SerialIntlID,
		*/
		//Pickuprequestid:      international.Pickupagentid,
		//Articleid:            international.Articleid,
		//Articlestate:         international.Articlestate,
		//Articletype:          international.Articletype,
		//Articlecontent:       international.Articlecontent,
		//Articleimageid:       international.Articleimageid,
		//Articlepickupcharges: international.Articlepickupcharges,
		//Ispremailing:         international.Ispremailing,
		//Isparcelpacking:      international.Isparcelpacking,
		//Createddatetime:           international.Createddatetime,
		//Modifieddatetime:          international.Modifieddatetime,
		//Customerdacpickup:         international.Customerdacpickup,
		//Addresstype:               international.Addresstype,
		//Bkgtransactionid:          international.Bkgtransactionid,
		/*
			Origincountrycode:      international.Origincountrycode,
			Destinationcountrycode: international.Destinationcountrycode,
		*/
		//Physicalweight:            international.Physicalweight,
		/*
			Mailclass:   international.Mailclass,
			Contenttype: international.Contenttype,
		*/
		//Shape:                     international.Shape,
		//Dimensionlength:           international.Dimensionlength,
		//Dimensionbreadth:          international.Dimensionbreadth,
		//Dimensionheight:           international.Dimensionheight,
		//Volumetricweight:          international.Volumetricweight,
		//Chargedweight:             international.Chargedweight,
		//Mailservicetypecode:       international.Mailservicetypecode,
		//Bkgtype:                   international.Bkgtype,
		//Mailform:                  international.Mailform,
		//Isprepaid:                 international.Isprepaid,
		//Prepaymenttype:            international.Prepaymenttype,
		//Valueofprepayment:         international.Valueofprepayment,
		//Vpcodflag:                 international.Vpcodflag,
		//Valueforvpcod:             international.Valueforvpcod,
		//Insuranceflag:             international.Insuranceflag,
		//Insurancetype:             international.Insurancetype,
		//Valueinsurance:            international.Valueinsurance,
		//Acknowledgementpod:        international.Acknowledgementpod,
		//Instructionsrts:           international.Instructionsrts,
		//Addressrefsender:          international.Addressrefsender,
		//Addressrefreceiver:        international.Addressrefreceiver,
		//Addressrefsenderaltaddr:   international.Addressrefsenderaltaddr,
		//Addressrefreceiveraktaddr: international.Addressrefreceiveraktaddr,
		//Barcodenumber:             international.Barcodenumber,
		//Pickupflag:                international.Pickupflag,
		//Basetariff:                international.Basetariff,
		//Tax:                       international.Tax,
		//Totaltariff:               international.Totaltariff,
		//Modeofpayment:             international.Modeofpayment,
		//Paymenttranid:             international.Paymenttranid,
		//Status:                    international.Status,
		//Createdon:                 international.Createdon,
		//Createdby:                 international.Createdby,
		//Updatedon:                 international.Updatedon,
		//Updatedby:                 international.Updatedby,
		//Authorisedon:              international.Authorisedon,
		//Authorisedby:              international.Authorisedby,
		//Facilityid:                international.Facilityid,
		//Reqipaddress: international.Reqipaddress,
		//Bookingchannel:            international.Bookingchannel,
		/*
			Consignmentvalue:    international.Consignmentvalue,
			Mailexporttype:      international.Mailexporttype,
			Pbefilingtype:       international.Pbefilingtype,
			Declaration1:        international.Declaration1,
			Declaration23:       international.Declaration23,
			Declaration4:        international.Declaration4,
			Selffilingcusbroker: international.Selffilingcusbroker,
			Cusbrokerlicno:      international.Cusbrokerlicno,
			Cusbrokername:       international.Cusbrokername,
			Cusbrokeraddress:    international.Cusbrokeraddress,
		*/
		//Customerid:          international.Customerid,
		//Contractnumber:      international.Contractnumber,
		/*
			Gstn:    international.Gstn,
			Ibccode: international.Ibccode,
			Lut:     international.Lut,
			Adcode:  international.Adcode,
		*/
		//Isparcel: international.Isparcel,
		//Iscod:    international.Iscod,
		//SUBPIECE//
		//Intlid:                              repository.SerialIntlID,
		/*
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
		*/
		//Modeofpayment:                       subpiece.Modeofpayment,
		/*
			Paymenttransactionid: subpiece.Paymenttransactionid,
		*/
		//Createdon:                           subpiece.Createdon,
		//Createdby:                           subpiece.Createdby,
		//Updatedon:                           subpiece.Updatedon,
		//Updatedby:                           subpiece.Updatedby,
		//Authorisedon:                        subpiece.Authorisedon,
		//Authorisedby:                        subpiece.Authorisedby,
		//Facilityid:                          subpiece.Facilityid,
		/*
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
		*/
		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		Tariffid:            tariff.Tariffid,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
		Paymentid:       payment.Paymentid,
	}
}

type RaisedPickupRequestResponseDom struct {
	//pickupmain
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
	Customername              string    `json:"customername"`
	// Customername              null.String `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
	//DOMESTIC//
	Domid int `json:"domid"`
	//Pickuprequestid           int       `json:"pickuprequestid"`
	Articleid            string  `json:"articleid"`
	Articlestate         string  `json:"articlestate"`
	Articletype          string  `json:"articletype"`
	Articlecontent       string  `json:"articlecontent"`
	Articleimageid       int     `json:"articleimageid"`
	Articlepickupcharges float64 `json:"articlepickupcharges"`
	Ispremailing         bool    `json:"ispremailing"`
	Isparcelpacking      bool    `json:"isparcelpacking"`
	//Createddatetime           time.Time `json:"createddatetime"`
	//Modifieddatetime          time.Time `json:"modifieddatetime"`
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
	//Customerid                string    `json:"customerid"`
	Contractnumber string `json:"contractnumber"`
	Isparcel       bool   `json:"isparcel"`
	Iscod          bool   `json:"iscod"`
	//TARIFF//
	//Pickuprequestid     int     `json:"pickuprequestid"`
	//Articleid           string  `json:"articleid"`
	Totalamount     float64 `json:"totalamount"`
	Pickupcharges   float64 `json:"pickupcharges"`
	Registrationfee float64 `json:"registrationfee"`
	Postage         float64 `json:"postage"`
	Ackorpodfee     float64 `json:"ackorpodfee"`
	//Valueinsurance      float64 `json:"valueinsurance"`
	//Valueforvpcod       float64 `json:"valueforvpcod"`
	Doordeliverycharges float64 `json:"doordeliverycharges"`
	Packingfee          float64 `json:"packingfee"`
	Cgst                float64 `json:"cgst"`
	Sgst                float64 `json:"sgst"`
	Othercharges        float64 `json:"othercharges"`
	Tariffid            int     `json:"tariffid"`
	//PAYMENT
	//Paymenttranid   string    `json:"paymenttranid"`
	//Pickuprequestid int       `json:"pickuprequestid"`
	//Articleid       string    `json:"articleid"`
	Paymenttype string `json:"paymenttype"`
	//Modeofpayment   string    `json:"modeofpayment"`
	//Paymentstatus   string    `json:"paymentstatus"`
	Paymentdatetime time.Time `json:"paymentdatetime"`
	Paidamount      float64   `json:"paidamount"`
	Paymentid       int       `json:"paymentid"`
}

type RaisedPickupRequestResponseInt struct {
	//pickupmain
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
	Intlid                    int       `json:"intlid"`
	//Pickuprequestid           int       `json:"pickuprequestid"`
	Articleid            string  `json:"articleid"`
	Articlestate         string  `json:"articlestate"`
	Articletype          string  `json:"articletype"`
	Articlecontent       string  `json:"articlecontent"`
	Articleimageid       int     `json:"articleimageid"`
	Articlepickupcharges float64 `json:"articlepickupcharges"`
	Ispremailing         bool    `json:"ispremailing"`
	Isparcelpacking      bool    `json:"isparcelpacking"`
	//Createddatetime           time.Time `json:"createddatetime"`
	//Modifieddatetime          time.Time `json:"modifieddatetime"`
	Customerdacpickup         string    `json:"customerdacpickup"`
	Addresstype               string    `json:"addresstype"`
	Bkgtransactionid          string    `json:"bkgtransactionid"`
	Origincountrycode         int       `json:"origincountrycode"`
	Destinationcountrycode    int       `json:"destinationcountrycode"`
	Physicalweight            float64   `json:"physicalweight"`
	Mailclass                 string    `json:"mailclass"`
	Contenttype               string    `json:"contenttype"`
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
	Consignmentvalue          int       `json:"consignmentvalue"`
	Mailexporttype            string    `json:"mailexporttype"`
	Pbefilingtype             string    `json:"pbefilingtype"`
	Declaration1              string    `json:"declaration1"`
	Declaration23             string    `json:"declaration23"`
	Declaration4              string    `json:"declaration4"`
	Selffilingcusbroker       string    `json:"selffilingcusbroker"`
	Cusbrokerlicno            string    `json:"cusbrokerlicno"`
	Cusbrokername             string    `json:"cusbrokername"`
	Cusbrokeraddress          string    `json:"cusbrokeraddress"`
	//Customerid          string `json:"customerid"`
	Contractnumber string `json:"contractnumber"`
	Gstn           string `json:"gstn"`
	Ibccode        string `json:"ibccode"`
	Lut            string `json:"lut"`
	Adcode         string `json:"adcode"`
	Isparcel       bool   `json:"isparcel"`
	Iscod          bool   `json:"iscod"`
	//SUBPEICE//
	Subid int `json:"subid"`
	//Intlid                              int       `json:"intlid"`
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
	//Modeofpayment                       string    `json:"modeofpayment"`
	Paymenttransactionid string `json:"paymenttransactionid"`
	//Createdon                           time.Time `json:"createdon"`
	//Createdby                           string    `json:"createdby"`
	//Updatedon                           time.Time `json:"updatedon"`
	//Updatedby                           string    `json:"updatedby"`
	//Authorisedon                        time.Time `json:"authorisedon"`
	//Authorisedby                        string    `json:"authorisedby"`
	//Facilityid                          string    `json:"facilityid"`
	Ipaddress          string `json:"ipaddress"`
	Bookingchanneltype string `json:"bookingchanneltype"`
	//TARIFF//
	//Pickuprequestid     int     `json:"pickuprequestid"`
	//Articleid           string  `json:"articleid"`
	Totalamount     float64 `json:"totalamount"`
	Pickupcharges   float64 `json:"pickupcharges"`
	Registrationfee float64 `json:"registrationfee"`
	Postage         float64 `json:"postage"`
	Ackorpodfee     float64 `json:"ackorpodfee"`
	//Valueinsurance      float64 `json:"valueinsurance"`
	//Valueforvpcod       float64 `json:"valueforvpcod"`
	Doordeliverycharges float64 `json:"doordeliverycharges"`
	Packingfee          float64 `json:"packingfee"`
	Cgst                float64 `json:"cgst"`
	Sgst                float64 `json:"sgst"`
	Othercharges        float64 `json:"othercharges"`
	//PAYMENT
	//Paymenttranid   string    `json:"paymenttranid"`
	//Pickuprequestid int       `json:"pickuprequestid"`
	//Articleid       string    `json:"articleid"`
	Paymenttype string `json:"paymenttype"`
	//Modeofpayment   string    `json:"modeofpayment"`
	//Paymentstatus   string    `json:"paymentstatus"`
	Paymentdatetime time.Time `json:"paymentdatetime"`
	Paidamount      float64   `json:"paidamount"`
}

/*
func newBulkPickupmainResponse(pickupmains []domain.Pickupmain) []pickupmainResponse {
	var responses []pickupmainResponse

	for _, pickupmain := range pickupmains {
		response := pickupmainResponse{
			Pickuprequestid: repository.GeneratedPickupRequestID,
			Customerid:      pickupmain.Customerid,
			Pickupdroptype:            pickupmain.Pickupdroptype,
			Pickuplocation:            pickupmain.Pickuplocation,
			Droplocation:              pickupmain.Droplocation,
			Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
			Pickupscheduledate:        pickupmain.Pickupscheduledate,
			Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
			Pickupagentid:             pickupmain.Pickupagentid,
			Pickupfacilityid:          pickupmain.Pickupfacilityid,
			Pickupstatus:              pickupmain.Pickupstatus,
			Paymentstatus:             pickupmain.Paymentstatus,
			Createddatetime:           pickupmain.Createddatetime,
			Pickupaddress:             pickupmain.Pickupaddress,
			Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
			Pickuplong:                pickupmain.Pickuplong,
			Pickuplat:                 pickupmain.Pickuplat,
			Modifieddatetime:          pickupmain.Modifieddatetime,
		}

		responses = append(responses, response)
	}

	return responses
}
*/

/////////////////////////////////////////////////////////////////////////

type AddressdetailsResponse struct {
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

// update api response
func newAddressdetailsResponse(addressdetails *domain.Addressdetails) AddressdetailsResponse {
	return AddressdetailsResponse{
		Customerid:   addressdetails.Customerid,
		Addressid:    addressdetails.Addressid,
		Firstname:    addressdetails.Firstname,
		Lastname:     addressdetails.Lastname,
		Addressline1: addressdetails.Addressline1,
		Addressline2: addressdetails.Addressline2,
		Landmark:     addressdetails.Landmark,
		City:         addressdetails.City,
		State:        addressdetails.State,
		Country:      addressdetails.Country,
		Pincode:      addressdetails.Pincode,
		Mobilenumber: addressdetails.Mobilenumber,
		Emailid:      addressdetails.Emailid,
		Geocode:      addressdetails.Geocode,
		Addresstype:  addressdetails.Addresstype,
		Fromtopickup: addressdetails.Fromtopickup,
		Isverified:   addressdetails.Isverified,
	}
}

// create address api response
func newCreatedAddressdetailsResponse(addressdetails *domain.Addressdetails) AddressdetailsResponse {
	return AddressdetailsResponse{
		Customerid:   addressdetails.Customerid,
		Addressid:    repository.GeneratedAddressID,
		Firstname:    addressdetails.Firstname,
		Lastname:     addressdetails.Lastname,
		Addressline1: addressdetails.Addressline1,
		Addressline2: addressdetails.Addressline2,
		Landmark:     addressdetails.Landmark,
		City:         addressdetails.City,
		State:        addressdetails.State,
		Country:      addressdetails.Country,
		Pincode:      addressdetails.Pincode,
		Mobilenumber: addressdetails.Mobilenumber,
		Emailid:      addressdetails.Emailid,
		Geocode:      addressdetails.Geocode,
		Addresstype:  addressdetails.Addresstype,
		Fromtopickup: addressdetails.Fromtopickup,
		Isverified:   addressdetails.Isverified,
	}
}

//////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////// fetch pickuprequest

// response function to fetch pickuprequest details based on pickuprequestid
func newFetchPickupRequestResponse(pickupmain *domain.Pickupmain, domestic *domain.Domesticarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseDom {
	return RaisedPickupRequestResponseDom{
		Pickuprequestid:           pickupmain.Pickuprequestid,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		Customername:              pickupmain.Customername,
		Customermobilenumber:      pickupmain.Customermobilenumber,
		Assigneddatetime:          pickupmain.Assigneddatetime,
		//DOMESTIC//
		Domid:                domestic.Domid,
		Articleid:            domestic.Articleid,
		Articlestate:         domestic.Articlestate,
		Articletype:          domestic.Articletype,
		Articlecontent:       domestic.Articlecontent,
		Articleimageid:       domestic.Articleimageid,
		Articlepickupcharges: domestic.Articlepickupcharges,
		Ispremailing:         domestic.Ispremailing,
		Isparcelpacking:      domestic.Isparcelpacking,
		//Createddatetime:           domestic.Createddatetime,
		//Modifieddatetime:          domestic.Modifieddatetime,
		Customerdacpickup:         domestic.Customerdacpickup,
		Addresstype:               domestic.Addresstype,
		Bkgtransactionid:          domestic.Bkgtransactionid,
		Originpin:                 domestic.Originpin,
		Destinationpin:            domestic.Destinationpin,
		Physicalweight:            domestic.Physicalweight,
		Shape:                     domestic.Shape,
		Dimensionlength:           domestic.Dimensionlength,
		Dimensionbreadth:          domestic.Dimensionbreadth,
		Dimensionheight:           domestic.Dimensionheight,
		Volumetricweight:          domestic.Volumetricweight,
		Chargedweight:             domestic.Chargedweight,
		Mailservicetypecode:       domestic.Mailservicetypecode,
		Bkgtype:                   domestic.Bkgtype,
		Mailform:                  domestic.Mailform,
		Isprepaid:                 domestic.Isprepaid,
		Prepaymenttype:            domestic.Prepaymenttype,
		Valueofprepayment:         domestic.Valueofprepayment,
		Vpcodflag:                 domestic.Vpcodflag,
		Valueforvpcod:             domestic.Valueforvpcod,
		Insuranceflag:             domestic.Insuranceflag,
		Insurancetype:             domestic.Insurancetype,
		Valueinsurance:            domestic.Valueinsurance,
		Acknowledgementpod:        domestic.Acknowledgementpod,
		Instructionsrts:           domestic.Instructionsrts,
		Addressrefsender:          domestic.Addressrefsender,
		Addressrefreceiver:        domestic.Addressrefreceiver,
		Addressrefsenderaltaddr:   domestic.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: domestic.Addressrefreceiveraktaddr,
		Barcodenumber:             domestic.Barcodenumber,
		Pickupflag:                domestic.Pickupflag,
		Basetariff:                domestic.Basetariff,
		Tax:                       domestic.Tax,
		Totaltariff:               domestic.Totaltariff,
		Modeofpayment:             domestic.Modeofpayment,
		Paymenttranid:             domestic.Paymenttranid,
		Status:                    domestic.Status,
		Createdon:                 domestic.Createdon,
		Createdby:                 domestic.Createdby,
		Updatedon:                 domestic.Updatedon,
		Updatedby:                 domestic.Updatedby,
		Authorisedon:              domestic.Authorisedon,
		Authorisedby:              domestic.Authorisedby,
		Facilityid:                domestic.Facilityid,
		Reqipaddress:              domestic.Reqipaddress,
		Bookingchannel:            domestic.Bookingchannel,
		//Customerid:                domestic.Customerid,
		Contractnumber: domestic.Contractnumber,
		Isparcel:       domestic.Isparcel,
		Iscod:          domestic.Iscod,
		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		Tariffid:            tariff.Tariffid,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
		Paymentid:       payment.Paymentid,
	}
}

// includes subpiece
func newIntFetchPickupRequestResponse(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, subpiece []domain.SubPiecedetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseInt {
	return RaisedPickupRequestResponseInt{
		Pickuprequestid:           pickupmain.Pickuprequestid,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		//INTERNATIONAL //
		Intlid: international.Intlid,
		//Pickuprequestid:      international.Pickupagentid,
		Articleid:            international.Articleid,
		Articlestate:         international.Articlestate,
		Articletype:          international.Articletype,
		Articlecontent:       international.Articlecontent,
		Articleimageid:       international.Articleimageid,
		Articlepickupcharges: international.Articlepickupcharges,
		Ispremailing:         international.Ispremailing,
		Isparcelpacking:      international.Isparcelpacking,
		//Createddatetime:           international.Createddatetime,
		//Modifieddatetime:          international.Modifieddatetime,
		Customerdacpickup:         international.Customerdacpickup,
		Addresstype:               international.Addresstype,
		Bkgtransactionid:          international.Bkgtransactionid,
		Origincountrycode:         international.Origincountrycode,
		Destinationcountrycode:    international.Destinationcountrycode,
		Physicalweight:            international.Physicalweight,
		Mailclass:                 international.Mailclass,
		Contenttype:               international.Contenttype,
		Shape:                     international.Shape,
		Dimensionlength:           international.Dimensionlength,
		Dimensionbreadth:          international.Dimensionbreadth,
		Dimensionheight:           international.Dimensionheight,
		Volumetricweight:          international.Volumetricweight,
		Chargedweight:             international.Chargedweight,
		Mailservicetypecode:       international.Mailservicetypecode,
		Bkgtype:                   international.Bkgtype,
		Mailform:                  international.Mailform,
		Isprepaid:                 international.Isprepaid,
		Prepaymenttype:            international.Prepaymenttype,
		Valueofprepayment:         international.Valueofprepayment,
		Vpcodflag:                 international.Vpcodflag,
		Valueforvpcod:             international.Valueforvpcod,
		Insuranceflag:             international.Insuranceflag,
		Insurancetype:             international.Insurancetype,
		Valueinsurance:            international.Valueinsurance,
		Acknowledgementpod:        international.Acknowledgementpod,
		Instructionsrts:           international.Instructionsrts,
		Addressrefsender:          international.Addressrefsender,
		Addressrefreceiver:        international.Addressrefreceiver,
		Addressrefsenderaltaddr:   international.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: international.Addressrefreceiveraktaddr,
		Barcodenumber:             international.Barcodenumber,
		Pickupflag:                international.Pickupflag,
		Basetariff:                international.Basetariff,
		Tax:                       international.Tax,
		Totaltariff:               international.Totaltariff,
		Modeofpayment:             international.Modeofpayment,
		Paymenttranid:             international.Paymenttranid,
		Status:                    international.Status,
		Createdon:                 international.Createdon,
		Createdby:                 international.Createdby,
		Updatedon:                 international.Updatedon,
		Updatedby:                 international.Updatedby,
		Authorisedon:              international.Authorisedon,
		Authorisedby:              international.Authorisedby,
		Facilityid:                international.Facilityid,
		Reqipaddress:              international.Reqipaddress,
		Bookingchannel:            international.Bookingchannel,
		Consignmentvalue:          international.Consignmentvalue,
		Mailexporttype:            international.Mailexporttype,
		Pbefilingtype:             international.Pbefilingtype,
		Declaration1:              international.Declaration1,
		Declaration23:             international.Declaration23,
		Declaration4:              international.Declaration4,
		Selffilingcusbroker:       international.Selffilingcusbroker,
		Cusbrokerlicno:            international.Cusbrokerlicno,
		Cusbrokername:             international.Cusbrokername,
		Cusbrokeraddress:          international.Cusbrokeraddress,
		//Customerid:          international.Customerid,
		Contractnumber: international.Contractnumber,
		Gstn:           international.Gstn,
		Ibccode:        international.Ibccode,
		Lut:            international.Lut,
		Adcode:         international.Adcode,
		Isparcel:       international.Isparcel,
		Iscod:          international.Iscod,
		/*
			//SubPieces:      subpiece,
			//SUBPIECE//
			//Intlid:                              repository.SerialIntlID,
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
			//Modeofpayment:                       subpiece.Modeofpayment,
			Paymenttransactionid: subpiece.Paymenttransactionid,
			//Createdon:                           subpiece.Createdon,
			//Createdby:                           subpiece.Createdby,
			//Updatedon:                           subpiece.Updatedon,
			//Updatedby:                           subpiece.Updatedby,
			//Authorisedon:                        subpiece.Authorisedon,
			//Authorisedby:                        subpiece.Authorisedby,
			//Facilityid:                          subpiece.Facilityid,
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
		*/
		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
	}
}

func newIntFetchPickupRequestResponse1(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) RaisedPickupRequestResponseInt {
	return RaisedPickupRequestResponseInt{
		Pickuprequestid:           pickupmain.Pickuprequestid,
		Customerid:                pickupmain.Customerid,
		Pickupdroptype:            pickupmain.Pickupdroptype,
		Pickuplocation:            pickupmain.Pickuplocation,
		Droplocation:              pickupmain.Droplocation,
		Pickupscheduleslot:        pickupmain.Pickupscheduleslot,
		Pickupscheduledate:        pickupmain.Pickupscheduledate,
		Actualpickupdatetime:      pickupmain.Actualpickupdatetime,
		Pickupagentid:             pickupmain.Pickupagentid,
		Pickupfacilityid:          pickupmain.Pickupfacilityid,
		Pickupstatus:              pickupmain.Pickupstatus,
		Paymentstatus:             pickupmain.Paymentstatus,
		Createddatetime:           pickupmain.Createddatetime,
		Pickupaddress:             pickupmain.Pickupaddress,
		Domesticforeignidentifier: pickupmain.Domesticforeignidentifier,
		Pickuplong:                pickupmain.Pickuplong,
		Pickuplat:                 pickupmain.Pickuplat,
		Modifieddatetime:          pickupmain.Modifieddatetime,
		Pickuprequestedpincode:    pickupmain.Pickuprequestedpincode,
		//INTERNATIONAL //
		Intlid: international.Intlid,
		//Pickuprequestid:      international.Pickupagentid,
		Articleid:            international.Articleid,
		Articlestate:         international.Articlestate,
		Articletype:          international.Articletype,
		Articlecontent:       international.Articlecontent,
		Articleimageid:       international.Articleimageid,
		Articlepickupcharges: international.Articlepickupcharges,
		Ispremailing:         international.Ispremailing,
		Isparcelpacking:      international.Isparcelpacking,
		//Createddatetime:           international.Createddatetime,
		//Modifieddatetime:          international.Modifieddatetime,
		Customerdacpickup:         international.Customerdacpickup,
		Addresstype:               international.Addresstype,
		Bkgtransactionid:          international.Bkgtransactionid,
		Origincountrycode:         international.Origincountrycode,
		Destinationcountrycode:    international.Destinationcountrycode,
		Physicalweight:            international.Physicalweight,
		Mailclass:                 international.Mailclass,
		Contenttype:               international.Contenttype,
		Shape:                     international.Shape,
		Dimensionlength:           international.Dimensionlength,
		Dimensionbreadth:          international.Dimensionbreadth,
		Dimensionheight:           international.Dimensionheight,
		Volumetricweight:          international.Volumetricweight,
		Chargedweight:             international.Chargedweight,
		Mailservicetypecode:       international.Mailservicetypecode,
		Bkgtype:                   international.Bkgtype,
		Mailform:                  international.Mailform,
		Isprepaid:                 international.Isprepaid,
		Prepaymenttype:            international.Prepaymenttype,
		Valueofprepayment:         international.Valueofprepayment,
		Vpcodflag:                 international.Vpcodflag,
		Valueforvpcod:             international.Valueforvpcod,
		Insuranceflag:             international.Insuranceflag,
		Insurancetype:             international.Insurancetype,
		Valueinsurance:            international.Valueinsurance,
		Acknowledgementpod:        international.Acknowledgementpod,
		Instructionsrts:           international.Instructionsrts,
		Addressrefsender:          international.Addressrefsender,
		Addressrefreceiver:        international.Addressrefreceiver,
		Addressrefsenderaltaddr:   international.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: international.Addressrefreceiveraktaddr,
		Barcodenumber:             international.Barcodenumber,
		Pickupflag:                international.Pickupflag,
		Basetariff:                international.Basetariff,
		Tax:                       international.Tax,
		Totaltariff:               international.Totaltariff,
		Modeofpayment:             international.Modeofpayment,
		Paymenttranid:             international.Paymenttranid,
		Status:                    international.Status,
		Createdon:                 international.Createdon,
		Createdby:                 international.Createdby,
		Updatedon:                 international.Updatedon,
		Updatedby:                 international.Updatedby,
		Authorisedon:              international.Authorisedon,
		Authorisedby:              international.Authorisedby,
		Facilityid:                international.Facilityid,
		Reqipaddress:              international.Reqipaddress,
		Bookingchannel:            international.Bookingchannel,
		Consignmentvalue:          international.Consignmentvalue,
		Mailexporttype:            international.Mailexporttype,
		Pbefilingtype:             international.Pbefilingtype,
		Declaration1:              international.Declaration1,
		Declaration23:             international.Declaration23,
		Declaration4:              international.Declaration4,
		Selffilingcusbroker:       international.Selffilingcusbroker,
		Cusbrokerlicno:            international.Cusbrokerlicno,
		Cusbrokername:             international.Cusbrokername,
		Cusbrokeraddress:          international.Cusbrokeraddress,
		//Customerid:          international.Customerid,
		Contractnumber: international.Contractnumber,
		Gstn:           international.Gstn,
		Ibccode:        international.Ibccode,
		Lut:            international.Lut,
		Adcode:         international.Adcode,
		Isparcel:       international.Isparcel,
		Iscod:          international.Iscod,

		/*
			//SUBPIECE//
			//Intlid:                              repository.SerialIntlID,
			Identifierpieceid:                   subpiece.Identifierpieceid,
			Subpiececatproductcode:              subpiece.Subpiececatproductcode,
			Hscode:                              subpiece.Hscode,
			Productcustomstariffhead:            subpiece.Productcustomstariffhead,
			Productdescription:                  subpiece.Productdescription,
			Isocodefororigincountry:             subpiece.Isocodefororigincountry,
			Unitforsubpiecequantity:             subpiece.Unitforsubpiecequantity,
			Subpiecequantitycount:               subpiece.Subpiecequantitycount,
			Producttotalvalueasperinvoice:       subpiece.Producttotalvalueasperinvoice,
			Isocodeforcurrency:                  subpiece.Isocodeforcurrency,
			Subpieceweight:                      subpiece.Subpieceweight,
			Subpieceweightnett:                  subpiece.Subpieceweightnett,
			Productinvoicenumber:                subpiece.Productinvoicenumber,
			Productinvoicedate:                  subpiece.Productinvoicedate,
			Statusforecommerce:                  subpiece.Statusforecommerce,
			Urlforecommerceconsignment:          subpiece.Urlforecommerceconsignment,
			Ecommercepaymenttransactionid:       subpiece.Ecommercepaymenttransactionid,
			Ecommerceskuno:                      subpiece.Ecommerceskuno,
			Taxinvoicenumber:                    subpiece.Taxinvoicenumber,
			Taxinvoicedate:                      subpiece.Taxinvoicedate,
			Serialnumberforsubpieceintaxinvoice: subpiece.Serialnumberforsubpieceintaxinvoice,
			Valueofsubpieceaspertaxinvoice:      subpiece.Valueofsubpieceaspertaxinvoice,
			Assessablefreeonboardvalue:          subpiece.Assessablefreeonboardvalue,
			Isocodeforassessablecurrency:        subpiece.Isocodeforassessablecurrency,
			Exchangerateforasblcurr:             subpiece.Exchangerateforasblcurr,
			Assessableamount:                    subpiece.Assessableamount,
			Rateforexportduty:                   subpiece.Rateforexportduty,
			Exportdutyamount:                    subpiece.Exportdutyamount,
			Rateforcess:                         subpiece.Rateforcess,
			Cessamount:                          subpiece.Cessamount,
			Igstrate:                            subpiece.Igstrate,
			Igstamount:                          subpiece.Igstamount,
			Compensationrate:                    subpiece.Compensationrate,
			Compensationamount:                  subpiece.Compensationamount,
			Detailsofletterofundertakingorbond:  subpiece.Detailsofletterofundertakingorbond,
			//Modeofpayment:                       subpiece.Modeofpayment,
			Paymenttransactionid: subpiece.Paymenttransactionid,
			//Createdon:                           subpiece.Createdon,
			//Createdby:                           subpiece.Createdby,
			//Updatedon:                           subpiece.Updatedon,
			//Updatedby:                           subpiece.Updatedby,
			//Authorisedon:                        subpiece.Authorisedon,
			//Authorisedby:                        subpiece.Authorisedby,
			//Facilityid:                          subpiece.Facilityid,
			Ipaddress:          subpiece.Ipaddress,
			Bookingchanneltype: subpiece.Bookingchanneltype,
		*/
		//TARIFF
		//Pickuptariffuestid:     tariff.Pickuptariffuestid,
		//Articleid:           tariff.Articleid,
		Totalamount:     tariff.Totalamount,
		Pickupcharges:   tariff.Pickupcharges,
		Registrationfee: tariff.Registrationfee,
		Postage:         tariff.Postage,
		Ackorpodfee:     tariff.Ackorpodfee,
		//Valueinsurance:      tariff.Valueinsurance,
		//Valueforvpcod:       tariff.Valueforvpcod,
		Doordeliverycharges: tariff.Doordeliverycharges,
		Packingfee:          tariff.Packingfee,
		Cgst:                tariff.Cgst,
		Sgst:                tariff.Sgst,
		Othercharges:        tariff.Othercharges,
		//PAYMENT
		//Paymenttranid:   payment.Paymenttranid,
		//Pickuppaymentuestid: payment.Pickuppaymentuestid,
		//Articleid:       payment.Articleid,
		Paymenttype: payment.Paymenttype,
		//Modeofpayment:   payment.Modeofpayment,
		//Paymentstatus:   payment.Paymentstatus,
		Paymentdatetime: payment.Paymentdatetime,
		Paidamount:      payment.Paidamount,
	}
}

////////////////////////////Pickupschedueslots

type pickupscheduleslotsResponse struct {
	Pickupscheduleslotid int    `json:"pickupscheduleslotid"`
	Scheduleslots        string `json:"scheduleslot"`
}

func newPickupscheduleslotsResponse(pickupscheduleslots *domain.Pickupscheduleslots) pickupscheduleslotsResponse {
	return pickupscheduleslotsResponse{
		Pickupscheduleslotid: pickupscheduleslots.Pickupscheduleslotid,
		Scheduleslots:        pickupscheduleslots.Scheduleslots,
	}
}

//////////////////////////////////Pickupschedueslots

// /////////////////////////////// Unassigned api
type UnassignedDom struct {
	Pickuprequestid    int       `json:"pickuprequestid"`
	Pickupscheduleslot string    `json:"pickupscheduleslot"`
	Pickupscheduledate time.Time `json:"pickupscheduledate"`
	Pickupaddress      string    `json:"pickupaddress"`
	// Customername         null.String `json:"customername"`
	Customername         string  `json:"customername"`
	Customermobilenumber string  `json:"customermobilenumber"`
	Physicalweight       float64 `json:"physicalweight"`
	Volumetricweight     float64 `json:"volumetricweight"`
}

// Response function to get unassigned pickuprequestid details ( Domestic )
func newFetchPickupRequestResponseDom(pickupmain *domain.Pickupmain, domestic *domain.Domesticarticledetails) UnassignedDom {
	return UnassignedDom{
		Pickuprequestid:      pickupmain.Pickuprequestid,
		Pickupscheduleslot:   pickupmain.Pickupscheduleslot,
		Pickupscheduledate:   pickupmain.Pickupscheduledate,
		Pickupaddress:        pickupmain.Pickupaddress,
		Customername:         pickupmain.Customername,
		Customermobilenumber: pickupmain.Customermobilenumber,
		Physicalweight:       domestic.Physicalweight,
		Volumetricweight:     domestic.Volumetricweight,
	}
}

type UnassignedInt struct {
	Pickuprequestid    int       `json:"pickuprequestid"`
	Pickupscheduleslot string    `json:"pickupscheduleslot"`
	Pickupscheduledate time.Time `json:"pickupscheduledate"`
	Pickupaddress      string    `json:"pickupaddress"`
	// Customername         null.String `json:"customername"`
	Customername         string  `json:"customername"`
	Customermobilenumber string  `json:"customermobilenumber"`
	Physicalweight       float64 `json:"physicalweight"`
	Volumetricweight     float64 `json:"volumetricweight"`
}

// Response function to get unassigned pickuprequestid details ( International )
func newFetchPickupRequestResponseInt(pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails) UnassignedInt {
	return UnassignedInt{
		Pickuprequestid:      pickupmain.Pickuprequestid,
		Pickupscheduleslot:   pickupmain.Pickupscheduleslot,
		Pickupscheduledate:   pickupmain.Pickupscheduledate,
		Pickupaddress:        pickupmain.Pickupaddress,
		Customername:         pickupmain.Customername,
		Customermobilenumber: pickupmain.Customermobilenumber,
		Physicalweight:       international.Physicalweight,
		Volumetricweight:     international.Volumetricweight,
	}
}

///////////////////////////////// Unassigned api

// ////////////////////////////// pickuprequest based on facilityid
type PickuprequestBasicfac struct {
	Pickuprequestid    int       `json:"pickuprequestid"`
	Customerid         string    `json:"customerid"`
	Pickupscheduleslot string    `json:"pickupscheduleslot"`
	Pickupscheduledate time.Time `json:"pickupscheduledate"`
	Pickupagentid      int       `json:"pickupagentid"`
	Pickupaddress      string    `json:"pickupaddress"`
	Pickupstatus       string    `json:"pickupstatus"`
	Customername       string    `json:"customername"`
	// Customername         null.String `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Createddatetime      time.Time `json:"createddatetime"`
	Modifieddatetime     time.Time `json:"modifieddatetime"`
}

// Response function to display pickuprequest based on facilityid
func newPickuprequestBasicfac(pickupmain *domain.Pickupmain) PickuprequestBasicfac {
	return PickuprequestBasicfac{
		Pickuprequestid:      pickupmain.Pickuprequestid,
		Customerid:           pickupmain.Customerid,
		Pickupscheduleslot:   pickupmain.Pickupscheduleslot,
		Pickupscheduledate:   pickupmain.Pickupscheduledate,
		Pickupagentid:        pickupmain.Pickupagentid,
		Pickupaddress:        pickupmain.Pickupaddress,
		Pickupstatus:         pickupmain.Pickupstatus,
		Customername:         pickupmain.Customername,
		Customermobilenumber: pickupmain.Customermobilenumber,
		Createddatetime:      pickupmain.Createddatetime,
		Modifieddatetime:     pickupmain.Modifieddatetime,
	}
}

type PickuprequestBasiccus struct {
	Pickuprequestid    int       `json:"pickuprequestid"`
	Customerid         string    `json:"customerid"`
	Pickupscheduleslot string    `json:"pickupscheduleslot"`
	Pickupscheduledate time.Time `json:"pickupscheduledate"`
	Pickupagentid      int       `json:"pickupagentid"`
	Pickupaddress      string    `json:"pickupaddress"`
	Pickupstatus       string    `json:"pickupstatus"`
	Customername       string    `json:"customername"`
	// Customername         null.String `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Createddatetime      time.Time `json:"createddatetime"`
	Modifieddatetime     time.Time `json:"modifieddatetime"`
}

// Response function to display pickuprequest based on facilityid
func newPickuprequestBasiccus(pickupmain *domain.Pickupmain) PickuprequestBasiccus {
	return PickuprequestBasiccus{
		Pickuprequestid:      pickupmain.Pickuprequestid,
		Customerid:           pickupmain.Customerid,
		Pickupscheduleslot:   pickupmain.Pickupscheduleslot,
		Pickupscheduledate:   pickupmain.Pickupscheduledate,
		Pickupagentid:        pickupmain.Pickupagentid,
		Pickupaddress:        pickupmain.Pickupaddress,
		Pickupstatus:         pickupmain.Pickupstatus,
		Customername:         pickupmain.Customername,
		Customermobilenumber: pickupmain.Customermobilenumber,
		Createddatetime:      pickupmain.Createddatetime,
		Modifieddatetime:     pickupmain.Modifieddatetime,
	}
}

// ////////////////////////////// pickuprequest based on facilityid

// type CountDetails struct {
// 	UnassignedCount int
// 	AssignedCount   int
// 	PickedupCount   int
// 	CancelledCount  int
// }

// func (counts CountDetails) CountDetails {
// 	return CountDetails{
// 		UnassignedCount: counts.UnassignedCount,
// 		AssignedCount:   counts.AssignedCount,
// 		PickedupCount:   counts.PickedupCount,
// 		CancelledCount:  counts.CancelledCount,
// 	}
// }

/*
func newAddressdetailsResponse2(addressdetailsList []*domain.Addressdetails) []addressdetailsResponse {
	var responses []addressdetailsResponse

	for _, addressdetails := range addressdetailsList {
		response := addressdetailsResponse{
			Customerid:   addressdetails.Customerid,
			Addressid:    addressdetails.Addressid,
			Firstname:    addressdetails.Firstname,
			Lastname:     addressdetails.Lastname,
			Addressline1: addressdetails.Addressline1,
			Addressline2: addressdetails.Addressline2,
			Landmark:     addressdetails.Landmark,
			City:         addressdetails.City,
			State:        addressdetails.State,
			Country:      addressdetails.Country,
			Pincode:      addressdetails.Pincode,
			Mobilenumber: addressdetails.Mobilenumber,
			Emailid:      addressdetails.Emailid,
			Geocode:      addressdetails.Geocode,
			Addresstype:  addressdetails.Addresstype,
			Fromtopickup: addressdetails.Fromtopickup,
			Isverified:   addressdetails.Isverified,
		}
		responses = append(responses, response)
	}

	return responses
}
*/

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	//port.ErrDataNotFound: http.StatusNotFound,
	port.ErrDataNotFound:               http.StatusNoContent,
	port.ErrConflictingData:            http.StatusConflict,
	port.ErrInvalidCredentials:         http.StatusUnauthorized,
	port.ErrUnauthorized:               http.StatusUnauthorized,
	port.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	port.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	port.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	port.ErrInvalidToken:               http.StatusUnauthorized,
	port.ErrExpiredToken:               http.StatusUnauthorized,
	port.ErrForbidden:                  http.StatusForbidden,
	port.ErrNoUpdatedData:              http.StatusBadRequest,
	port.ErrInsufficientStock:          http.StatusBadRequest,
	port.ErrInsufficientPayment:        http.StatusBadRequest,
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, err)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
/*func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errRsp := newErrorResponse(err.Error())

	ctx.JSON(statusCode, errRsp)
}*/

// handleAbort sends an error response and aborts the request with the specified status code and error message
func handleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	rsp := newErrorResponse(err.Error())
	ctx.AbortWithStatusJSON(statusCode, rsp)
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

//testing response

// TestGetaddressdetailsbyid
type AddressdetailsResponseTest []struct {
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
