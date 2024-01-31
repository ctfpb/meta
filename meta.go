package meta

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"

	"gopkg.in/yaml.v3"
)

type Author struct {
	Name    string `yaml:"name"`
	Contact string `yaml:"contact"`
}

type Meta struct {
	Author    Author    `yaml:"author"`
	Task      Task      `yaml:"task"`
	Challenge Challenge `yaml:"challenge"`
	Skill     Skill     `yaml:"skill"`
	Deploy    Deploy    `yaml:"deploy"`
}

func New(name, contact string) *Meta { return &Meta{Author: Author{Name: name, Contact: contact}} }
func Empty() *Meta                   { return &Meta{} }
func Default() *Meta                 { return New("陌竹", "mozhu233@outlook.com") }

func (m *Meta) R() *Meta {
	n := *m
	return &n
}

func Template() string {
	m := Default()
	m.Task.Name = "challenge_game_2023_web_abc"
	m.Task.Type = "web"
	m.Task.Category = "Web"
	m.Task.Description = "这是一个模板"
	m.Task.Level = "easy"
	m.Task.Flag = "ctftrain{this_is_a_test_flag}"
	m.Task.Hints = []string{"这是一个模板", "没有提示"}

	m.Challenge.Name = "Web题目"
	m.Challenge.Refer = "https://www.ctfhub.com"
	m.Challenge.Tags = []string{"web", "2024"}

	m.Deploy = Deploy{
		CPU:      200,
		Mem:      256,
		Protocol: "tcp",
	}

	buf, err := yaml.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func (t *Meta) parseFormat() *Meta {
	t.Task.parseFormat()
	t.Challenge.parseFormat()
	return t
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
