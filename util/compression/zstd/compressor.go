package zstd

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/klauspost/compress/zstd"
)

const AlgoName = "zstd"

type Compressor struct{}

func (c Compressor) CompressData(data []byte) (io.ReadSeeker, error) {
	var b bytes.Buffer
	w, err := zstd.NewWriter(&b)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(data); err != nil {
		w.Close()
		return nil, err
	}
	w.Close()
	return bytes.NewReader(b.Bytes()), nil
}

func (c Compressor) DecompressData(src io.Reader) ([]byte, error) {
	r, err := zstd.NewReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	block, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return block, nil
}
