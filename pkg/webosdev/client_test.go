package webosdev

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type emuResponse struct {
	status int
	body   string
}

func newHTTPServer(t *testing.T, wantPath, wantToken string, emu emuResponse) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, wantPath, r.URL.Path)
		assert.Equal(t, wantToken, r.URL.Query().Get("sessionToken"))
		if emu.status != 0 {
			w.WriteHeader(emu.status)
		}
		_, _ = w.Write([]byte(emu.body))
	}))
}

func TestClient_request(t *testing.T) {
	tests := []struct {
		name    string
		emuRes  emuResponse
		want    *response
		wantErr require.ErrorAssertionFunc
	}{
		{
			"success",
			emuResponse{body: `{"result":"success","errorCode":"200","errorMsg":"GNL"}`},
			&response{Result: "success", ErrorCode: "200", ErrorMessage: "GNL"},
			require.NoError,
		},
		{"failure", emuResponse{status: http.StatusInternalServerError}, nil, require.Error},
		{"invalid JSON", emuResponse{body: `{"result":`}, nil, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newHTTPServer(t, "/test", "test", tt.emuRes)
			t.Cleanup(srv.Close)

			c := NewClient(
				WithSessionToken("test"),
				WithBaseURL(srv.URL),
			)
			got, err := c.request(context.Background(), "/test")
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

	t.Run("invalid base URL", func(t *testing.T) {
		c := NewClient(WithBaseURL("\x1B"))
		got, err := c.request(context.Background(), "/test")
		require.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("connection error", func(t *testing.T) {
		c := NewClient(WithBaseURL("http://"))
		got, err := c.request(context.Background(), "/test")
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestClient_ExtendSession(t *testing.T) {
	tests := []struct {
		name    string
		emuRes  emuResponse
		wantErr require.ErrorAssertionFunc
	}{
		{"success", emuResponse{body: `{"result":"success","errorCode":"200","errorMsg":"GNL"}`}, require.NoError},
		{"failure", emuResponse{body: `{"result":"fail","errorCode":"ERR_005","errorMsg":"Check user session"}`}, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newHTTPServer(t, "/secure/ResetDevModeSession.dev", "test", tt.emuRes)
			t.Cleanup(srv.Close)

			c := NewClient(
				WithSessionToken("test"),
				WithBaseURL(srv.URL),
			)
			tt.wantErr(t, c.ExtendSession(context.Background()))
		})
	}
}

func TestClient_CheckExpiration(t *testing.T) {
	tests := []struct {
		name    string
		emuRes  emuResponse
		want    time.Duration
		wantErr require.ErrorAssertionFunc
	}{
		{"failure", emuResponse{body: `{"result":"success","errorCode":"200","errorMsg":"1000:00:00"}`}, 1000 * time.Hour, require.NoError},
		{"success", emuResponse{body: `{"result":"fail","errorCode":"ERR_005","errorMsg":"Check user session"}`}, 0, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newHTTPServer(t, "/secure/CheckDevModeSession.dev", "test", tt.emuRes)
			t.Cleanup(srv.Close)

			c := NewClient(
				WithSessionToken("test"),
				WithBaseURL(srv.URL),
			)
			got, err := c.CheckExpiration(context.Background())
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseDuration(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr require.ErrorAssertionFunc
	}{
		{"1000h", args{"1000:00:00"}, 1000 * time.Hour, require.NoError},
		{"9999h59m59s", args{"9999:59:59"}, 9999*time.Hour + 59*time.Minute + 59*time.Second, require.NoError},
		{"bad number of semicolons", args{"1000:00"}, 0, require.Error},
		{"bad hour format", args{"test:00:00"}, 0, require.Error},
		{"bad minute format", args{"00:test:00"}, 0, require.Error},
		{"bad second format", args{"00:00:test"}, 0, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.args.str)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
