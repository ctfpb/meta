package meta

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"

	"gopkg.in/yaml.v3"
)

func New(name, contact string) *Meta { return &Meta{Author: &Author{Name: name, Contact: contact}} }
func Empty() *Meta                   { return &Meta{} }
func Default() *Meta                 { return New("陌竹", "mozhu233@outlook.com") }
func NewSkill(id, pid, tid string, leaf int64, image, name string, level TaskLevel) *Meta {
	return Default().NewSkill(id, pid, tid, leaf, image, name, level)
}

func ParseFromFile(fds ...multipart.File) ([]*Meta, error) {
	var data bytes.Buffer
	for _, fd := range fds {
		buf, err := io.ReadAll(fd)
		if err != nil {
			return nil, err
		}
		data.WriteString("---\n")
		data.Write(buf)
	}
	return ParseBytes(data.Bytes())
}

func ParseBytes(data []byte) ([]*Meta, error) {
	var metas []*Meta
	dec := yaml.NewDecoder(bytes.NewReader(data))
	for {
		meta := Meta{}
		err := dec.Decode(&meta)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
			break
		}
		metas = append(metas, meta.parseFormat())
	}
	if len(metas) == 0 {
		return metas, errors.New("failed to parse meta data")
	}
	return metas, nil
}
