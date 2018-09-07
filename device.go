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
	Status           string `json:"status" schema:"status,omitempty"`
	DeviceName       string `json:"deviceName" schema:"deviceName,omitempty"`
	FaceCount        int    `json:"faceCount" schema:"faceCount,omitempty"`
	DepartmentName   string `json:"departmentName" schema:"departmentName,omitempty"`
	ZoneNumber       int    `json:"zoneNumber" schema:"zoneNumber,omitempty"`
	CmdCount         int    `json:"cmdCount" schema:"cmdCount,omitempty"`
	TransactionCount int    `json:"transactionCount" schema:"transactionCount,omitempty"`
	Alias            string `json:"alias" schema:"alias,omitempty"`
	FpCount          int    `json:"fpCount" schema:"fpCount,omitempty"`
	LastActivity     string `json:"lastActivity" schema:"lastActivity,omitempty"`
	SN               string `json:"sn" schema:"sn,omitempty"`
	ZoneName         string `json:"zoneName" schema:"zoneName,omitempty"`
	FirmwareVersion  string `json:"firmwareVersion" schema:"firmwareVersion,omitempty"`
	UserCount        int    `json:"userCount" schema:"userCount,omitempty"`
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
