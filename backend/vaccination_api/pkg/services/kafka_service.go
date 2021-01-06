package services

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/divoc/api/config"
	"github.com/divoc/api/pkg/db"
	"github.com/divoc/api/swagger_gen/models"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var producer *kafka.Producer

var messages = make(chan Message)
var events = make(chan []byte)
var reportedSideEffects = make(chan []byte)

type Event struct {
	Date          time.Time   `json:"date"`
	Source        string      `json:"source"`
	TypeOfMessage string      `json:"type"`
	ExtraInfo     interface{} `json:"extra"`
}
type Message struct {
	UploadId []byte
	rowId    []byte
	payload  string
}
type ReportedSideEffectsEvent struct {
	models.SideEffectsResponse
	RecipientCertificateId string `json:"recipientCertificateId"`
}

func InitializeKafka() {
	servers := config.Config.Kafka.BootstrapServers
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": servers})
	if err != nil {
		panic(err)
	}
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  servers,
		"group.id":           "certify_ack",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": "false",
	})
	if err != nil {
		panic(err)
	}
	log.Infof("Connected to kafka on %s", servers)

	//defer func() {
	//	log.Info("Closing the producer!")
	//	producer.Close()
	//}()

	go func() {
		topic := config.Config.Kafka.CertifyTopic
		for {
			msg := <-messages
			if err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(msg.payload),
				Headers: []kafka.Header{
					{Key: "uploadId", Value: msg.UploadId},
					{Key: "rowId", Value: msg.rowId},
				},
			}, nil); err != nil {
				log.Infof("Error while publishing message to %s topic %+v", topic, msg)
			}
		}
	}()

	go func() {
		topic := config.Config.Kafka.EventsTopic
		for {
			msg := <-events
			if err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          msg,
			}, nil); err != nil {
				log.Infof("Error while publishing message to %s topic %+v", topic, msg)
			}
		}
	}()

	go func() {
		topic := config.Config.Kafka.ReportedSideEffectsTopic
		for {
			msg := <-reportedSideEffects
			if err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          msg,
			}, nil); err != nil {
				log.Infof("Error while publishing message to %s topic %+v", topic, msg)
			}
		}
	}()

	go func() {
		consumer.SubscribeTopics([]string{"certify_ack"}, nil)

		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				var message map[string]string
				json.Unmarshal(msg.Value, &message)
				// check the status
				// update that status to certifyErrorRows db
				log.Infof("Message on %s: %v \n", msg.TopicPartition, message)
				if message["rowId"] == "" {
					// ignoring rows which doesnt have rowId
					consumer.CommitMessage(msg)
				} else {
					rowId, e := strconv.ParseUint(message["rowId"], 10, 64)
					if e != nil {
						log.Errorf("Error occurred wile parsing rowId as int - %s", message["rowId"])
					} else {
						if message["status"] == "SUCCESS" {
							// if certificate created successfully
							// delete that row => as we no longer require that row
							db.DeleteCertifyUploadError(uint(rowId))
						} else if message["status"] == "FAILED" {
							// if certificate creation fails
							// update the status of the row to Failed
							db.UpdateCertifyUploadErrorStatusAndErrorMsg(uint(rowId), db.CERTIFY_UPLOAD_FAILED_STATUS, message["errorMsg"])
						}
						consumer.CommitMessage(msg)
					}
				}
			} else {
				// The client will automatically try to recover from all errors.
				log.Infof("Consumer error: %v \n", err)
			}
		}
	}()

	go func() {
		for e := range producer.Events() {
			log.Infof("%+v", e)
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Infof("Delivery failed: %v\n%+v", ev.TopicPartition, ev.Value)
				} else {
					log.Infof("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}

	}()
}

func PublishCertifyMessage(message []byte, uploadId []byte, rowId []byte) {
	messages <- Message{
		UploadId: uploadId,
		rowId:    rowId,
		payload:  string(message),
	}
}

func PublishEvent(event Event) {
	if messageJson, err := json.Marshal(event); err != nil {
		log.Errorf("Error in getting json of event %+v", event)
	} else {
		events <- messageJson
	}
}

func PublishReportedSideEffects(event ReportedSideEffectsEvent) {
	log.Infof("Publishing reported side effects")
	if messageJson, err := json.Marshal(event); err != nil {
		log.Errorf("Error in getting json of event %+v", event)
	} else {
		reportedSideEffects <- messageJson
	}
	log.Infof("Successfully published reported side Effects")
}
