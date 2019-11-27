package presenter

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"pretest-privyid/helper"
	"pretest-privyid/modules/v1/product/model"
	"pretest-privyid/modules/v1/product/usecase"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
type HTTPProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewHTTPHandler(ProductUsecase usecase.ProductUsecase) *HTTPProductHandler {
	return &HTTPProductHandler{
		ProductUsecase: ProductUsecase,
	}
}

func (h *HTTPProductHandler) MountProduct(group *echo.Group)  {
	group.POST("/product", h.CreateProduct)
}

func (h *HTTPProductHandler) CreateProduct(c echo.Context) error  {
	reqData := model.Product{}
	err := c.Bind(&reqData)
	if err != nil {
		log.Println("bind_request_data", err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter body kosong", nil))
	}

	saveResult := h.ProductUsecase.CreateProduct(reqData)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal menambah data category")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data, ok := saveResult.Result.(model.Product)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		return c.JSON(http.StatusOK, helper.ResponseDetailOutput(err.Error(), data))
	}
	return c.JSON(http.StatusCreated, data)
}