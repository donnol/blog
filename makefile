.PHONY: server

server:
	hugo server --baseURL=https://blog.jdscript.com --bind=0.0.0.0 --appendPort=false

# curl https://wasmtime.dev/install.sh -sSf | bash
# rustup target add wasm32-wasi
rustc_wasm_compile_run:
	cd demo/rust/mybin/src && \
	rustc main.rs --target wasm32-wasi && \
	wasmtime main.wasm
