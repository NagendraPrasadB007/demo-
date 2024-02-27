package repository

import (
	"context"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var SerialPickupRequestID int
var SerialDomID int
var SerialIntlID int
var SerialSubID int

var SerialSubIDs []int

type RaisepickuprequestRepository struct {
	db  *DB
	log *logger.Logger
}

func NewRaisepickuprequestRepository(db *DB, log *logger.Logger) *RaisepickuprequestRepository {
	return &RaisepickuprequestRepository{
		db,
		log,
	}
}

// Handler function to raise a pickuprequest ( Domestic )
func (rr *RaisepickuprequestRepository) CreatingnewrequestDom(gctx *gin.Context, pickupmain *domain.Pickupmain, domestic *domain.Domesticarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) (*domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//for rollback starting this should be added and queryrow and in the end also
	txOptions := pgx.TxOptions{}
	tx, err := rr.db.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}

	//pickupmain query
	pickupmainquery := psql.Insert("pickup_main").
		Columns("customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime").
		Values(pickupmain.Customerid, pickupmain.Pickupdroptype, pickupmain.Pickuplocation, pickupmain.Droplocation, pickupmain.Pickupscheduleslot, pickupmain.Pickupscheduledate, pickupmain.Actualpickupdatetime, pickupmain.Pickupagentid, pickupmain.Pickupfacilityid, pickupmain.Pickupstatus, pickupmain.Paymentstatus, pickupmain.Pickupaddress, pickupmain.Domesticforeignidentifier, pickupmain.Pickuplong, pickupmain.Pickuplat, pickupmain.Pickuprequestedpincode, pickupmain.Customername, pickupmain.Customermobilenumber, pickupmain.Assigneddatetime).Suffix("RETURNING pickuprequest_id")

	sql, args, err := pickupmainquery.ToSql()
	if err != nil {
		return nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&SerialPickupRequestID,
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
		// &pickupmain.Domesticforeignidentifier,
		// &pickupmain.Pickuplong,
		// &pickupmain.Pickuplat,
		// &pickupmain.Modifieddatetime,
		// &pickupmain.Pickuprequestedpincode,
	)
	if err != nil {
		return nil, err
	}

	// domestic articledetails

	domesticquery := psql.Insert("domestic_articledetails").
		Columns("pickuprequest_id", "article_id", "article_state", "article_type", "article_content", "articleimageid", "articlepickupcharges", "is_premailing", "is_parcelpacking", "customer_dac_pickup", "address_type", "bkg_transaction_id", "origin_pin", "destination_pin", "physical_weight", "shape", "dimension_length", "dimension_breadth", "dimension_height", "volumetric_weight", "charged_weight", "mail_service_typecode", "bkg_type", "mail_form", "is_prepaid", "prepayment_type", "value_of_prepayment", "vp_cod_flag", "value_for_vp_cod", "insurance_flag", "insurance_type", "value_insurance", "acknowledgement_pod", "instructions_rts", "address_ref_sender", "address_ref_receiver", "address_ref_sender_alt_addr", "address_ref_receiver_akt_addr", "barcode_number", "pickup_flag, base_tariff", "tax", "total_tariff", "mode_of_payment", "payment_tranid", "status", "created_by", "updated_by", "authorised_on", "authorised_by", "facility_id", "req_ip_address", "booking_channel", "customer_id", "contract_number", "isparcel", "iscod").
		Values(SerialPickupRequestID, domestic.Articleid, domestic.Articlestate, domestic.Articletype,
			domestic.Articlecontent, domestic.Articleimageid, domestic.Articlepickupcharges, domestic.Ispremailing,
			domestic.Isparcelpacking, domestic.Customerdacpickup, domestic.Addresstype, domestic.Bkgtransactionid,
			domestic.Originpin, domestic.Destinationpin, domestic.Physicalweight, domestic.Shape,
			domestic.Dimensionlength, domestic.Dimensionbreadth, domestic.Dimensionheight, domestic.Volumetricweight,
			domestic.Chargedweight, domestic.Mailservicetypecode, domestic.Bkgtype, domestic.Mailform, domestic.Isprepaid, domestic.Prepaymenttype, domestic.Valueofprepayment, domestic.Vpcodflag, domestic.Valueforvpcod,
			domestic.Insuranceflag, domestic.Insurancetype, domestic.Valueinsurance, domestic.Acknowledgementpod,
			domestic.Instructionsrts, domestic.Addressrefsender, domestic.Addressrefreceiver,
			domestic.Addressrefsenderaltaddr, domestic.Addressrefreceiveraktaddr, domestic.Barcodenumber,
			domestic.Pickupflag, domestic.Basetariff, domestic.Tax, domestic.Totaltariff, domestic.Modeofpayment,
			domestic.Paymenttranid, domestic.Status, domestic.Createdby, domestic.Updatedby, domestic.Authorisedon,
			domestic.Authorisedby, domestic.Facilityid, domestic.Reqipaddress, domestic.Bookingchannel,
			domestic.Customerid, domestic.Contractnumber, domestic.Isparcel, domestic.Iscod,
		).Suffix("RETURNING dom_id")

	sql, args, err = domesticquery.ToSql()
	if err != nil {
		return nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(&SerialDomID)
	err = tx.QueryRow(ctx, sql, args...).Scan(&SerialDomID)
	if err != nil {
		return nil, err
	}

	//Tariff details table

	tariffquery := psql.Insert("tariff_details").
		Columns("pickuprequest_id", "article_id", "total_amount", "pickup_charges", "registration_fee", "postage", "ack_pod_fee", "value_insurance", "value_for_vp_cod", "doordelivery_charges", "packing_fee", "cgst", "sgst", "other_charges").
		Values(
			SerialPickupRequestID, tariff.Articleid, tariff.Totalamount, tariff.Pickupcharges, tariff.Registrationfee, tariff.Postage, tariff.Ackorpodfee, tariff.Valueinsurance, tariff.Valueforvpcod, tariff.Doordeliverycharges, tariff.Packingfee, tariff.Cgst, tariff.Sgst, tariff.Othercharges,
		).Suffix(`RETURNING "pickuprequest_id","article_id","total_amount","pickup_charges","registration_fee","postage","ack_pod_fee","value_insurance","value_for_vp_cod","doordelivery_charges","packing_fee","cgst","sgst","other_charges","tariff_id"`)
	sql, args, err = tariffquery.ToSql()
	if err != nil {
		return nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&SerialPickupRequestID,
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
	)
	if err != nil {
		return nil, err
	}

	//payment details

	paymentquery := psql.Insert("payment_details").
		Columns("payment_tranid", "pickuprequest_id", "article_id", "payment_type", "mode_of_payment", "payment_status", "payment_datetime", "paid_amount").Values(payment.Paymenttranid, SerialPickupRequestID, payment.Articleid, payment.Paymenttype, payment.Modeofpayment, payment.Paymentstatus, payment.Paymentdatetime, payment.Paidamount).Suffix(`RETURNING "payment_tranid","pickuprequest_id","article_id","payment_type","mode_of_payment","payment_status","payment_datetime","paid_amount","payment_id"`)
	sql, args, err = paymentquery.ToSql()
	if err != nil {
		return nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&payment.Paymenttranid,
		&SerialPickupRequestID,
		&payment.Articleid,
		&payment.Paymenttype,
		&payment.Modeofpayment,
		&payment.Paymentstatus,
		&payment.Paymentdatetime,
		&payment.Paidamount,
		&payment.Paymentid,
	)
	if err != nil {
		return nil, err
	}

	//rollback
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	// If all queries are successful, commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return pickupmain, nil

}

// INTERNATIONAL
// Repo function to raise pickuprequest ( international )
func (rr *RaisepickuprequestRepository) CreatingnewrequestInt(gctx *gin.Context, pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) (*domain.Pickupmain, *domain.Internationalarticledetails, *domain.Tariffdetails, *domain.Paymentdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//for rollback starting this should be added and queryrow and in the end also
	txOptions := pgx.TxOptions{}
	tx, err := rr.db.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	//pickupmain query
	pickupmainquery := psql.Insert("pickup_main").
		Columns("customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime").
		Values(pickupmain.Customerid, pickupmain.Pickupdroptype, pickupmain.Pickuplocation, pickupmain.Droplocation, pickupmain.Pickupscheduleslot, pickupmain.Pickupscheduledate, pickupmain.Actualpickupdatetime, pickupmain.Pickupagentid, pickupmain.Pickupfacilityid, pickupmain.Pickupstatus, pickupmain.Paymentstatus, pickupmain.Pickupaddress, pickupmain.Domesticforeignidentifier, pickupmain.Pickuplong, pickupmain.Pickuplat, pickupmain.Pickuprequestedpincode, pickupmain.Customername, pickupmain.Customermobilenumber, pickupmain.Assigneddatetime).Suffix("RETURNING pickuprequest_id")

	sql, args, err := pickupmainquery.ToSql()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&SerialPickupRequestID,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	//international
	internationalquery := psql.Insert("international_articledetails").
		Columns("pickuprequest_id", "article_id", "article_state", "article_type", "article_content", "articleimageid", "articlepickupcharges", "is_premailing", "is_parcelpacking", "customer_dac_pickup", "address_type", "bkg_transaction_id", "origin_countrycode", "destination_countrycode", "physical_weight", "mail_class", "content_type", "shape", "dimension_length", "dimension_breadth", "dimension_height", "volumetric_weight", "charged_weight", "mail_service_typecode", "bkg_type", "mail_form", "is_prepaid", "prepayment_type", "value_of_prepayment", "vp_cod_flag", "value_for_vp_cod", "insurance_flag", "insurance_type", "value_insurance", "acknowledgement_pod", "instructions_rts", "address_ref_sender", "address_ref_receiver", "address_ref_sender_alt_addr", "address_ref_receiver_akt_addr", "barcode_number", "pickup_flag", "base_tariff", "tax", "total_tariff", "mode_of_payment", "payment_tranid", "status", "created_by", "updated_by", "authorised_on", "authorised_by", "facility_id", "req_ip_address", "booking_channel", "consignment_value", "mail_export_type", "pbe_filing_type", "declaration1", "declaration2_3", "declaration4", "selffiling_cusbroker", "cusbroker_licno", "cusbroker_name", "cusbroker_address", "customer_id", "contract_number", "gstn", "ibccode", "lut", "adcode", "isparcel", "iscod").
		Values(
			SerialPickupRequestID, international.Articleid, international.Articlestate, international.Articletype,
			international.Articlecontent, international.Articleimageid, international.Articlepickupcharges,
			international.Ispremailing, international.Isparcelpacking, international.Customerdacpickup,
			international.Addresstype, international.Bkgtransactionid, international.Origincountrycode,
			international.Destinationcountrycode, international.Physicalweight, international.Mailclass,
			international.Contenttype, international.Shape, international.Dimensionlength, international.Dimensionbreadth, international.Dimensionheight, international.Volumetricweight, international.Chargedweight, international.Mailservicetypecode, international.Bkgtype, international.Mailform, international.Isprepaid, international.Prepaymenttype, international.Valueofprepayment, international.Vpcodflag, international.Valueforvpcod, international.Insuranceflag, international.Insurancetype,
			international.Valueinsurance, international.Acknowledgementpod, international.Instructionsrts,
			international.Addressrefsender, international.Addressrefreceiver, international.Addressrefsenderaltaddr, international.Addressrefreceiveraktaddr, international.Barcodenumber, international.Pickupflag, international.Basetariff, international.Tax, international.Totaltariff, international.Modeofpayment, international.Paymenttranid, international.Status, international.Createdby, international.Updatedby, international.Authorisedon, international.Authorisedby, international.Facilityid, international.Reqipaddress, international.Bookingchannel, international.Consignmentvalue, international.Mailexporttype, international.Pbefilingtype, international.Declaration1, international.Declaration23, international.Declaration4, international.Selffilingcusbroker, international.Cusbrokerlicno, international.Cusbrokername, international.Cusbrokeraddress, international.Customerid, international.Contractnumber, international.Gstn, international.Ibccode, international.Lut, international.Adcode, international.Isparcel, international.Iscod,
		).Suffix("RETURNING intl_id")
	sql, args, err = internationalquery.ToSql()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(&SerialIntlID)
	err = tx.QueryRow(ctx, sql, args...).Scan(&SerialIntlID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	for _, subpiece := range international.SubPieces {
		subpiecequery := psql.Insert("sub_piece_details").
			Columns("intl_id", "identifier_piece_id", "subpiecec_cat_product_code", "hscode", "product_customs_tariff_head", "product_description", "iso_code_for_origin_country", "unit_for_sub_piece_quantity", "sub_piece_quantity_count", "product_total_value_as_per_invoice", "iso_code_for_currency", "sub_piece_weight", "sub_piece_weight_nett", "product_invoice_number", "product_invoice_date", "status_for_ecommerce", "url_for_ecommerce_consignment", "ecommerce_payment_transaction_id", "ecommerce_sku_no", "tax_invoice_number", "tax_invoice_date", "serial_number_for_sub_piece_in_tax_invoice", "value_of_sub_piece_as_per_tax_invoice", "assessable_free_on_board_value", "iso_code_for_assessable_currency", "exchange_rate_for_asbl_curr", "assessable_amount", "rate_for_export_duty", "export_duty_amount", "rate_for_cess", "cess_amount", "igst_rate", "igst_amount", "compensation_rate", "compensation_amount", "details_of_letter_of_undertaking_or_bond", "mode_of_payment", "payment_transaction_id", "created_by", "updated_by", "authorised_on", "authorised_by", "facility_id", "ip_address", "booking_channel_type").
			Values(
				SerialIntlID, subpiece.Identifierpieceid, subpiece.Subpiececatproductcode, subpiece.Hscode,
				subpiece.Productcustomstariffhead, subpiece.Productdescription, subpiece.Isocodefororigincountry,
				subpiece.Unitforsubpiecequantity, subpiece.Subpiecequantitycount, subpiece.Producttotalvalueasperinvoice,
				subpiece.Isocodeforcurrency, subpiece.Subpieceweight, subpiece.Subpieceweightnett,
				subpiece.Productinvoicenumber, subpiece.Productinvoicedate, subpiece.Statusforecommerce,
				subpiece.Urlforecommerceconsignment, subpiece.Ecommercepaymenttransactionid, subpiece.Ecommerceskuno,
				subpiece.Taxinvoicenumber, subpiece.Taxinvoicedate, subpiece.Serialnumberforsubpieceintaxinvoice,
				subpiece.Valueofsubpieceaspertaxinvoice, subpiece.Assessablefreeonboardvalue,
				subpiece.Isocodeforassessablecurrency, subpiece.Exchangerateforasblcurr, subpiece.Assessableamount,
				subpiece.Rateforexportduty, subpiece.Exportdutyamount, subpiece.Rateforcess, subpiece.Cessamount,
				subpiece.Igstrate, subpiece.Igstamount, subpiece.Compensationrate, subpiece.Compensationamount,
				subpiece.Detailsofletterofundertakingorbond, subpiece.Modeofpayment, subpiece.Paymenttransactionid,
				subpiece.Createdby, subpiece.Updatedby, subpiece.Authorisedon, subpiece.Authorisedby, subpiece.Facilityid,
				subpiece.Ipaddress, subpiece.Bookingchanneltype,
			).Suffix("RETURNING sub_id")
		sql, args, err = subpiecequery.ToSql()
		if err != nil {
			return nil, nil, nil, nil, err
		}
		rr.log.Debug("sql:", sql)
		rr.log.Debug("args:", args)

		//err = rr.db.QueryRow(ctx, sql, args...).Scan(&SerialSubID)
		err = tx.QueryRow(ctx, sql, args...).Scan(&SerialSubID)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		//for getting subid in repsonse uncomment below line
		SerialSubIDs = append(SerialSubIDs, SerialSubID)
	}

	/*
		//subpiece details

		subpiecequery := psql.Insert("sub_piece_details").
			Columns("intl_id", "identifier_piece_id", "subpiecec_cat_product_code", "hscode", "product_customs_tariff_head", "product_description", "iso_code_for_origin_country", "unit_for_sub_piece_quantity", "sub_piece_quantity_count", "product_total_value_as_per_invoice", "iso_code_for_currency", "sub_piece_weight", "sub_piece_weight_nett", "product_invoice_number", "product_invoice_date", "status_for_ecommerce", "url_for_ecommerce_consignment", "ecommerce_payment_transaction_id", "ecommerce_sku_no", "tax_invoice_number", "tax_invoice_date", "serial_number_for_sub_piece_in_tax_invoice", "value_of_sub_piece_as_per_tax_invoice", "assessable_free_on_board_value", "iso_code_for_assessable_currency", "exchange_rate_for_asbl_curr", "assessable_amount", "rate_for_export_duty", "export_duty_amount", "rate_for_cess", "cess_amount", "igst_rate", "igst_amount", "compensation_rate", "compensation_amount", "details_of_letter_of_undertaking_or_bond", "mode_of_payment", "payment_transaction_id", "created_by", "updated_by", "authorised_on", "authorised_by", "facility_id", "ip_address", "booking_channel_type").
			Values(
				SerialIntlID, subpiece.Identifierpieceid, subpiece.Subpiececatproductcode, subpiece.Hscode,
				subpiece.Productcustomstariffhead, subpiece.Productdescription, subpiece.Isocodefororigincountry,
				subpiece.Unitforsubpiecequantity, subpiece.Subpiecequantitycount, subpiece.Producttotalvalueasperinvoice,
				subpiece.Isocodeforcurrency, subpiece.Subpieceweight, subpiece.Subpieceweightnett,
				subpiece.Productinvoicenumber, subpiece.Productinvoicedate, subpiece.Statusforecommerce,
				subpiece.Urlforecommerceconsignment, subpiece.Ecommercepaymenttransactionid, subpiece.Ecommerceskuno,
				subpiece.Taxinvoicenumber, subpiece.Taxinvoicedate, subpiece.Serialnumberforsubpieceintaxinvoice,
				subpiece.Valueofsubpieceaspertaxinvoice, subpiece.Assessablefreeonboardvalue,
				subpiece.Isocodeforassessablecurrency, subpiece.Exchangerateforasblcurr, subpiece.Assessableamount,
				subpiece.Rateforexportduty, subpiece.Exportdutyamount, subpiece.Rateforcess, subpiece.Cessamount,
				subpiece.Igstrate, subpiece.Igstamount, subpiece.Compensationrate, subpiece.Compensationamount,
				subpiece.Detailsofletterofundertakingorbond, subpiece.Modeofpayment, subpiece.Paymenttransactionid,
				subpiece.Createdby, subpiece.Updatedby, subpiece.Authorisedon, subpiece.Authorisedby, subpiece.Facilityid,
				subpiece.Ipaddress, subpiece.Bookingchanneltype,
			).Suffix("RETURNING sub_id")
		sql, args, err = subpiecequery.ToSql()
		if err != nil {
			return nil, err
		}
		rr.log.Debug("sql:", sql)
		rr.log.Debug("args:", args)

		err = rr.db.QueryRow(ctx, sql, args...).Scan(&SerialSubID)
		if err != nil {
			return nil, err
		}
	*/

	//Tariff details table

	tariffquery := psql.Insert("tariff_details").
		Columns("pickuprequest_id", "article_id", "total_amount", "pickup_charges", "registration_fee", "postage", "ack_pod_fee", "value_insurance", "value_for_vp_cod", "doordelivery_charges", "packing_fee", "cgst", "sgst", "other_charges").
		Values(
			SerialPickupRequestID, tariff.Articleid, tariff.Totalamount, tariff.Pickupcharges, tariff.Registrationfee, tariff.Postage, tariff.Ackorpodfee, tariff.Valueinsurance, tariff.Valueforvpcod, tariff.Doordeliverycharges, tariff.Packingfee, tariff.Cgst, tariff.Sgst, tariff.Othercharges,
		).Suffix(`RETURNING "pickuprequest_id","article_id","total_amount","pickup_charges","registration_fee","postage","ack_pod_fee","value_insurance","value_for_vp_cod","doordelivery_charges","packing_fee","cgst","sgst","other_charges","tariff_id"`)
	sql, args, err = tariffquery.ToSql()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&SerialPickupRequestID,
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
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// row, err := rr.db.Query(ctx, sql, args...)
	// if err != nil {
	// 	return nil, err
	// }
	// u, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Tariffdetails])
	// if err != nil {
	// 	return nil, err
	// }

	//payment details

	paymentquery := psql.Insert("payment_details").
		Columns("payment_tranid", "pickuprequest_id", "article_id", "payment_type", "mode_of_payment", "payment_status", "payment_datetime", "paid_amount").Values(payment.Paymenttranid, SerialPickupRequestID, payment.Articleid, payment.Paymenttype, payment.Modeofpayment, payment.Paymentstatus, payment.Paymentdatetime, payment.Paidamount).Suffix(`RETURNING "payment_tranid","pickuprequest_id","article_id","payment_type","mode_of_payment","payment_status","payment_datetime","paid_amount","payment_id"`)
	sql, args, err = paymentquery.ToSql()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rr.log.Debug("sql:", sql)
	rr.log.Debug("args:", args)

	//err = rr.db.QueryRow(ctx, sql, args...).Scan(
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&payment.Paymenttranid,
		&SerialPickupRequestID,
		&payment.Articleid,
		&payment.Paymenttype,
		&payment.Modeofpayment,
		&payment.Paymentstatus,
		payment.Paymentdatetime,
		&payment.Paidamount,
		&payment.Paymentid,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	//rollback
	if err != nil {
		tx.Rollback(ctx)
		return nil, nil, nil, nil, err
	}

	// If all queries are successful, commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return pickupmain, international, tariff, payment, nil
}

// func GetNextSubID(db *sql.DB) (int, error) {
// 	var subid int

// 	// Execute a SQL query to get the next subid from a sequence named 'sub_id'
// 	query := "SELECT nextval('sub_id')"

// 	err := db.QueryRow(query).Scan(&subid)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return subid, nil
// }
