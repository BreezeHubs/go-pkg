package ginpkg

// 实现思路：对gin的responseWriter进行包装， 每次写往请求方写响应数据的时候，将响应数据返回出去

// 定义一个新的CustomResponseWriter，通过组合方式持有一个gin.ResponseWriter和response body缓存

//type CustomResponseWriter struct {
//	gin.ResponseWriter
//	body *bytes.Buffer
//}
//
//func (w CustomResponseWriter) Write(b []byte) (int, error) {
//	w.body.Write(b)
//	return w.ResponseWriter.Write(b)
//}
//
//func (w CustomResponseWriter) WriteString(s string) (int, error) {
//	w.body.WriteString(s)
//	return w.ResponseWriter.WriteString(s)
//}
//
//
//func AccessLogHandler(db *pkg.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		startTime := time.Now()
//		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
//		c.Writer = blw
//
//		// request body
//		data, _ := io.ReadAll(c.Request.Body)
//		c.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 这点很重要，把字节流重新放回 body 中
//
//		c.Next()
//
//		// 获取响应信息
//		responseDuration := time.Since(startTime)
//		var responseTime string
//		if responseDuration < time.Millisecond {
//			responseTime = fmt.Sprintf("%dμs", responseDuration.Microseconds())
//		} else if responseDuration < time.Second {
//			responseTime = fmt.Sprintf("%dms", responseDuration.Milliseconds())
//		} else {
//			responseTime = fmt.Sprintf("%.2fs", responseDuration.Seconds())
//		}
//
//		logInfo := &model.ResponseLog{
//			Url:         c.Request.URL.String(),
//			Method:      c.Request.Method,
//			Status:      c.Writer.Status(),
//			ContentType: c.ContentType(),
//			Param:       *(*string)(unsafe.Pointer(&data)),
//			Content:     blw.body.String(),
//			SpendTime:   responseTime,
//			IP:          c.ClientIP(),
//			UserAgent:   c.Request.UserAgent(),
//		}
//		filterLogInfo(logInfo)
//		if err := db.Engine.Model(&model.ResponseLog{}).Create(logInfo).Error; err != nil {
//			log.Println("写入 response log 失败", err)
//		}
//		//fmt.Sprintf("[%s] url=%s, status=%d, resp=%s", c.Request.Method, c.Request.URL, c.Writer.Status(), blw.body.String())
//	}
//}
