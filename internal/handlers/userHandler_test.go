package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/lauro-ss/api_with_goe/internal/data"
)

func Setup() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		var u data.User
		json.NewDecoder(r.Body).Decode(&u)
		w.Write([]byte(fmt.Sprintf("%v %v", u.Name, u.Password)))
	})
	mux.HandleFunc("GET /user/:id", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 1)
		w.Write([]byte("UserGet " + r.URL.Query().Get("id")))
	})
	mux.HandleFunc("PUT /user/:id", func(w http.ResponseWriter, r *http.Request) {
		var u data.User
		json.NewDecoder(r.Body).Decode(&u)
		w.Write([]byte(fmt.Sprintf("%v %v", u.Name, u.Password)))
	})
	//mux.HandleFunc("DELETE /user/:id", handlers.UserDelete(userRepository))
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &s
}

func TestGetUserById(t *testing.T) {
	server := Setup()

	tests := []struct {
		name     string
		testCase func(t *testing.T) *http.Request
		exCode   int
		exBody   string
	}{
		{
			name: "Test Case for Getting User",
			testCase: func(t *testing.T) *http.Request {
				request, err := http.NewRequest(http.MethodGet, "/user/:id", nil)
				if err != nil {
					t.Fatal(err)
				}

				q := request.URL.Query()
				q.Add("id", "1")
				request.URL.RawQuery = q.Encode()
				return request
			},
			exCode: http.StatusOK,
			exBody: "UserGet 1",
		},
		{
			name: "Test Case for Creating a User",
			testCase: func(t *testing.T) *http.Request {
				b, err := json.Marshal(data.User{Name: "Lauro", Password: "123456"})
				if err != nil {
					t.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(b))
				if err != nil {
					t.Fatal(err)
				}

				return request
			},
			exCode: http.StatusOK,
			exBody: "Lauro 123456",
		},

		{
			name: "Test Case for Updating a User",
			testCase: func(t *testing.T) *http.Request {
				b, err := json.Marshal(data.User{Name: "Lauro Santana", Password: "123456"})
				if err != nil {
					t.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPut, "/user/:id", bytes.NewReader(b))
				if err != nil {
					t.Fatal(err)
				}

				return request
			},
			exCode: http.StatusOK,
			exBody: "Lauro Santana 123456",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRecorder()
			server.Handler.ServeHTTP(r, test.testCase(t))

			if r.Result().StatusCode != test.exCode {
				t.Errorf("Expected %v, got %v", test.exCode, r.Result().StatusCode)
			}

			if r.Body.String() != test.exBody {
				t.Errorf("Expected %v, got %v", test.exBody, r.Body.String())
			}
		})
	}
}
