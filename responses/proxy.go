package responses

type ProxyElement struct {
  Proxy    *ProxyResponse        `json:"proxy"`
  Behavior *Behavior             `json:"_behaviors,omitempty"`
}

type ProxyResponse struct {
  To         string            `json:"to,omitempty"`
  Mode       string            `json:"mode,omitempty"`
}

func (i ProxyElement) Type() string {
  return "proxy"
}

type ProxyBuilder struct {
  proxy         *ProxyResponse
  waitTime   int
  decorateFn string
}

func (builder *ProxyBuilder) To(to string) *ProxyBuilder {
  builder.proxy.To = to
  return builder
}

func (builder *ProxyBuilder) Mode(mode string) *ProxyBuilder {
  builder.proxy.Mode = mode
  return builder
}

func (builder *ProxyBuilder) Wait(waitTime int) *ProxyBuilder {
  if waitTime > 0 {
    builder.waitTime = waitTime
  }
  return builder
}

func (builder *ProxyBuilder) Decorate(decorateFn string) *ProxyBuilder {
  builder.decorateFn = decorateFn
  return builder
}

func (builder *ProxyBuilder) Build() ProxyElement {
  proxy := ProxyElement{Proxy: builder.proxy}
  if builder.waitTime > 0 || len(builder.decorateFn) > 0 {
    proxy.Behavior = &Behavior{builder.waitTime, builder.decorateFn}
  }
  return proxy
}

func Proxy() *ProxyBuilder {
  return &ProxyBuilder{proxy: &ProxyResponse{}}
}
