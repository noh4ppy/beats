package dns

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mkdns "github.com/miekg/dns"
)

var (
	_defaultIgnoreFilter = []string{
		".arpa",
	}
)

// loads the ignore filter from the config.
func init() {
	fp, err := os.OpenFile("white.conf", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		_defaultIgnoreFilter = append(_defaultIgnoreFilter, strings.TrimSpace(line))
	}

	for _, v := range _defaultIgnoreFilter {
		fmt.Printf("load %s\n", v)
	}
}

func filter(name string, qtype uint16) (bool, error) {

	for _, ignore := range _defaultIgnoreFilter {
		if strings.HasSuffix(trimRightDot(name), ignore) {
			return true, nil
		}
	}

	switch qtype {
	case mkdns.TypeA:
		return false, nil
	case mkdns.TypeAAAA:
		return false, nil
	case mkdns.TypeCNAME:
		return false, nil
	default:
		return true, nil
	}
}
