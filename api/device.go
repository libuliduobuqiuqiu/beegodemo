package api

import (
	"beegodemo/types"
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
)

type Resp struct {
	Code int           `json:"code"`
	Info string        `json:"info,omitempty"`
	Err  string        `json:"err,omitempty"`
	Data []interface{} `json:"data,omitempty"`
}

type DeviceController struct {
	web.Controller
}

func (d *DeviceController) ListDevice() {
	d.Ctx.WriteString("success")
}

func (d *DeviceController) GetDevice() {
	name := d.Ctx.Input.Param(":name")
	device := types.Device{}
	device.Name = name
	d.Ctx.JSONResp(device)
}

func (d *DeviceController) CreateDevice() {

	tmpDevice := &types.Device{}
	err := d.BindJSON(tmpDevice)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("%v\n", tmpDevice)
	d.Ctx.JSONResp(Resp{Code: 200, Info: "success", Data: []interface{}{tmpDevice}})
}

func (d *DeviceController) UpdateDevice() {
	d.Ctx.JSONResp(Resp{Code: 200, Info: "success"})
}

func (d *DeviceController) DeleteDevice() {
	d.Ctx.JSONResp(Resp{Code: 200, Info: "success"})
}
