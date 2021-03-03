package qst

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"github.com/zew/go-questionnaire/css"
)

func wrap(w io.Writer, tagName, forVal, className, style, content string) {

	forKeyVal := ""
	if forVal != "" {
		forKeyVal = fmt.Sprintf("for='%v'", forVal)

	}

	if style != "" {
		styles := strings.Split(style, ";")
		style = strings.Join(styles, ";\n\t\t")
		style = fmt.Sprintf("%v%v%v", "\n\t\t", style, "\n\t\t")
	}

	fmt.Fprintf(w,
		"<%v %v class='%v' style='%v' >\n%v\n</%v>\n",
		tagName,
		forKeyVal,
		className,
		style,
		content,
		tagName,
	)

	// fmt.Fprintf(w, "<!-- /%v -->\n", className)

}
func divWrap(w io.Writer, className, style, content string) {
	wrap(w, "div", "", className, style, content)
}

func (inp inputT) labelDescription(w io.Writer, langCode string) {

	if inp.Label.Empty() && inp.Desc.Empty() {
		return
	}

	if inp.IsHidden() {
		return
	}

	if inp.IsLayout() {
		if !inp.Label.Empty() {
			fmt.Fprint(w, inp.Label.Tr(langCode))
		}
		if !inp.Desc.Empty() {
			fmt.Fprint(w, inp.Desc.Tr(langCode))
		}
		return
	}

	// classes are only for font-size, font-weight
	// inline-block styles are applied in outer wrapper
	if !inp.Label.Empty() {
		fmt.Fprintf(w, " <span class='input-label-text'       >%v</span>", inp.Label.Tr(langCode))
	}
	if !inp.Desc.Empty() {
		fmt.Fprintf(w, " <span class='input-description-text' >%v</span>", inp.Desc.Tr(langCode))
	}

}

// ShortSuffix appends suffix to the input - using &nbsp;
// for longer text, reverse order of label and control;
// sadly, &nbsp; has no effect;
// we have to use white-space: nowrap; on the grid-item
func (inp *inputT) ShortSuffix(ctrl string, langCode string) string {

	if inp.Suffix.Empty() {
		return ctrl
	}

	ctrl = strings.Trim(ctrl, "\r\n\t ")
	ctrl = fmt.Sprintf("%v%v", ctrl, inp.Suffix.TrSilent(langCode)) // &nbsp; is broken anyway

	return ctrl
}

