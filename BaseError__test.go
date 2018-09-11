package commons_test

import (
	"Usemtech/commons.git"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBaseError(t *testing.T) {
	/***************************************************************************
	CASES XX
	(01) NewError
	(02) NewDecline
	(03) NewFailure
	(04) NewFatal
	(05) Stringer
	***************************************************************************/
	t.Run("01: NewError", func(t *testing.T) {
		// 1. prepare test-data
		cat := commons.ErrDecline
		rsn := "some reason"
		msg := "some message"

		// 2. execute test
		actual := commons.NewError(cat, rsn, msg)

		// 3. check outcome
		require.Equal(t, cat, actual.Cat)
		require.Equal(t, rsn, actual.Rsn)
		require.Equal(t, msg, actual.Msg)
		require.NotNil(t, actual.File)
		require.NotNil(t, actual.Func)
		require.NotZero(t, actual.Line)
	})

	t.Run("02: NewDecline", func(t *testing.T) {
		// 1. prepare test-data
		cat := commons.ErrDecline
		rsn := "some reason"
		msg := "some message"

		// 2. execute test
		actual := commons.NewDecline(rsn, msg)

		// 3. check outcome
		require.Equal(t, cat, actual.Cat)
		require.Equal(t, rsn, actual.Rsn)
		require.Equal(t, msg, actual.Msg)
		require.NotNil(t, actual.File)
		require.NotNil(t, actual.Func)
		require.NotZero(t, actual.Line)
	})

	t.Run("03: NewFailure", func(t *testing.T) {
		// 1. prepare test-data
		cat := commons.ErrFailure
		rsn := "some reason"
		msg := "some message"

		// 2. execute test
		actual := commons.NewFailure(rsn, msg)

		// 3. check outcome
		require.Equal(t, cat, actual.Cat)
		require.Equal(t, rsn, actual.Rsn)
		require.Equal(t, msg, actual.Msg)
		require.NotNil(t, actual.File)
		require.NotNil(t, actual.Func)
		require.NotZero(t, actual.Line)
	})

	t.Run("04: NewFatal", func(t *testing.T) {
		// 1. prepare test-data
		cat := commons.ErrFatal
		rsn := "some reason"
		msg := "some message"

		// 2. execute test
		actual := commons.NewFatal(rsn, msg)

		// 3. check outcome
		require.Equal(t, cat, actual.Cat)
		require.Equal(t, rsn, actual.Rsn)
		require.Equal(t, msg, actual.Msg)
		require.NotNil(t, actual.File)
		require.NotNil(t, actual.Func)
		require.NotZero(t, actual.Line)
	})

	t.Run("05: Stringer", func(t *testing.T) {
		// 1. prepare test-data
		cat := commons.ErrDecline
		rsn := "some reason"
		msg := "some message"
		err := commons.NewError(cat, rsn, msg)

		// 2. execute test
		expected := fmt.Sprintf("^%s\\: %s -- %s \\(.+\\.go\\:\\d+\\)$", cat, rsn, msg)
		actual := fmt.Sprint(err)

		// 3. check outcome
		require.Regexp(t, expected, actual)
	})
}
