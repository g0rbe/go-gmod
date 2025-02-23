//go:build linux

package freax_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func TestContainsLine(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	if _, err := file.WriteString("a\nb\nc\nd\n"); err != nil {
		t.Fatalf("Failed to write to testfile: %s\n", err)
	}

	exist, err := freax.FileContainsLine("testfile", "d")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	if !exist {
		t.Fatalf("\"d\" should exist in testfile\n")
	}
}

func TestCountLines(t *testing.T) {

	randLen := rand.Intn(10240)

	file, err := os.OpenFile("testfile", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Failed to create test file: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	for i := 0; i < randLen; i++ {
		file.Write([]byte{'a', '\n'})
	}

	n, err := freax.FileCountLines("testfile")
	if err != nil {
		t.Fatalf("Failed to count: %s\n", err)
	}

	if n != randLen {
		t.Fatalf("Invalid result: want: %d, got: %d\n", randLen, n)
	}
}

func BenchmarkContainsLine(b *testing.B) {

	file, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to create testfile: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	if _, err := file.WriteString("a\nb\nc\nd\n"); err != nil {
		b.Fatalf("Failed to write to testfile: %s\n", err)
	}

	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.FileContainsLine("testfile", "d")
	}

}

func TestIsExists(t *testing.T) {

	testName := "testfile.txt"

	file, err := os.OpenFile(testName, os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		t.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Failed to close %s: %s\n", testName, err)
	}

	ok, err := freax.IsPathExists(testName)
	if err != nil {
		t.Fatalf("FAIL: failed to check if exist: %s\n", err)
	}

	if !ok {
		t.Fatalf("FAIL: %s should exists!\n", testName)
	}

	err = os.Remove(testName)
	if err != nil {
		t.Fatalf("Failed to remove %s: %s\n", testName, err)
	}
}

func BenchmarkIsExistsTrue(b *testing.B) {

	testName := "testfile.txt"

	file, err := os.OpenFile(testName, os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		b.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	err = file.Close()
	if err != nil {
		b.Fatalf("Failed to close %s: %s\n", testName, err)
	}

	clean := func() {
		err = os.Remove(testName)
		if err != nil {
			b.Fatalf("Failed to remove %s: %s\n", testName, err)
		}
	}
	b.Cleanup(clean)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.IsPathExists(testName)
	}

}

func BenchmarkIsExistsFalse(b *testing.B) {

	testName := "testfile.txt"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.IsPathExists(testName)
	}

}
