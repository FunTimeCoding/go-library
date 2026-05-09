package constant

const (
	Name = "gopostgresd"

	ListInstances = "list_instances"
	UseInstance   = "use_instance"
	Query         = "query"
	Explain       = "explain"
	ListSchemas   = "list_schemas"
	ListTables    = "list_tables"
	DescribeTable = "describe_table"
	ListIndexes   = "list_indexes"
	TableSizes    = "table_sizes"

	ServerInstructions = "Multi-instance PostgreSQL access. Call use_instance to select a database before querying — the active instance persists across tool calls within a session. Call list_instances to see available databases."
)
