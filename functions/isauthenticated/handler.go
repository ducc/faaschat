package function

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
	"time"
)

var errTokenNotFound = errors.New("user token not found")

// Handle a serverless request
func Handle(req []byte) (string, error) {
	isAuthenticatedRequest := &IsAuthenticatedRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), isAuthenticatedRequest); err != nil {
		return "", err
	}

	if isAuthenticatedRequest.Token == "" {
		return "", errors.New("invalid token")
	}

	if isAuthenticatedRequest.UserId == "" {
		return "", errors.New("invalid user id")
	}

	ctx := context.Background()

	timeCreated, err := getUserToken(ctx, isAuthenticatedRequest.UserId, isAuthenticatedRequest.Token)
	if err != nil {
		if err == errTokenNotFound {
			// token does not exist, lets set it as expired
			timeCreated = time.Now().AddDate(-1, 0, 0)
		} else {
			return "", err
		}
	}

	var authenticated bool

	// tokens last 30 days
	if timeCreated.AddDate(0, 0, 30).After(time.Now()) {
		authenticated = true
	} else {
		if err := deleteUserToken(ctx, isAuthenticatedRequest.UserId, isAuthenticatedRequest.Token); err != nil {
			return "", err
		}
	}

	isAuthenticatedResponse := &IsAuthenticatedResponse{
		Authenticated: authenticated,
	}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(isAuthenticatedResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func getUserToken(ctx context.Context, userID, token string) (time.Time, error) {
	db, err := getDB()
	if err != nil {
		return time.Time{}, err
	}

	stmt, err := db.PrepareContext(ctx, `select time_created from user_tokens where user_id = $1 and token = $2`)
	if err != nil {
		return time.Time{}, err
	}

	row := stmt.QueryRowContext(ctx, userID, token)

	var timeCreated time.Time
	if err := row.Scan(&timeCreated); err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, errTokenNotFound
		}

		return time.Time{}, err
	}

	return timeCreated, nil
}

func deleteUserToken(ctx context.Context, userID, token string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(ctx, `delete from user_tokens where user_id = $1 and token = $2`)
	if err != nil {
		return err
	}

	if _, err = stmt.ExecContext(ctx, userID, token); err != nil {
		return err
	}

	return nil
}
