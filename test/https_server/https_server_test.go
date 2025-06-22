package https_server

import (
	"discuss/internal/https_server"
	"fmt"
	"testing"
)

func TestHttpsServer(t *testing.T) {

	if https_server.GE.Run(fmt.Sprintf("%s:%d", "127.0.0.1", 8090)) != nil {
		t.Errorf("https server run failed")
	}
	fmt.Println("HTTPS server is running on http://127.0.0.1:8090")

}
