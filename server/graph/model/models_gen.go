// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AdminLoginInput struct {
	AdminSecret string `json:"admin_secret"`
}

type AdminLoginResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Message     string  `json:"message"`
	AccessToken *string `json:"access_token"`
	ExpiresAt   *int64  `json:"expires_at"`
	User        *User   `json:"user"`
}

type Config struct {
	AdminSecret                *string  `json:"ADMIN_SECRET"`
	DatabaseType               *string  `json:"DATABASE_TYPE"`
	DatabaseURL                *string  `json:"DATABASE_URL"`
	DatabaseName               *string  `json:"DATABASE_NAME"`
	SMTPHost                   *string  `json:"SMTP_HOST"`
	SMTPPort                   *string  `json:"SMTP_PORT"`
	SMTPUsername               *string  `json:"SMTP_USERNAME"`
	SMTPPassword               *string  `json:"SMTP_PASSWORD"`
	SenderEmail                *string  `json:"SENDER_EMAIL"`
	JwtType                    *string  `json:"JWT_TYPE"`
	JwtSecret                  *string  `json:"JWT_SECRET"`
	AllowedOrigins             []string `json:"ALLOWED_ORIGINS"`
	AuthorizerURL              *string  `json:"AUTHORIZER_URL"`
	AppURL                     *string  `json:"APP_URL"`
	RedisURL                   *string  `json:"REDIS_URL"`
	CookieName                 *string  `json:"COOKIE_NAME"`
	ResetPasswordURL           *string  `json:"RESET_PASSWORD_URL"`
	DisableEmailVerification   *bool    `json:"DISABLE_EMAIL_VERIFICATION"`
	DisableBasicAuthentication *bool    `json:"DISABLE_BASIC_AUTHENTICATION"`
	DisableMagicLinkLogin      *bool    `json:"DISABLE_MAGIC_LINK_LOGIN"`
	DisableLoginPage           *bool    `json:"DISABLE_LOGIN_PAGE"`
	Roles                      []string `json:"ROLES"`
	ProtectedRoles             []string `json:"PROTECTED_ROLES"`
	DefaultRoles               []string `json:"DEFAULT_ROLES"`
	JwtRoleClaim               *string  `json:"JWT_ROLE_CLAIM"`
	GoogleClientID             *string  `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret         *string  `json:"GOOGLE_CLIENT_SECRET"`
	GithubClientID             *string  `json:"GITHUB_CLIENT_ID"`
	GithubClientSecret         *string  `json:"GITHUB_CLIENT_SECRET"`
	FacebookClientID           *string  `json:"FACEBOOK_CLIENT_ID"`
	FacebookClientSecret       *string  `json:"FACEBOOK_CLIENT_SECRET"`
	OrganizationName           *string  `json:"ORGANIZATION_NAME"`
	OrganizationLogo           *string  `json:"ORGANIZATION_LOGO"`
}

type DeleteUserInput struct {
	Email string `json:"email"`
}

type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type ForgotPasswordInput struct {
	Email string `json:"email"`
}

type LoginInput struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type MagicLinkLoginInput struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

type Meta struct {
	Version                      string `json:"version"`
	IsGoogleLoginEnabled         bool   `json:"is_google_login_enabled"`
	IsFacebookLoginEnabled       bool   `json:"is_facebook_login_enabled"`
	IsGithubLoginEnabled         bool   `json:"is_github_login_enabled"`
	IsEmailVerificationEnabled   bool   `json:"is_email_verification_enabled"`
	IsBasicAuthenticationEnabled bool   `json:"is_basic_authentication_enabled"`
	IsMagicLinkLoginEnabled      bool   `json:"is_magic_link_login_enabled"`
}

type ResendVerifyEmailInput struct {
	Email      string `json:"email"`
	Identifier string `json:"identifier"`
}

type ResetPasswordInput struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Response struct {
	Message string `json:"message"`
}

type SignUpInput struct {
	Email           string   `json:"email"`
	GivenName       *string  `json:"given_name"`
	FamilyName      *string  `json:"family_name"`
	MiddleName      *string  `json:"middle_name"`
	Nickname        *string  `json:"nickname"`
	Gender          *string  `json:"gender"`
	Birthdate       *string  `json:"birthdate"`
	PhoneNumber     *string  `json:"phone_number"`
	Picture         *string  `json:"picture"`
	Password        string   `json:"password"`
	ConfirmPassword string   `json:"confirm_password"`
	Roles           []string `json:"roles"`
}

type UpdateConfigInput struct {
	AdminSecret                *string  `json:"ADMIN_SECRET"`
	DatabaseType               *string  `json:"DATABASE_TYPE"`
	DatabaseURL                *string  `json:"DATABASE_URL"`
	DatabaseName               *string  `json:"DATABASE_NAME"`
	SMTPHost                   *string  `json:"SMTP_HOST"`
	SMTPPort                   *string  `json:"SMTP_PORT"`
	SenderEmail                *string  `json:"SENDER_EMAIL"`
	SenderPassword             *string  `json:"SENDER_PASSWORD"`
	JwtType                    *string  `json:"JWT_TYPE"`
	JwtSecret                  *string  `json:"JWT_SECRET"`
	AllowedOrigins             []string `json:"ALLOWED_ORIGINS"`
	AuthorizerURL              *string  `json:"AUTHORIZER_URL"`
	AppURL                     *string  `json:"APP_URL"`
	RedisURL                   *string  `json:"REDIS_URL"`
	CookieName                 *string  `json:"COOKIE_NAME"`
	ResetPasswordURL           *string  `json:"RESET_PASSWORD_URL"`
	DisableEmailVerification   *bool    `json:"DISABLE_EMAIL_VERIFICATION"`
	DisableBasicAuthentication *bool    `json:"DISABLE_BASIC_AUTHENTICATION"`
	DisableMagicLinkLogin      *bool    `json:"DISABLE_MAGIC_LINK_LOGIN"`
	DisableLoginPage           *bool    `json:"DISABLE_LOGIN_PAGE"`
	Roles                      []string `json:"ROLES"`
	ProtectedRoles             []string `json:"PROTECTED_ROLES"`
	DefaultRoles               []string `json:"DEFAULT_ROLES"`
	JwtRoleClaim               *string  `json:"JWT_ROLE_CLAIM"`
	GoogleClientID             *string  `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret         *string  `json:"GOOGLE_CLIENT_SECRET"`
	GithubClientID             *string  `json:"GITHUB_CLIENT_ID"`
	GithubClientSecret         *string  `json:"GITHUB_CLIENT_SECRET"`
	FacebookClientID           *string  `json:"FACEBOOK_CLIENT_ID"`
	FacebookClientSecret       *string  `json:"FACEBOOK_CLIENT_SECRET"`
	OrganizationName           *string  `json:"ORGANIZATION_NAME"`
	OrganizationLogo           *string  `json:"ORGANIZATION_LOGO"`
}

