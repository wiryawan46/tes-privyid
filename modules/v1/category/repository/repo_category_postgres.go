package repository

import (
	"database/sql"
	"log"
	"pretest-privyid/modules/v1/category/model"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type CategoryRepoPostgres struct {
	dbConn *sql.DB
}

func NewWorkCategoryRepoPostgres(dbConn *sql.DB) *CategoryRepoPostgres {
	return &CategoryRepoPostgres{dbConn: dbConn}
}

func (conn *CategoryRepoPostgres) CreateCategory(param model.Category) ResultRepository {
	output := ResultRepository{}

	query := "INSERT INTO category(name) VALUES($1)"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(param.Name)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: param}
	return output
}