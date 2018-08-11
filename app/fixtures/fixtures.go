package fixtures

import (
	"github.com/urbn/ordernumbergenerator/app"
	"net/http"
	"net/http/httptest"
)

var ApplyResultAN = app.MongoDocument{
	"AN",
	"an",
	"US-NV",
	1,
}

var MockOrderDaoError = app.Error{
	400,
	http.StatusBadRequest,
	"Unable to connect to MongoDB",
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
