# study-helm
KubernetesのHelmの学習リポジトリ

## helmリポジトリ

```sh
helm repo add study-helm https://shota-takaoka.github.io/study-helm/
```

## Usage

### インストール

```sh
./install
```

### アンインストール

```sh
./uninstall
```

### リクエスト

```sh
curl hello.info:30080/hello
```

## Helmコマンド

### Helmチャートの作成

```sh
helm create charts/[作成するチャート名]
```


