# Welcome to gofs 👋
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/floholz/gofs?label=latest)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/floholz/gofs/go.yml)
![GitHub License](https://img.shields.io/github/license/floholz/gofs)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/floholz/gofs?logo=go&labelColor=gray&label=%20)

---

> This is a simple fileserver tool to expose a file or directory locally. 
> It also has the option to rebind the url to a different path when exposing. 
> It is written in golang.


## Install

```sh
go get github.com/floholz/gofs
```

## Usage

To expose the current directory, simply run the package. By default, the directory will be exposed at `http://localhost:8080/`

```sh
go run github.com/floholz/gofs
```

Instead of the active directory, you can choose a file or directory to expose by passing it as an argument.  

```sh
go run github.com/floholz/gofs ~/my-direcory/my-file.txt
```

To set the url, your file or directory should be exposed to, use the `--url` or `-u` flag.

```sh
go run github.com/floholz/gofs ~/my-direcory/my-file.txt --url 0.0.0.0:3003/path-to-file/file.json
```

By leaving the `hostname`, `port` or `path` empty, it will be set to its default value.

| **URL part**       | **default value**        |
|--------------------|--------------------------|
| scheme             | _is always set to_ http: |
| hostname           | localhost                |
| port               | 8080                     |
| path _(directory)_ | /                        |
| path _(file)_      | /_filename.ext_          |

### Examples for `--url` parsing

| **argument**      | **`--url` value**              | **result**                              |
|-------------------|--------------------------------|-----------------------------------------|
| ~/my-dir/         |                                | http://localhost:8080/                  |
| ~/my-dir/file.txt |                                | http://localhost:8080/file.txt          |
| ~/my-dir/         | 192.168.0.10                   | http://192.168.0.10:8080/               |
| ~/my-dir/         | 192.168.0.10:1234              | http://192.168.0.10:1234/               |
| ~/my-dir/         | 192.168.0.10:1234/custom/path/ | http://192.168.0.10:1234/custom/path/   |
| ~/my-dir/file.txt | :1234                          | http://localhost:1234/file.txt          |
| ~/my-dir/file.txt | :1234/path/to-file.json        | http://localhost:1234/path/to-file.json |
| ~/my-dir/         | /custom/path/                  | http://localhost:8080/custom/path/      |


---

### 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/floholz/gofs/issues). 


### 📝 License

Copyright © 2024 [floholz](https://github.com/floholz).

This project is [MIT](./LICENSE) licensed.

---

### Show your support

Give a ⭐ if this project helped you!