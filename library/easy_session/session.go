package easy_session

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var sessionStore = customSession{}

type customSession struct {
	ctx      *gin.Context
	cSession sessions.Session
}

func NewCustomSession(ctx *gin.Context) customSession {
	sessionStore.ctx = ctx
	sessionStore.cSession = sessions.Default(ctx)
	return sessionStore
}

func (m customSession) Set(key string, data interface{}) (err error) {
	switch data.(type) {
	case string, int, int64, int32, int8, float32:
	default:
		bytes, _ := json.Marshal(data)
		data = string(bytes)
	}
	m.cSession.Set(key, data)
	err = m.cSession.Save()
	if err != nil {
		return
	}
	return
}

func (m customSession) Get(key string) (val interface{}) {
	return m.cSession.Get(key)
}

func (m customSession) Delete(key string) (err error) {
	m.cSession.Delete(key)
	return m.cSession.Save()
}
