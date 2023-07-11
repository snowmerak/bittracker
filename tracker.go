package bittracker

const (
	mask1 = uint8(0b10000000)
	mask2 = uint8(0b11000000)
	mask3 = uint8(0b11100000)
	mask4 = uint8(0b11110000)
	mask5 = uint8(0b11111000)
	mask6 = uint8(0b11111100)
	mask7 = uint8(0b11111110)
	mask8 = uint8(0b11111111)
)

/*
BitTracker is a simple bit tracker for a byte array.
It's used to track the bits of a byte array.
*/
type BitTracker struct {
	data []byte
}

func NewBitTracker(data []byte) *BitTracker {
	return &BitTracker{data}
}

// GetBit returns the bit at the given index.
func (bt *BitTracker) GetBit(index int) bool {
	if index < 0 || index >= len(bt.data)*8 {
		return false
	}

	byteIndex := index / 8
	bitIndex := index%8 - 1
	if bitIndex < 0 {
		bitIndex = 7
	}

	return bt.data[byteIndex]&(1<<uint(bitIndex)) != 0
}

// SetBit sets the bit at the given index.
func (bt *BitTracker) SetBit(index int, value bool) {
	if index < 0 || index >= len(bt.data)*8 {
		return
	}

	byteIndex := index / 8
	bitIndex := index%8 - 1
	if bitIndex < 0 {
		bitIndex = 7
	}

	if value {
		bt.data[byteIndex] |= 1 << uint(bitIndex)
	} else {
		bt.data[byteIndex] &= ^(1 << uint(bitIndex))
	}
}

// ToggleBit toggles the bit at the given index.
func (bt *BitTracker) ToggleBit(index int) {
	if index < 0 || index >= len(bt.data)*8 {
		return
	}

	byteIndex := index / 8
	bitIndex := index % 8

	bt.data[byteIndex] ^= 1 << uint(bitIndex)
}

// GetRange returns the []byte in the given range.
// The range is inclusive.
// If the range is invalid, it returns nil.
func (bt *BitTracker) GetRange(start, end int) []byte {
	if start < 0 || start > len(bt.data)*8 {
		return nil
	}

	if end < 0 || end > len(bt.data)*8 {
		return nil
	}

	if start > end {
		return nil
	}

	remains := end - start

	sourceByteIndex := start / 8
	sourceBitIndex := start % 8

	targetByteIndex := 0
	targetBitIndex := 0

	resultLength := remains / 8
	if remains%8 != 0 {
		resultLength++
	}

	result := make([]byte, resultLength)
	// copy from start bit index in start byte index until end bit index in end byte index to result
	for {
		min := 8 - sourceBitIndex
		if min > 8-targetBitIndex {
			min = 8 - targetBitIndex
		}
		if remains < min {
			min = remains
		}

		mask := uint8(0)
		switch min {
		case 0:
			mask = 0
		case 1:
			mask = mask1
		case 2:
			mask = mask2
		case 3:
			mask = mask3
		case 4:
			mask = mask4
		case 5:
			mask = mask5
		case 6:
			mask = mask6
		case 7:
			mask = mask7
		case 8:
			mask = mask8
		}
		result[targetByteIndex] |= (bt.data[sourceByteIndex] << sourceBitIndex & mask) >> targetBitIndex

		remains -= min
		if remains <= 0 {
			break
		}

		sourceBitIndex += min
		targetBitIndex += min

		if sourceBitIndex >= 8 {
			sourceBitIndex = 0
			sourceByteIndex++
		}
		if targetBitIndex >= 8 {
			targetBitIndex = 0
			targetByteIndex++
		}
	}

	return result
}
