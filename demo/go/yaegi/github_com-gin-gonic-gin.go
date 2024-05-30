// Code generated by 'yaegi extract github.com/gin-gonic/gin'. DO NOT EDIT.

package yaegi

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"go/constant"
	"go/token"
	"net"
	"net/http"
	"reflect"
)

func init() {
	Symbols["github.com/gin-gonic/gin/gin"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AuthProxyUserKey":                       reflect.ValueOf(constant.MakeFromLiteral("\"proxy_user\"", token.STRING, 0)),
		"AuthUserKey":                            reflect.ValueOf(constant.MakeFromLiteral("\"user\"", token.STRING, 0)),
		"BasicAuth":                              reflect.ValueOf(gin.BasicAuth),
		"BasicAuthForProxy":                      reflect.ValueOf(gin.BasicAuthForProxy),
		"BasicAuthForRealm":                      reflect.ValueOf(gin.BasicAuthForRealm),
		"Bind":                                   reflect.ValueOf(gin.Bind),
		"BindKey":                                reflect.ValueOf(constant.MakeFromLiteral("\"_gin-gonic/gin/bindkey\"", token.STRING, 0)),
		"BodyBytesKey":                           reflect.ValueOf(constant.MakeFromLiteral("\"_gin-gonic/gin/bodybyteskey\"", token.STRING, 0)),
		"ContextKey":                             reflect.ValueOf(constant.MakeFromLiteral("\"_gin-gonic/gin/contextkey\"", token.STRING, 0)),
		"ContextRequestKey":                      reflect.ValueOf(gin.ContextRequestKey),
		"CreateTestContext":                      reflect.ValueOf(gin.CreateTestContext),
		"CreateTestContextOnly":                  reflect.ValueOf(gin.CreateTestContextOnly),
		"CustomRecovery":                         reflect.ValueOf(gin.CustomRecovery),
		"CustomRecoveryWithWriter":               reflect.ValueOf(gin.CustomRecoveryWithWriter),
		"DebugMode":                              reflect.ValueOf(constant.MakeFromLiteral("\"debug\"", token.STRING, 0)),
		"DebugPrintFunc":                         reflect.ValueOf(&gin.DebugPrintFunc).Elem(),
		"DebugPrintRouteFunc":                    reflect.ValueOf(&gin.DebugPrintRouteFunc).Elem(),
		"Default":                                reflect.ValueOf(gin.Default),
		"DefaultErrorWriter":                     reflect.ValueOf(&gin.DefaultErrorWriter).Elem(),
		"DefaultWriter":                          reflect.ValueOf(&gin.DefaultWriter).Elem(),
		"Dir":                                    reflect.ValueOf(gin.Dir),
		"DisableBindValidation":                  reflect.ValueOf(gin.DisableBindValidation),
		"DisableConsoleColor":                    reflect.ValueOf(gin.DisableConsoleColor),
		"EnableJsonDecoderDisallowUnknownFields": reflect.ValueOf(gin.EnableJsonDecoderDisallowUnknownFields),
		"EnableJsonDecoderUseNumber":             reflect.ValueOf(gin.EnableJsonDecoderUseNumber),
		"EnvGinMode":                             reflect.ValueOf(constant.MakeFromLiteral("\"GIN_MODE\"", token.STRING, 0)),
		"ErrorLogger":                            reflect.ValueOf(gin.ErrorLogger),
		"ErrorLoggerT":                           reflect.ValueOf(gin.ErrorLoggerT),
		"ErrorTypeAny":                           reflect.ValueOf(gin.ErrorTypeAny),
		"ErrorTypeBind":                          reflect.ValueOf(gin.ErrorTypeBind),
		"ErrorTypeNu":                            reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"ErrorTypePrivate":                       reflect.ValueOf(gin.ErrorTypePrivate),
		"ErrorTypePublic":                        reflect.ValueOf(gin.ErrorTypePublic),
		"ErrorTypeRender":                        reflect.ValueOf(gin.ErrorTypeRender),
		"ForceConsoleColor":                      reflect.ValueOf(gin.ForceConsoleColor),
		"IsDebugging":                            reflect.ValueOf(gin.IsDebugging),
		"Logger":                                 reflect.ValueOf(gin.Logger),
		"LoggerWithConfig":                       reflect.ValueOf(gin.LoggerWithConfig),
		"LoggerWithFormatter":                    reflect.ValueOf(gin.LoggerWithFormatter),
		"LoggerWithWriter":                       reflect.ValueOf(gin.LoggerWithWriter),
		"MIMEHTML":                               reflect.ValueOf(constant.MakeFromLiteral("\"text/html\"", token.STRING, 0)),
		"MIMEJSON":                               reflect.ValueOf(constant.MakeFromLiteral("\"application/json\"", token.STRING, 0)),
		"MIMEMultipartPOSTForm":                  reflect.ValueOf(constant.MakeFromLiteral("\"multipart/form-data\"", token.STRING, 0)),
		"MIMEPOSTForm":                           reflect.ValueOf(constant.MakeFromLiteral("\"application/x-www-form-urlencoded\"", token.STRING, 0)),
		"MIMEPlain":                              reflect.ValueOf(constant.MakeFromLiteral("\"text/plain\"", token.STRING, 0)),
		"MIMETOML":                               reflect.ValueOf(constant.MakeFromLiteral("\"application/toml\"", token.STRING, 0)),
		"MIMEXML":                                reflect.ValueOf(constant.MakeFromLiteral("\"application/xml\"", token.STRING, 0)),
		"MIMEXML2":                               reflect.ValueOf(constant.MakeFromLiteral("\"text/xml\"", token.STRING, 0)),
		"MIMEYAML":                               reflect.ValueOf(constant.MakeFromLiteral("\"application/x-yaml\"", token.STRING, 0)),
		"Mode":                                   reflect.ValueOf(gin.Mode),
		"New":                                    reflect.ValueOf(gin.New),
		"PlatformCloudflare":                     reflect.ValueOf(constant.MakeFromLiteral("\"CF-Connecting-IP\"", token.STRING, 0)),
		"PlatformFlyIO":                          reflect.ValueOf(constant.MakeFromLiteral("\"Fly-Client-IP\"", token.STRING, 0)),
		"PlatformGoogleAppEngine":                reflect.ValueOf(constant.MakeFromLiteral("\"X-Appengine-Remote-Addr\"", token.STRING, 0)),
		"Recovery":                               reflect.ValueOf(gin.Recovery),
		"RecoveryWithWriter":                     reflect.ValueOf(gin.RecoveryWithWriter),
		"ReleaseMode":                            reflect.ValueOf(constant.MakeFromLiteral("\"release\"", token.STRING, 0)),
		"SetMode":                                reflect.ValueOf(gin.SetMode),
		"TestMode":                               reflect.ValueOf(constant.MakeFromLiteral("\"test\"", token.STRING, 0)),
		"Version":                                reflect.ValueOf(constant.MakeFromLiteral("\"v1.10.0\"", token.STRING, 0)),
		"WrapF":                                  reflect.ValueOf(gin.WrapF),
		"WrapH":                                  reflect.ValueOf(gin.WrapH),

		// type definitions
		"Accounts":           reflect.ValueOf((*gin.Accounts)(nil)),
		"Context":            reflect.ValueOf((*gin.Context)(nil)),
		"ContextKeyType":     reflect.ValueOf((*gin.ContextKeyType)(nil)),
		"Engine":             reflect.ValueOf((*gin.Engine)(nil)),
		"Error":              reflect.ValueOf((*gin.Error)(nil)),
		"ErrorType":          reflect.ValueOf((*gin.ErrorType)(nil)),
		"H":                  reflect.ValueOf((*gin.H)(nil)),
		"HandlerFunc":        reflect.ValueOf((*gin.HandlerFunc)(nil)),
		"HandlersChain":      reflect.ValueOf((*gin.HandlersChain)(nil)),
		"IRouter":            reflect.ValueOf((*gin.IRouter)(nil)),
		"IRoutes":            reflect.ValueOf((*gin.IRoutes)(nil)),
		"LogFormatter":       reflect.ValueOf((*gin.LogFormatter)(nil)),
		"LogFormatterParams": reflect.ValueOf((*gin.LogFormatterParams)(nil)),
		"LoggerConfig":       reflect.ValueOf((*gin.LoggerConfig)(nil)),
		"Negotiate":          reflect.ValueOf((*gin.Negotiate)(nil)),
		"OptionFunc":         reflect.ValueOf((*gin.OptionFunc)(nil)),
		"Param":              reflect.ValueOf((*gin.Param)(nil)),
		"Params":             reflect.ValueOf((*gin.Params)(nil)),
		"RecoveryFunc":       reflect.ValueOf((*gin.RecoveryFunc)(nil)),
		"ResponseWriter":     reflect.ValueOf((*gin.ResponseWriter)(nil)),
		"RouteInfo":          reflect.ValueOf((*gin.RouteInfo)(nil)),
		"RouterGroup":        reflect.ValueOf((*gin.RouterGroup)(nil)),
		"RoutesInfo":         reflect.ValueOf((*gin.RoutesInfo)(nil)),
		"Skipper":            reflect.ValueOf((*gin.Skipper)(nil)),

		// interface wrapper definitions
		"_IRouter":        reflect.ValueOf((*_github_com_gin_gonic_gin_IRouter)(nil)),
		"_IRoutes":        reflect.ValueOf((*_github_com_gin_gonic_gin_IRoutes)(nil)),
		"_ResponseWriter": reflect.ValueOf((*_github_com_gin_gonic_gin_ResponseWriter)(nil)),
	}
}

