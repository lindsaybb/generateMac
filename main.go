package main

import (
	"crypto/rand"
	"fmt"
	"flag"
	"strings"
)

var (
	helpFlag = flag.Bool("h", false, "Show this help")
	knownPrefix = flag.String("p", "28:99:c7", "Specify the OUI (first six hex) with colon (:) or dash (-) separators")
	numFlag = flag.Int("n", 1, "Number of entries to generate")
)

const usage = "`macGenerator` [options]"

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}
	var prefix []string
	if *knownPrefix != "" {
		prefix = strings.Split(*knownPrefix, ":")
		if len(prefix) < 3 {
			prefix = strings.Split(*knownPrefix, "-")
			if len(prefix) < 3 {
				fmt.Printf("Unknown OUI supplied: %s\n", *knownPrefix)
				return
			}
		}	
	} else {
		prefix = []string{"", "", ""}
	}
	

	switch {
	case *numFlag < 1:
		return
	case *numFlag > 1000:
		return
	default:
		generateMac(prefix, *numFlag)
	}
}

// controls iterations and prints directly to stdout
func generateMac(p []string, n int) {
	var buf []byte
	var hasPrefix bool
	if p[0] != "" {
		hasPrefix = true
	}

	for i := 0; i < n; i++ {
		buf = randomSixHex()
		if hasPrefix {
			fmt.Printf("%s:%s:%s:%02x:%02x:%02x\n", p[0], p[1], p[2], buf[3], buf[4], buf[5])
		} else {
			fmt.Printf("%02x:%02x:%02x:%02x:%02x:%02x\n", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])		
		}
	}
}

func randomSixHex() []byte {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	buf[0] |= 2
	return buf
}
