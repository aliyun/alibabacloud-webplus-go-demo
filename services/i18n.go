package services

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	LOCALE_MAP map[string]map[string]string
)

func LoadLocales() {
	files, _ := ioutil.ReadDir("locales")
	LOCALE_MAP = make(map[string]map[string]string)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".properties") {
			continue
		}
		code := strings.Split(file.Name(), ".")[0]
		LOCALE_MAP[code] = loadLocale("locales/" + file.Name())
	}
}

func loadLocale(fileName string) map[string]string {
	m := make(map[string]string)
	content, _ := ioutil.ReadFile(fileName)
	for _, l := range strings.Split(string(content), "\n") {
		if !strings.Contains(l, "=") {
			continue
		}
		sp := strings.SplitN(strings.TrimSpace(l), "=", 2)
		m[sp[0]] = sp[1]
	}
	return m
}

func format(ptn string, args []string) string {
	result := ptn
	for i, arg := range args {
		p := "{" + strconv.Itoa(i) + "}"
		result = strings.ReplaceAll(result, p, arg)
	}

	return result
}

func I18n(lang string, str string) string {
	sp := strings.Split(str, " ")
	args := make([]string, len(sp)-1)
	for i, s := range sp[1:] {
		args[i] = strings.Trim(s, "[]")
	}

	return format(LOCALE_MAP[lang][sp[0]], args)
}
