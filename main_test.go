package main

import (
	"context"
	"testing"
)

func BenchmarkFindByUUIDV7NewRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readNewerRecordsfile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByUUIDv7(collection, records[n].UUIDv7)
	}
}

func BenchmarkFindByUUIDV7OldRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readOlderRecordsFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByUUIDv7(collection, records[n].UUIDv7)
	}
}

func BenchmarkFindByUUIDV7RandomRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readRandomFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByUUIDv7(collection, records[n].UUIDv7)
	}
}

func BenchmarkFindByUUIDV4NewRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readNewerRecordsfile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByUUIDv4(collection, records[n].UUIDv4)
	}
}

func BenchmarkFindByUUIDV4OldRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readOlderRecordsFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByUUIDv4(collection, records[n].UUIDv4)
	}
}

func BenchmarkFindByUUIDV4RandomRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readRandomFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByUUIDv4(collection, records[n].UUIDv4)
	}
}

// --
func BenchmarkFindByMongoIDNewRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readNewerRecordsfile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByMongoID(collection, records[n].Id)
	}
}

func BenchmarkFindByMongoIDOldRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readOlderRecordsFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByMongoID(collection, records[n].Id)
	}
}

func BenchmarkFindByMongoIDRandomRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readRandomFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByMongoID(collection, records[n].Id)
	}
}

// ---
func BenchmarkFindByNumericIDNewRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readNewerRecordsfile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByNumID(collection, records[n].NumericID)
	}
}

func BenchmarkFindByNumericIDOldRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readOlderRecordsFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByNumID(collection, records[n].NumericID)
	}
}

func BenchmarkFindByNumericIDRandomRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readRandomFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByNumID(collection, records[n].NumericID)
	}
}

func BenchmarkFindByNumericStringIDNewRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readNewerRecordsfile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByNumIDString(collection, records[n].NumericIDStr)
	}
}

func BenchmarkFindByNumericIDStringOldRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readOlderRecordsFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByNumIDString(collection, records[n].NumericIDStr)
	}
}

func BenchmarkFindByNumericIDStringRandomRecords(b *testing.B) {
	b.StopTimer()
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())
	records := readRandomFile()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		FindByNumIDString(collection, records[n].NumericIDStr)
	}
}
