// Code generated by generate1024.go. DO NOT EDIT.

package mlkem

import (
	"crypto/internal/fips140"
	"crypto/internal/fips140/drbg"
	"crypto/internal/fips140/sha3"
	"crypto/internal/fips140/subtle"
	"errors"
)

// A DecapsulationKey1024 is the secret key used to decapsulate a shared key from a
// ciphertext. It includes various precomputed values.
type DecapsulationKey1024 struct {
	d [32]byte // decapsulation key seed
	z [32]byte // implicit rejection sampling seed

	ρ [32]byte // sampleNTT seed for A, stored for the encapsulation key
	h [32]byte // H(ek), stored for ML-KEM.Decaps_internal

	encryptionKey1024
	decryptionKey1024
}

// Bytes returns the decapsulation key as a 64-byte seed in the "d || z" form.
//
// The decapsulation key must be kept secret.
func (dk *DecapsulationKey1024) Bytes() []byte {
	var b [SeedSize]byte
	copy(b[:], dk.d[:])
	copy(b[32:], dk.z[:])
	return b[:]
}

// EncapsulationKey returns the public encapsulation key necessary to produce
// ciphertexts.
func (dk *DecapsulationKey1024) EncapsulationKey() *EncapsulationKey1024 {
	return &EncapsulationKey1024{
		ρ:                 dk.ρ,
		h:                 dk.h,
		encryptionKey1024: dk.encryptionKey1024,
	}
}

// An EncapsulationKey1024 is the public key used to produce ciphertexts to be
// decapsulated by the corresponding [DecapsulationKey1024].
type EncapsulationKey1024 struct {
	ρ [32]byte // sampleNTT seed for A
	h [32]byte // H(ek)
	encryptionKey1024
}

// Bytes returns the encapsulation key as a byte slice.
func (ek *EncapsulationKey1024) Bytes() []byte {
	// The actual logic is in a separate function to outline this allocation.
	b := make([]byte, 0, EncapsulationKeySize1024)
	return ek.bytes(b)
}

func (ek *EncapsulationKey1024) bytes(b []byte) []byte {
	for i := range ek.t {
		b = polyByteEncode(b, ek.t[i])
	}
	b = append(b, ek.ρ[:]...)
	return b
}

// encryptionKey1024 is the parsed and expanded form of a PKE encryption key.
type encryptionKey1024 struct {
	t [k1024]nttElement         // ByteDecode₁₂(ek[:384k])
	a [k1024 * k1024]nttElement // A[i*k+j] = sampleNTT(ρ, j, i)
}

// decryptionKey1024 is the parsed and expanded form of a PKE decryption key.
type decryptionKey1024 struct {
	s [k1024]nttElement // ByteDecode₁₂(dk[:decryptionKey1024Size])
}

// GenerateKey1024 generates a new decapsulation key, drawing random bytes from
// a DRBG. The decapsulation key must be kept secret.
func GenerateKey1024() (*DecapsulationKey1024, error) {
	// The actual logic is in a separate function to outline this allocation.
	dk := &DecapsulationKey1024{}
	return generateKey1024(dk)
}

func generateKey1024(dk *DecapsulationKey1024) (*DecapsulationKey1024, error) {
	var d [32]byte
	drbg.Read(d[:])
	var z [32]byte
	drbg.Read(z[:])
	kemKeyGen1024(dk, &d, &z)
	if err := fips140.PCT("ML-KEM PCT", func() error { return kemPCT1024(dk) }); err != nil {
		// This clearly can't happen, but FIPS 140-3 requires us to check.
		panic(err)
	}
	fips140.RecordApproved()
	return dk, nil
}

// GenerateKeyInternal1024 is a derandomized version of GenerateKey1024,
// exclusively for use in tests.
func GenerateKeyInternal1024(d, z *[32]byte) *DecapsulationKey1024 {
	dk := &DecapsulationKey1024{}
	kemKeyGen1024(dk, d, z)
	return dk
}

// NewDecapsulationKey1024 parses a decapsulation key from a 64-byte
// seed in the "d || z" form. The seed must be uniformly random.
func NewDecapsulationKey1024(seed []byte) (*DecapsulationKey1024, error) {
	// The actual logic is in a separate function to outline this allocation.
	dk := &DecapsulationKey1024{}
	return newKeyFromSeed1024(dk, seed)
}

func newKeyFromSeed1024(dk *DecapsulationKey1024, seed []byte) (*DecapsulationKey1024, error) {
	if len(seed) != SeedSize {
		return nil, errors.New("mlkem: invalid seed length")
	}
	d := (*[32]byte)(seed[:32])
	z := (*[32]byte)(seed[32:])
	kemKeyGen1024(dk, d, z)
	if err := fips140.PCT("ML-KEM PCT", func() error { return kemPCT1024(dk) }); err != nil {
		// This clearly can't happen, but FIPS 140-3 requires us to check.
		panic(err)
	}
	fips140.RecordApproved()
	return dk, nil
}

