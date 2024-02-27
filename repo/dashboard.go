package repository

import (
	"context"
	"errors"
	"log"
	"net/http"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type DashboardRepository struct {
	db  *DB
	log *logger.Logger
}

func NewDashboardRepository(db *DB, log *logger.Logger) *DashboardRepository {
	return &DashboardRepository{
		db,
		log,
	}
}

type Counts struct {
	UnassignedCount int
	AssignedCount   int
	PickedupCount   int
	CancelledCount  int
}

// Repo function to get the count of Unassigned , Assigned , pickedup, cancelled , total sum
func (dr *DashboardRepository) GetCountByID(gctx *gin.Context, Pickupfacilityid string) (Counts, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// check wheather facility id exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickup_facility_id": Pickupfacilityid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return Counts{}, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return Counts{}, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return Counts{}, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "FacilityId " + Pickupfacilityid + " does not exist",
		})
		return Counts{}, errors.New("Facilityid does not exist")
	}

	//count sapi starts here

	//today := time.Now().Format("2006-01-02")
	query := psql.Select(`COUNT(CASE WHEN pickup_status = 'Unassigned' THEN 1 END) AS unassigned_count,
	COUNT(CASE WHEN pickup_status = 'Assigned' THEN 1 END) AS assigned_count,
	COUNT(CASE WHEN pickup_status = 'Pickedup' THEN 1 END) AS pickedup_count,
	COUNT(CASE WHEN pickup_status = 'Cancelled' THEN 1 END) AS cancelled_count`).
		From("pickup_main").
		Where(sq.Eq{"pickup_facility_id": Pickupfacilityid})
		//.
		//Where(sq.Expr("DATE_TRUNC('day', created_datetime) = ?", today))

	sql, args, err := query.ToSql()
	if err != nil {
		return Counts{}, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	var counts Counts
	err = dr.db.QueryRow(ctx, sql, args...).Scan(
		&counts.UnassignedCount,
		&counts.AssignedCount,
		&counts.PickedupCount,
		&counts.CancelledCount,
	)
	if err != nil {
		log.Fatal(err)
	}

	return counts, nil

}

