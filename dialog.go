package widgets

import (
	"github.com/guigui-gui/guigui"
	"github.com/hajimehoshi/dialog"
)

type MessageDialog struct {
	guigui.DefaultWidget
	init                bool
	dialog              dialog.MsgBuilder
	yesno_c             chan bool
	on_yes_no_result_fn func(context *guigui.Context, value bool)
}

func (d *MessageDialog) Title(title string) {
	d.dialog.Dlg.Title = title
}

func (d *MessageDialog) Message(message string) {
	d.dialog.Msg = message
}

func (d *MessageDialog) Error() {
	go d.dialog.Error()
}

func (d *MessageDialog) Info() {
	go d.dialog.Info()
}

// Must wait until last YesNo dialog finishes before calling.
func (d *MessageDialog) YesNo(fn func(context *guigui.Context, value bool)) {
	d.on_yes_no_result_fn = fn
	go func() {
		value := d.dialog.YesNo()
		d.yesno_c <- value
	}()
}

func (d *MessageDialog) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if !d.init {
		d.yesno_c = make(chan bool, 1)
		d.init = true
	}

	select {
	case val := <-d.yesno_c:
		if d.on_yes_no_result_fn != nil {
			d.on_yes_no_result_fn(context, val)
		}
	default:
	}
	return nil
}

type BrowseDialog struct {
	guigui.DefaultWidget
	init bool
}

func (d *BrowseDialog) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	return nil
}

type FileDialog struct {
	guigui.DefaultWidget
	init bool
}

func (d *FileDialog) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	return nil
}
