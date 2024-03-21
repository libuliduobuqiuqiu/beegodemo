package api

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {

	web.CtrlGet("/api/device", (*DeviceController).ListDevice)
	web.CtrlGet("/api/device/:name", (*DeviceController).GetDevice)
	web.CtrlPost("/api/device", (*DeviceController).CreateDevice)
	web.CtrlPut("/api/device", (*DeviceController).UpdateDevice)
	web.CtrlDelete("/api/device", (*DeviceController).DeleteDevice)
	web.InsertFilter("/*", web.BeforeRouter, logRequest)
}

func logRequest(ctx *context.Context) {
	logs.Info("request incoming: ", ctx.Request.URL)
}
