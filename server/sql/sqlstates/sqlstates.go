package sqlstates

const (
	// Integrity Constraint Violation (Class 23)
	UniqueViolation     = "23505" // Unique constraint violation
	ForeignKeyViolation = "23503" // Foreign key constraint violation
	CheckViolation      = "23514" // Check constraint violation
	NotNullViolation    = "23502" // Not-null constraint violation
	ExclusionViolation  = "23P01" // Exclusion constraint violation

	// Invalid Transaction State (Class 25)
	InvalidTransactionState = "25000" // Invalid transaction state

	// Insufficient Privilege (Class 28)
	InsufficientPrivilege = "42501" // Insufficient privilege

	// Syntax Error or Access Rule Violation (Class 42)
	SyntaxError        = "42601" // Syntax error
	InvalidCatalogName = "3D000" // Invalid catalog name
	InvalidSchemaName  = "3F000" // Invalid schema name

	// No Data (Class 02)
	NoData = "02000" // No data found (similar to sql.ErrNoRows)

	// Connection Exception (Class 08)
	ConnectionException    = "08000" // Connection exception
	ConnectionDoesNotExist = "08003" // Connection does not exist
	ConnectionFailure      = "08006" // Connection failure

)
