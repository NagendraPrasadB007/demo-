package repository

import (
	"context"
	"errors"
	"net/http"
	"pickupmanagement/core/domain"
	"pickupmanagement/core/port"
	"pickupmanagement/logger"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type FetchdetailsRepository struct {
	db  *DB
	log *logger.Logger
}

func NewFetchdetailsRepository(db *DB, log *logger.Logger) *FetchdetailsRepository {
	return &FetchdetailsRepository{
		db,
		log,
	}
}

func (ar *FetchdetailsRepository) Identify(ctx context.Context, Pickuprequestid int) (*domain.Pickupmain, error) {
	ar.log.Debug("insided identify function ")
	var pickupmain domain.Pickupmain
	query := psql.
		Select("pickuprequest_id", "domestic_foreign_identifier").
		//Select("*").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": Pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	err = ar.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		// &pickupmain.Customerid,
		// &pickupmain.Pickupdroptype,
		// &pickupmain.Pickuplocation,
		// &pickupmain.Droplocation,
		// &pickupmain.Pickupscheduleslot,
		// &pickupmain.Pickupscheduledate,
		// &pickupmain.Actualpickupdatetime,
		// &pickupmain.Pickupagentid,
		// &pickupmain.Pickupfacilityid,
		// &pickupmain.Pickupstatus,
		// &pickupmain.Paymentstatus,
		// &pickupmain.Createddatetime,
		// &pickupmain.Pickupaddress,
		&pickupmain.Domesticforeignidentifier,
		// &pickupmain.Pickuplong,
		// &pickupmain.Pickuplat,
		// &pickupmain.Modifieddatetime,
		// &pickupmain.Pickuprequestedpincode,
		// &pickupmain.Customername,
		// &pickupmain.Customermobilenumber,
		// &pickupmain.Assigneddatetime,
	)

	return &pickupmain, err

}

