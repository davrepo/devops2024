init:
	python -c"from minitwit import init_db; init_db()"

build:
	gcc flag_tool.c -lsqlite3 -o flag_tool

clean:
	rm flag_tool
