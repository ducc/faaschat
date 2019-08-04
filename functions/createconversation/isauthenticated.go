package function

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getIsAuthenticatedEndpoint() string {
	return os.Getenv("ENDPOINT_ISAUTHENTICATED")
}

func isAuthenticated(userID, token string) (bool, error) {
	body := bytes.NewBuffer([]byte{})
	if err := (&jsonpb.Marshaler{}).Marshal(body, &IsAuthenticatedRequest{
		Token:  token,
		UserId: userID,
	}); err != nil {
		return false, fmt.Errorf("marshalling request: %s", err.Error())
	}

	res, err := http.Post(getIsAuthenticatedEndpoint(), "application/json", body)
	if err != nil {
		return false, fmt.Errorf("posting to url: %s", err.Error())
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("error closing response body: %s\n", err.Error())
		}
	}()

	isAuthenticatedResponse := &IsAuthenticatedResponse{}
	if err := jsonpb.Unmarshal(res.Body, isAuthenticatedResponse); err != nil {
		return false, fmt.Errorf("unmarshalling response: %s", err.Error())
	}

	return isAuthenticatedResponse.Authenticated, nil
}
