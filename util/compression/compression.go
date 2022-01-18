package compression

import (
	"io"

	"github.com/longhorn/backupstore/util/compression/gzip"
	"github.com/longhorn/backupstore/util/compression/lz4"
	"github.com/longhorn/backupstore/util/compression/zstd"
)

type Compressor interface {
	CompressData(data []byte) (io.ReadSeeker, error)
	DecompressData(src io.Reader) ([]byte, error)
}

var Compressors = map[string]Compressor{
	gzip.AlgoName: gzip.Compressor{},
	zstd.AlgoName: zstd.Compressor{},
	lz4.AlgoName:  lz4.Compressor{},
}
