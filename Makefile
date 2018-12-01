source: underscore.go

underscore.js:
	curl -kL http://underscorejs.org/underscore.js > $@
	
# requires https://github.com/jteeuwen/go-bindata
underscore.go: underscore.js
	go-bindata -o="./underscore.go" -pkg="underscore" underscore.js