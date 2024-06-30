# study-helm
KubernetesのHelmの学習リポジトリ

## helmリポジトリ

以下の手順でお試しデプロイ可能

```sh
# helmリポジトリの追加
helm repo add study-helm https://shota-takaoka.github.io/study-helm/

# helmリポジトリ内のチャート内の一覧
helm search repo study-helm

# 追加したhelmリポジトリからチャートをインストールする
helm install test-app1 study-helm/test-app1

# ワーカーノードのINTERNAL-IPを確認
kubectl get node -o wide

# サーバーが起動しているかを確認
curl [ワーカーノードのINTERNAL-IP]:31080

# チャートのアンインストール
helm uninstall test-app1 
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


