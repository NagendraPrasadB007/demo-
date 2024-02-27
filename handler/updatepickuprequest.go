package handler

import (
	"net/http"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdatedetailsHandler struct {
	svc repo.UpdatedetailsRepository
	log *logger.Logger
}

func NewUpdatedetailsHandler(svc repo.UpdatedetailsRepository, log *logger.Logger) *UpdatedetailsHandler {
	return &UpdatedetailsHandler{
		svc,
		log,
	}
}

/*
type SubPiece1 struct {
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
*/

type updatepickuprequest struct {
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
	Pickupaddress             string    `json:"pickupaddress"`
	Domesticforeignidentifier string    `json:"domesticforeignidentifier" oneof=domestic international"`
	Pickuplong                string    `json:"pickuplong"`
	Pickuplat                 string    `json:"pickuplat"`
	Pickuprequestedpincode    string    `json:"pickuprequestedpincode"`
	// Customername              null.String `json:"customername"`
	Customername         string    `json:"customername"`
	Customermobilenumber string    `json:"customermobilenumber"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
	//domestic //
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
	//Customerid                string    `json:"customerid"`
	Contractnumber string `json:"contractnumber"`
	Isparcel       bool   `json:"isparcel"`
	Iscod          bool   `json:"iscod"`
	//INTERNNATIONAL //
	//Pickuprequestid           int       `json:"pickuprequestid"`
	//Articleid                 string    `json:"articleid"`
	//Articlestate              string    `json:"articlestate"`
	//Articletype               string    `json:"articletype"`
	//Articlecontent            string    `json:"articlecontent"`
	//Articleimageid            int       `json:"articleimageid"`
	//Articlepickupcharges      float64   `json:"articlepickupcharges"`
	//Ispremailing              bool      `json:"ispremailing"`
	//Isparcelpacking           bool      `json:"isparcelpacking"`
	//Createddatetime           time.Time `json:"createddatetime"`
	//Modifieddatetime          time.Time `json:"modifieddatetime"`
	//Customerdacpickup         string    `json:"customerdacpickup"`
	//Addresstype               string    `json:"addresstype"`
	//Bkgtransactionid          string    `json:"bkgtransactionid"`
	Origincountrycode      int `json:"origincountrycode"`
	Destinationcountrycode int `json:"destinationcountrycode"`
	//Physicalweight            float64   `json:"physicalweight"`
	Mailclass   string `json:"mailclass"`
	Contenttype string `json:"contenttype"`
	//Shape                     string    `json:"shape"`
	//Dimensionlength           float64   `json:"dimensionlength"`
	//Dimensionbreadth          float64   `json:"dimensionbreadth"`
	//Dimensionheight           float64   `json:"dimensionheight"`
	//Volumetricweight          float64   `json:"volumetricweight"`
	//Chargedweight             float64   `json:"chargedweight"`
	//Mailservicetypecode       string    `json:"mailservicetypecode"`
	//Bkgtype                   string    `json:"bkgtype"`
	//Mailform                  string    `json:"mailform"`
	//Isprepaid                 bool      `json:"isprepaid"`
	//Prepaymenttype            string    `json:"prepaymenttype"`
	//Valueofprepayment         float64   `json:"valueofprepayment"`
	//Vpcodflag                 bool      `json:"vpcodflag"`
	//Valueforvpcod             float64   `json:"valueforvpcod"`
	//Insuranceflag             bool      `json:"insuranceflag"`
	//Insurancetype             string    `json:"insurancetype"`
	//Valueinsurance            float64   `json:"valueinsurance"`
	//Acknowledgementpod        bool      `json:"acknowledgementpod"`
	//Instructionsrts           string    `json:"instructionsrts"`
	//Addressrefsender          string    `json:"addressrefsender"`
	//Addressrefreceiver        string    `json:"addressrefreceiver"`
	//Addressrefsenderaltaddr   string    `json:"addressrefsenderaltaddr"`
	//Addressrefreceiveraktaddr string    `json:"addressrefreceiveraktaddr"`
	//Barcodenumber             string    `json:"barcodenumber"`
	//Pickupflag                bool      `json:"pickupflag"`
	//Basetariff                float64   `json:"basetariff"`
	//Tax                       float64   `json:"tax"`
	//Totaltariff               float64   `json:"totaltariff"`
	//Modeofpayment             string    `json:"modeofpayment"`
	//Paymenttranid             string    `json:"paymenttranid"`
	//Status                    string    `json:"status"`
	//Createdon                 time.Time `json:"createdon"`
	//Createdby                 string    `json:"createdby"`
	//Updatedon                 time.Time `json:"updatedon"`
	//Updatedby                 string    `json:"updatedby"`
	//Authorisedon              time.Time `json:"authorisedon"`
	//Authorisedby              string    `json:"authorisedby"`
	//Facilityid                string    `json:"facilityid"`
	//Reqipaddress              string    `json:"reqipaddress"`
	//Bookingchannel            string    `json:"bookingchannel"`
	Consignmentvalue    int    `json:"consignmentvalue"`
	Mailexporttype      string `json:"mailexporttype"`
	Pbefilingtype       string `json:"pbefilingtype"`
	Declaration1        string `json:"declaration1"`
	Declaration23       string `json:"declaration23"`
	Declaration4        string `json:"declaration4"`
	Selffilingcusbroker string `json:"selffilingcusbroker"`
	Cusbrokerlicno      string `json:"cusbrokerlicno"`
	Cusbrokername       string `json:"cusbrokername"`
	Cusbrokeraddress    string `json:"cusbrokeraddress"`
	//Customerid                string    `json:"customerid"`
	//Contractnumber            string    `json:"contractnumber"`
	Gstn    string `json:"gstn"`
	Ibccode string `json:"ibccode"`
	Lut     string `json:"lut"`
	Adcode  string `json:"adcode"`
	//SubPieces []SubPiece `json:"subpieces"`
	//Isparcel                  bool      `json:"isparcel"`
	//Iscod                     bool      `json:"iscod"`

	//subpiecedetails//
	//Intlid                              int       `json:"intlid"`
	//adding subid
	Subid                               int       `json:"subid"`
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

	//tarriffdetails//
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
	//paymentdetails//
	//Paymenttranid   string    `json:"paymenttranid"`
	//Pickuprequestid int       `json:"pickuprequestid"`
	//Articleid       string    `json:"articleid"`
	Paymenttype string `json:"paymenttype"`
	//Modeofpayment   string    `json:"modeofpayment"`
	//Paymentstatus   string    `json:"paymentstatus"`
	Paymentdatetime time.Time `json:"paymentdatetime"`
	Paidamount      float64   `json:"paidamount"`
}

// Handler function to update pickuprequest details based on pickuprequestid

// UpdatePickuprequest godoc
//
//	@Summary		Update a Pickuprequest
//	@Description	update a pickuprequest based on pickuprequestid
//	@Tags			Pickuprequest
//	@Accept			json
//	@Produce		json
//	@Param			pickuprequestid				path		int					true	"update pickuprequest details"
//	@Param			updatepickuprequest	body		updatepickuprequest	true	"updatepickuprequest"
//	@Success 200 {string} string "Pickuprequest details Updated successfully"
//	@Failure		400						{object}	errorValidResponse			"Validation error"
//	@Failure		401						{object}	errorValidResponse			"Unauthorized error"
//	@Failure		403						{object}	errorValidResponse			"Forbidden error"
//	@Failure		404						{object}	errorValidResponse			"Data not found error"
//	@Failure		409						{object}	errorValidResponse			"Data conflict error"
//	@Failure		500						{object}	errorValidResponse			"Internal server error"
//	@Router			/updatedetails/{pickuprequestid} [put]
func (uh *UpdatedetailsHandler) updatedetailsbypickuprequestid(ctx *gin.Context) {
	var req updatepickuprequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, err)
		uh.log.Debug("error occured during binding:", err)
		return
	}
	if !handleValidation(ctx, req) {
		return
	}

	idStr := ctx.Param("pickuprequestid")
	// here idStr is of type string i am converting string into int because pickuprequestid in my DB is int
	pickuprequestID, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(ctx, err)
		uh.log.Debug("error occured while converting addressid:", err)
		return
	}

	pickupmain := domain.Pickupmain{
		Pickuprequestid:           pickuprequestID,
		Customerid:                req.Customerid,
		Pickupdroptype:            req.Pickupdroptype,
		Pickuplocation:            req.Pickuplocation,
		Droplocation:              req.Droplocation,
		Pickupscheduleslot:        req.Pickupscheduleslot,
		Pickupscheduledate:        req.Pickupscheduledate,
		Actualpickupdatetime:      req.Actualpickupdatetime,
		Pickupagentid:             req.Pickupagentid,
		Pickupfacilityid:          req.Pickupfacilityid,
		Pickupstatus:              req.Pickupstatus,
		Paymentstatus:             req.Paymentstatus,
		Pickupaddress:             req.Pickupaddress,
		Domesticforeignidentifier: req.Domesticforeignidentifier,
		Pickuplong:                req.Pickuplong,
		Pickuplat:                 req.Pickuplat,
		Pickuprequestedpincode:    req.Pickuprequestedpincode,
		Customername:              req.Customername,
		Customermobilenumber:      req.Customermobilenumber,
		Assigneddatetime:          req.Assigneddatetime,
	}

	domestic := domain.Domesticarticledetails{
		Articleid:            req.Articleid,
		Articlestate:         req.Articlestate,
		Articletype:          req.Articletype,
		Articlecontent:       req.Articlecontent,
		Articleimageid:       req.Articleimageid,
		Articlepickupcharges: req.Articlepickupcharges,
		Ispremailing:         req.Ispremailing,
		Isparcelpacking:      req.Isparcelpacking,
		//Createddatetime:           req.Createddatetime,
		//Modifieddatetime:          req.Modifieddatetime,
		Customerdacpickup:         req.Customerdacpickup,
		Addresstype:               req.Addresstype,
		Bkgtransactionid:          req.Bkgtransactionid,
		Originpin:                 req.Originpin,
		Destinationpin:            req.Destinationpin,
		Physicalweight:            req.Physicalweight,
		Shape:                     req.Shape,
		Dimensionlength:           req.Dimensionlength,
		Dimensionbreadth:          req.Dimensionbreadth,
		Dimensionheight:           req.Dimensionheight,
		Volumetricweight:          req.Volumetricweight,
		Chargedweight:             req.Chargedweight,
		Mailservicetypecode:       req.Mailservicetypecode,
		Bkgtype:                   req.Bkgtype,
		Mailform:                  req.Mailform,
		Isprepaid:                 req.Isprepaid,
		Prepaymenttype:            req.Prepaymenttype,
		Valueofprepayment:         req.Valueofprepayment,
		Vpcodflag:                 req.Vpcodflag,
		Valueforvpcod:             req.Valueforvpcod,
		Insuranceflag:             req.Insuranceflag,
		Insurancetype:             req.Insurancetype,
		Valueinsurance:            req.Valueinsurance,
		Acknowledgementpod:        req.Acknowledgementpod,
		Instructionsrts:           req.Instructionsrts,
		Addressrefsender:          req.Addressrefsender,
		Addressrefreceiver:        req.Addressrefreceiver,
		Addressrefsenderaltaddr:   req.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: req.Addressrefreceiveraktaddr,
		Barcodenumber:             req.Barcodenumber,
		Pickupflag:                req.Pickupflag,
		Basetariff:                req.Basetariff,
		Tax:                       req.Tax,
		Totaltariff:               req.Totaltariff,
		Modeofpayment:             req.Modeofpayment,
		Paymenttranid:             req.Paymenttranid,
		Status:                    req.Status,
		//Createdon:                 req.Createdon,
		Createdby: req.Createdby,
		//Updatedon:                 req.Updatedon,
		Updatedby:      req.Updatedby,
		Authorisedon:   req.Authorisedon,
		Authorisedby:   req.Authorisedby,
		Facilityid:     req.Facilityid,
		Reqipaddress:   req.Reqipaddress,
		Bookingchannel: req.Bookingchannel,
		Customerid:     req.Customerid,
		Contractnumber: req.Contractnumber,
		Isparcel:       req.Isparcel,
		Iscod:          req.Iscod,
	}

	international := domain.Internationalarticledetails{
		Pickuprequestid:      req.Pickupagentid,
		Articleid:            req.Articleid,
		Articlestate:         req.Articlestate,
		Articletype:          req.Articletype,
		Articlecontent:       req.Articlecontent,
		Articleimageid:       req.Articleimageid,
		Articlepickupcharges: req.Articlepickupcharges,
		Ispremailing:         req.Ispremailing,
		Isparcelpacking:      req.Isparcelpacking,
		//Createddatetime:           req.Createddatetime,
		//Modifieddatetime:          req.Modifieddatetime,
		Customerdacpickup:         req.Customerdacpickup,
		Addresstype:               req.Addresstype,
		Bkgtransactionid:          req.Bkgtransactionid,
		Origincountrycode:         req.Origincountrycode,
		Destinationcountrycode:    req.Destinationcountrycode,
		Physicalweight:            req.Physicalweight,
		Mailclass:                 req.Mailclass,
		Contenttype:               req.Contenttype,
		Shape:                     req.Shape,
		Dimensionlength:           req.Dimensionlength,
		Dimensionbreadth:          req.Dimensionbreadth,
		Dimensionheight:           req.Dimensionheight,
		Volumetricweight:          req.Volumetricweight,
		Chargedweight:             req.Chargedweight,
		Mailservicetypecode:       req.Mailservicetypecode,
		Bkgtype:                   req.Bkgtype,
		Mailform:                  req.Mailform,
		Isprepaid:                 req.Isprepaid,
		Prepaymenttype:            req.Prepaymenttype,
		Valueofprepayment:         req.Valueofprepayment,
		Vpcodflag:                 req.Vpcodflag,
		Valueforvpcod:             req.Valueforvpcod,
		Insuranceflag:             req.Insuranceflag,
		Insurancetype:             req.Insurancetype,
		Valueinsurance:            req.Valueinsurance,
		Acknowledgementpod:        req.Acknowledgementpod,
		Instructionsrts:           req.Instructionsrts,
		Addressrefsender:          req.Addressrefsender,
		Addressrefreceiver:        req.Addressrefreceiver,
		Addressrefsenderaltaddr:   req.Addressrefsenderaltaddr,
		Addressrefreceiveraktaddr: req.Addressrefreceiveraktaddr,
		Barcodenumber:             req.Barcodenumber,
		Pickupflag:                req.Pickupflag,
		Basetariff:                req.Basetariff,
		Tax:                       req.Tax,
		Totaltariff:               req.Totaltariff,
		Modeofpayment:             req.Modeofpayment,
		Paymenttranid:             req.Paymenttranid,
		Status:                    req.Status,
		//Createdon:                 req.Createdon,
		Createdby: req.Createdby,
		//Updatedon:                 req.Updatedon,
		Updatedby:           req.Updatedby,
		Authorisedon:        req.Authorisedon,
		Authorisedby:        req.Authorisedby,
		Facilityid:          req.Facilityid,
		Reqipaddress:        req.Reqipaddress,
		Bookingchannel:      req.Bookingchannel,
		Consignmentvalue:    req.Consignmentvalue,
		Mailexporttype:      req.Mailexporttype,
		Pbefilingtype:       req.Pbefilingtype,
		Declaration1:        req.Declaration1,
		Declaration23:       req.Declaration23,
		Declaration4:        req.Declaration4,
		Selffilingcusbroker: req.Selffilingcusbroker,
		Cusbrokerlicno:      req.Cusbrokerlicno,
		Cusbrokername:       req.Cusbrokername,
		Cusbrokeraddress:    req.Cusbrokeraddress,
		Customerid:          req.Customerid,
		Contractnumber:      req.Contractnumber,
		Gstn:                req.Gstn,
		Ibccode:             req.Ibccode,
		Lut:                 req.Lut,
		Adcode:              req.Adcode,
		Isparcel:            req.Isparcel,
		Iscod:               req.Iscod,
	}

	subpiece := domain.SubPiecedetails{
		//added subid
		Subid:                               req.Subid,
		Identifierpieceid:                   req.Identifierpieceid,
		Subpiececatproductcode:              req.Subpiececatproductcode,
		Hscode:                              req.Hscode,
		Productcustomstariffhead:            req.Productcustomstariffhead,
		Productdescription:                  req.Productdescription,
		Isocodefororigincountry:             req.Isocodefororigincountry,
		Unitforsubpiecequantity:             req.Unitforsubpiecequantity,
		Subpiecequantitycount:               req.Subpiecequantitycount,
		Producttotalvalueasperinvoice:       req.Producttotalvalueasperinvoice,
		Isocodeforcurrency:                  req.Isocodeforcurrency,
		Subpieceweight:                      req.Subpieceweight,
		Subpieceweightnett:                  req.Subpieceweightnett,
		Productinvoicenumber:                req.Productinvoicenumber,
		Productinvoicedate:                  req.Productinvoicedate,
		Statusforecommerce:                  req.Statusforecommerce,
		Urlforecommerceconsignment:          req.Urlforecommerceconsignment,
		Ecommercepaymenttransactionid:       req.Ecommercepaymenttransactionid,
		Ecommerceskuno:                      req.Ecommerceskuno,
		Taxinvoicenumber:                    req.Taxinvoicenumber,
		Taxinvoicedate:                      req.Taxinvoicedate,
		Serialnumberforsubpieceintaxinvoice: req.Serialnumberforsubpieceintaxinvoice,
		Valueofsubpieceaspertaxinvoice:      req.Valueofsubpieceaspertaxinvoice,
		Assessablefreeonboardvalue:          req.Assessablefreeonboardvalue,
		Isocodeforassessablecurrency:        req.Isocodeforassessablecurrency,
		Exchangerateforasblcurr:             req.Exchangerateforasblcurr,
		Assessableamount:                    req.Assessableamount,
		Rateforexportduty:                   req.Rateforexportduty,
		Exportdutyamount:                    req.Exportdutyamount,
		Rateforcess:                         req.Rateforcess,
		Cessamount:                          req.Cessamount,
		Igstrate:                            req.Igstrate,
		Igstamount:                          req.Igstamount,
		Compensationrate:                    req.Compensationrate,
		Compensationamount:                  req.Compensationamount,
		Detailsofletterofundertakingorbond:  req.Detailsofletterofundertakingorbond,
		Modeofpayment:                       req.Modeofpayment,
		Paymenttransactionid:                req.Paymenttransactionid,
		//Createdon:                           req.Createdon,
		Createdby: req.Createdby,
		//Updatedon:                           req.Updatedon,
		Updatedby:          req.Updatedby,
		Authorisedon:       req.Authorisedon,
		Authorisedby:       req.Authorisedby,
		Facilityid:         req.Facilityid,
		Ipaddress:          req.Ipaddress,
		Bookingchanneltype: req.Bookingchanneltype,
	}
	uh.log.Debug(subpiece)

	tariff := domain.Tariffdetails{
		Articleid:           req.Articleid,
		Totalamount:         req.Totalamount,
		Pickupcharges:       req.Pickupcharges,
		Registrationfee:     req.Registrationfee,
		Postage:             req.Postage,
		Ackorpodfee:         req.Ackorpodfee,
		Valueinsurance:      req.Valueinsurance,
		Valueforvpcod:       req.Valueforvpcod,
		Doordeliverycharges: req.Doordeliverycharges,
		Packingfee:          req.Packingfee,
		Cgst:                req.Cgst,
		Sgst:                req.Sgst,
		Othercharges:        req.Othercharges,
	}

	payment := domain.Paymentdetails{
		Paymenttranid: req.Paymenttranid,
		//Pickuprequestid: req.Pickuprequestid,
		Articleid:       req.Articleid,
		Paymenttype:     req.Paymenttype,
		Modeofpayment:   req.Modeofpayment,
		Paymentstatus:   req.Paymentstatus,
		Paymentdatetime: req.Paymentdatetime,
		Paidamount:      req.Paidamount,
	}

	//checking wheather pickuprequestid is domestic or international
	returnvalue, err := uh.svc.IdentifyUpdate(ctx, pickuprequestID)
	if err != nil {
		uh.log.Debug(err)
		return
	}

	result := returnvalue.Domesticforeignidentifier
	uh.log.Debug(result)

	if result == "domestic" {

		err = uh.svc.UpdatepickuprequestdetailsDom(ctx, &pickupmain, &domestic, &tariff, &payment)
		if err != nil {
			handledbError(ctx, err)
			uh.log.Debug("error occured while repo calling:", err.Error())
		}
		handleSuccess(ctx, "Pickuprequest details Updated successfully")

	} else {
		uh.log.Debug("entered into else loop")

		// Retrieve sub_id from the payload
		subID, err := getSubIDFromPayload(subpiece)
		if err != nil {
			handledbError(ctx, err)
			uh.log.Debug("error occurred while getting sub_id:", err.Error())
			return
		}

		// Check if sub_id is valid for international requests
		if subID == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "subid is required in the request payload",
			})
			return
		}

		/*
			// Retrieve all SubIDs associated with the intl_id
			associatedSubIDs, err := uh.svc.GetAssociatedSubIDs(ctx, international.Intlid)
			if err != nil {
				handledbError(ctx, err)
				uh.log.Debug("error occurred while retrieving associated SubIDs:", err.Error())
				return
			}

			uh.log.Debug("associatedSubIDs : ", associatedSubIDs)

			// Check if the provided subID is in the list of associated SubIDs
			if !isSubIDInList(subID, associatedSubIDs) {
				// Respond with a JSON error message
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid sub_id for the provided intl_id",
				})
				return
			}
		*/

		err = uh.svc.UpdatepickuprequestdetailsInt(ctx, &pickupmain, &international, &subpiece, &tariff, &payment)
		if err != nil {
			handledbError(ctx, err)
			uh.log.Debug("error occured while repo calling:", err.Error())
			return
		}

		handleSuccess(ctx, "Pickuprequest details Updated successfully")

	}
}

func getSubIDFromPayload(subpiece domain.SubPiecedetails) (int, error) {
	// Assuming sub_id is a field in the SubPiece structure
	subID := subpiece.Subid

	return subID, nil
}

/*
func isSubIDInList(subID int, subIDs []int) bool {
	// Check if subID is in the list of associated SubIDs
	for _, id := range subIDs {
		if subID == id {
			return true
		}
	}
	return false
}
*/
