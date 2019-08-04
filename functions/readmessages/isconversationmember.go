package function

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"log"
	"net/http"
	"os"
)

func getIsConversationMemberEndpoint() string {
	return os.Getenv("ENDPOINT_ISCONVERSATIONMEMBER")
}

func isConversationMember(userID, token, conversationID string) (bool, error) {
	body := bytes.NewBuffer([]byte{})
	if err := (&jsonpb.Marshaler{}).Marshal(body, &IsConversationMemberRequest{
		Token:          token,
		UserId:         userID,
		ConversationId: conversationID,
	}); err != nil {
		return false, fmt.Errorf("marshalling request: %s", err.Error())
	}

	res, err := http.Post(getIsConversationMemberEndpoint(), "application/json", body)
	if err != nil {
		return false, fmt.Errorf("posting to url: %s", err.Error())
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("error closing response body: %s\n", err.Error())
		}
	}()

	isConversationMemberResponse := &IsConversationMemberResponse{}
	if err := jsonpb.Unmarshal(res.Body, isConversationMemberResponse); err != nil {
		return false, fmt.Errorf("unmarshalling response: %s", err.Error())
	}

	return isConversationMemberResponse.Member, nil
}
