+++
title = "DBeaver wordwrap and autoformat by default"
categories = ["zet"]
tags = ["zet"]
slug = "DBeaver-wordwrap-and-autoformat-by-default"
date = "2025-05-08 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# DBeaver wordwrap and autoformat by default

Make the `values` panel auto-format and wordwrap by default. This requires changes to the config files in DBeaver's workdirectory.

Make sure DBeaver is closed before doing this.

Edit these files:

- [workdirectory]/workspace6/.metadata/.plugins/org.eclipse.core.runtime/.settings/org.eclipse.ui.workbench.prefs

```
# add these to bottom of file 
lineNumberRuler=true
wordwrap.enabled=true
```

- [workdirectory]/workspace6/.metadata/.plugins/org.jkiss.dbeaver.ui.editors.data/dialog_settings.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<section name="Workbench">
  <section name="AbstractTextPanelEditor">
    # these two lines
    <item key="content.text.editor.auto-format" value="true"/>
    <item key="content.text.editor.word-wrap" value="true"/>
  </section>
</section>
```

ref: https://superuser.com/a/1827163/862055

Tags:

    #sql

[workdirectory]: https://github.com/dbeaver/dbeaver/wiki/Workspace-Location#default-location-of-the-workspace

