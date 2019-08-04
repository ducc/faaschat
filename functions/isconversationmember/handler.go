package function

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	isMemberRequest := &IsConversationMemberRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), isMemberRequest); err != nil {
		return "", err
	}

	if isMemberRequest.UserId == "" {
		return "", errors.New("user id is empty")
	}

	if isMemberRequest.Token == "" {
		return "", errors.New("token is empty")
	}

	if isMemberRequest.ConversationId == "" {
		return "", errors.New("conversation id is empty")
	}

	if authenticated, err := isAuthenticated(isMemberRequest.UserId, isMemberRequest.Token); err != nil {
		return "", err
	} else if !authenticated {
		return "", errors.New("not authenticated")
	}

	ctx := context.Background()

	isMember, err := getIsMember(ctx, isMemberRequest.ConversationId, isMemberRequest.UserId)
	if err != nil {
		return "", err
	}

	isMemberResponse := &IsConversationMemberResponse{
		Member: isMember,
	}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(isMemberResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func getIsMember(ctx context.Context, conversationID, userID string) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	stmt, err := db.PrepareContext(ctx, `select user_id from conversation_users where conversation_id = $1 and user_id = $2 limit 1`)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRowContext(ctx, conversationID, userID)

	var unusedUserID string
	if err := row.Scan(&unusedUserID); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
