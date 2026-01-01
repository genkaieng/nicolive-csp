package nicolive

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/genkaieng/nicolive-csp/gen/pb"
	"google.golang.org/protobuf/proto"
)

func decodeChunkedEntry(r io.Reader) (<-chan *pb.ChunkedEntry, error) {
	ch := make(chan *pb.ChunkedEntry)
	go func() {
		defer close(ch)

		br := bufio.NewReader(r)
		for {
			size, err := binary.ReadUvarint(br)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			body := make([]byte, size)
			_, err = io.ReadFull(br, body)
			if err != nil {
				panic(err)
			}

			var chunkedEntry pb.ChunkedEntry
			if err := proto.Unmarshal(body, &chunkedEntry); err != nil {
				panic(err)
			}
			ch <- &chunkedEntry
		}
	}()
	return ch, nil
}

func decodeChunkedMessage(r io.Reader) (<-chan *pb.ChunkedMessage, error) {
	ch := make(chan *pb.ChunkedMessage)
	go func() {
		defer close(ch)

		br := bufio.NewReader(r)
		for {
			size, err := binary.ReadUvarint(br)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			body := make([]byte, size)
			_, err = io.ReadFull(br, body)
			if err != nil {
				panic(err)
			}

			var chunkedMessage pb.ChunkedMessage
			if err := proto.Unmarshal(body, &chunkedMessage); err != nil {
				panic(err)
			}
			ch <- &chunkedMessage
		}
	}()
	return ch, nil
}
