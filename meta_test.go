package meta

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestTemplate(t *testing.T) {
	fmt.Println(Template())
}

func TestR(t *testing.T) {
	m := Default()
	n := m.R()
	m.Author.Name = "aaaaa"
	fmt.Printf("%p %+v\n", &m, m)
	fmt.Printf("%p %+v\n", &n, n)
}

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
		fmt.Println("==========")
		buf, _ := yaml.Marshal(meta)
		fmt.Println(string(buf))
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
				t.Log(err)
				return nil
			}
			metas, err := ParseBytes(buf)
			if err != nil {
				t.Log(err)
				return nil
			}
			for _, meta := range metas {
				err = meta.Check()
				if err != nil {
					t.Log(err)
					continue
				}
				meta.ParseFormat()
				_, err := yaml.Marshal(meta)
				if err != nil {
					t.Log(err)
					continue
				}
				if strings.Contains(path, "ok") {
					t.Log("data in filename =", path, " is ok!")
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
