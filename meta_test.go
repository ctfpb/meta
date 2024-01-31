package meta

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"
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

func TestParseYamls(t *testing.T) {
	var data bytes.Buffer
	for _, name := range []string{"web1.yaml", "web2.yaml"} {
		buf, err := os.ReadFile(path.Join(".", "test", name))
		if err != nil {
			t.Fatal(err)
		}
		data.WriteString("---\n")
		data.Write(buf)
	}
	metas, err := ParseBytes(data.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	for _, meta := range metas {
		fmt.Println("==========")
		fmt.Println(meta)
	}
}
