package repository

import (
	"context"
	"errors"
	"io/ioutil"
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

var GeneratedImageID int

type ImageRepository struct {
	db  *DB
	log *logger.Logger
}

func NewImageRepository(db *DB, log *logger.Logger) *ImageRepository {
	return &ImageRepository{
		db,
		log,
	}
}

// Repo function to upload the image of open article
func (ir *ImageRepository) SaveImage(gctx *gin.Context, image *domain.Image, filePath string) (*domain.Image, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the filePath already exists
	existsQuery := psql.Select("COUNT(*)").
		From("image").
		Where(sq.Eq{"filePath": filePath})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := ir.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count > 0 {
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.New("unexpected context type")
		}
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Image with the given filePath already exists",
		})
		return nil, errors.New("image with the given filePath already exists")
	}

	//post api starts here

	// If the filePath doesn't exist, proceed with the insertion
	query := psql.Insert("image").
		Columns("filename", "size", "mimetype", "filePath", "pickuprequest_id").
		Values(image.Filename, image.Size, image.Mimetype, filePath, image.Pickuprequestid).
		Suffix("RETURNING imageid")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	ir.log.Debug("sql:", sql)
	ir.log.Debug("args:", args)

	err = ir.db.QueryRow(ctx, sql, args...).Scan(
		&GeneratedImageID,
	)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// Repo function to get the image of the pickuprequest based on imageid
func (ir *ImageRepository) GetImagedetailsByID(gctx *gin.Context, imageid int) ([]byte, *domain.Image, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// first checking wheather imageid is exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("image").
		Where(sq.Eq{"imageid": imageid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, nil, existsErr
	}

	var count int
	if err := ir.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, nil, err
	}

	if count == 0 {
		// Assuming that ctx is of type *gin.Context
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			// Handle the case where ctx is not of the expected type
			return nil, nil, errors.New("unexpected context type")
		}

		// Use ginContext.JSON instead of ctx.JSON
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Imageid " + strconv.Itoa(imageid) + " does not exist",
		})

		return nil, nil, errors.New("Imageid does not exist")
	}

	// get api starts here

	var imageinfo domain.Image

	query := psql.Select("imageid", "filename", "size", "mimetype", "created_datetime", "filepath", "pickuprequest_id").
		From("image").
		Where(sq.Eq{"imageid": imageid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, err
	}

	ir.log.Debug("sql:", sql)
	ir.log.Debug("args:", args)

	err = ir.db.QueryRow(ctx, sql, args...).Scan(
		&imageinfo.Imageid,
		&imageinfo.Filename,
		&imageinfo.Size,
		&imageinfo.Mimetype,
		&imageinfo.Createddatetime,
		&imageinfo.Filepath,
		&imageinfo.Pickuprequestid,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil, port.ErrDataNotFound
		}
		return nil, nil, err
	}

	imageBytes, err := ioutil.ReadFile(imageinfo.Filepath)
	if err != nil {
		return nil, nil, err
	}

	return imageBytes, &imageinfo, nil
}
