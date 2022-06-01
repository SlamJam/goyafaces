package providers

import (
	"bytes"

	"github.com/SlamJam/goyafaces/pkg/app/config"
	"github.com/corpix/revip"
)

func ProvideConfig() (*config.Config, error) {
	c := config.Config{
		Foo: &config.Foo{
			Bar: "bar",
			Qux: true,
		},
		Baz: 666,
	}

	_, err := revip.Load(
		&c,
		revip.FromReader(
			bytes.NewBuffer([]byte(`{"foo":{"qux": false}}`)),
			revip.JsonUnmarshaler,
		),
		revip.FromReader(
			bytes.NewBuffer([]byte(`box = [666,777,888]`)),
			revip.TomlUnmarshaler,
		),
		revip.FromEnviron("revip"),
	)
	if err != nil {
		return nil, err
	}

	err = revip.Postprocess(
		&c,
		revip.WithDefaults(),
		revip.WithValidation(),
		revip.WithExpansion(),
	)

	return &c, err
}
