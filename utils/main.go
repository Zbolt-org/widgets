package utils

import (
	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

func SmallPadding(context *guigui.Context) guigui.Padding {
	u := basicwidget.UnitSize(context) / 2
	return guigui.Padding{u, u, u, u}
}

func MediumPadding(context *guigui.Context) guigui.Padding {
	u := basicwidget.UnitSize(context)
	return guigui.Padding{u, u, u, u}
}

func LargePadding(context *guigui.Context) guigui.Padding {
	u := basicwidget.UnitSize(context) * 2
	return guigui.Padding{u, u, u, u}
}

func SmallGap(context *guigui.Context) int {
	u := basicwidget.UnitSize(context) / 4
	return u
}

func MediumGap(context *guigui.Context) int {
	u := basicwidget.UnitSize(context)
	return u / 2
}

func LargeGap(context *guigui.Context) int {
	u := basicwidget.UnitSize(context)
	return u
}

func VerticalCenter(widget guigui.Widget) guigui.LinearLayout {
	layout := guigui.LinearLayout{
		Direction: guigui.LayoutDirectionVertical,
		Items: []guigui.LinearLayoutItem{
			{Size: guigui.FlexibleSize(1)},
			{Widget: widget},
			{Size: guigui.FlexibleSize(1)},
		},
	}
	return layout
}
