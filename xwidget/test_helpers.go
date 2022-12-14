package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	input2 "wirwl/input"
)

/*
	When opening input in the running application, the last typed in character is still handled which normally means it

gets typed into the input. There is code that prevents this but as this situation doesn't happen when running the
test code any string typed into the input needs an additional character at the beginning as if the bug happened.
*/
func TypeIntoFocusable(focusable fyne.Focusable, chars string) {
	fixedChars := " " + chars
	for _, char := range fixedChars {
		focusable.TypedRune(char)
	}
}

func (inputField *InputField) Type(chars string) {
	TypeIntoFocusable(inputField, chars)
}

func (dialog *FormDialog) Type(chars string) {
	TypeIntoFocusable(dialog.currentWidget(), chars)
}

func SimulateKeyPressOnTestCanvas(key fyne.KeyName) {
	SimulateKeyPressOn(test.Canvas(), key)
}

/*
Should behave mostly the same as the actual key presses in the application except it shouldn't be used for testing
typing in some text that has keys like Space, Escape etc. For that just use TypeIntoFocusable.
Overall, this should be used for testing key presses that are supposed to trigger some kind of actions.
*/
func SimulateKeyPressOn(canvas fyne.Canvas, key fyne.KeyName) {
	event := &fyne.KeyEvent{Name: key}
	focused := canvas.Focused()
	if focused != nil {
		focused.TypedKey(event)
	}
	//Focused object has to be retrieved again in case it has been unfocused somewhere after key was pressed
	focused = canvas.Focused()
	if focused != nil {
		/*This is a very naive way of retrieving a character to type in and will work incorrectly for e.g. Escape key
		which will send 'E' character to the focused object */
		character := []rune(event.Name)[0]
		focused.TypedRune(character)
	}
}

func ContainsWidget(content fyne.CanvasObject, searchedWidget interface{}) bool {
	for _, existingWidget := range content.(*fyne.Container).Objects {
		if existingWidget == searchedWidget {
			return true
		}
	}
	return false
}

func ContainsLabelWithSameText(content fyne.CanvasObject, searchedText string) bool {
	for _, existingWidget := range content.(*fyne.Container).Objects {
		if existingWidget.(*widget.Label).Text == searchedText {
			return true
		}
	}
	return false
}

func GetLabelFromContent(content fyne.CanvasObject, labelText string) *widget.Label {
	for _, existingWidget := range content.(*fyne.Container).Objects {
		if existingWidget.(*widget.Label).Text == labelText {
			return existingWidget.(*widget.Label)
		}
	}
	return nil
}

func GetLabelPositionInContent(content fyne.CanvasObject, labelText string) int {
	for position, existingWidget := range content.(*fyne.Container).Objects {
		if existingWidget.(*widget.Label).Text == labelText {
			return position
		}
	}
	return -1
}

func getInputHandlerForTesting() input2.Handler {
	keymap := make(map[input2.Action]input2.KeyCombination)
	//Default keys are the same as if they were set by default config
	keymap[input2.MoveDownAction] = input2.SingleKeyCombination(fyne.KeyJ)
	keymap[input2.MoveUpAction] = input2.SingleKeyCombination(fyne.KeyK)
	keymap[input2.EnterInputModeAction] = input2.SingleKeyCombination(fyne.KeyI)
	keymap[input2.ExitInputModeAction] = input2.SingleKeyCombination(fyne.KeyEscape)
	keymap[input2.ExitTableAction] = input2.SingleKeyCombination(fyne.KeySpace)
	keymap[input2.ConfirmAction] = input2.SingleKeyCombination(fyne.KeyReturn)
	keymap[input2.CancelAction] = input2.SingleKeyCombination(fyne.KeyEscape)
	return input2.NewHandler(keymap)
}

func getOneInputFieldForDialogTesting() []*FormDialogFormItem {
	var elements []*FormDialogFormItem
	elements = append(elements, newFormDialogFormItem("first", NewInputField(test.Canvas(), getInputHandlerForTesting())))
	return elements
}

func getTwoInputFieldsForFormDialogTesting() []*FormDialogFormItem {
	var elements []*FormDialogFormItem
	elements = append(elements, newFormDialogFormItem("first", NewInputField(test.Canvas(), getInputHandlerForTesting())))
	elements = append(elements, newFormDialogFormItem("second", NewInputField(test.Canvas(), getInputHandlerForTesting())))
	return elements
}

func getThreeInputFieldsForFormDialogTesting() []*FormDialogFormItem {
	elements := []*FormDialogFormItem{}
	elements = append(elements, newFormDialogFormItem("first", NewInputField(test.Canvas(), getInputHandlerForTesting())))
	elements = append(elements, newFormDialogFormItem("second", NewInputField(test.Canvas(), getInputHandlerForTesting())))
	elements = append(elements, newFormDialogFormItem("third", NewInputField(test.Canvas(), getInputHandlerForTesting())))
	return elements
}
