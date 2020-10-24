// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1603402565, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"UTF-8\">\n    <title>⚡ Send me some Sats</title>\n    <link href=\"https://fonts.googleapis.com/css2?family=Source+Code+Pro:wght@400;700&display=swap\" rel=\"stylesheet\" />\n    <style>\n      html {\n        width: 100%;\n        height: 100%;\n        font-size: 1vw;\n      }\n      body {\n        background-image: linear-gradient(60deg, #3d3393 0%, #2b76b9 37%, #2cacd1 65%, #35eb93 100%);\n        font-family: 'Source Code Pro', monospace;\n      }\n      .wrapper {\n        font-weight: 400;\n        line-height: 1.25;\n        margin-bottom: 1rem;\n        font-size: 4em;\n        color: #fff;\n        width: 640px;\n        margin: 1em auto;\n        text-align: center;\n      }\n      .note {\n        font-size: 0.2em;\n      }\n      .wrapper h1 {\n        font-weight: 700;\n      }\n      .wrapper h2 {\n        font-weight: 400;\n      }\n      .wrapper a, .wrapper a:visited, .wrapper a:active {\n        color: #fff;\n        text-decoration:none;\n      }\n      .wrapper p {\n        margin: 0.5em 0;\n      }\n      input, input:focus {\n        outline: none;\n        background-color: transparent;\n        font-size: 1em;\n        font-weight: 700;\n        border: none;\n        border-bottom: 1px solid #fff;\n        text-align: center;\n        color: #fff;\n      }\n      ::placeholder {\n        color: #fff;\n        opacity: 0.3;\n      }\n      input.amount {\n        width: 300px;\n      }\n      input.memo {\n        width: 100%;\n      }\n      button {\n        color: #fff;\n        font-weight: 400;\n        font-size: 0.5em;\n        line-height: 1.5;\n        font-weight: 400;\n        width: 100%;\n        cursor: pointer;\n        border: none;\n        border-radius: 30px;\n        background-color: #2b76b9;\n      }\n      h2.lnme-value, h1.lnme-memo {\n        font-size: 2em;\n        font-weight: 400;\n      }\n\n      .lnme-wrapper {\n        margin-top: 1em;\n        font-size: 0.5em;\n        width: 100%;\n        white-space: nowrap;\n      }\n      .lnme-details {\n        text-overflow: ellipsis;\n        overflow: hidden;\n      }\n      .lnme-link {\n        text-overflow: ellipsis;\n        overflow: hidden;\n      }\n      .lnme-copy {\n        cursor: pointer;\n      }\n    </style>\n  </head>\n  <body>\n    <div class=\"wrapper\">\n      <div class=\"form\" id=\"form\">\n        <p>\n          Send me<br>\n          <input type=\"number\" placeholder=\"10000\" class=\"amount\" id=\"amount\" autocomplete=\"off\" min=\"100\"> Sats\n          <br>\n          for\n          <br>\n          <input type=\"text\" class=\"memo\" id=\"memo\" placeholder=\"message\" autocomplete=\"off\">\n        </p>\n\n        <button id=\"send-button\">⚡ Send ⚡</button>\n\n        <p id=\"onchain\" class=\"note\">\n          <a href=\"#\" id=\"get-new-address\" class=\"onchain\">Prefer onchain Bitcoin? Click here!</a>\n        </p>\n\n        <div id=\"lnme-wrapper\" class=\"lnme-wrapper\" style=\"display:none\">\n          <div class=\"lnme-qrcode\"></div>\n          <div class=\"lnme-details\">\n            <a href=\"#\" class=\"lnme-link\"><span class=\"lnme-payment-request\"></span></a>\n            <div class=\"lnme-copy\">\n              <svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"feather feather-copy\"><rect x=\"9\" y=\"9\" width=\"13\" height=\"13\" rx=\"2\" ry=\"2\"></rect><path d=\"M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1\"></path></svg>\n            </div>\n          </div>\n        </div>\n\n        <div id=\"loader\" style=\"display:none\">\n          <svg version=\"1.1\" id=\"loader-1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" x=\"0px\" y=\"0px\" width=\"40px\" height=\"40px\" viewBox=\"0 0 50 50\" style=\"enable-background:new 0 0 50 50;\" xml:space=\"preserve\">\n            <path fill=\"#000\" d=\"M43.935,25.145c0-10.318-8.364-18.683-18.683-18.683c-10.318,0-18.683,8.365-18.683,18.683h4.068c0-8.071,6.543-14.615,14.615-14.615c8.072,0,14.615,6.543,14.615,14.615H43.935z\">\n              <animateTransform attributeType=\"xml\" attributeName=\"transform\" type=\"rotate\" from=\"0 25 25\" to=\"360 25 25\" dur=\"0.6s\" repeatCount=\"indefinite\"/>\n            </path>\n          </svg>\n        </div>\n      </div>\n    </div>\n\n    <script src=\"/lnme/lnme.js\" id=\"lnme-script\"></script>\n    <script>\n      document.getElementById(\"get-new-address\").addEventListener('click', function(e) {\n        e.preventDefault();\n        var lnme = new LnMe({});\n        lnme.newAddress().then(address => {\n          document.getElementById(\"onchain\").innerHTML = address;\n        });\n      });\n\n      var urlParams = new URLSearchParams(window.location.search);\n      if (urlParams.get('amount')) {\n        document.getElementById('amount').value = urlParams.get('amount');\n      }\n      if (urlParams.get('memo')) {\n        document.getElementById('memo').value = urlParams.get('memo');\n      }\n      document.getElementById('send-button').addEventListener('click', function(e) {\n        e.preventDefault();\n        document.getElementById('loader').style.display = 'block';\n\n        var amountElement = document.getElementById('amount');\n        var memoElement = document.getElementById('memo');\n\n        e.target.setAttribute('disabled', true);\n        amountElement.setAttribute('disabled', true);\n        memoElement.setAttribute('disabled', true);\n\n        var lnme = new LnMe({ value: amountElement.value, memo: memoElement.value, target: document.getElementById('lnme-wrapper') });\n\n        lnme.showPaymentRequest = function() {\n          document.getElementById('send-button').style.display = 'none';\n          document.getElementById('lnme-wrapper').style.display = 'block';\n          this.populatePaymentRequest();\n        }\n        lnme.thanks = function() {\n          document.getElementById('form').innerHTML = '<h1 class=\"lnme-headline\">Thanks!</h1>';\n        }\n        lnme.request();\n      });\n    </script>\n  </body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1603402565, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`files/root`, &embedded.EmbeddedBox{
		Name: `files/root`,
		Time: time.Unix(1603402565, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
		},
	})
}

