package new_redis

import (
	"fmt"
	"net"
)

type Server struct {
	db      *Database
	clients []*Client

	conf *Config
}

func NewServer(c *Config) (*Server, error) {
	if c == nil {
		return &Server{
			conf: DefaultConfig(),
		}, nil
	}

	return &Server{conf: c}, nil
}

func (s *Server) ListenAndServe() error {
	addr := fmt.Sprintf("%s:%d", s.conf.host, s.conf.port)
	l, e := net.Listen(s.conf.proto, addr)
	if e != nil {
		return e
	}
	return s.Serve(l)
}

// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each.  The service goroutines read requests and
// then call srv.Handler to reply to them.
func (s *Server) Serve(l net.Listener) error {
	defer l.Close()
	for {
		rw, err := l.Accept()
		if err != nil {
			return err
		}
		go s.ServeClient(rw)
	}
}

// Serve starts a new redis session, using `conn` as a transport.
// It reads commands using the redis protocol, passes them to `handler`,
// and returns the result.
func (s *Server) ServeClient(conn net.Conn) (err error) {
	defer func() {
		if err != nil {
			fmt.Fprintf(conn, "-%s\n", err)
		}
		conn.Close()
	}()

	var clientAddr string

	switch co := conn.(type) {
	case *net.UnixConn:
		f, err := conn.(*net.UnixConn).File()
		if err != nil {
			return err
		}
		clientAddr = f.Name()
	default:
		clientAddr = co.RemoteAddr().String()
	}

	for {
		client := NewClient([]byte(conn))
		client.ParseRequest().Exec()
	}
	return nil
}
