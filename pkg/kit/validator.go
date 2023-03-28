package kit

//gin > 1.4.0

//将验证器错误翻译成中文

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func Init() error {
	//注册翻译器
	translator := zh.New()
	uni = ut.New(translator, translator)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	// 注册电话号码验证器
	if err := validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		// 验证电话号码
		phone := fl.Field().Uint()
		if len(strconv.FormatUint(phone, 10)) != 11 {
			return false
		}
		return true // 这里需要根据实际情况返回 true 或 false
	}); err != nil {
		panic(err)
	}
	//注册翻译器
	return zhTranslations.RegisterDefaultTranslations(validate, trans)
}

func UnmarshalJSON(c *gin.Context, param interface{}) bool {
	if err := c.ShouldBindJSON(param); err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal") {
			// 参数转换错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
			return false
		}
	}
	return true
}
func BindJson(c *gin.Context, param interface{}) bool {
	if err := c.ShouldBindJSON(param); err != nil {
		fmt.Println("errsssss", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": Translate(err)})
		return false
	}
	return true
}

// Translate 翻译错误信息
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		// 自定义电话号码错误信息
		if err.Tag() == "phone" {
			result[err.Field()] = append(result[err.Field()], "电话号码格式错误")
			continue
		}
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}
