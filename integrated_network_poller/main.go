package main

import (
	"io"
	"net"
	"os"
	"syscall"
	"time"
)

func checkClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}

const timeout = 3 * time.Second

func ReadFile(c *net.UnixConn, timeout time.Duration) (*os.File, error) {
	oob := make([]byte, 64)

	if timeout > 0 {
		deadline := time.Now().Add(timeout)
		if err := c.SetReadDeadline(deadline); err != nil {
			return nil, err
		}
	}

	_, oobn, flags, _, err := c.ReadMsgUnix(nil, oob)
	if err != nil {
		return nil, err
	}
	if flags != 0 || oobn <= 0 {
		panic("ReadMsgUnix: flags != 0 || oobn <= 0")
	}

	// file descriptors are now open in this process

	scm, err := syscall.ParseSocketControlMessage(oob[:oobn])
	if err != nil {
		return nil, err
	}
	if len(scm) != 1 {
		panic("invalid scm message")
	}

	fds, err := syscall.ParseUnixRights(&scm[0])
	if err != nil {
		return nil, err
	}
	if len(fds) != 1 {
		panic("invalid scm message")
	}

	return os.NewFile(uintptr(fds[0]), ""), nil
}

func WriteFile(c *net.UnixConn, file *os.File, timeout time.Duration) error {
	if timeout > 0 {
		deadline := time.Now().Add(timeout)
		if err := c.SetWriteDeadline(deadline); err != nil {
			return err
		}
	}
	oob := syscall.UnixRights(int(file.Fd()))
	_, _, err := c.WriteMsgUnix(nil, oob, nil)
	return err
}

func ReadUnixConn(c *net.UnixConn, timeout time.Duration) (*net.UnixConn, error) {
	file, err := ReadFile(c, timeout)
	if err != nil {
		return nil, err
	}
	defer checkClose(file)
	fdConn, err := net.FileConn(file)
	if err != nil {
		return nil, err
	}
	return fdConn.(*net.UnixConn), nil
}

func WriteUnixConn(c *net.UnixConn, timeout time.Duration) (*net.UnixConn, error) {
	local, remote := UnixPair()
	defer checkClose(remote)
	file, err := remote.File()
	if err != nil {
		return nil, err
	}
	defer checkClose(file)
	if err := WriteFile(c, file, timeout); err != nil {
		defer checkClose(local)
		return nil, err
	}
	return local, nil
}

func unixPair(typ int) (*net.UnixConn, *net.UnixConn) {
	fds, err := syscall.Socketpair(syscall.AF_UNIX, typ, 0)
	if err != nil {
		panic(os.NewSyscallError("socketpair", err))
	}
	file := os.NewFile(uintptr(fds[0]), "")
	c0, err := net.FileConn(file)
	if err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
	file = os.NewFile(uintptr(fds[1]), "")
	c1, err := net.FileConn(file)
	if err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
	return c0.(*net.UnixConn), c1.(*net.UnixConn)
}

func UnixPair() (*net.UnixConn, *net.UnixConn) {
	return unixPair(syscall.SOCK_STREAM)
}

func newClientServerConn() []*net.UnixConn {
	c1, c2 := UnixPair()

	ch := make(chan struct{})
	var remote *net.UnixConn
	go func() {
		var err error
		remote, err = WriteUnixConn(c2, timeout)
		if err != nil {
			panic(err)
		}
		ch <- struct{}{}
	}()

	local, err := ReadUnixConn(c1, timeout)
	if err != nil {
		panic(err)
	}

	<-ch

	return []*net.UnixConn{c1, c2, local, remote}
}

func main() {
	for i := 0; i < 10; i++ {
		conns := newClientServerConn()
		for _, c := range conns {
			if err := c.Close(); err != nil {
				panic(err)
			}
		}
	}
}
