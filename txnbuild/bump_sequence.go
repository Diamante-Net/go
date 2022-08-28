package txnbuild

import (
	"github.com/diamnet/go/support/errors"
	"github.com/diamnet/go/xdr"
)

// BumpSequence represents the Diamnet bump sequence operation. See
// https://developers.diamnet.org/docs/start/list-of-operations/
type BumpSequence struct {
	BumpTo        int64
	SourceAccount string
}

// BuildXDR for BumpSequence returns a fully configured XDR Operation.
func (bs *BumpSequence) BuildXDR() (xdr.Operation, error) {
	opType := xdr.OperationTypeBumpSequence
	xdrOp := xdr.BumpSequenceOp{BumpTo: xdr.SequenceNumber(bs.BumpTo)}
	body, err := xdr.NewOperationBody(opType, xdrOp)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "failed to build XDR OperationBody")
	}
	op := xdr.Operation{Body: body}
	SetOpSourceAccount(&op, bs.SourceAccount)
	return op, nil
}

// FromXDR for BumpSequence initialises the txnbuild struct from the corresponding xdr Operation.
func (bs *BumpSequence) FromXDR(xdrOp xdr.Operation) error {
	result, ok := xdrOp.Body.GetBumpSequenceOp()
	if !ok {
		return errors.New("error parsing bump_sequence operation from xdr")
	}

	bs.SourceAccount = accountFromXDR(xdrOp.SourceAccount)
	bs.BumpTo = int64(result.BumpTo)
	return nil
}

// Validate for BumpSequence validates the required struct fields. It returns an error if any of the fields are
// invalid. Otherwise, it returns nil.
func (bs *BumpSequence) Validate() error {
	err := validateAmount(bs.BumpTo)
	if err != nil {
		return NewValidationError("BumpTo", err.Error())
	}
	return nil
}

// GetSourceAccount returns the source account of the operation, or the empty string if not
// set.
func (bs *BumpSequence) GetSourceAccount() string {
	return bs.SourceAccount
}
