package HexTool

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

func ZlibDecompress(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func ZlibCompress(data []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}
