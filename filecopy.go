package main

import (
	"flag"
	"io"
	"log"
	"os"
)

// Copy a file
func main() {

	ifname := flag.String("if", "", "Dateiname der Quelldatei mit Pfadangabe")
	ofname := flag.String("of", "", "Dateiname der Zieldatei mit Pfadangabe")
	flag.Parse()

	if (*ifname != "") && (*ofname != "") {
		// Open original file
		originalFile, err := os.Open(*ifname)
		if err != nil {
			log.Fatal(err)
		}
		defer originalFile.Close()

		// Create new file
		newFile, err := os.Create(*ofname)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		// Copy the bytes to destination from source
		bytesWritten, err := io.Copy(newFile, originalFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Copied %d bytes.", bytesWritten)

		// Commit the file contents
		// Flushes memory to disk
		err = newFile.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}
}
