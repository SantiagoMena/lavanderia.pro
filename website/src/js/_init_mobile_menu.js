(function($){
	var $ttDesctopMenu = $('#tt-nav'),
		mobileMenuToggle = $('#tt-menu-toggle');

	if ($ttDesctopMenu && mobileMenuToggle){
		var ttDesktopMenu = $ttDesctopMenu.find('ul').first().children().clone();
		$('#mobile-menu').find('ul').append(ttDesktopMenu);
		mobileMenuToggle.initMM({
			enable_breakpoint: true,
			mobile_button: true,
			breakpoint: 1360
		});
	};
})(jQuery);
