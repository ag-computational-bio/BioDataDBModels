package client

import (
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestGRPCEndpointsClients_OutGoingContextAccessToken(t *testing.T) {
	testStruct := GRPCEndpointsClients{}
	tokenString := "testToken"

	outctx := testStruct.OutGoingContext(tokenString, AccessToken)

	metadata, ok := metadata.FromOutgoingContext(outctx)
	if !ok {
		t.Fatal("Error on outgoing metadata extraction")
	}

	accessToken := metadata.Get("AccessToken")

	if accessToken[0] != tokenString {
		t.Errorf("Bad context in outgoing context for AccessToken token type")
	}

}

func TestGRPCEndpointsClients_OutGoingContextUserAPIToken(t *testing.T) {
	testStruct := GRPCEndpointsClients{}
	tokenString := "testToken"

	outctx := testStruct.OutGoingContext(tokenString, UserAPIToken)

	metadata, ok := metadata.FromOutgoingContext(outctx)
	if !ok {
		t.Fatal("Error on outgoing metadata extraction")
	}

	accessToken := metadata.Get("UserAPIToken")

	if accessToken[0] != tokenString {
		t.Errorf("Bad context in outgoing context for AccessToken token type")
	}

}
