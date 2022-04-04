/*
 * Copyright (C) 2022 Hopsworks AB
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301,
 * USA.
 */
#ifndef PKR_RESPONSE
#define PKR_RESPONSE

#include <stdint.h>
#include <cstring>
#include <string>
#include "src/rdrslib.h"
#include "src/status.hpp"
#include "src/error-strs.h"

using namespace std;

class PKRResponse {

private:
  char *respBuff;
  uint32_t capacity = 512; //TODO FIX ME
  uint32_t writeHeader = 0;

public:

  /**
   * Get maximum capacity of the response buffer 
   *
   * @return max capacity
   */
  uint32_t getMaxCapacity();

  /**
   * Get remaining capacity of the response buffer 
   *
   * @return remaining capacity
   */
  uint32_t getRemainingCapacity();

  /**
   * Append to response buffer
   */
  RS_Status append_string(string str, bool appendComma); 

  /**
   * Append to response buffer
   */
  RS_Status append_cstring(const char* str, bool appendComma); 

  PKRResponse(char *respBuff);

  char *getResponseBuffer();

  /**
   * Get write header location
   */
  uint32_t getWriteHeader();

  /**
   * Append to response buffer
   */
  RS_Status append_iu32(uint32_t num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_i32(int num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_i64(long long num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_iu64(unsigned long long num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_i8(char num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_iu8(unsigned char num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_i16(short int num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_iu16(unsigned short int num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_i24(int num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_iu24(unsigned int num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_f32(float num, bool appendComma);

  /**
   * Append to response buffer
   */
  RS_Status append_d64(double num, bool appendComma);

  /**
   * Append to response buffer. Append 
   */
  RS_Status append_char(const char *from_buffer, uint32_t from_length, CHARSET_INFO *from_cs, bool appendComma);

  /**
   * Append null. Used to terminate string response message
   */
  RS_Status appendNULL();
};

#endif
