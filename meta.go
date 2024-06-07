package meta

import (
	"encoding/json"
	"io"

	pb "buf.build/gen/go/ctfhub/meta/protocolbuffers/go"
)

func New(name, contact string) *pb.Meta {
	return &pb.Meta{Author: &pb.Author{Name: name, Contact: contact}}
}
func Empty() *pb.Meta   { return &pb.Meta{} }
func Default() *pb.Meta { return New("陌竹", "mozhu233@outlook.com") }

func Parse(data []byte) (*pb.Meta, error) {
	meta := &pb.Meta{}
	err := json.Unmarshal(data, meta)
	if err != nil {
		return nil, err
	}
	if err = Verify(meta); err != nil {
		return nil, err
	}
	return ParseFormatMeta(meta), nil
}

func ParseFromFile(fds ...io.ReadSeekCloser) ([]*pb.Meta, error) {
	var metas = []*pb.Meta{}
	for _, fd := range fds {
		buf, err := io.ReadAll(fd)
		if err != nil {
			return nil, err
		}
		meta, err := Parse(buf)
		if err != nil {
			return nil, err
		}
		metas = append(metas, meta)
	}
	return metas, nil
}
