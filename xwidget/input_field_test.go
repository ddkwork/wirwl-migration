package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
	"testing"
	input2 "wirwl/input"
)

func TestThatFunctionGetsCalledOnConfirm(t *testing.T) {
	functionExecuted := false
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	inputField.SetOnConfirm(func() { functionExecuted = true })
	inputField.canvas.Focus(inputField)
	SimulateKeyPressOnTestCanvas(fyne.KeyReturn)
	assert.Equal(t, true, functionExecuted)
}

func TestThatFunctionGetsCalledOnCancel(t *testing.T) {
	functionExecuted := false
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	inputField.SetOnExitInputModeFunction(func() { functionExecuted = true })
	inputField.canvas.Focus(inputField)
	SimulateKeyPressOnTestCanvas(fyne.KeyEscape)
	assert.True(t, functionExecuted)
}

func TestThatTypingWorks(t *testing.T) {
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	inputField.FocusLost()
	inputField.FocusGained()
	inputField.Type("some value")
	assert.Equal(t, "some value", inputField.Text)
}

func TestThatFunctionsAreNotNil(t *testing.T) {
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	assert.NotNil(t, inputField.OnConfirm)
}

func TestEnteringIntoInputMode(t *testing.T) {
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	inputField.EnterInputMode()
	assert.Equal(t, inputField.bgRenderer.BackgroundColor(), theme.BackgroundColor())
	assert.Equal(t, inputField, test.Canvas().Focused())
	assert.Equal(t, inputField, inputField.canvas.Focused())
}

func TestHighlightingAndUnhiglighting(t *testing.T) {
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	//InputField needs to be placed into a test window, otherwise renderer doesn't work properly and marking sets background color again
	test.NewApp().NewWindow("").SetContent(inputField)
	assert.Equal(t, inputField.bgRenderer.BackgroundColor(), theme.BackgroundColor())
	inputField.Highlight()
	assert.Equal(t, inputField.bgRenderer.BackgroundColor(), theme.FocusColor())
	inputField.Unhighlight()
	assert.Equal(t, inputField.bgRenderer.BackgroundColor(), theme.BackgroundColor())
}

func TestThatWhenExitingInputModeWithTwoKeyCombinationNeitherKeyOfCombinationGetsLeftInFieldsText(t *testing.T) {
	keymap := make(map[input2.Action]input2.KeyCombination)
	keymap[input2.ExitInputModeAction] = input2.TwoKeyCombination(fyne.KeyJ, fyne.KeyJ)
	inputHandler := input2.NewHandler(keymap)
	inputField := NewInputField(test.Canvas(), inputHandler)
	inputField.canvas.Focus(inputField)
	inputField.Type("abc")
	SimulateKeyPressOnTestCanvas(fyne.KeyJ)
	SimulateKeyPressOnTestCanvas(fyne.KeyJ)
	assert.Equal(t, "abc", inputField.Text)
}

func TestThatInputFieldDoesNotInputIfFilterFunctionIsSet(t *testing.T) {
	inputField := NewInputField(test.Canvas(), getInputHandlerForTesting())
	inputField.SetRuneFilteringFunction(func(r rune) bool {
		return r == 'a'
	})
	TypeIntoFocusable(inputField, "a5pi25a1!'';.Aa")
	assert.Equal(t, "aaa", inputField.Text)
}
