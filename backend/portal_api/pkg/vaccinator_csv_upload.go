package pkg

import (
	"errors"
	"github.com/divoc/kernel_library/services"
	"github.com/divoc/portal-api/config"
	"github.com/divoc/portal-api/pkg/db"
	"github.com/divoc/portal-api/swagger_gen/models"
	"strconv"
	"strings"
)

type VaccinatorCSV struct {
	CSVMetadata
}

func (vaccinatorCSV VaccinatorCSV) CreateCsvUploadHistory() *db.CSVUploads {
	return vaccinatorCSV.CSVMetadata.CreateCsvUploadHistory("Vaccinator")
}

func (vaccinatorCSV VaccinatorCSV) ValidateRow() []string {
	requiredHeaders := strings.Split(config.Config.Vaccinator.Upload.Required, ",")
	return vaccinatorCSV.CSVMetadata.ValidateRow(requiredHeaders)
}

func (vaccinatorCSV VaccinatorCSV) CreateCsvUpload() error {
	data := vaccinatorCSV.Data
	serialNum, err := strconv.ParseInt(data.Text("serialNum"), 10, 64)
	if err != nil {
		return err
	}
	mobileNumber := data.Text("mobileNumber")
	nationalIdentifier := data.Text("nationalIdentifier")
	code := data.Text("code")
	name := data.Text("name")
	status := data.Text("status")
	facilityIds := strings.Split(data.Text("facilityIds"), ",")
	averageRating := 0.0
	trainingCertificate := ""
	vaccinator := models.Vaccinator{
		SerialNum:           &serialNum,
		MobileNumber:        &mobileNumber,
		NationalIdentifier:  &nationalIdentifier,
		Code:                &code,
		Name:                &name,
		Status:              &status,
		FacilityIds:         facilityIds,
		AverageRating:       &averageRating,
		Signatures:          []*models.Signature{},
		TrainingCertificate: &trainingCertificate,
	}
	//services.MakeRegistryCreateRequest(vaccinator, "Vaccinator")
	errorVaccinator := services.CreateNewRegistry(vaccinator, "Vaccinator")
	if errorVaccinator != nil {
		errmsg := errorVaccinator.Error()
		if strings.Contains(errmsg, "Detail:") {
			split := strings.Split(errmsg, "Detail:")
			if len(split) > 0 {
				m1 := split[len(split)-1]
				return errors.New(m1)
			}
		}
		return errors.New(errmsg)
	}
	return nil
}

func (vaccinatorCSV VaccinatorCSV) SaveCsvErrors(rowErrors []string, csvUploadHistoryId uint) {
	vaccinatorCSV.CSVMetadata.SaveCsvErrors(rowErrors, csvUploadHistoryId)
}
