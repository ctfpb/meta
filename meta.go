package meta

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"

	"gopkg.in/yaml.v3"
)

type Meta struct {
	Author struct {
		Name    string `yaml:"name,omitempty"`
		Contact string `yaml:"contact,omitempty"`
	} `yaml:"author,omitempty"`
	Task struct {
		Name          string   `yaml:"name"`
		Type          string   `yaml:"type"`
		Description   string   `yaml:"description,omitempty"`
		Level         string   `yaml:"level,omitempty"`
		Flag          string   `yaml:"flag,omitempty"`
		AttachmentURL string   `yaml:"attachment_url,omitempty"`
		Hints         []string `yaml:"hints,omitempty"`
	} `yaml:"task"`
	Challenge struct {
		Name  string   `yaml:"name,omitempty"`
		Refer string   `yaml:"refer,omitempty"`
		Tags  []string `yaml:"tags,omitempty"`
	} `yaml:"challenge,omitempty"`
}

func Template() string {
	m := &Meta{}

	m.Author.Name = "陌竹"
	m.Author.Contact = "mozhu233@outlook.com"
	m.Task.Name = "challenge_game_2023_web_abc"
	m.Task.Type = "Web"
	m.Task.Description = "这是一个模板"
	m.Task.Level = "easy"
	m.Task.Flag = "ctftrain{this_is_a_test_flag}"
	m.Task.Hints = []string{"这是一个模板", "没有提示"}
	m.Challenge.Name = "Web题目"
	m.Challenge.Refer = "https://www.ctfhub.com"
	m.Challenge.Tags = []string{"模板"}

	buf, err := yaml.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func ParseMetaFromFile(fds ...multipart.File) ([]*Meta, error) {
	var data bytes.Buffer
	for _, fd := range fds {
		buf, err := io.ReadAll(fd)
		if err != nil {
			return nil, err
		}
		data.WriteString("---\n")
		data.Write(buf)
	}
	return ParseMetas(data.Bytes())
}

func ParseMetas(data []byte) ([]*Meta, error) {
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
		metas = append(metas, &meta)
	}
	if len(metas) == 0 {
		return metas, errors.New("failed to parse meta data")
	}
	return metas, nil
}
