package views

import (
	"github.com/derailed/k9s/internal/resource"
	"github.com/gdamore/tcell"
)

type nodeView struct {
	*resourceView
}

func newNodeView(t string, app *appView, list resource.List, c colorerFn) resourceViewer {
	v := nodeView{newResourceView(t, app, list, c).(*resourceView)}
	{
		v.extraActionsFn = v.extraActions
		v.switchPage("no")
	}

	return &v
}

func (v *nodeView) extraActions(aa keyActions) {
	aa[KeyShiftQ] = newKeyAction("Sort CPU%", v.sortColCmd(6, false), true)
	aa[KeyShiftW] = newKeyAction("Sort MEM%", v.sortColCmd(7, false), true)
	aa[KeyShiftA] = newKeyAction("Sort RCPU%", v.sortColCmd(8, false), true)
	aa[KeyShiftS] = newKeyAction("Sort RMEM%", v.sortColCmd(9, false), true)
	aa[KeyShiftC] = newKeyAction("Sort CPU", v.sortColCmd(10, false), true)
	aa[KeyShiftM] = newKeyAction("Sort MEM", v.sortColCmd(11, false), true)
}

func (v *nodeView) sortColCmd(col int, asc bool) func(evt *tcell.EventKey) *tcell.EventKey {
	return func(evt *tcell.EventKey) *tcell.EventKey {
		t := v.getTV()
		t.sortCol.index, t.sortCol.asc = t.nameColIndex()+col, asc
		t.refresh()

		return nil
	}
}
