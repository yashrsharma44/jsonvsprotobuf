// Benchmark testing to measure the performance of marshaling and unmarshaling of ProtoBuf, JSON and XML
package main

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/golang/protobuf/proto"
)

var order = AddressBook{People: []*Person{
	{
		Name:  "John Doe",
		Id:    101,
		Email: "john@example.com",
	},
	{
		Name: "Jane Doe",
		Id:   102,
	},
	{
		Name:  "Jack Doe",
		Id:    201,
		Email: "jack@example.com",
		Phones: []*Person_PhoneNumber{
			{Number: "555-555-5555", Type: Person_WORK},
		},
	},
	{
		Name:  "Jack Buck",
		Id:    301,
		Email: "buck@example.com",
		Phones: []*Person_PhoneNumber{
			{Number: "555-555-0000", Type: Person_HOME},
			{Number: "555-555-0001", Type: Person_MOBILE},
			{Number: "555-555-0002", Type: Person_WORK},
		},
	},
	{
		Name:  "Janet Doe",
		Id:    1001,
		Email: "janet@example.com",
		Phones: []*Person_PhoneNumber{
			{Number: "555-777-0000"},
			{Number: "555-777-0001", Type: Person_HOME},
		},
	},
}}

// var order1 = &Order{
// 	ID:        "101",
// 	Status:    "Created",
// 	OrderItems: []*OrderItem{
// 		&OrderItem{
// 			Code:      "knd100",
// 			Name:      "Kindle Voyage",
// 			UnitPrice: 220,
// 			Quantity:  1,
// 		},
// 		&OrderItem{

// 			Code:      "kc101",
// 			Name:      "Kindle Voyage SmartShell Case",
// 			UnitPrice: 10,
// 			Quantity:  2,
// 		},
// 	},
// }

// Benchmark Proto3 Marshal
func BenchmarkOrderProto3Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&order)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark JSON Marshal
func BenchmarkOrderJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&order)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark XML Marshal
func BenchmarkOrderXMLMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(&order)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark Proto3 Unmarshal
func BenchmarkOrderProto3Unmarshal(b *testing.B) {
	data, err := proto.Marshal(&order)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}
	for i := 0; i < b.N; i++ {
		var order AddressBook
		err := proto.Unmarshal(data, &order)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark JSON Unmarshal
func BenchmarkOrderJSONUnmarshal(b *testing.B) {
	data, err := json.Marshal(&order)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}
	for i := 0; i < b.N; i++ {
		var order AddressBook
		err := json.Unmarshal(data, &order)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark XML Unmarshal
func BenchmarkOrderXMLUnmarshal(b *testing.B) {
	data, err := xml.Marshal(&order)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}
	for i := 0; i < b.N; i++ {
		var order AddressBook
		err := xml.Unmarshal(data, &order)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}