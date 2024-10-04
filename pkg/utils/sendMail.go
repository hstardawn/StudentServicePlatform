package utils

import (
	"crypto/tls"
	"fmt"
	"mime"

	"gopkg.in/gomail.v2"
)

func SendMail(email string, name string, content string) {
	message := `
	<html>
	<head>
		<meta charset="UTF-8">
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; /* 更换字体 */
				line-height: 1.6;
				background-color: #f9f9f9; /* 更柔和的背景色 */
				color: #333; /* 深色文本 */
				padding: 20px;
			}
			.container {
				background: #ffffff; /* 白色背景 */
				padding: 20px;
				border-radius: 10px;
				box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
			}
			h1 {
				color: #2E8B57; /* 深绿色 */
				text-align: center;
				margin-bottom: 20px;
			}
			h2 {
				color: #333; /* 深色小标题 */
				margin-top: 20px;
			}
			p {
				margin: 10px 0;
				color: #555; /* 深灰色文本 */
			}
			.button {
				display: inline-block;
				padding: 12px 20px;
				margin: 20px 0;
				background-color: #E8D098; /* 蓝色 */
				color: white;
				text-decoration: none;
				border-radius: 5px;
				box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
				transition: background-color 0.3s ease;
				font-weight: bold; /* 加粗字体 */
			}
			.button:hover {
				background-color:#ffffff; /* 悬停时的更深蓝色 */
			}
			.footer {
				margin-top: 20px;
				font-size: 0.9em;
				color: #777; /* 轻灰色 */
				text-align: center;
			}
			.divider {
				height: 1px;
				background-color: #e0e0e0;
				margin: 20px 0;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>学生服务平台反馈</h1>
			<p>亲爱的 %s,</p>
			<p style="text-indent:2em">%s</p> 
			<div class="divider"></div>
			<h2>更多信息</h2>
			<p style="text-indent:2em">如需更多信息，请点击以下链接：</p>
			<a href="https://example.com" class="button">访问我们的平台</a>
			<p style="text-indent:2em">祝好！</p>
		</div>
		<div class="footer">
			<p>如有任何问题，请随时与我们联系。</p>
			<p>谢谢！</p>
		</div>
	</body>
	</html>
	`

    host := "smtp.qq.com"
    port := 465 // 使用 SSL 端口
    userName := "2039210588@qq.com"
    password := "xixbkschclgtbddg"
    
    m := gomail.NewMessage()
    m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "学生服务平台") + " <2039210588@qq.com>") // 使用完整的发件人格式
    m.SetHeader("To", email) // 使用参数传入的 email
    m.SetHeader("Subject", "学生服务平台反馈") // 邮件主题
    m.SetBody("text/html", fmt.Sprintf(message, name, content))

    d := gomail.NewDialer(host, port, userName, password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 注意：生产环境中应使用有效证书

    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Error sending email:", err) // 改为优雅的错误处理
        return
    }
}