package testing

import (
	"chi-example/utils"
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	var responsePaths []string
	resultOk := [][]string{{"path1"}, {"path2"}}
	responsePaths = append(responsePaths, "path1", "path2")
	chunks := utils.ChunkSlice(responsePaths, 1)
	if !reflect.DeepEqual(chunks, resultOk) {
		t.Errorf("ChunkSlide was incorrect, got %s expexted %s", chunks, resultOk)
	}
}
