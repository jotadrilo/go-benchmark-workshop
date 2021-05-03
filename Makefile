EXAMPLES=\
	0-concepts\
	1-concepts\
	2-sort\
	3-seen-map\
	4-rand

CAT := $(shell command -v bat 2>/dev/null || command -v cat)
DIFF := $(shell command -v diff -ur)

all: $(EXAMPLES)

clean:
	@rm -rf results

%: results/%/bench.txt
	@$(CAT) $^

results/%/bench.txt: examples/%/bench_test.go
	@mkdir -p $$(echo $@ | rev | cut -d/ -f2- | rev)
	@echo "go test -test.bench=. ./$? >$@"
	@go test -test.bench=. ./$? | grep Benchmark >$@

%-show: examples/%/bench_test.go
	@$(CAT) $^

%-stats: results/%/bench.txt
	@echo "benchstat -delta-test none $^ >results/$@"
	@benchstat -delta-test none $^ >results/$@
	@$(CAT) $^ results/$@

%-benchmark: results/%/bench.txt

.PRECIOUS: results/%/bench.txt
