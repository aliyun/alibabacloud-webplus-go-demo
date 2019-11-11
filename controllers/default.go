package controllers

import (
	"alibaba/com/webplusdemo/mustache"
	"alibaba/com/webplusdemo/services"
	"github.com/astaxie/beego"
	"os"
	"regexp"
	"strings"
)

type MainController struct {
	beego.Controller
}

func getLambdas(lang string) map[string]func(string) string {
	lambdas := make(map[string]func(string) string)
	lambdas["i18n"] = func(s string) string {
		return services.I18n(lang, s)
	}

	return lambdas
}

func format(ptn string, args map[string]string) string {
	result := ptn
	re, _ := regexp.Compile("\\${.*?}")
	for _, s := range re.FindAllString(ptn, -1) {
		k := strings.Trim(s, "${}")
		if val, ok := args[k]; ok {
			result = strings.ReplaceAll(result, s, val)
		}
	}
	return result
}

func getData() map[string]interface{} {
	config := services.LoadConfig()
	envVars := make(map[string]string)

	for _, v := range os.Environ() {
		sp := strings.SplitN(v, "=", 2)
		envVars[sp[0]] = sp[1]
	}

	ctx := map[string]interface{}{
		"siteId":             config.Site.Id,
		"quickstartDocUrl":   config.Quickstart.Doc.Url,
		"quickstartRepoName": config.Quickstart.Repo.Name,
		"quickstartRepoUrl":  config.Quickstart.Repo.Url,
		"appUrl":             format(config.App.Url, envVars),
		"envUrl":             format(config.Env.Url, envVars),
		"nextStep":           config.Next.Step.Show,
		"nextStepPackageUrl": config.Next.Step.Package.Url,
		"consoleUrl":         config.WebPlus.Console.Url,
		"envs": map[string]interface{}{
			"appRegionId": envVars["WP_APP_REGION_ID"],
			"appId":       envVars["WP_APP_ID"],
			"appName":     envVars["WP_APP_NAME"],
			"envId":       envVars["WP_ENV_ID"],
			"envName":     envVars["WP_ENV_NAME"],
			"fromCLI":     "CLI" == envVars["WP_CHANGE_TRIGGER_FROM"],
			"fromConsole": "Console" == envVars["WP_CHANGE_TRIGGER_FROM"],
		},
	}

	return ctx
}

func (c *MainController) Get() {
	lang := "en"
	if strings.HasPrefix(c.Ctx.Input.Header("Accept-Language"), "zh") {
		lang = "zh"
	}

	txt := mustache.RenderFile(beego.BConfig.WebConfig.ViewsPath+"/index.mustache", getLambdas(lang), getData())
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	c.Ctx.Output.Body([]byte(txt))
}
