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
	"flag"
	"fmt"
	"time"
)

const pathToDB = "/tmp/rocksdb_data"

// uuid generates a UUID using the C shared library.
// It returns a Go string.
func uuid() string {
	return C.GoString(C._go_uuid())
}

func memoryLeak() {
	var uuid *C.uchar
	uuid = (*C.uchar)(C.malloc(16))
	_ = uuid

	//for {
	//	var uuid *C.uchar
	//	uuid = (*C.uchar)(C.malloc(16))
	//	_ = uuid
	//
	//	time.Sleep(time.Millisecond * 10)
	//}
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

func main() {
	uuids := flag.Int("uuids", 1, "")
	flag.Parse()

	// and now it's simple to use
	for i := 0; i < *uuids; i++ {
		myuuid := uuid() // this is a go string now
		fmt.Println(myuuid)
	}

	//memoryLeak()
	//func1()

	//rocksdb()
}
