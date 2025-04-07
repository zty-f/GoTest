package util

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "a68a9b525e89094490271c5ab4b521baa9dcba3164290f6c3fb0970a3a050fc8"
	fmt.Println(Md51(str)) // 13cef9838e01192159be4d8282a75ca9
	fmt.Println(Md52(str)) // 13cef9838e01192159be4d8282a75ca9
	fmt.Println(Md53(str)) // 13cef9838e01192159be4d8282a75ca9
	fmt.Println(Md54(str)) // 13cef9838e01192159be4d8282a75ca9
	fmt.Println(Md55(str)) // 13cef9838e01192159be4d8282a75ca9
	fmt.Println(Md56(str)) // 13cef9838e01192159be4d8282a75ca9
}
