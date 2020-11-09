package main

import(
	"encoding/json"
	"net/http"
	"log"
)

type Server struct{
	Service ServiceRegistry
	Database Database
}

type ServiceRegistry struct{
	Name string `json:"Name"`
	Hostname string `json:"hostname"`
	IpAddr string `json:"ipAddr"`
	Status string `json:"status"`
	Port int `json:"port"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}

func (self *ServiceRegistry) startServer() {
	http.HandleFunc("/service_registry", ServiceRegistryAPI)
	log.Fatal(http.ListenAndServe(":50001", nil))
}

func ServiceRegistryAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		service := ServiceRegistry{}
		err := decoder.Decode(&service)
		checkErr(err)

		log.Println(service)
	}
}