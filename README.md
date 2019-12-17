# gohazelcastqueue
golang hazelcast queue client

# build
make install clean build

# hazelcast server
docker run ---name hazelcast d -e JAVA_OPTS="-Dhazelcast.local.publicAddress=192.168.178.120:5701" -e MIN_HEAP_SIZE="512M" -e JAVA_OPTS="-Xms512M -Xmx1024M" -p 5701:5701 hazelcast/hazelcast:3.12.5

# run
./tasksource/tasksource

./tasksink/tasksink
