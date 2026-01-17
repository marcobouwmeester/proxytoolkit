package bruno

type BrunoRequest struct {
	Meta     MetaBlock
	Request  RequestBlock
	Params   *ParamsBlock
	Headers  *HeadersBlock
	Body     *BodyBlock
	Settings *SettingsBlock
}

type MetaBlock struct {
	Name string
	Type string // always "http"
	Seq  int
}

type RequestBlock struct {
	Method string // GET, POST, PUT, ...
	URL    string
	Auth   *string // e.g. "inherit"
}

type ParamsBlock struct {
	Query map[string]string
	Path  map[string]string
}

type HeadersBlock struct {
	Values map[string]string
}

type BodyBlock struct {
	Type string // json, text, form, none
	Raw  string // raw body content
}

type SettingsBlock struct {
	EncodeURL bool
	Timeout   int
}
