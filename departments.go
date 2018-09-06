package wdmsapi

type DepartmnetParams struct {
	ZoneNumber int `schema:"zoneNumber,omitempty"`
	Page       int `schema:"page,omitempty"`
	PerPage    int `schema:"per_page,omitempty"`
}

type DepartmentData struct {
	DepartmentNumber int    `json:"departmentNumber"`
	ZoneName         string `json:"zoneName"`
	DepartmentName   string `json:"departmentName"`
}

type DepartmentRecords struct {
	Data       []DepartmentData `json:"data"`
	Page       int              `json:"page"`
	TotalItems int              `json:"totalItems"`
	Code       int              `json:"code"`
	TotalPages int              `json:"totalPages"`
	PageSize   int              `json:"pageSize"`
	NextUrl    string           `json:"next_url"`
	Message    string           `json:"message"`
}

type DepartmentUpsert struct {
	Data []DepartmentData `json:"Data"`
}

type DepartmentDelete struct {
	DepartmentNumber []int `json:"departmentNumber"`
}
