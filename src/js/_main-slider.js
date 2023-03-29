(function($){

	$(window).on('load', function () {
		$('body').addClass('btn-animation');
	});

	var sliderWrapper = $("#js-mainSlider"),
		srcAjaxDesktop = "ajax-content/mainslider-desktop.html",
		srcAjaxMobile = "ajax-content/mainslider-mobile.html";


	function debouncer(func, timeout) {
		var timeoutID, timeout = timeout || 500;
		return function() {
			var scope = this,
				args = arguments;
			clearTimeout(timeoutID);
			timeoutID = setTimeout(function() {
				func.apply(scope, Array.prototype.slice.call(args));
			}, timeout);
		}
	};

	function includeLayout(){
		sliderWrapper.find('.mainSlider-wrapper').empty();
		if($(window).width() > 767){
			ajaxInclude(srcAjaxDesktop);
		} else{
			ajaxInclude(srcAjaxMobile);
		};
	};
	function ajaxInclude(value){
		$.ajax({
			url: value,
			success: function(data) {
				var $item = $(data);
				sliderWrapper.find('.mainSlider-wrapper').append($item);
				initMainSlider();
			}
		});
	};
	includeLayout();
	$(window).resize(debouncer(function(e) {
		includeLayout()
	}));

	function initMainSlider(){
		var checkInit = sliderWrapper.find('.main-slider'),
			bubbles = $('#bubbles');

		if(checkInit.hasClass('slick-initialized')){
			checkInit.slick("unslick");
		};

		function initTilt(){
			if(window.innerWidth > 1024){
				sliderWrapper.find('.js-rotation').tilt({
					perspective: 1000
				});
			}
		};
		if (sliderWrapper.length){
			initTilt();
			$(window).resize(debouncer(function(e){
				initTilt();
			}));
		};
		sliderWrapper.find('.main-slider').on('init', function(event, slick, nextSlide){
			setTimeout(function () {
				sliderWrapper.addClass('show');
			}, 500);
			bubbles.addClass('start');
		});
		sliderWrapper.find('.main-slider').slick({
			slidesToShow: 1,
			slidesToScroll: 1,
			autoplaySpeed: 4300,
			speed: 1200,
			autoplay: true,
			arrows: false,
			dots: false,
			fade: true,
			responsive: [
				{
					breakpoint: 550,
					settings: {
						speed: 700,
					}
				}
			]
		});
		sliderWrapper.on('beforeChange', function(event, slick, nextSlide, currentSlide){
			$(this).removeClass('start');
			bubbles.removeClass('start');
		});
		sliderWrapper.on('afterChange', function(event, slick, nextSlide, currentSlide){
			$(this).addClass('start');
			bubbles.addClass('start');
		});
	};
})(jQuery);



