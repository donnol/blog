package sql

type Builder interface {
	Build() (string, []any)
}

// select * from table where field = value
func Select(cols ...string) *selectBuilder {
	b := &selectBuilder{cols: cols}
	return b
}

// insert into table (columns) values (values)
func Insert() *insertBuilder {
	return nil
}

// update table set columns = values where field = value
func Update() *updateBuilder {
	return nil
}

// delete from table where field = value
func Delete() *deleteBuilder {
	return nil
}

type selectBuilder struct {
	cols  []string
	table string
	conds []cond
}

type cond struct {
	operate string
	field   string
	value   any
}

func (s *selectBuilder) From(table string) *selectBuilder {
	s.table = table
	return s
}

func (s *selectBuilder) Where(operate, field string, value any) *selectBuilder {
	s.conds = append(s.conds, cond{
		operate: operate,
		field:   field,
		value:   value,
	})
	return s
}

func (s *selectBuilder) Build() (string, []any) {
	return "", nil
}

type insertBuilder struct {
}

func (s *insertBuilder) Build() (string, []any) {
	return "", nil
}

type updateBuilder struct {
}

func (s *updateBuilder) Build() (string, []any) {
	return "", nil
}

type deleteBuilder struct {
}

func (s *deleteBuilder) Build() (string, []any) {
	return "", nil
}
