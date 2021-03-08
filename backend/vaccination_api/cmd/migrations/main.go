package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/divoc/api/config"
	"github.com/divoc/api/pkg"
	"github.com/divoc/api/pkg/db"
	"github.com/divoc/api/pkg/models"
	models2 "github.com/divoc/api/swagger_gen/models"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

var enrollmentsCreated = make(map[string]map[int]string)
var enrollmentsRowId = make(map[string]map[int]string)
var enrollmentDuplicates = make(map[string]map[int]int)

func main() {
	config.Initialize()
	db.Init()
	clickhouseClient := initClickhouse()
	duplicatesFile, _ := os.OpenFile("./cmd/migrations/duplicates.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	for _, filePath := range []string{"./cmd/migrations/sample.csv"} {
		processedFile, err := os.OpenFile(filePath+"processed", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		file, _ := os.Open(filePath)
		data := pkg.NewScanner(file)
		for data.Scan() {
			rowID := data.Text("ID")
			name := data.Text("name")
			certificate := data.Text("certificate")
			mobile := data.Text("mobile")
			preEnrollmentCode := data.Text("preEnrollmentCode")
			osCreatedAt := data.Text("osCreatedAt")
			certificateId := data.Text("certificateId")
			recordedDose := data.Text("dose")
			var cert models.Certificate
			if err := json.Unmarshal([]byte(certificate), &cert); err == nil {
				dose := cert.Evidence[0].Dose
				if recordedDose != strconv.Itoa(dose) {
					if _, alreadyProcessed := enrollmentsCreated[preEnrollmentCode][dose]; alreadyProcessed {
						currentMatched := []string{rowID, preEnrollmentCode, strconv.Itoa(dose), osCreatedAt}
						if enrollmentDuplicates[preEnrollmentCode][dose] == 1 {
							firstMatched := []string{enrollmentsRowId[preEnrollmentCode][dose], preEnrollmentCode, strconv.Itoa(dose), enrollmentsCreated[preEnrollmentCode][dose]}
							if _, err = duplicatesFile.WriteString(strings.Join(firstMatched, ",") + "\n"); err != nil {
								log.Error(err)
							}
							if _, err = duplicatesFile.WriteString(strings.Join(currentMatched, ",") + "\n"); err != nil {
								log.Error(err)
							}
						} else {
							if _, err = duplicatesFile.WriteString(strings.Join(currentMatched, ",") + "\n"); err != nil {
								log.Error(err)
							}
						}
						enrollmentDuplicates[preEnrollmentCode][dose] += 1
					} else {
						enrollmentsCreated[preEnrollmentCode] = map[int]string{}
						enrollmentsRowId[preEnrollmentCode] = map[int]string{}
						enrollmentDuplicates[preEnrollmentCode] = map[int]int{}
						enrollmentsCreated[preEnrollmentCode][dose] = osCreatedAt
						enrollmentsRowId[preEnrollmentCode][dose] = rowID
						enrollmentDuplicates[preEnrollmentCode][dose] = 1
					}
					updateDose(rowID, dose)
					certifiedMessage := models.CertifiedMessage{
						Name:              name,
						Contact:           nil,
						Mobile:            mobile,
						PreEnrollmentCode: preEnrollmentCode,
						CertificateId:     certificateId,
						Certificate:       &cert,
						Meta:              models2.CertificationRequestV2Meta{},
					}
					data, err := json.Marshal(certifiedMessage)
					if err != nil {
						log.Errorf("certification marshal error %+v", err)
					}
					_ = saveCertifiedEventV1(clickhouseClient, string(data))
				}
			} else {
				log.Errorf("Error in getting certificate %+v", err)
			}
			if _, err = processedFile.WriteString(data.Row[0] + "\n"); err != nil {
				log.Error(err)
			}
		}
		_ = processedFile.Close()
		_ = file.Close()
	}
	_ = duplicatesFile.Close()
}

func updateDose(rowId string, dose int) {
	db.GetDB().Exec(`UPDATE "V_VaccinationCertificate" SET dose=? WHERE ID=?`, dose, rowId)
}

func initClickhouse() *sql.DB {
	dbConnectionInfo := config.Config.Clickhouse.Dsn
	log.Infof("Using the db %s", dbConnectionInfo)
	connect, err := sql.Open("clickhouse", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Errorf("Error in pinging the server %s", err)
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Errorf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)

		} else {
			log.Error(err)
		}
		panic("Error in pinging the database")
	}
	log.Infof("%+v", connect)
	return connect
}
func saveCertifiedEventV1(connect *sql.DB, msg string) error {
	var certifiedMessage models.CertifiedMessage
	if err := json.Unmarshal([]byte(msg), &certifiedMessage); err != nil {
		log.Errorf("Kafka message unmarshalling error %+v", err)
		return errors.New("kafka message unmarshalling failed")
	}
	if certifiedMessage.Certificate == nil {
		log.Infof("Ignoring invalid message %+v", msg)
		return nil
	}
	if certifiedMessage.Meta.VaccinationApp == nil {
		certifiedMessage.Meta.VaccinationApp = &models2.CertificationRequestV2MetaVaccinationApp{}
	}
	// push to click house - todo: batch it
	var (
		tx, _     = connect.Begin()
		stmt, err = tx.Prepare(`INSERT INTO certifiedv1 
	(  certificateId,
  preEnrollmentCode,
  dt,

  age,
  gender,
  district,
  state,
  
  batch,
  vaccine,
  manufacturer,
  vaccinationDate,
  effectiveStart,
  effectiveUntil,
  dose,
  totalDoses,
  
  facilityName,
  facilityCountryCode,
  facilityState,
  facilityDistrict,
  facilityPostalCode,
  vaccinatorName,
  
  vaccinationAppName,
  vaccinationAppVersion,
  vaccinationAppType,
  vaccinationAppDevice,
  vaccinationAppDeviceOs,
  vaccinationAppOSVersion,
  vaccinationAppMode,
  vaccinationAppConnectionType,
  
  facilityType,
  paymentType,
  registrationCategory,
  registrationDataMode,
  sessionDurationInMinutes,
  uploadTimestamp,
  verificationAttempts,
  verificationDurationInSeconds,
  waitForVaccinationInMinutes) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	)

	if err != nil {
		log.Infof("Error in preparing stmt %+v", err)
	}
	//todo collect n messages and batch write to analytics db.
	credentialSubject := certifiedMessage.Certificate.CredentialSubject
	age, _ := strconv.Atoi(credentialSubject.Age)
	evidence := certifiedMessage.Certificate.Evidence[0]
	if _, err := stmt.Exec(
		certifiedMessage.CertificateId,
		certifiedMessage.PreEnrollmentCode,
		time.Now(),

		age,
		credentialSubject.Gender,
		credentialSubject.Address.District,
		credentialSubject.Address.AddressRegion,

		evidence.Batch,
		evidence.Vaccine,
		evidence.Manufacturer,
		evidence.Date,
		evidence.EffectiveStart,
		evidence.EffectiveUntil,
		evidence.Dose,
		evidence.TotalDoses,

		evidence.Facility.Name,
		"IN",
		evidence.Facility.Address.AddressRegion,
		evidence.Facility.Address.District,
		certifiedMessage.Certificate.GetFacilityPostalCode(),
		evidence.Verifier.Name,

		certifiedMessage.Meta.VaccinationApp.Name,
		certifiedMessage.Meta.VaccinationApp.Version,
		certifiedMessage.Meta.VaccinationApp.Type,
		certifiedMessage.Meta.VaccinationApp.Device,
		certifiedMessage.Meta.VaccinationApp.DeviceOS,
		certifiedMessage.Meta.VaccinationApp.OSVersion,
		certifiedMessage.Meta.VaccinationApp.AppMode,
		certifiedMessage.Meta.VaccinationApp.ConnectionType,

		certifiedMessage.Meta.FacilityType,
		certifiedMessage.Meta.PaymentType,
		certifiedMessage.Meta.RegistrationCategory,
		certifiedMessage.Meta.RegistrationDataMode,
		certifiedMessage.Meta.SessionDurationInMinutes,
		getDate(certifiedMessage.Meta.UploadTimestamp.String()),
		certifiedMessage.Meta.VerificationAttempts,
		certifiedMessage.Meta.VerificationDurationInSeconds,
		certifiedMessage.Meta.WaitForVaccinationInMinutes,
	); err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func getDate(dateTime string) time.Time {
	if dTime, err := time.Parse(time.RFC3339, dateTime); err == nil {
		return dTime
	}
	return time.Now()
}
