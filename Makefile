wasm-compile:
	cd webassembly/scripts && ./compile.sh

wasm-serve:
	cd webassembly/scripts && ./run_browser.sh

run-browser: wasm-compile wasm-serve

run-node: 
	cd webassembly/scripts && ./run_node.sh

run-tests:
	cd webassembly/scripts && ./run_tests.sh
	
