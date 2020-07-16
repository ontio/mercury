package packager

type Packager interface {
	// PackMessage Pack a message for one or more recipients.
	//
	// Args:
	//
	// envelope: The message to pack
	//
	// Returns:
	//
	// []byte: The packed message
	//
	// error: error
	PackMessage(envelope *Envelope) ([]byte, error)

	// UnpackMessage Unpack a message.
	//
	// Args:
	//
	// encMessage: The encrypted message
	//
	// Returns:
	//
	// envelope: unpack message
	//
	// error: error
	UnpackMessage(encMessage []byte) (*Envelope, error)
}
