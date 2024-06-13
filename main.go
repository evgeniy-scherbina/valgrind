package main

// #include <uuid/uuid.h>
// #include <stdlib.h>
//
// // create a uuid function in C to return a uuid char*
// char* _go_uuid() {
//   uuid_t uuid;
//   uuid_generate_random(uuid);
//   char *str = malloc(37);
//   uuid_unparse_lower(uuid, str);
//   return str;
// }
import "C"
import (
	"fmt"
	"log"
	"time"

	"github.com/linxGnu/grocksdb"
)

const pathToDB = "/tmp/rocksdb_data"

// uuid generates a UUID using the C shared library.
// It returns a Go string.
func uuid() string {
	return C.GoString(C._go_uuid())
}

func memoryLeak() {
	for {
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		_ = uuid

		time.Sleep(time.Millisecond * 10)
	}
}

func func1() {
	func2()
}

func func2() {
	for {
		var uuid *C.uchar
		uuid = (*C.uchar)(C.malloc(16))
		_ = uuid

		time.Sleep(time.Millisecond * 10)
	}
}

func rocksdb() {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	cache := grocksdb.NewLRUCache(1 << 30)
	bbto.SetBlockCache(cache)
	bbto.SetFilterPolicy(grocksdb.NewBloomFilter(10))

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	opts.EnableStatistics()
	opts.SetStatsDumpPeriodSec(1)
	db, err := grocksdb.OpenDb(opts, pathToDB)
	if err != nil {
		log.Fatal(err)
	}

	simulateLoad(db)
}

func simulateLoad(db *grocksdb.DB) {
	ro := grocksdb.NewDefaultReadOptions()
	wo := grocksdb.NewDefaultWriteOptions()

	defaultKey := []byte("foo")
	defaultValue := []byte("bar")

	// if ro and wo are not used again, be sure to Close them.
	err := db.Put(wo, defaultKey, defaultValue)
	if err != nil {
		log.Fatal(err)
	}

	for {
		value, err := db.Get(ro, defaultKey)
		if err != nil {
			log.Fatal(err)
		}
		_ = value
		//defer value.Free()

		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	// and now it's simple to use
	myuuid := uuid() // this is a go string now
	fmt.Println(myuuid)

	memoryLeak()
	//func1()

	//rocksdb()
}
