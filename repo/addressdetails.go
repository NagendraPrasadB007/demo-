package repository

import (
	"context"
	"errors"
	"net/http"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	"reflect"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

var GeneratedAddressID int

type AddressdetailsRepository struct {
	db  *DB
	log *logger.Logger
}

func NewAddressdetailsRepository(db *DB, log *logger.Logger) *AddressdetailsRepository {
	return &AddressdetailsRepository{
		db,
		log,
	}
}

// Handler Function used to fecth address of customer based on customerid
// *domain.Addressdetails //for fetching single record
func (ar *AddressdetailsRepository) GetAddressdetailsByID(gctx *gin.Context, customerid string) ([]domain.Addressdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the customerid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("address_details").
		Where(sq.Eq{"customer_id": customerid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil, existsErr
	}

	var count int
	if err := ar.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return nil, err
	}

	if count == 0 {
		// Assuming that ctx is of type *gin.Context
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			// Handle the case where ctx is not of the expected type
			return nil, errors.New("unexpected context type")
		}

		// Use ginContext.JSON instead of ctx.JSON
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "CustomerID " + customerid + " does not exist",
		})

		return nil, errors.New("customerid does not exist")
	}

	// get api starts here

	var addressdetails domain.Addressdetails
	//extra for fetching multiple records
	var addressdetailss []domain.Addressdetails

	//ar.log.Debug("customerid:", customerid)

	query := psql.Select("customer_id", "address_id", "firstname", "lastname", "addressline1", "addressline2", "landmark", "city", "state", "country", "pincode", "mobilenumber", "email_id", "geo_code", "address_type", "fromtopickup", "is_verified").
		From("address_details").
		Where(sq.Eq{"customer_id": customerid})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	//starts here for fetching multiple records
	rows, err := ar.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&addressdetails.Customerid,
			&addressdetails.Addressid,
			&addressdetails.Firstname,
			&addressdetails.Lastname,
			&addressdetails.Addressline1,
			&addressdetails.Addressline2,
			&addressdetails.Landmark,
			&addressdetails.City,
			&addressdetails.State,
			&addressdetails.Country,
			&addressdetails.Pincode,
			&addressdetails.Mobilenumber,
			&addressdetails.Emailid,
			&addressdetails.Geocode,
			&addressdetails.Addresstype,
			&addressdetails.Fromtopickup,
			&addressdetails.Isverified,
		)
		if err != nil {
			return nil, err
		}

		addressdetailss = append(addressdetailss, addressdetails)
	}

	return addressdetailss, nil
	//ends here

	/* for single response
	err = ar.db.QueryRow(ctx, sql, args...).Scan(
		&addressdetails.Customerid,
		&addressdetails.Addressid,
		&addressdetails.Firstname,
		&addressdetails.Lastname,
		&addressdetails.Addressline1,
		&addressdetails.Addressline2,
		&addressdetails.Landmark,
		&addressdetails.City,
		&addressdetails.State,
		&addressdetails.Country,
		&addressdetails.Pincode,
		&addressdetails.Mobilenumber,
		&addressdetails.Emailid,
		&addressdetails.Geocode,
		&addressdetails.Addresstype,
		&addressdetails.Fromtopickup,
		&addressdetails.Isverified,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, port.ErrDataNotFound
		}
		return nil, err
	}

	return &addressdetails, nil
	*/

}

func (ar *AddressdetailsRepository) GetAddressdetailsByaddressid(gctx *gin.Context, addressid int) (domain.Addressdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var addressdetails domain.Addressdetails

	query := psql.Select("customer_id", "address_id", "firstname", "lastname", "addressline1", "addressline2", "landmark", "city", "state", "country", "pincode", "mobilenumber", "email_id", "geo_code", "address_type", "fromtopickup", "is_verified").
		From("address_details").
		Where(sq.Eq{"address_id": addressid})

	sql, args, err := query.ToSql()

	err = ar.db.QueryRow(ctx, sql, args...).Scan(
		&addressdetails.Customerid,
		&addressdetails.Addressid,
		&addressdetails.Firstname,
		&addressdetails.Lastname,
		&addressdetails.Addressline1,
		&addressdetails.Addressline2,
		&addressdetails.Landmark,
		&addressdetails.City,
		&addressdetails.State,
		&addressdetails.Country,
		&addressdetails.Pincode,
		&addressdetails.Mobilenumber,
		&addressdetails.Emailid,
		&addressdetails.Geocode,
		&addressdetails.Addresstype,
		&addressdetails.Fromtopickup,
		&addressdetails.Isverified,
	)
	if err != nil {
		return addressdetails, err
	}

	return addressdetails, nil

}

