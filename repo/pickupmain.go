package repository

import (
	"context"
	"errors"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	"time"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

var GeneratedPickupRequestID int

type PickupmainRepository struct {
	db  *DB
	log *logger.Logger
}

func NewPickupmainRepository(db *DB, log *logger.Logger) *PickupmainRepository {
	return &PickupmainRepository{
		db,
		log,
	}
}

// Repo function to create a new request in pickupmain table
func (pr *PickupmainRepository) Pickupmain(gctx *gin.Context, pickupmain *domain.Pickupmain) (*domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// for inserting into created_datetime and modified_datetime i added CURRENT_TIMESTAMP in DB
	pr.log.Info("entered pickupmain repo method")
	query := psql.Insert("pickup_main").
		Columns("customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime").
		Values(pickupmain.Customerid, pickupmain.Pickupdroptype, pickupmain.Pickuplocation, pickupmain.Droplocation, pickupmain.Pickupscheduleslot, pickupmain.Pickupscheduledate, pickupmain.Actualpickupdatetime, pickupmain.Pickupagentid, pickupmain.Pickupfacilityid, pickupmain.Pickupstatus, pickupmain.Paymentstatus, pickupmain.Pickupaddress, pickupmain.Domesticforeignidentifier, pickupmain.Pickuplong, pickupmain.Pickuplat, pickupmain.Pickuprequestedpincode, pickupmain.Customername, pickupmain.Customermobilenumber, pickupmain.Assigneddatetime).Suffix(`RETURNING "pickuprequest_id","customer_id","pickup_drop_type","pickup_location","drop_location","pickup_schedule_slot","pickup_schedule_date","actual_pickup_datetime","pickupagent_id","pickup_facility_id","pickup_status","payment_status","created_datetime","pickup_address","domestic_foreign_identifier","pickup_long","pickup_lat","modified_datetime","pickuprequestedpincode","customer_name","customer_mobilenumber","assigned_datetime"`)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	pr.log.Debug("sql:", sql)
	pr.log.Debug("args:", args)

	err = pr.db.QueryRow(ctx, sql, args...).Scan(
		&GeneratedPickupRequestID,
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

	if err != nil {
		return nil, err
	}

	return pickupmain, nil
}

// Repo function to create bulk insertion in pickupmain table
func (pr *PickupmainRepository) PickupmainBatch(gctx *gin.Context, pickupmains []domain.Pickupmain) ([]domain.Pickupmain, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pr.log.Info("entered pickupmainBatch repo method")

	// Create a slice to store the generated IDs
	var generatedIDs []int

	query := psql.Insert("pickup_main").
		Columns("customer_id", "pickup_drop_type", "pickup_location", "drop_location", "pickup_schedule_slot", "pickup_schedule_date", "actual_pickup_datetime", "pickupagent_id", "pickup_facility_id", "pickup_status", "payment_status", "pickup_address", "domestic_foreign_identifier", "pickup_long", "pickup_lat", "pickuprequestedpincode", "customer_name", "customer_mobilenumber", "assigned_datetime").Suffix("RETURNING pickuprequest_id")

	for _, pickupmain := range pickupmains {
		query = query.Values(
			pickupmain.Customerid,
			pickupmain.Pickupdroptype,
			pickupmain.Pickuplocation,
			pickupmain.Droplocation,
			pickupmain.Pickupscheduleslot,
			pickupmain.Pickupscheduledate,
			pickupmain.Actualpickupdatetime,
			pickupmain.Pickupagentid,
			pickupmain.Pickupfacilityid,
			pickupmain.Pickupstatus,
			pickupmain.Paymentstatus,
			pickupmain.Pickupaddress,
			pickupmain.Domesticforeignidentifier,
			pickupmain.Pickuplong,
			pickupmain.Pickuplat,
			pickupmain.Pickuprequestedpincode,
			&pickupmain.Customername,
			&pickupmain.Customermobilenumber,
			&pickupmain.Assigneddatetime,
		)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	pr.log.Debug("sql:", sql)
	pr.log.Debug("args:", args)

	rows, err := pr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var generatedID int
		err := rows.Scan(&generatedID)
		if err != nil {
			return nil, err
		}
		generatedIDs = append(generatedIDs, generatedID)
	}

	for i, generatedID := range generatedIDs {
		pickupmains[i].Pickuprequestid = generatedID
	}

	pr.log.Debug("generatedIDs:", generatedIDs)

	return pickupmains, nil

}

/*

// AssigningPickuprequests
func (pr *PickupmainRepository) AssigningPickuprequests(ctx context.Context, pickupmain *domain.Pickupmain) error {

	query := psql.Update("pickup_main").
		Set("pickupagent_id", pickupmain.Pickupagentid).
		Where(sq.Eq{"pickuprequest_id": pickupmain.Pickuprequestid})

	// if pickupmain.Pickupagentid != "" {
	// 	query = query.Set("pickupagent_id", pickupmain.Pickupagentid)
	// }

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	pr.log.Debug("sql:", sql)
	pr.log.Debug("args:", args)

	_, err = pr.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil

}

*/

// Repo function to assign pickuprequest to pickupagent
func (pr *PickupmainRepository) AssignPickupAgents(gctx *gin.Context, requestIDs []int, agentIDs []int) error {
	// Ensure requestIDs and agentIDs have the same length

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if len(requestIDs) != len(agentIDs) {
		return errors.New("number of request IDs and agent IDs must match")
	}

	//work
	for i := range requestIDs {
		query := psql.Update("pickup_main").
			Set("pickup_status", "Assigned").
			Set("assigned_datetime", squirrel.Expr("CURRENT_TIMESTAMP")).
			Set("pickupagent_id", agentIDs[i]).
			Where(sq.Eq{"pickuprequest_id": requestIDs[i]})

		sql, args, err := query.ToSql()
		if err != nil {
			return err
		}

		pr.log.Debug("sql :", sql)
		pr.log.Debug("args :", args)

		_, err = pr.db.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
	}

	return nil
}
