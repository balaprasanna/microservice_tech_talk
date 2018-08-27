C_ID := $(shell docker ps -aq)

build:
	docker build -t newsapp-nodejs:latest .

run:
	docker run -d -p 5000:5000 -e PORT=5000 newsapp-nodejs:latest

stop:
	@docker stop $(C_ID)

remove:
	@docker rm $(C_ID)

clean: stop remove

test:
	@curl http://localhost:5000/
	@curl http://localhost:5000/api/v1/time
	@curl http://localhost:5000/api/v1/source

loadtest:
	@for i in `seq 1 $(n)`; do curl http://localhost:5000/; done;
	@for i in `seq 1 $(n)`; do curl http://localhost:5000/api/v1/time; done;
	@for i in `seq 1 $(n)`; do curl http://localhost:5000/api/v1/source; done;