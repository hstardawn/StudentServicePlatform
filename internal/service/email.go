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
	m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "学生服务平台")+" <2039210588@qq.com>") // 使用完整的发件人格式
	m.SetHeader("To", email)                                                             // 使用参数传入的 email
	m.SetHeader("Subject", "学生服务平台反馈")                                                   // 邮件主题
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
func VerifyVerificationCode(email string, code string) bool {
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
func GetEmailTemplate(msg string) string {
	return "<html lang='en'><head><meta charset='UTF-8'><meta name='viewport'content='width=device-width, initial-scale=1.0'><title>Document</title><style>.rainbow{background-image:-webkit-linear-gradient(left,#147B96,#E6D205 25%,#147B96 50%,#E6D205 75%,#147B96);-webkit-text-fill-color:transparent;-webkit-background-clip:text;-webkit-background-size:200%100%;-webkit-animation:maskedAnimation 4s infinite linear;font-size:18px}@keyframes maskedAnimation{0%{background-position:0 0}100%{background-position:-100%0}}</style></head><body><table bgcolor='#15171e'style='background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'class='responsive-table'><tbody><tr><td bgcolor='#15171e'style='background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#15171e'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'><tbody><tr><td class='header-section-padding-0px'style='height: 50px;'align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td class='header-padding-0px-center'bgcolor='#15171e'><table width='100%'cellspacing='0'cellpadding='0'border='0'style='width: 100%; min-width: 100%; border-spacing: 0; border-collapse: collapse; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word;'><tbody><tr><td><img src='http://media.blizzard.com/emails/global/shared-components/spacer.gif'style='line-height: 0; visibility: hidden; min-width: 600px; height: 1px; max-height: 1px; min-height: 1px; text-decoration: none; -ms-interpolation-mode: bicubic; border: none;'height='1'></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#15171e'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'><tbody><tr><td class='header-section-padding-25px-15px-25px'align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td class='logo'style='width: 600px; text-align: center'bgcolor='#15171e'><center><a href='https://battle.net'target='_blank'alias='Image Link (Battle.net Logo)'rel='noopener'><img src='https://bnetproduct-a.akamaihd.net/fef/98ade86ccaa737c0a348f470d2beb3a7-logo.png'width='341'height='86'style='font-family: 'noto=''sans',='' 'open=''sans'='' ,='' frutiger,='' 'frutiger=''linotype'='' univers,='' 'helvetica=''neue'='' helvetica,='' arial,='' 'gill=''sans=''mt'='' 'myriad=''pro'='' myriad,='' 'dejavu=''condensed'='' 'liberation='''nimbus='' l'='''malgun='' gothic'='''microsoft='' yahei'=''applesdgothicneo,=''applegothic,=''dotum,=''jhenghei'='' 'hiragino=''kaku=''gothic=''pron=''w3'='' osaka,='' メイリオ,='' meiryo,='' 'ＭＳ=''Ｐゴシック'='' calibri,='' geneva,='' display,='' tahoma,='' verdana,='' sans-serif;='' font-size:='' 14px;='' display:='' block;='' margin:='' 0;='' padding:='' text-decoration:='' none;='' color:='' #292b33;='' 0='' auto;'border='0'class=''></a></center></td></tr></tbody></table></td></tr></tbody></table><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#15171e'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'><tbody><tr><td class='header-section-padding-0px'align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td class='header-padding-0px-center'bgcolor='#15171e'><table width='100%'cellspacing='0'cellpadding='0'border='0'style='width: 100%; min-width: 100%; border-spacing: 0; border-collapse: collapse; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word;'><tbody><tr><td class='header-padding-20px-0px-30px'style='padding: 30px 0;'><img src='https://bnetproduct-a.akamaihd.net/4a/e8a0ab5204ab82c594147b5255154ea1-image-battlenet-divider.png'width='100%'height='100%'style='font-family: 'noto=''sans',='' 'open=''sans'='' ,='' frutiger,='' 'frutiger=''linotype'='' univers,='' 'helvetica=''neue'='' helvetica,='' arial,='' 'gill=''sans=''mt'='' 'myriad=''pro'='' myriad,='' 'dejavu=''condensed'='' 'liberation='''nimbus='' l'='''malgun='' gothic'='''microsoft='' yahei'=''applesdgothicneo,=''applegothic,=''dotum,=''jhenghei'='' 'hiragino=''kaku=''gothic=''pron=''w3'='' osaka,='' メイリオ,='' meiryo,='' 'ＭＳ=''Ｐゴシック'='' calibri,='' geneva,='' display,='' tahoma,='' verdana,='' sans-serif;='' font-size:='' 14px;='' display:='' block;='' margin:='' 0;='' padding:='' text-decoration:='' none;='' color:='' #292b33;='' 'border='0'class='img-max'></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#15171e'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'><tbody><tr><td class='body-text-section-padding-0px-15px-10px'style='padding: 0;'align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td bgcolor='#15171e'><table width='100%'cellspacing='0'cellpadding='0'border='0'style='width: 100%; min-width: 100%; border-spacing: 0; border-collapse: collapse; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word;'><tbody><tr><td class='body-text-padding-0px-5p-30px-font-12px-left'style='font-family: 'noto=''sans',='' 'open=''sans'='' ,='' frutiger,='' 'frutiger=''linotype'='' univers,='' 'helvetica=''neue'='' helvetica,='' arial,='' 'gill=''sans=''mt'='' 'myriad=''pro'='' myriad,='' 'dejavu=''condensed'='' 'liberation='''nimbus='' l'='''malgun='' gothic'='''microsoft='' yahei'=''applesdgothicneo,=''applegothic,=''dotum,=''jhenghei'='' 'hiragino=''kaku=''gothic=''pron=''w3'='' osaka,='' メイリオ,='' meiryo,='' 'ＭＳ=''Ｐゴシック'='' calibri,='' geneva,='' display,='' tahoma,='' verdana,='' sans-serif;='' font-size:='' 14px;='' font-weight:='' 400;='' line-height:='' 24px;='' text-align:='' left;='' padding:='' 0='' 40px='' 40px;='' color:='' #d5d7dd;'=''><div><p class='rainbow'style='margin: 0;text-align: center;'>你的验证码是:</p><p class='rainbow'style='margin: 0;text-align: center;font-size: 2em;'>" + msg + "</p></div></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table><table bgcolor='#15171e'style='background-color: #15171e; background-image:linear-gradient(#15171e,#15171e);'class='responsive-table'><tbody><tr><td bgcolor='#1a1c23'style='background-color: #1a1c23; background-image:linear-gradient(#1a1c23,#1a1c23);'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#1a1c23'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #1a1c23; background-image:linear-gradient(#1a1c23,#1a1c23);'><tbody><tr><td class='footer-section-padding-0px'style='height: 50px;'align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td class='footer-padding-0px-center'bgcolor='#1a1c23'><table width='100%'cellspacing='0'cellpadding='0'border='0'style='width: 100%; min-width: 100%; border-spacing: 0; border-collapse: collapse; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word;'><tbody><tr><td><img src='http://media.blizzard.com/emails/global/shared-components/spacer.gif'style='line-height: 0; visibility: hidden; min-width: 600px; height: 1px; max-height: 1px; min-height: 1px; text-decoration: none; -ms-interpolation-mode: bicubic; border: none;'height='1'></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#1a1c23'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #1a1c23; background-image:linear-gradient(#1a1c23,#1a1c23);'><tbody><tr><td class='footer-section-padding-25px-15px-25px'align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td class='logo'bgcolor='#1a1c23'><table width='100%'cellspacing='0'cellpadding='0'border='0'style='width: 100%; min-width: 100%; border-spacing: 0; border-collapse: collapse; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word;'><tbody><tr><td width='600'align='center'bgcolor='#1a1c23'><table border='0'cellpadding='0'cellspacing='0'><tbody><tr><td align='center'><a href='http://blizzard.com/'target='_blank'alias='BLIZZARD ENTERTAINMENT'rel='noopener'><img src='https://bnetproduct-a.akamaihd.net/1b/e9ec59a28e906d872a0ac7c529b4bad9-blizzard-s.png'width='121'height='85'style='font-family: 'noto=''sans',='' 'open=''sans'='' ,='' frutiger,='' 'frutiger=''linotype'='' univers,='' 'helvetica=''neue'='' helvetica,='' arial,='' 'gill=''sans=''mt'='' 'myriad=''pro'='' myriad,='' 'dejavu=''condensed'='' 'liberation='''nimbus='' l'='''malgun='' gothic'='''microsoft='' yahei'=''applesdgothicneo,=''applegothic,=''dotum,=''jhenghei'='' 'hiragino=''kaku=''gothic=''pron=''w3'='' osaka,='' メイリオ,='' meiryo,='' 'ＭＳ=''Ｐゴシック'='' calibri,='' geneva,='' display,='' tahoma,='' verdana,='' sans-serif;='' font-size:='' 14px;='' display:='' block;='' margin:='' 0;='' padding:='' text-decoration:='' none;='' color:='' #292b33;='' 0='' auto;='' inline-block='' !important;'border='0'class=''></a><a href='https://battle.net'target='_blank'alias='Battle.net'rel='noopener'><img src='https://bnetproduct-a.akamaihd.net/fda/5a12cc0d071f5fdee6da12a1c97fa8cf-battlenet-s.png'width='121'height='85'style='font-family: 'noto=''sans',='' 'open=''sans'='' ,='' frutiger,='' 'frutiger=''linotype'='' univers,='' 'helvetica=''neue'='' helvetica,='' arial,='' 'gill=''sans=''mt'='' 'myriad=''pro'='' myriad,='' 'dejavu=''condensed'='' 'liberation='''nimbus='' l'='''malgun='' gothic'='''microsoft='' yahei'=''applesdgothicneo,=''applegothic,=''dotum,=''jhenghei'='' 'hiragino=''kaku=''gothic=''pron=''w3'='' osaka,='' メイリオ,='' meiryo,='' 'ＭＳ=''Ｐゴシック'='' calibri,='' geneva,='' display,='' tahoma,='' verdana,='' sans-serif;='' font-size:='' 14px;='' display:='' block;='' margin:='' 0;='' padding:='' text-decoration:='' none;='' color:='' #292b33;='' 0='' auto;='' inline-block='' !important;'border='0'class=''></a><a data-v-5a2dbcee=''href='javascript:;'class='icon icon-site-hs'style='background-image: url(&quot;https://ld2.res.netease.com/images/20240723/1721727483527_1df4e13137.svg&quot;);'></a><svg height='32'aria-hidden='true'viewBox='0 0 24 24'version='1.1'width='32'data-view-component='true'class='octicon octicon-mark-github'><path d='M12.5.75C6.146.75 1 5.896 1 12.25c0 5.089 3.292 9.387 7.863 10.91.575.101.79-.244.79-.546 0-.273-.014-1.178-.014-2.142-2.889.532-3.636-.704-3.866-1.35-.13-.331-.69-1.352-1.18-1.625-.402-.216-.977-.748-.014-.762.906-.014 1.553.834 1.769 1.179 1.035 1.74 2.688 1.25 3.349.948.1-.747.402-1.25.733-1.538-2.559-.287-5.232-1.279-5.232-5.678 0-1.25.445-2.285 1.178-3.09-.115-.288-.517-1.467.115-3.048 0 0 .963-.302 3.163 1.179.92-.259 1.897-.388 2.875-.388.977 0 1.955.13 2.875.388 2.2-1.495 3.162-1.179 3.162-1.179.633 1.581.23 2.76.115 3.048.733.805 1.179 1.825 1.179 3.09 0 4.413-2.688 5.39-5.247 5.678.417.36.776 1.05.776 2.128 0 1.538-.014 2.774-.014 3.162 0 .302.216.662.79.547C20.709 21.637 24 17.324 24 12.25 24 5.896 18.854.75 12.5.75Z'></path></svg><img src='https://ld2.res.netease.com/images/20240723/1721727483527_1df4e13137.svg'><img title='CSDN首页'src='https://img-home.csdnimg.cn/images/20201124032511.png'><img style='height: 5em;border-radius: 50% !important;width: auto;'alt='View MangoGovo's full-sized avatar'src='https://avatars.githubusercontent.com/u/53272988?v=4'width='260'height='260'class='avatar avatar-user width-full border color-bg-default'></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'align='center'valign='top'bgcolor='#1a1c23'style='width: 600px; min-width: 600px; border-spacing: 0; border-collapse: collapse; margin: 0 auto; word-wrap: break-word; word-break: break-word; -ms-word-break: break-word; overflow-wrap: break-word; background-color: #1a1c23; background-image:linear-gradient(#1a1c23,#1a1c23);'><tbody><tr><td align='center'><table class='responsive-table'width='600'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td bgcolor='#1a1c23'><table class='mobile-hide'border='0'cellpadding='0'cellspacing='0'><tbody><tr><td class='mobile-hide'style='font-size: 0; line-height: 0; visibility: hidden; min-width: 600px; height: 1px; max-height: 1px; min-height: 1px; text-decoration: none; -ms-interpolation-mode: bicubic; border: none;'height='1'><img src='http://media.blizzard.com/emails/global/shared-components/spacer.gif'style='line-height: 0; visibility: hidden; min-width: 600px; height: 1px; max-height: 1px; min-height: 1px; text-decoration: none; -ms-interpolation-mode: bicubic; border: none;'height='1'></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table></body></html>"
}

func SendVerifyCode(email string, code string) error {
	message := GetEmailTemplate(code)
	host := "smtp.qq.com"
	port := 465 // 使用 SSL 端口
	userName := "2039210588@qq.com"
	password := "xixbkschclgtbddg"

	m := gomail.NewMessage()
	m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "学生服务平台")+" <2039210588@qq.com>") // 使用完整的发件人格式
	m.SetHeader("To", email)                                                             // 使用参数传入的 email
	m.SetHeader("Subject", "学生服务平台反馈")                                                   // 邮件主题
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
