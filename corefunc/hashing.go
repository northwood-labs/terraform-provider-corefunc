// Copyright 2024-2025, Northwood Labs, LLC <license@northwood-labs.com>
// Copyright 2023-2025, Ryan Parman <rparman@northwood-labs.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package corefunc

import (
	"crypto/hmac"
	"crypto/md5"  // lint:not_crypto
	"crypto/sha1" // lint:not_crypto
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"runtime"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

/*
HashMD5 returns the MD5 hash of the input string as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The MD5 hash as a hexadecimal string.
*/
func HashMD5(input string) string {
	return hex.EncodeToString(hashMd5Bytes([]byte(input)))
}

/*
Base64HashMD5 returns the MD5 hash of the input string, as bytes, as a
base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The MD5 hash as a base64-encoded string.
*/
func Base64HashMD5(input string) string {
	return base64.StdEncoding.EncodeToString(hashMd5Bytes([]byte(input)))
}

/*
HashSHA1 returns the SHA-1 hash of the input string as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-1 hash as a hexadecimal string.
*/
func HashSHA1(input string) string {
	return hex.EncodeToString(hashSha1Bytes([]byte(input)))
}

/*
Base64HashSHA1 returns the SHA-1 hash of the input string, as bytes, as a
base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-1 hash as a base64-encoded string.
*/
func Base64HashSHA1(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha1Bytes([]byte(input)))
}

/*
HashSHA256 returns the SHA-2/256-bit hash of the input string as a hexadecimal
string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-256 hash as a hexadecimal string.
*/
func HashSHA256(input string) string {
	return hex.EncodeToString(hashSha256Bytes([]byte(input)))
}

/*
Base64HashSHA256 returns the SHA-2/256-bit hash of the input string, as bytes,
as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-256 hash as a base64-encoded string.
*/
func Base64HashSHA256(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha256Bytes([]byte(input)))
}

/*
HashSHA384 returns the SHA-2/384-bit hash of the input string as a hexadecimal
string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-384 hash as a hexadecimal string.
*/
func HashSHA384(input string) string {
	return hex.EncodeToString(hashSha384Bytes([]byte(input)))
}

/*
Base64HashSHA384 returns the SHA-2/384-bit hash of the input string, as bytes,
as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-384 hash as a base64-encoded string.
*/
func Base64HashSHA384(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha384Bytes([]byte(input)))
}

/*
HashSHA512 returns the SHA-2/512-bit hash of the input string as a hexadecimal
string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-512 hash as a hexadecimal string.
*/
func HashSHA512(input string) string {
	return hex.EncodeToString(hashSha512Bytes([]byte(input)))
}

/*
Base64HashSHA512 returns the SHA-2/512-bit hash of the input string, as bytes,
as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA-512 hash as a base64-encoded string.
*/
func Base64HashSHA512(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha512Bytes([]byte(input)))
}

/*
HashSHA3x256 returns the SHA-3/256-bit hash of the input string as a
hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA3-256 hash as a hexadecimal string.
*/
func HashSHA3x256(input string) string {
	return hex.EncodeToString(hashSha3_256Bytes([]byte(input)))
}

/*
Base64HashSHA3x256 returns the SHA-3/256-bit hash of the input string, as
bytes, as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA3-256 hash as a base64-encoded string.
*/
func Base64HashSHA3x256(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha3_256Bytes([]byte(input)))
}

/*
HashSHA3x384 returns the SHA-3/384-bit hash of the input string as a
hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA3-384 hash as a hexadecimal string.
*/
func HashSHA3x384(input string) string {
	return hex.EncodeToString(hashSha3_384Bytes([]byte(input)))
}

/*
Base64HashSHA3x384 returns the SHA-3/384-bit hash of the input string, as
bytes, as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA3-384 hash as a base64-encoded string.
*/
func Base64HashSHA3x384(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha3_384Bytes([]byte(input)))
}

/*
HashSHA3x512 returns the SHA-3/512-bit hash of the input string as a
hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA3-512 hash as a hexadecimal string.
*/
func HashSHA3x512(input string) string {
	return hex.EncodeToString(hashSha3_512Bytes([]byte(input)))
}

/*
Base64HashSHA3x512 returns the SHA-3/512-bit hash of the input string, as
bytes, as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - Returns (string): The SHA3-512 hash as a base64-encoded string.
*/
func Base64HashSHA3x512(input string) string {
	return base64.StdEncoding.EncodeToString(hashSha3_512Bytes([]byte(input)))
}

/*
HashBcrypt returns the bcrypt hash of the input string as a byte slice.

----

  - input (string): The input string to hash.

  - cost (int): The bcrypt cost parameter (work factor). Common values are
    between 10 and 14. Higher values are more secure but take longer to compute.
    The default value is 10.

  - Returns (string, error): The bcrypt hash as a byte slice, or an error if
    the hashing operation fails.
*/
func HashBcrypt(input string, cost ...int) (string, error) {
	if len(cost) > 1 {
		return "", errors.New("only one cost parameter is allowed") // lint:allow_errorf
	}

	c := bcrypt.DefaultCost
	if len(cost) == 1 {
		c = cost[0]
	}

	out, err := bcrypt.GenerateFromPassword([]byte(input), c)
	if err != nil {
		return "", fmt.Errorf("could not generate bcrypt hash: %w", err)
	}

	return string(out), nil
}

