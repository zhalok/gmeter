package helpers

import (
	"fmt"
	"net/url"
	"time"
)

func ExtractProtocl(testUrl string)(string,error) {

	u, err := url.Parse(testUrl)
	if err != nil {
		
		return "", fmt.Errorf("error parsing url: %w", err)
	}
	protocol := u.Scheme

	return protocol, nil
}

func ExtractHostname(testUrl string) (string, error) {
	u, err := url.Parse(testUrl)
	if err != nil {
		return "", fmt.Errorf("error parsing url: %w", err)
	}
	hostname := u.Hostname()
	return hostname, nil
}

func GetFormattedTimeStampString() string {
	t := time.Now()
	ts := fmt.Sprintf("%02d.%02d.%04d.%02d.%02d.%02d.%03d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1e6)

	return ts
}