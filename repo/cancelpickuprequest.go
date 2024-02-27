package repository

import (
	"context"
	"errors"
	"net/http"
	"pickupmanagement/logger"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type CancelpickuprequestRepository struct {
	db  *DB
	log *logger.Logger
}

func NewCancelpickuprequestRepository(db *DB, log *logger.Logger) *CancelpickuprequestRepository {
	return &CancelpickuprequestRepository{
		db,
		log,
	}
}

// Repo function used to cancel pickuprequest based on pickuprequestid
func (cr *CancelpickuprequestRepository) Cancelpickuprequestdetails(gctx *gin.Context, pickuprequestid int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the pickuprequestid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("pickup_main").
		Where(sq.Eq{"pickuprequest_id": pickuprequestid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return existsErr
	}

	var count int
	if err := cr.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return errors.New("unexpected context type")
		}

		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "PickuprequestID " + strconv.Itoa(pickuprequestid) + " does not exist",
		})

		//return nil, nil, nil, nil, nil, nil, errors.New("addressid does not exist")
		return errors.New("pickuprequestid does not exist")
	}

	
	query := psql.Update("pickup_main").
		Set("pickup_status", "Cancelled").
		Where(sq.Eq{"pickuprequest_id": pickuprequestid})

	cancelsql, cancelargs, err := query.ToSql()
	if err != nil {
		return err
	}

	cr.log.Debug("cancelsql:", cancelsql)
	cr.log.Debug("cancelargs:", cancelargs)

	_, err = cr.db.Exec(ctx, cancelsql, cancelargs...)
	if err != nil {
		return err
	}

	return nil

}
