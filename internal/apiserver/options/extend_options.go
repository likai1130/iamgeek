package options

import "time"

var ExtendConf *ExtendOptions

/**
	  sms:
    sign-name: "SIG20210825B1BAB9"
    template-code: "UTA21082708A7F9"
    public-key: "4aoTwELq1JMIPOCu7KPj9G2QmlYjmgFIy4Q6H4nNY"
    private-key: "KAImWt7Rucu62z9IitWmDX40HqVQ4DQj5HQ2gGj2aFUn4d2FXWbOzol6YfAuTFXC7d"
    project-id: "org-vhxub0"
  mongodb:
    hosts: [ "127.0.0.1:27017" ]
    max-pool-size: "100"
    username: ""
    password: ""
    dbname: "winery-p2p-market"
    replica-set: ""
    #primary,primaryPreferred,secondary,secondaryPreferred,nearest
    read-preference: ""
*/
type ExtendOptions struct {
	SMS     SMS            `json:"sms"`
	MongoDB MongoDBOptions `json:"mongodb"`
}

type SMS struct {
	SignName     string `json:"sign-name" mapstructure:"sign-name"`
	TemplateCode string `json:"template-code" mapstructure:"template-code"`
	PublicKey    string `json:"public-key" mapstructure:"public-key"`
	PrivateKey   string `json:"private-key" mapstructure:"private-key"`
	ProjectId    string `json:"project-id" mapstructure:"project-id"`
}

type MongoDBOptions struct {
	Hosts                   []string      `json:"hosts" mapstructure:"hosts"`
	UserName                string        `json:"username" mapstructure:"user_name"`
	Password                string        `json:"password" mapstructure:"password"`
	MaxPoolSize             uint64        `json:"max-pool-size" mapstructure:"max-pool-size"`
	DbName                  string        `json:"dbname" mapstructure:"dbname"`
	ReplicaSet              string        `json:"replica-set" mapstructure:"replica-set"`
	ReadPreference          string        `json:"read-preference" mapstructure:"read-preference"`
	ServerSelectionTimeoutS time.Duration `json:"server-selection-timeout-s"`
}
