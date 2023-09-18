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
	metas, err := ParseMetas(data.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	for _, meta := range metas {
		fmt.Println("==========")
		fmt.Println(meta)
	}
}
