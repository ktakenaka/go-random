package translator

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	localizerJA = func() *i18n.Localizer {
		bundleJA := i18n.NewBundle(language.Japanese)
		bundleJA.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundleJA.MustLoadMessageFile("app/errors/translator/app_error.ja.toml")
		return i18n.NewLocalizer(bundleJA, "ja")
	}()
	localizerEN = func() *i18n.Localizer {
		bundleEN := i18n.NewBundle(language.English)
		bundleEN.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundleEN.MustLoadMessageFile("app/errors/translator/app_error.en.toml")
		return i18n.NewLocalizer(bundleEN, "en")
	}()
)

// Localize translate
func Localize(lang, msgID string, data map[string]interface{}) (string, error) {
	if lang == "en" {
		return localizerEN.Localize(&i18n.LocalizeConfig{
			MessageID:    msgID,
			TemplateData: data,
		})
	}
	return localizerJA.Localize(&i18n.LocalizeConfig{
		MessageID:    msgID,
		TemplateData: data,
	})
}
