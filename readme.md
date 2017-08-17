# Server Push Example using GO

This simple app exemplifies the HTTP/2 Server Push technique.
See:
* [https://www.smashingmagazine.com/2017/04/guide-http2-server-push/](https://www.smashingmagazine.com/2017/04/guide-http2-server-push/)

The benefit of this technique is to improve the performance and caching of resources. With this we are able to inform the browser of additional resources required beforehand.

## Example
To start the application you have to have the GO Lang installed in your machine.
If you do have it, just run `go run Server.go`.

The Web Application is composed of 3 static files:
* index.html
* script.js
* styles.css

You can load the index.html through 2 different paths:
* [http://localhost:9000/preload](http://localhost:9000/preload) Where the `index.html` will be served, along with `script.js` and `styles.css` which are preloaded. Check the Response Headers, you should be able to see the `Link` header containing the preload metadata.
* [http://localhost:9000/postload](http://localhost:9000/postload) - Where the `index.html` will be served and both `script.js` and the `styles.css` will be required afterwards (normal way of requiring resources)




