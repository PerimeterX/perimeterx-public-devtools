.PHONY: all images

images:
	docker build -t collect collect
	docker build -t build build

feed.mmdb: feed.csv
	docker run --rm -it -v ${PWD}:/out build /out/feed.csv /out/feed.mmdb

feed.csv:
	docker run --rm -it collect | tee feed.csv

run: feed.mmdb
	cp feed.mmdb use
	docker build -t use use
	docker run --rm -it -p 8080:8080 use
