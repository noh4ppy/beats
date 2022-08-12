package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	data := `eJzMl0tv3LYTwO/7KQZ7+Sf410pToJc9FEhrFwhQGEGc9tDLhiJHK9YUR50h11Y+fUFK2pW8j7iP2PHFC2o4/M2Twwu4xW4FxssCINjgcAXLy+ub5QLAoGi2bbDkV3B5fXMhLWpbWQ24RR+gsuiMFAsYfq0WAAAX4FWDo8r0F7oWV7Bhiu2wMpWf7qmc2kihYqiJbVDBbnEnM+opiRwqP1mfcf4w+QDwJnFntdDDd9ZvINQqQKgRGKUlb9KaIG+RwQooDyNBBxXxTGHaZahR1mdkiIIGrM/rf0aUBFEsThjGqCOLJb9WW2WdKt0XMe+uxlAjw3DcNpNxBxLbljhkG8fzB/aZ1mxY749HmGJQLKP5onHSzqaEM5ZRB8lLQ7wCQRtZIoKaKewt3nnAdQW8P+0OylDKnTQ3JQT6YPXaqKC+cE6OmIOJmrxYgyyTlJV5xHZ4Jw3QNepb6zdrYyXF/Ynilc+aBmymRezGqxAZYaucNSqdBFSNxcTdSXMCR69VQLM+4o//1h7yrstElWUJ8P3r76DsAsoIyti6Du6Q5yFhDJE9miMm7PoEBmfWrYuyJo+nQD/UCFhVqFM3hEDthcMturELvcAPv1y+hKQFyCM0xAhOleiKGc/PxID3qmkdfpO5077/v04NDpYVUVEqLjbklN8UxJtimepiOV2Y60tYqRayAoMBubG+d1qvGzQ1KFAxNamlYlOiMWhAU9sNvpsp7JUl6TqEdvXqVRtLZ7XEqrL3mWAiPliyAtWoT+QLTUW8PeJq5eUOWdaaog8HKeLIbx6XH8laH5sSOaEzCkXWuVaJjaQSDcr6/VXw0XgphsM/9tddsTiCN9wzFuUAjso/UIdHpq8Hxay6ESTFQYGxOrc1ztcYoNL15GYTzF97lz8MRo9+zKF74mdy6h5g59iHSk8IQqM6IM7/PAUo0+2nXUxZabDFfhCgwwtRk6/sJvKuQb1T+hZDiSqcd1GRVs4V9nSUCAR3tdU1hNrKQ3dAi5y8IUerYPhRaGo+A5RidA4ofU8W5mp8EIFDqqMw1zfnEbRTIucYskCCSD35X4C8vT4CYozt7/onq7jdif+05PbIz1Rye4DPlNyB4FOV3P7gr6TkJkDPVXIThK+l5IoQ3OMTNdgmpUtA3iqXMAQ1eSPjnHkkZCnPSgStdD0bcQFKrNJkZANITdGZJGasaMUGTQG/I1OaQiMKNKj8fpbN3uj3zF9JfTr3RxXHa/KsLx68JD7jjIzRy5TjmDVaX2SJhJ2rqep9M+xIRSYzdX259XmnvNkHf6pzF+UDG6gNxRbz4+9cSl2lVBrkjubK8tvlceWG/u4w/7YCwdDPtYGVF9U32igoKaVvrn46YQjehzVrMicr9Oo+oE/tanxoQJI+bIN7s358c/nb1fubE8ZF067Ffjp8r5xr4u8xvxmI/yfw6+U7aFXnSBlIiuCF9f2L5GWxWPwVAAD//ysvXY8=`
	bs, err := decodeData(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

func decodeData(data string) ([]byte, error) {

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(decoded)
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}
