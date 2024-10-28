package repository

import (
	"net/http"
	"strconv"

	"github.com/tools/iservice/dto"
	"github.com/tools/iservice/model"
	"github.com/tools/iservice/util"
)

type ServicePartyRepo struct{}

func (spr *ServicePartyRepo) GetServicePartyCode() dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	var id dto.PartyCode

	db := util.CreateConnectionUsingGormToSitlPosSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	getCode := db.Raw("select coalesce ((max(party_code) + 1), 60000001) as \"party_code\" from iservice.service_party").Find(&id)

	if getCode.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong"
		r.Data = getCode.Error.Error()
	} else if id.PartyCode == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Unable to get service party id"
	} else {
		r.IsSuccess = true
		r.StatusCode = http.StatusOK
		r.Message = "New ID found"
		r.Data = id
	}
	return r
}

func (spr *ServicePartyRepo) GetAllServiceParty() dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	var data []model.ServiceParty

	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	getAllParty := db.Order("party_code").Find(&data)

	if getAllParty.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong"
		r.Data = getAllParty.Error.Error()
	} else if getAllParty.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "No party found"
	} else {
		r.IsSuccess = true
		r.StatusCode = http.StatusOK
		r.Message = "Total Data " + strconv.Itoa(len(data))
		r.Data = data
	}
	return r
}

func (spr *ServicePartyRepo) GetAllServicePartyByName() dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	var data []model.ServiceParty

	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	getAllParty := db.Order("party_name").Find(&data)

	if getAllParty.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong"
		r.Data = getAllParty.Error.Error()
	} else if getAllParty.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "No party found"
	} else {
		r.IsSuccess = true
		r.StatusCode = http.StatusOK
		r.Message = "Total Data " + strconv.Itoa(len(data))
		r.Data = data
	}
	return r
}

func (spr *ServicePartyRepo) GetOneParty(id dto.PartyCode) dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	var output model.ServiceParty
	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	getEmp := db.Where("party_code", id.PartyCode).Find(&output)

	if getEmp.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong"
		r.Data = getEmp.Error.Error()
	} else if getEmp.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "No party found"
	} else {
		r.IsSuccess = true
		r.StatusCode = http.StatusOK
		r.Message = "Data found "
		r.Data = output
	}

	return r
}

func (spr *ServicePartyRepo) CreateParty(i model.ServiceParty) dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	var openbalances []model.Openbalance
	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	tx := db.Begin()
	var PartyCode int
	_ = db.Raw("select coalesce ((max(party_code) + 1), 60000001) as \"party_code\" from iservice.service_party").Find(&PartyCode)

	if i.PartyCode != PartyCode {
		i.PartyCode = PartyCode
	}
	i.Id = i.PartyCode - 60000000
	i.PartyName = util.TrimString(i.PartyName)
	i.PartyAddress = util.TrimString(i.PartyAddress)
	i.PartyPhone = util.TrimString(i.PartyPhone)
	i.PartyEmail = util.TrimString(i.PartyEmail)

	r.IsSuccess, r.Message = ValidateServiceParty(i)

	if !r.IsSuccess {
		r.StatusCode = http.StatusBadRequest
		return r
	}

	// check if phone number already exist or not
	var checkPhone model.ServiceParty
	checkPhoneNo := tx.Where("party_phone like '%" + i.PartyPhone + "%'").Find(&checkPhone)

	if checkPhoneNo.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusInternalServerError
		r.Message = "Something went wrong"
		r.Data = checkPhoneNo.Error.Error()
		return r
	} else if checkPhoneNo.RowsAffected != 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusConflict
		r.Message = "Phone Number Already Exists - " + checkPhone.PartyName
		return r
	}

	checkId := tx.Where("party_code", i.PartyCode).Find(&checkPhone)
	if checkId.Error != nil {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong"
		r.Data = checkId.Error.Error()
		return r
	} else if checkId.RowsAffected != 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Invalid Party Code. Please try again."
		return r
	}

	// get financial year and company info
	yearList, compList := util.Company_Year_into(db)

	if len(yearList) == 0 {
		r.StatusCode = http.StatusNotFound
		r.IsSuccess = false
		r.Message = "No year found"
		return r
	}
	if len(compList) == 0 {
		r.StatusCode = http.StatusNotFound
		r.IsSuccess = false
		r.Message = "No Company Found"
		return r
	}
	// count := 0
	// prepare for iservice.openbalance data
	for _, x := range compList {
		for _, y := range yearList {
			openbalance := model.Openbalance{
				Accid:       i.PartyCode,
				Openbalance: 0.00,
				Headname:    i.PartyName,
				Parent:      60000000,
				Compid:      x,
				Yearid:      y,
			}
			openbalances = append(openbalances, openbalance)
		}
	}

	// insert data into iservice.openbalance table
	result1 := tx.Create(&openbalances)
	if result1.Error != nil || int(result1.RowsAffected) != len(compList)*len(yearList) {
		r.StatusCode = http.StatusBadRequest
		r.IsSuccess = false
		r.Message = strconv.Itoa(len(openbalances)) + "/" + strconv.Itoa(len(compList)*len(yearList)) + " were added"
		r.Data = result1.Error
		tx.Rollback()
		return r
	}

	// create Service Party
	insert := tx.Create(&i)
	if insert.Error != nil || insert.RowsAffected == 0 {
		r.StatusCode = http.StatusBadRequest
		r.IsSuccess = false
		r.Message = "Service party create failed"
		r.Data = insert.Error
		tx.Rollback()
		return r
	}

	r.IsSuccess = true
	r.StatusCode = http.StatusOK
	r.Message = "Service party created"
	r.Data = i
	tx.Commit()

	return r
}