// Repo function to fetch pickuprequest details based on pickuprequestid
func (ar *FetchdetailsRepository) GetdetailsByPickuprequestID(gctx *gin.Context, Pickuprequestid int) (*domain.Pickupmain, *domain.Domesticarticledetails, *domain.Tariffdetails, *domain.Paymentdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the pickuprequestid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": Pickuprequestid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, nil, nil, nil, existsErr
	}

	var count int
	if err := ar.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, nil, nil, nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, nil, nil, nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Pickuprequestid " + strconv.Itoa(Pickuprequestid) + " does not exist",
		})

		return nil, nil, nil, nil, errors.New("Pickuprequestid does not exist")
	}

	var pickupmain domain.Pickupmain
	var domestic domain.Domesticarticledetails
	var tariff domain.Tariffdetails
	var payment domain.Paymentdetails

	query := psql.
		// Select("pickuprequest_id", "customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "created_datetime", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "modified_datetime", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime", "dom_id", "pickuprequest_id", "article_id", "article_state", "article_type", "article_content", "articleimageid", "articlepickupcharges", "is_premailing", "is_parcelpacking", "created_datetime", "modified_datetime", "customer_dac_pickup", "address_type", "bkg_transaction_id", "origin_pin", "destination_pin", "physical_weight", "shape", "dimension_length", "dimension_breadth", "dimension_height", "volumetric_weight", "charged_weight", "mail_service_typecode", "bkg_type", "mail_form", "is_prepaid", "prepayment_type", "value_of_prepayment", "vp_cod_flag", "value_for_vp_cod", "insurance_flag", "insurance_type", "value_insurance", "acknowledgement_pod", "instructions_rts", "address_ref_sender", "address_ref_receiver", "address_ref_sender_alt_addr", "address_ref_receiver_akt_addr", "barcode_number", "pickup_flag", "base_tariff", "tax", "total_tariff", "mode_of_payment", "payment_tranid", "status", "created_on", "created_by", "updated_on", "updated_by", "authorised_on", "authorised_by", "facility_id", "req_ip_address", "booking_channel", "customer_id", "contract_number", "isparcel", "iscod", "pickuprequest_id", "article_id", "total_amount", "pickup_charges", "registration_fee", "postage", "ack_pod_fee", "value_insurance", "value_for_vp_cod", "doordelivery_charges", "packing_fee", "cgst", "sgst", "other_charges", "tariff_id", "payment_tranid", "pickuprequest_id", "article_id", "payment_type", "mode_of_payment", "payment_status", "payment_datetime", "paid_amount", "payment_id").
		Select("*").
		From("pickup_main").
		Join("domestic_articledetails ON pickup_main.pickuprequest_id = domestic_articledetails.pickuprequest_id").
		Join("tariff_details ON pickup_main.pickuprequest_id = tariff_details.pickuprequest_id").
		Join("payment_details ON pickup_main.pickuprequest_id = payment_details.pickuprequest_id").
		Where(sq.Eq{"pickup_main.pickuprequest_id": Pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	err = ar.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		&pickupmain.Customerid,
		&pickupmain.Pickupdroptype,
		&pickupmain.Pickuplocation,
		&pickupmain.Droplocation,
		&pickupmain.Pickupscheduleslot,
		&pickupmain.Pickupscheduledate,
		&pickupmain.Actualpickupdatetime,
		&pickupmain.Pickupagentid,
		&pickupmain.Pickupfacilityid,
		&pickupmain.Pickupstatus,
		&pickupmain.Paymentstatus,
		&pickupmain.Createddatetime,
		&pickupmain.Pickupaddress,
		&pickupmain.Domesticforeignidentifier,
		&pickupmain.Pickuplong,
		&pickupmain.Pickuplat,
		&pickupmain.Modifieddatetime,
		&pickupmain.Pickuprequestedpincode,
		&pickupmain.Customername,
		&pickupmain.Customermobilenumber,
		&pickupmain.Assigneddatetime,

		&domestic.Domid,
		&domestic.Pickuprequestid,
		&domestic.Articleid,
		&domestic.Articlestate,
		&domestic.Articletype,
		&domestic.Articlecontent,
		&domestic.Articleimageid,
		&domestic.Articlepickupcharges,
		&domestic.Ispremailing,
		&domestic.Isparcelpacking,
		&domestic.Createddatetime,
		&domestic.Modifieddatetime,
		&domestic.Customerdacpickup,
		&domestic.Addresstype,
		&domestic.Bkgtransactionid,
		&domestic.Originpin,
		&domestic.Destinationpin,
		&domestic.Physicalweight,
		&domestic.Shape,
		&domestic.Dimensionlength,
		&domestic.Dimensionbreadth,
		&domestic.Dimensionheight,
		&domestic.Volumetricweight,
		&domestic.Chargedweight,
		&domestic.Mailservicetypecode,
		&domestic.Bkgtype,
		&domestic.Mailform,
		&domestic.Isprepaid,
		&domestic.Prepaymenttype,
		&domestic.Valueofprepayment,
		&domestic.Vpcodflag,
		&domestic.Valueforvpcod,
		&domestic.Insuranceflag,
		&domestic.Insurancetype,
		&domestic.Valueinsurance,
		&domestic.Acknowledgementpod,
		&domestic.Instructionsrts,
		&domestic.Addressrefsender,
		&domestic.Addressrefreceiver,
		&domestic.Addressrefsenderaltaddr,
		&domestic.Addressrefreceiveraktaddr,
		&domestic.Barcodenumber,
		&domestic.Pickupflag,
		&domestic.Basetariff,
		&domestic.Tax,
		&domestic.Totaltariff,
		&domestic.Modeofpayment,
		&domestic.Paymenttranid,
		&domestic.Status,
		&domestic.Createdon,
		&domestic.Createdby,
		&domestic.Updatedon,
		&domestic.Updatedby,
		&domestic.Authorisedon,
		&domestic.Authorisedby,
		&domestic.Facilityid,
		&domestic.Reqipaddress,
		&domestic.Bookingchannel,
		&domestic.Customerid,
		&domestic.Contractnumber,
		&domestic.Isparcel,
		&domestic.Iscod,

		&tariff.Pickuprequestid,
		&tariff.Articleid,
		&tariff.Totalamount,
		&tariff.Pickupcharges,
		&tariff.Registrationfee,
		&tariff.Postage,
		&tariff.Ackorpodfee,
		&tariff.Valueinsurance,
		&tariff.Valueforvpcod,
		&tariff.Doordeliverycharges,
		&tariff.Packingfee,
		&tariff.Cgst,
		&tariff.Sgst,
		&tariff.Othercharges,
		&tariff.Tariffid,

		&payment.Paymenttranid,
		&payment.Pickuprequestid,
		&payment.Articleid,
		&payment.Paymenttype,
		&payment.Modeofpayment,
		&payment.Paymentstatus,
		&payment.Paymentdatetime,
		&payment.Paidamount,
		&payment.Paymentid,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil, nil, nil, port.ErrDataNotFound
		}
		return nil, nil, nil, nil, err
	}

	return &pickupmain, &domestic, &tariff, &payment, nil

}

