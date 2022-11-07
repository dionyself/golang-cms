##########
GOLANG CMS
##########


Open source Content Management System based on the BeeGO framework, inspired by Django CMS.


********
Features
********

* Hierarchical categories
* Extensive support for multilingual websites.  #TODO
* Use the content blocks (placeholders) in your own Templates
* Edit content directly in the frontend on your pages.  #TODO
* Navigation rendering and extending from your apps.
* SEO friendly urls.
* Mobile support.

You can define editable areas, called placeholders, in your templates which you fill
with many different so called CMS content plugins.
A list of all the plugins will be found here:

`3rd party plugins <http://www.3party-cms.com/golang-cms/>`_ #TODO

Unable to find a suitable plugin for you needs? Writing your own is very simple.

More information on `our website <http://www.golang-cms-url.org>`_.  #TODO

***********
Quick Start
***********

To compile and run the installer execute. (you will need to use GO v1.7.4 or later)

- go get -u  github.com/dionyself/beego
- go get -u  github.com/beego/bee
- go get -u  github.com/dionyself/golang-cms
- cd $GOPATH/src/github.com/dionyself/golang-cms
- bee run

Browse 127.0.0.1:8080 to see GolangCMS running.
Login details. user: test, password: test

- To create new articles visit http://127.0.0.1:8080/article/0/edit
- To view an article visit http://127.0.0.1:8080/article/<article_id>/show
- ex. http://127.0.0.1:8080/article/2/show

Note: You will be running a pre-alpha version in testmode.

*************
Documentation
*************

The current state of the project 'pre-alpha' version,
Only Linux based OS are supported, please report any bug you find.
Please head over to our `documentation <http://docs.goland-cms.org/>`_ for all
the details on how to install, extend and use the golang CMS.

http://docs.golang-cms-url.org/en/latest/introduction/index.html  #TODO

****
Demo
****

You will need to run the demo locally.
If you are using Docker you can run a golang-cms instance on port 8080
just run:

- docker run -p 8080:8080 dionyself/golang-cms:latest
Browse 127.0.0.1:8080 to see the demo.

if you can't see the demo please contact me.

To run unittests.

 - goconvey $GOPATH/src/github.com/dionyself/golang-cms/tests/

Bitcoin Donations.

.. image:: https://scontent.fsdq1-2.fna.fbcdn.net/v/t39.30808-6/313430088_2599438973524215_7879645149016306497_n.jpg?_nc_cat=100&ccb=1-7&_nc_sid=730e14&_nc_eui2=AeEphyk5KATsx2ad8SuxMKSUodzbm3Kxjk2h3NubcrGOTY3pYJVIVQvrKvA3TQe2Ui4CpanN7BHRZ93_KaJGhiHh&_nc_ohc=i409G20bmsAAX-Momwn&_nc_zt=23&_nc_ht=scontent.fsdq1-2.fna&oh=00_AfCIT6Bhsql0O55E3IfrNy2JqG4KAItmFMwZHkxAQADq5g&oe=636D99DE

XMR (Monero) Donations.

.. image:: https://scontent.fsdq1-1.fna.fbcdn.net/v/t39.30808-6/313864958_2599496130185166_4335848911691923872_n.jpg?_nc_cat=107&ccb=1-7&_nc_sid=730e14&_nc_eui2=AeGLnUVv6hW29LoGaFv7C_Q6iKwTeG849lyIrBN4bzj2XK-GDHh9rQaoLA2kLDrlWQr7bRDpVlsYp0jrgoKOJadi&_nc_ohc=aSjd4-gCmOkAX_ZoNve&_nc_zt=23&_nc_ht=scontent.fsdq1-1.fna&oh=00_AfAZP701kDSzpDin2RBLH0CzTGhVAuNAOmJz0NqqWMe31Q&oe=636D09E5
   :height: 300px
   :width: 300 px

Paypal

.. image:: https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif
   :height: 100px
   :width: 200 px
   :scale: 50 %
   :alt: alternate text
   :align: right
   :target: https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=L4H5TUWZTZERS
