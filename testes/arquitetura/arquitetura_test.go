package arquitetura

import (
	"runtime"
	"testing"
)

func TestDpendende(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Parallel() //permite executar o teste em paralelo
		t.Skip("Não funciona em arquitetura amd64")
	}

	//...

	t.Fail()
}
