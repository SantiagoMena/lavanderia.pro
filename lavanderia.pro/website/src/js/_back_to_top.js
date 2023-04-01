(function($){
	var ptBackToTop =  $('#js-backtotop'),
		$window = $(window);

	if (ptBackToTop.length){
		initbacktotop();
	};
	function initbacktotop(){
		ptBackToTop.on('click',  function(e) {
			$('html, body').animate({
				scrollTop: 0
			}, 500);
			return false;
		});
		$window.scroll(function(){
			$window.scrollTop() > 500 ? ptBackToTop.stop(true.false).addClass('pt-show') : ptBackToTop.stop(true.false).removeClass('pt-show');
		});
	};
})(jQuery);
