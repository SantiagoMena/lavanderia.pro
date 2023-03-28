(function($){
	$('[data-slick]').slick({
		lazyLoad: 'progressive',
		dots: true,
		arrows: false,
		infinite: true,
		speed: 300,
		autoplay:true,
		adaptiveHeight: true,
		slidesToScroll: 1,
		pauseOnFocus:false,
		pauseOnHover: false
	});
})(jQuery);
(function($){
	function initSliderCarusel(){
		var blogSlider = $('#blog-slider');

		if (!blogSlider.length) return false;


		var layout = blogSlider.find('.blog-slider__layout'),
			ImgWrapper = blogSlider.find('.tt-item-wrapper');

		layout.slick({
			dots: true,
			arrows: false,
			infinite: true,
			speed: 300,
			slidesToShow: 1,
			slidesToScroll: 1,
			adaptiveHeight: true,
			autoplay:true,
			autoplaySpeed:4500,
			pauseOnFocus:false,
			pauseOnHover: false
		});
		layout.on('afterChange', function(event, slick, currentSlide, nextSlide){
			var currentIndex = currentSlide;
			currentIndex++;
			ImgWrapper.find('.number-' + currentIndex).addClass('active').siblings().removeClass('active');
		});

		var numberImg = (function () {
			ImgWrapper.each(function(){
				$('.tt-item', this).each(function(i){
					$(this).addClass('number-' + (i+1));
				})
			})
		}());
	};
	initSliderCarusel();
})(jQuery);
(function($){
	function initSliderCarusel(){
		var slick04 = $('#tt-pageContent .js-init-carusel-tablet'),
			width = window.innerWidth || document.body.clientWidth;

		if (!slick04.length) return false;
		if (width <= 1024){
			slick04.slick({
				lazyLoad: 'progressive',
				dots: true,
				arrows: false,
				infinite: true,
				speed: 300,
				slidesToShow: 2,
				slidesToScroll: 2,
				adaptiveHeight: true,
				autoplay:true,
				autoplaySpeed:4500,
				pauseOnFocus:false,
				pauseOnHover: false,
				responsive: [
				{
					breakpoint: 650,
					settings: {
						slidesToShow: 1,
						slidesToScroll: 1
					}
				}
			  ]
			});
		} else {
			slick04.filter('.slick-initialized').slick('unslick');
		}
	};
	initSliderCarusel();
	$(window).resize(debouncer(function(e){
		initSliderCarusel();
	}));
})(jQuery);