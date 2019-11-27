package repository

import (
	"database/sql"
	"log"
	"pretest-privyid/modules/v1/product/model"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ProductRepoPostgres struct {
	dbConn *sql.DB
}

func NewProductRepoPostgres(dbConn *sql.DB) *ProductRepoPostgres {
	return &ProductRepoPostgres{dbConn: dbConn}
}

func (conn *ProductRepoPostgres) CreateProduct(param model.Product) ResultRepository {
	output := ResultRepository{}

	query := "INSERT INTO product(name, description) VALUES($1, $2) RETURNING id"
	id := 0
	errorDB := conn.dbConn.QueryRow(query, param.Name, param.Description).Scan(&id)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	for i := 0; i < len(param.Categories); i++ {
		query := "INSERT INTO category_product(product_id, category_id) VALUES($1,$2)"
		sqlStmt, errorDB := conn.dbConn.Prepare(query)
		if errorDB != nil {
			log.Println("Error prepare query : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		result, errorExecute := sqlStmt.Exec(id, param.Categories[i].ID)
		if errorExecute != nil {
			log.Println("Error executing query : ", errorExecute.Error())
			output = ResultRepository{Error: errorExecute}
			return output
		}
		result.RowsAffected()
	}

	output = ResultRepository{Result: param}
	return output
}

func (conn *ProductRepoPostgres) UploadImage(productId string, param model.Image) ResultRepository {
	output := ResultRepository{}

	query := "INSERT INTO image(name, file) VALUES($1, $2) RETURNING id"
	id := 0
	errorDB := conn.dbConn.QueryRow(query, param.Name, param.File).Scan(&id)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	sqlQuery := "INSERT INTO product_image(product_id, image_id) VALUES($1,$2)"
	sqlStmt, errorDB := conn.dbConn.Prepare(sqlQuery)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(productId, id)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: result}
	return output
}