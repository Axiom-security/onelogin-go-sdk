package mocks

import (
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/api"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/authentication"
	"net/http"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

type MockAuthenticator struct {
	GetTokenFunc         func() (string, error)
	NewAuthenticatorFunc func() *authentication.Authenticator
}

func (m *MockAuthenticator) GetToken() (string, error) {
	return m.GetTokenFunc()
}

func (m *MockAuthenticator) NewAuthenticator() *authentication.Authenticator {
	return &authentication.Authenticator{}
}

func CreateMockClient() *api.Client {
	mockClient := &MockHttpClient{}
	mockAuth := &MockAuthenticator{}

	mockAuth.GetTokenFunc = func() (string, error) {
		return "mockToken", nil
	}

	auth := authentication.NewAuthenticator("test", nil)
	client := &api.Client{
		HttpClient: mockClient,
		Auth:       auth,
		OLdomain:   "https://api.onelogin.com",
	}

	return client
}
