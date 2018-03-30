package processing

// Processing
// Node test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"testing"
)

func TestNodeNew(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
	n, err := NewNode("58psJTjw9DnKx1Fo2aL93bPJYGKgReeKEDxrATQ9ozwyxCYJfJXa7fCLGyQ2bhMbx1yy2ASfwx3DD5aFWkhwj3eCjLkoqzGzyfvHhC6sBQyxm9zC6SNysM1twtucoNGx3FfXFvh9GZff3chAV", defaultAuthoritiesListPath)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("my URL: ", n.my.Url)
	}
}
