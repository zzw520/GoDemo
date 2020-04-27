package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MessageConfig struct {
	Sms Sms `yaml:"sms"`
}

type Sms struct {
	SmsTopics map[string]SmsTopic `yaml:"smsTopics"`
}

type SmsTopic struct {
	Provider string   `yaml:"provider"`
	Endpoint string   `yaml:"endpoint"`
	To       []string `yaml:"to"`
}

func main() {
	bytes, err := ioutil.ReadFile("yaml/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	var messageConfig MessageConfig
	if err = unmarshal(bytes, &messageConfig, false); err != nil {
		log.Fatal(err.Error())
		return
	}
	i := messageConfig.Sms.SmsTopics["ec"]
	log.Print(i)

}

func unmarshal(in []byte, out interface{}, isStrict bool) (err error) {
	if in == nil {
		err = errors.New("can't unmarshal empty byte")
		return err
	}
	if isStrict == true {
		err = yaml.UnmarshalStrict(in, out)
		if err != nil {
			return err
		}
	} else {
		err = yaml.Unmarshal(in, out)
		if err != nil {
			return err
		}
	}
	return nil
}
