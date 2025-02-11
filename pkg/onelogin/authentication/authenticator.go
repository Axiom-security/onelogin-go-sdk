package authentication

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"net/http"
	"os"
	"strings"

	olError "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/error"
)

const (
	TkPath     string = "/auth/oauth2/v2/token"
	RevokePath string = "/auth/oauth2/revoke"
)

type Authenticator struct {
	accessToken         string
	subdomain           string
	expiresIn           int
	accountId           string
	credentialsOverride *mod.APICredentials
}

func NewAuthenticator(subdomain string, credentialsOverride *mod.APICredentials) *Authenticator {
	return &Authenticator{subdomain: subdomain, credentialsOverride: credentialsOverride}
}

func (a *Authenticator) GenerateToken() error {
	// Read & Check environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")

	if a.credentialsOverride != nil {
		clientID = a.credentialsOverride.ClientID
		clientSecret = a.credentialsOverride.ClientSecret
	}

	if len(clientID) == 0 {
		return olError.NewAuthenticationError("Missing ONELOGIN_CLIENT_ID Env Variable")
	}

	if len(clientSecret) == 0 {
		return olError.NewAuthenticationError("Missing ONELOGIN_CLIENT_SECRET Env Variable")
	}

	// Construct the authentication URL
	authURL := fmt.Sprintf("https://%s.onelogin.com%s", a.subdomain, TkPath)

	// Create authentication request payload
	data := map[string]string{
		"grant_type": "client_credentials",
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return olError.NewSerializationError("Unable to convert payload to JSON")
	}

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return olError.NewRequestError("Failed to create authentication request")
	}

	// Add authorization header with base64-encoded credentials
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCredentials))
	req.Header.Add("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return olError.NewRequestError("Failed to send authentication request")
	}

	// Parse the authentication response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return olError.NewSerializationError("Failed to read authentication response")
	}

	// Check if authentication failed
	if resp.StatusCode != http.StatusOK {
		return olError.NewAuthenticationError("Authentication failed")
	}

	// Extract access token from the response
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return olError.NewAuthenticationError("Authentication Failed at Endpoint")
	}

	accountId, ok := result["account_id"].(float64)
	if ok {
		a.accountId = fmt.Sprintf("%d", int(accountId))
	}

	expiresIn, ok := result["expires_in"].(float64)
	if ok {
		a.expiresIn = int(expiresIn)
	}

	// Store access token
	a.accessToken = accessToken

	return nil
}

func (a *Authenticator) RevokeToken(token *string) error {
	// Read environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")

	if a.credentialsOverride != nil {
		clientID = a.credentialsOverride.ClientID
		clientSecret = a.credentialsOverride.ClientSecret
	}

	// Check if required environment variables are missing
	if clientID == "" || clientSecret == "" {
		return errors.New("missing client ID, client secret, or subdomain")
	}

	// Construct the revoke URL
	revokeURL := fmt.Sprintf("%s.onelogin.com%s", a.subdomain, RevokePath)

	// Create revoke request payload
	data := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: *token,
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to create revocation request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", revokeURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("failed to create revocation request: %w", err)
	}

	// Add authorization header with base64-encoded credentials
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCredentials))
	req.Header.Add("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to revoke: %w", err)
	}

	// Check if revocation failed
	if resp.StatusCode != http.StatusOK {
		return olError.NewAuthenticationError("Revocation failed")
	}

	// Success condition feedback
	fmt.Println("Revocation successful")

	return nil
}

func (a *Authenticator) GetToken() (string, error) {
	return a.accessToken, nil
}

func (a *Authenticator) GetAccountId() string {
	return a.accountId
}
