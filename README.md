# widgets

[![Go Reference](https://pkg.go.dev/badge/github.com/Zbolt-org/widgets.svg)](https://pkg.go.dev/github.com/Zbolt-org/widgets)

A collection of reusable widgets for [Guigui](https://github.com/guigui-gui/guigui), the Go GUI framework built on [Ebitengine](https://ebitengine.org/). The widgets are extracted from [Zbolt](https://github.com/udan-jayanith/Zbolt) and are meant to be dropped into any Guigui application.

> **Status:** pre-1.0. Guigui itself is alpha, so expect the API of this module to change between releases. Pin a version in your `go.mod`.
>
> **Not supported yet:** disabling a widget. See [Known limitations](#known-limitations).

<img width="940" height="807" alt="preview" src="https://github.com/user-attachments/assets/eec77cea-6430-4970-aa5f-fca8d72bc77d" />

## Contents

- [Installation](#installation)
- [Run the gallery](#run-the-gallery)
- [Available widgets](#available-widgets)
- [Known limitations](#known-limitations)
- [License](#license)

## Installation

```sh
go get github.com/Zbolt-org/widgets
```

```go
import "github.com/Zbolt-org/widgets"
```

The module builds on `github.com/guigui-gui/guigui` and `github.com/hajimehoshi/ebiten/v2`. Both are pulled in through `go.mod`. See `go.mod` for the minimum Go version.

## Run the gallery

The `gallery` directory is a small application that shows every widget in one window. It is the quickest way to see what each widget looks like and how it is wired up, and its source is the reference for the examples in this README.

```sh
git clone https://github.com/Zbolt-org/widgets.git
cd widgets
go run ./gallery
```

Run it from the repository root. The gallery loads its icons and sample image by relative path, so starting it from another directory will fail to find them.

The gallery contains:

- a toggle that switches between light and dark mode
- an icon, an info hint and an infinite loading bar
- an image wrapped in a skeleton loader, with a toggle that flips its loading state
- a text input with a context menu
- a button that adds tabs to a tab container, and a link

## Available widgets

| Widget | Use it for |
| --- | --- |
| [`LabeledToggle`](#labeledtoggle) | An on/off switch with a text label |
| [`Button`](#button) | A button with an optional icon and tooltip |
| [`TextInput`](#textinput) | A text field with a customizable context menu |
| [`Link`](#link) | Clickable text that opens a URL |
| [`Icon`](#icon) | Showing an icon, optionally clickable |
| [`InfoHint`](#infohint) | A small hint that explains something nearby |
| [`Text`](#text) | A wrapper around `basicwidget.Text` |
| [`SkeletonLoader`](#skeletonloader) | A placeholder shown while content is loading |
| [`InfiniteLoadingBar`](#infiniteloadingbar) | A loading bar for work of unknown duration |
| [`TabContainer`](#tabcontainer) | Closable, movable tabs, each with its own content |
| [`LazyWidget`](#lazywidget) | A helper that creates a child widget on first use |

Icons and images are read from an `fs.FS`. The gallery passes `nil` and loads from disk relative to the working directory. In your own application, pass an `fs.FS` such as an `embed.FS` so the files travel with the binary.

---

### LabeledToggle

A `basicwidget.Toggle` with a text label next to it.

**Use it when** you need a labeled boolean setting, for example "Follow redirects" or "Color mode".

- `SetLabel(label string)` sets the text.
- `Toggle()` returns the underlying `*basicwidget.Toggle`. Use it to read or set the value and to subscribe to changes.

```go
darkMode widgets.LabeledToggle

// in Build
r.darkMode.SetLabel("Color mode")
r.darkMode.Toggle().OnValueChanged(func(context *guigui.Context, value bool) {
	// react to the new value
})
```

### Button

A `basicwidget.Button` that can show an icon and a tooltip. Because it embeds `basicwidget.Button`, everything the basic button offers (`SetText`, `OnDown`, `OnUp`, and so on) is available too.

**Use it when** a button needs a picture or a hover explanation, for example an "add" button with a plus icon.

- `SetIcon(name string, fs fs.FS) error` sets the icon.
- `SetTooltip(text string)` sets the text shown when hovering the button.

```go
addButton widgets.Button

// in Build
if err := r.addButton.SetIcon("icons/add.png", nil); err != nil {
	return err
}
r.addButton.SetText("Add a new tab")
r.addButton.SetTooltip("Add a new tab to the tab container widget")
r.addButton.OnDown(func(context *guigui.Context) {
	// handle the click
})
```

### TextInput

A `basicwidget.TextInput` with a context menu (right click). The default entries are copy, paste, cut, select all, undo and redo. You can disable any of them and add your own. Because it embeds `basicwidget.TextInput`, the usual text input methods are available too.

**Use it when** you want a text field with a context menu you control, for example a request body editor that adds extra actions next to copy and paste.

- `SetIcon(name string, fs fs.FS) error` shows an icon in the input.
- `SetDisable(item, disable)` and `Disable(item)` turn a default menu entry off or report whether it is off. The items are `ContextMenuItemCopy`, `ContextMenuItemPaste`, `ContextMenuItemCut`, `ContextMenuItemSelectAll`, `ContextMenuItemUndo` and `ContextMenuItemRedo`.
- `SetAditionalItems([]basicwidget.PopupMenuItem[int])` adds your own entries to the menu, and `OnAdditionalItemSelected(fn)` is called with the index of the chosen one.
- `SetTextContent(widgets.TextContent)` and `TextContent()` set and read the kind of content the input holds: `TextContentPlainText` (the default) or `TextContentHTML`.

```go
bodyInput widgets.TextInput

// in Build
if err := r.bodyInput.SetIcon("icons/input.png", nil); err != nil {
	return err
}
r.bodyInput.SetDisable(widgets.ContextMenuItemCut, true)
r.bodyInput.SetAditionalItems([]basicwidget.PopupMenuItem[int]{
	{Text: "Clear"},
})
r.bodyInput.OnAdditionalItemSelected(func(context *guigui.Context, index int) {
	// index of the selected additional item
})
```

### Link

Text that opens a URL when clicked. It embeds `basicwidget.Text`, so `SetValue` sets the visible text. When the visible text differs from the URL, a tooltip shows the URL on hover.

**Use it when** you want an inline link, such as a documentation or "learn more" link.

- `SetURL_Str(u string)` sets the target from a string, and `SetURL(u *url.URL) error` sets it from a parsed URL.
- `SetBaseStyle(style *basicwidget.TextStyle)` changes how the text is drawn.

```go
docsLink widgets.Link

// in Build
r.docsLink.SetValue("Click here to open the link")
r.docsLink.SetURL_Str("https://example.com/docs")
```

> **Warning:** a `Link` can also open file paths. Only pass URLs you trust, and be careful with values that come from user input or from remote content.

### Icon

Draws an icon, and optionally reacts to clicks.

**Use it when** you need a standalone icon, for example a close icon in a toolbar, an icon in a form row, or a clickable icon button.

- `Load(icon string, fs fs.FS, size image.Point) error` loads the icon. The gallery passes `nil` and `image.Point{}` for the last two arguments.
- `SetClickable(bool)` and `Clickable()` control whether it reacts to the pointer, and `OnClick(fn)` registers the handler. A clickable icon also shows the pointer cursor.
- `HasIcon()` reports whether an icon is loaded, and `Size(context)` returns its size.

```go
closeIcon widgets.Icon

// in Build
if err := r.closeIcon.Load("icons/close.png", nil, image.Point{}); err != nil {
	return err
}
r.closeIcon.SetClickable(true)
r.closeIcon.OnClick(func(context *guigui.Context, icon *widgets.Icon) {
	// handle the click
})
```

### InfoHint

A small information indicator that shows a hint text.

**Use it when** a setting or field needs a short explanation without taking up space on the screen, typically next to a label in a form.

- `SetText(text string)` sets the hint text.

```go
hint widgets.InfoHint

// in Build
r.hint.SetText("This is an information hint.")
```

### Text

A widget that wraps a `basicwidget.Text` and exposes it through `Text()`.

**Use it when** you want to place a text widget where a `widgets` widget is expected, and still configure the inner `basicwidget.Text` directly.

```go
title widgets.Text

// in Build
r.title.Text().SetValue("Request headers")
```

### SkeletonLoader

A generic wrapper, `SkeletonLoader[T]`, that shows a loading placeholder in place of any Guigui widget `T` while it is loading.

**Use it when** content arrives asynchronously, such as an image, a response body or a list, and you want the layout to keep its shape while waiting.

- `SetLoading(bool)` and `Loading()` control and report the loading state.
- `Widget()` returns the wrapped widget `T`, which you configure as usual.

The loader has no intrinsic size, so give it one. The gallery wraps it in `guigui.WidgetWithSize` and sets a fixed width:

```go
type Root struct {
	guigui.DefaultWidget

	picture guigui.WidgetWithSize[*widgets.SkeletonLoader[*basicwidget.Image]]
}

// in Build
r.picture.SetFixedWidth(basicwidget.UnitSize(context) * 8)
r.picture.Widget().Widget().SetImage(img)  // configure the wrapped basicwidget.Image
r.picture.Widget().SetLoading(isLoading)   // show or hide the placeholder
```

`r.picture.Widget()` is the `*SkeletonLoader`, and `r.picture.Widget().Widget()` is the wrapped `*basicwidget.Image`.

### InfiniteLoadingBar

A loading bar that animates continuously, with no progress value.

**Use it when** something is in progress and you cannot tell how long it will take, such as sending a request or waiting for a connection. It needs no configuration.

```go
loadingBar widgets.InfiniteLoadingBar

// in Build: add it only while work is in progress
if r.busy {
	adder.AddWidget(&r.loadingBar)
}
```

### TabContainer

A tab bar together with the content of the selected tab. Tabs can be closable and movable (drag to reorder), and each tab carries its own widget.

`TabContainer[T]` is generic over the type of the value attached to each tab, which must be `comparable`. Use `struct{}` when you do not need one, or something like a string or integer ID to identify tabs.

**Use it when** the user can open several documents or views at once, such as multiple requests in an API client.

A tab is described by a `TabItem[T]`:

| Field | Meaning |
| --- | --- |
| `Text` | The tab title |
| `Icon` | A `TabIcon{Icon: "path", Fs: fsys}` shown in the tab |
| `Closable` | Whether the tab has a close button |
| `Movable` | Whether the tab can be reordered |
| `Widget` | The content shown when the tab is selected |
| `Value` | A value of type `T` that you attach to the tab |

`TabContainer` does not own the list of tabs. You keep the slice, pass it with `SetTabItems`, and update it yourself in the callbacks:

- `OnTabClose(fn)` is called when a tab is closed. Remove the tab from your slice.
- `OnTabSwitch(fn)` is called when two tabs swap places. Swap them in your slice.
- `OnTabSelect(fn)` is called when a tab is selected.
- `SelectTabItem(index)` selects a tab, and `SelectedTabItem()` returns the selected index and whether there is one.
- `SetUnselectedWidget(widget)` sets the widget shown when no tab is selected.

Give every tab its own widget instance. A widget can only live in one place in the tree.

```go
type Root struct {
	guigui.DefaultWidget

	addButton widgets.Button
	tabs      widgets.TabContainer[struct{}]
	tabItems  []widgets.TabItem[struct{}]
	tabCount  int
}

// in Build
r.addButton.SetText("Add a new tab")
r.addButton.OnDown(func(context *guigui.Context) {
	r.tabCount++

	content := &basicwidget.Text{}
	content.SetValue(fmt.Sprintf("Tab %d", r.tabCount))

	r.tabItems = append(r.tabItems, widgets.TabItem[struct{}]{
		Text:     fmt.Sprintf("Tab %d", r.tabCount),
		Icon:     widgets.TabIcon{Icon: "icons/info-tip.png"},
		Closable: true,
		Movable:  true,
		Widget:   content,
	})
})

r.tabs.SetTabItems(r.tabItems)
r.tabs.OnTabClose(func(context *guigui.Context, index int, item widgets.TabItem[struct{}]) {
	r.tabItems = slices.Delete(r.tabItems, index, index+1)
})
r.tabs.OnTabSwitch(func(context *guigui.Context, index1 int, item1 widgets.TabItem[struct{}], index2 int, item2 widgets.TabItem[struct{}]) {
	r.tabItems[index1], r.tabItems[index2] = r.tabItems[index2], r.tabItems[index1]
})
adder.AddWidget(&r.tabs)
```

The container usually takes the remaining space, so lay it out with a flexible size:

```go
// in Layout
(guigui.LinearLayout{
	Direction: guigui.LayoutDirectionVertical,
	Items: []guigui.LinearLayoutItem{
		{Widget: &r.addButton},
		{Widget: &r.tabs, Size: guigui.FlexibleSize(1)},
	},
}).LayoutWidgets(context, widgetBounds.Bounds(), layouter)
```

### LazyWidget

`LazyWidget[T]` is a small generic helper that holds a child widget of type `T` and creates it the first time `Widget()` is called. It is useful when building your own wrapper widgets that need a child but should still work from their zero value.

```go
type Wrapper struct {
	guigui.DefaultWidget

	inner widgets.LazyWidget[*basicwidget.Text]
}

// anywhere in Build or Layout
r.inner.Widget().SetValue("created on first use")
```

### Image helpers

The package also exports a few helpers for loading images from an `fs.FS`:

- `OpenImage(file_path string, fs fs.FS) (*ImageDetails, error)` opens an image and returns an `*ImageDetails`.
- `ImageDetails` holds the decoded `*ebiten.Image` (`Img`), extra data (`Details`) and a `Cache` flag.
- `ChangeImageDetails(file_path string, fs fs.FS, details any) (ok bool)` updates the details of an image.

See the [package documentation](https://pkg.go.dev/github.com/Zbolt-org/widgets) for the full API.

## Known limitations

### Disabling widgets is not supported yet

Guigui lets you disable a widget with `context.SetEnabled(&widget, false)`. The widgets in this module do not support this yet. A widget you disable this way may not look or behave as a disabled widget should, so do not rely on it to block input or to signal that a control is unavailable.

The behavior of disabling a widget may change in a future release, including how each widget looks and responds when disabled. If your application depends on it, check the release notes when you update and test the affected screens.

Until it is supported, you can keep a widget out of the user's way by not adding it in `Build` (a widget that is not added is not drawn or sent input), or by ignoring its events in your own callbacks.

`TextInput.SetDisable` and `TextInput.Disable` are a separate feature. They turn individual context menu entries on or off, not the whole widget.

## License

Released under the [BSD 3-Clause License](LICENSE).

`lazy-widget.go` is derived from code by The Guigui Authors and keeps its original Apache-2.0 notice.
