package custom_session

import (
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

func (m customSession) Set(key string, val interface{}) (err error) {
	m.cSession.Set(key, val)
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
