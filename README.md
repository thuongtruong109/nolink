<div align="center">
    <h1><img src="public/logo.png" alt="logo"> nolink</h1>
    <img alt="Build status" src="https://img.shields.io/github/actions/workflow/status/thuongtruong1009/short1url/build.yml?logo=GitHub&label=build">
    <img alt="Test status" src="https://img.shields.io/github/actions/workflow/status/thuongtruong1009/short1url/test.yml?logo=GitHub&label=test">
    <img alt="CircleCI status" src="https://circleci.com/gh/circleci/circleci-docs.svg?style=svg">
    <a href="https://github.com/thuongtruong109/nolink/pkgs/container/nolink-api"><img alt="Automated api build" src="https://img.shields.io/docker/automated/thuongtruong1009/nolink-api?logo=Docker&label=server"></a>
    <a href="https://github.com/thuongtruong109/nolink/pkgs/container/nolink-client"><img alt="Automated client build" src="https://img.shields.io/docker/automated/thuongtruong1009/nolink-client?logo=Docker&label=client"></a>
    <a href="https://github.com/thuongtruong109/nolink/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/thuongtruong109/nolink"></a>
    <a href="https://paypal.me/thuongtruong1009" rel="nofollow"><img src="https://img.shields.io/badge/Donate-PayPal-ff3f59.svg" style="max-width: 100%;"></a>
</div>

## 🌳 Description

This is a simple URL shortener service, which helps you shorten your long URL to share repidly to external. It is written in Golang and uses Redis as database. Other hand, it also provides some services such as QR code generator, barcode generator, etc.

## 💫 Preview

![Preview image](public/preview.png)

## 💥 What's new

- [x] Shorten URL
- [x] Redirect to os.Getenv("ORIGIN")al URL
- [x] Expiration time (default 1 day)
- [x] QR code generator (custom color, download image, copy to clipboard)
- [x] Barcode generator
- [x] Statistics
- [x] Rate limit
- [x] Reverse proxy
- [x] Multi Dockerize layers
- [x] CI/CD build and deploy image
- [ ] Custom expiration time
- [ ] Unit test
- [ ] Caching

## 🏗️ Architecture

![Image](public/architecture.png)

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. If you like my work, please star 🌟 this repository.

## 🪪 License

**Nolink** is an [MIT-licensed](LICENSE) by <a href="https://github.com/thuongtruong109">thuongtruong109</a>

<!-- ## References

[Ref1](https://liamhieuvu.com/url-shortener-with-golang-and-mysql)
[Go on K8s](https://www.callicoder.com/deploy-multi-container-go-redis-app-kubernetes/)
[Nginx cache](https://vietnix.vn/cau-hinh-cache-nginx/)
[Nginx refs](https://github.dev/veryacademy/yt-nginx-mastery-series)
-->
