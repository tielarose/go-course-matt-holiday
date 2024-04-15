package hello

import "testing"

func TestSayHello(t *testing.T){
    // want := "Hello, test!"
    // got := Say([]string{"test"})

    // if want != got {
    //     t.Errorf("wanted %s, got %s", want, got)
    // }

    // subtests is a list of dictionaries
    // each dict:
    //    items: list of names
    //    result: what we'd expect the result string to be
    
    subtests := []struct{
        items []string
        result string
    }{
        {
            result: "Hello, world!",
        },
        {
            items: []string{"Tiela"},
            result: "Hello, Tiela!",
        },
        {
            items: []string{"Tiela", "Pati", "Joseph"},
            result: "Hello, Tiela, Pati, Joseph!",
        },
    }

    for _, st := range subtests {
        if s := Say(st.items); s != st.result {
            t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
        }
    }
}