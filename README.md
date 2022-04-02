# http-server-test
This project contains few basic HTTP servers. The goals of it is to test how HTTP server behaves when runs on 
different platforms:
* On a bare metal (as a container run via podman)
* On a virtual machine (as a container run via podman or on OCP cluster as a pod)

# Build instructions

Build the container:
```bash
podman build -t http-server-test:latest . -f Dockerfile
```

Run the container:
```bash
podman run -it --rm -p 9090:9090 http-server-test:latest
```

Verify the container:
```bash
curl http://localhost:9090
```

Deploy on OCP (Update image in deployment):
```bash
```bash
kubectl apply -f http-server-deploy.yaml
kubectl wait --for=condition=Available deploy/http-server-test -n default --timeout=120s
```

Use jmeter to test the service:
```bash
JMETER_HOME=/home/test/apache-jmeter-5.4.1
TEST_PLAN=/tmp/jmeter-http-server-test-plan.jmx
test_dir=$(mktemp -d)
test_route=$(kubectl get routes http-server-test --no-headers -o=custom-columns=HOST:.spec.host)
mkdir -p $test_dir/results
JVM_ARGS="-Xms4g -Xmx64g -Xss250k -XX:MaxMetaspaceSize=1g" $JMETER_HOME/bin/jmeter.sh -n -l $test_dir/results.csv \
    -f -e -o $test_dir/results/ -t $TEST_PLAN \
    -JUSERS=10 \
    -JRAMP_UP_TIME=100 \
    -JITERATIONS=2 \
    -JHTTP_SERVER=$test_route |& tee -a $test_dir/summary.txt
```
-----
Compare with Nginx:
```bash
JMETER_HOME=/home/test/apache-jmeter-5.4.1
TEST_PLAN=/tmp/jmeter-http-server-test-plan.jmx
test_dir=$(mktemp -d)
kubectl apply -f nginx-deploy.yaml
kubectl wait --for=condition=Available deploy/nginx -n default --timeout=120s
test_route=$(kubectl get routes nginx --no-headers -o=custom-columns=HOST:.spec.host)
JVM_ARGS="-Xms4g -Xmx64g -Xss250k -XX:MaxMetaspaceSize=1g" $JMETER_HOME/bin/jmeter.sh\
    -n -l $test_dir/results.csv \
    -f -e -o $test_dir/results/ \
    -t $TEST_PLAN \
    -JUSERS=20000 \
    -JRAMP_UP_TIME=2000 \
    -JITERATIONS=100 \
    -JHTTP_SERVER=$test_route |& tee -a $test_dir/summary.txt
```