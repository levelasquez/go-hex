redis:
	docker run --rm -it \
	-p 6379:6379 \
	--name go-hex-redis \
	redis:5

mongo:
	docker run --rm -it \
	-p 27017:27017 \
	--name go-hex-mongo \
	mongo:4
