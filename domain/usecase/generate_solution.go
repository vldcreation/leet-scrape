package usecase

import (
	"fmt"
	"strings"

	template_engine "text/template"

	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/ISKalsi/leet-scrape/v2/domain/service"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/util"
)

var (
	PATH_FILES = "../../domain/usecase/templates/%s.tmpl"
)

type GenerateSolutionFile struct {
	writer      service.FileWriter
	question    *entity.Question
	fileName    string
	path        string
	boilerplate string
	language    string
	template    string
}

func NewGenerateSolutionFile(q *entity.Question, w service.FileWriter, fileName string, path string, boilerplate string, lang string, template string) *GenerateSolutionFile {
	return &GenerateSolutionFile{
		writer:      w,
		question:    q,
		fileName:    fileName,
		path:        path,
		boilerplate: boilerplate,
		language:    lang,
		template:    template,
	}
}

func (uc *GenerateSolutionFile) Execute() error {
	if len(uc.question.CodeSnippets) == 0 {
		return errors.NoCodeSnippetsFound
	}

	solution := &entity.Solution{}
	for _, snippet := range uc.question.CodeSnippets {
		if strings.EqualFold(snippet.Lang, uc.language) {
			data := uc.boilerplate + snippet.Code
			data = util.FixNewlineAndTabs(data)

			// check if choose template
			solution.CodeSnippet = data
			if uc.template != "" {
				tmpl, err := template_engine.ParseFiles(fmt.Sprintf(PATH_FILES, uc.template))
				if err != nil {
					return err
				}
				var builder strings.Builder
				err = tmpl.Execute(&builder, solution)
				if err != nil {
					return err
				}
				solution.CodeSnippet = builder.String()
			}

			ext := fileExtensions[snippet.Lang]
			if ext == "" {
				return errors.ExtensionNotFound
			}
			fileNameWithExtension := uc.fileName + "." + ext
			return uc.writer.WriteDataToFile(fileNameWithExtension, uc.path, solution.String())
		}
	}
	return errors.SnippetNotFoundInGivenLang
}

var fileExtensions = map[string]string{
	"C++":        "cpp",
	"Java":       "java",
	"Python":     "py",
	"Python3":    "py",
	"C":          "c",
	"C#":         "cs",
	"JavaScript": "js",
	"Ruby":       "rb",
	"Swift":      "swift",
	"Go":         "go",
	"Scala":      "scala",
	"Kotlin":     "kt",
	"Rust":       "rs",
	"PHP":        "php",
	"TypeScript": "ts",
	"Racket":     "rkt",
	"ErLang":     "erl",
	"Elixir":     "ex",
}
