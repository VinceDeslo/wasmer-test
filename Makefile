wasm-compile:
	@cd webassembly/scripts && ./compile.sh

wasmer-serve:
	@cd webassembly/scripts && ./run_wasmer.sh

run-wasmer: wasm-compile wasmer-serve

run-node: 
	@cd webassembly/scripts && ./run_node.sh

run-tests:
	@cd webassembly/scripts && ./run_tests.sh
