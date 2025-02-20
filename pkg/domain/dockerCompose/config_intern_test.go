package dockerCompose

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitStringInPortMapping(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s                string
		expectedSrc      string
		expectedDst      string
		expectedProtocol string
	}{
		{
			s:                "53:53",
			expectedSrc:      "53",
			expectedDst:      "53",
			expectedProtocol: "",
		},
		{
			s:                "0.0.0.0:53:53",
			expectedSrc:      "0.0.0.0:53",
			expectedDst:      "53",
			expectedProtocol: "",
		},
		{
			s:                "0.0.0.0:53:10.11.12.13:53",
			expectedSrc:      "0.0.0.0:53",
			expectedDst:      "10.11.12.13:53",
			expectedProtocol: "",
		},
		{
			s:                "0.0.0.0:53:10.11.12.13:53/tcp",
			expectedSrc:      "0.0.0.0:53",
			expectedDst:      "10.11.12.13:53",
			expectedProtocol: "tcp",
		},
	}

	for i, testCase := range testCases {
		actualSrc, actualDst, actualProtocol := splitStringInPortMapping(testCase.s)
		require.Equal(testCase.expectedSrc, actualSrc, "TestCase %v", i)
		require.Equal(testCase.expectedDst, actualDst, "TestCase %v", i)
		require.Equal(testCase.expectedProtocol, actualProtocol, "TestCase %v", i)
	}
}

func TestPort_DstIP(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s              string
		expectedBool   bool
		expectedString string
	}{
		{
			s:              "",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53/tcp",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53/udp",
			expectedBool:   false,
			expectedString: "",
		},

		{
			s:              "0.0.0.0:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:0.0.0.0:53",
			expectedBool:   true,
			expectedString: "0.0.0.0",
		},
		{
			s:              "53:0.0.0.0:53/tcp",
			expectedBool:   true,
			expectedString: "0.0.0.0",
		},
		{
			s:              "53:0.0.0.0:53/udp",
			expectedBool:   true,
			expectedString: "0.0.0.0",
		},

		{
			s:              "10.11.12.13:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:10.11.12.13:53",
			expectedBool:   true,
			expectedString: "10.11.12.13",
		},
		{
			s:              "53:10.11.12.13:53/tcp",
			expectedBool:   true,
			expectedString: "10.11.12.13",
		},
		{
			s:              "53:10.11.12.13:53/udp",
			expectedBool:   true,
			expectedString: "10.11.12.13",
		},
	}

	for i, testCase := range testCases {
		p := port(testCase.s)
		require.Equal(testCase.expectedBool, p.existsDstIP(), "TestCase %v", i)
		require.Equal(testCase.expectedString, p.getDstIP(), "TestCase %v", i)
	}
}

func TestPort_DstPort(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s              string
		expectedBool   bool
		expectedString string
	}{
		{
			s:              "",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:53/tcp",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:53/udp",
			expectedBool:   true,
			expectedString: "53",
		},

		{
			s:              "53:0.0.0.0:53",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:0.0.0.0:53/tcp",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:0.0.0.0:53/udp",
			expectedBool:   true,
			expectedString: "53",
		},

		{
			s:              "53:10.11.12.13:53",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:10.11.12.13:53/tcp",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:10.11.12.13:53/udp",
			expectedBool:   true,
			expectedString: "53",
		},
	}

	for i, testCase := range testCases {
		p := port(testCase.s)
		require.Equal(testCase.expectedBool, p.existsDstPort(), "TestCase %v", i)
		require.Equal(testCase.expectedString, p.getDstPort(), "TestCase %v", i)
	}
}

func TestPort_Protocol(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s              string
		expectedBool   bool
		expectedString string
	}{
		{
			s:              "0",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53/tcp",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53/udp",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53/tcp",
			expectedBool:   true,
			expectedString: "tcp",
		},
		{
			s:              "53:53/udp",
			expectedBool:   true,
			expectedString: "udp",
		},
		{
			s:              "0.0.0.0:53:53/tcp",
			expectedBool:   true,
			expectedString: "tcp",
		},
		{
			s:              "0.0.0.0:53:53/udp",
			expectedBool:   true,
			expectedString: "udp",
		},
		{
			s:              "0.0.0.0:53:53/tcp",
			expectedBool:   true,
			expectedString: "tcp",
		},
		{
			s:              "0.0.0.0:53:11.12.13.14:53/tcp",
			expectedBool:   true,
			expectedString: "tcp",
		},
		{
			s:              "0.0.0.0:53:11.12.13.14:53/udp",
			expectedBool:   true,
			expectedString: "udp",
		},
	}

	for i, testCase := range testCases {
		p := port(testCase.s)
		require.Equal(testCase.expectedBool, p.existsProtocol(), "TestCase %v", i)
		require.Equal(testCase.expectedString, p.getProtocol(), "TestCase %v", i)
	}
}

func TestPort_SrcIP(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s              string
		expectedBool   bool
		expectedString string
	}{
		{
			s:              "",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53/tcp",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53/udp",
			expectedBool:   false,
			expectedString: "",
		},

		{
			s:              "0.0.0.0:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "0.0.0.0:53:53",
			expectedBool:   true,
			expectedString: "0.0.0.0",
		},
		{
			s:              "0.0.0.0:53:53/tcp",
			expectedBool:   true,
			expectedString: "0.0.0.0",
		},
		{
			s:              "0.0.0.0:53:53/udp",
			expectedBool:   true,
			expectedString: "0.0.0.0",
		},

		{
			s:              "10.11.12.13:53",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "10.11.12.13:53:53",
			expectedBool:   true,
			expectedString: "10.11.12.13",
		},
		{
			s:              "10.11.12.13:53:53/tcp",
			expectedBool:   true,
			expectedString: "10.11.12.13",
		},
		{
			s:              "10.11.12.13:53:53/udp",
			expectedBool:   true,
			expectedString: "10.11.12.13",
		},
	}

	for i, testCase := range testCases {
		p := port(testCase.s)
		require.Equal(testCase.expectedBool, p.existsSrcIP(), "TestCase %v", i)
		require.Equal(testCase.expectedString, p.getSrcIP(), "TestCase %v", i)
	}
}

func TestPort_SrcPort(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		s              string
		expectedBool   bool
		expectedString string
	}{
		{
			s:              "",
			expectedBool:   false,
			expectedString: "",
		},
		{
			s:              "53:53",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:53/tcp",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "53:53/udp",
			expectedBool:   true,
			expectedString: "53",
		},

		{
			s:              "0.0.0.0:53:53",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "0.0.0.0:53:53/tcp",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "0.0.0.0:53:53/udp",
			expectedBool:   true,
			expectedString: "53",
		},

		{
			s:              "10.11.12.13:53:53",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "10.11.12.13:53:53/tcp",
			expectedBool:   true,
			expectedString: "53",
		},
		{
			s:              "10.11.12.13:53:53/udp",
			expectedBool:   true,
			expectedString: "53",
		},
	}

	for i, testCase := range testCases {
		p := port(testCase.s)
		require.Equal(testCase.expectedBool, p.existsSrcPort(), "TestCase %v", i)
		require.Equal(testCase.expectedString, p.getSrcPort(), "TestCase %v", i)
	}
}
