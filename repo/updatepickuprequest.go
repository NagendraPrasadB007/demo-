package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type UpdatedetailsRepository struct {
	db  *DB
	log *logger.Logger
}

func NewUpdatedetailsRepository(db *DB, log *logger.Logger) *UpdatedetailsRepository {
	return &UpdatedetailsRepository{
		db,
		log,
	}
}

func (ur *UpdatedetailsRepository) IdentifyUpdate(gctx *gin.Context, Pickuprequestid int) (*domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pickupmain domain.Pickupmain
	query := psql.
		Select("pickuprequest_id", "domestic_foreign_identifier").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": Pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	ur.log.Debug("sql:", sql)
	ur.log.Debug("args:", args)

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
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

// Repo function to update pickuprequest details based on pickuprequestid
func (ur *UpdatedetailsRepository) UpdatepickuprequestdetailsDom(gctx *gin.Context, pickupmain *domain.Pickupmain, domestic *domain.Domesticarticledetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the pickuprequestid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		//return nil, nil, nil, nil, nil, nil, existsErr
		return existsErr
	}

	var count int
	if err := ur.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			//return nil, nil, nil, nil, nil, nil, errors.New("unexpected context type")
			return errors.New("unexpected context type")
		}

		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("PickuprequestID %d does not exist", pickupmain.Pickuprequestid),
		})

		//return nil, nil, nil, nil, nil, nil, errors.New("addressid does not exist")
		return errors.New("addressid does not exist")
	}

	// update api starts here

	pickupquery := psql.Update("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid}) //.
		//Suffix("RETURNING *")

	if pickupmain.Customerid != "" {
		pickupquery = pickupquery.Set("customer_id", pickupmain.Customerid)
	}
	if pickupmain.Pickupdroptype != "" {
		pickupquery = pickupquery.Set("pickup_drop_type", pickupmain.Pickupdroptype)
	}
	if pickupmain.Pickuplocation != "" {
		pickupquery = pickupquery.Set("pickup_location", pickupmain.Pickuplocation)
	}
	if pickupmain.Droplocation != "" {
		pickupquery = pickupquery.Set("drop_location", pickupmain.Droplocation)
	}
	if pickupmain.Pickupscheduleslot != "" {
		pickupquery = pickupquery.Set("pickup_schedule_slot", pickupmain.Pickupscheduleslot)
	}
	if !pickupmain.Pickupscheduledate.IsZero() {
		pickupquery = pickupquery.Set("pickup_schedule_date", pickupmain.Pickupscheduledate)
	}
	if !pickupmain.Actualpickupdatetime.IsZero() {
		pickupquery = pickupquery.Set("actual_pickup_datetime", pickupmain.Actualpickupdatetime)
	}
	if pickupmain.Pickupagentid != 0 {
		pickupquery = pickupquery.Set("pickupagent_id", pickupmain.Pickupagentid)
	}
	if pickupmain.Pickupfacilityid != "" {
		pickupquery = pickupquery.Set("pickup_facility_id", pickupmain.Pickupfacilityid)
	}
	if pickupmain.Pickupstatus != "" {
		pickupquery = pickupquery.Set("pickup_status", pickupmain.Pickupstatus)
	}
	if pickupmain.Paymentstatus != "" {
		pickupquery = pickupquery.Set("payment_status", pickupmain.Paymentstatus)
	}
	if pickupmain.Pickupaddress != "" {
		pickupquery = pickupquery.Set("pickup_address", pickupmain.Pickupaddress)
	}
	if pickupmain.Domesticforeignidentifier != "" {
		pickupquery = pickupquery.Set("domestic_foreign_identifier", pickupmain.Domesticforeignidentifier)
	}
	if pickupmain.Pickuplong != "" {
		pickupquery = pickupquery.Set("pickup_long", pickupmain.Pickuplong)
	}
	if pickupmain.Pickuplat != "" {
		pickupquery = pickupquery.Set("pickup_lat", pickupmain.Pickuplat)
	}
	if pickupmain.Pickuprequestedpincode != "" {
		pickupquery = pickupquery.Set("pickuprequestedpincode", pickupmain.Pickuprequestedpincode)
	}
	if pickupmain.Customername != "" {
		pickupquery = pickupquery.Set("customer_name", pickupmain.Customername)
	}

	//null.String handled way
	// if pickupmain.Customername.Valid && pickupmain.Customername.String != "" {
	// 	pickupquery = pickupquery.Set("customer_name", pickupmain.Customername.String)
	// }

	// if pickupmain.Customername != (sql.NullString{}) && pickupmain.Customername.Valid && pickupmain.Customername.String != "" {
	// 	pickupquery = pickupquery.Set("customer_name", pickupmain.Customername.String)
	// }

	//pickupquery = pickupquery.Set("customer_name", pickupmain.Customername)
	if pickupmain.Customermobilenumber != "" {
		pickupquery = pickupquery.Set("customer_mobilenumber", pickupmain.Customermobilenumber)
	}
	if !pickupmain.Assigneddatetime.IsZero() {
		pickupquery = pickupquery.Set("assigned_datetime", pickupmain.Assigneddatetime)
	}

	pickupmainsql, pickupmainargs, err := pickupquery.ToSql()
	if err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	ur.log.Debug("sql:", pickupmainsql)
	ur.log.Debug("args:", pickupmainargs)

	_, err = ur.db.Exec(ctx, pickupmainsql, pickupmainargs...)
	if err != nil {
		ur.log.Debug("error occured in executing query:", err.Error())
		return err
	}

	//update query for domestic_articledetails table

	domesticquery := psql.Update("domestic_articledetails").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid}) //.
		//Suffix("RETURNING *")

	if domestic.Articleid != "" {
		domesticquery = domesticquery.Set("article_id", domestic.Articleid)
	}
	if domestic.Articlestate != "" {
		domesticquery = domesticquery.Set("article_state", domestic.Articlestate)
	}
	if domestic.Articletype != "" {
		domesticquery = domesticquery.Set("article_type", domestic.Articletype)
	}
	if domestic.Articlecontent != "" {
		domesticquery = domesticquery.Set("article_content", domestic.Articlecontent)
	}
	if domestic.Articleimageid != 0 {
		domesticquery = domesticquery.Set("articleimageid", domestic.Articleimageid)
	}
	if domestic.Articlepickupcharges != 0 {
		domesticquery = domesticquery.Set("articlepickupcharges", domestic.Articlepickupcharges)
	}
	if domestic.Ispremailing {
		domesticquery = domesticquery.Set("is_premailing", domestic.Ispremailing)
	}
	if domestic.Isparcelpacking {
		domesticquery = domesticquery.Set("is_parcelpacking", domestic.Isparcelpacking)
	}
	if domestic.Customerdacpickup != "" {
		domesticquery = domesticquery.Set("customer_dac_pickup", domestic.Customerdacpickup)
	}
	if domestic.Addresstype != "" {
		domesticquery = domesticquery.Set("address_type", domestic.Addresstype)
	}
	if domestic.Bkgtransactionid != "" {
		domesticquery = domesticquery.Set("bkg_transaction_id", domestic.Bkgtransactionid)
	}
	if domestic.Originpin != 0 {
		domesticquery = domesticquery.Set("origin_pin", domestic.Originpin)
	}
	if domestic.Destinationpin != 0 {
		domesticquery = domesticquery.Set("destination_pin", domestic.Destinationpin)
	}
	if domestic.Physicalweight != 0 {
		domesticquery = domesticquery.Set("physical_weight", domestic.Physicalweight)
	}
	if domestic.Shape != "" {
		domesticquery = domesticquery.Set("shape", domestic.Shape)
	}
	if domestic.Dimensionlength != 0 {
		domesticquery = domesticquery.Set("dimension_length", domestic.Dimensionlength)
	}
	if domestic.Dimensionbreadth != 0 {
		domesticquery = domesticquery.Set("dimension_breadth", domestic.Dimensionbreadth)
	}
	if domestic.Dimensionheight != 0 {
		domesticquery = domesticquery.Set("dimension_height", domestic.Dimensionheight)
	}
	if domestic.Volumetricweight != 0 {
		domesticquery = domesticquery.Set("volumetric_weight", domestic.Volumetricweight)
	}
	if domestic.Chargedweight != 0 {
		domesticquery = domesticquery.Set("charged_weight", domestic.Chargedweight)
	}
	if domestic.Mailservicetypecode != "" {
		domesticquery = domesticquery.Set("mail_service_typecode", domestic.Mailservicetypecode)
	}
	if domestic.Bkgtype != "" {
		domesticquery = domesticquery.Set("bkg_type", domestic.Bkgtype)
	}
	if domestic.Mailform != "" {
		domesticquery = domesticquery.Set("mail_form", domestic.Mailform)
	}
	if domestic.Isprepaid {
		domesticquery = domesticquery.Set("is_prepaid", domestic.Isprepaid)
	}
	if domestic.Prepaymenttype != "" {
		domesticquery = domesticquery.Set("prepayment_type", domestic.Prepaymenttype)
	}
	if domestic.Valueofprepayment != 0 {
		domesticquery = domesticquery.Set("value_of_prepayment", domestic.Valueofprepayment)
	}
	if domestic.Vpcodflag {
		domesticquery = domesticquery.Set("vp_cod_flag", domestic.Vpcodflag)
	}
	if domestic.Valueforvpcod != 0 {
		domesticquery = domesticquery.Set("value_for_vp_cod", domestic.Valueforvpcod)
	}
	if domestic.Insuranceflag {
		domesticquery = domesticquery.Set("insurance_flag", domestic.Insuranceflag)
	}
	if domestic.Insurancetype != "" {
		domesticquery = domesticquery.Set("insurance_type", domestic.Insurancetype)
	}
	if domestic.Valueinsurance != 0 {
		domesticquery = domesticquery.Set("value_insurance", domestic.Valueinsurance)
	}
	if domestic.Acknowledgementpod {
		domesticquery = domesticquery.Set("acknowledgement_pod", domestic.Acknowledgementpod)
	}
	if domestic.Instructionsrts != "" {
		domesticquery = domesticquery.Set("instructions_rts", domestic.Instructionsrts)
	}
	if domestic.Addressrefsender != "" {
		domesticquery = domesticquery.Set("address_ref_sender", domestic.Addressrefsender)
	}
	if domestic.Addressrefreceiver != "" {
		domesticquery = domesticquery.Set("address_ref_receiver", domestic.Addressrefreceiver)
	}
	if domestic.Addressrefsenderaltaddr != "" {
		domesticquery = domesticquery.Set("address_ref_sender_alt_addr", domestic.Addressrefsenderaltaddr)
	}
	if domestic.Addressrefreceiveraktaddr != "" {
		domesticquery = domesticquery.Set("address_ref_receiver_akt_addr", domestic.Addressrefreceiveraktaddr)
	}
	if domestic.Barcodenumber != "" {
		domesticquery = domesticquery.Set("barcode_number", domestic.Barcodenumber)
	}
	if domestic.Pickupflag {
		domesticquery = domesticquery.Set("pickup_flag", domestic.Pickupflag)
	}
	if domestic.Basetariff != 0 {
		domesticquery = domesticquery.Set("base_tariff", domestic.Basetariff)
	}
	if domestic.Tax != 0 {
		domesticquery = domesticquery.Set("tax", domestic.Tax)
	}
	if domestic.Totaltariff != 0 {
		domesticquery = domesticquery.Set("total_tariff", domestic.Totaltariff)
	}
	if domestic.Modeofpayment != "" {
		domesticquery = domesticquery.Set("mode_of_payment", domestic.Modeofpayment)
	}
	if domestic.Paymenttranid != "" {
		domesticquery = domesticquery.Set("payment_tranid", domestic.Paymenttranid)
	}
	if domestic.Status != "" {
		domesticquery = domesticquery.Set("status", domestic.Status)
	}
	// if !domestic.Createdon.IsZero() {
	// 	domesticquery = domesticquery.Set("createdon", domestic.Createdon)
	// }
	if domestic.Createdby != "" {
		domesticquery = domesticquery.Set("created_by", domestic.Createdby)
	}
	// if !domestic.Updatedon.IsZero() {
	// 	domesticquery = domesticquery.Set("updatedon", domestic.Updatedon)
	// }
	if domestic.Updatedby != "" {
		domesticquery = domesticquery.Set("updated_by", domestic.Updatedby)
	}
	if !domestic.Authorisedon.IsZero() {
		domesticquery = domesticquery.Set("authorised_on", domestic.Authorisedon)
	}
	if domestic.Authorisedby != "" {
		domesticquery = domesticquery.Set("authorised_by", domestic.Authorisedby)
	}
	if domestic.Facilityid != "" {
		domesticquery = domesticquery.Set("facility_id", domestic.Facilityid)
	}
	if domestic.Reqipaddress != "" {
		domesticquery = domesticquery.Set("req_ip_address", domestic.Reqipaddress)
	}
	if domestic.Bookingchannel != "" {
		domesticquery = domesticquery.Set("booking_channel", domestic.Bookingchannel)
	}
	if domestic.Customerid != "" {
		domesticquery = domesticquery.Set("customer_id", domestic.Customerid)
	}
	if domestic.Contractnumber != "" {
		domesticquery = domesticquery.Set("contract_number", domestic.Contractnumber)
	}
	if domestic.Isparcel {
		domesticquery = domesticquery.Set("isparcel", domestic.Isparcel)
	}
	if domestic.Iscod {
		domesticquery = domesticquery.Set("iscod", domestic.Iscod)
	}

	domesticsql, domesticargs, err := domesticquery.ToSql()
	if err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	ur.log.Debug("sql:", domesticsql)
	ur.log.Debug("args:", domesticargs)

	_, err = ur.db.Exec(ctx, domesticsql, domesticargs...)
	if err != nil {
		ur.log.Debug("error occured in executing query:", err.Error())
		return err
	}

	//tariff update query

	tariffquery := psql.Update("tariff_details").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	if tariff.Articleid != "" {
		tariffquery = tariffquery.Set("article_id", tariff.Articleid)
	}
	if tariff.Totalamount != 0 {
		tariffquery = tariffquery.Set("total_amount", tariff.Totalamount)
	}
	if tariff.Pickupcharges != 0 {
		tariffquery = tariffquery.Set("pickup_charges", tariff.Pickupcharges)
	}
	if tariff.Registrationfee != 0 {
		tariffquery = tariffquery.Set("registration_fee", tariff.Registrationfee)
	}
	if tariff.Postage != 0 {
		tariffquery = tariffquery.Set("postage", tariff.Postage)
	}
	if tariff.Ackorpodfee != 0 {
		tariffquery = tariffquery.Set("ack_pod_fee", tariff.Ackorpodfee)
	}
	if tariff.Valueinsurance != 0 {
		tariffquery = tariffquery.Set("value_insurance", tariff.Valueinsurance)
	}
	if tariff.Valueforvpcod != 0 {
		tariffquery = tariffquery.Set("value_for_vp_cod", tariff.Valueforvpcod)
	}
	if tariff.Doordeliverycharges != 0 {
		tariffquery = tariffquery.Set("doordelivery_charges", tariff.Doordeliverycharges)
	}
	if tariff.Packingfee != 0 {
		tariffquery = tariffquery.Set("packing_fee", tariff.Packingfee)
	}
	if tariff.Cgst != 0 {
		tariffquery = tariffquery.Set("cgst", tariff.Cgst)
	}
	if tariff.Sgst != 0 {
		tariffquery = tariffquery.Set("sgst", tariff.Sgst)
	}
	if tariff.Othercharges != 0 {
		tariffquery = tariffquery.Set("other_charges", tariff.Othercharges)
	}

	tariffsql, tariffargs, err := tariffquery.ToSql()
	if err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	ur.log.Debug("tariffsql:", tariffsql)
	ur.log.Debug("tariffargs:", tariffargs)

	_, err = ur.db.Exec(ctx, tariffsql, tariffargs...)
	if err != nil {
		ur.log.Debug("error occured in executing query:", err.Error())
		return err
	}

	//payment

	paymentquery := psql.Update("payment_details").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	if payment.Paymenttranid != "" {
		paymentquery = paymentquery.Set("payment_tranid", payment.Paymenttranid)
	}
	if payment.Articleid != "" {
		paymentquery = paymentquery.Set("article_id", payment.Articleid)
	}
	if payment.Paymenttype != "" {
		paymentquery = paymentquery.Set("payment_type", payment.Paymenttype)
	}
	if payment.Modeofpayment != "" {
		paymentquery = paymentquery.Set("mode_of_payment", payment.Modeofpayment)
	}
	if payment.Paymentstatus != "" {
		paymentquery = paymentquery.Set("payment_status", payment.Paymentstatus)
	}
	if !payment.Paymentdatetime.IsZero() {
		paymentquery = paymentquery.Set("payment_datetime", payment.Paymentdatetime)
	}
	if payment.Paidamount != 0 {
		paymentquery = paymentquery.Set("paid_amount", payment.Paidamount)
	}

	paymentsql, paymentargs, err := paymentquery.ToSql()
	if err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	ur.log.Debug("paymentsql:", paymentsql)
	ur.log.Debug("paymentargs:", paymentargs)

	_, err = ur.db.Exec(ctx, paymentsql, paymentargs...)
	if err != nil {
		ur.log.Debug("error occured in executing query:", err.Error())
		return err
	}

	return nil

}

