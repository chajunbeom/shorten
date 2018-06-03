package utils_test

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"shorten/utils"
	"strings"
	"testing"
	"time"
)

// 구현한 Base64 인코더와 Builtin 라이브러리 인코더 값을 비교하는 테스트
func TestBase64EncodingWithLibrary(t *testing.T) {
	testcase := 10
	testdata := []string{}
	// test data generation
	for i := 0; i < testcase; i++ {
		testdata = append(testdata, fmt.Sprintf("%d", time.Now().UnixNano()/1000))
	}

	lib := base64.NewEncoding(utils.Base64MappingTable)
	for _, data := range testdata {
		enc := utils.Base64Encode([]byte(data))
		ans := make([]byte, lib.EncodedLen(len(data)))
		lib.Encode(ans, []byte(data))

		// 직접 구현한 base64 인코더는 패딩을 생성하지 않는다.
		if string(enc) != strings.Trim(string(ans), "=") {
			t.Error(string(enc), "!=", string(ans))
		}
	}
	fmt.Println(t.Name() + ":OK")
}

// 구현한 CRC32 해시 생성기와 Builtin 라이브러리 인코더 값을 비교하는 테스트
func TestCRC32HashCodeWithLibrary(t *testing.T) {
	testcase := 10
	testdata := []string{}
	// test data generation
	for i := 0; i < testcase; i++ {
		testdata = append(testdata, fmt.Sprintf("%d", time.Now().UnixNano()/1000))
	}

	for _, v := range testdata {
		ans := crc32.ChecksumIEEE([]byte(v)) // CRC32 표준 해시 생성기
		testValue := utils.CRC32([]byte(v))
		if ans != testValue {
			t.Fail()
		}
	}
	fmt.Println(t.Name() + ":OK")
}
