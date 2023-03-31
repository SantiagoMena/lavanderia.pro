(function($){
	var $html = $('html');

	var cssFix = function(){
		var u = navigator.userAgent.toLowerCase(),
			is = function(t) {
				return (u.indexOf(t) != -1)
			};
		$html.addClass([
			(!(/opera|webtv/i.test(u)) && /msie (\d)/.test(u)) ? ('ie ie' + RegExp.$1) :
			is('firefox/2') ? 'gecko ff2' :
			is('firefox/3') ? 'gecko ff3' :
			is('gecko/') ? 'gecko' :
			is('opera/9') ? 'opera opera9' : /opera (\d)/.test(u) ? 'opera opera' + RegExp.$1 :
			is('konqueror') ? 'konqueror' :
			is('applewebkit/') ? 'webkit safari' :
			is('mozilla/') ? 'gecko' : '',
			(is('x11') || is('linux')) ? ' linux' :
			is('mac') ? ' mac' :
			is('win') ? ' win' : ''
		].join(''));
	}();
	// identify touch device
	function is_touch_device(){
		return !!('ontouchstart' in window) || !!('onmsgesturechange' in window);
	};
	if(is_touch_device()){
		$html.addClass('touch-device');
	};
	if (/Edge/.test(navigator.userAgent)){
		$html.addClass('edge');
	};
	var isiPad = navigator.userAgent.match(/iPad/i) != null;
	if (isiPad){
		$html.addClass('ipad');
	}
})(jQuery);
