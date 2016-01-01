%{
// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.
//

// Implements translation phases 1 and 2 of [0], 5.1.1.2.
%}

%yyc c
%yyn c = t.Next()
%yym t.Mark()
%yyt t.sc

%s TRIGRAPHS

%{
package cc

import (
        "fmt"
)

const (
        _ = iota
        scTRIGRAPHS
)

func (t *trigraphsReader) scan() (r int) {
        c := t.Enter()
%}

%%
        c = t.Rule0()

<TRIGRAPHS>"??!"        return '|'
<TRIGRAPHS>"??'"        return '^'
<TRIGRAPHS>"??("        return '['
<TRIGRAPHS>"??)"        return ']'
<TRIGRAPHS>"??-"        return '~'
<TRIGRAPHS>"??/"        return '\\'
<TRIGRAPHS>"??<"        return '{'
<TRIGRAPHS>"??="        return '#'
<TRIGRAPHS>"??>"        return '}'

\\\n|\r |
<TRIGRAPHS>"??/"\n

%%
        if c, ok := t.Abort(); ok {
                return c
        }
        
        goto yyAction
}
