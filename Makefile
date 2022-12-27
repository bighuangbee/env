clean:
	docker images -q -f dangling=true | xargs docker rmi -f
	docker volume rm $(docker volume ls -qf dangling=true)
