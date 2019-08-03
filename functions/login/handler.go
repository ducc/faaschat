package function

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/google/uuid"
)

var errUserNotFound = errors.New("user not found")

// Handle a serverless request
func Handle(req []byte) (string, error) {
	loginRequest := &LoginRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), loginRequest); err != nil {
		return "", err
	}

	if loginRequest.UserName == "" {
		return "", errors.New("user name is empty")
	}

	ctx := context.Background()
	userID, err := getUserID(ctx, loginRequest.UserName)
	if err != nil {
		return "", err
	}

	token := uuid.New().String()
	if err := setUserToken(ctx, userID, token); err != nil {
		return "", nil
	}

	loginResponse := &LoginResponse{
		Token: token,
	}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(loginResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func getUserID(ctx context.Context, name string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}

	stmt, err := db.PrepareContext(ctx, `select user_id from users where user_name = $1 limit 1`)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRow(name)

	var userID string
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errUserNotFound
		}

		return "", err
	}

	return userID, nil
}

func setUserToken(ctx context.Context, userID, token string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(ctx, `upsert into user_tokens values ($1, $2)`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, userID, token)
	if err != nil {
		return err
	}

	return nil
}
