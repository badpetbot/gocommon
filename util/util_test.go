package util

import(
  "regexp"
  "strings"
  "testing"
  "time"
)

func TestSpecificRand(t *testing.T) {

  t.Run("100 times: 0 to 10 without 10", func(t *testing.T) {

    for i := 0; i < 100; i++ {

      got := SpecificRand(0, 10)

      if got < 0 || got >= 10 {
        t.Errorf("Wanted 0 <= N < 10, got %d", got)
        return
      }
    }
  })

  t.Run("100 times: 25 to 50 without 50", func(t *testing.T) {

    for i := 0; i < 100; i++ {

      got := SpecificRand(25, 50)

      if got < 25 || got >= 50 {
        t.Errorf("Wanted 25 <= N < 50, got %d", got)
        return
      }
    }
  })
}

func TestSpecificRandDuration(t *testing.T) {

  t.Run("100 times: time.Nanosecond to time.Microsecond without time.Microsecond", func(t *testing.T) {

    for i := 0; i < 100; i++ {

      got := SpecificRandDuration(time.Nanosecond, time.Microsecond)

      if got < time.Nanosecond || got >= time.Microsecond {
        t.Errorf("Wanted time.Nanosecond <= N < time.Microsecond, got %d", got)
        return
      }
    }
  })

  t.Run("100 times: time.Microsecond to time.Millisecond without time.Millisecond", func(t *testing.T) {

    for i := 0; i < 100; i++ {

      got := SpecificRandDuration(time.Microsecond, time.Millisecond)

      if got < time.Microsecond || got >= time.Millisecond {
        t.Errorf("Wanted time.Microsecond <= N < time.Millisecond, got %d", got)
        return
      }
    }
  })
}

func TestSubstring(t *testing.T) {

  full := "Test this substring function!"

  t.Run("from-beginning", func(t *testing.T) {

    got := Substring(full, 0, 10)
    want := "Test this "
    if got != want {
      t.Errorf("Wanted %q but got %q", want, got)
    }
  })

  t.Run("to-end", func(t *testing.T) {

    got := Substring(full, 10, 0)
    want := "substring function!"
    if got != want {
      t.Errorf("Wanted %q but got %q", want, got)
    }
  })

  t.Run("from-beginning-to-end", func(t *testing.T) {

    got := Substring(full, 0, 0)
    want := full
    if got != want {
      t.Errorf("Wanted %q but got %q", want, got)
    }
  })

  t.Run("negative-start-index", func(t *testing.T) {

    func() {

      defer func() {
        if r := recover(); r == nil {

          t.Errorf("Negative start index did not panic.")
        }
      }()

      Substring(full, -1, 5)
    }()
  })

  t.Run("negative-end-index", func(t *testing.T) {

    func() {

      defer func() {
        if r := recover(); r == nil {

          t.Errorf("Negative end index did not panic.")
        }
      }()

      Substring(full, 0, -1)
    }()
  })

  t.Run("end-index-too-high", func(t *testing.T) {

    got := Substring(full, 0, 100)
    want := full
    if got != want {
      t.Errorf("Wanted %q but got %q", want, got)
    }
  })
}

func TestSplitStringByLength(t *testing.T) {

  full := "Test splitting this string by its length. The splits should only be about 50 characters long, or less."

  t.Run("without-whitespace", func(t *testing.T) {

    got := SplitStringByLength(full, 50, false)
    want := []string{
      "Test splitting this string by its length. The spli",
      "ts should only be about 50 characters long, or les",
      "s.",
    }

    if len(got) != len(want) {
      t.Errorf("Wanted:\n%q\nbut got:\n%q", strings.Join(want, "\n"), strings.Join(got, "\n"))
    }

    for i := 0; i < len(got); i++ {

      if got[i] != want[i] {
        t.Errorf("Element %d: Wanted %q but got %q", i, strings.Join(want, "\n"), strings.Join(got, "\n"))
        break
      }
    }
  })

  t.Run("with-whitespace", func(t *testing.T) {

    got := SplitStringByLength(full, 50, true)
    want := []string{
      "Test splitting this string by its length. The",
      "splits should only be about 50 characters long,",
      "or less.",
    }
    
    if len(got) != len(want) {
      t.Errorf("Wanted:\n%q\nbut got:\n%q", strings.Join(want, "\n"), strings.Join(got, "\n"))
    }

    for i := 0; i < len(got); i++ {

      if got[i] != want[i] {
        t.Errorf("Element %d: Wanted %q but got %q", i, strings.Join(want, "\n"), strings.Join(got, "\n"))
        break
      }
    }
  })
}

func TestToUtcMilliseconds(t *testing.T) {

  y2k := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

  got := ToUtcMilliseconds(y2k)
  var want int64 = 946684800000

  if got != want {

    t.Errorf("Wanted %d but got %d", want, got)
  }
}

func TestFromUtcMilliseconds(t *testing.T) {

  var y2k int64 = 946684800000

  got := FromUtcMilliseconds(y2k)
  want := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

  if got != want {

    t.Errorf("Wanted %q but got %q", want.Format("2006-01-02T15:04:05Z"), got.Format("2006-01-02T15:04:05Z"))
  }
}

func TestIntMin(t *testing.T) {

  t.Run("minimum-first", func(t *testing.T) {

    got := IntMin(0, 5)
    want := 0

    if got != want {

      t.Errorf("Wanted %d but got %d", want, got)
    }
  })

  t.Run("maximum-first", func(t *testing.T) {

    got := IntMin(5, 0)
    want := 0

    if got != want {

      t.Errorf("Wanted %d but got %d", want, got)
    }
  })
}

func TestIntMax(t *testing.T) {

  t.Run("minimum-first", func(t *testing.T) {

    got := IntMax(0, 5)
    want := 5

    if got != want {

      t.Errorf("Wanted %d but got %d", want, got)
    }
  })

  t.Run("maximum-first", func(t *testing.T) {

    got := IntMax(5, 0)
    want := 5

    if got != want {

      t.Errorf("Wanted %d but got %d", want, got)
    }
  })
}

func TestRandomSha1(t *testing.T) {

  got := len(RandomSha1())
  want := 40

  if got != want {

    t.Errorf("Wanted hash of length %d but got %d", want, got)
  }
}

func TestRandomSha256(t *testing.T) {

  got := len(RandomSha256())
  want := 64

  if got != want {

    t.Errorf("Wanted hash of length %d but got %d", want, got)
  }
}

func TestRandomUUID(t *testing.T) {

  r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

  got := RandomUUID()

  if !r.MatchString(got) {

    t.Errorf("Wanted valid UUID like %q but got %q", "123e4567-e89b-12d3-a456-426655440000", got)
  }
}