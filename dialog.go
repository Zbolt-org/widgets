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

type browse_results struct {
	file_path string
	err       error
}

type DirectoryDialog struct {
	guigui.DefaultWidget
	init                bool
	dialog              dialog.DirectoryBuilder
	on_browse_result_fn func(context *guigui.Context, file_path string, err error)
	browse_result_c     chan browse_results
}

func (b *DirectoryDialog) Title(title string) {
	b.dialog.Dlg.Title = title
}

func (b *DirectoryDialog) SetStartDir(dir string) {
	b.dialog.StartDir = dir
}

func (b *DirectoryDialog) Browse(fn func(context *guigui.Context, file_path string, err error)) {
	b.on_browse_result_fn = fn
	go func() {
		file_path, err := b.dialog.Browse()
		b.browse_result_c <- browse_results{
			file_path: file_path,
			err:       err,
		}
	}()
}

func (d *DirectoryDialog) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if !d.init {
		d.browse_result_c = make(chan browse_results, 1)
		d.init = true
	}

	select {
	case val := <-d.browse_result_c:
		if d.on_browse_result_fn != nil {
			d.on_browse_result_fn(context, val.file_path, val.err)
		}
	default:
	}
	return nil
}

type FileDialog struct {
	guigui.DefaultWidget
	init   bool
	dialog *dialog.FileBuilder

	on_loadsave_result_fn func(context *guigui.Context, file_path string, err error)
	loadsave_c            chan browse_results
}

func (d *FileDialog) Title(title string) {
	d.dialog = d.dialog.Title(title)
}

func (d *FileDialog) Filter(desc string, extensions ...string) {
	d.dialog = d.dialog.Filter(desc, extensions...)
}

func (d *FileDialog) SetStartDir(startDir string) {
	d.dialog = d.dialog.SetStartDir(startDir)
}

func (d *FileDialog) SetStartFile(startFile string) {
	d.dialog = d.dialog.SetStartFile(startFile)
}

func (d *FileDialog) Load(fn func(context *guigui.Context, file_path string, err error)) {
	d.on_loadsave_result_fn = fn
	go func() {
		file_path, err := d.dialog.Load()
		d.loadsave_c <- browse_results{
			file_path: file_path,
			err:       err,
		}
	}()
}

func (d *FileDialog) Save(fn func(context *guigui.Context, file_path string, err error)) {
	d.on_loadsave_result_fn = fn
	go func() {
		file_path, err := d.dialog.Save()
		d.loadsave_c <- browse_results{
			file_path: file_path,
			err:       err,
		}
	}()
}

func (d *FileDialog) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if !d.init {
		d.loadsave_c = make(chan browse_results, 1)
		d.init = true
	}

	select {
	case val := <-d.loadsave_c:
		if d.on_loadsave_result_fn != nil {
			d.on_loadsave_result_fn(context, val.file_path, val.err)
		}
	default:
	}
	return nil
}
