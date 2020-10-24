Go y WebAssembly: interactuando con la API JavaScript de tu Navegador Web
=====================================================

[WebAssembly (abreviado WASM)](https://webassembly.org/) promete llevar a un
siguiente nivel la programación sobre navegadores web. Hasta hace poco, los
navegadores solo eran capaces de ejecutar programas escritos en JavaScript/ECMAScript
(en adelante, nos referiremos a este lenguaje como JS).
Si se deseaba programar para el navegador en otros lenguajes, o en versiones muy
modernas de JS, era necesario un transpilador que convertía el código
a una versión de JavaScript aceptada por el navegador.

JavaScript se interpreta en tiempo de ejecución. Interpretar cadenas de texto
con el programa JS es un proceso relativamente lento, que causa que
las web se vuelvan realmente lentas a medida que se vuelven más sofisticadas.

WebAssembly, sin embargo, es una especificación de instrucciones sencillas, que
se especifican mediante código binario (como si de las instrucciones de una
CPU física se tratara). Esto no solo disminuye enormemente el tiempo de carga
(los programas a cargar son mucho más pequeños), sino que el tiempo de interpretación
y optimización del programa en tiempo de ejecución también se reduce.

Otra ventaja es que ahora es posible programar en prácticamente cualquier lenguaje
las aplicaciones que se ejecutarán directamente sobre el navegador web; incluso
en lenguajes que hasta ahora se consideraban más orientados a sistemas de bajo
nivel, como C, C++ o Rust. Esto permite llevar al navegador un gran número de
aplicaciones antiguas escritas en estos lenguajes (videojuegos, aplicaciones
multimedia...).

Sin embargo, la mayoría de bibliotecas y APIs del Navegador, imprescindibles
para interactuar con éste, siguen proporcionando interfaces JavaScript únicamente.
Este capítulo explica cómo utilizar la librería estándar para la interconexión
de Go con las APIs de JS, permitiendo instanciar objetos, invocar métodos y
funciones, acceder a sus propiedades o incluso proporcionar funciones que se
usarán como [_callback_](https://es.wikipedia.org/wiki/Callback_(inform%C3%A1tica)).

Go provee el paquete experimental `syscall/js`, que facilita la creación de
aplicaciones para el navegador web sin necesidad de ningún transpilador JS;
simplemente especificando al compilador de Go que la arquitectura de destino
será WASM. Este capítulo explicará cómo compilar de Go a WASM y cómo configurar
una aplicación web para cargar nuestros programas WASM.

## Preparando el entorno

Ejecutar un proyecto de WebAssembly en Go requiere tres archivos:

* Un archivo con extensión `.wasm`, que contiene los datos binarios de WebAssembly.
  Este archivo será generado por el compilador de Go.
* Un fichero `wasm_exec.js`, proporcionado por el equipo de Go para poder cargar
  el archivo `.wasm` dentro de la página web.
* Un fichero HTML que carga el archivo `wasm_exec.js` y lo configura para 
  cargar y ejecutar el código `.wasm` en el navegador.

Además, será necesario un servidor Web para cargar los anteriores archivos mediante
HTTP o HTTPS. Se puede utilizar cualquier servidor HTTP simple, aunque en este
capítulo crearemos nuestro servidor HTTP para probar el código de los tutoriales
sin tener que instalar ni configurar software adicional en nuestro ordenador.

### 1. Creación del proyecto

Nuestro proyecto Go/WASM tendrá la siguiente estructura de directorios y
archivos:

```
.
├── server.go
├── site
│   ├── index.html
│   ├── main.wasm
│   └── wasm_exec.js
└── src
    └── main.go
```

* El directorio `site` contendrá los archivos mínimos indispensables para ejecutar
  un programa WASM en el navegador, tal y como se mencionó anteriormente.
* El programa `server.go` es un sencillo navegador Web de pruebas que servirá
  para cargar desde el navegador Web los archivos de la carpeta `site`.
* El directorio `src` contiene el código Go que se compilará en el archivo
  `site/main.wasm`.


### Download `wasm_exec.js`

The `wasm_exec.js` file is available in your standard Go installation. Just copy
it into the `site` folder with the following command:

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./site/
```

### Create your HTML file

For example, let's create an `index.html` file under the `site` folder:

```html
<html>
<head>
  <meta charset="utf-8"/>
  <script src="wasm_exec.js"></script>
  <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
        .then((result) => {
          go.run(result.instance);
        });
    </script>
</head>
<body>
</body>
</html>
```

The file is empty. It just loads the `wasm_exec.js` and fetches the `main.wasm`
file that is created in the following section.

### Compile your code into `main.wasm`

Let's create a dummy Go program in the `./main-wasm/main.go` path of
your project:

```go
package main

import "log"

func main() {
    log.Println("Hello Gophers!")
}
```

And let's compile it into the `./site/main.wasm` binary file:

```
GOOS=js GOARCH=wasm go build -o ./site/main.wasm ./wasm-main/.
```

Please observe that you need to set the `GOOS` and `GOARCH` environment
variables to `js` and `wasm`, respectively.

### Execute your Go WebAssembly program

You will need a Web server to allow fetching all the information. Modern IDEs
like IntelliJ IDEA bring their own bundled server, so you can preview your local
files as if they were in a remote server.

If you are using a plain text editor and don't want to install any web server,
the following `server.go` file in your project root will do the job:

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)
// super-simple debug server to test our Go WASM files
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        if req.RequestURI == "/" {
            req.RequestURI = "/index.html"
        }
        file, err := os.Open(filepath.Join("./site", req.RequestURI))
        if err == nil {
            io.Copy(w, file)
        }
    })
    fmt.Println(http.ListenAndServe(":8080", nil))
}
```

When you run the server (e.g. `go run server.go`) and go to
`http://localhost:8080` in your browser, you should see an empty screen. But if
you open the _Console_ in the _developer tools_ section of your browser,
you should see that the `log` command in the `./main-wasm/main.go` has been
executed:

![](/static/assets/2020/03/go_wasm/log_console.png)

To allow Go interacting with a Web Page and read/write contents in the actual
HTML document, let's see some methods of the `syscall/js` library.

## `syscall/js` basic functionalities

Let's walk through the basic functionalities of `syscall/js` with a simple
example:

```go
 1: func main() {
 2: 	window := js.Global()
 3: 	doc := window.Get("document")
 4:	body := doc.Get("body")
 5:	div := doc.Call("createElement", "div")
 6:	div.Set("textContent", "hello!!")
 7:	body.Call("appendChild", div)
 8:	body.Set("onclick",
 9:		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
10:			div := doc.Call("createElement", "div")
11:			div.Set("textContent", "click!!")
12:			body.Call("appendChild", div)
13:			return nil
14:		}))
15:	<-make(chan struct{})
16: }
```

This code adds to the HTML document a `<div>` element containing the `Hello!!`
message. In addition, the program is subscribed to the `onclick` event of the
document, and each time the user clicks the document, a new `<div>` is added,
containing the `click!!` text.

![](/static/assets/2020/03/go_wasm/result.png)

The functions used in this example are:

`js.Global()` in Line 1 returns the global object, usually the equivalent
to the JavaScript `window` object: the _root_ object that will allow you accessing
all the other elements in your page.

`js.Global()` returns a `js.Value`: a struct that can store any JavaScript type.
You will get used to work with `js.Value`, as it's what most functions and
properties return.

The `Get` method invoked on a `js.Value` returns another `js.Value`
belonging to the property passed as argument. For example, the `Get` invocations
in lines 3, and 4. 

The opposite of `Get` is the `Set` function, which receives two arguments: the
name of a property, and its new value. The value doesn't
need to be a `js.Value` instance: you can pass numbers or strings, like in
lines 6 and 11, and even instances of `js.Func` (lines 8-9), that specify
a function to be assigned to this property. In the example of lines 8-9,
a given Go function is assigned to the `onclick` event.

Finally, the example code also uses the `Call` method of `js.Value` to invoke
methods of a given type. `Call` requires the name of a function as the first
argument, following by a variable number of arguments. Examples of `Call` are
seen in lines 5, 7, 10 an 12. As for `Set`, the arguments can be native Go types
or other `js.Value` or `js.Func`.

The simplest way to instantiate a `js.Func` is by means of the `js.FuncOf` Go
auxiliary function.

## To know more

This introduction tutorial does not cover many other functionalities, as
instantiating JavaScript objects. For more details, please check the
[`syscall/js` package documentation](https://golang.org/pkg/syscall/js/).

The examples of this blog post are available in
[my Github repo](https://github.com/mariomac/go-wasm-tutorial).



 

