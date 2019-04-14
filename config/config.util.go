package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
)

var config Config

func ReadConfiguration() {

	yamlFile, errFile := ioutil.ReadFile("config.yml")
	if errFile != nil {
		fmt.Println("Error reading configuration file! " + errFile.Error())
		os.Exit(1)
	}
	errMarshal := yaml.Unmarshal(yamlFile, &config)
	if errMarshal != nil {
		fmt.Println("Error parsing yaml file! " + errMarshal.Error())
		os.Exit(2)
	}
	overrideWithEnv()
	fmt.Println("Configuration was loaded!")
}

func overrideWithEnv() {
	serviceName, serviceNamePresent := os.LookupEnv("SERVICE_NAME")
	if serviceNamePresent {
		config.Service.Name = serviceName
	}
	serviceVersion, serviceVersionPresent := os.LookupEnv("SERVICE_VERSION")
	if serviceVersionPresent {
		config.Service.Version = serviceVersion
	}
	serviceEnv, serviceEnvPresent := os.LookupEnv("SERVICE_ENV")
	if serviceEnvPresent {
		config.Service.Env = serviceEnv
	}

	serverPort, serverPortPresent := os.LookupEnv("SERVER_PORT")
	if serverPortPresent {
		port, portParseErr := strconv.Atoi(serverPort)
		if portParseErr != nil {
			fmt.Println("Port cannot be parsed to integer!")
			os.Exit(2)
		}
		config.Server.Port = port
	}

	serverBaseUrl, serverBaseUrlPresent := os.LookupEnv("SERVER_BASEURL")
	if serverBaseUrlPresent {
		config.Server.BaseUrl = serverBaseUrl
	}

	serverSessionDuration, serverSessionDurationPresent := os.LookupEnv("SERVER_SESSIONDURATION")
	if serverSessionDurationPresent {
		duration, durationParseErr := strconv.Atoi(serverSessionDuration)
		if durationParseErr != nil {
			fmt.Println("Session duration cannot be parsed to integer!")
			os.Exit(2)
		}
		if duration <= 0 {
			duration = 3600
		}
		config.Server.SessionDuration = duration
	}

	//datasourceType, datasourceTypePresent := os.LookupEnv("DATASOURCE_TYPE")
	//if datasourceTypePresent {
	//	config.Datasource.Type = datasourceType
	//}

	datasourceLocation, datasourceLocationPresent := os.LookupEnv("DATASOURCE_LOCATION")
	if datasourceLocationPresent {
		config.Datasource.Location = datasourceLocation
	}

}

func GetConfiguration() Config {
	return config
}
