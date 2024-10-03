package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
)

func main() {
	in := os.Args[1]

	mh, err := tryParseAsMH(in)
	if err != nil {
		fmt.Printf("not a base64 raw mh: %s\n", err)
	} else {
		c := cid.NewCidV1(uint64(multicodec.Raw), mh)
		fmt.Printf("canonical form: %s\n", c)
		return
	}
	mh, err = tryParseAsB58MH(in)
	if err != nil {
		fmt.Printf("not a formatted mh: %s\n", err)
	} else {
		c := cid.NewCidV1(uint64(multicodec.Raw), mh)
		fmt.Printf("canonical form: %s\n", c)
		return
	}
}

func tryParseAsMH(s string) (multihash.Multihash, error) {
	r, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	_, err = multihash.Decode(r)
	if err != nil {
		return nil, err
	}
	return multihash.Multihash(r), nil
}

func tryParseAsB58MH(s string) (multihash.Multihash, error) {
	mh, err := multihash.FromB58String(s)
	if err != nil {
		return nil, err
	}
	return mh, nil
}
