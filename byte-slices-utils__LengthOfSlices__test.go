package commons_test

import (
	"testing"

	"github.com/Usemtech/commons"

	"github.com/stretchr/testify/require"
)

func TestLengthOfByteSlices(t *testing.T) {
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
		expected := 0
		actual := commons.LengthOfByteSlices()

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("02: one empty slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{[]byte{}}

		// 2. execute test
		expected := 0
		actual := commons.LengthOfByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("03: many empty slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{[]byte{}, []byte{}, []byte{}, []byte{}}

		// 2. execute test
		expected := 0
		actual := commons.LengthOfByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})

	t.Run("04: one filled slice", func(t *testing.T) {
		// 1. prepare test-data
		slices := [][]byte{[]byte{byte(5), byte(5), byte(5)}}

		// 2. execute test
		expected := 3
		actual := commons.LengthOfByteSlices(slices...)

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
		expected := 11
		actual := commons.LengthOfByteSlices(slices...)

		// 3. check outcome
		require.Equal(t, expected, actual)
	})
}
