package translator

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	configRoot = "config/locale"
)
const (
	enLang = "en"
)

var (
	localizerJA = func() *i18n.Localizer {
		bundleJA := i18n.NewBundle(language.Japanese)
		bundleJA.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		files, err := ioutil.ReadDir(configRoot + "/ja")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			bundleJA.MustLoadMessageFile(configRoot + "/ja/" + file.Name())
		}
		return i18n.NewLocalizer(bundleJA, "ja")
	}()
	localizerEN = func() *i18n.Localizer {
		bundleEN := i18n.NewBundle(language.English)
		bundleEN.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		files, err := ioutil.ReadDir(configRoot + "/en")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			bundleEN.MustLoadMessageFile(configRoot + "/en/" + file.Name())
		}
		return i18n.NewLocalizer(bundleEN, "en")
	}()
)

// Arg for localization
type Arg struct {
	Lang  string
	MsgID string
	Data  map[string]interface{}
}

// Localize translate
func Localize(arg Arg) (string, error) {
	cnf := &i18n.LocalizeConfig{
		MessageID: arg.MsgID,
	}
	if arg.Data != nil {
		cnf.TemplateData = arg.Data
	}

	if arg.Lang == enLang {
		return localizerEN.Localize(cnf)
	}
	return localizerJA.Localize(cnf)
}

// LocalizeField translate a field
func LocalizeField(lang, field string) (string, error) {
	cnf := &i18n.LocalizeConfig{MessageID: field}
	if lang == enLang {
		return localizerEN.Localize(cnf)
	}
	return localizerJA.Localize(cnf)
}
