(function($){
	$('body').on('shown.bs.modal', function (e){
		if($('body').hasClass('mm-open')){
			$('#mobile-menu .mm-close').trigger("click");
		};
		return false;
	});
})(jQuery);


