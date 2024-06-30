# 参考資料：https://qiita.com/prodigy413/items/

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

kubectl create ns ingress-system
helm install obi-test ingress-nginx/ingress-nginx -n ingress-system