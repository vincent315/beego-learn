package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["myproject/controllers:TestController"] = append(beego.GlobalControllerRouter["myproject/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myproject/controllers:TestController"] = append(beego.GlobalControllerRouter["myproject/controllers:TestController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myproject/controllers:TestController"] = append(beego.GlobalControllerRouter["myproject/controllers:TestController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myproject/controllers:TestController"] = append(beego.GlobalControllerRouter["myproject/controllers:TestController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/test`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myproject/controllers:TestController"] = append(beego.GlobalControllerRouter["myproject/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
