package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_Init(t *testing.T) {

	c := &Session{
		CookieLifeTime: "100",
		CookiePersist:  "true",
		CookieName:     "celeritas",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager
	ses := c.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {

		fmt.Println("For Loop:", rv.Kind(), rv.Type(), rv)
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("Invalid kind or type ")
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong lind returned testing cookie session. Ex[ected,", reflect.ValueOf(sm).Kind(), " got:", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("wrong lind returned testing cookie session. Ex[ected,", reflect.ValueOf(sm).Type(), " got:", sessType)
	}
}
