package constant

var CheckinQuestionSheetNameArr = []string{
	CheckinQuestionIdiom,
	CheckinQuestionArithmetic,
	CheckinQuestionEnglish,
	CheckinQuestionFun,
}

const (
	CheckinQuestionTypeIdiom      = 1 // 成语
	CheckinQuestionTypeArithmetic = 2 // 口算
	CheckinQuestionTypeEnglish    = 3 // 英语
	CheckinQuestionTypeFun        = 4 // 趣味题
	CheckinQuestionIdiom          = "成语"
	CheckinQuestionArithmetic     = "口算"
	CheckinQuestionEnglish        = "英语"
	CheckinQuestionFun            = "趣味题"
	CheckinQuestionStatusEnable   = 1 // 启用
	CheckinQuestionStatusDisable  = 2 // 禁用
)

var CheckinQuestionOptionMap = map[string]string{
	"0": "A",
	"1": "B",
	"2": "C",
	"3": "D",
}
