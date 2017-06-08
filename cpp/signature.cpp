#include "signature.h"
#include <openssl/md5.h>
#include <openssl/sha.h>
#include <boost/uuid/uuid.hpp>
#include <boost/uuid/uuid_generators.hpp>
#include <boost/uuid/uuid_io.hpp>
#include <string>

namespace Tan {

std::string uuid() {
  boost::uuids::random_generator gen;
  boost::uuids::uuid id = gen();
  return boost::uuids::to_string(id);
}

// bsd checksum
// 16位，按字节做一次加法以及一次循环右移位
// 避免溢出
int bsdChecksum(const char *s) {
  int ret = 0;
  while (*s) {
    ret = (ret >> 1) + ((ret & 1) << 15);
    ret += *s;
    ret &= 0xffff;
  }
  return ret;
}

// TODO: CRC
// 基本原理: 多项式模2同余除法
// 比如二进制串: A1A2A3...An
//               B1B2B3...Bn

// TODO: pearson hashing

// TODO: Jenkins hash

// TODO: Murmurhash

std::string sha256(const void *data, size_t len) {
  char buf[2];
  unsigned char hash[SHA256_DIGEST_LENGTH];
  SHA256_CTX sha256;
  SHA256_Init(&sha256);
  SHA256_Update(&sha256, data, len);
  SHA256_Final(hash, &sha256);
  std::string res = "";
  for (int i = 0; i < SHA256_DIGEST_LENGTH; i++) {
    sprintf(buf, "%02x", hash[i]);
    res = res + buf;
  }
  return res;
}

std::string md5(const void *data, int len) {
  char buf[2];
  unsigned char hash[MD5_DIGEST_LENGTH];
  MD5_CTX md5;
  MD5_Init(&md5);
  MD5_Update(&md5, data, len);
  MD5_Final(hash, &md5);
  std::string res = "";
  for (int i = 0; i < MD5_DIGEST_LENGTH; i++) {
    sprintf(buf, "%02x", hash[i]);
    res = res + buf;
  }
  return res;
}
}
