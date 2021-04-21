EXAMPLES=\
	0-concepts\
	1-concepts\
	2-sort\
	3-map

CAT := $(shell command -v bat 2>/dev/null || command -v cat)
DIFF := $(shell command -v diff -ur)

all: $(EXAMPLES)

clean:
	@rm -rf results

%: results/%-1/bench.txt results/%-2/bench.txt
	@$(CAT) results/$@-1/bench.txt results/$@-2/bench.txt

results/%/bench.txt: examples/%/bench_test.go
	@mkdir -p $$(echo $@ | rev | cut -d/ -f2- | rev)
	@echo "go test -test.bench=. ./$? >$@"
	@go test -test.bench=. ./$? | grep Benchmark >$@

%-show: examples/%-1/bench_test.go examples/%-2/bench_test.go
	@$(CAT) $^
	@$(DIFF) $^ | $(CAT)

%-stats: results/%-1/bench.txt results/%-2/bench.txt
	@echo "benchstat -delta-test none $^ >results/$@"
	@benchstat -delta-test none $^ >results/$@
	@$(CAT) $^ results/$@

.PRECIOUS: results/%/bench.txt
