package meta

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestParseYamlBytes(t *testing.T) {
	var data bytes.Buffer
	err := filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".yaml" {
			buf, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			data.WriteString("---\n")
			data.Write(buf)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	metas, err := ParseBytes(data.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	for _, meta := range metas {
		_, err = yaml.Marshal(meta)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestParseEveryYaml(t *testing.T) {
	err := filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".yaml" {
			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			metas, err := ParseBytes(buf)
			if err != nil {
				if !strings.Contains(path, "err") {
					t.Log(path)
					t.Fatal(err)
				}
			}
			for _, meta := range metas {
				err = meta.Check()
				if err != nil {
					if !strings.Contains(path, "err") {
						t.Log(path)
						t.Fatal(err)
					}
					continue
				}
				meta.ParseFormat()
				_, err := yaml.Marshal(meta)
				if err != nil {
					if !strings.Contains(path, "err") {
						t.Log(path)
						t.Fatal(err)
					}
					continue
				}
				t.Log("data in filename =", path, " is ok!")
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseExample(t *testing.T) {
	buf, err := os.ReadFile("test/example.yaml")
	if err != nil {
		t.Fatal(err)
	}
	metas, err := ParseBytes(buf)
	if err != nil {
		t.Fatal(err)
	}
	if len(metas) == 0 {
		t.Fail()
	}
	meta := metas[0]
	if meta.Task == nil {
		t.Fail()
	}
	if meta.Task.Id != "2022_hitcon_web_rce-me" {
		t.Fail()
	}
}
