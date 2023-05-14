package catalog

type Flavour struct {
	UID             string `json:"uid"`
	Architecture    string `json:"architecture"`
	OperatingSystem string `json:"operatingsystem"`
	CPUOffer        string `json:"cpuoffer"`
	MemoryOffer     string `json:"memoryoffer"`
	PodsOffer       []Plan `json:"podsoffer"`
}

type Plan struct {
	Name      string `json:"name"`
	Available bool   `json:"available"`
	Pods      int    `json:"pods"`
}

type Request struct {
	Architecture  string `json:"architecture"`
	OS            string `json:"os"`
	CPURequest    string `json:"cpuRequest"`
	MemoryRequest string `json:"memoryRequest"`
	PodsPlan      string `json:"podsPlan"`
}