//includes subpiece details table
// func (ar *FetchdetailsRepository) GetIntdetailsByPickuprequestID(ctx context.Context, Pickuprequestid int) (*domain.Pickupmain, *domain.Internationalarticledetails, *domain.SubPiecedetails, *domain.Tariffdetails, *domain.Paymentdetails, error) {

// Repo function to fetch pickuprequest details based on pickuprequestid
func (ar *FetchdetailsRepository) GetIntdetailsByPickuprequestID(gctx *gin.Context, Pickuprequestid int) (*domain.Pickupmain, *domain.Internationalarticledetails, []domain.SubPiecedetails, *domain.Tariffdetails, *domain.Paymentdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//mne
	// Check if the pickuprequestid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": Pickuprequestid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, nil, nil, nil, nil, existsErr
	}

	var count int
	if err := ar.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if count == 0 {
		// Assuming that ctx is of type *gin.Context
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			// Handle the case where ctx is not of the expected type
			return nil, nil, nil, nil, nil, errors.New("unexpected context type")
		}

		// Use ginContext.JSON instead of ctx.JSON
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Pickuprequestid " + strconv.Itoa(Pickuprequestid) + " does not exist",
		})

		return nil, nil, nil, nil, nil, errors.New("Pickuprequestid does not exist")
	}

	// get api starts here

	var pickupmain domain.Pickupmain
	var international domain.Internationalarticledetails
	var subpiece domain.SubPiecedetails
	var tariff domain.Tariffdetails
	var payment domain.Paymentdetails

	query := psql.
		// Select("pickuprequest_id", "customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "created_datetime", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "modified_datetime", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime", "intl_id", "pickuprequest_id", "article_id", "article_state", "article_type", "article_content", "articleimageid", "articlepickupcharges", "is_premailing", "is_parcelpacking", "created_datetime", "modified_datetime", "customer_dac_pickup", "address_type", "bkg_transaction_id", "origin_countrycode", "destination_countrycode", "physical_weight", "mail_class", "content_type", "shape", "dimension_length", "dimension_breadth", "dimension_height", "volumetric_weight", "charged_weight", "mail_service_typecode", "bkg_type", "mail_form", "is_prepaid", "prepayment_type", "value_of_prepayment", "vp_cod_flag", "value_for_vp_cod", "insurance_flag", "insurance_type", "value_insurance", "acknowledgement_pod", "instructions_rts", "address_ref_sender", "address_ref_receiver", "address_ref_sender_alt_addr", "address_ref_receiver_akt_addr", "barcode_number", "pickup_flag", "base_tariff", "tax", "total_tariff", "mode_of_payment", "payment_tranid", "status", "created_on", "created_by", "updated_on", "updated_by", "authorised_on", "authorised_by", "facility_id", "req_ip_address", "booking_channel", "consignment_value", "mail_export_type", "pbe_filing_type", "declaration1", "declaration2_3", "declaration4", "selffiling_cusbroker", "cusbroker_licno", "cusbroker_name", "cusbroker_address", "customer_id", "contract_number", "gstn", "ibccode", "lut", "adcode", "isparcel", "iscod", "pickuprequest_id", "article_id", "total_amount", "pickup_charges", "registration_fee", "postage", "ack_pod_fee", "value_insurance", "value_for_vp_cod", "doordelivery_charges", "packing_fee", "cgst", "sgst", "other_charges", "tariff_id", "payment_tranid", "pickuprequest_id", "article_id", "payment_type", "mode_of_payment", "payment_status", "payment_datetime", "paid_amount", "payment_id").
		Select("*").
		From("pickup_main").
		Join("international_articledetails ON pickup_main.pickuprequest_id = international_articledetails.pickuprequest_id").
		//Join("sub_piece_details ON international_articledetails.intl_id = sub_piece_details.intl_id").
		Join("tariff_details ON pickup_main.pickuprequest_id = tariff_details.pickuprequest_id").
		Join("payment_details ON pickup_main.pickuprequest_id = payment_details.pickuprequest_id").
		Where(sq.Eq{"pickup_main.pickuprequest_id": Pickuprequestid})

	// subquery := psql.Select("*").From("sub_piece_details").Where(sq.Eq{"intl_id": international.Intlid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	err = ar.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		&pickupmain.Customerid,
		&pickupmain.Pickupdroptype,
		&pickupmain.Pickuplocation,
		&pickupmain.Droplocation,
		&pickupmain.Pickupscheduleslot,
		&pickupmain.Pickupscheduledate,
		&pickupmain.Actualpickupdatetime,
		&pickupmain.Pickupagentid,
		&pickupmain.Pickupfacilityid,
		&pickupmain.Pickupstatus,
		&pickupmain.Paymentstatus,
		&pickupmain.Createddatetime,
		&pickupmain.Pickupaddress,
		&pickupmain.Domesticforeignidentifier,
		&pickupmain.Pickuplong,
		&pickupmain.Pickuplat,
		&pickupmain.Modifieddatetime,
		&pickupmain.Pickuprequestedpincode,
		&pickupmain.Customername,
		&pickupmain.Customermobilenumber,
		&pickupmain.Assigneddatetime,

		&international.Intlid,
		&international.Pickuprequestid,
		&international.Articleid,
		&international.Articlestate,
		&international.Articletype,
		&international.Articlecontent,
		&international.Articleimageid,
		&international.Articlepickupcharges,
		&international.Ispremailing,
		&international.Isparcelpacking,
		&international.Createddatetime,
		&international.Modifieddatetime,
		&international.Customerdacpickup,
		&international.Addresstype,
		&international.Bkgtransactionid,
		&international.Origincountrycode,
		&international.Destinationcountrycode,
		&international.Physicalweight,
		&international.Mailclass,
		&international.Contenttype,
		&international.Shape,
		&international.Dimensionlength,
		&international.Dimensionbreadth,
		&international.Dimensionheight,
		&international.Volumetricweight,
		&international.Chargedweight,
		&international.Mailservicetypecode,
		&international.Bkgtype,
		&international.Mailform,
		&international.Isprepaid,
		&international.Prepaymenttype,
		&international.Valueofprepayment,
		&international.Vpcodflag,
		&international.Valueforvpcod,
		&international.Insuranceflag,
		&international.Insurancetype,
		&international.Valueinsurance,
		&international.Acknowledgementpod,
		&international.Instructionsrts,
		&international.Addressrefsender,
		&international.Addressrefreceiver,
		&international.Addressrefsenderaltaddr,
		&international.Addressrefreceiveraktaddr,
		&international.Barcodenumber,
		&international.Pickupflag,
		&international.Basetariff,
		&international.Tax,
		&international.Totaltariff,
		&international.Modeofpayment,
		&international.Paymenttranid,
		&international.Status,
		&international.Createdon,
		&international.Createdby,
		&international.Updatedon,
		&international.Updatedby,
		&international.Authorisedon,
		&international.Authorisedby,
		&international.Facilityid,
		&international.Reqipaddress,
		&international.Bookingchannel,
		&international.Consignmentvalue,
		&international.Mailexporttype,
		&international.Pbefilingtype,
		&international.Declaration1,
		&international.Declaration23,
		&international.Declaration4,
		&international.Selffilingcusbroker,
		&international.Cusbrokerlicno,
		&international.Cusbrokername,
		&international.Cusbrokeraddress,
		&international.Customerid,
		&international.Contractnumber,
		&international.Gstn,
		&international.Ibccode,
		&international.Lut,
		&international.Adcode,
		&international.Isparcel,
		&international.Iscod,
		/*
			&subpiece.Subid,
			&subpiece.Intlid,
			&subpiece.Identifierpieceid,
			&subpiece.Subpiececatproductcode,
			&subpiece.Hscode,
			&subpiece.Productcustomstariffhead,
			&subpiece.Productdescription,
			&subpiece.Isocodefororigincountry,
			&subpiece.Unitforsubpiecequantity,
			&subpiece.Subpiecequantitycount,
			&subpiece.Producttotalvalueasperinvoice,
			&subpiece.Isocodeforcurrency,
			&subpiece.Subpieceweight,
			&subpiece.Subpieceweightnett,
			&subpiece.Productinvoicenumber,
			&subpiece.Productinvoicedate,
			&subpiece.Statusforecommerce,
			&subpiece.Urlforecommerceconsignment,
			&subpiece.Ecommercepaymenttransactionid,
			&subpiece.Ecommerceskuno,
			&subpiece.Taxinvoicenumber,
			&subpiece.Taxinvoicedate,
			&subpiece.Serialnumberforsubpieceintaxinvoice,
			&subpiece.Valueofsubpieceaspertaxinvoice,
			&subpiece.Assessablefreeonboardvalue,
			&subpiece.Isocodeforassessablecurrency,
			&subpiece.Exchangerateforasblcurr,
			&subpiece.Assessableamount,
			&subpiece.Rateforexportduty,
			&subpiece.Exportdutyamount,
			&subpiece.Rateforcess,
			&subpiece.Cessamount,
			&subpiece.Igstrate,
			&subpiece.Igstamount,
			&subpiece.Compensationrate,
			&subpiece.Compensationamount,
			&subpiece.Detailsofletterofundertakingorbond,
			&subpiece.Modeofpayment,
			&subpiece.Paymenttransactionid,
			&subpiece.Createdon,
			&subpiece.Createdby,
			&subpiece.Updatedon,
			&subpiece.Updatedby,
			&subpiece.Authorisedon,
			&subpiece.Authorisedby,
			&subpiece.Facilityid,
			&subpiece.Ipaddress,
			&subpiece.Bookingchanneltype,
		*/

		&tariff.Pickuprequestid,
		&tariff.Articleid,
		&tariff.Totalamount,
		&tariff.Pickupcharges,
		&tariff.Registrationfee,
		&tariff.Postage,
		&tariff.Ackorpodfee,
		&tariff.Valueinsurance,
		&tariff.Valueforvpcod,
		&tariff.Doordeliverycharges,
		&tariff.Packingfee,
		&tariff.Cgst,
		&tariff.Sgst,
		&tariff.Othercharges,
		&tariff.Tariffid,

		&payment.Paymenttranid,
		&payment.Pickuprequestid,
		&payment.Articleid,
		&payment.Paymenttype,
		&payment.Modeofpayment,
		&payment.Paymentstatus,
		&payment.Paymentdatetime,
		&payment.Paidamount,
		&payment.Paymentid,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil, nil, nil, nil, port.ErrDataNotFound
		}
		return nil, nil, nil, nil, nil, err
	}

	//getting subid details based on intlid

	subquery := psql.Select("sub_id", "intl_id", "identifier_piece_id", "subpiecec_cat_product_code", "hscode", "product_customs_tariff_head", "product_description", "iso_code_for_origin_country", "unit_for_sub_piece_quantity", "sub_piece_quantity_count", "product_total_value_as_per_invoice", "iso_code_for_currency", "sub_piece_weight", "sub_piece_weight_nett", "product_invoice_number", "product_invoice_date", "status_for_ecommerce", "url_for_ecommerce_consignment", "ecommerce_payment_transaction_id", "ecommerce_sku_no", "tax_invoice_number", "tax_invoice_date", "serial_number_for_sub_piece_in_tax_invoice", "value_of_sub_piece_as_per_tax_invoice", "assessable_free_on_board_value", "iso_code_for_assessable_currency", "exchange_rate_for_asbl_curr", "assessable_amount", "rate_for_export_duty", "export_duty_amount", "rate_for_cess", "cess_amount", "igst_rate", "igst_amount", "compensation_rate", "compensation_amount", "details_of_letter_of_undertaking_or_bond", "mode_of_payment", "payment_transaction_id", "created_on", "created_by", "updated_on", "updated_by", "authorised_on", "authorised_by", "facility_id", "ip_address", "booking_channel_type").From("sub_piece_details").Where(sq.Eq{"intl_id": international.Intlid})
	ar.log.Debug("intl id is :", international.Intlid)
	sql, args, err = subquery.ToSql()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	rows, err := ar.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	for rows.Next() {
		//var subpiece domain.SubPiecedetails
		//ar.log.Debug("entered rows.next method")
		err := rows.Scan(
			&subpiece.Subid,
			&subpiece.Intlid,
			&subpiece.Identifierpieceid,
			&subpiece.Subpiececatproductcode,
			&subpiece.Hscode,
			&subpiece.Productcustomstariffhead,
			&subpiece.Productdescription,
			&subpiece.Isocodefororigincountry,
			&subpiece.Unitforsubpiecequantity,
			&subpiece.Subpiecequantitycount,
			&subpiece.Producttotalvalueasperinvoice,
			&subpiece.Isocodeforcurrency,
			&subpiece.Subpieceweight,
			&subpiece.Subpieceweightnett,
			&subpiece.Productinvoicenumber,
			&subpiece.Productinvoicedate,
			&subpiece.Statusforecommerce,
			&subpiece.Urlforecommerceconsignment,
			&subpiece.Ecommercepaymenttransactionid,
			&subpiece.Ecommerceskuno,
			&subpiece.Taxinvoicenumber,
			&subpiece.Taxinvoicedate,
			&subpiece.Serialnumberforsubpieceintaxinvoice,
			&subpiece.Valueofsubpieceaspertaxinvoice,
			&subpiece.Assessablefreeonboardvalue,
			&subpiece.Isocodeforassessablecurrency,
			&subpiece.Exchangerateforasblcurr,
			&subpiece.Assessableamount,
			&subpiece.Rateforexportduty,
			&subpiece.Exportdutyamount,
			&subpiece.Rateforcess,
			&subpiece.Cessamount,
			&subpiece.Igstrate,
			&subpiece.Igstamount,
			&subpiece.Compensationrate,
			&subpiece.Compensationamount,
			&subpiece.Detailsofletterofundertakingorbond,
			&subpiece.Modeofpayment,
			&subpiece.Paymenttransactionid,
			&subpiece.Createdon,
			&subpiece.Createdby,
			&subpiece.Updatedon,
			&subpiece.Updatedby,
			&subpiece.Authorisedon,
			&subpiece.Authorisedby,
			&subpiece.Facilityid,
			&subpiece.Ipaddress,
			&subpiece.Bookingchanneltype,
		)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		international.SubPieces = append(international.SubPieces, subpiece)
		//ar.log.Debug("subpiece:", subpiece)
		ar.log.Debug("international.SubPieces:", international.SubPieces)
	}

	//return &pickupmain, &international, &subpiece, &tariff, &payment, nil
	return &pickupmain, &international, international.SubPieces, &tariff, &payment, nil

}

