package mandel_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/linuxfreak003/mandel"
	. "github.com/onsi/gomega"
)

func TestWrite(t *testing.T) {
	G := NewGomegaWithT(t)
	x, y := mandel.FindInterestingPoint(2.0, 2.0)
	m := mandel.NewGenerator(800, 800, x, y, 200.0)
	m.SetLimit(256)
	m.Generate()
	w := &bytes.Buffer{}
	err := m.Write(w)
	G.Expect(err).To(BeNil())
	err = ioutil.WriteFile("test.png", w.Bytes(), 0644)
	G.Expect(err).To(BeNil())
}
