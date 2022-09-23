package constant

type SelectColumns []string

var (
	DEFAULT_SELECT_COLUMNS = SelectColumns([]string{"created_at", "updated_at"})
)
