package function

import (
	"bytes"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	isAuthenticatedRequest := &IsAuthenticatedRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), isAuthenticatedRequest); err != nil {
		return "", err
	}

	if isAuthenticatedRequest.Token == "" {
		return "", errors.New("invalid token")
	}

	// todo check if token is in database
	// if token has expired, delete it and return false
	// return true if valid token

	isAuthenticatedResponse := &IsAuthenticatedResponse{
		Authenticated: false, // todo
	}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(isAuthenticatedResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}
