package countOcc

import (
     "os"
    "fmt"
    "bufio"
    "bytes"
    //"log"
)

func ReadFASTA(sequence_file string) []byte {
    f, err := os.Open(sequence_file)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
    defer f.Close()
    br := bufio.NewReader(f)
    byte_array := bytes.Buffer{}
    var isPrefix bool
    _, isPrefix, err = br.ReadLine()
    if err != nil || isPrefix {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
    //fmt.Printf("%s",line)
    var line []byte
    for {
        line, isPrefix, err = br.ReadLine()
        if err != nil || isPrefix {
        break
        } else {
            byte_array.Write(line)
        }
    }
    return []byte(byte_array.String())
}
