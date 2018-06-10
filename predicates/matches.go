package predicates

import "encoding/json"

type matches struct {
	req Request
}

func (p matches) Type() string {
	return "matches"
}

func (p matches) MarshalJSON() ([]byte, error) {
	requestBytes, _ := json.Marshal(p.req)

	requestJSON := string(requestBytes)
	matchesJSON := " { \"matches\" : " + requestJSON + "}"

	return []byte(matchesJSON), nil
}

type MatchesBuilder struct {
	matches matches
}

func Matches() *MatchesBuilder {
	return &MatchesBuilder{matches: matches{req: Request{}}}
}

func (builder *MatchesBuilder) Path(path string) *MatchesBuilder {
	builder.matches.req.Path = path
	return builder
}

func (builder *MatchesBuilder) Method(method string) *MatchesBuilder {
	builder.matches.req.Method = method
	return builder
}

func (builder *MatchesBuilder) Header(header string, value string) *MatchesBuilder {
	if builder.matches.req.Headers == nil {
		builder.matches.req.Headers = map[string]string{}
	}
	builder.matches.req.Headers[header] = value
	return builder
}

func (builder *MatchesBuilder) Query(param string, value string) *MatchesBuilder {
	if builder.matches.req.QueryParams == nil {
		builder.matches.req.QueryParams = map[string]string{}
	}

	builder.matches.req.QueryParams[param] = value
	return builder
}

func (builder *MatchesBuilder) Body(body string) *MatchesBuilder {
	builder.matches.req.Body = body
	return builder
}

func (builder *MatchesBuilder) Build() matches {
	return builder.matches
}
