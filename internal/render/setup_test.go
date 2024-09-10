package render

import (
	"encoding/gob"
	"net/http"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/babtun123/bookings/internal/config"
	"github.com/babtun123/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	// Change this to true when in Production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

}
