package reader

import "fmt"

var (
	forgotPasswordText = `
	<!doctype html>
	<html lang="en">
	<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta http-equiv="x-ua-compatible" content="ie=edge" />
	<title>Password reset</title>
	<meta name="x-apple-disable-message-reformatting" />
	<style>
		:root {
			--color-bg: #ffffff;
			--color-fg: #2d3748;
		}
		a { 
			text-decoration: none; 
		}
		img { 
			border: 0; 
			line-height: 100%%; 
			vertical-align: middle; 
		}
		table { 
			border-collapse: collapse; 
		}
		.pos-middle {
			margin: 0 auto !important;
			margin-bottom: 10px !important;
		}
		.auth-btn {
			display: inline-block !important;
			width: auto !important;
			padding: 12px 18px !important;
			background: #2d3748 !important;
			color: #ffffff !important;
			border: 1.5px solid #2d3748 !important;
			font-size: 14px !important;
			font-weight: 700 !important;
			letter-spacing: 0.02em !important;
			text-transform: none !important;
			cursor: pointer !important;
			text-align: center !important;
			transition: background 0.2s, color 0.2s;
			border-radius: 0 !important;
			font-family: 'Inter', 'Helvetica Neue', Arial, sans-serif !important;
		}
		.auth-btn:hover,
		.auth-btn:focus {
			background: #2d3748 !important;
			color: #ffffff !important;
			outline: 1px solid #2d3748 !important;
			transform: translateY(-1px) !important;
			transition: all 0.2s ease !important;
		}
		.logo {
			display: block !important;
			margin: 0 auto !important;
			width: 100px !important;
			height: 38px !important;
		}
		@media (prefers-color-scheme: dark) {
		.card { 
			background: #ffffff !important; 
		}
		.text { 
			text-align: center !important;
			color: #2d3748 !important; 
		}
		.muted { color: #9CA3AF !important; }
		.divider { border-color: #374151 !important; }
		}
		@media only screen and (max-width: 620px) {
		.container { width: 100% !important; padding: 0 16px !important; }
		.card { padding: 20px !important; }
		.btn { display: block !important; width: 100%% !important; text-align: center !important; }
		}
	</style>
	</head>
	<body style="margin:0; padding:0; background:#f6f8fb;">
	<div style="display:none; overflow:hidden; line-height:1px; opacity:0; max-height:0; max-width:0;">
		Password reset link
	</div>

	<table role="presentation" width="100%%" style="width:100%%; background:#ffffff;" cellpadding="0" cellspacing="0">
		<tr>
		<td align="center" style="padding: 32px 12px;">
			<table role="presentation" class="container" width="600" style="width:600px; max-width:100%;" cellpadding="0" cellspacing="0">
			<tr>
				<td class="card" style="background:#ffffff; border:1px solid #e2e8f0; box-shadow:0 2px 8px rgba(0,0,0,0.05); padding:28px;">

				<table class="table-head" role="presentation" width="100%%" cellpadding="0" cellspacing="0">
					<tr>
					<td style="padding-bottom: 12px;">
                    	<img src="https://e-toolnet.ru/assets/mainlogo.svg" alt="Logo" class="logo" />
					</td>
					</tr>
				</table>

				<hr class="divider" style="border:0; border-top:1px solid #e2e8f0; margin: 4px 0 20px%%" />

				<h1 class="text" style="margin:0 0 12px; font-family:Inter, 'Helvetica Neue', Arial, sans-serif; font-size:22px; line-height:1.3; color:#2d3748; font-weight:700;">
					Password reset
				</h1>
				<p class="text" style="margin:0 0 16px; font-family:Inter, 'Helvetica Neue', Arial, sans-serif; font-size:14px; line-height:1.6; color:#2d3748;">
					You requested a password reset for your account. Click the button below to continue. If you did not request this, please ignore this email.
				</p>

				<table class="pos-middle" role="presentation" cellpadding="0" cellspacing="0" style="margin: 22px 0;">
					<tr>
					<td>
						<a class="auth-btn" href="%[1]s" target="_blank" rel="noopener" >
						Reset password
						</a>
					</td>
					</tr>
				</table>

				<p class="muted" style="margin:0 0 8px; font-family:Inter, 'Helvetica Neue', Arial, sans-serif; font-size:12px; line-height:1.6; color:#a0aec0;">
					If the button doesn't work, copy and paste the link into your browser:
				</p>
				<p style="margin:0 0 16px; word-break:break-all; font-family:Inter, 'Helvetica Neue', Arial, sans-serif; font-size:12px; line-height:1.6; color:#2d3748;">
					<a href="%[1]s" target="_blank" rel="noopener" style="color:#2d3748; text-decoration:underline;">%[1]s</a>
				</p>

				<hr class="divider" style="border:0; border-top:1px solid #e2e8f0; margin: 16px 0;" />

				<p class="muted" style="margin:0; font-family:Inter, 'Helvetica Neue', Arial, sans-serif; font-size:12px; line-height:1.6; color:#a0aec0;">
					This is an automated email, please do not reply. If you have any questions, visit the support page.
				</p>

				</td>
			</tr>
			</table>
		</td>
		</tr>
	</table>
	</body>
	</html>`
)

func getForgotPasswordText(domain, pathToToken, token string) string {
	tmp := fmt.Sprintf("%v%v%v", domain, pathToToken, token)
	return fmt.Sprintf(forgotPasswordText, tmp)
}
