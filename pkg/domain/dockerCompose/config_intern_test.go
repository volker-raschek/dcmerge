package dockerCompose

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPort_ExistsProtocol(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s              string
		expectedResult bool
	}{
		{
			s:              "0",
			expectedResult: false,
		},
		{
			s:              "53",
			expectedResult: false,
		},
		{
			s:              "53/tcp",
			expectedResult: true,
		},
		{
			s:              "53/udp",
			expectedResult: true,
		},
		{
			s:              "53:53",
			expectedResult: false,
		},
		{
			s:              "53:53/udp",
			expectedResult: true,
		},
		{
			s:              "53:53/tcp",
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		p := port(testCase.s)
		require.Equal(testCase.expectedResult, p.existsProtocol())
	}
}
