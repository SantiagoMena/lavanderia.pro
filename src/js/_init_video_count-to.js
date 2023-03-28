(function($){
	var objCounter = $('#counter-js');
	if(!objCounter.length) return;

	var $window = $(window);

	$window.scroll(function(){
		var ttCounterObj =  $('.tt-counter');

		ttCounterObj.each(function(){
			var cPos = $(this).offset().top,
				topWindow = $window.scrollTop();

			if(cPos < topWindow + 800) {
				$(this).countTo().removeClass('tt-counter');
			}
		});
	});
})(jQuery);

