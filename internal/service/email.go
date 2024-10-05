package service

import (
	Redis "StudentServicePlatform/internal/pkg/redis"
	"crypto/tls"
	"fmt"
	"math/rand"
	"mime"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
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

var (
    verificationCodeCache = cache.New(5*time.Minute, 10*time.Minute)
)

// EmailService defines the email service interface
// type EmailService interface {
//     SendVerificationCode( to string) error
//     VerifyVerificationCode(email string, code string) (bool)
// }

// // emailService implements the EmailService interface
// type emailService struct {
// }

// // NewEmailService creates a new email service
// func NewEmailService() EmailService {
//     return &emailService{}
// }

// // SendVerificationCode sends a verification code to the user's email
// func (e *emailService) SendVerificationCode( to string) error {
//     code := generateVerificationCode()
//     err := e.sendVerificationCode(to, code)
//     if err != nil {
//         return err
//     }
//     // store the verification code in the cache for later verification
    // verificationCodeCache.Set(to, code, cache.DefaultExpiration)
    // return nil
// }

// // sendVerificationCode sends the verification code to the specified email address
// func (e *emailService) sendVerificationCode(to string, code string) error {
//     // Create a new email instance
//     m := gomail.NewMessage()
//     m.SetHeader("From", "学生服务平台 <2039210588@qq.com>")
//     m.SetHeader("To", to)
//     m.SetHeader("Subject", "Verification Code")
//     // Set the email HTML content
//     m.SetBody("text/html", fmt.Sprintf(`
//         <h1>Verification Code</h1>
//         <p>Your verification code is: <strong>%s</strong></p>
//     `, code))

//     // Send the email (using QQ SMTP)
//     d := gomail.NewDialer("smtp.qq.com", 587, "2039210588@qq.com", "xixbkschclgtbddg")
//     if err := d.DialAndSend(m); err != nil {
//         return err // If there is an error while sending, return the error
//     }
//     return nil // Email sent successfully
// }

// generateVerificationCode generates a random 6-digit verification code
func GenerateVerificationCode() string {
    rand.Seed(time.Now().UnixNano())
    code := fmt.Sprintf("%06d", rand.Intn(1000000))
    return code
}

// VerifyVerificationCode verifies the verification code sent to the user
func VerifyVerificationCode(email string, code string) (bool) {
 // 从内存缓存中获取验证码
	cachedCode, found := verificationCodeCache.Get(email)
	if !found {
		fmt.Println("Verification code not found for email:", email)
		return false
	}

	// 从Redis中获取验证码
	redisCode, err := Redis.RedisClient.Get(ctx, email).Result()
	if err == redis.Nil {
		fmt.Println("Redis中未找到验证码")
		return false
	} else if err != nil {
		fmt.Println("Redis错误:", err)
		return false
	}

	// 比较缓存中的验证码和Redis中的验证码
	if cachedCode != code && redisCode != code {
		return false
	}
	return true
}
func GetEmailTemplate(message string) string {
	return "<html><head><body><table class=m-shell border=0 width=775 cellspacing=0 cellpadding=0><tbody><tr><td class=td style=width:775px;min-width:775px;font-size:0;line-height:0;padding:0;margin:0;font-weight:normal><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"p-80 mpy-35 mpx-15\" style=padding:80px bgcolor=#212429><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"img pb-45\" style=font-size:0;line-height:0;text-align:left;padding-bottom:45px><a href=http://www.myzjut.org target=_blank rel=noopener><img src=https://www.yuanshen.com/images/ys.96a55539.png width=150 height=\"\" border=0></a><tr><td><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"title-36 pb-30 c-grey6 fw-b\" style=\"font-size:36px;line-height:42px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;padding-bottom:30px;color:#bfbfbf;font-weight:bold\">您好！</table><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"text-18 c-white pb-20\" style=\"font-size:18px;line-height:25px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;color:#dbdbdb;padding-bottom:20px\">欢迎您注册学生管理平台</table><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"text-18 c-white pb-20\" style=\"font-size:18px;line-height:25px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;color:#dbdbdb;padding-bottom:20px\">下面是您的验证码</table><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"text-18 c-white pb-20\" style=\"font-size:18px;line-height:25px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;color:#dbdbdb;padding-bottom:20px\">祝您生活愉快！</table><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"pb-70 mpb-50\" style=padding-bottom:20px><table border=0 width=100% cellspacing=0 cellpadding=0 bgcolor=#17191c><tbody><tr><td class=\"py-30 px-56\" style=\"padding:30px 56px 30px 56px\"><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td style=\"font-size:25px;line-height:30px;font-family:Arial,sans-serif,'Motiva Sans';color:#f1f1f1;text-align:center;letter-spacing:1px\">验证码<tr><td style=\"font-size:25px;line-height:30px;font-family:Arial,sans-serif,'Motiva Sans';color:#3a9aed;text-align:center;letter-spacing:1px\">" + message + "<tr><td style=padding-bottom:16px>&nbsp;<tr><tr></table></table></table><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"text-18 c-white pb-20\" style=\"font-size:18px;line-height:25px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;color:#dbdbdb;padding-bottom:20px\">这是一份由系统自动发送的邮件。<td class=\"text-18 c-white pb-20\" style=\"font-size:18px;line-height:25px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;color:#dbdbdb;padding-bottom:20px\">如果您没有注册但是收到了这份邮件，我也不知道为什么。</table></table><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=pt-30 style=padding-top:30px><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=img style=font-size:0;line-height:0;text-align:left bgcolor=#3a9aed width=3>&nbsp;<td class=img style=font-size:0;line-height:0;text-align:left width=37>&nbsp;<td><table border=0 width=100% cellspacing=0 cellpadding=0><tbody><tr><td class=\"text-16 py-20 c-grey4 fallback-font\" style=\"font-size:16px;line-height:22px;font-family:Arial,sans-serif,'Motiva Sans';text-align:left;padding-top:20px;padding-bottom:20px;color:#f1f1f1\">祝您愉快，<br>学生服务平台</table></table></table></table></table>"
}

func SendVerifyCode(email string, code string)  error{
	message := GetEmailTemplate(code)
    host := "smtp.qq.com"
    port := 465 // 使用 SSL 端口
    userName := "2039210588@qq.com"
    password := "xixbkschclgtbddg"
    
    m := gomail.NewMessage()
    m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "学生服务平台") + " <2039210588@qq.com>") // 使用完整的发件人格式
    m.SetHeader("To", email) // 使用参数传入的 email
    m.SetHeader("Subject", "学生服务平台反馈") // 邮件主题
    m.SetBody("text/html", message)

    d := gomail.NewDialer(host, port, userName, password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 注意：生产环境中应使用有效证书

    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Error sending email:", err) // 改为优雅的错误处理
        return err
	}

	verificationCodeCache.Set(email, code, cache.DefaultExpiration)
	err := Redis.RedisClient.Set(ctx, email, code, 5*time.Minute).Err()
    if err != nil {
        return err
    }
    return nil
}