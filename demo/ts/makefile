.PHONY:

# tsc will build the ts source file and generate js file.
build:
	tsc --lib es2015,es5,dom --outDir target type.ts join.ts result.ts

install_types_node:
	sudo npm install -g @types/node

run:
	ts-node --files type.ts
	ts-node --files join.ts
	ts-node --files result.ts
