// Code generated by fastssz. DO NOT EDIT.
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the VoteSlashing object
func (v *VoteSlashing) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the VoteSlashing object to a target array
func (v *VoteSlashing) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'Vote1'
	dst = ssz.WriteOffset(dst, offset)
	if v.Vote1 == nil {
		v.Vote1 = new(MultiValidatorVote)
	}
	offset += v.Vote1.SizeSSZ()

	// Offset (1) 'Vote2'
	dst = ssz.WriteOffset(dst, offset)
	if v.Vote2 == nil {
		v.Vote2 = new(MultiValidatorVote)
	}
	offset += v.Vote2.SizeSSZ()

	// Field (0) 'Vote1'
	if dst, err = v.Vote1.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Vote2'
	if dst, err = v.Vote2.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the VoteSlashing object
func (v *VoteSlashing) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Vote1'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Offset (1) 'Vote2'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'Vote1'
	{
		buf = tail[o0:o1]
		if v.Vote1 == nil {
			v.Vote1 = new(MultiValidatorVote)
		}
		if err = v.Vote1.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Vote2'
	{
		buf = tail[o1:]
		if v.Vote2 == nil {
			v.Vote2 = new(MultiValidatorVote)
		}
		if err = v.Vote2.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the VoteSlashing object
func (v *VoteSlashing) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'Vote1'
	if v.Vote1 == nil {
		v.Vote1 = new(MultiValidatorVote)
	}
	size += v.Vote1.SizeSSZ()

	// Field (1) 'Vote2'
	if v.Vote2 == nil {
		v.Vote2 = new(MultiValidatorVote)
	}
	size += v.Vote2.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the VoteSlashing object
func (v *VoteSlashing) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the VoteSlashing object with a hasher
func (v *VoteSlashing) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Vote1'
	if err = v.Vote1.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Vote2'
	if err = v.Vote2.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the RANDAOSlashing object
func (r *RANDAOSlashing) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(r)
}

// MarshalSSZTo ssz marshals the RANDAOSlashing object to a target array
func (r *RANDAOSlashing) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'RandaoReveal'
	dst = append(dst, r.RandaoReveal[:]...)

	// Field (1) 'Slot'
	dst = ssz.MarshalUint64(dst, r.Slot)

	// Field (2) 'ValidatorPubkey'
	dst = append(dst, r.ValidatorPubkey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the RANDAOSlashing object
func (r *RANDAOSlashing) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 152 {
		return ssz.ErrSize
	}

	// Field (0) 'RandaoReveal'
	copy(r.RandaoReveal[:], buf[0:96])

	// Field (1) 'Slot'
	r.Slot = ssz.UnmarshallUint64(buf[96:104])

	// Field (2) 'ValidatorPubkey'
	copy(r.ValidatorPubkey[:], buf[104:152])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the RANDAOSlashing object
func (r *RANDAOSlashing) SizeSSZ() (size int) {
	size = 152
	return
}

// HashTreeRoot ssz hashes the RANDAOSlashing object
func (r *RANDAOSlashing) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(r)
}

// HashTreeRootWith ssz hashes the RANDAOSlashing object with a hasher
func (r *RANDAOSlashing) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	hh.PutBytes(r.RandaoReveal[:])

	// Field (1) 'Slot'
	hh.PutUint64(r.Slot)

	// Field (2) 'ValidatorPubkey'
	hh.PutBytes(r.ValidatorPubkey[:])

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the ProposerSlashing object
func (p *ProposerSlashing) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the ProposerSlashing object to a target array
func (p *ProposerSlashing) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'BlockHeader1'
	if p.BlockHeader1 == nil {
		p.BlockHeader1 = new(BlockHeader)
	}
	if dst, err = p.BlockHeader1.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'BlockHeader2'
	if p.BlockHeader2 == nil {
		p.BlockHeader2 = new(BlockHeader)
	}
	if dst, err = p.BlockHeader2.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Signature1'
	dst = append(dst, p.Signature1[:]...)

	// Field (3) 'Signature2'
	dst = append(dst, p.Signature2[:]...)

	// Field (4) 'ValidatorPublicKey'
	dst = append(dst, p.ValidatorPublicKey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ProposerSlashing object
func (p *ProposerSlashing) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 984 {
		return ssz.ErrSize
	}

	// Field (0) 'BlockHeader1'
	if p.BlockHeader1 == nil {
		p.BlockHeader1 = new(BlockHeader)
	}
	if err = p.BlockHeader1.UnmarshalSSZ(buf[0:372]); err != nil {
		return err
	}

	// Field (1) 'BlockHeader2'
	if p.BlockHeader2 == nil {
		p.BlockHeader2 = new(BlockHeader)
	}
	if err = p.BlockHeader2.UnmarshalSSZ(buf[372:744]); err != nil {
		return err
	}

	// Field (2) 'Signature1'
	copy(p.Signature1[:], buf[744:840])

	// Field (3) 'Signature2'
	copy(p.Signature2[:], buf[840:936])

	// Field (4) 'ValidatorPublicKey'
	copy(p.ValidatorPublicKey[:], buf[936:984])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ProposerSlashing object
func (p *ProposerSlashing) SizeSSZ() (size int) {
	size = 984
	return
}

// HashTreeRoot ssz hashes the ProposerSlashing object
func (p *ProposerSlashing) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the ProposerSlashing object with a hasher
func (p *ProposerSlashing) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'BlockHeader1'
	if err = p.BlockHeader1.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'BlockHeader2'
	if err = p.BlockHeader2.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature1'
	hh.PutBytes(p.Signature1[:])

	// Field (3) 'Signature2'
	hh.PutBytes(p.Signature2[:])

	// Field (4) 'ValidatorPublicKey'
	hh.PutBytes(p.ValidatorPublicKey[:])

	hh.Merkleize(indx)
	return
}