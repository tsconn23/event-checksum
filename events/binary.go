package events

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/ugorji/go/codec"
)

func NewBinaryEvent() (data []byte, err error) {
	imgPath, err := findPath()
	if err != nil {
		return
	}
	imgPath +=  "/lebowski.jpg"
	file, err := os.Open(imgPath)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	timestamp := makeTimestamp()
	deviceName := "RandomDevice-2"
	evt := models.Event{ Created:timestamp, Modified:timestamp, Device:deviceName }
	readings := []models.Reading{}
	readings = append(readings, models.Reading{Created:timestamp, Modified:timestamp, Device:deviceName, Name:"Reading2", Value:"789"})
	readings = append(readings, models.Reading{Created:timestamp, Modified:timestamp, Device:deviceName, Name:"Reading1", Value:"XYZ"})
	readings = append(readings, models.Reading{Created:timestamp, Modified:timestamp, Device:deviceName, Name:"Reading1", BinaryValue:bytes})
	evt.Readings = readings

	var handle codec.CborHandle
	data = make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&data, &handle)

	err = enc.Encode(evt)
	return
}

func findPath() (path string, err error) {
	exec, err := os.Executable()
	if err != nil {
		return
	}
	path = filepath.Dir(exec)
	path += "/img"
	return
}