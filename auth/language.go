package auth

import (
	"golang.org/x/text/language"
	"invest/utils/constants"
	"invest/utils/message"

	"net/http"
)

func Parse_prefered_language_of_user(w http.ResponseWriter, r *http.Request) (message.Msg) {
	var supported = []language.Tag{
		language.Kazakh,
		language.Russian,
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
	}

	lang, err := r.Cookie(constants.CookieLanguageKeyWord)
	if err != nil {
		lang = &http.Cookie{}
	}

	accept := r.Header.Get(constants.HeaderContentLanguage)
	if accept == "" {
		r.Header.Get(constants.HeaderAcceptLanguage)
	}

	tag, _ := language.MatchStrings(language.NewMatcher(supported), lang.String(), accept)

	var user_language string
	switch tag {
	case language.Kazakh:
		user_language = constants.ContentLanguageKk
	case language.English:
		user_language = constants.ContentLanguageEn
	default:
		user_language = constants.ContentLanguageRu
	}

	r.Header.Set(constants.HeaderContentLanguage, user_language)
	r.Header.Set(constants.HeaderAcceptLanguage, user_language)

	return message.Msg{}
}