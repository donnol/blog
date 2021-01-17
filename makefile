.PHONY: server

server:
	hugo server --baseURL=https://blog.jdscript.com --bind=0.0.0.0 --appendPort=false
