package stringToNumber

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestStringInt_UnmarshalJSON(t *testing.T) {
	rawJson := []byte(`[
	  {
		"id": 1,
		"address": {
		  "city_id": 5,
		  "street": "Satbayev"
		},
		"Age": 20
	  },
	  {
		"id": 1,
		"address": {
		  "city_id": 6,
		  "street": "Al-Farabi"
		},
		"Age": "32"
	  }
	]`)
	expectedUsers := []User{
		{
			ID: MyInt64(1),
			Address: Address{
				CityID: MyInt64(5),
				Street: "Satbayev",
			},
			Age: MyInt(20),
		},
		{
			ID: MyInt64(1),
			Address: Address{
				CityID: MyInt64(6),
				Street: "Al-Farabi",
			},
			Age: MyInt(32),
		},
	}

	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		t.Errorf("Error occured during Unmarshaling JSON")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Incorrect result, expected:\n %#v\n got:\n %#v", expectedUsers, users)
	}
}

func TestStringInt64_UnmarshalJSON(t *testing.T) {
	rawJson := []byte(`[
	  {
		"id": "1",
		"address": {
		  "city_id": 5,
		  "street": "Satbayev"
		},
		"Age": 20
	  },
	  {
		"id": 1,
		"address": {
		  "city_id": "6",
		  "street": "Al-Farabi"
		},
		"Age": 32
	  }
	]`)
	expectedUsers := []User{
		{
			ID: MyInt64(1),
			Address: Address{
				CityID: MyInt64(5),
				Street: "Satbayev",
			},
			Age: MyInt(20),
		},
		{
			ID: MyInt64(1),
			Address: Address{
				CityID: MyInt64(6),
				Street: "Al-Farabi",
			},
			Age: MyInt(32),
		},
	}

	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		t.Errorf("Error occured during Unmarshaling JSON")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Incorrect result, expected:\n %#v\n got:\n %#v", expectedUsers, users)
	}
}

func TestStringInt_UnmarshalXML(t *testing.T) {
	var rawXML = []byte(`<Users>
		<User>
		  <id>0</id>
		  <address>
			 <city_id>5</city_id>
			 <street>Satbayev</street>
		  </address>
		  <age>20</age>
	   </User>
	   <User>
		  <id>1</id>
		  <address>
			 <city_id>6</city_id>
			 <street>Al-Farabi</street>
		  </address>
		  <age>"32"</age>
	   </User>
	</Users>`)
	expectedUsers := Users{
		Users: []User{
			{
				ID: MyInt64(0),
				Address: Address{
					CityID: MyInt64(5),
					Street: "Satbayev",
				},
				Age: MyInt(20),
			},
			{
				ID: MyInt64(1),
				Address: Address{
					CityID: MyInt64(6),
					Street: "Al-Farabi",
				},
				Age: MyInt(32),
			},
		},
	}
	var users Users
	if err := xml.Unmarshal(rawXML, &users); err != nil {
		t.Errorf("Error occured during Unmarshaling JSON")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Incorrect result, expected:\n %#v\n got:\n %#v", expectedUsers, users)
	}
}

func TestStringInt64_UnmarshalXML(t *testing.T) {
	var rawXML = []byte(`<Users>
		<User>
		  <id>"0"</id>
		  <address>
			 <city_id>5</city_id>
			 <street>Satbayev</street>
		  </address>
		  <age>20</age>
	   </User>
	   <User>
		  <id>1</id>
		  <address>
			 <city_id>"6"</city_id>
			 <street>Al-Farabi</street>
		  </address>
		  <age>32</age>
	   </User>
	</Users>`)
	expectedUsers := Users{
		Users: []User{
			{
				ID: MyInt64(0),
				Address: Address{
					CityID: MyInt64(5),
					Street: "Satbayev",
				},
				Age: MyInt(20),
			},
			{
				ID: MyInt64(1),
				Address: Address{
					CityID: MyInt64(6),
					Street: "Al-Farabi",
				},
				Age: MyInt(32),
			},
		},
	}
	var users Users
	if err := xml.Unmarshal(rawXML, &users); err != nil {
		t.Errorf("Error occured during Unmarshaling JSON")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("Incorrect result, expected:\n %#v\n got:\n %#v", expectedUsers, users)
	}
}