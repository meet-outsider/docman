# Standard Go Project Layout
 ```
 ├── cmd/
│   ├── app/
│   │   ├── main.go
│   │   └── ...
│   └── ...
├── internal/
│   ├── pkg1/
│   │   ├── ...
│   │   └── ...
│   ├── pkg2/
│   │   ├── ...
│   │   └── ...
│   └── ...
├── pkg/
│   ├── lib1/
│   │   ├── ...
│   │   └── ...
│   ├── lib2/
│   │   ├── ...
│   │   └── ...
│   └── ...
├── vendor/
│   ├── ...
├── web/
│   ├── html/
│   │   ├── ...
│   │   └── ...
│   ├── css/
│   │   ├── ...
│   │   └── ...
│   ├── js/
│   │   ├── ...
│   │   └── ...
│   └── ...
├── configs/
│   ├── app.yaml
│   ├── log.yaml
│   └── ...
├── scripts/
│   ├── initdb.sql
│   ├── deploy.sh
│   └── ...
├── test/
│   ├── ...
├── docs/
│   ├── ...
├── LICENSE
├── README.md
└── go.mod

 ```
其中：

- cmd 目录包含应用程序的主要入口点，每个应用程序都应该有一个子目录。
- internal 目录包含应用程序的私有代码，这些代码不应该被其他应用程序使用。
- pkg 目录包含应用程序的公共库代码，可以被其他应用程序使用。
- vendor 目录包含应用程序的依赖库代码。
- web 目录包含应用程序的 Web 资源，例如 HTML、CSS 和 JavaScript。
- configs 目录包含应用程序的配置文件。
- scripts 目录包含用于构建、部署和管理应用程序的脚本。
- test 目录包含应用程序的测试代码。
- docs 目录包含应用程序的文档。
- LICENSE 文件包含应用程序的许可证。
- README.md 文件包含应用程序的说明文档。
- go.mod 文件包含应用程序的模块依赖关系。
