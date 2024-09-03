package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/storage/mongodb"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	configFile = "config.yaml"

	envKey = "TS_ENV" // env variable name to get deployment environment
)

type serviceIdentifier string

type configuration struct {
	Env      string    `yaml:"env"`
	Secret   string    `yaml:"secret" json:"secret"`
	PortHTTP uint16    `yaml:"portHTTP"`
	PortGRPC uint16    `yaml:"portGRPC"`
	Database *database `yaml:"database" json:"database"`
	Cors     *cors     `yaml:"cors" json:"cors"`
}

type SinglePaymentLink struct {
	BillingPeriod string  `json:"billingPeriod" yaml:"billingPeriod"`
	Currency      string  `json:"currency" yaml:"currency"`
	PriceID       string  `json:"priceID" yaml:"priceID"`
	PaymentLink   string  `json:"paymentLink" yaml:"paymentLink"`
	Amount        float64 `json:"amount" yaml:"amount"`
}

type serviceConfig struct {
	URL  string `yaml:"url" json:"url"`
	Port uint16 `yaml:"port" json:"port"`
}

type database struct {
	MongoDB *mongoDB `yaml:"mongoDB" json:"mongoDB"`
}

type cors struct {
	AllowedOrigins []string `yaml:"allowedOrigins" json:"allowedOrigins"`
	AllowedMethods []string `yaml:"allowedMethods" json:"allowedMethods"`
	AllowedHeaders []string `yaml:"allowedHeaders" json:"allowedHeaders"`
	ExposeHeaders  []string `yaml:"exposeHeaders" json:"exposeHeaders"`
}

type mongoDB struct {
	Username                string `yaml:"username" json:"username"`
	Password                string `yaml:"password" json:"password"`
	DBName                  string `yaml:"dbName" json:"dbName"`
	Endpoint                string `yaml:"endpoint" json:"endpoint"`
	Client                  *mongo.Client
	TriggerAPIRenewalSecret string `yaml:"triggerAPIRenewalSecret" json:"triggerAPIRenewalSecret"`
}

type dynamoDB struct {
	Region string `yaml:"region" json:"region"`
	Tables tables `yaml:"tables" json:"tables"`
}

type tables struct {
	Users         string `yaml:"users" json:"users"`
	APITokens     string `yaml:"apiTokens" json:"apiTokens"`
	APILog        string `yaml:"apiLog" json:"apiLog"`
	APIConfig     string `yaml:"apiConfig" json:"apiConfig"`
	Teams         string `yaml:"teams" json:"teams"`
	TeamAPIConfig string `yaml:"teamAPIConfig" json:"teamAPIConfig"`
}

type aws struct {
	Region     string      `yaml:"region"`
	S3         *s3         `yaml:"s3" json:"s3"`
	Cognito    *cognito    `yaml:"cognito" json:"cognito"`
	SQS        sqs         `yaml:"sqs" json:"sqs"`
	Cloudfront *cloudfront `yaml:"cloudfront" json:"cloudfront"`
	Cloudwatch *cloudwatch `yaml:"cloudwatch" json:"cloudwatch"`
}

type s3 struct {
	FalconImageBucketName      string `yaml:"falconImageBucketName" json:"falconImageBucketName"`
	ReportBucketName           string `yaml:"reportBucketName" json:"reportBucketName"`
	TempBucketName             string `yaml:"tempBucketName" json:"tempBucketName"`
	ExportThreatFeedBucketName string `yaml:"exportThreatFeedBucketName" json:"exportThreatFeedBucketName"`
}

type cloudflare struct {
	Turnstile turnstile         `yaml:"turnstile" json:"turnstile"`
	Worker    *cloudflareWorker `yaml:"worker" json:"worker"`
}

