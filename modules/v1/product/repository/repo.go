package repository

import "pretest-privyid/modules/v1/product/model"

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ResultRepository struct {
	Result interface{}
	Error  error
}

type ProductRepository interface {
	CreateProduct(param model.Product) ResultRepository
	UploadImage(productId string, param model.Image) ResultRepository
}