// kemKeyGen1024 generates a decapsulation key.
//
// It implements ML-KEM.KeyGen_internal according to FIPS 203, Algorithm 16, and
// K-PKE.KeyGen according to FIPS 203, Algorithm 13. The two are merged to save
// copies and allocations.
func kemKeyGen1024(dk *DecapsulationKey1024, d, z *[32]byte) {
	dk.d = *d
	dk.z = *z

	g := sha3.New512()
	g.Write(d[:])
	g.Write([]byte{k1024}) // Module dimension as a domain separator.
	G := g.Sum(make([]byte, 0, 64))
	ρ, σ := G[:32], G[32:]
	dk.ρ = [32]byte(ρ)

	A := &dk.a
	for i := byte(0); i < k1024; i++ {
		for j := byte(0); j < k1024; j++ {
			A[i*k1024+j] = sampleNTT(ρ, j, i)
		}
	}

	var N byte
	s := &dk.s
	for i := range s {
		s[i] = ntt(samplePolyCBD(σ, N))
		N++
	}
	e := make([]nttElement, k1024)
	for i := range e {
		e[i] = ntt(samplePolyCBD(σ, N))
		N++
	}

	t := &dk.t
	for i := range t { // t = A ◦ s + e
		t[i] = e[i]
		for j := range s {
			t[i] = polyAdd(t[i], nttMul(A[i*k1024+j], s[j]))
		}
	}

	H := sha3.New256()
	ek := dk.EncapsulationKey().Bytes()
	H.Write(ek)
	H.Sum(dk.h[:0])
}

// kemPCT1024 performs a Pairwise Consistency Test per FIPS 140-3 IG 10.3.A
// Additional Comment 1: "For key pairs generated for use with approved KEMs in
// FIPS 203, the PCT shall consist of applying the encapsulation key ek to
// encapsulate a shared secret K leading to ciphertext c, and then applying
// decapsulation key dk to retrieve the same shared secret K. The PCT passes if
// the two shared secret K values are equal. The PCT shall be performed either
// when keys are generated/imported, prior to the first exportation, or prior to
// the first operational use (if not exported before the first use)."
func kemPCT1024(dk *DecapsulationKey1024) error {
	ek := dk.EncapsulationKey()
	K, c := ek.Encapsulate()
	K1, err := dk.Decapsulate(c)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(K, K1) != 1 {
		return errors.New("mlkem: PCT failed")
	}
	return nil
}

// Encapsulate generates a shared key and an associated ciphertext from an
// encapsulation key, drawing random bytes from a DRBG.
//
// The shared key must be kept secret.
func (ek *EncapsulationKey1024) Encapsulate() (sharedKey, ciphertext []byte) {
	// The actual logic is in a separate function to outline this allocation.
	var cc [CiphertextSize1024]byte
	return ek.encapsulate(&cc)
}

func (ek *EncapsulationKey1024) encapsulate(cc *[CiphertextSize1024]byte) (sharedKey, ciphertext []byte) {
	var m [messageSize]byte
	drbg.Read(m[:])
	// Note that the modulus check (step 2 of the encapsulation key check from
	// FIPS 203, Section 7.2) is performed by polyByteDecode in parseEK1024.
	fips140.RecordApproved()
	return kemEncaps1024(cc, ek, &m)
}

// EncapsulateInternal is a derandomized version of Encapsulate, exclusively for
// use in tests.
func (ek *EncapsulationKey1024) EncapsulateInternal(m *[32]byte) (sharedKey, ciphertext []byte) {
	cc := &[CiphertextSize1024]byte{}
	return kemEncaps1024(cc, ek, m)
}

// kemEncaps1024 generates a shared key and an associated ciphertext.
//
// It implements ML-KEM.Encaps_internal according to FIPS 203, Algorithm 17.
func kemEncaps1024(cc *[CiphertextSize1024]byte, ek *EncapsulationKey1024, m *[messageSize]byte) (K, c []byte) {
	g := sha3.New512()
	g.Write(m[:])
	g.Write(ek.h[:])
	G := g.Sum(nil)
	K, r := G[:SharedKeySize], G[SharedKeySize:]
	c = pkeEncrypt1024(cc, &ek.encryptionKey1024, m, r)
	return K, c
}

// NewEncapsulationKey1024 parses an encapsulation key from its encoded form.
// If the encapsulation key is not valid, NewEncapsulationKey1024 returns an error.
func NewEncapsulationKey1024(encapsulationKey []byte) (*EncapsulationKey1024, error) {
	// The actual logic is in a separate function to outline this allocation.
	ek := &EncapsulationKey1024{}
	return parseEK1024(ek, encapsulationKey)
}

