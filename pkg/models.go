package catalog

// Flavour represents a subset of a node's resources
type Flavour struct {
	UID             string `json:"uid"`
	Architecture    string `json:"architecture"`
	OperatingSystem string `json:"operatingsystem"`
	CPUOffer        string `json:"cpuoffer"`
	MemoryOffer     string `json:"memoryoffer"`
	PodsOffer       []Plan `json:"podsoffer"`
}

// PodsPlan represents a plan for which is possibile to have a specific amount of available pods
type Plan struct {
	Name      string `json:"name"`
	Available bool   `json:"available"`
	Pods      int    `json:"pods"`
}

// Request represents a request to match specific Flavours
type Request struct {
	Architecture  string `json:"architecture"`
	OS            string `json:"os"`
	CPURequest    string `json:"cpuRequest"`
	MemoryRequest string `json:"memoryRequest"`
	PodsPlan      string `json:"podsPlan"`
}
