package errors

const (
	InvalidURL Error = iota + 100
	LoginRequired
	FlagMissing
	FileGeneration
	QuestionNotFound
	QuestionIdOutOfRange
	NoDailyChallenge
	Unexpected
	NotImplemented
	NoCodeSnippetsFound
	ExtensionNotFound
	SnippetNotFoundInGivenLang
	BadAppCommand
	InvalidSearchMethod
	FailedToWriteImage
	FailedToCreateDirectory
	FailedToDownloadImage
	FileTemplateInvalid
	FileTemplateDoesntExist
)

const repoUrl = "https://github.com/ISKalsi/leet-scrape"

type Error int

func (e Error) Error() string {
	return e.GetMessage("leetscrape")
}

func (e Error) GetCode() int {
	return int(e)
}

func (e Error) GetMessage(cli string) string {
	switch e {
	case InvalidURL:
		return "Please enter valid leetcode url"
	case LoginRequired:
		return "The part of leetcode you're trying to access requires login. Try entering the name of problem instead.  Enter \"" + cli + " help\""
	case FlagMissing:
		return "At least one flag is required. Enter \"" + cli + " help\""
	case FileGeneration:
		return "Something went wrong during the generation of solution template file"
	case QuestionNotFound:
		return "Question not found. Please check the spellings."
	case QuestionIdOutOfRange:
		return "Question number is out of range of the given list of questions on Leetcode. Please check what you entered"
	case NoDailyChallenge:
		return "No daily challenges exist for the current month. Check back later"
	case NoCodeSnippetsFound:
		return "No code snippets found. Are you sure you entered the right command ? (\"" + cli + " sol\" and not \"leetscrape ques\")"
	case SnippetNotFoundInGivenLang:
		return "Couldn't find the code snippet in your programming language. There could be 2 reasons -\n1. Spelling mistake. Generally available options: C++, C, C#, Kotlin, Java, Python, Python3, Swift, Go, PHP, Racket, Rust, Ruby, JavaScript, TypeScript, Scala, ErLang, Elixir\n2. Snippet is not available on leetcode in your programming language"
	case ExtensionNotFound:
		return "Couldn't find the code snippet in your programming language. There could be 2 reasons -\n1. Spelling mistake. Generally available options: C++, C, C#, Kotlin, Java, Python, Python3, Swift, Go, PHP, Racket, Rust, Ruby, JavaScript, TypeScript, Scala, ErLang, Elixir\n2. Snippet is not available on leetcode in your programming language"
	case InvalidSearchMethod:
		return "Problem with the search medium. Neither name, url nor number"
	case BadAppCommand:
		return "Internal error. Bad app command"
	case FailedToWriteImage:
		return "Failed to write image asset to file"
	case FailedToDownloadImage:
		return "Failed to download image asset"
	case FailedToCreateDirectory:
		return "Failed to create folders for the given path. Try executing the command with administrator privileges"
	case Unexpected:
		return "This was not supposed to happen. Please raise an issue on the Github repo - " + repoUrl
	case NotImplemented:
		return "This feature is yet to be implemented. Keep track on " + repoUrl
	case FileTemplateDoesntExist:
		return "File Template doesn't exist. Please raise an issue on the Github repo - " + repoUrl
	case FileTemplateInvalid:
		return "File template invalid. Please raise an issue on the Github repo - " + repoUrl
	default:
		return "!!! I was not prepared for this (>.<) Please raise an issue on the Github repo - " + repoUrl
	}
}