/*

type Addressvalues struct {
	Customerid   string
	Firstname    string
	Lastname     string
	Addressline1 string
	Addressline2 string
	Landmark     string
	City         string
	State        string
	Country      string
	Pincode      string
	Mobilenumber string
	Emailid      string
	Geocode      string
	Addresstype  string
	Fromtopickup string
	Isverified   bool
}

func generateMapFromStruct(instance interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.Indirect(reflect.ValueOf(instance))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("insert")
		if tag != "" {
			result[tag] = val.Field(i).Interface()
		}
	}

	return result
}

func (ar *AddressdetailsRepository) executequery(builder sq.Sqlizer, s interface{}, ctx context.Context) (interface{}, error) {
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	row, err := ar.db.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	u, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Addressdetails])
	if err != nil {
		return nil, err
	}
	return &u, nil

}
*/

// Repo Function used to create address of customer
func (ar *AddressdetailsRepository) CreateAddressForCustomer(gctx *gin.Context, address *domain.Addressdetails) (*domain.Addressdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//if agin geo code is same it will show error
	// existsQuery := psql.Select("COUNT(*)").
	// 	From("address_details").
	// 	Where(sq.Eq{"geo_code": address.Geocode})

	// existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	// if existsErr != nil {
	// 	return nil, existsErr
	// }

	// var count int
	// if err := ar.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
	// 	return nil, err
	// }

	// if count == 0 {
	// 	// Assuming that ctx is of type *gin.Context
	// 	ginContext, ok := ctx.(*gin.Context)
	// 	if !ok {
	// 		// Handle the case where ctx is not of the expected type
	// 		return nil, errors.New("unexpected context type")
	// 	}

	// 	// Use ginContext.JSON instead of ctx.JSON
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{
	// 		"message": fmt.Sprintf("Geocode already exists"),
	// 	})

	// 	return nil, errors.New("Geocode already exists")
	// }

	// if count > 0 {
	// 	return nil, errors.New("geo code already exists")
	// }

	//post query starts here

	query := psql.Insert("address_details").
		//Columns(result).
		Columns("customer_id", "firstname", "lastname", "addressline1", "addressline2", "landmark", "city", "state", "country", "pincode", "mobilenumber", "email_id", "geo_code", "address_type", "fromtopickup", "is_verified").
		Values(address.Customerid, address.Firstname, address.Lastname, address.Addressline1, address.Addressline2, address.Landmark, address.City, address.State, address.Country, address.Pincode, address.Mobilenumber, address.Emailid, address.Geocode, address.Addresstype, address.Fromtopickup, address.Isverified).Suffix("RETURNING address_id")

	//query := psql.Insert("address_details").SetMap(generateMapFromStruct(address)).Suffix("RETURNING address_id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	err = ar.db.QueryRow(ctx, sql, args...).Scan(&GeneratedAddressID)
	if err != nil {
		return nil, err
	}

	return address, nil

}

// Repo Function used to update address of customer based on addressid
func (ar *AddressdetailsRepository) UpdateAddress(gctx *gin.Context, address *domain.Addressdetails) (*domain.Addressdetails, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Check if the addressid exists or not
	// existsQuery := psql.Select("COUNT(*)").
	// 	From("address_details").
	// 	Where(sq.Eq{"address_id": address.Addressid})

	// existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	// if existsErr != nil {
	// 	return nil, existsErr
	// }

	// var count int
	// if err := ar.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
	// 	return nil, err
	// }

	// if count == 0 {
	// 	// Assuming that ctx is of type *gin.Context
	// 	ginContext, ok := ctx.(*gin.Context)
	// 	if !ok {
	// 		// Handle the case where ctx is not of the expected type
	// 		return nil, errors.New("unexpected context type")
	// 	}

	// 	// Use ginContext.JSON instead of ctx.JSON
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "AddressID " + strconv.Itoa(address.Addressid) + " does not exist",
	// 	})

	// 	return nil, errors.New("addressid does not exist")
	// }

	query := psql.Update("address_details").SetMap(generateMapFromStructForUpdate(address)).Where(sq.Eq{"address_id": address.Addressid}).Suffix("returning *")

	/*
		// First Approach
			query := psql.Update("address_details").
				Set("customer_id", sq.Expr("COALESCE(?, customer_id)", address.Customerid)).
				Set("firstname", sq.Expr("COALESCE(?, firstname)", address.Firstname)).
				Set("lastname", sq.Expr("COALESCE(?, lastname)", address.Lastname)).
				Set("addressline1", sq.Expr("COALESCE(?, addressline1)", address.Addressline1)).
				Set("addressline2", sq.Expr("COALESCE(?, addressline2)", address.Addressline2)).
				Set("landmark", sq.Expr("COALESCE(?, landmark)", address.Landmark)).
				Set("city", sq.Expr("COALESCE(?, city)", address.City)).
				Set("state", sq.Expr("COALESCE(?, state)", address.State)).
				Set("country", sq.Expr("COALESCE(?, country)", address.Country)).
				Set("pincode", sq.Expr("COALESCE(?, pincode)", address.Pincode)).
				Set("mobilenumber", sq.Expr("COALESCE(?, mobilenumber)", address.Mobilenumber)).
				Set("email_id", sq.Expr("COALESCE(?, email_id)", address.Emailid)).
				Set("geo_code", sq.Expr("COALESCE(?, geo_code)", address.Geocode)).
				Set("address_type", sq.Expr("COALESCE(?, address_type)", address.Addresstype)).
				Set("fromtopickup", sq.Expr("COALESCE(?, fromtopickup)", address.Fromtopickup)).
				Set("is_verified", sq.Expr("COALESCE(?, is_verified)", address.Isverified)).
				Where(sq.Eq{"address_id": address.Addressid}).
			Suffix("RETURNING *")
	*/

	//update query begins here

	/*
		query := psql.Update("address_details").
			Where(sq.Eq{"address_id": address.Addressid}).
			Suffix(`RETURNING "customer_id", "address_id", "firstname", "lastname", "addressline1", "addressline2", "landmark", "city", "state", "country", "pincode", "mobilenumber", "email_id", "geo_code", "address_type", "fromtopickup", "is_verified"`)

		if address.Customerid != "" {
			query = query.Set("customer_id", address.Customerid)
		}
		if address.Firstname != "" {
			query = query.Set("firstname", address.Firstname)
		}
		if address.Lastname != "" {
			query = query.Set("lastname", address.Lastname)
		}
		if address.Addressline1 != "" {
			query = query.Set("addressline1", address.Addressline1)
		}
		if address.Addressline2 != "" {
			query = query.Set("addressline2", address.Addressline2)
		}
		if address.Landmark != "" {
			query = query.Set("landmark", address.Landmark)
		}
		if address.City != "" {
			query = query.Set("city", address.City)
		}
		if address.State != "" {
			query = query.Set("state", address.State)
		}
		if address.Country != "" {
			query = query.Set("country", address.Country)
		}
		if address.Pincode != "" {
			query = query.Set("pincode", address.Pincode)
		}
		if address.Mobilenumber != "" {
			query = query.Set("mobilenumber", address.Mobilenumber)
		}
		if address.Emailid != "" {
			query = query.Set("email_id", address.Emailid)
		}
		if address.Geocode != "" {
			query = query.Set("geo_code", address.Geocode)
		}
		if address.Addresstype != "" {
			query = query.Set("address_type", address.Addresstype)
		}
		if address.Fromtopickup != "" {
			query = query.Set("fromtopickup", address.Fromtopickup)
		}
		query = query.Set("is_verified", address.Isverified)
	*/

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	err = ar.db.QueryRow(ctx, sql, args...).Scan(
		&address.Customerid,
		&address.Addressid,
		&address.Firstname,
		&address.Lastname,
		&address.Addressline1,
		&address.Addressline2,
		&address.Landmark,
		&address.City,
		&address.State,
		&address.Country,
		&address.Pincode,
		&address.Mobilenumber,
		&address.Emailid,
		&address.Geocode,
		&address.Addresstype,
		&address.Fromtopickup,
		&address.Isverified,
	)
	if err != nil {
		return nil, err
	}

	return address, nil

}

func generateMapFromStructForUpdate(instance interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.Indirect(reflect.ValueOf(instance))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tagValue := field.Tag.Get("update")
		if tagValue != "" {
			result[tagValue] = val.Field(i).Interface()
		}
	}

	return result
}

