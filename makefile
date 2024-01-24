.PHONY: server

server:
	hugo server --baseURL=https://blog.jdscript.com --bind=0.0.0.0 --appendPort=false

# curl https://wasmtime.dev/install.sh -sSf | bash
# rustup target add wasm32-wasi
rustc_wasm_compile_run:
	cd demo/rust/mybin/src && \
	rustc main.rs --target wasm32-wasi && \
	wasmtime main.wasm

rust_bindgen:
	bindgen demo/c/shift.h -o demo/rust/mybin/src/bindings.rs

build_object:
	clang src/shift.c -c

rust_build_with_link:
	RUSTFLAGS="-l static=/home/jd/Project/jd/blog/demo/rust/mybin/shift.o" cargo build 
