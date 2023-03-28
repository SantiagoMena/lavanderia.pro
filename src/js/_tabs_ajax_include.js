(function($){
	var $window = $(window);
	(function(){
		var objTabsAjax = $('#tt-pageContent .tt-ajax-tabs');
		if(!objTabsAjax.length) return;

		$window.resize(debouncer(function(e){
			switchAjaxTabs();
		}));
		$window.on('load', function(){
			switchAjaxTabs();
		});

		function switchAjaxTabs(){
			setTimeout(function(){
				$('#tt-pageContent .tt-ajax-tabs').each(function(){
					$(this).removeAttr("style");
					var value =  $(this).innerHeight();
					$(this).css({
						'height': value
					});
				});
			}, 350);
		};
	}());

	//tabs init carusel
	$('a[data-toggle="tab"]').length && $('body').on('shown.bs.tab', 'a[data-toggle="tab"]', function (e) {

		// switch animation
		var tttabsLayout = $(this).closest('.tt-ajax-tabs').find('.tab-content');
		if (tttabsLayout.length) {
			tttabsLayout.fadeTo(0,0);
			setTimeout(function(){
				tttabsLayout.fadeTo(170,1);
			}, 350);
		};

		var srcInclude = $(this).data("ajax-include") || "false",
			idInclude = $(this).attr("href") || "false";

		idInclude = idInclude.replace(/#/g, '');

		if(srcInclude !== "false" && !idInclude !== "false" && !$(this).hasClass('include')){
			$(this).addClass('include');
			$.ajax({
				url: srcInclude,
				success: function(data) {
					var $item = $(data),
						$this = $("#" + idInclude);

					$this.append($item);
					$('#tt-pageContent .js-accordeon:not(.init-accordeon)').accordeon();

					// new LazyLoad();
					var objAjax = $this.closest('.tt-ajax-tabs'),
						objAjaxValueOld = objAjax.innerHeight();

					setTimeout(function(){
						objAjax.removeAttr("style");
						var objAjaxValue =  objAjax.innerHeight();
						if(objAjaxValue < objAjaxValueOld){
							objAjax.css({
								'height': objAjaxValue
							});
						};
					}, 1000);
				}
			});
		};
		return false;
	});
})(jQuery);
