package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// GetRandomNumber: Gets a random number between given low and high values
func GetRandomNumber(_low int64, _high int64) int64 {

	var randomNumber int64
	randomNumber = -1

	if _low >= 0 && _low <= math.MaxInt64 && _high >= 0 && _high <= math.MaxInt64 {

		if _low < _high {
			rand.Seed(time.Now().UnixNano())
			randomNumber = rand.Int63n(_high-_low+1) + _low
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_low:", _low, "_high:", _high)
	}

	return randomNumber
}

// GetRandomData: Gets a byte array of given size filled with random data
func GetRandomData(_size int64) []byte {

	var randomData []byte
	randomData = nil

	if _size > 0 && _size <= math.MaxInt64 {

		randomData = make([]byte, _size)

		if int64(len(randomData)) != _size {
			fmt.Println(UI_RandomDataError)
		} else {
			_, err := rand.Read(randomData)

			if err != nil {
				fmt.Println(UI_RandomDataError, _size, err.Error())
			}

			if int64(len(randomData)) != _size {
				fmt.Println(UI_RandomDataError)
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_size:", _size)
	}

	return randomData
}

// ConvertMegabytesToBytes: Converts given file size in MB to bytes
func ConvertMegabytesToBytes(_sizeInMB int64) int64 {

	var sizeInBytes int64 = -1

	if _sizeInMB > 0 && _sizeInMB <= math.MaxInt64 {

		sizeInBytes = _sizeInMB * MB_IN_BYTES

		// If conversion has produced an invalid value, return error code
		if sizeInBytes < 0 || sizeInBytes > math.MaxInt64 {

			fmt.Println(UI_InvalidFileSizeError, sizeInBytes)
			sizeInBytes = -1
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_sizeInMB:", _sizeInMB)
	}

	return sizeInBytes
}
