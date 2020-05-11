package lt2020

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	// qst.RadioVali = "mustRadioGroup"
	qst.RadioVali = ""
	qst.CSSLabelHeader = ""
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("example")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	q.LangCodesOrder = []string{"en", "de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Beispielumfrage", "en": "Example survey"}

	for i1 := 0; i1 < 3; i1++ {
		page := q.AddPage()
		gr := page.AddGroup()
		gr.Cols = 3
		inp := gr.AddInput()
		inp.Name = fmt.Sprintf("name%v", i1)
		inp.Type = "text"
		inp.Label = trl.S{"de": "Vorname", "en": "first name"}
		inp.MaxChars = 10
	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
