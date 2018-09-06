package wdmsapi

type responseData struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type FingerPrintParams struct {
	EmpID   int `schema:"empID,omitempty"`
	Page    int `schema:"page,omitempty"`
	PerPage int `schema:"per_page,omitempty"`
}

type FingerPrint struct {
	FpTemplate string `json:"fpTemplate"`
	Valid      int    `json:"valid"`
	FingerID   int    `json:"fingerID"`
	EmpID      int    `json:"empID"`
	Pin        string `json:"pin"`
}

type FingerPrintRecords struct {
	Data       []FingerPrint `json:"data"`
	Page       int           `json:"page"`
	TotalItems int           `json:"totalItems"`
	Code       int           `json:"code"`
	TotalPages int           `json:"totalPages"`
	PageSize   int           `json:"pageSize"`
	NextURL    string        `json:"next_url"`
	Message    string        `json:"message"`
}

type FaceTemplateParams struct {
	EmpID   int `schema:"empID,omitempty"`
	Page    int `schema:"page,omitempty"`
	PerPage int `schema:"per_page,omitempty"`
}

type FaceTemplate struct {
	FaceTemplate string `json:"faceTemplate"`
	FaceID       int    `json:"faceID"`
	Valid        int    `json:"valid"`
	EmpID        int    `json:"empID"`
	Pin          string `json:"pin"`
}

type FaceTemplateRecords struct {
	Data       []FaceTemplate `json:"data"`
	Page       int            `json:"page"`
	TotalItems int            `json:"totalItems"`
	Code       int            `json:"code"`
	TotalPages int            `json:"totalPages"`
	PageSize   int            `json:"pageSize"`
	NextURL    string         `json:"next_url"`
	Message    string         `json:"message"`
}

type SendEmployeesParams struct {
	Sn    []string `json:"sn"`
	EmpID []int    `json:"empID"`
}

const (
	FormatDateTime   = "2006-01-02 15:04:05"
	FormatDateTimeTZ = "2006-01-02 15:04:05 -0700"
)