//
//
// GroupHTMLGridBased renders a group of inputs to GroupHTMLGridBased
func (q QuestionnaireT) GroupHTMLGridBased(pageIdx, grpIdx int) string {

	wCSS := &strings.Builder{}
	gr := q.Pages[pageIdx].Groups[grpIdx]

	//
	//
	if gr.Style == nil {
		gr.Style = css.NewStylesResponsive()
	}
	gr.Style.Desktop.BoxStyle.Display = "grid"
	if gr.Style.Desktop.GridContainerStyle.AutoFlow == "" {
		gr.Style.Desktop.GridContainerStyle.AutoFlow = "row"
	}
	if gr.Style.Desktop.GridContainerStyle.TemplateColumns == "" {
		gr.Style.Desktop.GridContainerStyle.TemplateColumns = strings.Repeat("1fr ", gr.Cols)
	}
	if gr.Style.Desktop.GridContainerStyle.ColumnGap == "" {
		gr.Style.Desktop.GridContainerStyle.ColumnGap = "0.4rem"
	}
	if gr.Style.Desktop.GridContainerStyle.RowGap == "" {
		gr.Style.Desktop.GridContainerStyle.RowGap = "0.8rem"
	}
	gridContainerClass := fmt.Sprintf("pg%02v-grp%02v", pageIdx, grpIdx)
	fmt.Fprint(wCSS, gr.Style.CSS(gridContainerClass))

	//
	//
	wInner := &strings.Builder{} // inside the container

	for inpIdx, inp := range gr.Inputs {
		if inp.Type == "dyn-composite-scalar" {
			continue
		}
		if inp.Type == "dyn-composite" {
			continue
		}

		if inp.Style == nil {
			inp.Style = css.NewStylesResponsive()

			// input wrapper is item      to group
			inp.Style.Desktop.GridItemStyle.Col = fmt.Sprintf("auto / span %v", inp.ColSpanLabel+inp.ColSpanControl)

			// input wrapper is container to label and control
			inp.Style.Desktop.BoxStyle.Display = "grid"
			inp.Style.Desktop.GridContainerStyle.AutoFlow = "row"
			inp.Style.Desktop.GridContainerStyle.TemplateColumns = strings.Repeat("1fr ", inp.ColSpanLabel+inp.ColSpanControl)

		}
		gridItemClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v", pageIdx, grpIdx, inpIdx)
		fmt.Fprint(wCSS, inp.Style.CSS(gridItemClass))

		wInp := &strings.Builder{} // label and control of input

		{

			if inp.ColSpanLabel > 0 {
				wLbl := &strings.Builder{}
				if inp.StyleLbl == nil {
					inp.StyleLbl = css.NewStylesResponsive()
					inp.StyleLbl.Desktop.GridItemStyle.Col = fmt.Sprintf("auto / span %v", inp.ColSpanLabel)
					inp.StyleLbl.Desktop.GridItemStyle.AlignSelf = "center"
					// inp.StyleLbl.Desktop.GridItemStyle.Order = 2  // label post control
				}

				lblClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v-lbl", pageIdx, grpIdx, inpIdx)
				fmt.Fprint(wCSS, inp.StyleLbl.CSS(lblClass))

				inp.labelDescription(wLbl, q.LangCode)
				if inp.IsLayout() {
					divWrap(wInp, lblClass+" grid-item-lvl-2", "", wLbl.String())
				} else {
					wrap(wInp, "label", inp.Name, lblClass+" grid-item-lvl-2", "", wLbl.String())
				}
			}

			if inp.ColSpanControl > 0 {
				wCtl := &strings.Builder{}
				if inp.StyleCtl == nil {
					inp.StyleCtl = css.NewStylesResponsive()
					inp.StyleCtl.Desktop.GridItemStyle.Col = fmt.Sprintf("auto / span %v", inp.ColSpanControl)
					inp.StyleCtl.Desktop.GridItemStyle.AlignSelf = "center"

					if !inp.IsLayout() {
						inp.StyleCtl.Desktop.TextStyle.WhiteSpace = "nowrap" // prevent suffix from being wrapped
					}
				}
				if inp.Type == "radiogroup" || inp.Type == "checkboxgroup" {
					inp.StyleCtl.Desktop.BoxStyle.Display = "grid"
					inp.StyleCtl.Desktop.GridContainerStyle.AutoFlow = "row"
					inp.StyleCtl.Desktop.GridContainerStyle.TemplateColumns = strings.Repeat("1fr ", len(inp.Radios))
				}
				if inp.Type == "radio" || inp.Type == "checkbox" {
					inp.StyleCtl.Desktop.GridItemStyle.JustifySelf = "center"
					inp.StyleCtl.Desktop.GridItemStyle.AlignSelf = "center"
					inp.StyleCtl.Desktop.TextStyle.AlignHorizontal = "center"
				}

				ctlClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v-ctl", pageIdx, grpIdx, inpIdx)
				fmt.Fprint(wCSS, inp.StyleCtl.CSS(ctlClass))

				fmt.Fprint(wCtl, q.InputHTMLGrid(pageIdx, grpIdx, inpIdx))
				divWrap(wInp, ctlClass+" grid-item-lvl-2", "", wCtl.String())
			}

		}

		//
		divWrap(wInner, gridItemClass+" grid-item-lvl-1", "", wInp.String())
	}

	//
	//
	wContainer := &strings.Builder{}
	divWrap(wContainer, gridContainerClass+" grid-container", "", wInner.String())

	w := &strings.Builder{}
	fmt.Fprint(w, css.StyleTag(wCSS.String()))
	fmt.Fprint(w, wContainer.String())
	return w.String()

}

