package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
	"image/color"
	"strconv"
	"testing"
)

func TestThatHeaderRowBorderIsDrawnCorrectly(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	rectangle := renderer.headerRowBorder
	assert.Equal(t, expectedTableWidth, rectangle.Size().Width)
	assert.Equal(t, expectedHeaderHeight, rectangle.Size().Height)
	assert.Equal(t, float32(2), rectangle.StrokeWidth)
	assert.Equal(t, color.Black, rectangle.StrokeColor)
	assert.Equal(t, color.Transparent, rectangle.FillColor)
}

func TestThatThereIsCorrectAmountOfDataRowBorders(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	assert.Equal(t, testRowAmount, len(renderer.dataRowsBorders))
}

func TestThatAllDataRowBordersHaveCorrectSize(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	for i, border := range renderer.dataRowsBorders {
		assert.Equal(t, expectedTableWidth, border.Size().Width, "Border with number "+strconv.Itoa(i)+" does not have the correct width")
		assert.Equal(t, expectedRowHeight, border.Size().Height, "Border with number "+strconv.Itoa(i)+" does not have the correct height")
	}
}

func TestThatAllDataBordersHaveCorrectPosition(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	expectedPosition := fyne.NewPos(0, expectedHeaderHeight)
	for i, border := range renderer.dataRowsBorders {
		assert.Equal(t, expectedPosition, border.Position(), "Border with number "+strconv.Itoa(i)+" does not have correct position")
		expectedPosition = expectedPosition.Add(fyne.NewPos(0, expectedRowHeight))
	}
}

func TestThatAllDataBordersHaveCorrectProperties(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	for i, border := range renderer.dataRowsBorders {
		assert.Equal(t, float32(2), border.StrokeWidth, "Border with number "+strconv.Itoa(i)+" does not have correct stroke width")
		assert.Equal(t, color.Black, border.StrokeColor, "Border with number "+strconv.Itoa(i)+" does not have correct stroke color")
		assert.Equal(t, color.Transparent, border.FillColor, "Border with number "+strconv.Itoa(i)+" does not have correct fill color")
	}
}

func TestThatThereIsCorrectAmountOfColumnBorders(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	assert.Equal(t, testColumnAmount, len(renderer.columnBorders))
}

func TestThatAllColumnBordersHaveCorrectSize(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	for i, border := range renderer.columnBorders {
		assert.Equal(t, renderer.table.columnLabels[i].Size().Width+expectedPadding, border.Size().Width, "Border with number "+strconv.Itoa(i)+" does not have the correct width")
		assert.Equal(t, expectedHeaderHeight+expectedRowHeight*testRowAmount, border.Size().Height, "Border with number "+strconv.Itoa(i)+" does not have the correct height")
	}
}

func TestThatAllColumnBordersHaveCorrectPosition(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	expectedPosition := fyne.NewPos(0, 0)
	for i, border := range renderer.columnBorders {
		assert.Equal(t, expectedPosition, border.Position(), "Border with number "+strconv.Itoa(i)+" does not have correct position")
		expectedPosition = expectedPosition.Add(fyne.NewPos(renderer.table.columnLabels[i].Size().Width+expectedPadding, 0))
	}
}

func TestThatAllColumnBordersHaveCorrectProperties(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	for i, border := range renderer.columnBorders {
		assert.Equal(t, float32(2), border.StrokeWidth, "Border with number "+strconv.Itoa(i)+" does not have correct stroke width")
		assert.Equal(t, color.Black, border.StrokeColor, "Border with number "+strconv.Itoa(i)+" does not have correct stroke color")
		assert.Equal(t, color.Transparent, border.FillColor, "Border with number "+strconv.Itoa(i)+" does not have correct fill color")
	}
}

func TestThatFocusedBorderHasCorrectProperties(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	assert.Equal(t, expectedTableWidth, renderer.focusedBorder.Size().Width)
	assert.Equal(t, expectedTableHeight, renderer.focusedBorder.Size().Height)
	assert.Equal(t, float32(3), renderer.focusedBorder.StrokeWidth)
	assert.Equal(t, theme.FocusColor(), renderer.focusedBorder.StrokeColor)
	assert.Equal(t, color.Transparent, renderer.focusedBorder.FillColor)
}

func TestThatFocusedBorderHidesAndShowsCorrectly(t *testing.T) {
	renderer := createDefaultTableRendererForTesting()
	assert.True(t, renderer.focusedBorder.Hidden)
	renderer.table.EnterInputMode()
	assert.True(t, renderer.focusedBorder.Visible())
	renderer.table.ExitInputMode()
	assert.True(t, renderer.focusedBorder.Hidden)
}

func TestThatAfterRowIsAddedThereIsCorrectAmountOfRowBorders(t *testing.T) {
	table := createDefaultTableForTesting()
	table.AddRow(createTestTableRow(testColumnAmount))
	renderer := test.WidgetRenderer(table).(*tableRenderer)
	assert.Equal(t, testRowAmount+1, len(renderer.dataRowsBorders))
}
