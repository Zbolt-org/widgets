package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"slices"
	"widgets"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/ebiten/v2"
)

func load_image() (*ebiten.Image, error) {
	file, err := os.Open("./gallery/image.jpeg")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

var (
	tab_items []widgets.TabItem[struct{}] = []widgets.TabItem[struct{}]{}
	tab_count int                         = len(tab_items)
)

type Root struct {
	guigui.DefaultWidget
	background basicwidget.Background

	form basicwidget.Form

	labeled_toggle_text basicwidget.Text
	toggle              widgets.LabeledToggle

	icon_widget_text basicwidget.Text
	close            widgets.Icon

	infotip_widget_text basicwidget.Text
	info_tip            widgets.InfoHint

	skeleton_loader_text basicwidget.Text
	image                *ebiten.Image
	image_wdiget         guigui.WidgetWithSize[*widgets.SkeletonLoader[*basicwidget.Image]]

	set_image_load_text basicwidget.Text
	load_image_toggle   widgets.LabeledToggle

	infinite_loading_bar_text basicwidget.Text
	infinite_loadingbar       widgets.InfiniteLoadingBar

	text_input_text basicwidget.Text
	text_input      widgets.TextInput

	add_tab_text basicwidget.Text
	add_tab_btn  basicwidget.Button

	tab_container widgets.TabContainer[struct{}]
	layout        guigui.LinearLayout
}

func (r *Root) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	adder.AddWidget(&r.background)

	r.toggle.SetLabel("Color mode")
	r.toggle.Toggle().OnValueChanged(func(context *guigui.Context, value bool) {
		if value {
			context.SetPreferredColorMode(ebiten.ColorModeDark)
		} else {
			context.SetPreferredColorMode(ebiten.ColorModeLight)
		}
	})

	err := r.close.Load("icons/close.png", nil, image.Point{})
	if err != nil {
		return err
	}
	r.info_tip.SetText("This is a infomation hint.")

	if r.image == nil {
		img, err := load_image()
		if err != nil {
			panic(err.Error())
		}
		r.image = img
	}

	r.image_wdiget.SetFixedWidth(basicwidget.UnitSize(context) * 8)
	r.image_wdiget.Widget().Widget().SetImage(r.image)
	r.load_image_toggle.Toggle().OnValueChanged(func(context *guigui.Context, value bool) {
		r.image_wdiget.Widget().SetLoading(value)
	})

	r.labeled_toggle_text.SetValue("Labeled toggle")
	r.icon_widget_text.SetValue("Icon widget")
	r.infotip_widget_text.SetValue("Infotip widget")
	r.skeleton_loader_text.SetValue("Skeleton loader")
	r.set_image_load_text.SetValue("Set image loading state")

	r.infinite_loading_bar_text.SetValue("Infinite loadingbar")

	r.text_input_text.SetValue("TextInput widget with a contextmenu")

	r.add_tab_text.SetValue("Add tab")
	r.add_tab_btn.OnDown(func(context *guigui.Context) {
		tab_count++
		widget := basicwidget.Text{}
		widget.SetHorizontalAlign(basicwidget.HorizontalAlignCenter)
		widget.SetVerticalAlign(basicwidget.VerticalAlignMiddle)
		widget.SetValue(fmt.Sprintf("Tab %v", tab_count))
		tab_items = append(tab_items, widgets.TabItem[struct{}]{
			Text:     fmt.Sprintf("Tab %v", tab_count),
			Movable:  true,
			Closable: true,
			Widget:   &widget,
		})
	})
	r.add_tab_btn.SetText("Add a new tab")

	r.form.SetItems([]basicwidget.FormItem{
		{
			PrimaryWidget:   &r.labeled_toggle_text,
			SecondaryWidget: &r.toggle,
		},
		{
			PrimaryWidget:   &r.icon_widget_text,
			SecondaryWidget: &r.close,
		},
		{
			PrimaryWidget:   &r.infotip_widget_text,
			SecondaryWidget: &r.info_tip,
		},
		{
			PrimaryWidget:   &r.skeleton_loader_text,
			SecondaryWidget: &r.image_wdiget,
		},
		{
			PrimaryWidget:   &r.set_image_load_text,
			SecondaryWidget: &r.load_image_toggle,
		},
		{
			PrimaryWidget:   &r.infinite_loading_bar_text,
			SecondaryWidget: &r.infinite_loadingbar,
		},
		{
			PrimaryWidget:   &r.text_input_text,
			SecondaryWidget: &r.text_input,
		},
		{
			PrimaryWidget:   &r.add_tab_text,
			SecondaryWidget: &r.add_tab_btn,
		},
	})
	adder.AddWidget(&r.form)

	r.tab_container.SetTabItems(tab_items)
	r.tab_container.OnTabClose(func(context *guigui.Context, index int, data widgets.TabItem[struct{}]) {
		tab_items = slices.Delete(tab_items, index, index+1)
	})
	r.tab_container.OnTabSwitch(func(context *guigui.Context, index1 int, data1 widgets.TabItem[struct{}], index2 int, data2 widgets.TabItem[struct{}]) {
		item1 := tab_items[index1]
		tab_items[index1] = tab_items[index2]
		tab_items[index2] = item1
	})
	adder.AddWidget(&r.tab_container)
	return nil
}

func (r *Root) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&r.background, widgetBounds.Bounds())

	u := basicwidget.UnitSize(context)
	r.layout = guigui.LinearLayout{
		Gap:       u / 4,
		Padding:   guigui.Padding{u, u, u, u},
		Direction: guigui.LayoutDirectionVertical,
		Items: []guigui.LinearLayoutItem{
			{
				Widget: &r.form,
			},
			{
				Widget: &r.tab_container,
				Size:   guigui.FlexibleSize(1),
			},
		},
	}

	r.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func main() {
	if err := guigui.Run(&Root{}, &guigui.RunOptions{
		Title: "My App",
	}); err != nil {
		log.Fatal(err)
	}
}
