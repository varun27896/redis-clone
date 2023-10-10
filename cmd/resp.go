package resp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Resp() {
	input := "$5\r\nVarun\r\n"
	//We create a string and convert it to a bufio buffer.
	reader := bufio.NewReader(strings.NewReader(input))

	//next we read the data type, which is the first byte from the buffer
	b, _ := reader.ReadByte()

	if b != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	//read the number to determine the size of the string
	size, _ := reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	// consume /r/n
	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))

}
