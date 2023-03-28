(function($){
	var hoverPromo = (function(){
		$(document).on('mouseenter mouseleave click', '#tt-pageContent .js-handler', function(e) {
			var $this = $(this),
				objHeight = $this.height();

			if (e.type === 'mouseenter') {
				onHover();
				} else if (e.type === 'mouseleave' && e.relatedTarget){
				offHover();
			};

			function onHover(e){
				$this.addClass('active');
				$this.height(objHeight);
				return false;
			};
			function offHover(e){
				$this.removeClass('active');
				$this.removeAttr('style');
				return false;
			};
		});
	}());
})(jQuery);
