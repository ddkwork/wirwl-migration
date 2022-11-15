package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

const HeaderWidth float32 = 30.0
const RowHeight float32 = 22.0
const widthBetweenColumns float32 = 90.0

type tableRenderer struct {
	table           *Table
	headerRowBorder *canvas.Rectangle
	dataRowsBorders []*canvas.Rectangle
	columnBorders   []*canvas.Rectangle
	focusedBorder   *canvas.Rectangle
	borderColor     color.Color
}

func newTableRenderer(table *Table) *tableRenderer {
	dataRowsBorders := createBorders(len(table.rowData))
	return &tableRenderer{
		table:           table,
		headerRowBorder: canvas.NewRectangle(color.Black),
		dataRowsBorders: dataRowsBorders,
		columnBorders:   createBorders(table.columnAmount()),
		focusedBorder:   canvas.NewRectangle(color.Transparent),
		borderColor:     color.Black,
	}
}

func createBorders(amount int) []*canvas.Rectangle {
	var borders []*canvas.Rectangle
	for i := 1; i <= amount; i++ {
		borders = append(borders, canvas.NewRectangle(color.Black))
	}
	return borders
}

func (renderer *tableRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor() //todo add ohther
}

func (renderer *tableRenderer) Destroy() {
	//No resources to clear
}

func (renderer *tableRenderer) Layout(_ fyne.Size) {
	renderer.renderHeader()
	renderer.renderData()
	renderer.renderFocusedBorder()
}

func (renderer *tableRenderer) renderHeader() {
	renderer.renderHeaderColumnLabels()
	renderer.renderHeaderRowBorder()
}

func (renderer *tableRenderer) renderHeaderColumnLabels() {
	position := fyne.NewPos(widthBetweenColumns/2, 0)
	for _, columnLabel := range renderer.table.columnLabels {
		label := columnLabel.(*widget.Label)
		label.TextStyle.Bold = true
		label.Move(position)
		size := fyne.NewSize(label.MinSize().Width, HeaderWidth)
		label.Resize(size)
		position = position.Add(fyne.NewPos(size.Width+widthBetweenColumns, 0))
	}
}

func (renderer *tableRenderer) renderHeaderRowBorder() {
	renderer.headerRowBorder.Move(fyne.NewPos(0, 0))
	headerRowRectangleSize := fyne.NewSize(renderer.tableWidth(), HeaderWidth)
	renderer.setBorderProperties(renderer.headerRowBorder)
	renderer.headerRowBorder.Resize(headerRowRectangleSize)
}

// Should only be called after header column labels have been rendered, otherwise width will be wrong
func (renderer *tableRenderer) tableWidth() float32 {
	var tableWidth float32 = 0.0
	for _, columnLabel := range renderer.table.columnLabels {
		tableWidth += columnLabel.Size().Width + widthBetweenColumns
	}
	return tableWidth
}

func (renderer *tableRenderer) tableHeight() float32 {
	//All data rows have the same height
	return HeaderWidth + float32(len(renderer.table.rowData))*RowHeight
}

func (renderer *tableRenderer) renderData() {
	renderer.renderCellsContent()
	renderer.renderDataRowsBorders()
	renderer.renderColumnBorders()
}

func (renderer *tableRenderer) renderCellsContent() {
	position := fyne.NewPos(widthBetweenColumns/2, HeaderWidth)
	for _, row := range renderer.table.rowData {
		size := fyne.NewSize(0, RowHeight)
		for i, cellContent := range row {
			columnWidth := renderer.table.columnLabels[i].Size().Width
			size := fyne.NewSize(columnWidth, RowHeight)
			cellContent.Resize(size)
			cellContent.Move(position)
			renderer.setContentProperties(cellContent, renderer.table.columnData[i].Type)
			position = position.Add(fyne.NewPos(size.Width+widthBetweenColumns, 0))
		}
		position = position.Subtract(fyne.NewPos(position.X, 0))
		position = position.Add(fyne.NewPos(widthBetweenColumns/2, size.Height))
	}
}

func (renderer *tableRenderer) setContentProperties(content fyne.CanvasObject, columnType ColumnType) {
	switch columnType {
	case TextColumn:
		contentLabel := content.(*widget.Label)
		contentLabel.Wrapping = fyne.TextWrapWord
		contentLabel.Alignment = fyne.TextAlignCenter
	}
}

func (renderer *tableRenderer) renderDataRowsBorders() {
	size := fyne.NewSize(renderer.tableWidth(), RowHeight)
	position := fyne.NewPos(0, HeaderWidth)
	for _, border := range renderer.dataRowsBorders {
		border.Move(position)
		border.Resize(size)
		renderer.setBorderProperties(border)
		position = position.Add(fyne.NewPos(0, RowHeight))
	}
}

func (renderer *tableRenderer) renderColumnBorders() {
	position := fyne.NewPos(0, 0)
	for columnNum, border := range renderer.columnBorders {
		columnWidth := renderer.table.columnLabels[columnNum].Size().Width + widthBetweenColumns
		size := fyne.NewSize(columnWidth, renderer.tableHeight())
		border.Move(position)
		border.Resize(size)
		renderer.setBorderProperties(border)
		position = position.Add(fyne.NewPos(columnWidth, 0))
	}
}

func (renderer *tableRenderer) setBorderProperties(border *canvas.Rectangle) {
	border.StrokeWidth = 2
	border.FillColor = color.Transparent
	border.StrokeColor = renderer.borderColor
}

func (renderer *tableRenderer) renderFocusedBorder() {
	if renderer.table.focused {
		renderer.focusedBorder.Show()
	} else {
		renderer.focusedBorder.Hide()
	}
	renderer.focusedBorder.StrokeWidth = 3
	renderer.focusedBorder.FillColor = color.Transparent
	renderer.focusedBorder.StrokeColor = theme.FocusColor()
	renderer.focusedBorder.Move(fyne.NewPos(0, 0))
	size := fyne.NewSize(renderer.tableWidth(), renderer.tableHeight())
	renderer.focusedBorder.Resize(size)
}

func (renderer *tableRenderer) MinSize() fyne.Size {
	return fyne.NewSize(renderer.tableWidth(), renderer.tableHeight())
}

func (renderer *tableRenderer) Objects() []fyne.CanvasObject {
	objects := []fyne.CanvasObject{}
	for _, row := range renderer.table.rowData {
		objects = append(objects, row...)
	}
	objects = append(objects, renderer.table.columnLabels...)
	objects = append(objects, renderer.headerRowBorder)
	objects = append(objects, convertRectanglesToCanvasObjects(renderer.dataRowsBorders)...)
	objects = append(objects, convertRectanglesToCanvasObjects(renderer.columnBorders)...)
	objects = append(objects, renderer.focusedBorder)
	return objects
}

func convertRectanglesToCanvasObjects(rectangles []*canvas.Rectangle) []fyne.CanvasObject {
	objects := []fyne.CanvasObject{}
	for _, rectangle := range rectangles {
		objects = append(objects, rectangle)
	}
	return objects
}

func (renderer *tableRenderer) Refresh() {
	renderer.recreateDataBorders()
	//The size can be anything as it is ignored by renderer
	renderer.Layout(fyne.NewSize(0, 0))
}

func (renderer *tableRenderer) recreateDataBorders() {
	renderer.dataRowsBorders = createBorders(len(renderer.table.rowData))
}
