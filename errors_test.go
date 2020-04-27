package errors

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Format(t *testing.T) {
	label := errors.New("label-error")
	err := errors.New("error")

	var wraped error
	wraped = Wrap(err, label)
	require.Error(t, wraped)
	require.Equal(t, "label-error: error", wraped.Error())

	// check format
	require.Equal(t, `err: label-error: error`,
		fmt.Sprintf("err: %v", wraped))
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	require.Equal(t, fmt.Sprintf(`err: label-error: error:
    github.com/tico8/lerrors.Test_Format
        %s:17`, filename),
		fmt.Sprintf("err: %+v", wraped))
}

func Test_Wrap(t *testing.T) {
	label := errors.New("label-error")
	err1 := errors.New("error1")
	err2 := errors.New("error2")

	var wraped1 error
	wraped1 = Wrap(err1, label)
	require.Error(t, wraped1)
	require.Equal(t, "label-error: error1", wraped1.Error())

	var wraped2 error
	wraped2 = Wrap(err2, label)
	require.Error(t, wraped2)
	require.Equal(t, "label-error: error2", wraped2.Error())

	require.True(t, errors.Is(wraped1, label))
	require.True(t, errors.Is(wraped1, err1))
	require.False(t, errors.Is(wraped1, err2))
	require.True(t, errors.Is(wraped2, label))
	require.False(t, errors.Is(wraped2, err1))
	require.True(t, errors.Is(wraped2, err2))
	require.True(t, errors.Is(wraped1, wraped2))
	require.True(t, errors.Is(wraped2, wraped1))
}
