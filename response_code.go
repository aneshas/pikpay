package pikpay

const (
	Approved = 0
)

const (
	CardExpired = iota + 1001
	CardSuspicious
	CardSuspended
	CardStolen
	CardLost
)

const (
	CardNotFound = iota + 1011
	CardholderNotFound
	_
	AccountNotFound
	InvalidRequest
	InsufficientFunds
	PreviouslyReversed
	PreviouslyReversedPrime
	ActivityReversalPrevent
	ActivityVoidPrevent
	TransactionVoided
	PreauthNotAllowed
	Only3DAuthAllowed
	InstallmentsNotAllowed
	TransactionPreauthFrobidden
	InstallmentsNotAllowedZABA
)

const (
	TransactionDeclined = 1050
)

const (
	MissingFields = iota + 1802
	ExtraFields
	InvalidCardNumber
	_
	CardNotActive
	_
	CardNotConfigured
	_
	InvalidAmount
	SystemErrorDB
	SystemErrorTrn
	CardholderNotActive
	CardholderNotConfigured
	CardholderExpired
	OriginalNotFound
	UsageLimitReached
	ConfigurationError
	InvalidTerminal
	InactiveTerminal
	InvalidMerchant
	DuplicateEntity
	InvalidAcquirer
)
