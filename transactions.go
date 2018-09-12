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
	CheckTime       string    `json:"checkTime" bson:"-" schema:"-"`
	CheckTimeParsed time.Time `json:"check_time" bson:"check_time" schema:"-"`
	Resrved         string    `json:"reserved" bson:"reserved" schema:"-"`
	Name            string    `json:"name" bson:"name" schema:"-"`
	PIN             string    `json:"pin" bson:"pin" schema:"pin,omitempty"`
	CheckTYpe       string    `json:"checkType" bson:"check_type" schema:"-"`
	ZoneNumber      int       `json:"zoneNumber" bson:"zone_number" schema:"zoneNumber,omitempty"`
	WorkCode        string    `json:"workCode" bson:"work_code" schema:"-"`
	VerifyCode      int       `json:"verifyCode" bson:"verify_code" schema:"-"`
	SN              string    `json:"SN" bson:"sn" schema:"SN,omitempty"`
	ID              int       `json:"id" bson:"id" schema:"id,omitempty"`
	TransactionPic  string    `json:"transactionPic" bson:"transaction_pic" schema:"-"`
	ZoneName        string    `json:"zoneName" bson:"zone_name"  schema:"-"`
	SyncFlag        *bool     `json:"-" bson:"-"  schema:"synced,omitempty"`
	Synced          bool      `json:"synced" bson:"synced"  schema:"-"`
	Subdomain       string    `json:"subdomain" bson:"subdomain"  schema:"-"`
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
