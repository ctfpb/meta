package meta

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseBytes(t *testing.T) {
	err := filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, "example") {
			return nil
		}
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			buf, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			meta, err := Parse(buf)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(meta)
			if err != nil {
				t.Fatal(err)
			}
			println(string(data))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestParseEveryFile(t *testing.T) {
	err := filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "example") {
			return nil
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			meta, err := Parse(buf)
			if err != nil {
				t.Fatal(err)
			}
			err = meta.Check()
			if err != nil {
				if !strings.Contains(path, "err") {
					t.Log(path)
					t.Fatal(err)
				}
			}
			meta.ParseFormat()
			data, err := json.Marshal(meta)
			if err != nil {
				if !strings.Contains(path, "err") {
					t.Log(path)
					t.Fatal(err)
				}
			}
			t.Log("data in filename =", path, " is ok!")
			println(string(data))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTemplate(t *testing.T) {
	println(Template())
}
