package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)
var trans ut.Translator
type LoginForm struct {
	User string `json:"user" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

func removeTopStruct(fields map[string]string) map[string]string{
	rsp:=map[string]string{}
	for field,err:=range fields{
		rsp[field[strings.Index(field,".")+1:]]=err
	}
	return rsp
}

func InitTrans(locale string)(err error){
	//修改gin框架中的validator 引擎属性，实现定制
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		//注册一个获取json的tag自定义办法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name:=strings.SplitN(field.Tag.Get("json"),",",2)[0]
			if name=="_"{
				return ""
			}
			return name
		})
		zhT:=zh.New()//
		enT:=en.New()//
		uni:=ut.New(enT,zhT,enT)
		trans,ok=uni.GetTranslator(locale)
		if !ok{
			return fmt.Errorf("uni.GetTranslator(%s)",locale)
		}
		switch  locale {
			case "en":
						en_translations.RegisterDefaultTranslations(v,trans)
			case "zh":
						zh_translations.RegisterDefaultTranslations(v,trans)
			default:
						en_translations.RegisterDefaultTranslations(v,trans)
			}
		return
	}
	return
}

func main(){
	InitTrans("zh")
	router:=gin.Default()
	router.POST("/loginJSON",func(c *gin.Context){
		var loginForm LoginForm
		if err:=c.ShouldBind(&loginForm);err!=nil{
			errs,ok:=err.(validator.ValidationErrors)
			if !ok{
				c.JSON(http.StatusOK,gin.H{
					"msg":err.Error(),
				})
			}

			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest,gin.H{
				"error":removeTopStruct(errs.Translate(trans)),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功",
		})
	})
	router.Run(":8087")
}
