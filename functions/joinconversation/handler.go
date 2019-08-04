package function

import (
	"bytes"
	"context"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	createRequest := &CreateConversationRequest{}
	if err := jsonpb.Unmarshal(bytes.NewReader(req), createRequest); err != nil {
		return "", err
	}

	if createRequest.UserId == "" {
		return "", errors.New("user id is empty")
	}

	if createRequest.Token == "" {
		return "", errors.New("token is empty")
	}

	if createRequest.ConversationName == "" {
		return "", errors.New("conversation name is empty")
	}

	if authenticated, err := isAuthenticated(createRequest.UserId, createRequest.Token); err != nil {
		return "", err
	} else if !authenticated {
		return "", errors.New("not authenticated")
	}

	ctx := context.Background()

	conversationID, err := createConversation(ctx, createRequest.ConversationName)
	if err != nil {
		return "", err
	}

	createResponse := &CreateConversationResponse{
		ConversationId: conversationID,
	}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(createResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func createConversation(ctx context.Context, name string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}

	stmt, err := db.PrepareContext(ctx, `insert into conversations (conversation_name) values ($1) returning conversation_id`)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRow(name)

	var conversationID string
	if err := row.Scan(&conversationID); err != nil {
		return "", err
	}

	return conversationID, nil
}
