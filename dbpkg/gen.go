package dbpkg

import "strings"

func GenSQlCaseString(field string, m map[string]string) string {
	/*
		( CASE
			WHEN `device_type` = 'ttl' THEN '弹弹乐'
			WHEN `device_type` = 'fwjt' THEN '凤舞九天'
			WHEN `device_type` = 'wzxg' THEN '王者峡谷'
		ELSE 'UNKNOWN' END
		)
	*/
	var strBuff strings.Builder
	strBuff.WriteString("( CASE ")
	for k, v := range m {
		strBuff.WriteString("WHEN `")
		strBuff.WriteString(field)
		strBuff.WriteString("`='")
		strBuff.WriteString(k)
		strBuff.WriteString("' THEN '")
		strBuff.WriteString(v)
		strBuff.WriteString("' ")
	}
	strBuff.WriteString(" ELSE 'UNKNOWN' END )")
	return strBuff.String()
}
