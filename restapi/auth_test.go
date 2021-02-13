package restapi_test

import (
	"bytes"
	"education/model"
	"education/restapi"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestFirebaseAuth(t *testing.T) {
// 	var jsonStr = []byte(`{"fireToken":"abcd"}`)
// 	req, err := http.NewRequest("GET", "/firebaseAuth", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(restapi.FirebaseAuth)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `{"id":4,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

func TestFirebaseAuth(t *testing.T) {
	var jsonStr = []byte(`{"fireToken":"eyJhbGciOiJSUzI1NiIsImtpZCI6IjYxMDgzMDRiYWRmNDc1MWIyMWUwNDQwNTQyMDZhNDFkOGZmMWNiYTgiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vZWR1Y2F0aW9uLW1vZGVjayIsImF1ZCI6ImVkdWNhdGlvbi1tb2RlY2siLCJhdXRoX3RpbWUiOjE2MTMyMjE1MDUsInVzZXJfaWQiOiIwT1ZvSjE3cE5NT3ZNUUZheG1Rc29yeTJGZWsyIiwic3ViIjoiME9Wb0oxN3BOTU92TVFGYXhtUXNvcnkyRmVrMiIsImlhdCI6MTYxMzIzMTYyNSwiZXhwIjoxNjEzMjM1MjI1LCJlbWFpbCI6ImpvcGFAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImpvcGFAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.I2i1fCNbKKK1CA2jtpkOGGlF6Uma7G2wsiJt35EfdMN4t9D9hbzSRA1ChCVpHdyxEAYOjo9TFW3AjIbjfaIc2DdNM4OTExwBh94EhTidNIQcqhv1ZLdeFD4rKLtqqHfF825eoo0dXJ7MKAhGD-NztwsLMh4156OjcdEc7o6_GeFO1f0jtdDS6O4kqOLxqJOqNFQVMXtyfF6ZtRfQ3Vsu--w398KuJ1-oZKxZn3z7LbGQ4Ow5XR1i_JGgpYkG2Mo5bxjp4D12K7EucnXkh5hpLEvLy0KJsiF8Xz80cVPzKvCiQNzcGkWGPhS-dx244FkMN9uUZmSIgkZ56bPw1fgIDw"}`)
	req, err := http.NewRequest("GET", "/firebaseAuth", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	g := restapi.NewResolver()
	handler := http.HandlerFunc(g.FirebaseAuth)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	user := &model.UserToken{}
	err = json.NewDecoder(rr.Body).Decode(user)
	if err != nil {
		t.Fatal(err)
	}
	expectedUser := &model.UserToken{
		User: model.User{
			DisplayName: "jopa",
		},
	}
	if expectedUser.User.DisplayName != user.User.DisplayName {
		t.Fatalf("Not valid user: %v", user)
	}
}
