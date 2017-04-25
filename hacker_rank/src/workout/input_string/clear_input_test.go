package input_string

import "testing"

func TestClear_inout(t *testing.T) {
  var expected = "Hello World"
  var text = "Hello World\n"

 var result = Clear_input(text)

 if result != expected {
   t.Error("expected", expected, "got", result)
 }
}
