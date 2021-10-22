package tools

import (
	"context"
	"runtime"

	"strings"

	"go.elastic.co/apm"
)

type apmSpan string

// Padroniza o ApmSpan pegando o nome do pacote e mêtodo no qual ele foi invocado.
func ApmSpan(c context.Context, self apmSpan) (*apm.Span, context.Context) {
	return apm.StartSpan(c, self.Name(), self.Package())
}

// Inicia uma instância com pacote e mêtodo de origem.
func Self() apmSpan {
	pc, _, _, _ := runtime.Caller(1)
	return apmSpan(runtime.FuncForPC(pc).Name())
}

func (s apmSpan) Name() string {
	split := strings.Split(string(s), "/")
	return split[len(split)-1:][0]
}

func (s apmSpan) Func() string {
	return strings.Split(s.Name(), ".")[1]
}

func (s apmSpan) Package() string {
	return strings.Split(s.Name(), ".")[0]
}
