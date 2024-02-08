package integration

import (
	"bytes"
	"log"
	"project.com/restful-api/utilities"

	"net/http"

	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"project.com/restful-api/controllers"
)

type getresponse struct {
	body string
	code int
}

type test struct {
	name        string
	environment string
	method      string
	rurl        string
	body        string
	expres      string
	expcode     int
}

var tests = []test{
	{name: "GET request", environment: "TESTING", method: http.MethodGet, rurl: "/healthz", expres: "", expcode: http.StatusOK},
}

func TestInitController(t *testing.T) {
	utilities.LoadLogger()
	utilities.InitializeStatsd()

	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Could not load environment file at given location " + err.Error())
	}

	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			router := initServer()
			body := bytes.NewBuffer([]byte(value.body))
			var request *http.Request
			if value.body == "" {
				request, _ = http.NewRequest(value.method, value.rurl, nil)
			} else {
				request, _ = http.NewRequest(value.method, value.rurl, body)
			}

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			got := getresponse{recorder.Body.String(), recorder.Code}
			want := getresponse{value.expres, value.expcode}

			assertCorrectResponse(t, got, want)
		})
	}
}

func assertCorrectResponse(t testing.TB, got getresponse, want getresponse) {
	t.Helper()
	if got.body != want.body || got.code != want.code {
		t.Errorf("got %v want %v", got, want)
	}
}

func initServer() *gin.Engine {
	router := gin.Default()
	controllers.InitControllers(router)

	return router
}
