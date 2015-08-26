package canonicaltest

import (
	"bytes"
	"encoding/json"
	"testing"

	cbor "github.com/whyrusleeping/cbor/go"
)

func RoundTripTest(t *testing.T, o1, o2 interface{}) bool {
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer
	var buf3 bytes.Buffer
	enc1 := cbor.NewEncoder(&buf1)
	enc2 := cbor.NewEncoder(&buf2)
	dec := cbor.NewDecoder(&buf3)

	if err := enc1.Encode(o1); err != nil {
		t.Error(err)
		return false
	}

	m1 := buf1.Bytes()
	if _, err := buf3.Write(m1); err != nil {
		t.Error(err)
		return false
	}

	if err := dec.Decode(o2); err != nil {
		t.Error(err)
		return false
	}

	if err := enc2.Encode(o2); err != nil {
		t.Error(err)
		return false
	}

	m2 := buf2.Bytes()
	if !bytes.Equal(m1, m2) {
		t.Error("marshalled values not equal")
		t.Log(m1)
		t.Log(m2)
		t.Log(string(m1))
		t.Log(string(m2))
		return false
	}

	return true
}

func OutputJSON(t *testing.T, thing interface{}) {
	b, err := json.Marshal(thing)
	if err != nil {
		t.Error(err)
		return
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, b, "", "\t"); err != nil {
		t.Error(err)
		return
	}

	t.Log(buf.String())
}
