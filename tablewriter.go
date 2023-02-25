package tablewriter

import (
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Table struct {
	t    *tablewriter.Table
	cols []int
}

type TableOption func(*Table)

func AlignRight(cols ...int) TableOption {
	return func(table *Table) {
		for _, col := range cols {
			table.cols[col] = tablewriter.ALIGN_RIGHT
		}
	}
}

func NewTable(writer io.Writer, header []string, opts ...TableOption) *Table {
	t := tablewriter.NewWriter(writer)
	t.SetHeader(header)
	t.SetBorders(tablewriter.Border{Top: true, Bottom: true})
	t.SetAutoWrapText(false)
	t.SetAutoFormatHeaders(true)
	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetRowSeparator("")
	t.SetHeaderLine(false)
	t.SetBorder(false)
	t.SetTablePadding("\t")
	t.SetNoWhiteSpace(true)

	cols := make([]int, len(header))
	for i := range cols {
		cols[i] = tablewriter.ALIGN_LEFT
	}

	table := &Table{
		t,
		cols,
	}
	for _, opt := range opts {
		opt(table)
	}

	t.SetColumnAlignment(table.cols)

	return table
}

func NewConsoleTable(header []string, opts ...TableOption) *Table {
	return NewTable(os.Stdout, header, opts...)
}

func (table *Table) Write(row []string) {
	table.t.Append(row)
}

func (table *Table) Render() {
	table.t.Render()
}