// Repo function to get all Assigned pickuprequest based on facilityid
func (dr *DashboardRepository) GetAssignedlistByID(gctx *gin.Context, Pickupfacilityid string) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// check wheather facility id exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickup_facility_id": Pickupfacilityid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "FacilityId " + Pickupfacilityid + " does not exist",
		})
		return nil, errors.New("facilityid does not exist")
	}

	//get api starts here

	query := psql.Select("pickuprequest_id").
		From("pickup_main").
		Where(squirrel.Eq{"pickup_facility_id": Pickupfacilityid, "pickup_status": "Assigned"})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	//starts here for fetching multiple records
	rows, err := dr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var pickupRequestIDs []int

	for rows.Next() {
		var pickupRequestID int
		if err := rows.Scan(&pickupRequestID); err != nil {
			return nil, err
		}
		pickupRequestIDs = append(pickupRequestIDs, pickupRequestID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pickupRequestIDs, nil

}

// calling repo function to get all unassigned pickuprequestid
func (dr *DashboardRepository) GetUnAssignedlistByID(gctx *gin.Context, Pickupfacilityid string) ([]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check wheather facility id exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickup_facility_id": Pickupfacilityid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "FacilityId " + Pickupfacilityid + " does not exist",
		})
		return nil, errors.New("Facilityid does not exist")
	}

	//get api starts here

	query := psql.Select("pickuprequest_id").
		From("pickup_main").
		Where(squirrel.Eq{"pickup_facility_id": Pickupfacilityid, "pickup_status": "Unassigned"})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	//starts here for fetching multiple records
	rows, err := dr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var pickupRequestIDs []int

	for rows.Next() {
		var pickupRequestID int
		if err := rows.Scan(&pickupRequestID); err != nil {
			return nil, err
		}
		pickupRequestIDs = append(pickupRequestIDs, pickupRequestID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pickupRequestIDs, nil

}

func (dr *DashboardRepository) Identify(gctx *gin.Context, Pickuprequestid int) (*domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var pickupmain domain.Pickupmain
	query := psql.
		Select("pickuprequest_id", "customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "created_datetime", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "modified_datetime", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": Pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	err = dr.db.QueryRow(ctx, sql, args...).Scan(
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
	)

	return &pickupmain, err

}

// calling repo function to get individual unassigned pickuprequestid details ( Domestic)
func (dr *DashboardRepository) GetcompletedetailsDom(gctx *gin.Context, pickuprequestid int) (*domain.Pickupmain, *domain.Domesticarticledetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pickupmain domain.Pickupmain
	var domestic domain.Domesticarticledetails

	dr.log.Debug("inside domestic before query")

	query := psql.
		Select("pm.pickuprequest_id", "pm.pickup_schedule_slot", "pm.pickup_schedule_date", "pm.pickup_address", "pm.customer_name", "pm.customer_mobilenumber", "d.physical_weight", "d.volumetric_weight").
		From("pickup_main pm").
		Join("domestic_articledetails d ON pm.pickuprequest_id = d.pickuprequest_id").
		Where(sq.Eq{"pm.pickuprequest_id": pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	err = dr.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		&pickupmain.Pickupscheduleslot,
		&pickupmain.Pickupscheduledate,
		&pickupmain.Pickupaddress,
		&pickupmain.Customername,
		&pickupmain.Customermobilenumber,

		&domestic.Physicalweight,
		&domestic.Volumetricweight,
	)
	if err != nil {
		return nil, nil, err
	}

	return &pickupmain, &domestic, nil

}

// calling repo function to get individual unassigned pickuprequestid details ( International )
func (dr *DashboardRepository) GetcompletedetailsInt(gctx *gin.Context, pickuprequestid int) (*domain.Pickupmain, *domain.Internationalarticledetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pickupmain domain.Pickupmain
	var international domain.Internationalarticledetails

	query := psql.
		Select("pm.pickuprequest_id", "pm.pickup_schedule_slot", "pm.pickup_schedule_date", "pm.pickup_address", "pm.customer_name", "pm.customer_mobilenumber", "d.physical_weight", "d.volumetric_weight").
		From("pickup_main pm").
		Join("international_articledetails d ON pm.pickuprequest_id = d.pickuprequest_id").
		Where(sq.Eq{"pm.pickuprequest_id": pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	err = dr.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		&pickupmain.Pickupscheduleslot,
		&pickupmain.Pickupscheduledate,
		&pickupmain.Pickupaddress,
		&pickupmain.Customername,
		&pickupmain.Customermobilenumber,

		&international.Physicalweight,
		&international.Volumetricweight,
	)
	if err != nil {
		return nil, nil, err
	}

	return &pickupmain, &international, nil

}

// calling repo function to get list of pickuprequestid based on facilityid
func (dr *DashboardRepository) GetPickuprequestlistByFacilityid(gctx *gin.Context, Pickupfacilityid string) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check wheather facility id exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickup_facility_id": Pickupfacilityid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "FacilityId " + Pickupfacilityid + " does not exist",
		})
		return nil, errors.New("Facilityid does not exist")
	}

	//get api starts here

	query := psql.Select("pickuprequest_id").
		From("pickup_main").
		Where(squirrel.Eq{"pickup_facility_id": Pickupfacilityid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	rows, err := dr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var pickupRequestIDs []int

	for rows.Next() {
		var pickupRequestID int
		if err := rows.Scan(&pickupRequestID); err != nil {
			return nil, err
		}
		pickupRequestIDs = append(pickupRequestIDs, pickupRequestID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pickupRequestIDs, nil

}

// calling repo function to get single pickuprequestid details
func (dr *DashboardRepository) Getcompletedetailsbasedobidfac(gctx *gin.Context, pickuprequestid int) (*domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pickupmain domain.Pickupmain

	query := psql.
		Select("pickuprequest_id", "customer_id", "pickup_schedule_slot", "pickup_schedule_date", "pickupagent_id", "pickup_address", "pickup_status", "customer_name", "customer_mobilenumber", "created_datetime", "modified_datetime").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	//var customerName interface{}
	err = dr.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		&pickupmain.Customerid,
		&pickupmain.Pickupscheduleslot,
		&pickupmain.Pickupscheduledate,
		&pickupmain.Pickupagentid,
		&pickupmain.Pickupaddress,
		&pickupmain.Pickupstatus,
		&pickupmain.Customername,
		&pickupmain.Customermobilenumber,
		&pickupmain.Createddatetime,
		&pickupmain.Modifieddatetime,
	)
	if err != nil {
		return nil, err
	}

	return &pickupmain, nil

}

func (dr *DashboardRepository) Getcompletedetailsbasedobidcus(gctx *gin.Context, pickuprequestid int) (*domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pickupmain domain.Pickupmain

	query := psql.
		Select("pickuprequest_id", "customer_id", "pickup_schedule_slot", "pickup_schedule_date", "pickupagent_id", "pickup_address", "pickup_status", "customer_name", "customer_mobilenumber", "created_datetime", "modified_datetime").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickuprequestid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	//var customerName interface{}
	err = dr.db.QueryRow(ctx, sql, args...).Scan(
		&pickupmain.Pickuprequestid,
		&pickupmain.Customerid,
		&pickupmain.Pickupscheduleslot,
		&pickupmain.Pickupscheduledate,
		&pickupmain.Pickupagentid,
		&pickupmain.Pickupaddress,
		&pickupmain.Pickupstatus,
		&pickupmain.Customername,
		&pickupmain.Customermobilenumber,
		&pickupmain.Createddatetime,
		&pickupmain.Modifieddatetime,
	)
	if err != nil {
		return nil, err
	}

	return &pickupmain, nil

}

// Repo Function to get all pickuprequest based on facilityid
func (dr *DashboardRepository) GetPickuprequestlistBypincode(gctx *gin.Context, Pickuprequestedpincode string) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// check wheather facility id exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickuprequestedpincode": Pickuprequestedpincode})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Pickuprequestedpinode " + Pickuprequestedpincode + " does not exist",
		})
		return nil, errors.New("Pickuprequestedpinode does not exist")
	}

	//get api starts here

	query := psql.Select("pickuprequest_id").
		From("pickup_main").
		Where(squirrel.Eq{"pickuprequestedpincode": Pickuprequestedpincode})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	//starts here for fetching multiple records
	rows, err := dr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var pickupRequestIDs []int

	for rows.Next() {
		var pickupRequestID int
		if err := rows.Scan(&pickupRequestID); err != nil {
			return nil, err
		}
		pickupRequestIDs = append(pickupRequestIDs, pickupRequestID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pickupRequestIDs, nil

}

// Repo Function to get all pickuprequestid assigned to pickupagentid
func (dr *DashboardRepository) GetPickuprequestbyPickupagentid(gctx *gin.Context, Pickupagentid int) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check wheather pickupagentid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickupagent_id": Pickupagentid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Pickupagentid " + strconv.Itoa(Pickupagentid) + " does not exist",
		})
		return nil, errors.New("Pickupagentid does not exist")
	}

	query := psql.Select("pickuprequest_id").From("pickup_main").Where(squirrel.Eq{"pickupagent_id": Pickupagentid})
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	rows, err := dr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var pickupRequestIDs []int

	for rows.Next() {
		var pickupRequestID int
		if err := rows.Scan(&pickupRequestID); err != nil {
			return nil, err
		}
		pickupRequestIDs = append(pickupRequestIDs, pickupRequestID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pickupRequestIDs, nil
}

// Repo Function to get all pickuprequestid raised by the customer
func (dr *DashboardRepository) GetPickuprequestbyCustomerid(gctx *gin.Context, Customerid string) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// check wheather pickupagentid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"customer_id": Customerid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := dr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Customerid " + Customerid + " does not exist",
		})
		return nil, errors.New("Customerid does not exist")
	}

	query := psql.Select("pickuprequest_id").From("pickup_main").Where(squirrel.Eq{"customer_id": Customerid})
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	dr.log.Debug("sql:", sql)
	dr.log.Debug("args:", args)

	rows, err := dr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var pickupRequestIDs []int

	for rows.Next() {
		var pickupRequestID int
		if err := rows.Scan(&pickupRequestID); err != nil {
			return nil, err
		}
		pickupRequestIDs = append(pickupRequestIDs, pickupRequestID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pickupRequestIDs, nil
}
