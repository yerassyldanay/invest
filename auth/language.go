package auth

import (
	"invest/utils"
	"net/http"
	"golang.org/x/text/language"
)

func Parse_prefered_language_of_user(w http.ResponseWriter, r *http.Request) (utils.Msg) {
	var supported = []language.Tag{
		language.Kazakh,
		language.Russian,
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
	}

	lang, err := r.Cookie(utils.CookieLanguageKeyWord)
	if err != nil {
		lang = &http.Cookie{}
	}

	accept := r.Header.Get(utils.HeaderContentLanguage)
	if accept == "" {
		r.Header.Get(utils.HeaderAcceptLanguage)
	}

	tag, _ := language.MatchStrings(language.NewMatcher(supported), lang.String(), accept)

	var user_language string
	switch tag {
	case language.Kazakh:
		user_language = utils.ContentLanguageKk
	case language.Russian:
		user_language = utils.ContentLanguageRu
	default:
		user_language = utils.ContentLanguageEn
	}

	r.Header.Set(utils.HeaderContentLanguage, user_language)
	r.Header.Set(utils.HeaderAcceptLanguage, user_language)

	return utils.Msg{}
}