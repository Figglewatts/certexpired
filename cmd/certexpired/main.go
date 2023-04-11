package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/figglewatts/certexpired/pkg/cert"
)

const (
	ExitCodeOk           = 0
	ExitCodeWithinExpiry = 1
	ExitCodeUsage        = 2
	ExitCodeError        = 3

	UsageText = "invalid usage: send piped input, or run: certexpired [options] ADDRESS..."
)

var (
	expiryThreshold = flag.Duration(
		"threshold", 24*time.Hour*30, "certificate expiry threshold",
	)
	verbose = flag.Bool("verbose", false, "verbose output")
)

func printStdErr(val interface{}) {
	_, err := fmt.Fprintln(os.Stderr, val)
	if err != nil {
		panic(err)
	}
}

func argsMain(addresses []string) int {
	var expiredAddresses []string
	for _, address := range addresses {
		expired, err := cert.Expired(address, *expiryThreshold, *verbose)
		if err != nil {
			printStdErr(
				fmt.Errorf(
					"checking certificate at %s: %w", address, err,
				),
			)
			return ExitCodeError
		}

		if expired {
			expiredAddresses = append(expiredAddresses, address)
		}
	}

	if len(expiredAddresses) == 0 {
		return ExitCodeOk
	} else {
		for _, address := range expiredAddresses {
			fmt.Println(address)
		}
		return ExitCodeWithinExpiry
	}
}

func pipedMain(inputReader io.Reader) int {
	var addresses []string
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		addresses = append(addresses, line)
	}

	if err := scanner.Err(); err != nil {
		printStdErr(fmt.Errorf("reading input: %w", err))
		return ExitCodeError
	}

	return argsMain(addresses)
}

func isPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func main() {
	flag.Parse()

	exitCode := ExitCodeOk
	if len(flag.Args()) > 0 {
		exitCode = argsMain(flag.Args())
	} else if isPipe() {
		// input has been piped into stdin
		exitCode = pipedMain(bufio.NewReader(os.Stdin))
	} else {
		exitCode = ExitCodeUsage
		printStdErr(UsageText)
	}
	os.Exit(exitCode)
}
