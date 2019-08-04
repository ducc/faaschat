package function

import (
	"bytes"
	"context"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	sendRequest := &SendMessageRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), sendRequest); err != nil {
		return "", err
	}

	if sendRequest.UserId == "" {
		return "", errors.New("user id is empty")
	}

	if sendRequest.Token == "" {
		return "", errors.New("token is empty")
	}

	if sendRequest.ConversationId == "" {
		return "", errors.New("conversation id is empty")
	}

	if sendRequest.Content == "" {
		return "", errors.New("content is empty")
	}

	if authenticated, err := isAuthenticated(sendRequest.UserId, sendRequest.Token); err != nil {
		return "", err
	} else if !authenticated {
		return "", errors.New("not authenticated")
	}

	if member, err := isConversationMember(sendRequest.UserId, sendRequest.Token, sendRequest.ConversationId); err != nil {
		return "", err
	} else if !member {
		return "", errors.New("not a member of the conversation")
	}

	ctx := context.Background()

	if err := sendMessage(ctx, sendRequest.UserId, sendRequest.ConversationId, sendRequest.Content); err != nil {
		return "", err
	}

	sendResponse := &SendMessageResponse{}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(sendResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func sendMessage(ctx context.Context, userID, conversationID, content string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(ctx, `insert into messages (user_id, conversation_id, content) values ($1, $2, $3)`)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, conversationID, userID, content); err != nil {
		return err
	}

	return nil
}
