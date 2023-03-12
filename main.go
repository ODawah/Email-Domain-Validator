package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/mail"
	"os"
	"strings"
)

func DomainExtractor(email string) (string, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}
	domain := strings.Split(email, "@")
	return domain[1], nil
}

func main() {
	fmt.Println(DomainExtractor("omar@gmail.com"))
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")
	for scanner.Scan() {
		domain, err := DomainExtractor(scanner.Text())
		if err == nil {
			checkDomain(domain)
		}
		continue
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error couldn't read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v,",
		domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
