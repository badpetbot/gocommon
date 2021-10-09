package util

import(

  // Import builtin packages.
  "crypto/sha1"
  "crypto/sha256"
  "encoding/hex"
  "fmt"
  "math/rand"
  "regexp"
  "strconv"
  "strings"
  "time"

  // Import 3rd party packages.
  "github.com/google/uuid"
)

const aRegexT = "a|4"
var   aRegexS = regexp.MustCompile(aRegexT)
const eRegexT = "e|3"
var   eRegexS = regexp.MustCompile(eRegexT)
const iRegexT = "i|l|1"
var   iRegexS = regexp.MustCompile(iRegexT)
const oRegexT = "o|0"
var   oRegexS = regexp.MustCompile(oRegexT)
const sRegexT = "s|5"
var   sRegexS = regexp.MustCompile(sRegexT)
const tRegexT = "t|7"
var   tRegexS = regexp.MustCompile(tRegexT)
const xRegexT = "\\-|\\_"
var   xRegexS = regexp.MustCompile(xRegexT)

// Regexify translates ordinary text into a string representation of a regular expression with
// letters and numbers similar to each other translated into non-capturing OR groups.
func Regexify(in string) (out string) {

  out = strings.ToLower(in)
  out = aRegexS.ReplaceAllString(out, "(?:"+aRegexT+")")
  out = eRegexS.ReplaceAllString(out, "(?:"+eRegexT+")")
  out = iRegexS.ReplaceAllString(out, "(?:"+iRegexT+")")
  out = oRegexS.ReplaceAllString(out, "(?:"+oRegexT+")")
  out = sRegexS.ReplaceAllString(out, "(?:"+sRegexT+")")
  out = tRegexS.ReplaceAllString(out, "(?:"+tRegexT+")")
  out = xRegexS.ReplaceAllString(out, "(?:"+xRegexT+")")

  return
}

// SpecificRand gets a random number in the specified range, min and max inclusive.
func SpecificRand(min, max int) int {

  rand.Seed(ToUtcMilliseconds(time.Now()))
  return rand.Intn(max-min) + min
}

// SpecificRand gets a random number in the specified range, min and max inclusive.
func SpecificRandDuration(min, max time.Duration) time.Duration {

  rand.Seed(ToUtcMilliseconds(time.Now()))
  return time.Duration(rand.Int63n(int64(max-min)) + int64(min))
}

// Substring returns a substring of the specified length, beginning at the specified location.
// If length is 0, returns a substring beginning at the specified location and ending at the input's end.
// Panics if the start or length are negative, or if the start plus the length is greater than the input's length.
func Substring(in string, start, length int) string {
  
  if (start > len(in)) { panic(fmt.Sprintf("Out of bounds: Substring start outside length of string, length %v.", len(in))) }
  
  if (start < 0 || length < 0) { panic(fmt.Sprintf("Out of bounds: Substring start and/or length negative.")) }
  
  lgth := length
  if (lgth == 0 || lgth > len(in)) { lgth = len(in) - start }
  
  return in[start:(start+lgth)]
}

// SplitStringByLength splits a string into multiple strings at most the specified length.
// If useWhitespace is true, will split around whitespace rather than in the middle of a word.
func SplitStringByLength(in string, length int, useWhitespace bool) []string {
  
  toRet := make([]string, 0)
  currentChar := 0
  for (currentChar < len(in)) {
    
    thisLen := IntMin(length, len(in) - currentChar)
    thisString := Substring(in, currentChar, thisLen)
    
    if (useWhitespace && currentChar + thisLen < len(in)) {
      
      wsLen := strings.LastIndex(thisString, " ")
      thisString = Substring(thisString, 0, wsLen)
      currentChar += wsLen + 1
    } else {
      
      currentChar += thisLen
    }
    
    toRet = append(toRet, thisString)
  }
  
  return toRet
}

// ToUtcMilliseconds gets milliseconds since Jan 1, 1970 12:00:00.0000.
func ToUtcMilliseconds(t time.Time) int64 {
  return t.UTC().UnixNano() / int64(time.Millisecond)
}

// FromUtcMilliseconds gets a time object using the given milliseconds.
func FromUtcMilliseconds(m int64) time.Time {
  return time.Unix(0, m * int64(time.Millisecond)).UTC()
}

// IntMin returns the lesser of two integers.
func IntMin(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// IntMax returns the greater of two integers.
func IntMax(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// RandomSha1 returns a random sha1 digest of the current time. This is not
// intended for uses which require many guaranteed unique sha1's quickly.
func RandomSha1() string {
  
  hash := sha1.Sum([]byte(strconv.FormatInt(ToUtcMilliseconds(time.Now().Add(time.Duration(SpecificRand(-10000, 10000)) * time.Millisecond)), 10)))
  return string(hex.EncodeToString(hash[:]))
}

// RandomSha256 returns a random sha256 digest of the current time. This is not
// intended for uses which require many guaranteed unique sha256's quickly.
func RandomSha256() string {
  
  hash := sha256.Sum256([]byte(strconv.FormatInt(ToUtcMilliseconds(time.Now().Add(time.Duration(SpecificRand(-10000, 10000)) * time.Millisecond)), 10)))
  return string(hex.EncodeToString(hash[:]))
}

// RandomUUID generates a random UUID string.
func RandomUUID() string {

  return uuid.New().String()
}