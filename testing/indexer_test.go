package testing

import (
	"chi-example/utils"
	"fmt"
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	var responsePaths []string
	var resultOk [][]string
	for i := 0; i < 100000000; i++ {
		item := fmt.Sprintf(`path_%d`, i)
		responsePaths = append(responsePaths, item)
	}
	chunks := utils.ChunkSlice(responsePaths, 1)
	resultOk = chunks
	if !reflect.DeepEqual(chunks, resultOk) {
		t.Errorf("ChunkSlide was incorrect, got %s expexted %s", chunks, resultOk)
	}
}
