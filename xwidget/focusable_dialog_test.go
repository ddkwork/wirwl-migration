package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatNewFocusableDialogIsHidden(t *testing.T) {
	dialog := NewFocusableDialog(test.Canvas())
	assert.Equal(t, true, dialog.Hidden)
}

func TestThatDisplayShowsFocusedFocusableDialogWithSpecifiedData(t *testing.T) {
	label := widget.NewLabel("some title")
	dialog := NewFocusableDialog(test.Canvas(), label)
	dialog.Display("some title")
	assert.True(t, dialog.Visible())
	assert.True(t, dialog.Focused())
	assert.Equal(t, "some title", dialog.Title())
	assert.True(t, ContainsWidget(dialog.Content, label))
}

func TestThatPressingAnyKeyHidesFocusableDialog(t *testing.T) {
	dialog := NewFocusableDialog(test.Canvas())
	dialog.Display("")
	SimulateKeyPressOnTestCanvas(fyne.KeyQ)
	assert.Equal(t, true, dialog.Hidden)
	assert.Equal(t, false, dialog.Focused())
	dialog = NewFocusableDialog(test.Canvas())
	dialog.Display("")
	SimulateKeyPressOnTestCanvas(fyne.Key1)
	assert.Equal(t, true, dialog.Hidden)
	assert.Equal(t, false, dialog.Focused())
}

func TestThatTitleIsDisplayedFirst(t *testing.T) {
	label := widget.NewLabel("some title")
	dialog := NewFocusableDialog(test.Canvas(), label)
	assert.Equal(t, dialog.Content.(*fyne.Container).Objects[1], label)
}

func TestThatAfterOnlyFirstHidingCallbackFunctionIsCalled(t *testing.T) {
	label := widget.NewLabel("some title")
	dialog := NewFocusableDialog(test.Canvas(), label)
	count := 0
	dialog.SetOneTimeOnHideCallback(func() {
		count++
	})
	dialog.Hide()
	assert.Equal(t, 1, count)
	dialog.Hide()
	assert.Equal(t, 1, count)
}

func TestThatFocusIsNotLostIfItWasSetInHidingCallbackFunctionWhenHidingOnKeyPress(t *testing.T) {
	input := NewInputField(test.Canvas(), getInputHandlerForTesting())
	dialog := NewFocusableDialog(test.Canvas(), input)
	dialog.SetOneTimeOnHideCallback(func() {
		test.Canvas().Focus(input)
	})
	dialog.Canvas.Focus(dialog)
	assert.NotEqual(t, input, test.Canvas().Focused())
	SimulateKeyPressOnTestCanvas(fyne.KeyE)
	assert.Equal(t, input, test.Canvas().Focused())
}
