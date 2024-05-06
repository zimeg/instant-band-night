package errors

var (
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

	// ErrPromptInterrupt denotes an interrupted prompt
	//
	// The empty message prints nothing but can be adjusted in code
	ErrPromptInterrupt = IBNError{
		Code: "err_prompt_interrupt",
	}
)
