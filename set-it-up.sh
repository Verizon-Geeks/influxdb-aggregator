rm -rf build/_output
operator-sdk build quay.io/example/app-operator
kubectl apply -f deploy/crds/influxdb.monitoring.tbh.io_influxdbaggregators_crd.yaml
kubectl apply -f deploy/service_account.yaml
kubectl apply -f deploy/role.yaml
kubectl apply -f deploy/role_binding.yaml
kubectl apply -f deploy/operator.yaml
kubectl apply -f deploy/crds/influxdb.monitoring.tbh.io_v1alpha1_influxdbaggregator_cr.yaml
