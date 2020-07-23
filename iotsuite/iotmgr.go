package iotsuite



import (
	"net/http"
)

func IotmgrRules(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/mme/rules"
	Get(httpClient, url)
}

func IotmgrTasks(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/mme/tasks"
	Get(httpClient, url)
}

func IotmgrDevices(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/di/devices"
	Get(httpClient, url)
}

func IotmgrGroups(httpClient *http.Client, conf *Configuration) {
	var url = "https://manager.eu-1.bosch-iot-suite.com/api/1/di/groups/directories"
	Get(httpClient, url)
}