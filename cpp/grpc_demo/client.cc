#include <string>
#include <unordered_map>
#include <memory>
#include <iostream>
#include <stdio.h>
#include <boost/lexical_cast.hpp>

#include <grpc/grpc.h>
#include <grpc++/client_context.h>
#include <grpc++/create_channel.h>
#include <grpc++/security/server_credentials.h>

#include "dictionary.grpc.pb.h"

using Tan::Word;
using Tan::Meaning;
using Tan::Dictionary;

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReader;
using grpc::ClientReaderWriter;
using grpc::ClientWriter;
using grpc::Status;



class Client {
  public:
    Client(std::shared_ptr<Channel> channel) : stub_(Dictionary::NewStub(channel)) {} 
    bool callGetOne(const Word &request, Meaning *response)
    {
      ClientContext context;
      Status status = stub_->GetOne(&context, request, response); 
      if (!status.ok()) {
        std::cout << "err when call the rpc" << std::endl; 
      }
      std::cout << "the meaning of word : " << request.data() << " is : " 
        << response->data() << std::endl;
      return status.ok();
    }
  private:
    std::unique_ptr<Dictionary::Stub> stub_;
};


int main(int argc, char *argv[])
{
  Client client(grpc::CreateChannel("localhost:8080", grpc::InsecureChannelCredentials())); 
  std::cout << "======== client issure requests to localhost:8080 ========" << std::endl;
  Word request;
  Meaning response;

  for (int i = 1; i <= 5000; ++i) {
    request.set_data(std::string("key") + boost::lexical_cast<std::string>(i));
    client.callGetOne(request, &response);
  }

  return 0;
}
