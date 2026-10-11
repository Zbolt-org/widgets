package widgets

import (
	"github.com/guigui-gui/guigui"
)

type tab_bar_widget[T comparable] struct {
	guigui.DefaultWidget

	tab_items guigui.WidgetSlice[*tab_item_widget[T]]
}
