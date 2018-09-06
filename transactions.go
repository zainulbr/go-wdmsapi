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
	CheckTime       string    `json:"checkTime" bson:"-"`
	CheckTimeParsed time.Time `json:"-" bson:"check_time"`
	Resrved         string    `json:"reserved" bson:"reserved"`
	Name            string    `json:"name" bson:"name"`
	PIN             string    `json:"pin" bson:"pin"`
	CheckTYpe       string    `json:"checkType" bson:"check_type"`
	ZoneNumber      int       `json:"zoneNumber" bson:"zone_number"`
	WorkCode        string    `json:"workCode" bson:"work_code"`
	VerifyCode      int       `json:"verifyCode" bson:"verify_code"`
	SN              string    `json:"SN" bson:"sn"`
	ID              int       `json:"id" bson:"id"`
	TransactionPic  string    `json:"transactionPic" bson:"transaction_pic"`
	ZoneName        string    `json:"zoneName" bson:"zone_name" `
	Synced          bool      `json:"-" bson:"synced" `
	Subdomain       string    `json:"-" bson:"subdomain" `
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
