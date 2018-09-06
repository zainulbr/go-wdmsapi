package wdmsapi

type Account struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type DeviceParams struct {
	SN               string `schema:"sn,omitempty"`
	Status           int    `schema:"status,omitempty"`
	DepartmentNumber int    `schema:"departmentNumbe,omitempty"`
	Page             int    `schema:"page,omitempty"`
	PerPage          int    `schema:"per_page,omitempty"`
}

type DeviceData struct {
	Status           string `json:"status"`
	DeviceName       string `json:"deviceName"`
	FaceCount        int    `json:"faceCount"`
	DepartmentName   string `json:"departmentName"`
	ZoneNumber       int    `json:"zoneNumber"`
	CmdCount         int    `json:"cmdCount"`
	TransactionCount int    `json:"transactionCount"`
	Alias            string `json:"alias"`
	FpCount          int    `json:"fpCount"`
	LastActivity     string `json:"lastActivity"`
	SN               string `json:"sn"`
	ZoneName         string `json:"zoneName"`
	FirmwareVersion  string `json:"firmwareVersion"`
	UserCount        int    `json:"userCount"`
}

type DeviceRecord struct {
	Data       []DeviceData `json:"data"`
	Page       int          `json:"page"`
	TotalItems int          `json:"totalItems"`
	Code       int          `json:"code"`
	TotalPages int          `json:"totalPages"`
	PageSize   int          `json:"pageSize"`
	NextUrl    string       `json:"next_url"`
	Message    string       `json:"message"`
}

type DeviceUpsert struct {
	DepartmentName string `json:"departmentName"`
	Alias          string `json:"alias"`
	SN             string `json:"sn"`
	MasterDevice   string `json:"masterDevice"`
	FacialDevice   string `json:"facialDevice"`
}

type DeviceDelete struct {
	SN []string `json:"sn"`
}
