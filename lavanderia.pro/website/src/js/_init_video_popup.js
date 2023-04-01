(function($){
	var videoPopup = $('#tt-pageContent .js-video-popup');
	if(!videoPopup.length) return;
	videoPopup.each(function(){
		$(this).magnificPopup({
			type: 'iframe',
			iframe: {
			patterns: {
				dailymotion: {
				index: 'dailymotion.com',
				id: function(url) {
					var m = url.match(/^.+dailymotion.com\/(video|hub)\/([^_]+)[^#]*(#video=([^_&]+))?/);
					if (m !== null) {
						if(m[4] !== undefined){
							return m[4];
						}
						return m[2];
					}
					return null;
				},
				src: 'https://www.dailymotion.com/embed/video/%id%'
				}
			}
			}
		});
	});
})(jQuery);

