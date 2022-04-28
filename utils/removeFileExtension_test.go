package utils_test

import (
	"testing"

	"github.com/dmsi-io/api-generator/utils"
	"github.com/stretchr/testify/assert"
)

func Test_RemoveFileExtension(t *testing.T) {
	filename := "file.txt"

	cleaned := utils.RemoveFileExtension(filename)

	assert.NotEqual(t, filename, cleaned)
	assert.Equal(t, "file", cleaned)
}

func Test_RemoveFileExtension_NoExtension(t *testing.T) {
	filename := "file"

	cleaned := utils.RemoveFileExtension(filename)

	assert.Equal(t, filename, cleaned)
}
