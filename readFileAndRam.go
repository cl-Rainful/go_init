package main

import (
	"fmt"
	"os"
	"time"
)

func readRamRandomly() {
	var index [100000]int
	var arr [100000]int

	t0 := time.Now()
	defer func() {

		fmt.Printf("随机读写内存%dms\n", time.Since(t0).Microseconds())
	}()

	for _, i := range index {
		_ = arr[i]
	}
}

func readDiskSequentially() {
	rootPath := "d:/"
	var arr []byte

	t0 := time.Now()
	defer func() {

		fmt.Printf("顺序读写磁盘%dms\n", time.Since(t0).Microseconds())
	}()
	fin, err := os.Open(rootPath + "/data/arr.bin")
	if err != nil {
		panic(err)
	}
	defer fin.Close()
	fin.Read(arr)
}