func (spr *ServicePartyRepo) UpdateParty(i dto.UpdateParty) dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	tx := db.Begin()

	i.PartyName = util.TrimString(i.PartyName)
	i.PartyAddress = util.TrimString(i.PartyAddress)
	i.PartyPhone = util.TrimString(i.PartyPhone)
	i.PartyEmail = util.TrimString(i.PartyEmail)

	//validate data
	r.IsSuccess, r.Message = ValidateServiceParty(i.ServiceParty)
	if !r.IsSuccess {
		r.StatusCode = http.StatusBadRequest
		return r
	} else if i.ChangeUser == "" {
		r.IsSuccess = false
		r.Message = "User name required"
		r.StatusCode = http.StatusBadRequest
		return r
	}
	// get old data
	var oldData model.ServiceParty
	checkId := db.Where("party_code", i.PartyCode).Find(&oldData)
	if checkId.Error != nil || checkId.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "User id not found"
		r.Data = checkId.Error
		return r
	}
	// create archive table
	var archiveTable model.ServicePartyArchive
	archiveTable.ServiceParty = oldData
	archiveTable.ChangeDate,archiveTable.ChangeTime = util.GenerateArchiveDateTime()
	archiveTable.ChangeEvent = "Update"
	archiveTable.ChangeUser = i.ChangeUser
	// get track id from archive table
	_ = db.Raw("select coalesce ((max(track_id) + 1), 1) as \"track_id\" from iservice.service_party_archive").Find(&archiveTable.TrackId)

	createArchive := tx.Create(&archiveTable)
	if createArchive.Error != nil || createArchive.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong. Please try again"
		r.Data = createArchive.Error
		tx.Rollback()
		return r
	}
	// if change in Party Name, change hadname of openbalance
	if oldData.PartyName != i.PartyName {
		changeOpenBalance := tx.Exec(`UPDATE iservice.openbalance SET headname=? WHERE accid=?`, i.PartyName, i.PartyCode)
		if changeOpenBalance.Error != nil || changeOpenBalance.RowsAffected == 0 {
			r.IsSuccess = false
			r.StatusCode = http.StatusNotFound
			r.Message = "Something went wrong. Please try again"
			r.Data = changeOpenBalance.Error
			tx.Rollback()
			return r
		}
	}

	// update main table
	update := map[string]interface{}{
		"party_name":    i.PartyName,
		"party_nid":     i.PartyNid,
		"party_phone":   i.PartyPhone,
		"party_email":   i.PartyEmail,
		"party_address": i.PartyAddress,
	}
	updateTable := tx.Model(&oldData).Where("party_code", i.PartyCode).Updates(update)
	if updateTable.Error != nil || updateTable.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong. Please try again"
		r.Data = updateTable.Error
		tx.Rollback()
		return r
	} else {
		r.IsSuccess = true
		r.StatusCode = http.StatusOK
		r.Message = "Update Successful"
		r.Data = i.ServiceParty
		tx.Commit()
	}
	return r
}

func (spr *ServicePartyRepo) DeleteOneParty(i dto.DeleteParty) dto.ResponseDtoV2 {
	var r dto.ResponseDtoV2
	db := util.CreateConnectionUsingGormToiServiceSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	tx := db.Begin()

	if i.ChangeUser == "" {
		r.IsSuccess = false
		r.Message = "User name required"
		r.StatusCode = http.StatusBadRequest
		return r
	}

	// get old data
	var oldData model.ServiceParty
	checkId := db.Where("party_code", i.PartyCode).Find(&oldData)
	if checkId.Error != nil || checkId.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "User id not found"
		r.Data = checkId.Error
		return r
	}

	// create archive table
	var archiveTable model.ServicePartyArchive
	archiveTable.ServiceParty = oldData
	archiveTable.ChangeDate,archiveTable.ChangeTime = util.GenerateArchiveDateTime()
	archiveTable.ChangeEvent = "Delete"
	archiveTable.ChangeUser = i.ChangeUser
	// get track id from archive table
	_ = db.Raw("select coalesce ((max(track_id) + 1), 1) as \"track_id\" from iservice.service_party_archive").Find(&archiveTable.TrackId)

	createArchive := tx.Create(&archiveTable)
	if createArchive.Error != nil || createArchive.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong. Please try again"
		r.Data = createArchive.Error
		tx.Rollback()
		return r
	}

	// delete iservice.openbalance data
	deleteOpenBalance := tx.Exec(`DELETE FROM iservice.openbalance WHERE accid=?`, i.PartyCode)
	if deleteOpenBalance.Error != nil || deleteOpenBalance.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong. Please try again"
		r.Data = deleteOpenBalance.Error
		tx.Rollback()
		return r
	}

	// delete main table
	delete := tx.Model(&oldData).Where("party_code", i.PartyCode).Delete(&oldData)
	if delete.Error != nil || delete.RowsAffected == 0 {
		r.IsSuccess = false
		r.StatusCode = http.StatusNotFound
		r.Message = "Something went wrong. Please try again"
		r.Data = delete.Error
		tx.Rollback()
		return r
	}

	// save data
	r.IsSuccess = true
	r.StatusCode = http.StatusOK
	r.Message = "Delete Successful"
	r.Data = oldData
	tx.Commit()

	return r
}


func ValidateServiceParty(i model.ServiceParty) (bool, string) {
	if i.PartyCode == 0 {
		return false, "Party code required"
	}
	if i.PartyName == "" {
		return false, "Party name required"
	}
	if i.PartyPhone == "" {
		return false, "Party phone required"
	}
	if !util.IsPhoneValid(i.PartyPhone) {
		return false, "Invalid phone number"
	}
	if i.PartyEmail == "" {
		return false, "Party email required"
	}
	if !util.ValidateEmail(i.PartyEmail) {
		return false, "Invalid email"
	}
	return true, "OK"
}
