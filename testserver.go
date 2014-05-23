package zkutils

import (
	"fmt"
	zk "github.com/koofr/gozk"
	"io/ioutil"
	"os"
)

type TestServer struct {
	server *zk.Server
	zkRun  string
}

func NewTestServer(port int) (s *TestServer, err error) {
	zkDir := os.Getenv("ZKROOT")

	if zkDir == "" {
		err = fmt.Errorf("ZKROOT env variable missing")
		return
	}

	zkRun, err := ioutil.TempDir("", "zkRun")

	if err != nil {
		return
	}

	server, err := zk.CreateServer(port, zkRun, zkDir)

	err = server.Start()

	if err != nil {
		return
	}

	s = &TestServer{
		server: server,
		zkRun:  zkRun,
	}

	return
}

func (s *TestServer) Stop() error {
	defer os.RemoveAll(s.zkRun)

	return s.server.Destroy()
}
