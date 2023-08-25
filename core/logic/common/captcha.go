package common

import (
	"csf/core/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

func init() {
	service.RegisterNewCaptcha(NewCaptchaService())
}

type sCaptchaService struct {
	ctx   *gin.Context
	store base64Captcha.Store
}

func NewCaptchaService() *sCaptchaService {
	return &sCaptchaService{
		store: base64Captcha.DefaultMemStore,
	}
}

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// CreateCaptcha 创建验证码
func (s *sCaptchaService) CreateCaptcha(ctx *gin.Context) (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	driver := e.DriverDigit
	cp := base64Captcha.NewCaptcha(driver, s.store)
	return cp.Generate()
}

// Verify 校验验证码
func (s *sCaptchaService) Verify(ctx *gin.Context, id, code string, clear bool) bool {
	return s.store.Verify(id, code, clear)
}
