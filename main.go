package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/test"
	"github.com/ddkwork/golibrary/src/fynelib/fyneTheme"
	"wirwl/input"
	"wirwl/packets"
	"wirwl/xwidget"
)

func main() {
	t := createDefaultTableForTesting()
	//t.AddRow(xwidget.TableRow{})
	a := app.NewWithID("com.mitmproxy")
	fyneTheme.Dark()
	//a.SetIcon(fyne.NewStaticResource("mitm", asserts.MitmBuf))
	fyneTheme.Dark()
	w := a.NewWindow("mitmproxy")
	//w.Resize(fyne.NewSize(2000, 680))
	w.Resize(fyne.NewSize(300, 400))
	w.SetMaster()
	w.CenterOnScreen()
	w.SetContent(t)
	w.ShowAndRun()
}

const (
	testLabelWidth  = 20
	testLabelHeight = 10
)

func createDefaultTableForTesting() *xwidget.Table {
	table := xwidget.NewTable(test.Canvas(), setInputHandler(),
		makeHeader(), mockRows())
	return table
}

func setInputHandler() input.Handler {
	keymap := make(map[input.Action]input.KeyCombination)
	keymap[input.MoveDownAction] = input.SingleKeyCombination(fyne.KeyJ)
	keymap[input.MoveUpAction] = input.SingleKeyCombination(fyne.KeyK)
	keymap[input.EnterInputModeAction] = input.SingleKeyCombination(fyne.KeyI)
	keymap[input.ExitInputModeAction] = input.SingleKeyCombination(fyne.KeyEscape)
	keymap[input.ExitTableAction] = input.SingleKeyCombination(fyne.KeySpace)
	keymap[input.ConfirmAction] = input.SingleKeyCombination(fyne.KeyReturn)
	keymap[input.CancelAction] = input.SingleKeyCombination(fyne.KeyEscape)
	return input.NewHandler(keymap)
}

func makeHeader() []xwidget.TableColumn {
	var data []xwidget.TableColumn
	header := Header()
	for _, s := range header {
		column := xwidget.TableColumn{Type: xwidget.TextColumn, Name: s}
		data = append(data, column)
	}
	return data
}

func Header() []string {
	return []string{
		packets.NamePacketField.Method(),
		packets.NamePacketField.Scheme(),
		packets.NamePacketField.Host(),
		packets.NamePacketField.Path(),
		packets.NamePacketField.ContentType(),
		packets.NamePacketField.ContentLength(),
		packets.NamePacketField.Status(),
		packets.NamePacketField.Notes(),
		packets.NamePacketField.Process(),
		packets.NamePacketField.PadTime(),
	}
}
