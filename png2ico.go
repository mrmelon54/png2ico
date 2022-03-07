package png2ico

import (
	"bytes"
	"encoding/binary"
)

type iconDir struct {
	reserved  uint16
	imageType uint16
	numImages uint16
}

type iconDirEntry struct {
	imageWidth   uint8
	imageHeight  uint8
	numColors    uint8
	reserved     uint8
	colorPlanes  uint16
	bitsPerPixel uint16
	sizeInBytes  uint32
	offset       uint32
}

func newIconDir() iconDir {
	return iconDir{imageType: 1, numImages: 1}
}

func newIconDirEntry() iconDirEntry {
	return iconDirEntry{colorPlanes: 1, bitsPerPixel: 32, offset: 22}
}

func ConvertPngToIco(im []byte, width, height int) ([]byte, error) {
	id := newIconDir()
	ide := newIconDirEntry()

	ide.sizeInBytes = uint32(len(im))
	ide.imageWidth = uint8(width)
	ide.imageHeight = uint8(height)
	bb := new(bytes.Buffer)

	err := binary.Write(bb, binary.LittleEndian, id)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bb, binary.LittleEndian, ide)
	if err != nil {
		return nil, err
	}

	bb.Write(im)
	return bb.Bytes(), nil
}
