package main

// simple echo server

import (
	_ "bytes"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/http2"
	"io"
	"net"
)

var lport string

func init() {
	flag.StringVar(&lport, "lport", "7777", "local port to listen at")
	log.SetLevel(log.DebugLevel)
}

var http2Preface = []byte(http2.ClientPreface)

func FrameTransfer(f *http2.Framer) {
	for {
		fr, err := f.ReadFrame()
		if err != nil {
			log.Errorf("read frame error %s", err)
			return
		} else {
			log.Debugf("frame: %v", fr)
		}
		hdr := fr.Header()
		switch v := fr.(type) {
		case *http2.SettingsFrame:
			settings := []http2.Setting{}
			v.ForeachSetting(func(s http2.Setting) error {
				settings = append(settings, s)
				return nil
			})
			f.WriteSettings(settings[:]...)
		case *http2.HeadersFrame:
			hfp := http2.HeadersFrameParam{
				BlockFragment: v.HeaderBlockFragment(),
				EndHeaders:    v.HeadersEnded(),
				EndStream:     v.StreamEnded(),
				Priority:      v.Priority,
				StreamID:      v.StreamID,
			}
			f.WriteHeaders(hfp)
		case *http2.DataFrame:
			f.WriteData(hdr.StreamID, hdr.Flags == http2.FlagDataEndStream, v.Data())
		case *http2.WindowUpdateFrame:
			f.WriteWindowUpdate(hdr.StreamID, v.Increment)
		case *http2.PingFrame:
			f.WritePing(v.IsAck(), v.Data)
		default:
			log.Errorf("not support this frame type")
			continue
		}

		/*
			// read out all response data
			if hdr.Type == http2.FrameData && hdr.Flags == http2.FlagDataEndStream {
				buf := make([]byte, 2048)
				n, err := remote.Read(buf)
				if err != nil {
					log.Errorf("read from remote failed %s", err)
				} else {
					log.Debugf("read from remote %d, content", n, buf[:n])
				}
				conn.Write(buf[:n])
				log.Debugf("write one response")
			}
		*/

	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	remote, _ := net.Dial("tcp", "127.0.0.1:50051")

	go io.Copy(remote, conn)
	go io.Copy(conn, remote)

	/*
		f := http2.NewFramer(remote, conn)

		// read preface
		preface := make([]byte, len(http2Preface))
		_, err := io.ReadFull(conn, preface)
		if err != nil {
			log.Errorf("read client preface error %s", err)
			return
		} else {
			if !bytes.Equal(preface, http2Preface) {
				log.Errorf("Preface not valid")
			}
			log.Debugf("got preface %s", preface)
		}
		remote.Write(preface)

		nf := http2.NewFramer(conn, remote)

		go FrameTransfer(f)
		go FrameTransfer(nf)

	*/
	select {}

}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", lport))
	if err != nil {
		fmt.Printf("listen err %v\n", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept err %v\n", err)
			break
		}
		fmt.Printf("got a conn\n")
		go handleConn(conn)
	}

}