type cloudwatch struct {
	LogGroupArgo         string `yaml:"logGroupArgo" json:"logGroupArgo"`
	LogStreamArgoWebhook string `yaml:"logStreamArgoWebhook" json:"logStreamArgoWebhook"`
	LogStreamArgoEmail   string `yaml:"logStreamArgoEmail" json:"logStreamArgoEmail"`
}

type cloudflareWorker struct {
	WebhookWorker webhookWorker `yaml:"webhookWorker" json:"webhookWorker"`
	CVEWorker     cveWorker     `yaml:"cveWorker" json:"cveWorker"`
}

type webhookWorker struct {
	URL    string `yaml:"url" json:"url"`
	Token  string `yaml:"token" json:"token"`
	Method string `yaml:"method" json:"method"`
}

type cveWorker struct {
	Secret string `yaml:"secret" json:"secret"`
}

type turnstile struct {
	Secret string `yaml:"secret" json:"secret"`
}

type cognito struct {
	UserPoolID   string `yaml:"userPoolID" json:"userPoolID"`
	ClientID     string `yaml:"clientID" json:"clientID"`
	Domain       string `yaml:"domain" json:"domain"`
	ClientSecret string `yaml:"clientSecret" json:"clientSecret"`
	CallbackURL  string `yaml:"callbackURL" json:"callbackURL"`
}

type sqs struct {
	FalconWorkerQueue falconWorkerQueue `yaml:"falconWorkerQueue" json:"falconWorkerQueue"`
	HawkQueue         hawkQueue         `yaml:"hawkQueue" json:"hawkQueue"`
}

type falconWorkerQueue struct {
	URL         string `yaml:"url" json:"url"`
	MaxMessages int    `yaml:"maxMessages" json:"maxMessages"`
}

type hawkQueue struct {
	URL         string `yaml:"url" json:"url"`
	MaxMessages int    `yaml:"maxMessages" json:"maxMessages"`
}

type acl struct {
	FeatureAccess *featureAccess `yaml:"featureAccess" json:"featureAccess"`
}

type featureAccess struct {
	ThreatFeed *threatFeed `yaml:"threatFeed" json:"threatFeed"`
}

type threatFeed struct {
	FreePlanDateLimit string `yaml:"freePlanDateLimit" json:"freePlanDateLimit"`
}

type cloudfront struct {
	FalconImagesBaseURL string `json:"falconImagesBaseURL" yaml:"falconImagesBaseURL"`
}

var config *configuration

func Init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	//env := os.Getenv(envKey)
	//if !utils.IsEmptyString(env) {
	//	config.Env = env
	//}

	//if utils.IsEnvSTG(config.Env) || utils.IsEnvProd(config.Env) {
	//	if err := overwriteConfigFromSecrets(); err != nil {
	//		panic(err)
	//	}
	//}

	initMongoDB()
}

func loadConfig() error {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	return nil
}

// overwriteConfigFromSecrets ... load configuration values from AWS Secrets Manager
//func overwriteConfigFromSecrets() error {
//	secretName := fmt.Sprintf("%s/%s", config.Env, service.Falcon)
//
//	aws := libsAWS.New(config.AWS.Region, context.Background())
//	secrets, err := aws.SecretsManager.GetSecret(context.TODO(), secretName)
//	if err != nil {
//		return err
//	}
//
//	err = json.Unmarshal([]byte(*secrets), config)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// Get ... return Configuration struct with values
func Get() *configuration {
	if config == nil {
		panic("Configuration not loaded. Application crashed!")
	}
	return config
}

//// GetServiceConfig ... returns url, port and other config for a grpc service
//func GetServiceConfig(identifier string) serviceConfig {
//	s := config.Services[serviceIdentifier(identifier)]
//	return s
//}

func initMongoDB() {
	mongoDBClient := mongodb.NewClient(
		config.Env,
		config.Database.MongoDB.Username,
		config.Database.MongoDB.Password,
		config.Database.MongoDB.DBName,
		config.Database.MongoDB.Endpoint)

	config.Database.MongoDB.Client = mongoDBClient
}
