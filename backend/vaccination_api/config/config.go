package config

import (
	"errors"
	"github.com/imroc/req"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

func Initialize() {
	err := configor.Load(&Config, "./config/application-default.yml") //"config/application.yml"

	if err != nil {
		panic("Unable to read configurations")
	}
	if Config.Keycloak.Enable && Config.Keycloak.Pubkey == "" {
		updatePublicKeyFromKeycloak()
	}
}

func updatePublicKeyFromKeycloak() error {
	url := Config.Keycloak.Url + "/realms/" + Config.Keycloak.Realm
	log.Info("Public key url ", url)
	resp, err := req.Get(url)
	if err != nil {
		return err
	}
	log.Infof("Got response %+v", resp.String())
	responseObject := map[string]interface{}{}
	if err := resp.ToJSON(&responseObject); err == nil {
		if publicKey, ok := responseObject["public_key"].(string); ok {
			Config.Keycloak.Pubkey = publicKey
		}
	}
	return errors.New("Unable to get public key from keycloak")
}

var Config = struct {
	Registry struct {
		Url               string `default:"localhost:8081" env:"REGISTRY_URL"`
		AddOperationId    string
		SearchOperationId string `default:"search"`
		ReadOperationId   string `default:"read"`
		ApiVersion        string `default:"1"`
	}
	Keycloak struct {
		Url                  string `env:"KEYCLOAK_URL"`
		AdminApiClientSecret string `env:"ADMIN_API_CLIENT_SECRET"`
		Pubkey               string `env:"PUBLIC_KEY"`
		Realm                string `default:"divoc"`
		AuthHeader           string `env:"AUTH_TOKEN"`
		RecipientGroupId     string `default:"recipient"`
		Enable               bool   `env:"ENABLE_KEYCLOAK" default:"true"`
	}
	Kafka struct {
		BootstrapServers         string `env:"KAFKA_BOOTSTRAP_SERVERS" yaml:"bootstrapServers"`
		CertifyTopic             string `default:"certify" yaml:"certifyTopic"`
		CertifiedTopic           string `default:"certified" yaml:"certifiedTopic"`
		EventsTopic              string `default:"events" yaml:"eventsTopic"`
		ReportedSideEffectsTopic string `default:"reported_side_effects" yaml:"reportedSideEffectsTopic"`
		EnableCertificateAck     bool   `default:"false" env:"ENABLE_CERTIFICATE_ACK"`
	}
	Database struct {
		Host     string `default:"0.0.0.0" yaml:"host" env:"DB_HOST"`
		Password string `default:"postgres" yaml:"password" env:"DB_PASSWORD"`
		User     string `default:"postgres" yaml:"user" env:"DB_USER"`
		Port     string `default:"5432" yaml:"port" env:"DB_PORT"`
		DBName   string `default:"postgres" yaml:"dbname" env:"DB_NAME"`
	}
	Redis struct {
		Url      string `env:"REDIS_URL" default:"redis://0.0.0.0:6379"`
		Password string `default:"" env:"REDIS_PASSWORD"`
		DB       int    `default:"0" env:"REDIS_DB"`
	}
	Certificate struct {
		Upload struct {
			Columns string `yaml:"columns"`
		}
	}
	Clickhouse struct {
		Dsn string `env:"CLICK_HOUSE_URL"`
	}
	Digilocker struct {
		AuthKeyName string `env:"DIGILOCKER_AUTHKEYNAME" default:"x-digilocker-hmac"`
		AuthHMACKey string `env:"DIGILOCKER_HMAC_AUTHKEY"`
		DocType     string `env:"DIGILOCKER_DOCTYPE"`
	}
}{}
