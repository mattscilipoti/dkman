# Derived from: https://github.com/chris-crone/containerized-go-dev/blob/main/Makefile
all: bin/dkman
test: lint unit-test

PLATFORM=local

.PHONY: bin/dkman
bin/dkman:
	@docker build . \
	--tag dkman \
	--target bin \
	--output bin/ \
	--platform ${PLATFORM}

# .PHONY: unit-test
# unit-test:
# 	@docker build . --target unit-test

# .PHONY: unit-test-coverage
# unit-test-coverage:
# 	@docker build . --target unit-test-coverage \
# 	--output coverage/
# 	cat coverage/cover.out

# .PHONY: lint
# lint:
# 	@docker build . --target lint