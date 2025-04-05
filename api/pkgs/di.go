package pkgs

import "github.com/samber/do"

type Di struct {
	injector *do.Injector
}

func NewDi() *Di {
	return &Di{
		injector: do.New(),
	}
}

// Provide registra uma função que retorna um valor para ser injetado
//
// Exemplo:
//
// pkgs.Provide(di, services.NewHealthcheckService)
func Provide[T any](d *Di, fn func(d *Di) (T, error)) {
	do.Provide(d.injector, func(_ *do.Injector) (T, error) {
		return fn(d)
	})
}

// Invoke injeta as dependências e chama a função que retorna o valor
//
// Exemplo:
//
// hc, err := pkgs.Invoke[handlers.HealthCheckHandler](di)
func Invoke[T any](d *Di) (T, error) {
	return do.Invoke[T](d.injector)
}
