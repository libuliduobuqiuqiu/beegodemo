package api

import (
	"fmt"

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
	web.InsertFilterChain("/*", logFilterChain)
}

func logFilterChain(next web.FilterFunc) web.FilterFunc {
	return func(ctx *context.Context) {

		next(ctx)

		requestLog := fmt.Sprintf(" %3d | %13v | %15s | %-7v %s",
			ctx.Output.Status, ctx.Output.Context.ResponseWriter.Elapsed, ctx.Request.Host, ctx.Request.Method, ctx.Request.URL)
		logs.Info(requestLog)
	}
}
