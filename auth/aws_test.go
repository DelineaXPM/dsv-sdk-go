package auth

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
)

func TestGetSTSHeaderAndBody(t *testing.T) {
	testCases := []struct {
		name              string
		getCallerIdentity func() *request.Request
		expectedHeader    string
		expectedBody      string
		expectedError     error
	}{
		{"good Response", func() *request.Request {
			return &request.Request{HTTPRequest: &http.Request{
				Header: http.Header{},
				Body:   io.NopCloser(strings.NewReader("Test data")),
			}}
		}, "e30=", "VGVzdCBkYXRh", nil},
	}

	for _, tt := range testCases {
		ath, _ := New(Config{Provider: AWS})
		ath.getCallerIdentity = tt.getCallerIdentity

		header, body, err := ath.GetSTSHeaderAndBody()

		if !errors.Is(err, tt.expectedError) {
			t.Error("unexpected err", err)
		}

		if !reflect.DeepEqual(header, tt.expectedHeader) {
			t.Error("unexpected header", header)
		}

		if !reflect.DeepEqual(body, tt.expectedBody) {
			t.Error("unexpected body", body)
		}
	}
}
