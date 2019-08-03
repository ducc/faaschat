package function

import (
	"bytes"
	"context"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	registerRequest := &RegisterRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), registerRequest); err != nil {
		return "", err
	}

	if registerRequest.UserName == "" {
		return "", errors.New("user name is empty")
	}

	if err := createUser(context.Background(), registerRequest.UserName); err != nil {
		return "", err
	}

	registerResponse := &RegisterResponse{}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(registerResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func createUser(ctx context.Context, name string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(ctx, `insert into users (user_name) values ($1)`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
