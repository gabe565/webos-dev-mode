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

func TestClient_ExtendSession(t *testing.T) {
	tests := []struct {
		name     string
		response string
		wantErr  require.ErrorAssertionFunc
	}{
		{"valid", `{"result":"success","errorCode":"200","errorMsg":"GNL"}`, require.NoError},
		{"failure", `{"result":"fail","errorCode":"ERR_005","errorMsg":"Check user session"}`, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/secure/ResetDevModeSession.dev", r.URL.Path)
				assert.Equal(t, "test", r.URL.Query().Get("sessionToken"))
				_, _ = w.Write([]byte(tt.response))
			}))
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
		name     string
		response string
		want     time.Duration
		wantErr  require.ErrorAssertionFunc
	}{
		{"1000h", `{"result":"success","errorCode":"200","errorMsg":"1000:00:00"}`, 1000 * time.Hour, require.NoError},
		{"999h59m59s", `{"result":"success","errorCode":"200","errorMsg":"999:59:59"}`, 999*time.Hour + 59*time.Minute + 59*time.Second, require.NoError},
		{"failure response", `{"result":"fail","errorCode":"ERR_005","errorMsg":"Check user session"}`, 0, require.Error},
		{"bad number of semicolons", `{"result":"success","errorCode":"200","errorMsg":"1000:00"}`, 0, require.Error},
		{"bad hour format", `{"result":"success","errorCode":"200","errorMsg":"test:00:00"}`, 0, require.Error},
		{"bad minute format", `{"result":"success","errorCode":"200","errorMsg":"00:test:00"}`, 0, require.Error},
		{"bad second format", `{"result":"success","errorCode":"200","errorMsg":"00:00:test"}`, 0, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/secure/CheckDevModeSession.dev", r.URL.Path)
				assert.Equal(t, "test", r.URL.Query().Get("sessionToken"))
				_, _ = w.Write([]byte(tt.response))
			}))
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
