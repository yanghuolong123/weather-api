all:
	make tar
tar: 
	../../bin/bee pack
	docker build -t weather-1.0 -f ./Dockerfile .
	docker save -o weather-1.0.tar weather-1.0
	tar -zcvf weather-1.0.tar.gz weather-1.0.tar
clean:
	rm app.tar.gz