/*
func (ar *FetchdetailsRepository) getSubpiecesByInternationalID(ctx context.Context, Intlid int) ([]*domain.SubPiecedetails, error) {
	query := psql.
		Select("*").
		From("sub_piece_details").
		Where(sq.Eq{"intl_id": Intlid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := ar.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subpieces []*domain.SubPiecedetails
	for rows.Next() {
		var subpiece domain.SubPiecedetails
		err := rows.Scan(
			&subpiece.Subid,
			&subpiece.Intlid,
			&subpiece.Identifierpieceid,
			&subpiece.Subpiececatproductcode,
			&subpiece.Hscode,
			&subpiece.Productcustomstariffhead,
			&subpiece.Productdescription,
			&subpiece.Isocodefororigincountry,
			&subpiece.Unitforsubpiecequantity,
			&subpiece.Subpiecequantitycount,
			&subpiece.Producttotalvalueasperinvoice,
			&subpiece.Isocodeforcurrency,
			&subpiece.Subpieceweight,
			&subpiece.Subpieceweightnett,
			&subpiece.Productinvoicenumber,
			&subpiece.Productinvoicedate,
			&subpiece.Statusforecommerce,
			&subpiece.Urlforecommerceconsignment,
			&subpiece.Ecommercepaymenttransactionid,
			&subpiece.Ecommerceskuno,
			&subpiece.Taxinvoicenumber,
			&subpiece.Taxinvoicedate,
			&subpiece.Serialnumberforsubpieceintaxinvoice,
			&subpiece.Valueofsubpieceaspertaxinvoice,
			&subpiece.Assessablefreeonboardvalue,
			&subpiece.Isocodeforassessablecurrency,
			&subpiece.Exchangerateforasblcurr,
			&subpiece.Assessableamount,
			&subpiece.Rateforexportduty,
			&subpiece.Exportdutyamount,
			&subpiece.Rateforcess,
			&subpiece.Cessamount,
			&subpiece.Igstrate,
			&subpiece.Igstamount,
			&subpiece.Compensationrate,
			&subpiece.Compensationamount,
			&subpiece.Detailsofletterofundertakingorbond,
			&subpiece.Modeofpayment,
			&subpiece.Paymenttransactionid,
			&subpiece.Createdon,
			&subpiece.Createdby,
			&subpiece.Updatedon,
			&subpiece.Updatedby,
			&subpiece.Authorisedon,
			&subpiece.Authorisedby,
			&subpiece.Facilityid,
			&subpiece.Ipaddress,
			&subpiece.Bookingchanneltype,
		)
		if err != nil {
			return nil, err
		}
		subpieces = append(subpieces, &subpiece)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subpieces, nil
}
*/