// Repo function to update pickuprequest details based on pickuprequestid
func (ur *UpdatedetailsRepository) UpdatepickuprequestdetailsInt(gctx *gin.Context, pickupmain *domain.Pickupmain, international *domain.Internationalarticledetails, subpiece *domain.SubPiecedetails, tariff *domain.Tariffdetails, payment *domain.Paymentdetails) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the pickuprequestid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		//return nil, nil, nil, nil, nil, nil, existsErr
		return existsErr
	}

	var count int
	if err := ur.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			//return nil, nil, nil, nil, nil, nil, errors.New("unexpected context type")
			return errors.New("unexpected context type")
		}

		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("PickuprequestID %d does not exist", pickupmain.Pickuprequestid),
		})

		//return nil, nil, nil, nil, nil, nil, errors.New("addressid does not exist")
		return errors.New("addressid does not exist")
	}

	// update api starts here

	pickupquery := psql.Update("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid}) //.

	if pickupmain.Customerid != "" {
		pickupquery = pickupquery.Set("customer_id", pickupmain.Customerid)
	}
	if pickupmain.Pickupdroptype != "" {
		pickupquery = pickupquery.Set("pickup_drop_type", pickupmain.Pickupdroptype)
	}
	if pickupmain.Pickuplocation != "" {
		pickupquery = pickupquery.Set("pickup_location", pickupmain.Pickuplocation)
	}
	if pickupmain.Droplocation != "" {
		pickupquery = pickupquery.Set("drop_location", pickupmain.Droplocation)
	}
	if pickupmain.Pickupscheduleslot != "" {
		pickupquery = pickupquery.Set("pickup_schedule_slot", pickupmain.Pickupscheduleslot)
	}
	if !pickupmain.Pickupscheduledate.IsZero() {
		pickupquery = pickupquery.Set("pickup_schedule_date", pickupmain.Pickupscheduledate)
	}
	if !pickupmain.Actualpickupdatetime.IsZero() {
		pickupquery = pickupquery.Set("actual_pickup_datetime", pickupmain.Actualpickupdatetime)
	}
	if pickupmain.Pickupagentid != 0 {
		pickupquery = pickupquery.Set("pickupagent_id", pickupmain.Pickupagentid)
	}
	if pickupmain.Pickupfacilityid != "" {
		pickupquery = pickupquery.Set("pickup_facility_id", pickupmain.Pickupfacilityid)
	}
	if pickupmain.Pickupstatus != "" {
		pickupquery = pickupquery.Set("pickup_status", pickupmain.Pickupstatus)
	}
	if pickupmain.Paymentstatus != "" {
		pickupquery = pickupquery.Set("payment_status", pickupmain.Paymentstatus)
	}
	if pickupmain.Pickupaddress != "" {
		pickupquery = pickupquery.Set("pickup_address", pickupmain.Pickupaddress)
	}
	if pickupmain.Domesticforeignidentifier != "" {
		pickupquery = pickupquery.Set("domestic_foreign_identifier", pickupmain.Domesticforeignidentifier)
	}
	if pickupmain.Pickuplong != "" {
		pickupquery = pickupquery.Set("pickup_long", pickupmain.Pickuplong)
	}
	if pickupmain.Pickuplat != "" {
		pickupquery = pickupquery.Set("pickup_lat", pickupmain.Pickuplat)
	}
	if pickupmain.Pickuprequestedpincode != "" {
		pickupquery = pickupquery.Set("pickuprequestedpincode", pickupmain.Pickuprequestedpincode)
	}
	// if pickupmain.Customername != "" {
	// 	pickupquery = pickupquery.Set("customer_name", pickupmain.Customername)
	// }
	pickupquery = pickupquery.Set("customer_name", pickupmain.Customername)
	if pickupmain.Customermobilenumber != "" {
		pickupquery = pickupquery.Set("customer_mobilenumber", pickupmain.Customermobilenumber)
	}
	if !pickupmain.Assigneddatetime.IsZero() {
		pickupquery = pickupquery.Set("assigned_datetime", pickupmain.Assigneddatetime)
	}

	pickupmainsql, pickupmainargs, err := pickupquery.ToSql()
	if err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		ur.log.Debug("error is occuring in tosql method")
		//return err
	}

	ur.log.Debug("sql:", pickupmainsql)
	ur.log.Debug("args:", pickupmainargs)

	_, err = ur.db.Exec(ctx, pickupmainsql, pickupmainargs...)
	if err != nil {
		ur.log.Debug("error is occuring during  exec")
		//return err
	}

	//international update query

	internationalquery := psql.Update("international_articledetails").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid}) //.
	//Suffix("RETURNING *")

	if international.Articleid != "" {
		internationalquery = internationalquery.Set("article_id", international.Articleid)
	}
	if international.Articlestate != "" {
		internationalquery = internationalquery.Set("article_state", international.Articlestate)
	}
	if international.Articletype != "" {
		internationalquery = internationalquery.Set("article_type", international.Articletype)
	}
	if international.Articlecontent != "" {
		internationalquery = internationalquery.Set("article_content", international.Articlecontent)
	}
	if international.Articleimageid != 0 {
		internationalquery = internationalquery.Set("articleimageid", international.Articleimageid)
	}
	if international.Articlepickupcharges != 0 {
		internationalquery = internationalquery.Set("articlepickupcharges", international.Articlepickupcharges)
	}

	//in postman if i send false it will not impact in DB
	// if international.Ispremailing {
	// 	internationalquery = internationalquery.Set("is_premailing", international.Ispremailing)
	// }

	//internationalquery = internationalquery.Set("is_premailing", international.Ispremailing)

	// if international.Ispremailing {
	// 	internationalquery = internationalquery.Set("is_premailing", true)
	// } else {
	// 	internationalquery = internationalquery.Set("is_premailing", false)
	// }

	// if international.Ispremailing {
	// 	internationalquery = internationalquery.Set("is_premailing", international.Ispremailing)
	// } else {
	// 	internationalquery = internationalquery.Set("is_premailing", international.Ispremailing)
	// }

	// if international.Isparcelpacking {
	// 	internationalquery = internationalquery.Set("is_parcelpacking", international.Isparcelpacking)
	// }
	//internationalquery = internationalquery.Set("is_parcelpacking", international.Isparcelpacking)
	// if international.Isparcelpacking {
	// 	internationalquery = internationalquery.Set("is_parcelpacking", true)
	// } else {
	// 	internationalquery = internationalquery.Set("is_parcelpacking", false)
	// }

	if international.Customerdacpickup != "" {
		internationalquery = internationalquery.Set("customer_dac_pickup", international.Customerdacpickup)
	}
	if international.Addresstype != "" {
		internationalquery = internationalquery.Set("address_type", international.Addresstype)
	}
	if international.Bkgtransactionid != "" {
		internationalquery = internationalquery.Set("bkg_transaction_id", international.Bkgtransactionid)
	}
	if international.Origincountrycode != 0 {
		internationalquery = internationalquery.Set("origin_countrycode", international.Origincountrycode)
	}
	if international.Destinationcountrycode != 0 {
		internationalquery = internationalquery.Set("destination_countrycode", international.Destinationcountrycode)
	}
	if international.Physicalweight != 0 {
		internationalquery = internationalquery.Set("physical_weight", international.Physicalweight)
	}
	if international.Mailclass != "" {
		internationalquery = internationalquery.Set("mail_class", international.Mailclass)
	}
	if international.Contenttype != "" {
		internationalquery = internationalquery.Set("content_type", international.Contenttype)
	}
	if international.Shape != "" {
		internationalquery = internationalquery.Set("shape", international.Shape)
	}
	if international.Dimensionlength != 0 {
		internationalquery = internationalquery.Set("dimension_length", international.Dimensionlength)
	}
	if international.Dimensionbreadth != 0 {
		internationalquery = internationalquery.Set("dimension_breadth", international.Dimensionbreadth)
	}
	if international.Dimensionheight != 0 {
		internationalquery = internationalquery.Set("dimension_height", international.Dimensionheight)
	}
	if international.Volumetricweight != 0 {
		internationalquery = internationalquery.Set("volumetric_weight", international.Volumetricweight)
	}
	if international.Chargedweight != 0 {
		internationalquery = internationalquery.Set("charged_weight", international.Chargedweight)
	}
	if international.Mailservicetypecode != "" {
		internationalquery = internationalquery.Set("mail_service_typecode", international.Mailservicetypecode)
	}
	if international.Bkgtype != "" {
		internationalquery = internationalquery.Set("bkg_type", international.Bkgtype)
	}
	if international.Mailform != "" {
		internationalquery = internationalquery.Set("mail_form", international.Mailform)
	}
	if international.Isprepaid {
		internationalquery = internationalquery.Set("is_prepaid", international.Isprepaid)
	}
	if international.Prepaymenttype != "" {
		internationalquery = internationalquery.Set("prepayment_type", international.Prepaymenttype)
	}
	if international.Valueofprepayment != 0 {
		internationalquery = internationalquery.Set("value_of_prepayment", international.Valueofprepayment)
	}
	if international.Vpcodflag {
		internationalquery = internationalquery.Set("vp_cod_flag", international.Vpcodflag)
	}
	if international.Valueforvpcod != 0 {
		internationalquery = internationalquery.Set("value_for_vp_cod", international.Valueforvpcod)
	}
	if international.Insuranceflag {
		internationalquery = internationalquery.Set("insurance_flag", international.Insuranceflag)
	}
	if international.Insurancetype != "" {
		internationalquery = internationalquery.Set("insurance_type", international.Insurancetype)
	}
	if international.Valueinsurance != 0 {
		internationalquery = internationalquery.Set("value_insurance", international.Valueinsurance)
	}
	if international.Acknowledgementpod {
		internationalquery = internationalquery.Set("acknowledgement_pod", international.Acknowledgementpod)
	}
	if international.Instructionsrts != "" {
		internationalquery = internationalquery.Set("instructions_rts", international.Instructionsrts)
	}
	if international.Addressrefsender != "" {
		internationalquery = internationalquery.Set("address_ref_sender", international.Addressrefsender)
	}
	if international.Addressrefreceiver != "" {
		internationalquery = internationalquery.Set("address_ref_receiver", international.Addressrefreceiver)
	}
	if international.Addressrefsenderaltaddr != "" {
		internationalquery = internationalquery.Set("address_ref_sender_alt_addr", international.Addressrefsenderaltaddr)
	}
	if international.Addressrefreceiveraktaddr != "" {
		internationalquery = internationalquery.Set("address_ref_receiver_akt_addr", international.Addressrefreceiveraktaddr)
	}
	if international.Barcodenumber != "" {
		internationalquery = internationalquery.Set("barcode_number", international.Barcodenumber)
	}
	if international.Pickupflag {
		internationalquery = internationalquery.Set("pickup_flag", international.Pickupflag)
	}
	if international.Basetariff != 0 {
		internationalquery = internationalquery.Set("base_tariff", international.Basetariff)
	}
	if international.Tax != 0 {
		internationalquery = internationalquery.Set("tax", international.Tax)
	}
	if international.Totaltariff != 0 {
		internationalquery = internationalquery.Set("total_tariff", international.Totaltariff)
	}
	if international.Modeofpayment != "" {
		internationalquery = internationalquery.Set("mode_of_payment", international.Modeofpayment)
	}
	if international.Paymenttranid != "" {
		internationalquery = internationalquery.Set("payment_tranid", international.Paymenttranid)
	}
	if international.Status != "" {
		internationalquery = internationalquery.Set("status", international.Status)
	}
	// if !international.Createdon.IsZero() {
	// 	internationalquery = internationalquery.Set("createdon", international.Createdon)
	// }
	if international.Createdby != "" {
		internationalquery = internationalquery.Set("created_by", international.Createdby)
	}
	// if !international.Updatedon.IsZero() {
	// 	internationalquery = internationalquery.Set("updatedon", international.Updatedon)
	// }
	if international.Updatedby != "" {
		internationalquery = internationalquery.Set("updated_by", international.Updatedby)
	}
	if !international.Authorisedon.IsZero() {
		internationalquery = internationalquery.Set("authorised_on", international.Authorisedon)
	}
	if international.Authorisedby != "" {
		internationalquery = internationalquery.Set("authorised_by", international.Authorisedby)
	}
	if international.Facilityid != "" {
		internationalquery = internationalquery.Set("facility_id", international.Facilityid)
	}
	if international.Reqipaddress != "" {
		internationalquery = internationalquery.Set("req_ip_address", international.Reqipaddress)
	}
	if international.Bookingchannel != "" {
		internationalquery = internationalquery.Set("booking_channel", international.Bookingchannel)
	}
	if international.Consignmentvalue != 0 {
		internationalquery = internationalquery.Set("consignment_value", international.Consignmentvalue)
	}
	if international.Mailexporttype != "" {
		internationalquery = internationalquery.Set("mail_export_type", international.Mailexporttype)
	}
	if international.Pbefilingtype != "" {
		internationalquery = internationalquery.Set("pbe_filing_type", international.Pbefilingtype)
	}
	if international.Declaration1 != "" {
		internationalquery = internationalquery.Set("declaration1", international.Declaration1)
	}
	if international.Declaration23 != "" {
		internationalquery = internationalquery.Set("declaration2_3", international.Declaration23)
	}
	if international.Declaration4 != "" {
		internationalquery = internationalquery.Set("declaration4", international.Declaration4)
	}
	if international.Selffilingcusbroker != "" {
		internationalquery = internationalquery.Set("selffiling_cusbroker", international.Selffilingcusbroker)
	}
	if international.Cusbrokerlicno != "" {
		internationalquery = internationalquery.Set("cusbroker_licno", international.Cusbrokerlicno)
	}
	if international.Cusbrokername != "" {
		internationalquery = internationalquery.Set("cusbroker_name", international.Cusbrokername)
	}
	if international.Cusbrokeraddress != "" {
		internationalquery = internationalquery.Set("cusbroker_address", international.Cusbrokeraddress)
	}
	if international.Customerid != "" {
		internationalquery = internationalquery.Set("customer_id", international.Customerid)
	}
	if international.Contractnumber != "" {
		internationalquery = internationalquery.Set("contract_number", international.Contractnumber)
	}
	if international.Gstn != "" {
		internationalquery = internationalquery.Set("gstn", international.Gstn)
	}
	if international.Ibccode != "" {
		internationalquery = internationalquery.Set("ibccode", international.Ibccode)
	}
	if international.Lut != "" {
		internationalquery = internationalquery.Set("lut", international.Lut)
	}
	if international.Adcode != "" {
		internationalquery = internationalquery.Set("adcode", international.Adcode)
	}
	if international.Isparcel {
		internationalquery = internationalquery.Set("isparcel", international.Isparcel)
	}
	if international.Iscod {
		internationalquery = internationalquery.Set("iscod", international.Iscod)
	}

	internationalsql, internationalargs, err := internationalquery.ToSql()
	if err != nil {
		ur.log.Debug("error in internationalquery tosql method")
		//return nil, nil, nil, nil, nil, nil, err
		//return err
	}

	ur.log.Debug("sql:", internationalsql)
	ur.log.Debug("args:", internationalargs)

	_, err = ur.db.Exec(ctx, internationalsql, internationalargs...)
	if err != nil {
		return err
	}

	//subpice update query

	// subpiecequery := psql.Update("sub_piece_details").
	// 	Where(sq.Eq{"international_articledetails.intl_id": pickupmain.Pickuprequestid}).
	// 	Join("international_articledetails ON sub_piece_details.intl_id = international_articledetails.intl_id")

	// Create a subquery to get the intl_id based on pickuprequest_id
	subquery := psql.Select("intl_id").From("international_articledetails").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	subsql, subargs, err := subquery.ToSql()
	if err != nil {
		//return nil, nil, nil, nil, nil, nil, err
		return err
	}

	ur.log.Debug("subsql:", subsql)
	ur.log.Debug("subargs:", subargs)

	var intlID int
	err = ur.db.QueryRow(ctx, subsql, subargs...).Scan(&intlID)
	if err != nil {
		return err
	}

	// Create the main update query for sub_piece_details
	// subpiecequery := psql.Update("sub_piece_details").
	// 	Where(sq.Eq{"intl_id": intlID})
	subpiecequery := psql.Update("sub_piece_details").
		Where(sq.Eq{"sub_id": subpiece.Subid})

	if subpiece.Identifierpieceid != 0 {
		subpiecequery = subpiecequery.Set("identifier_piece_id", subpiece.Identifierpieceid)
	}
	if subpiece.Subpiececatproductcode != "" {
		subpiecequery = subpiecequery.Set("subpiecec_cat_product_code", subpiece.Subpiececatproductcode)
	}
	if subpiece.Hscode != "" {
		subpiecequery = subpiecequery.Set("hscode", subpiece.Hscode)
	}
	if subpiece.Productcustomstariffhead != "" {
		subpiecequery = subpiecequery.Set("product_customs_tariff_head", subpiece.Productcustomstariffhead)
	}
	if subpiece.Productdescription != "" {
		subpiecequery = subpiecequery.Set("product_description", subpiece.Productdescription)
	}
	if subpiece.Isocodefororigincountry != "" {
		subpiecequery = subpiecequery.Set("iso_code_for_origin_country", subpiece.Isocodefororigincountry)
	}
	if subpiece.Unitforsubpiecequantity != "" {
		subpiecequery = subpiecequery.Set("unit_for_sub_piece_quantity", subpiece.Unitforsubpiecequantity)
	}
	if subpiece.Subpiecequantitycount != "" {
		subpiecequery = subpiecequery.Set("sub_piece_quantity_count", subpiece.Subpiecequantitycount)
	}
	if subpiece.Producttotalvalueasperinvoice != "" {
		subpiecequery = subpiecequery.Set("product_total_value_as_per_invoice", subpiece.Producttotalvalueasperinvoice)
	}
	if subpiece.Isocodeforcurrency != "" {
		subpiecequery = subpiecequery.Set("iso_code_for_currency", subpiece.Isocodeforcurrency)
	}
	if subpiece.Subpieceweight != "" {
		subpiecequery = subpiecequery.Set("sub_piece_weight", subpiece.Subpieceweight)
	}
	if subpiece.Subpieceweightnett != "" {
		subpiecequery = subpiecequery.Set("sub_piece_weight_nett", subpiece.Subpieceweightnett)
	}
	if subpiece.Productinvoicenumber != "" {
		subpiecequery = subpiecequery.Set("product_invoice_number", subpiece.Productinvoicenumber)
	}
	if !subpiece.Productinvoicedate.IsZero() {
		subpiecequery = subpiecequery.Set("product_invoice_date", subpiece.Productinvoicedate)
	}
	if subpiece.Statusforecommerce != "" {
		subpiecequery = subpiecequery.Set("status_for_ecommerce", subpiece.Statusforecommerce)
	}
	if subpiece.Urlforecommerceconsignment != "" {
		subpiecequery = subpiecequery.Set("url_for_ecommerce_consignment", subpiece.Urlforecommerceconsignment)
	}
	if subpiece.Ecommercepaymenttransactionid != "" {
		subpiecequery = subpiecequery.Set("ecommerce_payment_transaction_id", subpiece.Ecommercepaymenttransactionid)
	}
	if subpiece.Ecommerceskuno != "" {
		subpiecequery = subpiecequery.Set("ecommerce_sku_no", subpiece.Ecommerceskuno)
	}
	if subpiece.Taxinvoicenumber != "" {
		subpiecequery = subpiecequery.Set("tax_invoice_number", subpiece.Taxinvoicenumber)
	}
	if !subpiece.Taxinvoicedate.IsZero() {
		subpiecequery = subpiecequery.Set("tax_invoice_date", subpiece.Taxinvoicedate)
	}
	if subpiece.Serialnumberforsubpieceintaxinvoice != "" {
		subpiecequery = subpiecequery.Set("serial_number_for_sub_piece_in_tax_invoice", subpiece.Serialnumberforsubpieceintaxinvoice)
	}
	if subpiece.Valueofsubpieceaspertaxinvoice != 0 {
		subpiecequery = subpiecequery.Set("value_of_sub_piece_as_per_tax_invoice", subpiece.Valueofsubpieceaspertaxinvoice)
	}
	if subpiece.Assessablefreeonboardvalue != 0 {
		subpiecequery = subpiecequery.Set("assessable_free_on_board_value", subpiece.Assessablefreeonboardvalue)
	}
	if subpiece.Isocodeforassessablecurrency != "" {
		subpiecequery = subpiecequery.Set("iso_code_for_assessable_currency", subpiece.Isocodeforassessablecurrency)
	}
	if subpiece.Exchangerateforasblcurr != 0 {
		subpiecequery = subpiecequery.Set("exchange_rate_for_asbl_curr", subpiece.Exchangerateforasblcurr)
	}
	if subpiece.Assessableamount != 0 {
		subpiecequery = subpiecequery.Set("assessable_amount", subpiece.Assessableamount)
	}
	if subpiece.Rateforexportduty != 0 {
		subpiecequery = subpiecequery.Set("rate_for_export_duty", subpiece.Rateforexportduty)
	}
	if subpiece.Exportdutyamount != 0 {
		subpiecequery = subpiecequery.Set("export_duty_amount", subpiece.Exportdutyamount)
	}
	if subpiece.Rateforcess != 0 {
		subpiecequery = subpiecequery.Set("rate_for_cess", subpiece.Rateforcess)
	}
	if subpiece.Cessamount != 0 {
		subpiecequery = subpiecequery.Set("cess_amount", subpiece.Cessamount)
	}
	if subpiece.Igstrate != 0 {
		subpiecequery = subpiecequery.Set("igst_rate", subpiece.Igstrate)
	}
	if subpiece.Igstamount != 0 {
		subpiecequery = subpiecequery.Set("igst_amount", subpiece.Igstamount)
	}
	if subpiece.Compensationrate != 0 {
		subpiecequery = subpiecequery.Set("compensation_rate", subpiece.Compensationrate)
	}
	if subpiece.Compensationamount != 0 {
		subpiecequery = subpiecequery.Set("compensation_amount", subpiece.Compensationamount)
	}
	if subpiece.Detailsofletterofundertakingorbond {
		subpiecequery = subpiecequery.Set("details_of_letter_of_undertaking_or_bond", subpiece.Detailsofletterofundertakingorbond)
	}
	if subpiece.Modeofpayment != "" {
		subpiecequery = subpiecequery.Set("mode_of_payment", subpiece.Modeofpayment)
	}
	if subpiece.Paymenttransactionid != "" {
		subpiecequery = subpiecequery.Set("payment_transaction_id", subpiece.Paymenttransactionid)
	}
	// if !subpiece.Createdon.IsZero() {
	// 	subpiecequery = subpiecequery.Set("createdon", subpiece.Createdon)
	// }
	if subpiece.Createdby != "" {
		subpiecequery = subpiecequery.Set("created_by", subpiece.Createdby)
	}
	// if !subpiece.Updatedon.IsZero() {
	// 	subpiecequery = subpiecequery.Set("updatedon", subpiece.Updatedon)
	// }
	if subpiece.Updatedby != "" {
		subpiecequery = subpiecequery.Set("updated_by", subpiece.Updatedby)
	}
	if !subpiece.Authorisedon.IsZero() {
		subpiecequery = subpiecequery.Set("authorised_on", subpiece.Authorisedon)
	}
	if subpiece.Authorisedby != "" {
		subpiecequery = subpiecequery.Set("authorised_by", subpiece.Authorisedby)
	}
	if subpiece.Facilityid != "" {
		subpiecequery = subpiecequery.Set("facility_id", subpiece.Facilityid)
	}
	if subpiece.Ipaddress != "" {
		subpiecequery = subpiecequery.Set("ip_address", subpiece.Ipaddress)
	}
	if subpiece.Bookingchanneltype != "" {
		subpiecequery = subpiecequery.Set("booking_channel_type", subpiece.Bookingchanneltype)
	}

	subpiecesql, subpieceargs, err := subpiecequery.ToSql()
	if err != nil {
		ur.log.Debug(err)
		//return err
	}

	ur.log.Debug("subpiecesql:", subpiecesql)
	ur.log.Debug("subpieceargs:", subpieceargs)

	_, err = ur.db.Exec(ctx, subpiecesql, subpieceargs...)
	if err != nil {
		return err
	}

	//tariff update query

	tariffquery := psql.Update("tariff_details").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	if tariff.Articleid != "" {
		tariffquery = tariffquery.Set("article_id", tariff.Articleid)
	}
	if tariff.Totalamount != 0 {
		tariffquery = tariffquery.Set("total_amount", tariff.Totalamount)
	}
	if tariff.Pickupcharges != 0 {
		tariffquery = tariffquery.Set("pickup_charges", tariff.Pickupcharges)
	}
	if tariff.Registrationfee != 0 {
		tariffquery = tariffquery.Set("registration_fee", tariff.Registrationfee)
	}
	if tariff.Postage != 0 {
		tariffquery = tariffquery.Set("postage", tariff.Postage)
	}
	if tariff.Ackorpodfee != 0 {
		tariffquery = tariffquery.Set("ack_pod_fee", tariff.Ackorpodfee)
	}
	if tariff.Valueinsurance != 0 {
		tariffquery = tariffquery.Set("value_insurance", tariff.Valueinsurance)
	}
	if tariff.Valueforvpcod != 0 {
		tariffquery = tariffquery.Set("value_for_vp_cod", tariff.Valueforvpcod)
	}
	if tariff.Doordeliverycharges != 0 {
		tariffquery = tariffquery.Set("doordelivery_charges", tariff.Doordeliverycharges)
	}
	if tariff.Packingfee != 0 {
		tariffquery = tariffquery.Set("packing_fee", tariff.Packingfee)
	}
	if tariff.Cgst != 0 {
		tariffquery = tariffquery.Set("cgst", tariff.Cgst)
	}
	if tariff.Sgst != 0 {
		tariffquery = tariffquery.Set("sgst", tariff.Sgst)
	}
	if tariff.Othercharges != 0 {
		tariffquery = tariffquery.Set("other_charges", tariff.Othercharges)
	}

	tariffsql, tariffargs, err := tariffquery.ToSql()
	if err != nil {
		ur.log.Debug(err)
		//return err
	}

	ur.log.Debug("tariffsql:", tariffsql)
	ur.log.Debug("tariffargs:", tariffargs)

	_, err = ur.db.Exec(ctx, tariffsql, tariffargs...)
	if err != nil {
		return err
	}

	//payment

	paymentquery := psql.Update("payment_details").
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	if payment.Paymenttranid != "" {
		paymentquery = paymentquery.Set("payment_tranid", payment.Paymenttranid)
	}
	if payment.Articleid != "" {
		paymentquery = paymentquery.Set("article_id", payment.Articleid)
	}
	if payment.Paymenttype != "" {
		paymentquery = paymentquery.Set("payment_type", payment.Paymenttype)
	}
	if payment.Modeofpayment != "" {
		paymentquery = paymentquery.Set("mode_of_payment", payment.Modeofpayment)
	}
	if payment.Paymentstatus != "" {
		paymentquery = paymentquery.Set("payment_status", payment.Paymentstatus)
	}
	if !payment.Paymentdatetime.IsZero() {
		paymentquery = paymentquery.Set("payment_datetime", payment.Paymentdatetime)
	}
	if payment.Paidamount != 0 {
		paymentquery = paymentquery.Set("paid_amount", payment.Paidamount)
	}

	paymentsql, paymentargs, err := paymentquery.ToSql()
	if err != nil {
		ur.log.Debug(err)
		//return err
	}

	ur.log.Debug("paymentsql:", paymentsql)
	ur.log.Debug("paymentargs:", paymentargs)

	_, err = ur.db.Exec(ctx, paymentsql, paymentargs...)
	if err != nil {
		return err
	}

	return nil

}

/*

func (ur *UpdatedetailsRepository) GetAssociatedSubIDs(ctx context.Context, intlID int) ([]int, error) {

	// Use psql.Select to build the SELECT query
	selectQuery := psql.Select("sub_id").
		From("sub_piece_details").
		Where(sq.Eq{"intl_id": intlID})

	// Get the SQL query string and arguments
	selectSQL, selectArgs, err := selectQuery.ToSql()
	if err != nil {
		// Handle the error
		return nil, err
	}

	// Execute the query and retrieve rows
	rows, err := ur.db.Query(ctx, selectSQL, selectArgs...)
	if err != nil {
		// Handle the error
		return nil, err
	}
	defer rows.Close()

	var subIDs []int
	for rows.Next() {
		var subID int
		if err := rows.Scan(&subID); err != nil {
			return nil, err
		}
		subIDs = append(subIDs, subID)
	}

	return subIDs, nil
}
*/
