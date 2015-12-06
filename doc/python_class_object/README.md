[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Python: class, object

- [Object Oriented Programming](#object-oriented-programming)
- [class](#class)

[↑ top](#python-class-object)
<br><br><br><br><hr>


### Object Oriented Programming

*Steve Jobs* explains [here](http://www.edibleapple.com/2011/10/29/steve-jobs-explains-object-oriented-programming/):

> **_Objects_** are like people. They’re living, breathing 
> things that have knowledge inside them about 
> **how to do things** and **have memory** inside them
> so they can remember things.

An object can contains:

- how the object does things - `method`
- memory to store things

[↑ top](#python-class-object)
<br><br><br><br><hr>


#### class

```python
#!/usr/bin/python -u

class Mail:

    def __init__(self):
        self.html_list = []
        self.attachments = []

    def html(self, h): self.html_list.append(h)

    def attach(self, t):
        self.attachments.append(t)

    def h3(self, t):
        self.html("<h3>%s</h3>" % t)

    def p(self, t):
        self.html("<p>%s</p>" % t)

    def describe(self):
        print "To:", self.to
        print "Subject:", self.subject
        print "Body:", "\n".join([str(c) for c in self.html_list])
        print "Attach:", self.attachments



if __name__ == "__main__":
    m = Mail()
    m.to = "mail@mail.com"
    m.subject = "Hello"
    m.h3("Hello this is Python!")
    m.p("Thanks for your visit.")
    m.p("See you again!")
    m.attach("test.txt")

    m.describe()

    """
    To: mail@mail.com
    Subject: Hello
    Body: <h3>Hello this is Python!</h3>
    <p>Thanks for your visit.</p>
    <p>See you again!</p>
    Attach: ['test.txt']
    """

```

[↑ top](#python-class-object)
<br><br><br><br><hr>
