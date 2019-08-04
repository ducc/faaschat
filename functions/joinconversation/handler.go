package function

import (
	"bytes"
	"context"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	joinRequest := &JoinConversationRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), joinRequest); err != nil {
		return "", err
	}

	if joinRequest.UserId == "" {
		return "", errors.New("user id is empty")
	}

	if joinRequest.Token == "" {
		return "", errors.New("token is empty")
	}

	if joinRequest.ConversationId == "" {
		return "", errors.New("conversation id is empty")
	}

	if authenticated, err := isAuthenticated(joinRequest.UserId, joinRequest.Token); err != nil {
		return "", err
	} else if !authenticated {
		return "", errors.New("not authenticated")
	}

	ctx := context.Background()

	if err := joinConversation(ctx, joinRequest.UserId, joinRequest.ConversationId); err != nil {
		return "", err
	}

	joinResponse := &JoinConversationResponse{}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(joinResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func joinConversation(ctx context.Context, userID, conversationID string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(ctx, `insert into conversation_users (conversation_id, user_id) values ($1, $2)`)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, conversationID, userID); err != nil {
		return err
	}

	return nil
}
