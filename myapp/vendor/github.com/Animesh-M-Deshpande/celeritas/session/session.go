package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifeTime string
	CookieSecure   string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookiePersist  string
}

func (c *Session) InitSession() *scs.SessionManager {
	var secure bool

	// how long does session last

	minutes, err := strconv.Atoi(c.CookieLifeTime)

	if err != nil {
		minutes = 60
	}

	//shoud cookies persist
	/*
		if strings.ToLower(c.CookieSecure) == "true" {
			persist = true
		}
	*/
	// must cookies be secure

	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// whcih session store
	switch strings.ToLower(c.SessionType) {

	case "redis":

	case "mysql", "mariadb":

	default:
		//cookie
	}

	return session

}
