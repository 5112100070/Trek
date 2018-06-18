all: test build

run: build pre-deploy start

build:
	@echo " >> building binaries"
	@go build -o bin/trek src/cmd/app.go

start:
	@echo " >> starting binaries"
	@./bin/trek

pre-deploy:
	sudo mv bin/trek /usr/local/bin/.
	sudo cp -r files/WEB-INF/pages/. /var/www/trek/pages/.
	sudo cp -r files/WEB-INF/attr/scss/. /var/www/trek/scss/.
	sudo cp -r files/WEB-INF/attr/css/. /var/www/trek/css/.
	sudo cp -r files/WEB-INF/attr/js/. /var/www/trek/js/.
	sudo cp -r files/WEB-INF/attr/img/. /var/www/trek/img/.
	sudo cp -r files/WEB-INF/attr/vendor/. /var/www/trek/vendor/.
	sudo cp -r files/WEB-INF/attr/etc/. /var/www/trek/etc/.
	sudo cp -r files/WEB-INF/attr/files/. /var/www/trek/files/.
	