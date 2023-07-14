package main

import (
	"context"
	"fmt"
	gen "main/gen/go"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in allowed header, don't send the header
	return s, false
}

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.FormValue("auth_key")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Token must be filled"})
			return
		}
		if !tokenAuthorized(token) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Token invalido"})
			return
		}

		ctx.Next()
	}

}

func tokenAuthorized(token string) bool {
	value := GetValue(nil, "dh_"+token)
	if value == "" {
		return false
	}
	fmt.Println(value)
	//value := service.GetPg("frf", 4).Nonce
	return value != ""
}

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux(
		// convert header in response(going from gateway) from metadata received.
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}))
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := gen.RegisterAuthHandlerFromEndpoint(context.Background(), mux, "authserver:5052", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	mux2 := runtime.NewServeMux(
		// convert header in response(going from gateway) from metadata received.
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux2 *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux2, marshaler, writer, request, &newError)
		}))
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err2 := gen.RegisterBizHandlerFromEndpoint(context.Background(), mux2, "bizserver:5062", []grpc.DialOption{grpc.WithInsecure()})
	if err2 != nil {
		log.Fatal(err2)
	}
	// Creating a normal HTTP server
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(CheckBlocked())

	//server.Use(TokenAuthMiddleware())
	server.Group("/req_pq", ApplyLeakBucket()).Any("", gin.WrapH(mux))
	server.Group("/req_dh_params", ApplyLeakBucket()).Any("", gin.WrapH(mux))
	server.Group("/get_users", TokenAuthMiddleware()).Any("", gin.WrapH(mux2))
	server.Group("/get_users_sql", TokenAuthMiddleware()).Any("", gin.WrapH(mux2))

	// start server
	err = server.Run(":6433")
	if err != nil {
		log.Fatal(err)
	}
}
