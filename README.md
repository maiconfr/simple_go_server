# Simple GO server

This project contains the code to create a simple server using GO language.
This project follows [this tutorial](https://blog.logrocket.com/creating-a-web-server-with-golang/) written by [Michiel Mulders](https://blog.logrocket.com/author/michielmulders/).

### Goals

This project:
- Can accept GET requests;
- Will respond to a POST request;
- Has basic securities.


### Structure

The project is divided in "server.go" and the "static" folder.
- server.go: All the GO code is developed in this file.
- static folder: All the static code are included at this folder, where we have the HTML code.

### Running

To run this project, download this project and run the following line at the prompt:
```
go run server.go
```

The project uses the por 8080, so, to see the project running, you need access 
```
http://localhost:8080
```
With /hello the server will access the "index.html" file, and with /form the server will access the "form.html" file.
The /form contains a form, where you can complete and send the informations by POST, and the server will return a message with those informations.
