package stringToNumber

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
)

type MyInt64 int64
type MyInt int

func (myInt *MyInt) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case float64:
		*myInt = MyInt(v)
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*myInt = MyInt(n)
	}
	return nil
}

func (myInt64 *MyInt64) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case float64:
		*myInt64 = MyInt64(v)
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*myInt64 = MyInt64(n)
	}
	return nil
}

func (myInt *MyInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var item string
	d.DecodeElement(&item, &start)

	if item[0] == '"' && item[len(item)-1] == '"' {
		item = item[1 : len(item)-1]
	}

	n, err := strconv.Atoi(item)
	if err != nil {
		return err
	}

	*myInt = MyInt(n)
	return nil
}

func (myInt64 *MyInt64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var item string
	d.DecodeElement(&item, &start)

	if item[0] == '"' && item[len(item)-1] == '"' {
		item = item[1 : len(item)-1]
	}

	n, err := strconv.ParseInt(item, 0, 64)
	if err != nil {
		return err
	}

	*myInt64 = MyInt64(n)
	return nil
}

type Users struct {
	Users []User `json:"User" xml:"User"`
}

type User struct {
	ID      MyInt64 `json:"id" xml:"id"`
	Address Address `json:"address" xml:"address"`
	Age     MyInt   `json:"age" xml:"age"`
}

type Address struct {
	CityID MyInt64 `json:"city_id" xml:"city_id"`
	Street string  `json:"street" xml:"street"`
}
