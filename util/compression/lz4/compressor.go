package lz4

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/pierrec/lz4/v4"
)

const AlgoName = "lz4"

type Compressor struct{}

func (c Compressor) CompressData(data []byte) (io.ReadSeeker, error) {
	var b bytes.Buffer
	w := lz4.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		w.Close()
		return nil, err
	}
	w.Close()
	return bytes.NewReader(b.Bytes()), nil
}

func (c Compressor) DecompressData(src io.Reader) ([]byte, error) {
	r := lz4.NewReader(src)

	block, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return block, nil
}
