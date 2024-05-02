package meta

import (
	"encoding/json"
	"io"
)

func New(name, contact string) *Meta { return &Meta{Author: &Author{Name: name, Contact: contact}} }
func Empty() *Meta                   { return &Meta{} }
func Default() *Meta                 { return New("陌竹", "mozhu233@outlook.com") }

func ParseFromFile(fds ...io.ReadSeekCloser) ([]*Meta, error) {
	var metas = []*Meta{}
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

func Parse(data []byte) (*Meta, error) {
	meta := &Meta{}
	err := json.Unmarshal(data, meta)
	if err != nil {
		return nil, err
	}
	if err = meta.Check(); err != nil {
		return nil, err
	}
	return meta.ParseFormat(), nil
}
