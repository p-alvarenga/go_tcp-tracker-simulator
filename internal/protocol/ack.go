package protocol

type ACKPacket struct {
	Kind   PacketType
	Serial uint16
}

func CheckACK(ack *ACKPacket, pkt Packet) bool {
	return ack.Kind == pkt.Kind() && ack.Serial == pkt.GetSerial()
}
