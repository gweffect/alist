package tool_test

import (
	"testing"

	"github.com/gweffectx/safedav/internal/offline_download/tool"
)

func TestGetFiles(t *testing.T) {
	files, err := tool.GetFiles("..")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Log(file.Name, file.Size, file.Path, file.Modified)
	}
}
