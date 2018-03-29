package main

func Index(vs []string, t string) int {
  for i, v := range vs {
    if v == t {
      return i
    }
  }
  return -1
}

func Include(vs []string, t string) bool {
  return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
  for _, v := range vs {
    if f(v) {
      return true
    }
  }
  return false
}
