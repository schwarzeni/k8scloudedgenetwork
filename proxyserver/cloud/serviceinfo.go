package main

type ServiceInfo struct {
	Name  string `json:"name"`
	Addrs []Addr `json:"addrs"`
}

type Addr struct {
	NodeID string `json:"node_id"`
	IP     string `json:"ip"`
	Port   string `json:"port"`
}

var fakeServiceInfo = map[string]ServiceInfo{
	"xxxx.svc": {
		Name:  "xxxx.svc",
		Addrs: []Addr{{"xxxx", "10.211.55.2", "8080"}},
	},
	"oooo.svc": {
		Name:  "oooo.svc",
		Addrs: []Addr{{"oooo", "10.211.55.52", "8080"}},
	},
}

func GetServiceInfo(svcName string) (*ServiceInfo, bool, error) {
	svc, ok := fakeServiceInfo[svcName]
	return &svc, ok, nil
}
