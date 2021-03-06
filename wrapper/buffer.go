package wrapper

import (
	buffer "github.com/Dmitriy-Vas/wave_buffer"
	"time"
)

type Buffer struct {
	*buffer.DefaultBuffer
}

func (b *Buffer) Clone() buffer.PacketBuffer {
	buf := &Buffer{DefaultBuffer: b.DefaultBuffer.Clone().(*buffer.DefaultBuffer)}
	buf.Reset()
	return buf
}

func (b *Buffer) WriteBool(data []byte, value bool, index uint64) {
	var v byte = 0
	if value {
		v = 1
	}
	b.WriteByte(data, v, index)
}

func (b *Buffer) WriteByte(data []byte, value byte, index uint64) {
	b.DefaultBuffer.WriteShort(data, int16(value), index)
}

func (b *Buffer) WriteString(data []byte, value string, index uint64) {
	b.DefaultBuffer.WriteInt(data, int32(len(value)), index)
	b.DefaultBuffer.WriteString(data, value, index+4)
}

func (b *Buffer) WriteBytes(data []byte, value []byte, index uint64) {
	b.DefaultBuffer.WriteInt(data, int32(len(value)), index)
	b.DefaultBuffer.WriteBytes(data, value, index+4)
}

func WriteDate(buffer buffer.PacketBuffer, t time.Time) {
	str := t.Format("02/01/2006")
	buffer.WriteString(buffer.Bytes(), str, buffer.Index())
}

func (b *Buffer) ReadBool(data []byte, index uint64) bool {
	return b.ReadByte(data, index) != 0
}

func (b *Buffer) ReadByte(data []byte, index uint64) byte {
	return byte(b.DefaultBuffer.ReadShort(data, index))
}

func (b *Buffer) ReadString(data []byte, index uint64, length uint64) string {
	length = uint64(b.DefaultBuffer.ReadInt(data, index))
	return b.DefaultBuffer.ReadString(data, index+4, length)
}

func (b *Buffer) ReadBytes(data []byte, index uint64, length uint64) []byte {
	length = uint64(b.DefaultBuffer.ReadInt(data, index))
	return b.DefaultBuffer.ReadBytes(data, index+4, length)
}

func ReadDate(buffer buffer.PacketBuffer) time.Time {
	str := buffer.ReadString(buffer.Bytes(), buffer.Index(), 0)
	if str == "" {
		return time.Time{}
	}
	t, _ := time.Parse("02/01/2006", str)
	return t
}
