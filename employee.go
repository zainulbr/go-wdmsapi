package wdmsapi

type EmployeeParams struct {
	PIN              string `schema:"pin,omitempty"`
	DepartmentNumber int    `schema:"departmentNumbe,omitempty"`
	R                int    `schema:"r,omitempty"`
	ZoneNumber       int    `schema:"zoneNumber,omitempty"`
	Page             int    `schema:"page,omitempty"`
	PerPage          int    `schema:"per_page,omitempty"`
}

type EmployeeData struct {
	DepartmentNumber int    `json:"departmentNumber"`
	Name             string `json:"name"`
	Pin              string `json:"pin"`
	EmpID            int    `json:"empID"`
	Gender           string `json:"gender"`
	IdCard           string `json:"idCard"`
	Password         string `json:"password"`
	FaceCount        int    `json:"faceCount"`
	FPCount          int    `json:"fpCount"`
	DepartmentName   string `json:"departmentName"`
	ZoneName         string `json:"zoneName"`
}

type EmployeeRecords struct {
	Data       []EmployeeData `json:"data"`
	Page       int            `json:"page"`
	TotalItems int            `json:"totalItems"`
	Code       int            `json:"code"`
	TotalPages int            `json:"totalPages"`
	PageSize   int            `json:"pageSize"`
	NextUrl    string         `json:"next_url"`
	Message    string         `json:"message"`
}

type EmployeeUpsert struct {
	DepartmentNumber int    `json:"departmentNumber"`
	Name             string `json:"name"`
	Pin              string `json:"pin"`
	Gender           string `json:"gender"`
	IdCard           string `json:"idCard"`
	Password         string `json:"password"`
	Privilege        int    `json:"privilege"`
}

type EmployeeDelete struct {
	EmpID []int `json:"empID"`
}
