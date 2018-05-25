package qst

// For all multi lingual strings.
// Contains one value for each language code.
type transMapT map[string]string // type translation map

const noTrans = "Translation map not initialized."

// Tr translates by key.
// Defaults           to english.
// Defaults otherwise to first existing map key.
// Returns a 'signifiers string' if no translation exists,
// to help to uncover missing translations.
func (t transMapT) Tr(langCode string) string {
	if val, ok := t[langCode]; ok {
		return val
	}
	if val, ok := t["en"]; ok {
		return val
	}
	for _, val := range t {
		return val
	}
	if t == nil {
		return noTrans
	}
	return noTrans
}

// TrSilent gives no warning - if the translation is not set.
// Good if we do not require a translation string.
// Good for i.e. HTML title attribute - where errors are easy to overlook.
func (t transMapT) TrSilent(langCode string) string {
	ret := t.Tr(langCode)
	if ret == noTrans {
		return ""
	}
	return ret
}

// String is the default "stringer" implementation
func (t transMapT) String() string {
	return t.Tr("en")
}
