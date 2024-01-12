package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"goapi/pkg/logger"
	"io"
	"io/ioutil"

	"github.com/forgoer/openssl"
)

func AesCBCEncrypt(src, key, iv []byte, padding string) ([]byte, error) {
	data, err := openssl.AesCBCEncrypt(src, key, iv, padding)
	if err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(data)), nil
}

func AesCBCDecrypt(src, key, iv []byte, padding string) ([]byte, error) {
	data, err := hex.DecodeString(string(src))
	if err != nil {
		return nil, err
	}
	return openssl.AesCBCDecrypt(data, key, iv, padding)
}

func Md5(text string) string {
	hashMd5 := md5.New()
	_, err := io.WriteString(hashMd5, text)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Zip(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&b, 9)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func UnZip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer func(r *gzip.Reader) {
		err = r.Close()
		if err != nil {
			logger.Error(err)
		}
	}(r)
	unzipData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return unzipData, nil
}

func Password(pwd, pwdCode string) string {
	return Md5(pwd + pwdCode)
}