// _github_com_gin_gonic_gin_IRouter is an interface wrapper for IRouter type
type _github_com_gin_gonic_gin_IRouter struct {
	IValue        interface{}
	WAny          func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WDELETE       func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WGET          func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WGroup        func(a0 string, a1 ...gin.HandlerFunc) *gin.RouterGroup
	WHEAD         func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WHandle       func(a0 string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes
	WMatch        func(a0 []string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes
	WOPTIONS      func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WPATCH        func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WPOST         func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WPUT          func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WStatic       func(a0 string, a1 string) gin.IRoutes
	WStaticFS     func(a0 string, a1 http.FileSystem) gin.IRoutes
	WStaticFile   func(a0 string, a1 string) gin.IRoutes
	WStaticFileFS func(a0 string, a1 string, a2 http.FileSystem) gin.IRoutes
	WUse          func(a0 ...gin.HandlerFunc) gin.IRoutes
}

func (W _github_com_gin_gonic_gin_IRouter) Any(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WAny(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) DELETE(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WDELETE(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) GET(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WGET(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) Group(a0 string, a1 ...gin.HandlerFunc) *gin.RouterGroup {
	return W.WGroup(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) HEAD(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WHEAD(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) Handle(a0 string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes {
	return W.WHandle(a0, a1, a2...)
}
func (W _github_com_gin_gonic_gin_IRouter) Match(a0 []string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes {
	return W.WMatch(a0, a1, a2...)
}
func (W _github_com_gin_gonic_gin_IRouter) OPTIONS(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WOPTIONS(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) PATCH(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WPATCH(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) POST(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WPOST(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) PUT(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WPUT(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRouter) Static(a0 string, a1 string) gin.IRoutes {
	return W.WStatic(a0, a1)
}
func (W _github_com_gin_gonic_gin_IRouter) StaticFS(a0 string, a1 http.FileSystem) gin.IRoutes {
	return W.WStaticFS(a0, a1)
}
func (W _github_com_gin_gonic_gin_IRouter) StaticFile(a0 string, a1 string) gin.IRoutes {
	return W.WStaticFile(a0, a1)
}
func (W _github_com_gin_gonic_gin_IRouter) StaticFileFS(a0 string, a1 string, a2 http.FileSystem) gin.IRoutes {
	return W.WStaticFileFS(a0, a1, a2)
}
func (W _github_com_gin_gonic_gin_IRouter) Use(a0 ...gin.HandlerFunc) gin.IRoutes {
	return W.WUse(a0...)
}

// _github_com_gin_gonic_gin_IRoutes is an interface wrapper for IRoutes type
type _github_com_gin_gonic_gin_IRoutes struct {
	IValue        interface{}
	WAny          func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WDELETE       func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WGET          func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WHEAD         func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WHandle       func(a0 string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes
	WMatch        func(a0 []string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes
	WOPTIONS      func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WPATCH        func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WPOST         func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WPUT          func(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes
	WStatic       func(a0 string, a1 string) gin.IRoutes
	WStaticFS     func(a0 string, a1 http.FileSystem) gin.IRoutes
	WStaticFile   func(a0 string, a1 string) gin.IRoutes
	WStaticFileFS func(a0 string, a1 string, a2 http.FileSystem) gin.IRoutes
	WUse          func(a0 ...gin.HandlerFunc) gin.IRoutes
}

func (W _github_com_gin_gonic_gin_IRoutes) Any(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WAny(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) DELETE(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WDELETE(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) GET(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WGET(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) HEAD(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WHEAD(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) Handle(a0 string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes {
	return W.WHandle(a0, a1, a2...)
}
func (W _github_com_gin_gonic_gin_IRoutes) Match(a0 []string, a1 string, a2 ...gin.HandlerFunc) gin.IRoutes {
	return W.WMatch(a0, a1, a2...)
}
func (W _github_com_gin_gonic_gin_IRoutes) OPTIONS(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WOPTIONS(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) PATCH(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WPATCH(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) POST(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WPOST(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) PUT(a0 string, a1 ...gin.HandlerFunc) gin.IRoutes {
	return W.WPUT(a0, a1...)
}
func (W _github_com_gin_gonic_gin_IRoutes) Static(a0 string, a1 string) gin.IRoutes {
	return W.WStatic(a0, a1)
}
func (W _github_com_gin_gonic_gin_IRoutes) StaticFS(a0 string, a1 http.FileSystem) gin.IRoutes {
	return W.WStaticFS(a0, a1)
}
func (W _github_com_gin_gonic_gin_IRoutes) StaticFile(a0 string, a1 string) gin.IRoutes {
	return W.WStaticFile(a0, a1)
}
func (W _github_com_gin_gonic_gin_IRoutes) StaticFileFS(a0 string, a1 string, a2 http.FileSystem) gin.IRoutes {
	return W.WStaticFileFS(a0, a1, a2)
}
func (W _github_com_gin_gonic_gin_IRoutes) Use(a0 ...gin.HandlerFunc) gin.IRoutes {
	return W.WUse(a0...)
}

// _github_com_gin_gonic_gin_ResponseWriter is an interface wrapper for ResponseWriter type
type _github_com_gin_gonic_gin_ResponseWriter struct {
	IValue          interface{}
	WCloseNotify    func() <-chan bool
	WFlush          func()
	WHeader         func() http.Header
	WHijack         func() (net.Conn, *bufio.ReadWriter, error)
	WPusher         func() http.Pusher
	WSize           func() int
	WStatus         func() int
	WWrite          func(a0 []byte) (int, error)
	WWriteHeader    func(statusCode int)
	WWriteHeaderNow func()
	WWriteString    func(a0 string) (int, error)
	WWritten        func() bool
}

func (W _github_com_gin_gonic_gin_ResponseWriter) CloseNotify() <-chan bool {
	return W.WCloseNotify()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Flush() {
	W.WFlush()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Header() http.Header {
	return W.WHeader()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return W.WHijack()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Pusher() http.Pusher {
	return W.WPusher()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Size() int {
	return W.WSize()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Status() int {
	return W.WStatus()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Write(a0 []byte) (int, error) {
	return W.WWrite(a0)
}
func (W _github_com_gin_gonic_gin_ResponseWriter) WriteHeader(statusCode int) {
	W.WWriteHeader(statusCode)
}
func (W _github_com_gin_gonic_gin_ResponseWriter) WriteHeaderNow() {
	W.WWriteHeaderNow()
}
func (W _github_com_gin_gonic_gin_ResponseWriter) WriteString(a0 string) (int, error) {
	return W.WWriteString(a0)
}
func (W _github_com_gin_gonic_gin_ResponseWriter) Written() bool {
	return W.WWritten()
}
