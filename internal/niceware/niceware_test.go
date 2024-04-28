package niceware_test

import (
	"bytes"
	"slices"
	"testing"

	"github.com/Tethik/go-template/internal/niceware"
)

func TestParse(t *testing.T) {
	b, err := niceware.WordsToBytes([]string{"A"})
	if err != nil {
		t.Errorf("Error parsing empty byte slice")
	}
	if !bytes.Equal(b, []byte{0, 0}) {
		t.Errorf("Wrong byte slice")
	}

	b, err = niceware.WordsToBytes([]string{"zyzzyva"})
	if err != nil {
		t.Errorf("Error parsing empty byte slice")
	}
	if !bytes.Equal(b, []byte{0xff, 0xff}) {
		t.Errorf("Wrong byte slice")
	}

	b, err = niceware.WordsToBytes([]string{"A", "billet", "baiting", "glum", "Crawl", "writhing", "dePlanE", "zyzzYvA"})
	if err != nil {
		t.Errorf("Error parsing empty byte slice")
	}
	if !bytes.Equal(b, []byte{0, 0, 17, 212, 12, 140, 90, 247, 46, 83, 254, 60, 54, 169, 255, 255}) {
		t.Errorf("Wrong byte slice")
	}
}

func TestStringify(t *testing.T) {
	b, err := niceware.BytesToWords([]byte{0, 0})
	if err != nil {
		t.Errorf("Error parsing empty byte slice")
	}
	if !slices.Equal(b, []string{"a"}) {
		t.Errorf("Wrong word slice %s vs %s", b, "a")
	}

	b, err = niceware.BytesToWords([]byte{0xff, 0xff})
	if err != nil {
		t.Errorf("Error parsing empty byte slice")
	}
	if !slices.Equal(b, []string{"zyzzyva"}) {
		t.Errorf("Wrong word slice %s vs %s", b, "zyzzyva")
	}

	b, err = niceware.BytesToWords([]byte{0, 0, 17, 212, 12, 140, 90, 247, 46, 83, 254, 60, 54, 169, 255, 255})
	if err != nil {
		t.Errorf("Error parsing empty byte slice")
	}
	if !slices.Equal(b, []string{"a", "billet", "baiting", "glum", "crawl", "writhing", "deplane", "zyzzyva"}) {
		t.Errorf("Wrong word slice %s vs %s", b, "a billet baiting glum crawl writhing deplane zyzzyva")
	}
}

func TestShouldErrorOnNonEven(t *testing.T) {
	_, err := niceware.BytesToWords([]byte("hello"))
	if err == nil {
		t.Errorf("Expected error on odd length input")
	}
}

func TestBinarySearch(t *testing.T) {
	count := 0
	for i, word := range niceware.English {
		if niceware.BinarySearch(word) != i {
			count += 1
			t.Errorf("Binary search failed for word %s", word)
		}
	}
	if count > 0 {
		t.Errorf("Binary search failed for %d out of %d words", count, len(niceware.English))
	}
}
