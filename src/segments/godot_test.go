package segments

import (
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	mock2 "github.com/stretchr/testify/mock"

	"github.com/jandedobbeleer/oh-my-posh/src/mock"
	"github.com/jandedobbeleer/oh-my-posh/src/platform"
	"github.com/jandedobbeleer/oh-my-posh/src/properties"

	"github.com/stretchr/testify/assert"
)

func TestGodotSegmentEnabled(t *testing.T){
	cases := []struct {
		Case			string
		ExpectedOutput 	bool
	}

	assert.True(false)
}

func TestGodotSegmentContent(t *testing.T) {
	cases := []struct {
		Case                string
		ExpectedOutput      string
		VersionFileText     string
		ExpectedToBeEnabled bool
		VersionFileExists   bool
	} {

	}

	for _, tc := range cases {
		env := new(mock.MockedEnvironment)
		env.On("Error", mock2.Anything).Return()
		env.On("Debug", mock2.Anything)


	}

	
	assert.True(false)
}