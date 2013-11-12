onepagego : OnePageCRM API Wrapper written in Go
=================================================

This is an API wrapper package to interface with 
[OnePageCRM API](http://www.onepagecrm.com/api/sales-crm-api.html), written in Go.

This is currently a work-in-progress.

I am currently targeting the stable API. Though API v3 beta is out
already, it's incomplete and has quirks.

_Disclaimer_: I am not affiliated with OnePageCRM other than being one of
their users. This (onepagego) is not officially
endorsed by OnePageCRM team. That said, you should check out
[OnePageCRM](https://www.onepagecrm.com/) if you haven't yet.

Install
=======

You would want to do

    go get github.com/dotmanish/onepagego

to download and install the *onepagego* package.
This requires you to have configured GOPATH variable correctly in your
environment.


API Wrapper Package Usage
=========================

Typical usage would entail:

1. import "github.com/dotmanish/onepagego"

2. call InitOnePageWithUserPass()
   
3. call the Main APIs or Helper Functions. I am still in process of
adding the Main APIs and Helper Functions.
    

**Currently available APIs:**

**Initialization:** 

    InitOnePageWithUserPass

**Separate Auth APIs:**

    GetNewAuthKey



License
=======

Use of this source code is governed by a BSD (3-Clause) License.

Copyright 2013 Manish Malik (manishmalik.name)

All rights reserved.
    
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright notice,
      this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright notice,
      this list of conditions and the following disclaimer in the documentation
      and/or other materials provided with the distribution.
    * Neither the name of this program/product nor the names of its contributors may
      be used to endorse or promote products derived from this software without
      specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
