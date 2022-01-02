package sm3

import "hash"

const (
	Size      = 32 // The size of an SM3 checksum in bytes.
	BlockSize = 64 // The blocksize of SM3 in bytes.
	chunk     = BlockSize
)

var (
	iv = [8]uint32{
		0x7380166f,
		0x4914b2b9,
		0x172442d7,
		0xda8a0600,
		0xa96f30bc,
		0x163138aa,
		0xe38dee4d,
		0xb0fb0e4e,
	}
)

func (d *digest) Size() int {
	return Size
}

func (d *digest) BlockSize() int {
	return BlockSize
}

type digest struct {
	h   [8]uint32   //  8*32=256
	x   [chunk]byte //chunk=64, 64*8=512
	nx  int
	len uint64
}

// New returns a new hash.Hash computing the SM4 checksum.
func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}
func (d *digest) Reset() {
	copy(d.h[:], iv[:])
	d.nx = 0
	d.len = 0
}
func (d *digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 {
		n := copy(d.x[d.nx:], p)
		d.nx += n
		if d.nx == chunk {
			block(d, d.x[:])
			d.nx = 0
		}
		p = p[n:]
	}

	if len(p) >= chunk {
		n := len(p) &^ (chunk - 1)
		block(d, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		d.nx = copy(d.x[:], p)
	}
	return
}
func (d0 *digest) Sum(b []byte) []byte {
	d := *d0
	hash := d.checkSum()
	return append(b, hash[:]...)
}

func (d *digest) checkSum() [Size]byte {
	len := d.len

	// padding method like crypto/sha1
	var tmp [64]byte
	tmp[0] = 0x80

	if len%64 < 56 {
		_, _ = d.Write(tmp[0 : 56-len%64])
	} else {
		_, _ = d.Write(tmp[0 : 64+56-len%64])
	}

	len <<= 3 // Length in bits.
	Uint64Tobyte(tmp[:], len)
	_, _ = d.Write(tmp[0:8])

	if d.nx != 0 {
		panic("d.nx != 0")
	}

	var digest [Size]byte
	Uint32Tobyte(digest[0:], d.h[0])
	Uint32Tobyte(digest[4:], d.h[1])
	Uint32Tobyte(digest[8:], d.h[2])
	Uint32Tobyte(digest[12:], d.h[3])
	Uint32Tobyte(digest[16:], d.h[4])
	Uint32Tobyte(digest[20:], d.h[5])
	Uint32Tobyte(digest[24:], d.h[6])
	Uint32Tobyte(digest[28:], d.h[7])

	return digest
}
