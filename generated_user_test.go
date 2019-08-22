// Code generated by go-uaa/generator; DO NOT EDIT.

package uaa_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	uaa "github.com/cloudfoundry-community/go-uaa"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testUser(t *testing.T, when spec.G, it spec.S) {
	var (
		s       *httptest.Server
		handler http.Handler
		called  int
		a       *uaa.API
	)

	it.Before(func() {
		RegisterTestingT(t)
		called = 0
		s = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			called = called + 1
			Expect(handler).NotTo(BeNil())
			handler.ServeHTTP(w, req)
		}))
		c := &http.Client{Transport: http.DefaultTransport}
		u, _ := url.Parse(s.URL)
		a = &uaa.API{
			TargetURL:             u,
			AuthenticatedClient:   c,
			UnauthenticatedClient: c,
		}
	})

	it.After(func() {
		if s != nil {
			s.Close()
		}
	})

	when("GetUser()", func() {
		when("the user is returned from the server", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					Expect(req.Header.Get("Accept")).To(Equal("application/json"))
					Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(userResponse))
				})
			})
			it("gets the user from the UAA by ID", func() {
				user, err := a.GetUser("00000000-0000-0000-0000-000000000001")
				Expect(err).NotTo(HaveOccurred())
				Expect(user.ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			})
		})

		when("the server errors", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					Expect(req.Header.Get("Accept")).To(Equal("application/json"))
					Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
					w.WriteHeader(http.StatusInternalServerError)
				})
			})

			it("returns helpful error", func() {
				user, err := a.GetUser("00000000-0000-0000-0000-000000000001")
				Expect(err).To(HaveOccurred())
				Expect(user).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("An error occurred while calling"))
			})
		})

		when("the server returns unparsable users", func() {
			it.Before(func() {
				handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					Expect(req.Header.Get("Accept")).To(Equal("application/json"))
					Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{unparsable-json-response}"))
				})
			})

			it("returns helpful error", func() {
				user, err := a.GetUser("00000000-0000-0000-0000-000000000001")
				Expect(err).To(HaveOccurred())
				Expect(user).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("An unknown error occurred while parsing response from"))
				Expect(err.Error()).To(ContainSubstring("Response was {unparsable-json-response}"))
			})
		})
	})

	when("CreateUser()", func() {
		it("performs a POST with the user data and returns the created user", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.Header.Get("Content-Type")).To(Equal("application/json"))
				Expect(req.Method).To(Equal(http.MethodPost))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				defer req.Body.Close()
				body, _ := ioutil.ReadAll(req.Body)
				Expect(body).To(MatchJSON(testUserJSON))
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(userResponse))
			})

			created, err := a.CreateUser(testUserValue)
			Expect(called).To(Equal(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(created).NotTo(BeNil())
		})

		it("returns error when response cannot be parsed", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal(http.MethodPost))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{unparseable}"))
			})
			created, err := a.CreateUser(testUserValue)
			Expect(err).To(HaveOccurred())
			Expect(created).To(BeNil())
		})

		it("returns error when response is not 200 OK", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal(http.MethodPost))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				w.WriteHeader(http.StatusBadRequest)
			})
			created, err := a.CreateUser(testUserValue)
			Expect(err).To(HaveOccurred())
			Expect(created).To(BeNil())
		})
	})

	when("UpdateUser()", func() {
		it("performs a PUT with the user data and returns the updated user", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.Header.Get("Content-Type")).To(Equal("application/json"))
				Expect(req.Method).To(Equal(http.MethodPut))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
				defer req.Body.Close()
				body, _ := ioutil.ReadAll(req.Body)
				Expect(body).To(MatchJSON(testUserJSON))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userResponse))
			})

			updated, err := a.UpdateUser(testUserValue)
			Expect(called).To(Equal(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(updated).NotTo(BeNil())
		})

		it("returns error when response cannot be parsed", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal(http.MethodPut))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{unparseable}"))
			})
			updated, err := a.UpdateUser(testUserValue)
			Expect(err).To(HaveOccurred())
			Expect(updated).To(BeNil())
		})

		it("returns error when response is not 200 OK", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal(http.MethodPut))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
				w.WriteHeader(http.StatusBadRequest)
			})
			updated, err := a.UpdateUser(testUserValue)
			Expect(err).To(HaveOccurred())
			Expect(updated).To(BeNil())
		})
	})

	when("DeleteUser()", func() {
		it("errors when the userID is empty", func() {
			deleted, err := a.DeleteUser("")
			Expect(called).To(Equal(0))
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeNil())
		})

		it("performs a DELETE for the user", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.Method).To(Equal(http.MethodDelete))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userResponse))
			})

			deleted, err := a.DeleteUser("00000000-0000-0000-0000-000000000001")
			Expect(called).To(Equal(1))
			Expect(err).NotTo(HaveOccurred())
			Expect(deleted).NotTo(BeNil())
		})

		it("returns error when response cannot be parsed", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal(http.MethodDelete))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{unparseable}"))
			})
			deleted, err := a.DeleteUser("00000000-0000-0000-0000-000000000001")
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeNil())
		})

		it("returns error when response is not 200 OK", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Method).To(Equal(http.MethodDelete))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint + "/00000000-0000-0000-0000-000000000001"))
				w.WriteHeader(http.StatusBadRequest)
			})
			deleted, err := a.DeleteUser("00000000-0000-0000-0000-000000000001")
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeNil())
		})
	})

	when("ListAllUsers()", func() {
		it("can return multiple pages", func() {
			page1 := MultiPaginatedResponse(1, 1, 2, uaa.User{ID: "test-user-1"})
			page2 := MultiPaginatedResponse(2, 1, 2, uaa.User{ID: "test-user-2"})
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				w.WriteHeader(http.StatusOK)
				if called == 1 {
					Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
					Expect(req.URL.Query().Get("count")).To(Equal("100"))
					w.Write([]byte(page1))
				} else {
					Expect(req.URL.Query().Get("startIndex")).To(Equal("2"))
					Expect(req.URL.Query().Get("count")).To(Equal("1"))
					w.Write([]byte(page2))
				}
			})

			users, err := a.ListAllUsers("", "", "", "")
			Expect(err).NotTo(HaveOccurred())
			Expect(users[0].ID).To(Equal("test-user-1"))
			Expect(users[1].ID).To(Equal("test-user-2"))
			Expect(called).To(Equal(2))
		})

		it("returns an error when the endpoint doesn't respond", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				w.WriteHeader(http.StatusInternalServerError)
			})

			users, err := a.ListAllUsers("", "", "", "")
			Expect(err).To(HaveOccurred())
			Expect(users).To(BeNil())
			Expect(called).To(Equal(1))
		})
	})

	when("ListUsers()", func() {
		it("can accept a filter query to limit results", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal("id eq \"00000000-0000-0000-0000-000000000001\""))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userListResponse))
			})
			userList, _, err := a.ListUsers("id eq \"00000000-0000-0000-0000-000000000001\"", "", "", "", 1, 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(userList[0].ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			Expect(userList[1].ID).To(Equal("00000000-0000-0000-0000-000000000002"))
		})

		it("does not include the filter param if no filter exists", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal(""))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userListResponse))
			})
			userList, _, err := a.ListUsers("", "", "", "", 1, 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(userList[0].ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			Expect(userList[1].ID).To(Equal("00000000-0000-0000-0000-000000000002"))
		})

		it("can accept an attributes list", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal("id eq \"00000000-0000-0000-0000-000000000001\""))
				Expect(req.URL.Query().Get("attributes")).To(Equal("testField"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userListResponse))
			})
			userList, _, err := a.ListUsers("id eq \"00000000-0000-0000-0000-000000000001\"", "", "testField", "", 1, 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(userList[0].ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			Expect(userList[1].ID).To(Equal("00000000-0000-0000-0000-000000000002"))
		})

		it("can accept sortBy", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal(""))
				Expect(req.URL.Query().Get("sortBy")).To(Equal("testField"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userListResponse))
			})
			userList, _, err := a.ListUsers("", "testField", "", "", 1, 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(userList[0].ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			Expect(userList[1].ID).To(Equal("00000000-0000-0000-0000-000000000002"))
		})

		it("can accept sort order ascending/descending", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal(""))
				Expect(req.URL.Query().Get("sortBy")).To(Equal(""))
				Expect(req.URL.Query().Get("sortOrder")).To(Equal("ascending"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userListResponse))
			})
			userList, _, err := a.ListUsers("", "", "", uaa.SortAscending, 1, 100)
			Expect(err).NotTo(HaveOccurred())
			Expect(userList[0].ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			Expect(userList[1].ID).To(Equal("00000000-0000-0000-0000-000000000002"))
		})

		it("uses a startIndex of 1 if 0 is supplied", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(userListResponse))
			})
			userList, _, err := a.ListUsers("", "", "", "", 0, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(userList[0].ID).To(Equal("00000000-0000-0000-0000-000000000001"))
			Expect(userList[1].ID).To(Equal("00000000-0000-0000-0000-000000000002"))
		})

		it("returns an error when the endpoint doesn't respond", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal("id eq \"00000000-0000-0000-0000-000000000001\""))
				w.WriteHeader(http.StatusInternalServerError)
			})

			userList, _, err := a.ListUsers("id eq \"00000000-0000-0000-0000-000000000001\"", "", "", "", 1, 100)
			Expect(err).To(HaveOccurred())
			Expect(userList).To(BeNil())
		})

		it("returns an error when response is unparseable", func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				Expect(req.Header.Get("Accept")).To(Equal("application/json"))
				Expect(req.URL.Path).To(Equal(uaa.UsersEndpoint))
				Expect(req.URL.Query().Get("count")).To(Equal("100"))
				Expect(req.URL.Query().Get("startIndex")).To(Equal("1"))
				Expect(req.URL.Query().Get("filter")).To(Equal("id eq \"00000000-0000-0000-0000-000000000001\""))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{unparsable}"))
			})
			userList, _, err := a.ListUsers("id eq \"00000000-0000-0000-0000-000000000001\"", "", "", "", 1, 100)
			Expect(err).To(HaveOccurred())
			Expect(userList).To(BeNil())
		})
	})
}
