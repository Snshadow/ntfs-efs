package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/Snshadow/ntfs-efs"
)

func main() {
	writeRaw := flag.Bool("write-raw", false, "convert encrypted file to raw stream then write")
	readRaw := flag.Bool("read-raw", false, "read from stream of raw data then write encrypted file or directory")
	src := flag.String("source", "", "source file of data to read")
	target := flag.String("target", "", "target file to write data to")

	rw, err := ntfs_efs.NewRawReadWriter()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create EFSRawReadWriter: %v\n", err)
		os.Exit(2)
	}
}
