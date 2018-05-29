package rtfs_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/Temporal/rtfs"
)

const testPIN = "QmNZiPk974vDsPmQii3YbrMKfi12KTSNM7XMiYyiea4VYZ"

func TestInitialize(t *testing.T) {
	im := rtfs.Initialize("")
	_, err := im.Shell.ID()
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
}

/*
func TestPin(t *testing.T) {
	im := rtfs.Initialize("")
	err := im.Pin(testPIN)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("pin test successful")
}*/

func TestGetObjectFileSizeInBytes(t *testing.T) {
	im := rtfs.Initialize("")
	size, err := im.GetObjectFileSizeInBytes(testPIN)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("file size in bytes ", size)
}
