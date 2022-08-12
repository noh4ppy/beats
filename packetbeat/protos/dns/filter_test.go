package dns

import (
	"testing"

	mkdns "github.com/miekg/dns"
)

func TestFilter(t *testing.T) {
	ok, err := filter("1111.in-addr.arpa", mkdns.TypeA)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("expected true")
	}

	ok2, err := filter("xxxxx.baidu.com", mkdns.TypeTXT)
	if err != nil {
		t.Error(err)
	}
	if !ok2 {
		t.Error("expected true")
	}

	ok3, err := filter("dlp.keenon.com", mkdns.TypeHTTPS)
	if err != nil {
		t.Error(err)
	}
	if ok3 {
		t.Error("expected flase")
	}
}
