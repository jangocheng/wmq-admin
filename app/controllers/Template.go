package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type TemplateController struct {
	beego.Controller
	controllerName string
	methodName string
}

//执行前
func (this *TemplateController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.methodName = actionName
}

//执行后
func (this *TemplateController) finish() {
	
}

//渲染模板
func (this *TemplateController) display(tpl ...string) {
	this.Layout = "layout/default.html";
	this.Data["navName"] = this.controllerName;
	this.TplName = tpl[0];
}

//成功输出json
func (this *TemplateController) jsonSuccess(message string, redirect string, data interface{}) {
	this.jsonResult(1, message, redirect, data);
}

//错误输出json
func (this *TemplateController) jsonError(message string, redirect string, data interface{}) {
	this.jsonResult(0, message, redirect, data);
}

//输出json
func (this *TemplateController) jsonResult(code int, message string, redirect string, data interface{}) {
	body := make(map[string]interface{});
	body["code"] = code;
	body["message"] = message;
	body["redirect"] = redirect;
	body["data"] = data;

	this.Data["json"] = body;
	this.ServeJSON();
	this.StopRun();
}

//302跳转
func (this *TemplateController) redirect(tpl ...string) {

}

//是否是 post 请求
func (this *TemplateController) isPost() bool {
	return this.Ctx.Request.Method == "POST";
}