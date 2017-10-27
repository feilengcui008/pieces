#include <string>
#include <unordered_map>
#include <memory>
#include <iostream>
#include <stdio.h>

#include <grpc/grpc.h>
#include <grpc++/server.h>
#include <grpc++/server_builder.h>
#include <grpc++/server_context.h>
#include <grpc++/security/server_credentials.h>

#include "dictionary.grpc.pb.h"

using Tan::Word;
using Tan::Meaning;
using Tan::Dictionary;

using grpc::Status;
using grpc::Server;
using grpc::ServerContext;
using grpc::ServerWriter;
using grpc::ServerReader;
using grpc::ServerReaderWriter;
using grpc::ServerBuilder;


class DictionaryImplSync final : public Dictionary::Service {
  public:
    DictionaryImplSync(const std::string &filepath) : filepath_(filepath) 
  {
    fp_ = fopen(filepath_.c_str(), "r");
  }
    ~DictionaryImplSync() 
    {
      fclose(fp_);
    }
    Status GetOne(ServerContext *context, const Word *request, Meaning *response)
    {
      std::cout << "~~~~~ get client request for word : " << request->data() << std::endl;
      response->set_data(std::string(""));
      if (buffer_.find(request->data()) != buffer_.end()) {
        response->set_data(buffer_[request->data()]);
      } else {
        //fseek(fp_, 0, SEEK_SET);
        char *line = NULL;
        size_t size = 0;
        std::string key;
        std::string value;
        while (getline(&line, &size, fp_) != -1) {
          // get key 
          while (*line != ' ' && *line != '\t' && *line != '\n' && *line != '\0') {
            key.push_back(*line);
            line++;
          }
          // skip white space
          if (key == request->data()) {
            while (!(isalnum(*line))) {
              line++;    
            }
            // get value
            while (*line != ' ' && *line != '\t' && *line != '\n' && *line != '\0') {
              value.push_back(*line);
              line++;
            }
            response->set_data(value);
            buffer_[request->data()] = value;
            break;
          }

        }
      }
      return Status::OK;
    }

    Status SendStreamList(ServerContext *context, const Word *request, ServerWriter<Meaning> *response)
    {
      return Status::OK;
    }
    
    Status GetStreamList(ServerContext *context, ServerReader<Word> *request, Meaning *response)
    {
      return Status::OK;
    }

    Status StreamList(ServerContext *context, ServerReaderWriter<Meaning, Word> *rw)
    {
      return Status::OK;
    }

  private:
    const std::string filepath_;
    std::unordered_map<std::string, std::string> buffer_;
    FILE *fp_;

};


class DictionaryImplAsync final : public Dictionary::AsyncService {

};


void runServer(const std::string &filepath)
{
  std::string server_addr("localhost:8080");
  DictionaryImplSync service(filepath);
  ServerBuilder builder;
  builder.AddListeningPort(server_addr, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<Server> server(builder.BuildAndStart());  
  std::cout << "======== server on localhost:8080 ========" << std::endl;
  server->Wait();
}


int main(int argc, char *argv[])
{
  std::string path("dic.txt");
  runServer(path);
  return 0;
}
