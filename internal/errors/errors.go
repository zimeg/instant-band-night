package errors

var (
	// ErrBandNotFound denotes the band failed to be found
	ErrBandNotFound = IBNError{
		Code:    "err_band_not_found",
		Message: "The band of that order cannot be found",
	}

	// ErrConfusedFlags denotes flag values failed to be parsed
	ErrConfusedFlags = IBNError{
		Code:    "err_confused_flags",
		Message: "The flag values provided raise this error",
	}

	// ErrNotEnoughMusician denotes to few musicians remain
	ErrNotEnoughMusician = IBNError{
		Code:    "err_not_enough_musician",
		Message: "Not enough musician remains to create a band",
	}

	// ErrMissingFile denotes a missing file
	ErrMissingFile = IBNError{
		Code:    "err_missing_file",
		Message: "No matching file was found",
	}

	// ErrMissingInputMusicianInstrument denotes a missing instrument selection
	ErrMissingInputInstrument = IBNError{
		Code:    "err_missing_input_instrument",
		Message: "No instrument was provided",
	}

	// ErrMissingInputMusicianInstrument denotes a missing instrument selection
	ErrMissingInputInstrumentCount = IBNError{
		Code:    "err_missing_input_instrument_count",
		Message: "No instruments were counted",
	}

	// ErrMissingInputMusicianInstrument denotes a missing instrument selection
	ErrMissingInputMusicianInstrument = IBNError{
		Code:    "err_missing_input_musician_instrument",
		Message: "No instrument was picked",
	}

	// ErrMissingInputMusicianName denotes a missing name selection
	ErrMissingInputMusicianName = IBNError{
		Code:    "err_missing_input_musician_name",
		Message: "No name was signed",
	}

	// ErrMusicianNotFound denotes the musician failed to be found
	ErrMusicianNotFound = IBNError{
		Code:    "err_musician_not_found",
		Message: "The musician with that ID cannot be found",
	}

	// ErrPromptInterrupt denotes an interrupted prompt
	//
	// The empty message prints nothing but can be adjusted in code
	ErrPromptInterrupt = IBNError{
		Code: "err_prompt_interrupt",
	}
)
