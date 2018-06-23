package qst

import (
	"bytes"
	"fmt"

	"github.com/zew/go-questionaire/cfg"
)

const (
	htmlExample = `
		<ol class="progress progress--large">
			<li class="is-complete" data-step="1">
				Page 1
			</li>
			<li class="is-complete" data-step="2">
				Page 2
			</li>
			<li class="is-active" data-step="3">
				Page 3
			</li>
			<li data-step="4" class="progress__last">
				Confirm
			</li>
		</ol>
`
)

// ProgressBar generates a discrete position indicator
// https://fribly.com/2015/01/01/scalable-responsive-progress-bar/
//
// It should be clickable and jump to the indicated page.
// This means, we have to submit the form and submit the destination page.
func (q *QuestionaireT) ProgressBar() string {

	b := bytes.Buffer{}
	// b.WriteString(fmt.Sprintf(htmlExample))

	// This does not post the form :(
	//		location.href='%v?page=%v
	// We need
	//
	//		this.form.page.value='%v';
	// this.form.submit();  - not working for <li> element
	//		document.forms.frmMain.submit()
	//
	// Debug with
	// 		console.log('document.forms.frmMain.page.value: ',document.forms.frmMain.page.value);
	b.WriteString(fmt.Sprintf(`<input type="hidden" name="page" value="-1" >`))

	b.WriteString(fmt.Sprintf("<ol class='progress'>"))

	for idx, p := range q.Pages {

		if p.NoNavigation {
			continue
		}

		eff := p.Label.Tr(q.LangCode)
		if p.Short != nil {
			eff = p.Short.Tr(q.LangCode)
		}

		liClass := ""
		if idx < q.CurrPage {
			liClass = "is-complete"
		} else if idx == q.CurrPage {
			liClass = "is-active"
		}

		// Some positional finetuning
		sect := p.Section.TrSilent(q.LangCode)
		leftOrCenter := "text-align: left; width: 98%; transform: translate(25%, 0px);"
		if sect != "" {
			sect = fmt.Sprintf("<b>%v</b>", sect)
			sect += vspacer
			if idx == len(q.Pages)-1 {
				// last element more to the right
				leftOrCenter = "text-align: left; width: 98%; transform: translate(40%, 0px);"
			}
		} else if sect == "" {
			leftOrCenter = "transform: translate(0, 75%);"
		}

		// make hyperlinks to the pages
		onclick := fmt.Sprintf(` onclick="document.forms.frmMain.page.value='%v';document.forms.frmMain.submit();" style="cursor:pointer"  `, idx)
		if cfg.Get().AllowSkipForward == false && idx > q.CurrPage {
			onclick = ""
		}

		b.WriteString(
			fmt.Sprintf(`
					<li 
						%v
						class="%v" data-step="%v">
						<span style='display: inline-block; line-height: 95%%;  %v'>
							%v%v
						<span>
					</li> 
				`,
				onclick,
				liClass, p.NavigationalNum, //idx+1,
				leftOrCenter,
				sect, eff,
			),
		)

	}
	b.WriteString(fmt.Sprintf("</ol>"))

	return b.String()
}
