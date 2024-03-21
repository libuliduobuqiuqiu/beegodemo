package api

import "github.com/beego/beego/v2/server/web"

func init() {
	web.CtrlGet("/api/device", (*DeviceController).ListDevice)
	web.CtrlGet("/api/device/:name", (*DeviceController).GetDevice)
	web.CtrlPost("/api/device", (*DeviceController).CreateDevice)
	web.CtrlPut("/api/device", (*DeviceController).UpdateDevice)
	web.CtrlDelete("/api/device", (*DeviceController).DeleteDevice)
}
