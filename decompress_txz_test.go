package getter

import (
	"path/filepath"
	"testing"
)

func TestTarXzDecompressor(t *testing.T) {

	multiplePaths := []string{"dir/", "dir/test2", "test1"}

	cases := []TestDecompressCase{
		{
			"empty.tar.xz",
			false,
			true,
			nil,
			"",
			nil,
		},

		{
			"single.tar.xz",
			false,
			false,
			nil,
			"d3b07384d113edec49eaa6238ad5ff00",
			nil,
		},

		{
			"single.tar.xz",
			true,
			false,
			[]string{"file"},
			"",
			nil,
		},

		{
			"multiple.tar.xz",
			true,
			false,
			[]string{"file1", "file2"},
			"",
			nil,
		},

		{
			"multiple.tar.xz",
			false,
			true,
			nil,
			"",
			nil,
		},

		{
			"multiple_dir.tar.xz",
			true,
			false,
			multiplePaths,
			"",
			nil,
		},

	}

	for i, tc := range cases {
		cases[i].Input = filepath.Join("./testdata", "decompress-txz", tc.Input)
	}

	TestDecompressor(t, new(TarXzDecompressor), cases)
}