// InputHTMLGrid renders an input to HTML
func (q QuestionnaireT) InputHTMLGrid(pageIdx, grpIdx, inpIdx int) string {

	// gr := q.Pages[pageIdx].Groups[grpIdx]
	inp := *q.Pages[pageIdx].Groups[grpIdx].Inputs[inpIdx]
	nm := inp.Name
	ctrl := ""

	switch inp.Type {

	case "textblock":
		// no op

	case "button":
		ctrl = fmt.Sprintf(
			"<button type='submit' name='%v' value='%v' class='%v' accesskey='%v'><b>%v</b> %v</button>\n",
			inp.Name, inp.Response, inp.CSSControl, inp.AccessKey,
			inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
		)

	case "textarea":
		width := ""
		colsRows := fmt.Sprintf(" cols='%v' rows='1' ", inp.MaxChars+1)
		if inp.MaxChars > 80 {
			colsRows = fmt.Sprintf(" cols='80' rows='%v' ", inp.MaxChars/80+1)
			// width = fmt.Sprintf("width: %vem;", int(float64(80)*1.05))
			width = "width: 98%;"
		}
		ctrl += fmt.Sprintf("<textarea        name='%v' id='%v' title='%v %v' class='%v' style='%v' maxlength='%v' %v  autocomplete='off' >%v</textarea>\n",
			nm, nm, inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode), inp.CSSControl, width, inp.MaxChars, colsRows, inp.Response)

	case "dropdown":
		// i.DD = &DropdownT{}
		inp.DD.SetName(inp.Name)
		inp.DD.LC = q.LangCode
		inp.DD.SetTitle(inp.Label.TrSilent(q.LangCode) + " " + inp.Desc.TrSilent(q.LangCode))
		inp.DD.Select(inp.Response)
		inp.DD.SetAttr("class", inp.CSSControl)
		sort.Sort(inp.DD)

		ctrl += inp.DD.RenderStr()

	case "text", "number", "hidden", "checkbox", "radio":
		rspvl := inp.Response

		checked := ""
		if inp.Type == "checkbox" {
			if rspvl == ValSet {
				checked = "checked=\"checked\""
			}
			rspvl = ValSet
		}
		if inp.Type == "radio" {
			if rspvl == inp.ValueRadio {
				checked = "checked=\"checked\""
			}
			rspvl = inp.ValueRadio
		}

		width := fmt.Sprintf("width:%.2frem", float32(inp.MaxChars)*0.65)
		if inp.Type == "checkbox" || inp.Type == "radio" {
			width = ""
		}

		inputMode := ""
		if inp.Type == "number" {
			if inp.Step != 0 {
				if inp.Step >= 1 {
					inputMode = fmt.Sprintf(" step='%.0f'  ", inp.Step)
				} else {
					prec := int(math.Log10(1 / inp.Step))
					f := fmt.Sprintf(" step='%%.%vf'  ", prec)
					inputMode = fmt.Sprintf(f, inp.Step)
				}
			}
		}

		ctrl += fmt.Sprintf(
			`<input type='%v'  %v  
				name='%v' id='%v' title='%v %v' 
				class='%v' style='%v'  
				size='%v' maxlength=%v min='%v' max='%v'  %v  value='%v'   autocomplete='off'  />
			`,
			inp.Type, inputMode,
			nm, fmt.Sprintf("%v%v", nm, inp.ValueRadio), inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
			inp.CSSControl, width,
			inp.MaxChars, inp.MaxChars, inp.Min, inp.Max, checked,
			rspvl,
		)

		// the checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if inp.Type == "checkbox" {
			ctrl += fmt.Sprintf(
				"<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

	case "dyn-textblock":
		ctrl = fmt.Sprintf("<span class='%v'>%v</span>\n", inp.CSSLabel, inp.Label.Tr(q.LangCode))

	case "dyn-composite", "dyn-composite-scalar":
		// no op
		// rendered at group level -  rendered by composit

	default:
		ctrl = fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, inp.Type, implementedTypes)
	}

	//
	// common

	// append suffix
	ctrl = inp.ShortSuffix(ctrl, q.LangCode)

	if inp.ErrMsg.Set() {
		ctrl += fmt.Sprintf("<div class='error error-block >%v</div>\n", inp.ErrMsg.TrSilent(q.LangCode)) // ugly layout - but radiogroup and checkboxgroup won't have validation errors anyway
	}

	return ctrl

}