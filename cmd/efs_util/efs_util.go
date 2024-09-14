package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Snshadow/ntfs-efs"
	"github.com/Snshadow/ntfs-efs/cmd/utils"
)

func main() {
	writeRaw := flag.Bool("write-raw", false, "convert encrypted file to raw stream then write")
	toEncrypted := flag.Bool("to-encrypted", false, "read from stream of raw data then write encrypted file or directory")
	src := flag.String("src", "", "source file of data to read")
	target := flag.String("target", "", "target file to write data to")
	efsDir := flag.Bool("efs-dir", false, "process encrypted directory(and its files)")
	useStdin := flag.Bool("stdin", false, "use stardard input as source")
	useStdout := flag.Bool("stdout", false, "write to standard output")

	progName := filepath.Base(os.Args[0])

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s utilizes NTFS(New Technology File System) EFS(Encrypted File System) which some features.\nUsage:\n    copy encrypted file as raw data:\n\tas file: %s -write-raw -source [source file] -target [target file] or %s -write-raw [source file] [target file]\n\tto stdout: %s -write-raw -stdout -source [source file]\n    write encrypted file from raw data:\n\tfrom file: %s -to-encrypted -src [source file] -target [target file] or %s -to-encrypted [source file] [target file]\n\tto stdout: %s -to-encrypted -src [source file] or %s -to-encrypted [source file]\n\n", progName, progName, progName, progName, progName, progName, progName, progName)
		flag.PrintDefaults()

		// prevent window from closing immediately if the console was created for this process
		if utils.IsFromOwnConsole() {
			fmt.Fprintln(flag.CommandLine.Output(), "\nPress enter to close...")
			fmt.Scanln()
		}
	}

	flag.Parse()

	if !*toEncrypted && !*writeRaw {
		flag.Usage()
		os.Exit(1)
	}

	if *toEncrypted || *writeRaw {
		if *src == "" && (!*useStdin && *target == "") || *target == "" && (!*useStdout && *src == "") {
			flag.Usage()
			os.Exit(1)
		}

		rw, err := ntfs_efs.NewRawReadWriter()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create EFSRawReadWriter: %v\n", err)
			os.Exit(2)
		}

		var strm io.ReadWriter

		if *writeRaw {
			if *useStdout {
				strm = os.Stdout
			} else {
				strm, err = os.OpenFile(*target, os.O_RDWR|os.O_CREATE, 0777)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to open target file for write: %v\n", err)
					os.Exit(1)
				}
				defer strm.(*os.File).Close()
			}

			err = rw.ReadRaw(*src, strm)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write raw data from encrypted file: %v\n", err)
				os.Exit(1)
			}

			if !*useStdout {
				fmt.Printf("Wrote raw data from %s to %s\n", *src, *target)
			}
		} else if *toEncrypted {
			if *useStdin {
				strm = os.Stdin
			} else {
				strm, err = os.Open(*src)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to open encrypted file for reading: %v\n", err)
					os.Exit(1)
				}
				defer strm.(*os.File).Close()
			}

			err = rw.WriteRaw(*target, strm, *efsDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create encrypted file from raw data: %v\n", err)
				os.Exit(1)
			}

			if *useStdin {
				fmt.Printf("Wrote encrypted file using stdin to %s\n", *target)
			} else {
				fmt.Printf("Wrote encrypted file using %s to %s\n", *src, *target)
			}
		}

		return
	}
}
