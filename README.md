### nzen-iot-client-test

KubeEdge환경에서 IOT 데이터 전송 테스트 프로그램

가속도센서 데이터를 초단위로 broker로 전송한다. 

```shell
sudo groupadd docker
sudo usermod -aG docker $USER

sudo docker login 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry
Username: 
Password: 

sudo docker build -t 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/nzen-iot-client-test .
sudo docker push 44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/nzen-iot-client-test

kubectl create secret docker-registry ncr-secret \
  --docker-server=44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry \
  --docker-username= \
  --docker-password=
```