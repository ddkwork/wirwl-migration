package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"net/http"
	"strconv"
	"wirwl/packets"
	"wirwl/xwidget"
)

func fmtRow(row packets.Row) xwidget.TableRow {
	r := xwidget.TableRow{
		widget.NewLabel(row.Method),
		widget.NewLabel(row.Scheme),
		widget.NewLabel(row.Host),
		widget.NewLabel(row.Path),
		widget.NewLabel(row.ContentType),
		widget.NewLabel(fmt.Sprint(row.ContentLength)),
		widget.NewLabel(row.Status),
		widget.NewLabel(row.Note),
		widget.NewLabel(fmt.Sprint(row.PadTime)),
	}
	for i := 0; i < len(r); i++ {
		r[i].Resize(fyne.NewSize(xwidget.HeaderWidth, xwidget.RowHeight)) //not used ?
	}
	return r
}

func mockRows() []xwidget.TableRow {
	var labels []xwidget.TableRow
	for i := 0; i < 20; i++ {
		row := fmtRow(packets.Row{
			Method:        http.MethodConnect,
			Scheme:        "wss",
			Host:          "www.baidu.com",
			Path:          "cmsocket",
			ContentType:   "json",
			ContentLength: int64(i),
			Status:        "ok",
			Note:          "aes key",
			Process:       "steam.exe",
			PadTime:       10,
		})
		label := widget.NewLabel("row" + strconv.Itoa(i))
		label.Resize(fyne.NewSize(testLabelWidth, testLabelHeight))
		labels = append(labels, row)
	}
	return labels
}
