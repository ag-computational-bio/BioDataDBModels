package client

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ag-computational-bio/BioDataDBModels/go/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// TokenType Indicates the kind of token passed for authentication
type TokenType string

const (
	// AccessToken for Oauth2 authentication
	AccessToken TokenType = "AccessToken"
	// UserAPIToken for api token authentication
	UserAPIToken TokenType = "UserAPIToken"
)

// GRPCEndpointsClients Holds all grpc endpoint clients and necessary contexts
type GRPCEndpointsClients struct {
	DatasetBackend         api.DatasetServiceClient
	ProjectBackend         api.ProjectAPIClient
	LoadBackend            api.LoadServiceClient
	TokenBackend           api.APITokenServiceClient
	ObjectsBackend         api.ObjectsServiceClient
	GenericOutGoingContext context.Context
}

// New Creates and initializes a new set of GRPCEndpointsClients
func (clients *GRPCEndpointsClients) New(host string, port string, token string, tokentype string) error {
	ctx := context.Background()
	clients.GenericOutGoingContext = ctx

	var tlsConf tls.Config

	credentials := credentials.NewTLS(&tlsConf)

	dialOptions := grpc.WithTransportCredentials(credentials)
	if host == "localhost" {
		dialOptions = grpc.WithInsecure()
	}

	dialOptions = grpc.WithInsecure()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", host, port), dialOptions)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	clients.createDatasetClient(conn)
	clients.createProjectClient(conn)
	clients.createTokenClient(conn)
	clients.createLoadClient(conn)
	clients.createObjectsClient(conn)

	return nil
}

func (clients *GRPCEndpointsClients) createObjectsClient(conn *grpc.ClientConn) {
	clients.ObjectsBackend = api.NewObjectsServiceClient(conn)
}

func (clients *GRPCEndpointsClients) createLoadClient(conn *grpc.ClientConn) {
	clients.LoadBackend = api.NewLoadServiceClient(conn)
}

func (clients *GRPCEndpointsClients) createTokenClient(conn *grpc.ClientConn) {
	clients.TokenBackend = api.NewAPITokenServiceClient(conn)
}

func (clients *GRPCEndpointsClients) createDatasetClient(conn *grpc.ClientConn) {
	clients.DatasetBackend = api.NewDatasetServiceClient(conn)
}

func (clients *GRPCEndpointsClients) createProjectClient(conn *grpc.ClientConn) {
	clients.ProjectBackend = api.NewProjectAPIClient(conn)
}

// OutGoingContextFromToken Creates the required outgoing context for a call
func (clients *GRPCEndpointsClients) OutGoingContextFromToken(token string, tokentype TokenType) context.Context {
	mdMap := make(map[string]string)
	mdMap[string(tokentype)] = token
	tokenMetadata := metadata.New(mdMap)

	outgoingContext := metadata.NewOutgoingContext(context.TODO(), tokenMetadata)
	return outgoingContext
}

// GetAccessTokenFromGinContext Returns the access token of a gin context from cookie
// Token needs to be stored as "token"
func (clients *GRPCEndpointsClients) GetAccessTokenFromGinContext(c *gin.Context) string {
	rawTokenCookie, err := c.Request.Cookie("token")
	if err != nil && err != http.ErrNoCookie {
		log.Println(err.Error())
		c.AbortWithError(400, err)
	}

	if err == http.ErrNoCookie {
		log.Println("cookie not found")
		return ""
	}

	unescapedBase64Data, err := url.QueryUnescape(rawTokenCookie.Value)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/index")
		c.AbortWithError(400, err)
	}

	rawBytesDecoded, err := base64.StdEncoding.DecodeString(unescapedBase64Data)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/index")
		c.AbortWithError(400, err)
	}

	var token oauth2.Token

	err = json.Unmarshal([]byte(rawBytesDecoded), &token)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, err)
	}

	return token.AccessToken
}
