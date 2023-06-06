// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ClickHouse/ch-go/internal/gold"
)

func newByte16(v int) [16]byte {
	return [16]byte{0: byte(v)}
}

func TestColFixedStr16_DecodeColumn(t *testing.T) {
	t.Parallel()
	const rows = 50
	var data ColFixedStr16
	for i := 0; i < rows; i++ {
		v := newByte16(i)
		data.Append(v)
		require.Equal(t, v, data.Row(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		t.Parallel()
		gold.Bytes(t, buf.Buf, "col_byte16")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		var dec ColFixedStr16
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())

		require.Equal(t, ColumnTypeFixedString.With("16"), dec.Type())

	})
	t.Run("ZeroRows", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColFixedStr16
		require.NoError(t, dec.DecodeColumn(r, 0))
	})
	t.Run("EOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColFixedStr16
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.EOF)
	})
	t.Run("NoShortRead", func(t *testing.T) {
		var dec ColFixedStr16
		requireNoShortRead(t, buf.Buf, colAware(&dec, rows))
	})
	t.Run("ZeroRowsEncode", func(t *testing.T) {
		var v ColFixedStr16
		v.EncodeColumn(nil) // should be no-op
	})
}
func TestColFixedStr16Array(t *testing.T) {
	const rows = 50
	data := NewArrFixedStr16()
	for i := 0; i < rows; i++ {
		data.Append([][16]byte{
			newByte16(i),
			newByte16(i + 1),
			newByte16(i + 2),
		})
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		gold.Bytes(t, buf.Buf, "col_arr_byte16")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		dec := NewArrFixedStr16()
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())
		require.Equal(t, ColumnTypeFixedString.With("16").Array(), dec.Type())
	})
	t.Run("EOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		dec := NewArrFixedStr16()
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.EOF)
	})
}

func BenchmarkColFixedStr16_DecodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColFixedStr16
	for i := 0; i < rows; i++ {
		data = append(data, newByte16(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	br := bytes.NewReader(buf.Buf)
	r := NewReader(br)

	var dec ColFixedStr16
	if err := dec.DecodeColumn(r, rows); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		br.Reset(buf.Buf)
		r.raw.Reset(br)
		dec.Reset()

		if err := dec.DecodeColumn(r, rows); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColFixedStr16_EncodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColFixedStr16
	for i := 0; i < rows; i++ {
		data = append(data, newByte16(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		data.EncodeColumn(&buf)
	}
}
