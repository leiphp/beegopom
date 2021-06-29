package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"lxtkj/hellobeego/consts"
	"lxtkj/hellobeego/models"
	"net/http"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName string
	IsLogin   bool
	Loginuser interface{}
}

func (c *BaseController) Prepare()  {
	//获取session
	//loginuser := c.GetSession("user")
	//fmt.Println("loginuser---->", loginuser)

	//赋值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	//TODO 保存用户数据
	fmt.Println("beego:perpare"+c.controllerName+","+c.actionName)
	//检查sso令牌access_token
	//tokenMap := c.checkToken()
	//fmt.Println("tokenMap",tokenMap)

	user := c.auth()//验证登录
	//获取sessionc
	if user != (models.UserModel{}) {//通过判断struct是否初始化，判断user是否为空
		c.IsLogin = true
		c.Loginuser = user
	}else {
		c.IsLogin = false
		c.Loginuser = models.UserModel{}
	}

	//c.Data["Menu"] = models.MenuStruct()
	c.Data["Menu"] = models.MenuTreeStruct(user)
	c.Data["IsLogin"] = c.IsLogin
	c.Data["LoginUser"] = c.Loginuser
}

//设置模板
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(c.controllerName[0:len(c.controllerName) -10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" +actionName+".html"
	}
	_,found := c.Data["Footer"]
	if !found{
		c.Data["Footer"] = "menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}

func (c *BaseController) jsonResult(code consts.JsonResultCode, msg string, obj interface{}){
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) listJsonResult(code consts.JsonResultCode, msg string, count int64, obj interface{}) {
	r := &models.ListJsonResult{code, msg, count, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) auth() models.UserModel {
	user := c.GetSession("user")
	fmt.Println("user",user)
	if user == nil {
		c.Redirect("/login", 302)
		//c.Redirect("http://localhost:9096/authorize?client_id=test_client_3&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:8080/login", 302)//authorization_code才需要到认证中心授权
		c.StopRun()
		return models.UserModel{}
	} else {
		fmt.Println("get user:",user.(models.UserModel))
		return user.(models.UserModel)
	}
}

//效验sso分发的access_token有效性
func (c *BaseController) checkToken() interface{} {
	oauth_token := c.GetSession("oauth_token")
	fmt.Println("oauth_token",oauth_token)
	if oauth_token != nil {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/test", nil)
		//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Type","application/x-www-form-urlencoded")
		//req.BasicAuth()
		req.Header.Set("Authorization","Bearer "+oauth_token.(string))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer resp.Body.Close()
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println("resp str:",string(bs))//invalid access token
		m := make(map[string]interface{})
		err = json.Unmarshal(bs, &m)
		if err != nil {
			fmt.Println("Umarshal failed:", err)//Umarshal failed: invalid character 'i' looking for beginning of value
			c.Redirect("/login", 302)
			//c.Redirect("http://localhost:9096/authorize?client_id=test_client_3&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:8080/login", 302)//authorization_code才需要到认证中心授权
			c.StopRun()
		}
		fmt.Println("map:", m)//map: map[client_id:test_client_3 domain:http://localhost:8080 expires_in:7199 scope:all user_id:admin]
		return m
	} else {
		c.Redirect("/login", 302)
		//c.Redirect("http://localhost:9096/authorize?client_id=test_client_3&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:8080/login", 302)//authorization_code才需要到认证中心授权
		c.StopRun()
		return ""
	}
}