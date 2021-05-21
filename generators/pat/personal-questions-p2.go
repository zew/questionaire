package pat

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// PersonalQuestions2 - numbered 8-15
func PersonalQuestions2(q *qst.QuestionnaireT, vE VariableElements) error {

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.3rem 0 0"

	validatorInput := ""
	validatorRadio := ""
	if vE.AllMandatory {
		validatorInput = "must"
		validatorRadio = "mustRadioGroup"
	}

	{
		// page := q.AddPage()
		page := q.AddPage()
		// page.Label = trl.S{"de": "POP page"}
		// page.Short = trl.S{"de": "Stiftungen 1"}
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		{
			gr := page.AddGroup()
			gr.Cols = 12
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Bitte geben Sie Ihr Geschlecht an:
					</p>
				`, vE.NumberingQuestions+0),
				}
			}
			keyVals := []string{
				"male:Männlich",
				"female:Weiblich",
				"diverse:Divers",
			}

			for _, kv := range keyVals {
				sp := strings.Split(kv, ":")
				key := sp[0]
				val := sp[1]
				lbl := trl.S{"de": val}
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q8"
				rad.Validator = validatorRadio
				rad.ValueRadio = key
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}
		}

		//
		//
		{
			gr := page.AddGroup()
			gr.Cols = 12
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Bitte geben Sie Ihr Geburtsjahr an:
					</p>
				`, vE.NumberingQuestions+1),
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q9"
				inp.Validator = validatorInput
				inp.ColSpan = 4
				inp.ColSpanControl = 1

				inp.Min = 1900
				inp.Max = 2010
				inp.Step = 1
				inp.MaxChars = 5
			}
		}

		// separate header - since the states are vertically shown
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					In welchem Bundesland befindet sich Ihr Hauptwohnsitz?
					</p>
				`, vE.NumberingQuestions+2),
				}

			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			// gr.Vertical(8)
			for _, stt := range trl.FederalStatesGermanyISOs2 {
				lbl := stt.S
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q10"
				rad.Validator = validatorRadio
				rad.ValueRadio = strings.ToLower(stt.Key)
				rad.ColSpan = 1 // for vertical
				rad.ColSpan = 4 // horizontal
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Welcher ist Ihr höchster Bildungsabschluss?
					</p>
				`, vE.NumberingQuestions+3),
				}
			}

			keyVals := []string{
				"kein_abschluss:Kein Abschluss",
				"hauptschule:Haupt&shy;schul&shy;abschluss",
				"mittlere_reife:Mittlere Reife",
				"abitur:Abitur, (Fach-)Hoch&shy;schul&shy;reife",
				"hochschule:Universitäts- oder FH-Abschluss (Bachelor, Diplom, Master)",
				"promotion:Promotion",
			}

			for _, kv := range keyVals {
				sp := strings.Split(kv, ":")
				key := sp[0]
				val := sp[1]
				lbl := trl.S{"de": val}
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q11"
				rad.Validator = validatorRadio
				rad.ValueRadio = key
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}

		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		{
			gr := page.AddGroup()
			gr.Cols = 12
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Wie ist Ihr Familienstand?
					</p>
				`, vE.NumberingQuestions+4),
				}
			}

			keyVals := []string{
				"single:Alleinstehend",
				"engaged:Partnerschaft ohne Ehe",
				"married:Verheiratet",
			}

			for _, kv := range keyVals {
				sp := strings.Split(kv, ":")
				key := sp[0]
				val := sp[1]
				lbl := trl.S{"de": val}
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q12"
				rad.Validator = validatorRadio
				rad.ValueRadio = key
				rad.ColSpan = 6
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Wie viel Geld verdienen Sie persönlich monatlich 
					nach Abzug von Steuern und Sozialversicherungsbeiträgen?
					</p>
				`, vE.NumberingQuestions+5),
				}
			}

			keyVals := []string{
				"null_euro:0 Euro", // '0' would be empty
				"upto500:bis unter 500 Euro",
				"upto1000:500 bis unter 1.000 Euro",
				"upto1500:1.000 bis unter 1.500 Euro",
				"upto2000:1.500 bis unter 2.000 Euro",
				"upto3000:2.000 bis unter 3.000 Euro",
				"upto4000:3.000 bis unter 4.000 Euro",
				"upto5000:4.000 bis unter 5.000 Euro",
				"upto10000:5.000 bis unter 10.000 Euro",
				"over10000:Mehr als 10.000 Euro",
			}

			for _, kv := range keyVals {
				sp := strings.Split(kv, ":")
				key := sp[0]
				val := sp[1]
				lbl := trl.S{"de": val}
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q13"
				rad.Validator = validatorRadio
				rad.ValueRadio = key
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Mit welcher Partei fühlen Sie sich 
					aufgrund Ihrer Werte und Überzeugungen am ehesten verbunden? 
					<br>
					<i>Bitte beachten Sie, dass nur eine Antwort zulässig ist.</i>
					</p>
				`, vE.NumberingQuestions+6),
				}
			}
		}
		keyVals := []string{
			"cducsu:CDU/CSU",
			"linke:Die Linke",
			"spd:SPD",
			"gruene:Bündnis 90/Die Grünen",
			"fdp:FDP",
			"afd:AfD",
			// "other:Andere",
		}
		{
			for _, kv := range keyVals {
				gr := page.AddGroup()
				gr.Cols = 8
				// gr.Cols = 4
				gr.RandomizationGroup = 1
				gr.BottomVSpacers = 0
				{
					sp := strings.Split(kv, ":")
					key := sp[0]
					val := sp[1]
					lbl := trl.S{"de": val}
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "q14"
					rad.Validator = validatorRadio
					rad.ValueRadio = key
					rad.ColSpan = 4
					rad.ColSpanLabel = 4
					rad.ColSpanControl = 1
					rad.Label = lbl
					rad.StyleLbl = lblStyleRight
				}

			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			gr.BottomVSpacers = 0
			{
				key := "other"
				val := "Andere"
				lbl := trl.S{"de": val}

				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q14"
				rad.Validator = validatorRadio
				rad.ValueRadio = key
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q14_other_text"
					inp.ColSpan = 4
					inp.ColSpanControl = 1
					inp.MaxChars = 14
					inp.Validator = "otherParty"
				}

			}

		}

		//
		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": " &nbsp; ",
					"de": " &nbsp; ",
				}
			}
		}

		//
		{
			grStPage78 := css.NewStylesResponsive(nil)
			grStPage78.Desktop.StyleGridContainer.GapRow = "0.1rem"
			grStPage78.Desktop.StyleGridContainer.GapColumn = "0.01rem"

			gb := qst.NewGridBuilderRadiosWithValidator(
				[]float32{
					0, 1,
					0, 1,
					0, 1, // 3
					0, 1,
					0, 1, // 5
					0, 1,
					0, 1, // 7
					0, 1,
					0, 1, // 9
					0, 1,
					0, 1, // 11
					1.2, 1, // weiss nicht
				},
				labelsOneToSeven4,
				[]string{"q15"},
				[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "weiss_nicht"},
				[]trl.S{},
				validatorRadio,
			)
			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
					<p>
					<b>Frage&nbsp;%v.</b>
					In der Politik reden die Leute häufig von "links" und "rechts". 
					Wenn Sie die Skala hier benutzen, 
					wo würden Sie sich selbst einordnen, 
					wenn 1 "links" und 11 "rechts" ist? 
					Bitte geben Sie den Wert an, der auf Sie persönlich zutrifft.
					</p>
					<br>
				`, vE.NumberingQuestions+7),
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0"
			gr.BottomVSpacers = 4
		}

	}

	return nil
}