package validation

import (
	"strconv"
)

var key = map[string][2]string {
	"Title"			: { "001", "标题" },
	"Category"		: { "002", "节点" },
	"Tag" 			: { "003", "标签" },
	"Content"		: { "004", "内容" },
	"Abstract" 		: { "005", "摘要" },
	"Name"			: {	"006", "名" 	 },
	"DisplayName" 	: { "007", "别名" },
	"Description"	: { "008", "描述" },
}

var condition = map[string][2]string {
	"Required"	: { "001", "不能为空" },
	"MaxSize"	: { "002", "超出最大范围" },
}


var env = map[string][2]string {
	"Post"		: { "4001", "帖子" },
	"Cate"		: { "4002", "节点" },

}


func CustomErrCodeAndMessage (errEnv string,errKey string,errCondition string) (int,string) {

	codeEnv,okOne := env[errEnv]
	codeKey,okTwo := key[errKey]
	codeCondition,okThree := condition[errCondition]
	if !okOne || !okTwo || !okThree {
		return 4000000000,"系统内部错误"
	}
	codeString := codeEnv[0] + codeKey[0] + codeCondition[0]
	message := codeKey[1] + codeCondition[1]
	code,_:= strconv.Atoi(codeString)
	return code,message
}