type UpdateProfileInput struct {
	OldPassword        *string `json:"old_password"`
	NewPassword        *string `json:"new_password"`
	ConfirmNewPassword *string `json:"confirm_new_password"`
	Email              *string `json:"email"`
	GivenName          *string `json:"given_name"`
	FamilyName         *string `json:"family_name"`
	MiddleName         *string `json:"middle_name"`
	Nickname           *string `json:"nickname"`
	Gender             *string `json:"gender"`
	Birthdate          *string `json:"birthdate"`
	PhoneNumber        *string `json:"phone_number"`
	Picture            *string `json:"picture"`
}

type UpdateUserInput struct {
	ID          string    `json:"id"`
	Email       *string   `json:"email"`
	GivenName   *string   `json:"given_name"`
	FamilyName  *string   `json:"family_name"`
	MiddleName  *string   `json:"middle_name"`
	Nickname    *string   `json:"nickname"`
	Gender      *string   `json:"gender"`
	Birthdate   *string   `json:"birthdate"`
	PhoneNumber *string   `json:"phone_number"`
	Picture     *string   `json:"picture"`
	Roles       []*string `json:"roles"`
}

type User struct {
	ID                  string   `json:"id"`
	Email               string   `json:"email"`
	EmailVerified       bool     `json:"email_verified"`
	SignupMethods       string   `json:"signup_methods"`
	GivenName           *string  `json:"given_name"`
	FamilyName          *string  `json:"family_name"`
	MiddleName          *string  `json:"middle_name"`
	Nickname            *string  `json:"nickname"`
	PreferredUsername   *string  `json:"preferred_username"`
	Gender              *string  `json:"gender"`
	Birthdate           *string  `json:"birthdate"`
	PhoneNumber         *string  `json:"phone_number"`
	PhoneNumberVerified *bool    `json:"phone_number_verified"`
	Picture             *string  `json:"picture"`
	Roles               []string `json:"roles"`
	CreatedAt           *int64   `json:"created_at"`
	UpdatedAt           *int64   `json:"updated_at"`
}

type VerificationRequest struct {
	ID         string  `json:"id"`
	Identifier *string `json:"identifier"`
	Token      *string `json:"token"`
	Email      *string `json:"email"`
	Expires    *int64  `json:"expires"`
	CreatedAt  *int64  `json:"created_at"`
	UpdatedAt  *int64  `json:"updated_at"`
}

type VerifyEmailInput struct {
	Token string `json:"token"`
}
