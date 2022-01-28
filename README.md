# Tasks

- [x] How to show breeds in select
- [x] Show something for the categories to appear
- [x] Show images by default
- [x] Able to manipulate all data and fetch according to that
- [x] Images Per Page
- [x] Ascending should reorder the current ones
- [x] Show some message if nothing found
- [x] Modify Grid
- [x] Adding More, to show new images
- [x] No image found
- [x] Have the refresh button move along the scrollbar
- [x] Make it rersponsive
- [x] Lazy Loading
- [x] Frontend
- [ ] How to add api key in beego

# Refactor

- [ ] How to do it with View
- [ ] Change var to let in index.html
- [ ] Hover show link

# API

- To get cat images by categories, give the category id and it'll give you:
  `https://api.thecatapi.com/v1/images/search?category_ids=1`

- To get all categories
  `https://api.thecatapi.com/v1/categories`

- To get all categories
  `https://api.thecatapi.com/v1/breeds`

- Fetch by breeds:
  `https://api.thecatapi.com/v1/images/search?breed_ids=abob`

    <!-- https://documenter.getpostman.com/view/4016432/RWToRJCq#ea81771b-b042-42d1-ab7f-c75deb6bb259 -->

    <!-- https://documenter.getpostman.com/view/4016432/RWToRJCq -->

# What I've learned:

- Golang doesn't support any ternary operation.

- to make multidimensional array with its length coming from variables,. those variuables has tyo be constants:

```
Code:

const first = 2
		const sec = 3
		var twoD [first][sec]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)
```

```
Output:

2d:  [[0 1 2] [1 2 3]]
```

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

- How to comment in HTML:

```
<!-- ... -->
```

- [About Json](https://www.sohamkamani.com/golang/json/)
- [Using ajax](https://stackoverflow.com/questions/50128474/how-ajax-will-get-the-data-from-the-golang-code)

# Problems and it solution

- 2022/01/20 15:33:00 ERROR ▶ 0004 Failed to build the application: can't load package: package .: no Go files in /home/w3e63/Desktop/Beego. Solution: go inside the directory and then press

```
bee run
```

- [How to console.log select tag's value](https://ricardometring.com/getting-the-value-of-a-select-in-javascript)

- If any fields value is not showing from the fetched json response. Check if the field typy matches with the json property's value's type.

- How to pass golang values inside a property in tag -> `<div class="... bg-[url({{ .Url }})]"></div>`

- [Embedding in Go](https://www.youtube.com/watch?v=xba1mcOsWEI)

<!-- https://docs.thecatapi.com/ -->
<!-- https://docs.thecatapi.com/authentication -->

<!-- https://jsfiddle.net/adenF/njf4vts0/ -->t
<!-- https://www.youtube.com/watch?v=aYk8XAKxhxU&list=PLujhHB_uAFJws6Vv5q1KDoaQ4YcpS9UOm&index=5 -->
<!-- https://www.youtube.com/watch?v=0ub6BwdBwIY -->
<!-- https://www.youtube.com/results?search_query=golang+ajax -->
<!-- https://www.youtube.com/watch?v=F3tieL1lX1I -->
<!-- https://uploads-ssl.webflow.com/5e3de80322b300854230f11f/5e5bebb60e1706155d830222_bidroom-search-page-1280x960.jpeg -->
<!-- https://www.google.com/imgres?imgurl=http%3A%2F%2Fassets.uxbooth.com%2Fuploads%2F2017%2F05%2Fimage3.png&imgrefurl=https%3A%2F%2Fwww.uxbooth.com%2Farticles%2Fbest-practices-for-search%2F&tbnid=CILtm-2OFNMdRM&vet=12ahUKEwi40o_khML1AhUK-DgGHVICAt4QMygGegUIARDTAQ..i&docid=6CrS57Z9cO39XM&w=1999&h=730&itg=1&q=search%20bars%20web%20design&client=ubuntu&ved=2ahUKEwi40o_khML1AhUK-DgGHVICAt4QMygGegUIARDTAQ -->

<!-- https://www.youtube.com/watch?v=RQtgMCajbX8 -->

<!-- -->

<!-- https://www.youtube.com/watch?v=0lyngYIozro -->

<!-- https://cdn.dribbble.com/users/1498143/screenshots/9539471/media/a9eb6ed8dd9114790f7bec5d8f9b0843.png?compress=1&resize=400x300 -->



Hello