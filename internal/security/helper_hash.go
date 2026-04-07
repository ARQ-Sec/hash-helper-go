package security

import (
    "crypto/md5"
    "hash"
)

func legacyFactory() hashWrapper {
    return hashWrapper{newFunc: md5.New}
}

type hashWrapper struct {
    newFunc func() hash.Hash
}

func Digest(input []byte) []byte {
    digest := legacyFactory().newFunc()
    digest.Write(input)
    return digest.Sum(nil)
}
