package widgets

import (
	"reflect"
	"sync"

	"github.com/guigui-gui/guigui"
)

// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 The Guigui Authors
// Original code belongs to @hajime-hojshi
type LazyWidget[T guigui.Widget] struct {
	widget T
	once   sync.Once
}

func (l *LazyWidget[T]) Widget() T {
	l.once.Do(func() {
		t := reflect.TypeFor[T]()
		if t.Kind() == reflect.Pointer {
			l.widget = reflect.New(t.Elem()).Interface().(T)
		}
	})
	return l.widget
}
