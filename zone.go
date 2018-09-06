package wdmsapi

type ZoneParams struct {
	Page    int `schema:"page,omitempty"`
	PerPage int `schema:"per_page,omitempty"`
}

type ZoneData struct {
	ZoneNumber int    `json:"zoneNumber"`
	ZoneName   string `json:"zoneName"`
}

type ZoneRecord struct {
	Data       []ZoneData `json:"data"`
	Page       int        `json:"page"`
	TotalItems int        `json:"totalItems"`
	Code       int        `json:"code"`
	TotalPages int        `json:"totalPages"`
	PageSize   int        `json:"pageSize"`
	NextUrl    string     `json:"next_url"`
	Message    string     `json:"message"`
}

type ZoneUpsert struct {
	Data []ZoneData `json:"Data"`
}
