package day16

import "log"

const filePath = "input1.txt"

type bits []int

func (b *bits) version(head *int) (version int) {
	for i := *head; i < *head+3; i++ {
		version <<= 1
		version |= (*b)[i]
	}
	*head += 3
	return version
}

func (b *bits) typeID(head *int) (typeID int) {
	for i := *head; i < *head+3; i++ {
		typeID <<= 1
		typeID |= (*b)[i]
	}

	*head += 3
	return typeID
}

func (b *bits) literal(head *int) (literal int) {
	stop := false
	for !stop {
		if (*b)[*head] == 0 {
			stop = true
		}

		for j := *head + 1; j < *head+5; j++ {
			literal <<= 1
			literal |= (*b)[j]
		}

		*head += 5
	}

	return literal
}

func (b *bits) lengthTypeID(head *int) (lengthTypeID, length int) {
	if lengthTypeID = (*b)[*head]; lengthTypeID == 0 {
		for i := *head + 1; i <= *head+15; i++ {
			length <<= 1
			length |= (*b)[i]
		}
		*head += 15
	} else {
		for i := *head + 1; i <= *head+11; i++ {
			length <<= 1
			length |= (*b)[i]
		}
		*head += 11
	}
	*head += 1
	return
}

func (b *bits) operator(head *int) []packet {
	subPackets := make([]packet, 0)
	lengthTypeID, length := b.lengthTypeID(head)
	if lengthTypeID == 0 {
		bitsEnd := *head + length
		for *head < bitsEnd {
			subPackets = append(subPackets, toPacket(*b, head))
		}
	} else {
		for i := 0; i < length; i++ {
			subPackets = append(subPackets, toPacket(*b, head))
		}
	}
	return subPackets
}

type packet struct {
	version    int
	typeID     int
	literal    int
	subPackets []packet
}

func (p *packet) sumVersions() int {
	sum := p.version
	for _, sp := range p.subPackets {
		sum += sp.sumVersions()
	}
	return sum
}

func (p packet) value() int {
	switch p.typeID {
	case 0:
		v := 0
		for _, sp := range p.subPackets {
			v += sp.value()
		}
		return v
	case 1:
		v := 1
		for _, sp := range p.subPackets {
			v *= sp.value()
		}
		return v
	case 2:
		v := -1
		for _, sp := range p.subPackets {
			subV := sp.value()
			if v == -1 || subV < v {
				v = subV
			}
		}
		return v
	case 3:
		v := -1
		for _, sp := range p.subPackets {
			subV := sp.value()
			if v == -1 || subV > v {
				v = subV
			}
		}
		return v
	case 4:
		return p.literal
	case 5:
		if p.subPackets[0].value() > p.subPackets[1].value() {
			return 1
		}
		return 0
	case 6:
		if p.subPackets[0].value() < p.subPackets[1].value() {
			return 1
		}
		return 0
	case 7:
		if p.subPackets[0].value() == p.subPackets[1].value() {
			return 1
		}
		return 0
	default:
		log.Fatal("unexpected value for operator")
		return 0
	}
}

var hexToByte = map[string][]int{
	"0": {0, 0, 0, 0},
	"1": {0, 0, 0, 1},
	"2": {0, 0, 1, 0},
	"3": {0, 0, 1, 1},
	"4": {0, 1, 0, 0},
	"5": {0, 1, 0, 1},
	"6": {0, 1, 1, 0},
	"7": {0, 1, 1, 1},
	"8": {1, 0, 0, 0},
	"9": {1, 0, 0, 1},
	"A": {1, 0, 1, 0},
	"B": {1, 0, 1, 1},
	"C": {1, 1, 0, 0},
	"D": {1, 1, 0, 1},
	"E": {1, 1, 1, 0},
	"F": {1, 1, 1, 1},
}

func toPacket(bits bits, head *int) packet {
	var literal int
	var subPackets []packet

	v := bits.version(head)
	typeID := bits.typeID(head)

	if typeID == 4 {
		literal = bits.literal(head)
	} else {
		subPackets = bits.operator(head)
	}

	return packet{
		version:    v,
		typeID:     typeID,
		literal:    literal,
		subPackets: subPackets,
	}
}

func parseInput(lines []string) bits {
	b := make(bits, 0)
	for _, line := range lines {
		for _, hex := range line {
			if bits, ok := hexToByte[string(hex)]; ok {
				b = append(b, bits...)
			}
		}
	}

	return b
}

func exe2(lines []string) int {
	bits := parseInput(lines)

	head := 0
	p := toPacket(bits, &head)
	return p.value()
}

func exe1(lines []string) int {
	bits := parseInput(lines)

	head := 0
	p := toPacket(bits, &head)
	return p.sumVersions()
}
