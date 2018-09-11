package commons_test

import (
	"Usemtech/commons.git"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcatByteSlices(t *testing.T) {
	/***************************************************************************
	CASES XX
	(01) no argument
	(02) one empty slice
	(03) many empty slices
	(04) one filled slice
	(05) many filled slices
	***************************************************************************/
	t.Run("01: no argument", func(t *testing.T) {
		// 1. prepare test-data

		// 2. execute test
		expected := []byte{}
		actual := commons.ConcatByteSlices()

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("02: one empty slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{[]byte{}}

		// 2. execute test
		expected := []byte{}
		actual := commons.ConcatByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("03: many empty slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{[]byte{}, []byte{}, []byte{}, []byte{}}

		// 2. execute test
		expected := []byte{}
		actual := commons.ConcatByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("04: one filled slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{[]byte{byte(5), byte(5), byte(5)}}

		// 2. execute test
		expected := []byte{byte(5), byte(5), byte(5)}
		actual := commons.ConcatByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("05: many filled slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{
			[]byte{byte(5), byte(5), byte(5)},
			[]byte{byte(5)},
			[]byte{byte(5), byte(5), byte(5), byte(5), byte(5)},
			[]byte{byte(5), byte(5)}}

		// 2. execute test
		expected := []byte{
			byte(5), byte(5), byte(5),
			byte(5), byte(5), byte(5),
			byte(5), byte(5), byte(5),
			byte(5), byte(5)}
		actual := commons.ConcatByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})
}
