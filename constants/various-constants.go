package constants

type Category string

const (
	Business      Category = "Business"
	Entertainment          = "Entertainment"
	Health                 = "Health"
	Science                = "Science"
	Sports                 = "Sports"
	Technology             = "Technology"
)


type ErrorCode string

const (
	ApiKeyExhausted ErrorCode = "apiKeyExhausted"
	ApiKeyMissing = "apiKeyMissing"
	ApiKeyInvalid = "apiKeyInvalid"
	ApiKeyDisabled = "apiKeyDisabled"
	ParametersMissing = "parametersMissing"
	ParametersIncompatible = "parametersIncompatible"
	ParameterInvalid = "parameterInvalid"
	RateLimited = "rateLimited"
	RequestTimeout = "requestTimeout"
	SourcesTooMany = "sourcesTooMany"
	SourceDoesNotExist = "sourceDoesNotExist"
	SourceUnavailableSortedBy = "sourceUnavailableSortedBy"
	SourceTemporarilyUnavailable = "sourceTemporarilyUnavailable"
	UnexpectedError = "unexpectedError"
	UnknownError = "unknownError"
)


type SortBy string

const (
	/// <summary>
	/// Sort by publisher popularity
	/// </summary>
	Popularity SortBy = "popularity"
	/// <summary>
	/// Sort by article publish date (newest first)
	/// </summary>
	PublishedAt  = "publishedAt"
	/// <summary>
	/// Sort by relevancy to the Q param
	/// </summary>
	Relevancy = "relevancy"
)
