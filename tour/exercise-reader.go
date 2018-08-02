package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/31
 */

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	// 赋值并返回
	b[0] = 'A'
	return 1, nil
}

func main() {
	m := MyReader{}
	b := make([]byte, 1)
	for {
		n, err := m.Read(b)
		fmt.Println(n, b[:n])
		if err != nil {
			break
		}
	}

}
