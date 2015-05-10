##########
GOLANG CMS
##########


Open source enterprise content management system based on the BeeGO framework, inspired by django cms.


********
Features
********

* Hierarchical pages
* Extensive support for multilingual websites
* Multi site support
* Draft/Published workflows
* Undo/Redo
* Use the content blocks (placeholders) in your own apps (models)
* Use the content blocks (static placeholders) anywhere in your templates
* Edit content directly in the frontend on your pages
* Hierarchical content plugins (columns, style changes etc)
* Navigation rendering and extending from your apps
* SEO friendly urls
* Highly integrative into your own apps
* Mobile support

You can define editable areas, called placeholders, in your templates which you fill
with many different so called CMS content plugins.
A list of all the plugins will be found here:

`3rd party plugins <http://www.3party-cms.com/golang-cms/>`_

Should you be unable to find a suitable plugin for you needs, writing your own is very simple.

More information on `our website <http://www.golang-cms-url.org>`_.

*************
Documentation
*************

Please head over to our `documentation <http://docs.goland-cms.org/>`_ for all
the details on how to install, extend and use the goland CMS.

********
Tutorial
********

http://docs.golang-cms-url.org/en/latest/introduction/index.html

***********
Quick Start
***********

To complire and run the installer execute.

- go get -u  github.com/astaxie/beego
- go get -u  github.com/beego/bee
- go get -u  github.com/Shaked/gomobiledetect
- go get -u  github.com/garyburd/redigo/redis
- go get -u  github.com/dionyself/golang-cms
- cd $GOPATH/src/github.com/dionyself/golang-cms
- bee run
