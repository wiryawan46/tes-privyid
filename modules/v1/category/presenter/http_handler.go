package presenter

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"pretest-privyid/helper"
	"pretest-privyid/modules/v1/category/model"
	"pretest-privyid/modules/v1/category/usecase"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
type HTTPCategoryHandler struct {
	CategoryUsecase usecase.CategoryUsecase
}

func NewHTTPHandler(CategoryUsecase usecase.CategoryUsecase) *HTTPCategoryHandler {
	return &HTTPCategoryHandler{
		CategoryUsecase: CategoryUsecase,
	}
}

func (h *HTTPCategoryHandler) MountWorkField(group *echo.Group)  {
	group.POST("/category", h.CreateCategory)
}

func (h *HTTPCategoryHandler) CreateCategory(c echo.Context) error  {
	reqData := model.Category{}
	err := c.Bind(&reqData)
	if err != nil {
		log.Println("bind_request_data", err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter body kosong", nil))
	}

	saveResult := h.CategoryUsecase.CreateCategory(reqData)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal menambah data category")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data, ok := saveResult.Result.(model.Category)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		return c.JSON(http.StatusOK, helper.ResponseDetailOutput(err.Error(), data))
	}
	return c.JSON(http.StatusCreated, data)
}