// Repo Function used to delete address of customer based on addressid
func (ar *AddressdetailsRepository) DeleteAddressByAddressID(gctx *gin.Context, addressid int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the addressid exists or not
	existsQuery := psql.Select("COUNT(*)").
		From("address_details").
		Where(sq.Eq{"address_id": addressid})

	existsSQL, existsArgs, existsErr := existsQuery.ToSql()
	if existsErr != nil {
		return nil
	}

	var count int
	if err := ar.db.QueryRow(ctx, existsSQL, existsArgs...).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		// Assuming that ctx is of type *gin.Context
		ginContext, ok := ctx.(*gin.Context)
		if !ok {
			// Handle the case where ctx is not of the expected type
			return errors.New("unexpected context type")
		}

		// Use ginContext.JSON instead of ctx.JSON
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "AddressID " + strconv.Itoa(addressid) + " does not exist",
		})

		return errors.New("addressid does not exist")
	}

	// delete api starts here
	deleteQuery := psql.Delete("address_details").
		Where(sq.Eq{"address_id": addressid})

	sql, args, err := deleteQuery.ToSql()
	if err != nil {
		return err
	}

	ar.log.Debug("sql:", sql)
	ar.log.Debug("args:", args)

	_, err = ar.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil

}

//RR sir update

/*
type Addressvalues struct {
	Customerid   string
	Firstname    string
	Lastname     string
	Addressline1 string
	Addressline2 string
	Landmark     string
	City         string
	State        string
	Country      string
	Pincode      string
	Mobilenumber string
	Emailid      string
	Geocode      string
	Addresstype  string
	Fromtopickup string
	Isverified   bool
}
*/

