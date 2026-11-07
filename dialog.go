package widgets

import (
	"github.com/guigui-gui/guigui"
	"github.com/hajimehoshi/dialog"
)

// MessageDialog is a widget that displays a message dialog.
type MessageDialog struct {
	guigui.DefaultWidget
	init                bool
	dialog              dialog.MsgBuilder
	yesno_c             chan bool
	on_yes_no_result_fn func(context *guigui.Context, value bool)
}

// Sets the title of the dialog.
func (d *MessageDialog) Title(title string) {
	d.dialog.Dlg.Title = title
}

// Sets the message of the dialog.
func (d *MessageDialog) Message(message string) {
	d.dialog.Msg = message
}

// Opens the error dialog.
func (d *MessageDialog) Error() {
	go d.dialog.Error()
}

// Opens the info dialog.
func (d *MessageDialog) Info() {
	go d.dialog.Info()
}

// Opens the yes/no dialog.
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

// DirectoryDialog is a widget that opens a directory dialog.
type DirectoryDialog struct {
	guigui.DefaultWidget
	init                bool
	dialog              dialog.DirectoryBuilder
	on_browse_result_fn func(context *guigui.Context, file_path string, err error)
	browse_result_c     chan browse_results
}

// Sets the title of the directory dialog.
func (b *DirectoryDialog) Title(title string) {
	b.dialog.Dlg.Title = title
}

// Sets the start directory of the directory dialog.
func (b *DirectoryDialog) SetStartDir(dir string) {
	b.dialog.StartDir = dir
}

// Opens the directory dialog.
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

// FileDialog is a widget that opens a file dialog.
type FileDialog struct {
	guigui.DefaultWidget
	init   bool
	dialog *dialog.FileBuilder

	on_loadsave_result_fn func(context *guigui.Context, file_path string, err error)
	loadsave_c            chan browse_results
}

// Sets the title of the file dialog.
func (d *FileDialog) Title(title string) {
	d.dialog = d.dialog.Title(title)
}

// Filter adds a category of files to the types allowed by the dialog. Multiple
// calls to Filter are cumulative - any of the provided categories will be allowed.
// By default all files can be selected.
//
// The special extension '*' allows all files to be selected when the Filter is active.
func (d *FileDialog) Filter(desc string, extensions ...string) {
	d.dialog = d.dialog.Filter(desc, extensions...)
}

// SetStartDir sets the directory that the file dialog will open in.
func (d *FileDialog) SetStartDir(startDir string) {
	d.dialog = d.dialog.SetStartDir(startDir)
}

// SetStartFile sets the file that the file dialog will open with.
func (d *FileDialog) SetStartFile(startFile string) {
	d.dialog = d.dialog.SetStartFile(startFile)
}

// Load opens the file dialog in load mode.
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

// Save opens the file dialog in save mode.
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
