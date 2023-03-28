(function($){
	var objLayout = $('#js-stepbox__layout'),
		objNaw = $('#js-stepbox__nav'),
		lengthSlide = objNaw.find('li').length;

	if (!objLayout.length && !objNaw.length) return false;

	var switchNav = (function () {
		$('body').on('click','#js-stepbox__nav .stepbox-dots li', function(){
			if ($(this).hasClass('active')) return false;
			$(this).addClass('active').siblings().removeClass('active');
			var dataNumber = $(this).attr('data-number');
			switchSlides(dataNumber);
			objNaw.attr('data-number', dataNumber);
			return false;
		});
	}());
	function switchSlides(dataNumber){
		objLayout.find('[data-number="' + dataNumber + '"]').addClass('active').siblings().removeClass('active');
	};
	var autoSlide = setInterval(function () {
		var rollingSlide = objNaw.find('.stepbox-dots .active').next(),
			dataNumber = rollingSlide.attr('data-number') || false;

		if(!dataNumber){
			objNaw.find('[data-number="1"]').click();
		} else{
			objNaw.find('.stepbox-dots .active').next().click();
		}
	}, 4000);
})(jQuery);
