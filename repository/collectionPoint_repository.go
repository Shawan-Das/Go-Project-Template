package repository

import (
	"net/http"

	"github.com/tools/iservice/dto"
	"github.com/tools/iservice/util"
)

type CollectionPointRepo struct{}

type CollectionPoint struct {
	Compcode int    `json:"compcode"`
	Compname string `json:"compname"`
}

func (cpr *CollectionPointRepo) GetCollectionPointName() dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	var data []CollectionPoint

	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	query := `SELECT compcode, compname from common.compinfo WHERE compcode IN (1,2,6) ORDER BY compname`

	getData := db.Raw(query).Find(&data)

	if getData.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusBadRequest
		r.Message = "Somehting went wrong"
		r.Data = getData.Error.Error()
	} else if getData.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "No collection point found"
		// r.Data = getData.Error.Error()
	} else {
		r.IsSuccess = true
		r.StatusCode = http.StatusOK
		r.Message = "Collection point found "
		r.Data = data
	}

	return r
}

