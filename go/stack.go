package main

type Stack struct {
  _el []float64
}

func (s *Stack) push(items ...float64) {
  s._el = append(s._el, items...)
}

func (s *Stack) pop(num int) []float64 {
  var resp []float64
  for i := 0; i < num; i++ {
    var n = len((*s)._el) - 1;
    resp = append(resp, (*s)._el[n]);
    (*s)._el = (*s)._el[:n];
  }
  return resp;
}

func (s *Stack) size() int {
  return len((*s)._el);
}
