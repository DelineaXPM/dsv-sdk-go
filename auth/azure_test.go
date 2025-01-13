package auth

import (
	"errors"
	_"io"
	_"net/http"
	_"reflect"
	_"strings"
	"testing"
	
	"github.com/golang-jwt/jwt/v5"
)

func TestBuildAzureParams(t *testing.T) {
	var secretKey = []byte("secret-key")

	testCases := []struct {
		name          string
		GrantType     string
		ValidJwt      bool
		err1          error
		err2          error
		err3          error
	}{
		{"test0", "azure", true, nil, nil, errors.New("token signature is invalid: key is of invalid type: RSA verify expects *rsa.PublicKey")},
	}

	for _, tt := range testCases {
		ath, _ := New(Config{Provider: AZURE})
		data, err := ath.BuildAzureParams()
		if err != tt.err1 {
			t.Error("unexpected err", err)
		}

		token, err := jwt.Parse(data.Jwt, func(token *jwt.Token) (interface{}, error) {
			  return secretKey, nil
		   })
		  
		   if err != tt.err3 {
			t.Error("unexpected err1:", err)
		   }
		  
		   if !token.Valid {
			t.Error("invalid token:", err)
		   }
   
/*
		if !reflect.DeepEqual(header, tt.expectedHeader) {
			t.Error("unexpected header", header)
		}

		if !reflect.DeepEqual(body, tt.expectedBody) {
			t.Error("unexpected body", body)
		}
*/
	}
}