/*
HashArgon2id returns the Argon2id hash of the input string as a hexadecimal
string.

Time is fixed to 1. Memory is fixed to 64 MiB. Parallelism is set to the number
of CPU cores. Output length of the raw hash is fixed to 32 bytes before
hexadecimal encoding.

----

  - input (string): The input string to hash.

  - salt ([]byte): The salt to use for hashing. Must not be empty.

  - Returns (string, error): The Argon2id hash as a hexadecimal string, or an
    error if the hashing operation fails.
*/
func HashArgon2id(input string, salt []byte) (string, error) {
	hash, err := hashArgon2idBytes([]byte(input), salt)
	if err != nil {
		return "", fmt.Errorf("could not generate Argon2id hash: %w", err)
	}

	return hex.EncodeToString(hash), nil
}

/*
Base64HashArgon2id returns the Argon2id hash of the input string, as bytes,
as a base64-encoded string instead of as a hexadecimal string.

Time is fixed to 1. Memory is fixed to 64 MiB. Parallelism is set to the number
of CPU cores. Output length of the raw hash is fixed to 32 bytes before
base64 encoding.

----

  - input (string): The input string to hash.

  - salt ([]byte): The salt to use for hashing. Must not be empty.

  - Returns (string, error): The Argon2id hash as a base64-encoded string, or
    an error if the hashing operation fails.
*/
func Base64HashArgon2id(input string, salt []byte) (string, error) {
	hash, err := hashArgon2idBytes([]byte(input), salt)
	if err != nil {
		return "", fmt.Errorf("could not generate Argon2id hash: %w", err)
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}

/*
HashScrypt returns the scrypt hash of the input string as a hexadecimal string.

----

  - input (string): The input string to hash.

  - salt ([]byte): The salt to use for hashing. Must not be empty.

  - Returns (string, error): The scrypt hash as a hexadecimal string, or an
    error if the hashing operation fails.
*/
func HashScrypt(input string, salt []byte) (string, error) {
	hash, err := hashScryptBytes([]byte(input), salt)
	if err != nil {
		return "", fmt.Errorf("could not generate scrypt hash: %w", err)
	}

	return hex.EncodeToString(hash), nil
}

/*
Base64HashScrypt returns the scrypt hash of the input string, as bytes, as a
base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - salt ([]byte): The salt to use for hashing. Must not be empty.

  - Returns (string, error): The scrypt hash as a base64-encoded string, or an
    error if the hashing operation fails.
*/
func Base64HashScrypt(input string, salt []byte) (string, error) {
	hash, err := hashScryptBytes([]byte(input), salt)
	if err != nil {
		return "", fmt.Errorf("could not generate scrypt hash: %w", err)
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}

/*
HashHMACSHA256 returns the HMAC-SHA256 hash of the input string as a
hexadecimal string.

----

  - input (string): The input string to hash.

  - key (string): The secret key used for HMAC.

  - Returns (string): The HMAC-SHA256 hash as a hexadecimal string.
*/
func HashHMACSHA256(input, key string) string {
	return hex.EncodeToString(hashHMACSHA256Bytes([]byte(input), []byte(key)))
}

/*
Base64HashHMACSHA256 returns the HMAC-SHA256 hash of the input string, as
bytes, as a base64-encoded string instead of as a hexadecimal string.

----

  - input (string): The input string to hash.

  - key (string): The secret key used for HMAC.

  - Returns (string): The HMAC-SHA256 hash as a base64-encoded string.
*/
func Base64HashHMACSHA256(input, key string) string {
	return base64.StdEncoding.EncodeToString(hashHMACSHA256Bytes([]byte(input), []byte(key)))
}

// -----------------------------------------------------------------------------
// Internal hash functions

func hashMd5Bytes(input []byte) []byte {
	h := md5.New() // lint:not_crypto
	h.Write(input)

	return h.Sum(nil)
}

func hashSha1Bytes(input []byte) []byte {
	h := sha1.New() // lint:not_crypto
	h.Write(input)

	return h.Sum(nil)
}

func hashSha256Bytes(input []byte) []byte {
	h := sha256.New()
	h.Write(input)

	return h.Sum(nil)
}

func hashSha384Bytes(input []byte) []byte {
	h := sha512.New384()
	h.Write(input)

	return h.Sum(nil)
}

func hashSha512Bytes(input []byte) []byte {
	h := sha512.New()
	h.Write(input)

	return h.Sum(nil)
}

func hashSha3_256Bytes(input []byte) []byte {
	h := sha3.New256()

	_, _ = h.Write(input) // lint:allow_unhandled

	return h.Sum(nil)
}

func hashSha3_384Bytes(input []byte) []byte {
	h := sha3.New384()

	_, _ = h.Write(input) // lint:allow_unhandled

	return h.Sum(nil)
}

func hashSha3_512Bytes(input []byte) []byte {
	h := sha3.New512()

	_, _ = h.Write(input) // lint:allow_unhandled

	return h.Sum(nil)
}

func hashArgon2idBytes(input, salt []byte) ([]byte, error) {
	const (
		memoryBytes = 65536
		bitLength   = 32
	)

	if len(salt) == 0 {
		return nil, errors.New("salt must not be empty") // lint:allow_errorf
	}

	hash := argon2.IDKey(
		input,
		salt,
		1,
		memoryBytes,
		uint8(runtime.NumCPU()), // lint:allow_overflow_in_conversion
		bitLength,
	)

	return hash, nil
}

func hashScryptBytes(input, salt []byte) ([]byte, error) {
	const (
		N         = 1 << 15
		r         = 8
		p         = 1
		keyLength = 32
	)

	if len(salt) == 0 {
		return nil, errors.New("salt must not be empty") // lint:allow_errorf
	}

	hash, err := scrypt.Key(input, salt, N, r, p, keyLength)
	if err != nil {
		return nil, fmt.Errorf("could not generate scrypt hash: %w", err)
	}

	return hash, nil
}

func hashHMACSHA256Bytes(input, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(input)

	return h.Sum(nil)
}
