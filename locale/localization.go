package locale

/*
import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
)

var EnLocalizer *i18n.Localizer
var RuLocalizer *i18n.Localizer
var LocalizeTemplate = func(messageID string, template interface{}, lang string) string {
	// if Localizer isn't initialized, set up with system language
	switch lang {
	case "Russian":
		msg := RuLocalizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
			TemplateData:template })
		return msg
	default:
		msg := EnLocalizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
			TemplateData:template})
		return msg
	}
}

//var LocalizePlural = func(messageID string, template interface{}, pluralCount interface{}) string {
//
//	msg := Localizer.MustLocalize(&i18n.LocalizeConfig{
//		MessageID:    messageID,
//		TemplateData: template,
//		PluralCount:  pluralCount})
//	return msg
//}

var LocalizeMsgID = func(messageID string, lang string) string {
	switch lang {
	case "ru":
		msg := RuLocalizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID})
		return msg
	default:
		msg := EnLocalizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID})
		return msg
	}
}

func StartLocalizer() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	dir, err := os.Getwd()
	if err != nil {
		//log.Fatal(err)
	}
	bundle.MustLoadMessageFile(dir+"/locale/active.en.toml")
	bundle.MustLoadMessageFile(dir+"/locale/active.ru.toml")
	//
	//if lang == "" {
	//	lang = os.Getenv("LANG")
	//	// remove UTF-8 suffix from language if found
	//	if i := strings.Index(lang, ".UTF-8"); i != -1 {
	//		lang = lang[:i]
	//	}
	//}
	//Localizer = i18n.NewLocalizer(bundle, "en-US")
	EnLocalizer = i18n.NewLocalizer(bundle, "en")
	RuLocalizer = i18n.NewLocalizer(bundle, "ru")
	// Test translation
	// fmt.Println(Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloWorld"}))
}*/
