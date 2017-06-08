#ifndef _TAN_SIGNATURE_H_
#define _TAN_SIGNATURE_H_

#include <string>

/*
 * some utility functions for uuid, hash,
 * hmac, decode, encode, crypt etc
 *
 */

namespace Tan {

// boost uuid
std::string uuid();
// bsd checksum
int bsdChecksum(const char *s);
std::string sha256(const void *data, size_t len);
std::string md5(const void *data, size_t len);

}  // end namespace Tan

#endif  // end _TAN_SIGNATURE_H_
