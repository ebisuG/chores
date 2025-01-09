### Memo
Index.html is sent by server.js(*) on the toppest directory.\
Index.html contains div with "root" id, and entry-client.js script.\
Inside entry-server.js,  renderToString is used. This return html string.\
As documents say, renderToString doesn't wait data fetching. It looks renderToPipeableStream is necessary.\
Maybe as long as using renderToString, sending basic html first and js files second is what it can do.
