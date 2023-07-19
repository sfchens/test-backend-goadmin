package common_service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

type sComCaptchaService struct {
	ctx   *gin.Context
	store base64Captcha.Store
}

func NewComCaptchaService(ctx *gin.Context) *sComCaptchaService {
	return &sComCaptchaService{
		ctx:   ctx,
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

func (s sComCaptchaService) CreateCaptcha() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	driver := e.DriverDigit
	cp := base64Captcha.NewCaptcha(driver, s.store)
	return cp.Generate()
}

// Verify 校验验证码
func (s sComCaptchaService) Verify(id, code string, clear bool) bool {
	return s.store.Verify(id, code, clear)
}
