package domain

import (
	"bytes"
	"encoding/json"
)

type AccountKind int

const (
	AccountKindUndefined = iota
	AccountKindPerson
	AccountKindShelter
)

var kindToString = map[AccountKind]string{
	AccountKindUndefined: "undefined",
	AccountKindPerson:    "person",
	AccountKindShelter:   "shelter",
}

var stringToKind = make(map[string]AccountKind)

func init() {
	for kind, str := range kindToString {
		stringToKind[str] = kind
	}
}

func (kind AccountKind) String() string {
	return kindToString[kind]
}

func AccountKindFromString(s string) AccountKind {
	return stringToKind[s]
}

func (kind AccountKind) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(kind.String())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (kind *AccountKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*kind = stringToKind[s]
	return nil
}