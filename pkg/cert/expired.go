package cert

import (
	"crypto/tls"
	"fmt"
	"time"
)

func Expired(address string, expiryThreshold time.Duration, verbose bool) (
	expired bool, err error,
) {
	conn, err := tls.Dial("tcp", address, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return false, err
	}
	defer func(conn *tls.Conn) {
		if tempErr := conn.Close(); tempErr != nil && err == nil {
			err = tempErr
		}
	}(conn)

	cert := conn.ConnectionState().PeerCertificates[0]
	expiryDuration := cert.NotAfter.Sub(time.Now())
	if verbose {
		fmt.Printf("%s %s (%s)\n", address, cert.NotAfter, expiryDuration)
	}
	return expiryDuration < expiryThreshold, nil
}
