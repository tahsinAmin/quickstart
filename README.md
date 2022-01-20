# What I've learned:

- ```
    type MainController struct {
    web.Controller
    }
  ```

  The MainController is the first thing created. It contains an anonymous struct field of type web.Controller. This is called struct embedding and is the way that Go mimics inheritance. Because of this MainController automatically acquires all the methods of web.Controller.

- Beego registers the static directory as the static path by default. Registered rule: URL prefix with directory mapping

`StaticDir["/static"] = "static"`

You can register multiple static directories. For example two different download directories, download1 and download2, can be set using:

```
web.SetStaticPath("/down1", "download1")
web.SetStaticPath("/down2", "download2")
```

Visiting the URL http://localhost/down1/123.txt will request the file 123.txt in the download1 directory.
To remove the default /static -> static mapping use web.DelStaticPath("/static")

# Problems and it solution

- 2022/01/20 15:33:00 ERROR ▶ 0004 Failed to build the application: can't load package: package .: no Go files in /home/w3e63/Desktop/Beego. Solution: go inside the directory and then press

````

bee run

```

```
````
