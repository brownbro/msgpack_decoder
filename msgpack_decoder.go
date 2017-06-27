package main

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ugorji/go/codec"
)

var (
	v     interface{}
	input []byte
	json  []byte
	mh    = codec.MsgpackHandle{RawToString: true}
	jh    codec.JsonHandle
)

// DecodeMessagePack decodes MessagePack(binary) data to JSON
func DecodeMessagePack(buf []byte) error {
	err := codec.NewDecoderBytes(buf, &mh).Decode(&v)
	if err != nil {
		return err
	}

	err = codec.NewEncoderBytes(&json, &jh).Encode(&v)
	if err != nil {
		return err
	}

	return err
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	err = DecodeMessagePack(input)
	if err != nil {
		panic(err)
	}

	rep := regexp.MustCompile(`\.[0-9A-Za-z]+\z`)
	jsonfilename := rep.ReplaceAllString(os.Args[1], ".json")
	ioutil.WriteFile(jsonfilename, json, os.ModePerm)
}
