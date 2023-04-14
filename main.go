package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter domain name")

	domain := bufio.NewScanner(os.Stdin)
	domain.Scan()

	record, err := net.LookupMX(domain.Text())
	if err != nil {
		panic(err)

	}

	fmt.Printf("Hostname: %s and Preference number: %d\n", record[0].Host, record[0].Pref)

	spf, err := net.LookupTXT(domain.Text())
	if err != nil {
		panic(err)
	}

	for _, spfr := range spf {
		if strings.HasPrefix(spfr, "v=DMARC1") {
			fmt.Println(spfr)
		}

	}

}
