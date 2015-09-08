# Copyright 2015 The CC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all clean editor later nuke todo cover cpu edit internalError mem

grep=--include=*.l --include=*.y --include=*xerror* --include=*.go

all: editor
	go vet || true
	golint
	make todo

clean:
	go clean
	rm -f *~ cpu.test mem.test *.out

cover:
	t=$(shell tempfile) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu:
	go test -c -o cpu.test
	./cpu.test -noerr -test.cpuprofile cpu.out
	go tool pprof --lines cpu.test cpu.out

edit:
	gvim -p Makefile trigraphs.l scanner.l xerrors parser.y y.output *.go

editor:
	go generate
	gofmt -l -s -w *.go
	go test
	go install

internalError:
	egrep -ho '"internal error.*"' *.go | sort | cat -n

later:
	@grep -n $(grep) LATER * || true
	@grep -n $(grep) MAYBE * || true

mem:
	go test -c -o mem.test
	./mem.test -test.run NONE -test.bench BenchmarkPPParseCorpus -test.memprofile mem.out
	go tool pprof --lines --web --alloc_space mem.test mem.out

nuke: clean
	go clean -i
	rm -f *_string.go parser.go scanner.go trigraphs.go ast.go ast_test.go

todo:
	@grep -n $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * || true
	@grep -n $(grep) TODO * || true
	@grep -n $(grep) BUG * || true
	@grep -n $(grep) [^[:alpha:]]println * || true
