(function($){
	var methods = {
		init: function(options){
			return this.each(function(){
				var obj = $(this),
				objOpen = obj.find('.tt-item.tt-item__open'),
				objItemTitle = obj.find('.tt-item .tt-item__title');

				obj.addClass('init-accordeon');

				objOpen.find('.tt-item__content').slideDown(100);

				objItemTitle.on('click', function(){
					$(this).closest('.tt-item').siblings('.tt-item__open').find('.tt-item__content').slideToggle(200).closest('.tt-item').removeClass('tt-item__open');
					$(this).next().slideToggle(200).parent().toggleClass('tt-item__open');
				});
			});
		}
	};
	$.fn.accordeon = function(action){
		if(methods[action]){
			return methods[action].apply(this, Array.prototype.slice.call(arguments, 1));
		} else if(typeof action === 'object' || !action){
			return methods.init.apply(this, arguments);
		} else {
			console.info('Action ' +action+ 'not found this plugin');
			return this;
		}
	};
	$('#tt-pageContent .js-accordeon').accordeon();
})(jQuery);
