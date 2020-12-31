.PHONY: bootstrap
bootstrap:
	mkdir -p ./bin
	curl -L -o ./bin/client_generator.tar.gz https://github.com/go-generalize/api_gen/releases/latest/download/client_generator_$(shell uname -s)_$(shell uname -m).tar.gz
	curl -L -o ./bin/server_generator.tar.gz https://github.com/go-generalize/api_gen/releases/latest/download/server_generator_$(shell uname -s)_$(shell uname -m).tar.gz
	cd ./bin && \
		tar xzf server_generator.tar.gz && \
		tar xzf client_generator.tar.gz && \
		rm *.tar.gz
