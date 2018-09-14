package wdmsapi

import (
	"time"
)

// type TransactionsParams struct {
// 	PIN        interface{} `schema:"pin,omitempty"`
// 	ZoneNumber interface{} `schema:"zoneNumber,omitempty"`
// 	StartTIme  string      `schema:"startTime"`
// 	EndTime    string      `schema:"endTime"`
// 	SN         interface{} `schema:"sn,omitempty"`
// 	Page       interface{} `schema:"page,omitempty"`
// 	PerPage    interface{} `schema:"per_page,omitempty"`
// }

type TransactionsParams struct {
	PIN        string `schema:"pin,omitempty"`
	ZoneNumber int    `schema:"zoneNumber,omitempty"`
	StartTIme  string `schema:"startTime"`
	EndTime    string `schema:"endTime"`
	SN         string `schema:"sn,omitempty"`
	Page       int    `schema:"page,omitempty"`
	PerPage    int    `schema:"per_page,omitempty"`
}

type TransactionsData struct {
	CheckTime       string    `json:"checkTime,omitempty" bson:"-,omitempty" schema:"-"`
	CheckTimeParsed time.Time `json:"check_time,omitempty" bson:"check_time,omitempty" schema:"-"`
	Resrved         string    `json:"reserved,omitempty" bson:"reserved,omitempty" schema:"-"`
	Name            string    `json:"name,omitempty" bson:"name,omitempty" schema:"-"`
	PIN             string    `json:"pin,omitempty" bson:"pin,omitempty" schema:"pin,omitempty"`
	CheckTYpe       string    `json:"checkType,omitempty" bson:"check_type,omitempty" schema:"-"`
	ZoneNumber      int       `json:"zoneNumber,omitempty" bson:"zone_number,omitempty" schema:"zoneNumber,omitempty"`
	WorkCode        string    `json:"workCode,omitempty" bson:"work_code,omitempty" schema:"-"`
	VerifyCode      int       `json:"verifyCode,omitempty" bson:"verify_code,omitempty" schema:"-"`
	SN              string    `json:"SN,omitempty" bson:"sn,omitempty" schema:"SN,omitempty"`
	ID              int       `json:"id,omitempty" bson:"id,omitempty" schema:"id,omitempty"`
	TransactionPic  string    `json:"transactionPic,omitempty" bson:"transaction_pic,omitempty" schema:"-"`
	ZoneName        string    `json:"zoneName,omitempty" bson:"zone_name,omitempty" schema:"-"`
	SyncFlag        *bool     `json:"-" bson:"-" schema:"synced,omitempty"`
	Synced          bool      `json:"synced,omitempty" bson:"synced,omitempty" schema:"-"`
	Subdomain       string    `json:"subdomain,omitempty" bson:"subdomain,omitempty" schema:"-"`
}

type TransactionsRecords struct {
	Data       []TransactionsData `json:"data"`
	Page       int                `json:"page"`
	TotalItems int                `json:"totalItems"`
	Code       int                `json:"code"`
	TotalPages int                `json:"totalPages"`
	PageSize   int                `json:"pageSize"`
	NextUrl    string             `json:"next_url"`
	Message    string             `json:"message"`
}

type TransactionsDelete struct {
	StartTIme time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type TransactionsDeletes struct {
	Period []TransactionsDelete `json:"Period"`
}
