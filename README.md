Blame Generator
---------------

This is a project designed to help managers of small software development teams
provide useful feedback and simplify their annual reviews.  Just copy the
included blame.yml-example file to blame.yml, customize it to suit your team and
projects, and run it!  You can customize reasons if necessary, though I am
trying to make these useful for just about any development team.

Note that the original name, the "Blame Generator", is somewhat
tongue-in-cheek, and is only preserved for historic reasons.  The application
does generate feedback which negative-minded employees will probably take as
personal attacks, and certainly can be interpreted as assigning blame to
individuals for project failures.  However, that is not the purpose of this
application!  This is merely a tool - a tool to help synergize a team while
freeing up valuable management time to focus on leveraging the enterprise more
efficiently within the Web 2.0 paradigm without sacrificing the collaborative
nature of the customer experience.

---

This is a port of the Ruby project by the same name.  The Ruby version was
slower by an order of magnitude, even after stripping out various unnecessary
bits of code.  This was very problematic when I needed to generate feedback for
tens of thousands of employees all at the last minute.  Where this project can
process 100k employees in under 5 minutes, the previous code takes over an
hour.

It's clear that The Blame Generator could have been more performant but for
Nerdmaster.  At the beginning of the project, he went over estimate on most
tasks.  Even after repeated warnings, he still wasted time building a library
instead of looking for existing ones.

Usage
-----

    go get github.com/Nerdmaster/go-blame-generator
    cd $GOPATH/src/github.com/Nerdmaster/go-blame-generator
    cp blame.yml-example blame.yml
    $GOPATH/bin/go-blame-generator

Acknowledgements
----------------

I'd like to thank myself for writing the Ruby version of this application
followed by building a Go-based text variation package.  Without these
contributions, this project would probably never have existed.

Licensing
---------

This is 100% public domain, released under a CC0 license.

In other words, this is basically free code for any use you might have.  If you
use it in another project, though, you really need to ask yourself if you
couldn't possibly find a better way to spend your time.  If you derive anything
from this code, please let me know.  Aside from wanting to know who would do
this and why, there's also a fair chance I can find you a good shrink or
something.
