(function($){
	var nav = $('#tt-nav');

	if(!nav.length) return false;

	var determineSybMenu = (function(){
		nav.find('> ul > li').each(function(){
			if ($(this).children('ul').length != 0){
				$(this).addClass('subMenu');
			}
		});
	}());

	var determineActive = (function(){
		var location = window.location.href.split('#')[0],
			cur_url = location.split('/').pop() || 'index.html';

		nav.find('li').each(function(){
			var link = $(this).find('a').attr('href');
			if (cur_url == link){
				$(this).addClass('active').closest('.subMenu').addClass('active');
				nav.addClass('defined-item');
			}
		});
	}());

	var missingItemActive = (function(){
		if(!nav.hasClass('defined-item')){
			nav.find('> ul > li:first-child').addClass('active');
		}
	}());

	var hoverAddClass = (function(){
		nav.find('li').on( "mouseenter mouseleave", function( event ) {
			$(this).toggleClass('is-hover');
		});
	}());

})(jQuery);