func init() {

	// define files
	file4 := &embedded.EmbeddedFile{
		Filename:    "lnme.css",
		FileModTime: time.Unix(1602933916, 0),

		Content: string("@keyframes slideIn{0%{opacity:0;transform:scale(.85)}70%{opacity:1;transform:scale(1.03)}100%{transform:scale(1)}}@keyframes slideOut{from{opacity:1}to{opacity:0}}@keyframes rotate{from{transform:rotate(0)}to{transform:rotate(180deg)}}.jPopup{position:absolute;top:0;right:0;bottom:0;left:0;z-index:9999;max-width:100%;padding:50px 15px 15px;box-sizing:border-box}.jPopup:after{content:'';position:fixed;top:0;right:0;bottom:0;left:0;z-index:9998;background:#fff}.jPopup>.jCloseBtn{position:absolute;right:1rem;top:1rem;z-index:9999;outline:0;border:0;box-sizing:border-box;cursor:pointer;width:5rem;height:5rem;background:#f2f2f2;border-radius:50%}.jPopup>.jCloseBtn>.graphicIcon{width:2rem;height:2rem;position:relative;margin:0 auto}.jPopup>.jCloseBtn>.graphicIcon:after,.jPopup>.jCloseBtn>.graphicIcon:before{position:absolute;left:.9rem;content:'';height:2rem;width:.3rem;background-color:#adadad;border-radius:.5rem}.jPopup>.jCloseBtn>.graphicIcon:before{-ms-transform:rotate(45deg);transform:rotate(45deg)}.jPopup>.jCloseBtn>.graphicIcon:after{-ms-transform:rotate(-45deg);transform:rotate(-45deg)}.jPopup>.jCloseBtn:hover>.graphicIcon{animation:rotate 250ms ease-in}.jPopup>.jCloseBtn:active{-ms-transform:scale(.95);transform:scale(.95)}.jPopup>.content{top:50%;left:1.5rem;right:1.5rem;position:absolute;z-index:9999;-ms-transform:translateY(-50%);transform:translateY(-50%)}@media screen and (min-width:680px){.jPopup{padding:8rem 3rem 3rem}.jPopup>.jCloseBtn{width:6rem;height:6rem;right:2rem;top:2rem}.jPopup>.jCloseBtn:after{content:'(esc)';position:absolute;top:6.5rem;left:50%;font-size:1.1rem;-ms-transform:translateX(-50%);transform:translateX(-50%);color:#adadad;pointer-events:none}.jPopup>.jCloseBtn>.graphicIcon{width:3rem;height:3rem}.jPopup>.jCloseBtn>.graphicIcon:after,.jPopup>.jCloseBtn>.graphicIcon:before{left:1.4rem;height:3rem}.jPopup>.content{left:3rem;right:3rem}}.jPopupOpen,.jPopupOpen body{overflow:hidden}.jPopupOpen .jPopup{animation:slideIn .5s cubic-bezier(.34,.34,.26,.99)}.jPopupClosed .jPopup{animation:slideOut 250ms ease-in}\n\n.jPopup > .jCloseBtn {\n  z-index:10000;\n}\n\n.jPopup .lnme-payment-request {\n  width: 220px;\n  margin: 0 auto;\n  text-align: center;\n}\n.jPopup .lnme-payment-request h1, .jPopup .lnme-payment-request h2 {\n  text-align: center;\n}\n\n.jPopup .lnme-details {\n  width: 200px;\n  overflow: hidden;\n}\n\n.jPopup .lnme-copy {\n  cursor: pointer;\n  display: inline;\n}\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "lnme.js",
		FileModTime: time.Unix(1603529292, 0),

		Content: string("'use strict';\n/*! https://github.com/robiveli/jpopup */\n!function(e,n){void 0===e&&void 0!==window&&(e=window),\"function\"==typeof define&&define.amd?define([],function(){return e.jPopup=n()}):\"object\"==typeof module&&module.exports?module.exports=n():e.jPopup=n()}(this,function(){\"use strict\";var n,o,t=function(){var e=0<arguments.length&&void 0!==arguments[0]?arguments[0]:\"\";1==(n=0!=e.shouldSetHash)&&(o=void 0!==e.hashtagValue?e.hashtagValue:\"#popup\"),i(e.content).then(a).then(1==n&&s(!0))},i=function(e){return u.classList.add(\"jPopupOpen\"),Promise.resolve(document.body.insertAdjacentHTML(\"beforeend\",'<div class=\"jPopup\">\\n                <button type=\"button\" class=\"jCloseBtn\">\\n                    <div class=\"graphicIcon\"></div>\\n                </button>\\n                <div class=\"content\">'.concat(e,\"</div>\\n            </div>\")))},s=function(e){1==e?window.location.hash=o:window.history.back()},d=function(e){27==e.keyCode&&t.prototype.close(!0)},c=function(){window.location.hash!==o&&t.prototype.close(!1)},a=function(){document.getElementsByClassName(\"jCloseBtn\")[0].addEventListener(\"click\",function(){t.prototype.close(!0)}),window.addEventListener(\"keydown\",d),1==n&&window.addEventListener(\"hashchange\",c)},u=document.querySelector(\"html\");return t.prototype={close:function(e){u.classList.add(\"jPopupClosed\"),1==n&&(e&&s(!1),window.removeEventListener(\"hashchange\",c)),window.removeEventListener(\"keydown\",d),document.getElementsByClassName(\"jPopup\")[0].addEventListener(\"animationend\",function(e){e.target.parentNode.removeChild(this),u.classList.remove(\"jPopupClosed\"),u.classList.remove(\"jPopupOpen\")})},open:function(e){t(e)}},t});\n\n/*\n https://github.com/nimiq/qr-creator\n jquery-qrcode v0.14.0 - https://larsjung.de/jquery-qrcode/ */\n'use strict';let G=null;class H{}H.render=function(w,B){G(w,B)};self.QrCreator=H;\n(function(w){function B(t,c,a,e){var b={},h=w(a,c);h.u(t);h.J();e=e||0;var r=h.h(),d=h.h()+2*e;b.text=t;b.level=c;b.version=a;b.O=d;b.a=function(b,a){b-=e;a-=e;return 0>b||b>=r||0>a||a>=r?!1:h.a(b,a)};return b}function C(t,c,a,e,b,h,r,d,g,x){function u(b,a,f,c,d,r,g){b?(t.lineTo(a+r,f+g),t.arcTo(a,f,c,d,h)):t.lineTo(a,f)}r?t.moveTo(c+h,a):t.moveTo(c,a);u(d,e,a,e,b,-h,0);u(g,e,b,c,b,0,-h);u(x,c,b,c,a,h,0);u(r,c,a,e,a,0,h)}function z(t,c,a,e,b,h,r,d,g,x){function u(b,a,c,d){t.moveTo(b+c,a);t.lineTo(b,\na);t.lineTo(b,a+d);t.arcTo(b,a,b+c,a,h)}r&&u(c,a,h,h);d&&u(e,a,-h,h);g&&u(e,b,-h,-h);x&&u(c,b,h,-h)}function A(t,c){var a=c.fill;if(\"string\"===typeof a)t.fillStyle=a;else{var e=a.type,b=a.colorStops;a=a.position.map((b)=>Math.round(b*c.size));if(\"linear-gradient\"===e)var h=t.createLinearGradient.apply(t,a);else if(\"radial-gradient\"===e)h=t.createRadialGradient.apply(t,a);else throw Error(\"Unsupported fill\");b.forEach(([b,a])=>{h.addColorStop(b,a)});t.fillStyle=h}}function y(t,c){a:{var a=c.text,e=\nc.v,b=c.N,h=c.K,r=c.P;b=Math.max(1,b||1);for(h=Math.min(40,h||40);b<=h;b+=1)try{var d=B(a,e,b,r);break a}catch(J){}d=void 0}if(!d)return null;a=t.getContext(\"2d\");c.background&&(a.fillStyle=c.background,a.fillRect(c.left,c.top,c.size,c.size));e=d.O;h=c.size/e;a.beginPath();for(r=0;r<e;r+=1)for(b=0;b<e;b+=1){var g=a,x=c.left+b*h,u=c.top+r*h,p=r,q=b,f=d.a,k=x+h,m=u+h,D=p-1,E=p+1,n=q-1,l=q+1,y=Math.floor(Math.min(.5,Math.max(0,c.R))*h),v=f(p,q),I=f(D,n),w=f(D,q);D=f(D,l);var F=f(p,l);l=f(E,l);q=f(E,\nq);E=f(E,n);p=f(p,n);x=Math.round(x);u=Math.round(u);k=Math.round(k);m=Math.round(m);v?C(g,x,u,k,m,y,!w&&!p,!w&&!F,!q&&!F,!q&&!p):z(g,x,u,k,m,y,w&&p&&I,w&&F&&D,q&&F&&l,q&&p&&E)}A(a,c);a.fill();return t}var v={minVersion:1,maxVersion:40,ecLevel:\"L\",left:0,top:0,size:200,fill:\"#000\",background:null,text:\"no text\",radius:.5,quiet:0};G=function(t,c){var a={};Object.assign(a,v,t);a.N=a.minVersion;a.K=a.maxVersion;a.v=a.ecLevel;a.left=a.left;a.top=a.top;a.size=a.size;a.fill=a.fill;a.background=a.background;\na.text=a.text;a.R=a.radius;a.P=a.quiet;if(c instanceof HTMLCanvasElement){if(c.width!==a.size||c.height!==a.size)c.width=a.size,c.height=a.size;c.getContext(\"2d\").clearRect(0,0,c.width,c.height);y(c,a)}else t=document.createElement(\"canvas\"),t.width=a.size,t.height=a.size,a=y(t,a),c.appendChild(a)}})(function(){function w(c){var a=C.s(c);return{S:function(){return 4},b:function(){return a.length},write:function(c){for(var b=0;b<a.length;b+=1)c.put(a[b],8)}}}function B(){var c=[],a=0,e={B:function(){return c},\nc:function(b){return 1==(c[Math.floor(b/8)]>>>7-b%8&1)},put:function(b,h){for(var a=0;a<h;a+=1)e.m(1==(b>>>h-a-1&1))},f:function(){return a},m:function(b){var h=Math.floor(a/8);c.length<=h&&c.push(0);b&&(c[h]|=128>>>a%8);a+=1}};return e}function C(c,a){function e(b,h){for(var a=-1;7>=a;a+=1)if(!(-1>=b+a||d<=b+a))for(var c=-1;7>=c;c+=1)-1>=h+c||d<=h+c||(r[b+a][h+c]=0<=a&&6>=a&&(0==c||6==c)||0<=c&&6>=c&&(0==a||6==a)||2<=a&&4>=a&&2<=c&&4>=c?!0:!1)}function b(b,a){for(var f=d=4*c+17,k=Array(f),m=0;m<\nf;m+=1){k[m]=Array(f);for(var p=0;p<f;p+=1)k[m][p]=null}r=k;e(0,0);e(d-7,0);e(0,d-7);f=y.G(c);for(k=0;k<f.length;k+=1)for(m=0;m<f.length;m+=1){p=f[k];var q=f[m];if(null==r[p][q])for(var n=-2;2>=n;n+=1)for(var l=-2;2>=l;l+=1)r[p+n][q+l]=-2==n||2==n||-2==l||2==l||0==n&&0==l}for(f=8;f<d-8;f+=1)null==r[f][6]&&(r[f][6]=0==f%2);for(f=8;f<d-8;f+=1)null==r[6][f]&&(r[6][f]=0==f%2);f=y.w(h<<3|a);for(k=0;15>k;k+=1)m=!b&&1==(f>>k&1),r[6>k?k:8>k?k+1:d-15+k][8]=m,r[8][8>k?d-k-1:9>k?15-k:14-k]=m;r[d-8][8]=!b;if(7<=\nc){f=y.A(c);for(k=0;18>k;k+=1)m=!b&&1==(f>>k&1),r[Math.floor(k/3)][k%3+d-8-3]=m;for(k=0;18>k;k+=1)m=!b&&1==(f>>k&1),r[k%3+d-8-3][Math.floor(k/3)]=m}if(null==g){b=t.I(c,h);f=B();for(k=0;k<x.length;k+=1)m=x[k],f.put(4,4),f.put(m.b(),y.f(4,c)),m.write(f);for(k=m=0;k<b.length;k+=1)m+=b[k].j;if(f.f()>8*m)throw Error(\"code length overflow. (\"+f.f()+\">\"+8*m+\")\");for(f.f()+4<=8*m&&f.put(0,4);0!=f.f()%8;)f.m(!1);for(;!(f.f()>=8*m);){f.put(236,8);if(f.f()>=8*m)break;f.put(17,8)}var u=0;m=k=0;p=Array(b.length);\nq=Array(b.length);for(n=0;n<b.length;n+=1){var v=b[n].j,w=b[n].o-v;k=Math.max(k,v);m=Math.max(m,w);p[n]=Array(v);for(l=0;l<p[n].length;l+=1)p[n][l]=255&f.B()[l+u];u+=v;l=y.C(w);v=z(p[n],l.b()-1).l(l);q[n]=Array(l.b()-1);for(l=0;l<q[n].length;l+=1)w=l+v.b()-q[n].length,q[n][l]=0<=w?v.c(w):0}for(l=f=0;l<b.length;l+=1)f+=b[l].o;f=Array(f);for(l=u=0;l<k;l+=1)for(n=0;n<b.length;n+=1)l<p[n].length&&(f[u]=p[n][l],u+=1);for(l=0;l<m;l+=1)for(n=0;n<b.length;n+=1)l<q[n].length&&(f[u]=q[n][l],u+=1);g=f}b=g;f=\n-1;k=d-1;m=7;p=0;a=y.F(a);for(q=d-1;0<q;q-=2)for(6==q&&--q;;){for(n=0;2>n;n+=1)null==r[k][q-n]&&(l=!1,p<b.length&&(l=1==(b[p]>>>m&1)),a(k,q-n)&&(l=!l),r[k][q-n]=l,--m,-1==m&&(p+=1,m=7));k+=f;if(0>k||d<=k){k-=f;f=-f;break}}}var h=A[a],r=null,d=0,g=null,x=[],u={u:function(b){b=w(b);x.push(b);g=null},a:function(b,a){if(0>b||d<=b||0>a||d<=a)throw Error(b+\",\"+a);return r[b][a]},h:function(){return d},J:function(){for(var a=0,h=0,c=0;8>c;c+=1){b(!0,c);var d=y.D(u);if(0==c||a>d)a=d,h=c}b(!1,h)}};return u}\nfunction z(c,a){if(\"undefined\"==typeof c.length)throw Error(c.length+\"/\"+a);var e=function(){for(var b=0;b<c.length&&0==c[b];)b+=1;for(var r=Array(c.length-b+a),d=0;d<c.length-b;d+=1)r[d]=c[d+b];return r}(),b={c:function(b){return e[b]},b:function(){return e.length},multiply:function(a){for(var h=Array(b.b()+a.b()-1),c=0;c<b.b();c+=1)for(var g=0;g<a.b();g+=1)h[c+g]^=v.i(v.g(b.c(c))+v.g(a.c(g)));return z(h,0)},l:function(a){if(0>b.b()-a.b())return b;for(var c=v.g(b.c(0))-v.g(a.c(0)),h=Array(b.b()),\ng=0;g<b.b();g+=1)h[g]=b.c(g);for(g=0;g<a.b();g+=1)h[g]^=v.i(v.g(a.c(g))+c);return z(h,0).l(a)}};return b}C.s=function(c){for(var a=[],e=0;e<c.length;e++){var b=c.charCodeAt(e);128>b?a.push(b):2048>b?a.push(192|b>>6,128|b&63):55296>b||57344<=b?a.push(224|b>>12,128|b>>6&63,128|b&63):(e++,b=65536+((b&1023)<<10|c.charCodeAt(e)&1023),a.push(240|b>>18,128|b>>12&63,128|b>>6&63,128|b&63))}return a};var A={L:1,M:0,Q:3,H:2},y=function(){function c(b){for(var a=0;0!=b;)a+=1,b>>>=1;return a}var a=[[],[6,18],\n[6,22],[6,26],[6,30],[6,34],[6,22,38],[6,24,42],[6,26,46],[6,28,50],[6,30,54],[6,32,58],[6,34,62],[6,26,46,66],[6,26,48,70],[6,26,50,74],[6,30,54,78],[6,30,56,82],[6,30,58,86],[6,34,62,90],[6,28,50,72,94],[6,26,50,74,98],[6,30,54,78,102],[6,28,54,80,106],[6,32,58,84,110],[6,30,58,86,114],[6,34,62,90,118],[6,26,50,74,98,122],[6,30,54,78,102,126],[6,26,52,78,104,130],[6,30,56,82,108,134],[6,34,60,86,112,138],[6,30,58,86,114,142],[6,34,62,90,118,146],[6,30,54,78,102,126,150],[6,24,50,76,102,128,154],\n[6,28,54,80,106,132,158],[6,32,58,84,110,136,162],[6,26,54,82,110,138,166],[6,30,58,86,114,142,170]],e={w:function(b){for(var a=b<<10;0<=c(a)-c(1335);)a^=1335<<c(a)-c(1335);return(b<<10|a)^21522},A:function(b){for(var a=b<<12;0<=c(a)-c(7973);)a^=7973<<c(a)-c(7973);return b<<12|a},G:function(b){return a[b-1]},F:function(b){switch(b){case 0:return function(b,a){return 0==(b+a)%2};case 1:return function(b){return 0==b%2};case 2:return function(b,a){return 0==a%3};case 3:return function(b,a){return 0==\n(b+a)%3};case 4:return function(b,a){return 0==(Math.floor(b/2)+Math.floor(a/3))%2};case 5:return function(b,a){return 0==b*a%2+b*a%3};case 6:return function(b,a){return 0==(b*a%2+b*a%3)%2};case 7:return function(b,a){return 0==(b*a%3+(b+a)%2)%2};default:throw Error(\"bad maskPattern:\"+b);}},C:function(b){for(var a=z([1],0),c=0;c<b;c+=1)a=a.multiply(z([1,v.i(c)],0));return a},f:function(b,a){if(4!=b||1>a||40<a)throw Error(\"mode: \"+b+\"; type: \"+a);return 10>a?8:16},D:function(b){for(var a=b.h(),c=0,\nd=0;d<a;d+=1)for(var g=0;g<a;g+=1){for(var e=0,t=b.a(d,g),p=-1;1>=p;p+=1)if(!(0>d+p||a<=d+p))for(var q=-1;1>=q;q+=1)0>g+q||a<=g+q||(0!=p||0!=q)&&t==b.a(d+p,g+q)&&(e+=1);5<e&&(c+=3+e-5)}for(d=0;d<a-1;d+=1)for(g=0;g<a-1;g+=1)if(e=0,b.a(d,g)&&(e+=1),b.a(d+1,g)&&(e+=1),b.a(d,g+1)&&(e+=1),b.a(d+1,g+1)&&(e+=1),0==e||4==e)c+=3;for(d=0;d<a;d+=1)for(g=0;g<a-6;g+=1)b.a(d,g)&&!b.a(d,g+1)&&b.a(d,g+2)&&b.a(d,g+3)&&b.a(d,g+4)&&!b.a(d,g+5)&&b.a(d,g+6)&&(c+=40);for(g=0;g<a;g+=1)for(d=0;d<a-6;d+=1)b.a(d,g)&&!b.a(d+\n1,g)&&b.a(d+2,g)&&b.a(d+3,g)&&b.a(d+4,g)&&!b.a(d+5,g)&&b.a(d+6,g)&&(c+=40);for(g=e=0;g<a;g+=1)for(d=0;d<a;d+=1)b.a(d,g)&&(e+=1);return c+=Math.abs(100*e/a/a-50)/5*10}};return e}(),v=function(){for(var c=Array(256),a=Array(256),e=0;8>e;e+=1)c[e]=1<<e;for(e=8;256>e;e+=1)c[e]=c[e-4]^c[e-5]^c[e-6]^c[e-8];for(e=0;255>e;e+=1)a[c[e]]=e;return{g:function(b){if(1>b)throw Error(\"glog(\"+b+\")\");return a[b]},i:function(b){for(;0>b;)b+=255;for(;256<=b;)b-=255;return c[b]}}}(),t=function(){function c(b,c){switch(c){case A.L:return a[4*\n(b-1)];case A.M:return a[4*(b-1)+1];case A.Q:return a[4*(b-1)+2];case A.H:return a[4*(b-1)+3]}}var a=[[1,26,19],[1,26,16],[1,26,13],[1,26,9],[1,44,34],[1,44,28],[1,44,22],[1,44,16],[1,70,55],[1,70,44],[2,35,17],[2,35,13],[1,100,80],[2,50,32],[2,50,24],[4,25,9],[1,134,108],[2,67,43],[2,33,15,2,34,16],[2,33,11,2,34,12],[2,86,68],[4,43,27],[4,43,19],[4,43,15],[2,98,78],[4,49,31],[2,32,14,4,33,15],[4,39,13,1,40,14],[2,121,97],[2,60,38,2,61,39],[4,40,18,2,41,19],[4,40,14,2,41,15],[2,146,116],[3,58,36,\n2,59,37],[4,36,16,4,37,17],[4,36,12,4,37,13],[2,86,68,2,87,69],[4,69,43,1,70,44],[6,43,19,2,44,20],[6,43,15,2,44,16],[4,101,81],[1,80,50,4,81,51],[4,50,22,4,51,23],[3,36,12,8,37,13],[2,116,92,2,117,93],[6,58,36,2,59,37],[4,46,20,6,47,21],[7,42,14,4,43,15],[4,133,107],[8,59,37,1,60,38],[8,44,20,4,45,21],[12,33,11,4,34,12],[3,145,115,1,146,116],[4,64,40,5,65,41],[11,36,16,5,37,17],[11,36,12,5,37,13],[5,109,87,1,110,88],[5,65,41,5,66,42],[5,54,24,7,55,25],[11,36,12,7,37,13],[5,122,98,1,123,99],[7,73,\n45,3,74,46],[15,43,19,2,44,20],[3,45,15,13,46,16],[1,135,107,5,136,108],[10,74,46,1,75,47],[1,50,22,15,51,23],[2,42,14,17,43,15],[5,150,120,1,151,121],[9,69,43,4,70,44],[17,50,22,1,51,23],[2,42,14,19,43,15],[3,141,113,4,142,114],[3,70,44,11,71,45],[17,47,21,4,48,22],[9,39,13,16,40,14],[3,135,107,5,136,108],[3,67,41,13,68,42],[15,54,24,5,55,25],[15,43,15,10,44,16],[4,144,116,4,145,117],[17,68,42],[17,50,22,6,51,23],[19,46,16,6,47,17],[2,139,111,7,140,112],[17,74,46],[7,54,24,16,55,25],[34,37,13],[4,\n151,121,5,152,122],[4,75,47,14,76,48],[11,54,24,14,55,25],[16,45,15,14,46,16],[6,147,117,4,148,118],[6,73,45,14,74,46],[11,54,24,16,55,25],[30,46,16,2,47,17],[8,132,106,4,133,107],[8,75,47,13,76,48],[7,54,24,22,55,25],[22,45,15,13,46,16],[10,142,114,2,143,115],[19,74,46,4,75,47],[28,50,22,6,51,23],[33,46,16,4,47,17],[8,152,122,4,153,123],[22,73,45,3,74,46],[8,53,23,26,54,24],[12,45,15,28,46,16],[3,147,117,10,148,118],[3,73,45,23,74,46],[4,54,24,31,55,25],[11,45,15,31,46,16],[7,146,116,7,147,117],\n[21,73,45,7,74,46],[1,53,23,37,54,24],[19,45,15,26,46,16],[5,145,115,10,146,116],[19,75,47,10,76,48],[15,54,24,25,55,25],[23,45,15,25,46,16],[13,145,115,3,146,116],[2,74,46,29,75,47],[42,54,24,1,55,25],[23,45,15,28,46,16],[17,145,115],[10,74,46,23,75,47],[10,54,24,35,55,25],[19,45,15,35,46,16],[17,145,115,1,146,116],[14,74,46,21,75,47],[29,54,24,19,55,25],[11,45,15,46,46,16],[13,145,115,6,146,116],[14,74,46,23,75,47],[44,54,24,7,55,25],[59,46,16,1,47,17],[12,151,121,7,152,122],[12,75,47,26,76,48],\n[39,54,24,14,55,25],[22,45,15,41,46,16],[6,151,121,14,152,122],[6,75,47,34,76,48],[46,54,24,10,55,25],[2,45,15,64,46,16],[17,152,122,4,153,123],[29,74,46,14,75,47],[49,54,24,10,55,25],[24,45,15,46,46,16],[4,152,122,18,153,123],[13,74,46,32,75,47],[48,54,24,14,55,25],[42,45,15,32,46,16],[20,147,117,4,148,118],[40,75,47,7,76,48],[43,54,24,22,55,25],[10,45,15,67,46,16],[19,148,118,6,149,119],[18,75,47,31,76,48],[34,54,24,34,55,25],[20,45,15,61,46,16]],e={I:function(b,a){var e=c(b,a);if(\"undefined\"==\ntypeof e)throw Error(\"bad rs block @ typeNumber:\"+b+\"/errorCorrectLevel:\"+a);b=e.length/3;a=[];for(var d=0;d<b;d+=1)for(var g=e[3*d],h=e[3*d+1],t=e[3*d+2],p=0;p<g;p+=1){var q=t,f={};f.o=h;f.j=q;a.push(f)}return a}};return e}();return C}());\n\n// no static class variables because those are not supported on mobile\nvar paymentRequestTemplate = `<div id=\"lnme-wrapper\" class=\"lnme-wrapper\">\n    <h1 class=\"lnme-headline\"><span class=\"lnme-memo\"><span></h1>\n    <h2 class=\"lnme-headline\"><span class=\"lnme-value\"></span> Sats</h2>\n    <div class=\"lnme-qrcode\"></div>\n    <div class=\"lnme-details\">\n      <a href=\"#\" class=\"lnme-link\"><span class=\"lnme-payment-request\"></span></a>\n      <div class=\"lnme-copy\">\n        <svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"feather feather-copy\"><rect x=\"9\" y=\"9\" width=\"13\" height=\"13\" rx=\"2\" ry=\"2\"></rect><path d=\"M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1\"></path></svg>\n      </div>\n    </div>\n  </div>`;\nvar paymentConfirmationTemplate = `<h1 class=\"lnme-header lnme-confirmation\">Payment sent!</h1>`;\n\nclass LnMe {\n\n  constructor(options) {\n    options = options || {};\n    this.script = document.getElementById('lnme-script');\n    if (options.baseURL) {\n      this.baseURL = options.baseURL;\n    } else if (this.script && this.script.dataset.lnmeBaseUrl) {\n      this.baseURL = this.script.dataset.lnmeBaseUrl;\n    } else {\n      this.baseURL = `${document.location.protocol}//${document.location.host}`;\n    }\n    this.value = parseInt(options.value || 0);\n    this.memo = options.memo || '';\n    this.target = options.target;\n    this.loadStylesheet(); // load it early that styles are ready when the popup is opened\n  }\n\n  loadStylesheet() {\n    if (document.getElementById('lnme-style')) { return; }\n    // get the CSS file from the same source as the JS widget file\n    let source = `${this.baseURL}/lnme/lnme.css`;\n    let head = document.getElementsByTagName('head')[0];\n    let css = document.createElement('link');\n    css.id = \"lnme-style\";\n    css.rel = \"stylesheet\";\n    css.type = \"text/css\";\n    css.href = source\n    head.appendChild(css);\n  }\n\n  watchPayment() {\n    if (this.paymentWatcher) { window.clearInterval(this.paymentWatcher) }\n\n    return new Promise((resolve, reject) => {\n      this.paymentWatcher = window.setInterval(() => {\n        this._fetch(`${this.baseURL}/v1/invoice/${this.invoice.payment_hash}`)\n          .then((invoice) => {\n            if (invoice.settled) {\n              this.invoice.settled = true;\n              this.stopWatchingPayment();\n              resolve(this.invoice);\n            }\n          });\n      }, 2000);\n    });\n  }\n\n  stopWatchingPayment() {\n    window.clearInterval(this.paymentWatcher);\n    this.paymentWatcher = null;\n  }\n\n  payWithWebln() {\n    if (!webln.isEnabled) {\n      webln.enable().then((weblnResponse) => {\n        return webln.sendPayment(this.invoice.payment_request);\n      }).catch((e) => {\n        return this.showPaymentRequest();\n      })\n    } else {\n      return webln.sendPayment(this.invoice.payment_request);\n    }\n  }\n\n  populatePaymentRequest() {\n    document.querySelectorAll('.lnme-memo').forEach(e => {\n      e.innerHTML = this.memo;\n    });\n    document.querySelectorAll('.lnme-value').forEach(e => {\n      e.innerHTML = this.value;\n    });\n    document.querySelectorAll('.lnme-payment-request').forEach(e => {\n      e.innerHTML = this.invoice.payment_request;\n    });\n    document.querySelectorAll('.lnme-link').forEach(e => {\n      e.setAttribute('href', `lightning:${this.invoice.payment_request}`);\n    });\n    QrCreator.render({ text: this.invoice.payment_request, size: 128}, this.target.querySelector('.lnme-qrcode'));\n    this.target.querySelectorAll('.lnme-copy').forEach(element => {\n      element.addEventListener('click', (e) => {\n        navigator.clipboard.writeText(this.invoice.payment_request).then(() => {\n          alert('Copied to clipboad');\n        });\n      });\n    });\n  }\n\n  showPaymentRequest() {\n    this.render(LnMePaymentRequestTemplate);\n    this.populatePaymentRequest()\n    return Promise.resolve(); // be compatible to payWithWebln()\n  }\n\n  addInvoice() {\n    let args = {\n      method: 'POST',\n      mode: 'cors',\n      headers: { 'Content-Type': 'application/json' },\n      body: JSON.stringify({ memo: this.memo, value: this.value })\n    };\n    return this._fetch(\n      `${this.baseURL}/v1/invoices`,\n      args\n      ).then((invoice) => {\n        this.invoice = invoice;\n        return invoice;\n      });\n  }\n\n  newAddress() {\n    let args = {\n      method: 'POST',\n      mode: 'cors',\n      header: { 'Content-Type': 'application/json' }\n    };\n    return this._fetch(`${this.baseURL}/v1/newaddress`, args)\n      .then(address => {\n        this.address = address;\n        return address;\n      });\n  }\n\n  requestPayment() {\n    return this.addInvoice().then((invoice) => {\n      if (typeof webln !== 'undefined') {\n        return this.payWithWebln();\n      } else {\n        return this.showPaymentRequest();\n      }\n    });\n  }\n\n  request() {\n    return this.requestPayment().then(() => {\n      this.watchPayment().then((invoice) => {\n        this.thanks();\n      });\n    });\n  }\n\n  _fetch(url, args) {\n    return fetch(url, args).then((response) => {\n      if (response.ok) {\n        return response.json();\n      } else {\n        throw new Error(response);\n      }\n    });\n  }\n\n  closePopup() {\n    if (this.popup) {\n      this.popup.close();\n      this.popup = null;\n    }\n  }\n\n  render(content) {\n    this.closePopup();\n    this.popup = new jPopup({\n      content: content,\n      shouldSetHash: false\n    });\n    this.target = document.querySelector('.jPopup .content').firstElementChild;\n  }\n\n  thanks() {\n    this.target.innerHTML = LnMePaymentConfirmationTemplate;\n  }\n}\n"),
	}

	// define dirs
	dir3 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1603529292, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // "lnme.css"
			file5, // "lnme.js"

		},
	}

	// link ChildDirs
	dir3.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`files/assets`, &embedded.EmbeddedBox{
		Name: `files/assets`,
		Time: time.Unix(1603529292, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir3,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"lnme.css": file4,
			"lnme.js":  file5,
		},
	})
}