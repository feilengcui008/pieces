#!/bin/bash
protoc  --grpc_out=. --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin dictionary.proto
protoc --cpp_out=. dictionary.proto
g++ -std=c++11 -o server server.cc dictionary.grpc.pb.cc dictionary.pb.cc -lprotobuf -lgrpc -lgrpc++ -pthread
g++ -std=c++11 -o client client.cc dictionary.grpc.pb.cc dictionary.pb.cc -lprotobuf -lgrpc -lgrpc++ -pthread
