package function

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"github.com/gogo/protobuf/jsonpb"
	"time"
)

// Handle a serverless request
func Handle(req []byte) (string, error) {
	sendRequest := &ReadMessagesRequest{}
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

	if sendRequest.FromTimestamp == "" {
		return "", errors.New("from timestamp is empty")
	}

	if sendRequest.ToTimestamp == "" {
		return "", errors.New("to timestamp is empty")
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

	fromTimestamp, err := time.Parse(time.RFC3339, sendRequest.FromTimestamp)
	if err != nil {
		return "", err
	}

	toTimestamp, err := time.Parse(time.RFC3339, sendRequest.ToTimestamp)
	if err != nil {
		return "", err
	}

	messages, err := readMessages(ctx, sendRequest.ConversationId, fromTimestamp, toTimestamp)
	if err != nil {
		return "", err
	}

	readResponse := &ReadMessagesResponse{
		Messages: messages,
	}

	res, err := (&jsonpb.Marshaler{}).MarshalToString(readResponse)
	if err != nil {
		return "", err
	}

	return res, nil
}

func readMessages(ctx context.Context, conversationID string, fromTimestamp, toTimestamp time.Time) ([]*ReadMessagesResponse_Message, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, `select timestamp, user_id, content from messages where timestamp > $1 and timestamp <= $2 and conversation_id = $3`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, fromTimestamp, toTimestamp, conversationID)
	if err != nil {
		return nil, err
	}

	messages := make([]*ReadMessagesResponse_Message, 0)

	for rows.Next() {
		if err := rows.Err(); err != nil {
			if err == sql.ErrNoRows {
				break
			}

			return nil, err
		}

		message := &ReadMessagesResponse_Message{}
		var timestamp time.Time

		if err := rows.Scan(&timestamp, &message.UserId, &message.Content); err != nil {
			return nil, err
		}

		message.Timestamp = timestamp.Format(time.RFC3339)
		messages = append(messages, message)
	}

	return messages, nil
}
