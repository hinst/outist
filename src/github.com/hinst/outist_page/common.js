/// <reference path="jquery-3.0.0.js"/>
// I took this function from somewhere
HinstApp.prototype.getURLArgument = function (param) {
	var vars = {};
	window.location.href.replace(location.hash, '').replace(
		/[?&]+([^=&]+)=?([^&]*)?/gi, // regexp
		function (m, key, value) { // callback
			vars[key] = value !== undefined ? value : '';
		}
	);
	if (param) {
		return vars[param] ? decodeURIComponent(vars[param]) : null;
	}
	return vars;
};

// I took this from https://learn.javascript.ru/cookie
HinstApp.prototype.getCookie = function (name) {
	var matches = document.cookie.match(new RegExp(
	  "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
	));
	return matches ? decodeURIComponent(matches[1]) : undefined;
}

HinstApp.prototype.setCookie = function (name, value, options) {
	options = options || {};

	var expires = options.expires;

	if (typeof expires == "number" && expires) {
		var d = new Date();
		d.setTime(d.getTime() + expires * 1000);
		expires = options.expires = d;
	}
	if (expires && expires.toUTCString) {
		options.expires = expires.toUTCString();
	}

	value = encodeURIComponent(value);

	var updatedCookie = name + "=" + value;

	for (var propName in options) {
		updatedCookie += "; " + propName;
		var propValue = options[propName];
		if (propValue !== true) {
			updatedCookie += "=" + propValue;
		}
	}

	document.cookie = updatedCookie;
};

HinstApp.prototype.addEvent = function(oldEvent, newEvent) {
	return function() {
		oldEvent();
		newEvent();
	}
}

HinstApp.prototype.goTo = function (url) {
	document.location = url;
}

hinstApp = new HinstApp();