// parseEK1024 parses an encryption key from its encoded form.
//
// It implements the initial stages of K-PKE.Encrypt according to FIPS 203,
// Algorithm 14.
func parseEK1024(ek *EncapsulationKey1024, ekPKE []byte) (*EncapsulationKey1024, error) {
	if len(ekPKE) != EncapsulationKeySize1024 {
		return nil, errors.New("mlkem: invalid encapsulation key length")
	}

	h := sha3.New256()
	h.Write(ekPKE)
	h.Sum(ek.h[:0])

	for i := range ek.t {
		var err error
		ek.t[i], err = polyByteDecode[nttElement](ekPKE[:encodingSize12])
		if err != nil {
			return nil, err
		}
		ekPKE = ekPKE[encodingSize12:]
	}
	copy(ek.ρ[:], ekPKE)

	for i := byte(0); i < k1024; i++ {
		for j := byte(0); j < k1024; j++ {
			ek.a[i*k1024+j] = sampleNTT(ek.ρ[:], j, i)
		}
	}

	return ek, nil
}

// pkeEncrypt1024 encrypt a plaintext message.
//
// It implements K-PKE.Encrypt according to FIPS 203, Algorithm 14, although the
// computation of t and AT is done in parseEK1024.
func pkeEncrypt1024(cc *[CiphertextSize1024]byte, ex *encryptionKey1024, m *[messageSize]byte, rnd []byte) []byte {
	var N byte
	r, e1 := make([]nttElement, k1024), make([]ringElement, k1024)
	for i := range r {
		r[i] = ntt(samplePolyCBD(rnd, N))
		N++
	}
	for i := range e1 {
		e1[i] = samplePolyCBD(rnd, N)
		N++
	}
	e2 := samplePolyCBD(rnd, N)

	u := make([]ringElement, k1024) // NTT⁻¹(AT ◦ r) + e1
	for i := range u {
		u[i] = e1[i]
		for j := range r {
			// Note that i and j are inverted, as we need the transposed of A.
			u[i] = polyAdd(u[i], inverseNTT(nttMul(ex.a[j*k1024+i], r[j])))
		}
	}

	μ := ringDecodeAndDecompress1(m)

	var vNTT nttElement // t⊺ ◦ r
	for i := range ex.t {
		vNTT = polyAdd(vNTT, nttMul(ex.t[i], r[i]))
	}
	v := polyAdd(polyAdd(inverseNTT(vNTT), e2), μ)

	c := cc[:0]
	for _, f := range u {
		c = ringCompressAndEncode11(c, f)
	}
	c = ringCompressAndEncode5(c, v)

	return c
}

// Decapsulate generates a shared key from a ciphertext and a decapsulation key.
// If the ciphertext is not valid, Decapsulate returns an error.
//
// The shared key must be kept secret.
func (dk *DecapsulationKey1024) Decapsulate(ciphertext []byte) (sharedKey []byte, err error) {
	if len(ciphertext) != CiphertextSize1024 {
		return nil, errors.New("mlkem: invalid ciphertext length")
	}
	c := (*[CiphertextSize1024]byte)(ciphertext)
	// Note that the hash check (step 3 of the decapsulation input check from
	// FIPS 203, Section 7.3) is foregone as a DecapsulationKey is always
	// validly generated by ML-KEM.KeyGen_internal.
	return kemDecaps1024(dk, c), nil
}

// kemDecaps1024 produces a shared key from a ciphertext.
//
// It implements ML-KEM.Decaps_internal according to FIPS 203, Algorithm 18.
func kemDecaps1024(dk *DecapsulationKey1024, c *[CiphertextSize1024]byte) (K []byte) {
	fips140.RecordApproved()
	m := pkeDecrypt1024(&dk.decryptionKey1024, c)
	g := sha3.New512()
	g.Write(m[:])
	g.Write(dk.h[:])
	G := g.Sum(make([]byte, 0, 64))
	Kprime, r := G[:SharedKeySize], G[SharedKeySize:]
	J := sha3.NewShake256()
	J.Write(dk.z[:])
	J.Write(c[:])
	Kout := make([]byte, SharedKeySize)
	J.Read(Kout)
	var cc [CiphertextSize1024]byte
	c1 := pkeEncrypt1024(&cc, &dk.encryptionKey1024, (*[32]byte)(m), r)

	subtle.ConstantTimeCopy(subtle.ConstantTimeCompare(c[:], c1), Kout, Kprime)
	return Kout
}

// pkeDecrypt1024 decrypts a ciphertext.
//
// It implements K-PKE.Decrypt according to FIPS 203, Algorithm 15,
// although s is retained from kemKeyGen1024.
func pkeDecrypt1024(dx *decryptionKey1024, c *[CiphertextSize1024]byte) []byte {
	u := make([]ringElement, k1024)
	for i := range u {
		b := (*[encodingSize11]byte)(c[encodingSize11*i : encodingSize11*(i+1)])
		u[i] = ringDecodeAndDecompress11(b)
	}

	b := (*[encodingSize5]byte)(c[encodingSize11*k1024:])
	v := ringDecodeAndDecompress5(b)

	var mask nttElement // s⊺ ◦ NTT(u)
	for i := range dx.s {
		mask = polyAdd(mask, nttMul(dx.s[i], ntt(u[i])))
	}
	w := polySub(v, inverseNTT(mask))

	return ringCompressAndEncode1(nil, w)
}
