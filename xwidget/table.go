package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	input2 "wirwl/input"
)

type Table struct {
	widget.BaseWidget
	inputHandler input2.Handler
	columnData   []TableColumn
	columnLabels []fyne.CanvasObject
	rowData      []TableRow
	canvas       fyne.Canvas
	focused      bool
	onExit       func()
}

type TableColumn struct {
	Type ColumnType
	Name string
}

type ColumnType string

const (
	TextColumn  ColumnType = "TEXT_COLUMN"
	ImageColumn ColumnType = "IMAGE_COLUMN"
)

type TableRow []fyne.CanvasObject

func NewTable(canvas fyne.Canvas, inputHandler input2.Handler, columnData []TableColumn, rowData []TableRow) *Table {
	table := &Table{
		inputHandler: inputHandler,
		columnData:   columnData,
		columnLabels: createColumnLabels(columnData),
		rowData:      rowData,
		canvas:       canvas,
		focused:      false,
	}
	table.ExtendBaseWidget(table)
	table.inputHandler.BindFunctionToAction(table, input2.ExitTableAction, func() { table.onExit() })
	return table
}

func createColumnLabels(columnData []TableColumn) []fyne.CanvasObject {
	var labels []fyne.CanvasObject
	for _, column := range columnData {
		labels = append(labels, widget.NewLabel(column.Name))
	}
	return labels
}

func (table *Table) HeaderColumns() []fyne.CanvasObject {
	return table.columnLabels
}

func (table *Table) CreateRenderer() fyne.WidgetRenderer {
	return newTableRenderer(table)
}

func (table *Table) columnAmount() int {
	return len(table.columnData)
}

func (table *Table) FocusGained() {
	table.focused = true
}

func (table *Table) FocusLost() {
	table.focused = false
}

func (table *Table) Focused() bool {
	return table.focused
}

func (table *Table) TypedRune(rune) {
	//Table will not support any sort of typing therefore no implementation is needed
}

func (table *Table) TypedKey(event *fyne.KeyEvent) {
	table.inputHandler.HandleInNormalMode(table, event.Name)
}

func (table *Table) EnterInputMode() {
	table.canvas.Focus(table)
	table.Refresh()
}

func (table *Table) ExitInputMode() {
	table.canvas.Unfocus()
	table.Refresh()
}

func (table *Table) AddRow(row TableRow) {
	table.rowData = append(table.rowData, row)
	table.Refresh()
}

func (table *Table) SetOnExitCallbackFunction(function func()) {
	table.onExit = function
}
