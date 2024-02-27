package repository

import (
	"context"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type PickupscheduleslotsRepository struct {
	db  *DB
	log *logger.Logger
}

func NewPickupscheduleslotsRepository(db *DB, log *logger.Logger) *PickupscheduleslotsRepository {
	return &PickupscheduleslotsRepository{
		db,
		log,
	}
}

// Repo function to get all pickupschedulslots
func (psr *PickupscheduleslotsRepository) Listslots(gctx *gin.Context) ([]domain.Pickupscheduleslots, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var slot domain.Pickupscheduleslots
	var slots []domain.Pickupscheduleslots

	query := psql.Select("pickupscheduleslot_id", "schedule_slots").
		From("pickup_scheduleslots").
		OrderBy("pickupscheduleslot_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := psr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&slot.Pickupscheduleslotid,
			&slot.Scheduleslots,
		)
		if err != nil {
			return nil, err
		}

		slots = append(slots, slot)
	}

	return slots, nil
}

// repo function to get all pickup agents
func (psr *PickupscheduleslotsRepository) Pickupagentlist(gctx *gin.Context) ([]domain.Pickupagent, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var list domain.Pickupagent
	var lists []domain.Pickupagent

	query := psql.Select("pickupagent_id", "pickupagent_name").
		From("pickupagent").
		OrderBy("pickupagent_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := psr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&list.Pickupagentid,
			&list.Pickupagentname,
		)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	return lists, nil
}

// Repo function to get remarks
func (psr *PickupscheduleslotsRepository) Remarkslist(gctx *gin.Context) ([]domain.Remarks, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var list domain.Remarks
	var lists []domain.Remarks

	query := psql.Select("remarks_id", "remarks").
		From("remarks").
		OrderBy("remarks_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := psr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&list.Remarksid,
			&list.Remarks,
		)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	return lists, nil
}

// Repo function to get addresstype
func (psr *PickupscheduleslotsRepository) Addresstypelist(gctx *gin.Context) ([]domain.Addresstype, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var list domain.Addresstype
	var lists []domain.Addresstype

	//implement list and offset
	query := psql.Select("addresstype_id", "addresstype").
		From("addresstype").
		OrderBy("addresstype")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := psr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&list.Addresstypeid,
			&list.Addresstype,
		)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	return lists, nil
}
