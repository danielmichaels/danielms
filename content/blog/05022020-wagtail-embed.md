+++
title = "Wagtail embeded YouTube videos"
categories = ["python", "Django"]
tags = ["python"]
slug = "wagtail-embedurl-youtube-tags"
date = "2020-01-25"
draft = "false"
+++

## Wagtail Embed Video

![](/images/wagtail-logo.png 'wagtail icon')

Wagtail is a brilliant content management service built atop of Django. It comes with all of Django's functionality and Django Rest Framework built in for headless work making life a lot easier.

In all, I really love it, especially its [StreamFields][0]. But I did stumble on one component, getting embedded videos to work correctly.

### Embedding videos

This is harder than I feel it should be and poorly explained by the Wagtail documentation. Perhaps too harsh, in a sense - the documentation is actually really good. Except for in the circumstance whereby you want to render an embedded video inside your own HTML and CSS.

Before, covering that, again I should mention that wagtail's documentation is excellent and so too is their implementation of this Framework. For instance, you _can_ render an embed video 'block' by calling a template tag and creating the component in the model.

Except, it will __always__ render this block inside a custom CSS div which apparently cannot be overridden. I'd wager 8 out of 10 times this might mess with your sites style.

### Custom embedding of video content

There is a solution to this issue; just call the _video.url_ within your own implementation of an iframe element. Except, if its YouTube, and your are using a URL such as https://www.youtube.com/watch?v=xUWd3o6z2bk, it will likely not run. Why? It's the _watch_ url endpoint. YouTube, wants embedded videos to use https://www.youtube.com/embed?v=xUWd3o6z2bk instead. __Note__ the use of _embed_ in place of _watch_.

And, wagtail will not accept YouTube urls containing _embed_ causing a template error. So what's a person to do; custom template tags to override this blocker!

### Code or GTFO

Lets create a VideoBlock that will embed a video component into our template.

```python
# blocks.py
class VideoBlock(blocks.StructBlock):
    """Only used for Video Card modals."""
    video = EmbedBlock() # <-- the part we need

    class Meta:
        template = "streams/video_card_block.html"
        icon = "media"
        label = "Embed Video"
```

This will now give us a StreamField which we can then use to enter the url we wish to embed into the template. In this example the url we are entering into our StreamField video_card_block is https://www.youtube.com/watch?v=xUWd3o6z2bk

Next, in the video block template we'll use Bootstrap's embedded video component and inject the url into the _src_ attribute.

```html
<!--  video_card_block.html -->
{% load wagtailcore_tags wagtailembeds_tags %}

{% block content %}
    <div class="embed-responsive embed-responsive-16by9">
  <iframe class="embed-responsive-item" src="{{ self.video.url }}?rel=0" allowfullscreen></iframe>
</div>
{% endblock %}
```

Unfortunately, this will render it as seen below, and be blocked by YouTube. In fact when it is blocked, you will need to consult dev tools to see that wagtail did indeed render the HTML.

```html
<div class="embed-responsive embed-responsive-16by9">
  <iframe class="embed-responsive-item" src="https://www.youtube.com/watch/zpOULjyy-n8?rel=0" allowfullscreen></iframe>
</div>
```

To fix this we need to create a [custom template tag][1] which will intercept the _watch_ url and inject the _embed_ url into the template. This will make wagtail's embed [classifiers][2] happy and us too.

```python
import re
from django import template

register = template.Library()

@register.filter(name="embedurl")
def get_embed_url_with_parameters(url):
    if "youtube.com" in url or "youtu.be" in url:
        regex = r"(?:https:\/\/)?(?:www\.)?(?:youtube\.com|youtu\.be)\/(?:watch\?v=)?(.+)"  # Get video id from URL
        embed_url = re.sub(
            regex, r"https://www.youtube.com/embed/\1", url
        )  # Append video id to desired URL
        print(embed_url)
        embed_url_with_parameters = embed_url + "?rel=0"  # Add additional parameters
        return embed_url_with_parameters
    else:
        return None
```

This will now render the _embed_ url as expected and upon inspection with dev tools should look like so:

```html
<div class="embed-responsive embed-responsive-16by9">
  <iframe class="embed-responsive-item" src="https://www.youtube.com/embed/zpOULjyy-n8?rel=0" allowfullscreen></iframe>
</div>
```

### fixed

I Searched the internet for hours and only after piecing it together over time did I eventually unlock this riddle! This is still an [open][3] ticket on the tail issue tracker!

[0]: https://docs.djangoproject.com/en/dev/howto/custom-template-tags/
[1]: https://docs.djangoproject.com/en/dev/howto/custom-template-tags/
[2]: https://github.com/wagtail/wagtail/blob/master/wagtail/embeds/oembed_providers.py
[3]: https://github.com/wagtail/wagtail/issues/4127
