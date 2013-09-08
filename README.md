journal
=======

A simple go blogging platform backed by mirrors.  Currently, we run on [mirrors](https://github.com/hayesgm/mirrors) which requires DNSimple for mirroring.

# Building and Running

    go build journal.go

    ./journal --domain="www.example.com" --token="<DOMAIN KEY>"

# TODOs

* Currently, there is no way to add posts.  This would be considered important.
* Also, want to run on [Fiddler](https://github.com/hayesgm/fiddler) for auto-scaling