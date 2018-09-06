package wdmsapi

import (
	"errors"
	"net/http"

	helper "github.com/zainulbr/go-http-helper"
)

type wdms struct {
	cookie  []*http.Cookie
	baseurl string
}

func New(cookie []*http.Cookie, baseurl string) *wdms {
	return &wdms{cookie, baseurl}
}

func Login(p Account, baseurl string) ([]*http.Cookie, error) {
	r := helper.HTTPReq{Url: baseurl}
	r.Url = r.Url + "/accounts/login/"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}
	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	if data.Code != 200 {
		return nil, errors.New("login failed " + data.Message)
	}

	return resp.Cookie, nil
}

func (w *wdms) Logout(p Account) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/accounts/logout/"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}
	w.cookie = nil
	return &data, nil
}

func (w *wdms) GetDevices(p DeviceParams) (*DeviceRecord, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/devices"
	r.Body = p

	resp, err := r.GET()
	if err != nil {
		return nil, err
	}

	var data DeviceRecord
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) UpsertDevices(p DeviceUpsert) (*responseData, error) {

	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/devices"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil

}

func (w *wdms) DeleteDevices(p DeviceDelete) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/delete/devices"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) GetZone(p ZoneParams) (*ZoneRecord, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/zones"
	r.Body = p

	resp, err := r.GET()
	if err != nil {
		return nil, err
	}

	var data ZoneRecord
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) UpsertZone(p ZoneUpsert) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/zones"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (w *wdms) GetDepartments(p DepartmnetParams) (*DepartmentRecords, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/departments"
	r.Body = p

	resp, err := r.GET()
	if err != nil {
		return nil, err
	}

	var data DepartmentRecords
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) UpsertDepartments(p DepartmentUpsert) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/departments"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) DeleteDepartments(p DepartmentDelete) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/delete/departments"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) GetEmployee(p EmployeeParams) (*EmployeeRecords, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/employees"
	r.Body = p

	resp, err := r.GET()
	if err != nil {
		return nil, err
	}

	var data EmployeeRecords
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) UpsertEmployee(p EmployeeUpsert) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/employees"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) DeleteEmployee(p EmployeeDelete) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/delete/employees"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) GetTransactions(p TransactionsParams) (*TransactionsRecords, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/transactions"
	r.Body = p

	resp, err := r.GET()
	if err != nil {
		return nil, err
	}

	var data TransactionsRecords
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil

}

func (w *wdms) DeleteTransactions(p TransactionsDeletes) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/delete/transactions"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) GetFingerTemplates(p FingerPrintParams) (*FingerPrintRecords, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/fingertemplates"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data FingerPrintRecords
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) GetFaceTemplates(p FaceTemplateParams) (*FaceTemplateRecords, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/facetemplates"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data FaceTemplateRecords
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (w *wdms) SendEmployessToDevice(p SendEmployeesParams) (*responseData, error) {
	r := helper.HTTPReq{Url: w.baseurl, Cookie: w.cookie}
	r.Url = r.Url + "/empstodevs"
	r.Body = p

	resp, err := r.POST()
	if err != nil {
		return nil, err
	}

	var data responseData
	if err := resp.UnMarshall(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
