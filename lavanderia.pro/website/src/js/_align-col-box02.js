
(function($){
	function calcHeight(){
		var ttwindowWidth = window.innerWidth || $(window).width();
		if(ttwindowWidth <= 650){
			$('#tt-pageContent .box02').each(function(){
				$(this).find('.box02__content').attr('style', '');
			});
			return false;
		};

		$(this).find('.box02__content').attr('style', '');
		$('#tt-pageContent .box02').each(function(){
			var value = $(this).find('.box02__content').innerHeight();

			if($(this).hasClass('box02-notover')){
				var value = $(this).find('.box02__content').innerHeight();
			} else{
				var value = $(this).find('.box02__content').innerHeight() + 80;
			};

			$(this).find('.box02__img img:not(.tt-arrow)').css({
				height: value
			});

		});
	};

	$(window).on('load', function(){
		calcHeight();
	});
	$(window).resize(debouncer(function(e){
		calcHeight()
	}));
})(jQuery);
