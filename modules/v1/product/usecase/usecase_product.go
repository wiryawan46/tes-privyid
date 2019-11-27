package usecase

import (
	"fmt"
	"log"
	"pretest-privyid/modules/v1/product/model"
	"pretest-privyid/modules/v1/product/repository"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ProductUsecaseImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(ProductRepo repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{
		ProductRepository: ProductRepo,
	}
}

func (rp *ProductUsecaseImpl) CreateProduct(param model.Product) ResultUseCase {
	output := ResultUseCase{}

	saveResult := rp.ProductRepository.CreateProduct(param)
	if saveResult.Error != nil {
		log.Println("Error menyimpan data :", saveResult.Error.Error())
		output = ResultUseCase{Error: saveResult.Error}
		return output
	}
	data, ok := saveResult.Result.(model.Product)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	output = ResultUseCase{Result: data}
	return output
}