/*
func generateMapFromStruct(instance interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	//var result []map[string]interface{}

	fmt.Println("instance:", instance)

	val := reflect.Indirect(reflect.ValueOf(instance))

	fmt.Println("val:", val)
	typ := val.Type()
	fmt.Println("Struct:", typ)
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("update")

		fmt.Println("value:", val.Field(i).Interface())
		//if tag != "" {
		fmt.Println("Came inside")
		result[tag] = val.Field(i).Interface()
		//result = append(result, map[string]interface{}{tag: val.Field(i).Interface()})
		fmt.Println("result inside lopp", result)
		//}
	}
	//fmt.Println(result[tag])
	fmt.Println("result", result)
	return result
}
*/

// func generateMapFromStruct(instance interface{}) map[string]interface{} {
//     result := make(map[string]interface{})

//     val := reflect.Indirect(reflect.ValueOf(instance))
//     typ := val.Type()

//     // Define the order of columns
//     columnsOrder := []string{"customer_id", "firstname", "lastname",, "is_verified"}

//     for _, col := range columnsOrder {
//         for i := 0; i < val.NumField(); i++ {
//             field := typ.Field(i)
//             tag := field.Tag.Get("update")

//             if tag == col {
//                 result[tag] = val.Field(i).Interface()
//             }
//         }
//     }

//     return result
// }

// func (ar *AddressdetailsRepository) UpdateAddress(ctx context.Context, address *domain.Addressdetails) (*domain.Addressdetails, error) {

// 	ar.log.Debug("enetered repo layer")

// 	/*
// 		updateMap := generateMapFromStruct(address)

// 		if len(updateMap) == 0 {
// 			updateMap["dummy"] = "dummy" // Add a dummy update
// 		}

// 		updateBuilder := squirrel.Update("address_details").
// 			SetMap(updateMap).
// 			Where(squirrel.Eq{"address_id": address.Addressid}).
// 			Suffix("RETURNING *")
// 	*/

// 	//maps := generateMapFromStruct(address)

// 	updateBuilder := psql.Update("address_details").
// 		SetMap(generateMapFromStruct(address)).
// 		Where(squirrel.Eq{"address_id": address.Addressid}).
// 		Suffix("RETURNING *")

// 	ar.log.Debug("finished upadate query")

// 	sqlString, args, err := updateBuilder.ToSql()
// 	if err != nil {
// 		fmt.Println("Error generating SQL:", err)
// 		return nil, err
// 	}

// 	ar.log.Debug("sql:", sqlString)
// 	ar.log.Debug("args:", args)

// 	row, err := ar.db.Query(ctx, sqlString, args...)

// 	if err != nil {
// 		return nil, err
// 	}

// 	u, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Addressdetails])

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &u, nil

// }
